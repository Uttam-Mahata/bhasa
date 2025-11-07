package compiler

import (
	"bhasa/ast"
	"bhasa/code"
	"bhasa/lexer"
	"bhasa/object"
	"bhasa/parser"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Compiler compiles AST to bytecode
type Compiler struct {
	constants    []object.Object
	symbolTable  *SymbolTable
	scopes       []CompilationScope
	scopeIndex   int
	loopStack    []LoopContext       // track nested loops for break/continue
	moduleCache  map[string]bool     // track loaded modules to prevent circular imports
	moduleLoader ModuleLoader        // function to load module files
}

// LoopContext tracks loop start and break positions
type LoopContext struct {
	loopStart      int
	breakPositions []int
	contPositions  []int
}

// ModuleLoader is a function type for loading module source code
type ModuleLoader func(path string) (string, error)

// CompilationScope tracks instructions and jump positions
type CompilationScope struct {
	instructions        code.Instructions
	lastInstruction     EmittedInstruction
	previousInstruction EmittedInstruction
}

// EmittedInstruction tracks an emitted instruction
type EmittedInstruction struct {
	Opcode   code.Opcode
	Position int
}

// Bytecode represents compiled bytecode
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

// New creates a new Compiler
func New() *Compiler {
	mainScope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}

	symbolTable := NewSymbolTable()

	// Register built-in functions
	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	return &Compiler{
		constants:    []object.Object{},
		symbolTable:  symbolTable,
		scopes:       []CompilationScope{mainScope},
		scopeIndex:   0,
		moduleCache:  make(map[string]bool),
		moduleLoader: DefaultModuleLoader,
	}
}

// NewWithState creates a compiler with existing state
func NewWithState(s *SymbolTable, constants []object.Object) *Compiler {
	compiler := New()
	compiler.symbolTable = s
	compiler.constants = constants
	return compiler
}

// Compile compiles an AST node
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {

	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(code.OpPop)

	case *ast.InfixExpression:
		if node.Operator == "<" {
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}
			err = c.Compile(node.Left)
			if err != nil {
				return err
			}
			c.emit(code.OpGreaterThan)
			return nil
		}

		if node.Operator == "<=" {
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}
			err = c.Compile(node.Left)
			if err != nil {
				return err
			}
			c.emit(code.OpGreaterThanEqual)
			return nil
		}

		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(code.OpAdd)
		case "-":
			c.emit(code.OpSub)
		case "*":
			c.emit(code.OpMul)
		case "/":
			c.emit(code.OpDiv)
		case "%":
			c.emit(code.OpMod)
		case ">":
			c.emit(code.OpGreaterThan)
		case ">=":
			c.emit(code.OpGreaterThanEqual)
		case "==":
			c.emit(code.OpEqual)
		case "!=":
			c.emit(code.OpNotEqual)
		case "&&":
			c.emit(code.OpAnd)
		case "||":
			c.emit(code.OpOr)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "!":
			c.emit(code.OpBang)
		case "-":
			c.emit(code.OpMinus)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.IfExpression:
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)

		err = c.Compile(node.Consequence)
		if err != nil {
			return err
		}

		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}

		jumpPos := c.emit(code.OpJump, 9999)

		afterConsequencePos := len(c.currentInstructions())
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)

		if node.Alternative == nil {
			c.emit(code.OpNull)
		} else {
			err := c.Compile(node.Alternative)
			if err != nil {
				return err
			}

			if c.lastInstructionIs(code.OpPop) {
				c.removeLastPop()
			}
		}

		afterAlternativePos := len(c.currentInstructions())
		c.changeOperand(jumpPos, afterAlternativePos)

	case *ast.BlockStatement:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}

	case *ast.LetStatement:
		symbol := c.symbolTable.Define(node.Name.Value)
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}

		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index)
		} else {
			c.emit(code.OpSetLocal, symbol.Index)
		}

	case *ast.AssignmentStatement:
		symbol, ok := c.symbolTable.Resolve(node.Name.Value)
		if !ok {
			return fmt.Errorf("undefined variable %s", node.Name.Value)
		}

		err := c.Compile(node.Value)
		if err != nil {
			return err
		}

		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index)
		} else {
			c.emit(code.OpSetLocal, symbol.Index)
		}

	case *ast.WhileStatement:
		loopStart := len(c.currentInstructions())

		// Push loop context for break/continue
		loopCtx := LoopContext{loopStart: loopStart}
		c.loopStack = append(c.loopStack, loopCtx)

		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)

		err = c.Compile(node.Body)
		if err != nil {
			return err
		}

		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}

		c.emit(code.OpJump, loopStart)

		afterLoopPos := len(c.currentInstructions())
		c.changeOperand(jumpNotTruthyPos, afterLoopPos)

		// Patch all break statements
		ctx := c.loopStack[len(c.loopStack)-1]
		for _, pos := range ctx.breakPositions {
			c.changeOperand(pos, afterLoopPos)
		}

		// Patch all continue statements (go back to loop start)
		for _, pos := range ctx.contPositions {
			c.changeOperand(pos, loopStart)
		}

		// Pop loop context
		c.loopStack = c.loopStack[:len(c.loopStack)-1]

		c.emit(code.OpNull)

	case *ast.ForStatement:
		// Compile initialization
		if node.Init != nil {
			err := c.Compile(node.Init)
			if err != nil {
				return err
			}
			// Pop the result if it's an expression statement
			if _, ok := node.Init.(*ast.ExpressionStatement); ok {
				c.emit(code.OpPop)
			}
		}

		loopStart := len(c.currentInstructions())

		// Push loop context
		loopCtx := LoopContext{loopStart: loopStart}
		c.loopStack = append(c.loopStack, loopCtx)

		// Compile condition
		var jumpNotTruthyPos int
		if node.Condition != nil {
			err := c.Compile(node.Condition)
			if err != nil {
				return err
			}
			jumpNotTruthyPos = c.emit(code.OpJumpNotTruthy, 9999)
		}

		// Compile body
		err := c.Compile(node.Body)
		if err != nil {
			return err
		}

		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}

		// Continue statements jump here (before increment)
		continueTarget := len(c.currentInstructions())

		// Compile increment
		if node.Increment != nil {
			err := c.Compile(node.Increment)
			if err != nil {
				return err
			}
			// Pop the result
			if _, ok := node.Increment.(*ast.ExpressionStatement); ok {
				c.emit(code.OpPop)
			}
		}

		// Jump back to condition check
		c.emit(code.OpJump, loopStart)

		afterLoopPos := len(c.currentInstructions())

		// Patch condition jump
		if node.Condition != nil {
			c.changeOperand(jumpNotTruthyPos, afterLoopPos)
		}

		// Patch all break statements
		ctx := c.loopStack[len(c.loopStack)-1]
		for _, pos := range ctx.breakPositions {
			c.changeOperand(pos, afterLoopPos)
		}

		// Patch all continue statements (go to increment)
		for _, pos := range ctx.contPositions {
			c.changeOperand(pos, continueTarget)
		}

		// Pop loop context
		c.loopStack = c.loopStack[:len(c.loopStack)-1]

		c.emit(code.OpNull)

	case *ast.BreakStatement:
		if len(c.loopStack) == 0 {
			return fmt.Errorf("break statement outside loop")
		}
		// Emit a jump that will be patched later
		pos := c.emit(code.OpJump, 9999)
		// Record this position in the current loop context
		ctx := &c.loopStack[len(c.loopStack)-1]
		ctx.breakPositions = append(ctx.breakPositions, pos)

	case *ast.ContinueStatement:
		if len(c.loopStack) == 0 {
			return fmt.Errorf("continue statement outside loop")
		}
		// Emit a jump that will be patched later
		pos := c.emit(code.OpJump, 9999)
		// Record this position in the current loop context
		ctx := &c.loopStack[len(c.loopStack)-1]
		ctx.contPositions = append(ctx.contPositions, pos)

	case *ast.ImportStatement:
		// Evaluate the path expression (should be a string literal)
		if pathLit, ok := node.Path.(*ast.StringLiteral); ok {
			modulePath := pathLit.Value
			// Load and compile the module
			if err := c.LoadAndCompileModule(modulePath); err != nil {
				return fmt.Errorf("error importing module: %v", err)
			}
		} else {
			return fmt.Errorf("import path must be a string literal")
		}

	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined variable %s", node.Value)
		}

		c.loadSymbol(symbol)

	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))

	case *ast.StringLiteral:
		str := &object.String{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(str))

	case *ast.Boolean:
		if node.Value {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}

	case *ast.ArrayLiteral:
		for _, el := range node.Elements {
			err := c.Compile(el)
			if err != nil {
				return err
			}
		}

		c.emit(code.OpArray, len(node.Elements))

	case *ast.HashLiteral:
		keys := []ast.Expression{}
		for k := range node.Pairs {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})

		for _, k := range keys {
			err := c.Compile(k)
			if err != nil {
				return err
			}
			err = c.Compile(node.Pairs[k])
			if err != nil {
				return err
			}
		}

		c.emit(code.OpHash, len(node.Pairs)*2)

	case *ast.IndexExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Index)
		if err != nil {
			return err
		}

		c.emit(code.OpIndex)

	case *ast.FunctionLiteral:
		c.enterScope()

		for _, p := range node.Parameters {
			c.symbolTable.Define(p.Value)
		}

		err := c.Compile(node.Body)
		if err != nil {
			return err
		}

		if c.lastInstructionIs(code.OpPop) {
			c.replaceLastPopWithReturn()
		}
		if !c.lastInstructionIs(code.OpReturnValue) {
			c.emit(code.OpReturn)
		}

		freeSymbols := c.symbolTable.FreeSymbols
		numLocals := c.symbolTable.numDefinitions
		instructions := c.leaveScope()

		for _, s := range freeSymbols {
			c.loadSymbol(s)
		}

		compiledFn := &object.CompiledFunction{
			Instructions:  instructions,
			NumLocals:     numLocals,
			NumParameters: len(node.Parameters),
		}
		fnIndex := c.addConstant(compiledFn)
		c.emit(code.OpClosure, fnIndex, len(freeSymbols))

	case *ast.ReturnStatement:
		err := c.Compile(node.ReturnValue)
		if err != nil {
			return err
		}

		c.emit(code.OpReturnValue)

	case *ast.CallExpression:
		err := c.Compile(node.Function)
		if err != nil {
			return err
		}

		for _, a := range node.Arguments {
			err := c.Compile(a)
			if err != nil {
				return err
			}
		}

		c.emit(code.OpCall, len(node.Arguments))

	case *ast.ClassStatement:
		// Compile methods as closures
		methods := make(map[string]*object.Closure)
		for name, method := range node.Methods {
			// Compile each method
			err := c.Compile(method)
			if err != nil {
				return err
			}
			
			// For simplicity, we'll store the class definition as a constant
			// In a real implementation, methods would be compiled and stored properly
			_ = name // We'll handle this when implementing VM support
		}
		
		// Create class object and add as constant
		classObj := &object.Class{
			Name:    node.Name.Value,
			Methods: methods,
		}
		classConst := c.addConstant(classObj)
		
		// Define the class as a global variable
		symbol := c.symbolTable.Define(node.Name.Value)
		c.emit(code.OpConstant, classConst)
		
		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index)
		} else {
			c.emit(code.OpSetLocal, symbol.Index)
		}

	case *ast.NewExpression:
		// Compile the class identifier
		err := c.Compile(node.Class)
		if err != nil {
			return err
		}

		// Compile constructor arguments
		for _, arg := range node.Arguments {
			err := c.Compile(arg)
			if err != nil {
				return err
			}
		}

		// Emit new instance instruction
		c.emit(code.OpNewInstance, len(node.Arguments))

	case *ast.MemberAccessExpression:
		// Compile the object
		err := c.Compile(node.Object)
		if err != nil {
			return err
		}

		// Push property name as constant
		propertyName := &object.String{Value: node.Member.Value}
		propIndex := c.addConstant(propertyName)
		c.emit(code.OpConstant, propIndex)

		// Emit get property instruction
		c.emit(code.OpGetProperty)

	case *ast.ThisExpression:
		// Emit instruction to get current instance
		c.emit(code.OpThis)
	}

	return nil
}

// Bytecode returns the compiled bytecode
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.currentInstructions(),
		Constants:    c.constants,
	}
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	c.setLastInstruction(op, pos)

	return pos
}

func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.currentInstructions())
	updatedInstructions := append(c.currentInstructions(), ins...)

	c.scopes[c.scopeIndex].instructions = updatedInstructions

	return posNewInstruction
}

func (c *Compiler) setLastInstruction(op code.Opcode, pos int) {
	previous := c.scopes[c.scopeIndex].lastInstruction
	last := EmittedInstruction{Opcode: op, Position: pos}

	c.scopes[c.scopeIndex].previousInstruction = previous
	c.scopes[c.scopeIndex].lastInstruction = last
}

func (c *Compiler) lastInstructionIs(op code.Opcode) bool {
	if len(c.currentInstructions()) == 0 {
		return false
	}

	return c.scopes[c.scopeIndex].lastInstruction.Opcode == op
}

func (c *Compiler) removeLastPop() {
	last := c.scopes[c.scopeIndex].lastInstruction
	previous := c.scopes[c.scopeIndex].previousInstruction

	old := c.currentInstructions()
	new := old[:last.Position]

	c.scopes[c.scopeIndex].instructions = new
	c.scopes[c.scopeIndex].lastInstruction = previous
}

func (c *Compiler) replaceInstruction(pos int, newInstruction []byte) {
	ins := c.currentInstructions()

	for i := 0; i < len(newInstruction); i++ {
		ins[pos+i] = newInstruction[i]
	}
}

func (c *Compiler) changeOperand(opPos int, operand int) {
	op := code.Opcode(c.currentInstructions()[opPos])
	newInstruction := code.Make(op, operand)

	c.replaceInstruction(opPos, newInstruction)
}

func (c *Compiler) currentInstructions() code.Instructions {
	return c.scopes[c.scopeIndex].instructions
}

func (c *Compiler) enterScope() {
	scope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.scopeIndex++
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable)
}

func (c *Compiler) leaveScope() code.Instructions {
	instructions := c.currentInstructions()

	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	c.symbolTable = c.symbolTable.Outer

	return instructions
}

func (c *Compiler) replaceLastPopWithReturn() {
	lastPos := c.scopes[c.scopeIndex].lastInstruction.Position
	c.replaceInstruction(lastPos, code.Make(code.OpReturnValue))

	c.scopes[c.scopeIndex].lastInstruction.Opcode = code.OpReturnValue
}

func (c *Compiler) loadSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.emit(code.OpGetGlobal, s.Index)
	case LocalScope:
		c.emit(code.OpGetLocal, s.Index)
	case BuiltinScope:
		c.emit(code.OpGetBuiltin, s.Index)
	case FreeScope:
		c.emit(code.OpGetFree, s.Index)
	case FunctionScope:
		c.emit(code.OpCurrentClosure)
	}
}

// DefaultModuleLoader loads modules from the filesystem
// Supports both .ভাষা (Bengali) and .bhasa extensions
func DefaultModuleLoader(modulePath string) (string, error) {
	// Try different file extensions
	extensions := []string{".ভাষা", ".bhasa"}
	
	// Try different search paths
	searchPaths := []string{
		modulePath,            // Direct path
		"modules/" + modulePath, // modules directory
	}
	
	var fullPath string
	
	for _, basePath := range searchPaths {
		for _, ext := range extensions {
			// Try with extension if not already present
			testPath := basePath
			if !strings.HasSuffix(basePath, ext) {
				testPath = basePath + ext
			}
			
			// Check if file exists
			if _, statErr := os.Stat(testPath); statErr == nil {
				fullPath = testPath
				break
			}
		}
		if fullPath != "" {
			break
		}
	}
	
	if fullPath == "" {
		return "", fmt.Errorf("module not found: %s (tried .ভাষা and .bhasa extensions in current dir and modules/ dir)", modulePath)
	}
	
	// Read the file
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("error reading module %s: %v", fullPath, err)
	}
	
	return string(content), nil
}

// LoadAndCompileModule loads a module file, parses it, and compiles it
func (c *Compiler) LoadAndCompileModule(modulePath string) error {
	// Load module source code first (this will search in multiple locations)
	source, err := c.moduleLoader(modulePath)
	if err != nil {
		return err
	}
	
	// Use module path as cache key (simple approach)
	// Check if module is already loaded (circular dependency detection)
	if c.moduleCache[modulePath] {
		return nil // Already loaded, skip
	}
	
	// Mark as being loaded
	c.moduleCache[modulePath] = true
	
	// Parse the module
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		return fmt.Errorf("parser errors in module %s: %v", modulePath, p.Errors())
	}
	
	// Compile the module
	return c.Compile(program)
}

