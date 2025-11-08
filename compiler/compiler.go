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
		case "&":
			c.emit(code.OpBitAnd)
		case "|":
			c.emit(code.OpBitOr)
		case "^":
			c.emit(code.OpBitXor)
		case "<<":
			c.emit(code.OpLeftShift)
		case ">>":
			c.emit(code.OpRightShift)
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
		case "~":
			c.emit(code.OpBitNot)
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

		// If-else is an expression that MUST leave a value. If consequence doesn't
		// end with OpPop (meaning it's a statement, not expression), emit OpNull.
		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		} else {
			c.emit(code.OpNull)
		}

		jumpPos := c.emit(code.OpJump, 9999)

		afterConsequencePos := len(c.currentInstructions())
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)

		// Same for alternative - must leave a value
		if node.Alternative == nil {
			c.emit(code.OpNull)
		} else {
			err := c.Compile(node.Alternative)
			if err != nil {
				return err
			}

			if c.lastInstructionIs(code.OpPop) {
				c.removeLastPop()
			} else {
				c.emit(code.OpNull)
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
		// Define symbol with type annotation if present
		var symbol Symbol
		if node.TypeAnnot != nil {
			symbol = c.symbolTable.DefineWithType(node.Name.Value, node.TypeAnnot)
		} else {
			symbol = c.symbolTable.Define(node.Name.Value)
		}

		// If value is an EnumDefinition, set its name from the binding
		if enumDef, ok := node.Value.(*ast.EnumDefinition); ok {
			enumDef.Name = node.Name
		}

		err := c.Compile(node.Value)
		if err != nil {
			return err
		}

		// If type annotation is present, emit type check
		if node.TypeAnnot != nil {
			typeConstIndex := c.addConstant(&object.String{Value: node.TypeAnnot.String()})
			c.emit(code.OpAssertType, typeConstIndex)
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

	case *ast.MemberAssignmentStatement:
		// Compile the object expression
		err := c.Compile(node.Object)
		if err != nil {
			return err
		}

		// Push the field name as a constant
		nameConstant := c.addConstant(&object.String{Value: node.Member.Value})
		c.emit(code.OpConstant, nameConstant)

		// Compile the value to assign
		err = c.Compile(node.Value)
		if err != nil {
			return err
		}

		// Emit instruction to set struct field
		c.emit(code.OpSetStructField)

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

	case *ast.StructLiteral:
		// Sort field names for deterministic compilation
		fieldNames := make([]string, 0, len(node.Fields))
		for name := range node.Fields {
			fieldNames = append(fieldNames, name)
		}
		sort.Strings(fieldNames)

		// Compile field name-value pairs
		for _, name := range fieldNames {
			// Push field name as constant
			nameConstant := c.addConstant(&object.String{Value: name})
			c.emit(code.OpConstant, nameConstant)

			// Push field value
			err := c.Compile(node.Fields[name])
			if err != nil {
				return err
			}
		}

		// Create struct with number of fields
		c.emit(code.OpStruct, len(node.Fields)*2)

	case *ast.EnumDefinition:
		// Create EnumType object
		enumName := ""
		if node.Name != nil {
			enumName = node.Name.Value
		}

		// Build variants map
		variants := make(map[string]int)
		value := 0
		for _, variant := range node.Variants {
			if variant.Value != nil {
				value = *variant.Value
			}
			variants[variant.Name] = value
			value++
		}

		// Create EnumType object and add as constant
		enumType := &object.EnumType{
			Name:     enumName,
			Variants: variants,
		}
		enumTypeIndex := c.addConstant(enumType)
		c.emit(code.OpConstant, enumTypeIndex)

	case *ast.MemberAccessExpression:
		// Compile the object expression
		err := c.Compile(node.Object)
		if err != nil {
			return err
		}

		// Push the field name as a constant
		nameConstant := c.addConstant(&object.String{Value: node.Member.Value})
		c.emit(code.OpConstant, nameConstant)

		// Emit instruction to get struct field
		c.emit(code.OpGetStructField)

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

	case *ast.TypeCastExpression:
		// Compile the expression to cast
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}

		// Emit type cast opcode with target type
		typeConstIndex := c.addConstant(&object.String{Value: node.TargetType.String()})
		c.emit(code.OpTypeCast, typeConstIndex)

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

	// ========== OOP Compilation ==========

	case *ast.ClassDefinition:
		return c.compileClassDefinition(node)

	case *ast.InterfaceDefinition:
		return c.compileInterfaceDefinition(node)

	case *ast.NewExpression:
		return c.compileNewExpression(node)

	case *ast.ThisExpression:
		c.emit(code.OpGetThis)

	case *ast.SuperExpression:
		c.emit(code.OpGetSuper)

	case *ast.MethodCallExpression:
		return c.compileMethodCall(node)
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


// ============================================================================
// OOP Compilation Functions
// ============================================================================

// compileClassDefinition compiles a class definition
func (c *Compiler) compileClassDefinition(node *ast.ClassDefinition) error {
	// Create a new Class object
	class := &object.Class{
		Name:         node.Name.Value,
		SuperClass:   nil,
		Interfaces:   []*object.Interface{},
		Fields:       make(map[string]string),
		Methods:      make(map[string]*object.Method),
		Constructor:  nil,
		StaticFields: make(map[string]object.Object),
		IsAbstract:   node.IsAbstract,
		IsFinal:      node.IsFinal,
		FieldAccess:  make(map[string]string),
		FieldOrder:   []string{},
	}

	// Process fields
	for _, field := range node.Fields {
		class.Fields[field.Name] = field.TypeAnnot.String()
		class.FieldAccess[field.Name] = string(field.Access)
		class.FieldOrder = append(class.FieldOrder, field.Name)
	}

	// Compile constructor
	if len(node.Constructors) > 0 {
		// Use the first constructor (in simple implementation)
		constructor := node.Constructors[0]
		
		// Compile constructor as a function
		c.enterScope()
		
		// Define 'this' parameter
		c.symbolTable.Define("এই")
		
		// Define constructor parameters
		for _, param := range constructor.Parameters {
			c.symbolTable.Define(param.Value)
		}
		
		// Compile constructor body
		if constructor.Body != nil {
			if err := c.Compile(constructor.Body); err != nil {
				return err
			}
		}
		
		// Constructor should always return 'this' (এই)
		// If there's no explicit return, add implicit return of 'this'
		if !c.lastInstructionIs(code.OpReturnValue) {
			// Load 'এই' (this) parameter - it's always the first parameter (index 0)
			c.emit(code.OpGetLocal, 0)
			c.emit(code.OpReturnValue)
		}
		
		freeSymbols := c.symbolTable.FreeSymbols
		numLocals := c.symbolTable.numDefinitions
		instructions := c.leaveScope()
		
		compiledFn := &object.CompiledFunction{
			Instructions:  instructions,
			NumLocals:     numLocals,
			NumParameters: len(constructor.Parameters) + 1, // +1 for 'this'
		}
		
		fnIndex := c.addConstant(compiledFn)
		
		for _, s := range freeSymbols {
			c.loadSymbol(s)
		}
		
		class.Constructor = &object.Closure{
			Fn:   compiledFn,
			Free: make([]object.Object, len(freeSymbols)),
		}
		
		c.emit(code.OpClosure, fnIndex, len(freeSymbols))
		c.emit(code.OpDefineConstructor, fnIndex)
	}
	
	// Compile methods
	for _, method := range node.Methods {
		// Skip abstract methods (no body)
		if method.IsAbstract {
			continue
		}
		
		// Compile method as a function
		c.enterScope()
		
		// Define 'this' parameter
		c.symbolTable.Define("এই")
		
		// Define method parameters
		for _, param := range method.Parameters {
			c.symbolTable.Define(param.Value)
		}
		
		// Compile method body
		if method.Body != nil {
			if err := c.Compile(method.Body); err != nil {
				return err
			}
		}
		
		// Add implicit return if needed
		if !c.lastInstructionIs(code.OpReturnValue) {
			c.emit(code.OpNull)
			c.emit(code.OpReturnValue)
		}
		
		freeSymbols := c.symbolTable.FreeSymbols
		numLocals := c.symbolTable.numDefinitions
		instructions := c.leaveScope()
		
		compiledFn := &object.CompiledFunction{
			Instructions:  instructions,
			NumLocals:     numLocals,
			NumParameters: len(method.Parameters) + 1, // +1 for 'this'
		}
		
		fnIndex := c.addConstant(compiledFn)
		
		for _, s := range freeSymbols {
			c.loadSymbol(s)
		}
		
		closure := &object.Closure{
			Fn:   compiledFn,
			Free: make([]object.Object, len(freeSymbols)),
		}
		
		// Create method object
		methodObj := &object.Method{
			Name:       method.Name.Value,
			Access:     string(method.Access),
			IsStatic:   method.IsStatic,
			IsFinal:    method.IsFinal,
			IsAbstract: method.IsAbstract,
			Closure:    closure,
		}
		
		class.Methods[method.Name.Value] = methodObj
		
		// Emit method definition
		methodNameIndex := c.addConstant(&object.String{Value: method.Name.Value})
		c.emit(code.OpClosure, fnIndex, len(freeSymbols))
		c.emit(code.OpDefineMethod, methodNameIndex)
	}
	
	// Add class to constants
	classIndex := c.addConstant(class)
	c.emit(code.OpClass, classIndex)
	
	// Define class in symbol table
	symbol := c.symbolTable.Define(node.Name.Value)
	if symbol.Scope == GlobalScope {
		c.emit(code.OpSetGlobal, symbol.Index)
	} else {
		c.emit(code.OpSetLocal, symbol.Index)
	}
	
	return nil
}

// compileInterfaceDefinition compiles an interface definition
func (c *Compiler) compileInterfaceDefinition(node *ast.InterfaceDefinition) error {
	// Create interface object
	iface := &object.Interface{
		Name:             node.Name.Value,
		MethodSignatures: make(map[string][]string),
	}
	
	// Process method signatures
	for _, method := range node.Methods {
		paramTypes := []string{}
		for _, paramType := range method.ParameterTypes {
			paramTypes = append(paramTypes, paramType.String())
		}
		iface.MethodSignatures[method.Name.Value] = paramTypes
	}
	
	// Add interface to constants
	ifaceIndex := c.addConstant(iface)
	c.emit(code.OpInterface, ifaceIndex)
	
	// Define interface in symbol table
	symbol := c.symbolTable.Define(node.Name.Value)
	if symbol.Scope == GlobalScope {
		c.emit(code.OpSetGlobal, symbol.Index)
	} else {
		c.emit(code.OpSetLocal, symbol.Index)
	}
	
	return nil
}

// compileNewExpression compiles a new instance expression
func (c *Compiler) compileNewExpression(node *ast.NewExpression) error {
	// Compile constructor arguments first
	for _, arg := range node.Arguments {
		err := c.Compile(arg)
		if err != nil {
			return err
		}
	}

	// Load the class last (so it's on top of stack)
	err := c.Compile(node.ClassName)
	if err != nil {
		return err
	}

	// Emit OpNewInstance with argument count
	// Stack layout before OpNewInstance: [arg1, arg2, ..., argN, class]
	c.emit(code.OpNewInstance, len(node.Arguments))

	return nil
}

// compileMethodCall compiles a method call
func (c *Compiler) compileMethodCall(node *ast.MethodCallExpression) error {
	// Compile the object
	err := c.Compile(node.Object)
	if err != nil {
		return err
	}
	
	// Push method name as constant
	methodNameIndex := c.addConstant(&object.String{Value: node.MethodName.Value})
	c.emit(code.OpConstant, methodNameIndex)
	
	// Compile arguments
	for _, arg := range node.Arguments {
		err := c.Compile(arg)
		if err != nil {
			return err
		}
	}
	
	// Emit method call
	c.emit(code.OpCallMethod, len(node.Arguments))
	
	return nil
}
