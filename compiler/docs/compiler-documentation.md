# Bhasa Compiler Documentation

## Table of Contents

1. [Overview](#overview)
2. [Core Types](#core-types)
3. [Compiler Lifecycle](#compiler-lifecycle)
4. [Statement Compilation](#statement-compilation)
5. [Expression Compilation](#expression-compilation)
6. [Scope Management](#scope-management)
7. [Control Flow](#control-flow)
8. [Function Compilation](#function-compilation)
9. [OOP Compilation](#oop-compilation)
10. [Module System](#module-system)
11. [Helper Methods](#helper-methods)
12. [Optimization Techniques](#optimization-techniques)

---

## Overview

The Bhasa compiler translates Abstract Syntax Tree (AST) nodes into bytecode instructions. It's a **single-pass compiler** that generates stack-based bytecode for the Bhasa Virtual Machine.

### Key Responsibilities

- **AST → Bytecode**: Convert high-level AST to low-level bytecode
- **Symbol Management**: Track variables, functions, and their scopes
- **Constant Pooling**: Store and reuse literal values
- **Scope Tracking**: Manage nested scopes (global, local, closures)
- **Jump Patching**: Fix forward references in control flow
- **Closure Handling**: Capture free variables from outer scopes

---

## Core Types

### Compiler

```go
type Compiler struct {
    constants    []object.Object     // Constant pool for literals
    symbolTable  *SymbolTable        // Variable tracking
    scopes       []CompilationScope  // Stack of compilation scopes
    scopeIndex   int                 // Current scope index
    loopStack    []LoopContext       // Nested loop tracking
    moduleCache  map[string]bool     // Circular import prevention
    moduleLoader ModuleLoader        // Module loading function
}
```

**Fields**:

- **constants**: Pool of constant values (numbers, strings, functions, classes)
  - Indexed by 16-bit integers (max 65,535 constants)
  - Shared across all scopes
  
- **symbolTable**: Current symbol table (changes with scope)
  - Tracks variable names and their scopes
  - Forms a linked list for nested scopes
  
- **scopes**: Stack of compilation scopes
  - Each function/block gets its own scope
  - Scopes contain bytecode instructions
  
- **scopeIndex**: Index into scopes array
  - Points to current scope
  - Incremented on `enterScope()`, decremented on `leaveScope()`
  
- **loopStack**: Stack of loop contexts
  - Tracks break/continue positions
  - Popped when exiting loop
  
- **moduleCache**: Set of loaded module paths
  - Prevents circular imports
  - Cleared between compilations

- **moduleLoader**: Function to load module source
  - Default: loads from filesystem
  - Can be customized for testing

---

### CompilationScope

```go
type CompilationScope struct {
    instructions        code.Instructions  // Bytecode buffer
    lastInstruction     EmittedInstruction // Last emitted instruction
    previousInstruction EmittedInstruction // Previous instruction
}
```

**Purpose**: Represents one scope's compilation state

**Fields**:
- **instructions**: Bytecode generated in this scope
- **lastInstruction**: Used for optimization (e.g., pop removal)
- **previousInstruction**: Used when tracking instruction sequences

---

### EmittedInstruction

```go
type EmittedInstruction struct {
    Opcode   code.Opcode  // The opcode
    Position int          // Position in bytecode
}
```

**Purpose**: Tracks an emitted instruction for optimization

**Usage**:
- Removing unnecessary `OpPop` instructions
- Replacing `OpPop` with `OpReturnValue`
- Jump position calculations

---

### LoopContext

```go
type LoopContext struct {
    loopStart      int    // Position of loop start
    breakPositions []int  // Positions of break statements
    contPositions  []int  // Positions of continue statements
}
```

**Purpose**: Tracks loop information for break/continue

**Lifecycle**:
1. Pushed when entering loop
2. Break/continue positions added during body compilation
3. Positions patched after loop
4. Popped when exiting loop

---

### Bytecode

```go
type Bytecode struct {
    Instructions code.Instructions  // Compiled bytecode
    Constants    []object.Object    // Constant pool
}
```

**Purpose**: Final compilation output

**Usage**:
```go
bytecode := compiler.Bytecode()
vm := vm.New(bytecode)
vm.Run()
```

---

## Compiler Lifecycle

### 1. Creation

```go
func New() *Compiler
```

**Creates a new compiler with**:
- Empty constant pool
- New global symbol table
- One scope (global scope)
- Empty loop stack
- Empty module cache
- Default module loader

**Builtin Registration**:
- All builtin functions (len, print, etc.) are registered in symbol table
- Given BUILTIN scope
- Indexed by their position in builtins array

**Example**:
```go
compiler := compiler.New()
```

---

### 2. Compilation

```go
func (c *Compiler) Compile(node ast.Node) error
```

**Main compilation entry point**:
- Called recursively for all AST nodes
- Type switches on node type
- Emits bytecode instructions
- Updates symbol table

**Example**:
```go
program := parser.ParseProgram()
err := compiler.Compile(program)
if err != nil {
    // Handle compilation error
}
```

---

### 3. Bytecode Extraction

```go
func (c *Compiler) Bytecode() *Bytecode
```

**Extracts compiled bytecode**:
- Returns instructions and constants
- Called after compilation completes
- Can be called multiple times (idempotent)

**Example**:
```go
bytecode := compiler.Bytecode()
fmt.Println(bytecode.Instructions.String())
```

---

## Statement Compilation

### Program

```go
case *ast.Program:
    for _, s := range node.Statements {
        err := c.Compile(s)
        if err != nil {
            return err
        }
    }
```

**Process**: Compile each statement in sequence

---

### ExpressionStatement

```go
case *ast.ExpressionStatement:
    err := c.Compile(node.Expression)
    if err != nil {
        return err
    }
    c.emit(code.OpPop)
```

**Purpose**: Compile expression and discard result

**Why OpPop?**: Expression results must be removed from stack

**Example**:
```bhasa
5 + 3;  // Expression statement
```

**Bytecode**:
```
OpConstant 0
OpConstant 1
OpAdd
OpPop        // ← Discard result
```

---

### LetStatement

```go
case *ast.LetStatement:
    // 1. Define symbol
    var symbol Symbol
    if node.TypeAnnot != nil {
        symbol = c.symbolTable.DefineWithType(node.Name.Value, node.TypeAnnot)
    } else {
        symbol = c.symbolTable.Define(node.Name.Value)
    }
    
    // 2. Compile value
    err := c.Compile(node.Value)
    if err != nil {
        return err
    }
    
    // 3. Type check if annotated
    if node.TypeAnnot != nil {
        typeConstIndex := c.addConstant(&object.String{Value: node.TypeAnnot.String()})
        c.emit(code.OpAssertType, typeConstIndex)
    }
    
    // 4. Store variable
    if symbol.Scope == GlobalScope {
        c.emit(code.OpSetGlobal, symbol.Index)
    } else {
        c.emit(code.OpSetLocal, symbol.Index)
    }
```

**Steps**:
1. **Define symbol** in symbol table (get index and scope)
2. **Compile value** expression (leaves value on stack)
3. **Type assertion** if type annotation present (runtime check)
4. **Emit store** instruction (global or local based on scope)

**Example**:
```bhasa
ধরি x: পূর্ণসংখ্যা = 5;
```

**Bytecode**:
```
OpConstant 0         // Push 5
OpAssertType 1       // Check type (পূর্ণসংখ্যা)
OpSetGlobal 0        // Store in x
```

---

### AssignmentStatement

```go
case *ast.AssignmentStatement:
    // 1. Resolve symbol
    symbol, ok := c.symbolTable.Resolve(node.Name.Value)
    if !ok {
        return fmt.Errorf("undefined variable %s", node.Name.Value)
    }
    
    // 2. Compile value
    err := c.Compile(node.Value)
    if err != nil {
        return err
    }
    
    // 3. Store variable
    if symbol.Scope == GlobalScope {
        c.emit(code.OpSetGlobal, symbol.Index)
    } else {
        c.emit(code.OpSetLocal, symbol.Index)
    }
```

**Steps**:
1. **Resolve symbol** (must already exist)
2. **Compile value** expression
3. **Emit store** instruction

**Difference from LetStatement**: Variable must already be defined

---

### MemberAssignmentStatement

```go
case *ast.MemberAssignmentStatement:
    // Compile object
    err := c.Compile(node.Object)
    if err != nil {
        return err
    }
    
    // Push field name
    nameConstant := c.addConstant(&object.String{Value: node.Member.Value})
    c.emit(code.OpConstant, nameConstant)
    
    // Compile value
    err = c.Compile(node.Value)
    if err != nil {
        return err
    }
    
    // Set field
    c.emit(code.OpSetStructField)
```

**Purpose**: Assign to struct/class field

**Example**:
```bhasa
person.নাম = "নতুন নাম";
```

**Bytecode**:
```
OpGetGlobal 0        // Load person
OpConstant 1         // Push "নাম"
OpConstant 2         // Push "নতুন নাম"
OpSetStructField     // person.নাম = value
```

---

### ImportStatement

```go
case *ast.ImportStatement:
    if pathLit, ok := node.Path.(*ast.StringLiteral); ok {
        modulePath := pathLit.Value
        if err := c.LoadAndCompileModule(modulePath); err != nil {
            return fmt.Errorf("error importing module: %v", err)
        }
    } else {
        return fmt.Errorf("import path must be a string literal")
    }
```

**Purpose**: Load and compile external modules

**Features**:
- Circular import detection
- File extension support (.ভাষা, .bhasa)
- Module path search (current dir, modules/ dir)

**Example**:
```bhasa
অন্তর্ভুক্ত "গণিত";
```

---

### ReturnStatement

```go
case *ast.ReturnStatement:
    err := c.Compile(node.ReturnValue)
    if err != nil {
        return err
    }
    c.emit(code.OpReturnValue)
```

**Purpose**: Return value from function

**Example**:
```bhasa
ফেরত x + 5;
```

**Bytecode**:
```
OpGetLocal 0         // Load x
OpConstant 0         // Push 5
OpAdd                // Add
OpReturnValue        // Return result
```

---

## Expression Compilation

### Literals

#### IntegerLiteral

```go
case *ast.IntegerLiteral:
    integer := &object.Integer{Value: node.Value}
    c.emit(code.OpConstant, c.addConstant(integer))
```

**Process**:
1. Create Integer object
2. Add to constant pool
3. Emit OpConstant with index

---

#### StringLiteral

```go
case *ast.StringLiteral:
    str := &object.String{Value: node.Value}
    c.emit(code.OpConstant, c.addConstant(str))
```

**Same as IntegerLiteral** but with String object

---

#### Boolean

```go
case *ast.Boolean:
    if node.Value {
        c.emit(code.OpTrue)
    } else {
        c.emit(code.OpFalse)
    }
```

**Optimization**: Booleans use dedicated opcodes (no constant pool)

---

### Identifier

```go
case *ast.Identifier:
    symbol, ok := c.symbolTable.Resolve(node.Value)
    if !ok {
        return fmt.Errorf("undefined variable %s", node.Value)
    }
    c.loadSymbol(symbol)
```

**Process**:
1. Resolve symbol from symbol table
2. Load based on scope type (global, local, builtin, free)

---

### InfixExpression

```go
case *ast.InfixExpression:
    // Special handling for < and <=
    if node.Operator == "<" {
        err := c.Compile(node.Right)
        err = c.Compile(node.Left)
        c.emit(code.OpGreaterThan)
        return nil
    }
    
    // Normal operators
    err := c.Compile(node.Left)
    err = c.Compile(node.Right)
    
    switch node.Operator {
    case "+": c.emit(code.OpAdd)
    case "-": c.emit(code.OpSub)
    case "*": c.emit(code.OpMul)
    case "/": c.emit(code.OpDiv)
    // ... more operators
    }
```

**Key Feature**: `<` and `<=` are implemented as reversed `>` and `>=`

**Why?**: VM only implements `>` and `>=`, reduces opcodes

**Example**:
```bhasa
a < b  // Compiled as: b > a
```

---

### PrefixExpression

```go
case *ast.PrefixExpression:
    err := c.Compile(node.Right)
    if err != nil {
        return err
    }
    
    switch node.Operator {
    case "!": c.emit(code.OpBang)
    case "-": c.emit(code.OpMinus)
    case "~": c.emit(code.OpBitNot)
    }
```

**Example**:
```bhasa
!সত্য
```

**Bytecode**:
```
OpTrue
OpBang
```

---

### ArrayLiteral

```go
case *ast.ArrayLiteral:
    for _, el := range node.Elements {
        err := c.Compile(el)
        if err != nil {
            return err
        }
    }
    c.emit(code.OpArray, len(node.Elements))
```

**Process**:
1. Compile each element (pushes to stack)
2. Emit OpArray with element count
3. VM pops N elements and creates array

**Example**:
```bhasa
[1, 2, 3]
```

**Bytecode**:
```
OpConstant 0         // Push 1
OpConstant 1         // Push 2
OpConstant 2         // Push 3
OpArray 3            // Create array with 3 elements
```

---

### HashLiteral

```go
case *ast.HashLiteral:
    // Sort keys for deterministic compilation
    keys := []ast.Expression{}
    for k := range node.Pairs {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool {
        return keys[i].String() < keys[j].String()
    })
    
    // Compile key-value pairs
    for _, k := range keys {
        err := c.Compile(k)
        err = c.Compile(node.Pairs[k])
    }
    
    c.emit(code.OpHash, len(node.Pairs)*2)
```

**Key Feature**: Keys are sorted for **deterministic compilation**

**Why?**: Go maps have random iteration order

---

### IndexExpression

```go
case *ast.IndexExpression:
    err := c.Compile(node.Left)    // array/hash
    err = c.Compile(node.Index)    // index/key
    c.emit(code.OpIndex)
```

**Example**:
```bhasa
arr[2]
```

**Bytecode**:
```
OpGetGlobal 0        // Load arr
OpConstant 0         // Push 2
OpIndex              // Get arr[2]
```

---

### StructLiteral

```go
case *ast.StructLiteral:
    // Sort field names for deterministic compilation
    fieldNames := make([]string, 0, len(node.Fields))
    for name := range node.Fields {
        fieldNames = append(fieldNames, name)
    }
    sort.Strings(fieldNames)
    
    // Compile field name-value pairs
    for _, name := range fieldNames {
        nameConstant := c.addConstant(&object.String{Value: name})
        c.emit(code.OpConstant, nameConstant)
        
        err := c.Compile(node.Fields[name])
        if err != nil {
            return err
        }
    }
    
    c.emit(code.OpStruct, len(node.Fields)*2)
```

**Example**:
```bhasa
ব্যক্তি{নাম: "রহিম", বয়স: 30}
```

---

### MemberAccessExpression

```go
case *ast.MemberAccessExpression:
    // Compile object
    err := c.Compile(node.Object)
    
    // Push field name
    nameConstant := c.addConstant(&object.String{Value: node.Member.Value})
    c.emit(code.OpConstant, nameConstant)
    
    // Get field
    c.emit(code.OpGetStructField)
```

**Example**:
```bhasa
person.নাম
```

**Bytecode**:
```
OpGetGlobal 0        // Load person
OpConstant 1         // Push "নাম"
OpGetStructField     // Get person.নাম
```

---

### EnumDefinition

```go
case *ast.EnumDefinition:
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
    
    // Create EnumType and add as constant
    enumType := &object.EnumType{
        Name:     enumName,
        Variants: variants,
    }
    enumTypeIndex := c.addConstant(enumType)
    c.emit(code.OpConstant, enumTypeIndex)
```

**Example**:
```bhasa
ধরি দিক = গণনা {
    উত্তর,
    দক্ষিণ,
    পূর্ব,
    পশ্চিম
};
```

---

## Control Flow

### IfExpression

```go
case *ast.IfExpression:
    // 1. Compile condition
    err := c.Compile(node.Condition)
    
    // 2. Emit jump with placeholder
    jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)
    
    // 3. Compile consequence
    err = c.Compile(node.Consequence)
    
    // 4. Handle value (if-else is an expression)
    if c.lastInstructionIs(code.OpPop) {
        c.removeLastPop()
    } else {
        c.emit(code.OpNull)
    }
    
    // 5. Jump over alternative
    jumpPos := c.emit(code.OpJump, 9999)
    
    // 6. Patch first jump
    afterConsequencePos := len(c.currentInstructions())
    c.changeOperand(jumpNotTruthyPos, afterConsequencePos)
    
    // 7. Compile alternative or emit null
    if node.Alternative == nil {
        c.emit(code.OpNull)
    } else {
        err := c.Compile(node.Alternative)
        if c.lastInstructionIs(code.OpPop) {
            c.removeLastPop()
        } else {
            c.emit(code.OpNull)
        }
    }
    
    // 8. Patch second jump
    afterAlternativePos := len(c.currentInstructions())
    c.changeOperand(jumpPos, afterAlternativePos)
```

**Key Points**:
- If-else is an **expression** (must leave value on stack)
- Uses **OpPop removal** optimization
- **Jump patching** for forward references

**Example**:
```bhasa
যদি (x > 5) { 10 } নাহলে { 20 }
```

**Bytecode**:
```
OpGetGlobal 0       // Load x
OpConstant 0        // Push 5
OpGreaterThan       // Compare
OpJumpNotTruthy 10  // Jump if false
OpConstant 1        // Push 10
OpJump 13           // Skip alternative
OpConstant 2        // Push 20
```

---

### WhileStatement

```go
case *ast.WhileStatement:
    loopStart := len(c.currentInstructions())
    
    // Push loop context
    loopCtx := LoopContext{loopStart: loopStart}
    c.loopStack = append(c.loopStack, loopCtx)
    
    // Compile condition
    err := c.Compile(node.Condition)
    jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)
    
    // Compile body
    err = c.Compile(node.Body)
    if c.lastInstructionIs(code.OpPop) {
        c.removeLastPop()
    }
    
    // Jump back to start
    c.emit(code.OpJump, loopStart)
    
    // Patch exit jump
    afterLoopPos := len(c.currentInstructions())
    c.changeOperand(jumpNotTruthyPos, afterLoopPos)
    
    // Patch break statements
    ctx := c.loopStack[len(c.loopStack)-1]
    for _, pos := range ctx.breakPositions {
        c.changeOperand(pos, afterLoopPos)
    }
    
    // Patch continue statements
    for _, pos := range ctx.contPositions {
        c.changeOperand(pos, loopStart)
    }
    
    // Pop loop context
    c.loopStack = c.loopStack[:len(c.loopStack)-1]
    
    c.emit(code.OpNull)
```

**Example**:
```bhasa
যতক্ষণ (i < 10) {
    i = i + 1;
}
```

---

### ForStatement

```go
case *ast.ForStatement:
    // 1. Compile initialization
    if node.Init != nil {
        err := c.Compile(node.Init)
        if _, ok := node.Init.(*ast.ExpressionStatement); ok {
            c.emit(code.OpPop)
        }
    }
    
    loopStart := len(c.currentInstructions())
    
    // Push loop context
    loopCtx := LoopContext{loopStart: loopStart}
    c.loopStack = append(c.loopStack, loopCtx)
    
    // 2. Compile condition
    var jumpNotTruthyPos int
    if node.Condition != nil {
        err := c.Compile(node.Condition)
        jumpNotTruthyPos = c.emit(code.OpJumpNotTruthy, 9999)
    }
    
    // 3. Compile body
    err := c.Compile(node.Body)
    if c.lastInstructionIs(code.OpPop) {
        c.removeLastPop()
    }
    
    // Continue target (before increment)
    continueTarget := len(c.currentInstructions())
    
    // 4. Compile increment
    if node.Increment != nil {
        err := c.Compile(node.Increment)
        if _, ok := node.Increment.(*ast.ExpressionStatement); ok {
            c.emit(code.OpPop)
        }
    }
    
    // 5. Jump back to condition
    c.emit(code.OpJump, loopStart)
    
    afterLoopPos := len(c.currentInstructions())
    
    // Patch jumps
    if node.Condition != nil {
        c.changeOperand(jumpNotTruthyPos, afterLoopPos)
    }
    
    // Patch break/continue
    ctx := c.loopStack[len(c.loopStack)-1]
    for _, pos := range ctx.breakPositions {
        c.changeOperand(pos, afterLoopPos)
    }
    for _, pos := range ctx.contPositions {
        c.changeOperand(pos, continueTarget)
    }
    
    // Pop loop context
    c.loopStack = c.loopStack[:len(c.loopStack)-1]
    
    c.emit(code.OpNull)
```

**Key Difference from While**: Continue jumps to increment, not start

---

### BreakStatement

```go
case *ast.BreakStatement:
    if len(c.loopStack) == 0 {
        return fmt.Errorf("break statement outside loop")
    }
    // Emit jump with placeholder
    pos := c.emit(code.OpJump, 9999)
    // Record position in loop context
    ctx := &c.loopStack[len(c.loopStack)-1]
    ctx.breakPositions = append(ctx.breakPositions, pos)
```

**Patched**: After loop compilation, all breaks jump to end

---

### ContinueStatement

```go
case *ast.ContinueStatement:
    if len(c.loopStack) == 0 {
        return fmt.Errorf("continue statement outside loop")
    }
    pos := c.emit(code.OpJump, 9999)
    ctx := &c.loopStack[len(c.loopStack)-1]
    ctx.contPositions = append(ctx.contPositions, pos)
```

**Patched**: After loop compilation, all continues jump to:
- **While**: Loop start
- **For**: Increment statement

---

## Function Compilation

### FunctionLiteral

```go
case *ast.FunctionLiteral:
    // 1. Enter new scope
    c.enterScope()
    
    // 2. Define parameters
    for _, p := range node.Parameters {
        c.symbolTable.Define(p.Value)
    }
    
    // 3. Compile body
    err := c.Compile(node.Body)
    
    // 4. Handle implicit return
    if c.lastInstructionIs(code.OpPop) {
        c.replaceLastPopWithReturn()
    }
    if !c.lastInstructionIs(code.OpReturnValue) {
        c.emit(code.OpReturn)
    }
    
    // 5. Collect free variables
    freeSymbols := c.symbolTable.FreeSymbols
    numLocals := c.symbolTable.numDefinitions
    instructions := c.leaveScope()
    
    // 6. Load free variables
    for _, s := range freeSymbols {
        c.loadSymbol(s)
    }
    
    // 7. Create function object
    compiledFn := &object.CompiledFunction{
        Instructions:  instructions,
        NumLocals:     numLocals,
        NumParameters: len(node.Parameters),
    }
    fnIndex := c.addConstant(compiledFn)
    
    // 8. Emit closure
    c.emit(code.OpClosure, fnIndex, len(freeSymbols))
```

**Key Steps**:
1. **Enter scope**: New symbol table
2. **Define params**: Add to symbol table as locals
3. **Compile body**: Generate function bytecode
4. **Implicit return**: Add if missing
5. **Collect free vars**: Variables from outer scopes
6. **Load free vars**: Push onto stack
7. **Create function**: Add to constants
8. **Emit closure**: Create closure with free vars

**Example**:
```bhasa
ফাংশন(x, y) {
    ফেরত x + y;
}
```

---

### CallExpression

```go
case *ast.CallExpression:
    // 1. Compile function
    err := c.Compile(node.Function)
    
    // 2. Compile arguments
    for _, a := range node.Arguments {
        err := c.Compile(a)
    }
    
    // 3. Emit call
    c.emit(code.OpCall, len(node.Arguments))
```

**Stack Layout Before OpCall**:
```
[function, arg1, arg2, ..., argN]
```

**Stack Layout After OpCall**:
```
[returnValue]
```

---

### Closures

**Closure Example**:
```bhasa
ফাংশন(x) {
    ফাংশন(y) {
        ফেরত x + y;  // x is free variable
    }
}
```

**Outer Function**:
```
OpGetLocal 0         // Get x
OpClosure 0 1        // Create closure with 1 free var
OpReturnValue
```

**Inner Function** (in constants[0]):
```
OpGetFree 0          // Get x (free variable)
OpGetLocal 0         // Get y (parameter)
OpAdd
OpReturnValue
```

---

## OOP Compilation

### ClassDefinition

```go
func (c *Compiler) compileClassDefinition(node *ast.ClassDefinition) error {
    // 1. Create Class object
    class := &object.Class{
        Name: node.Name.Value,
        Fields: make(map[string]string),
        Methods: make(map[string]*object.Method),
        // ...
    }
    
    // 2. Process fields
    for _, field := range node.Fields {
        class.Fields[field.Name] = field.TypeAnnot.String()
        class.FieldAccess[field.Name] = string(field.Access)
        class.FieldOrder = append(class.FieldOrder, field.Name)
    }
    
    // 3. Compile constructor
    if len(node.Constructors) > 0 {
        // Enter scope
        c.enterScope()
        
        // Define 'this' (এই)
        c.symbolTable.Define("এই")
        
        // Define parameters
        for _, param := range constructor.Parameters {
            c.symbolTable.Define(param.Value)
        }
        
        // Compile body
        c.Compile(constructor.Body)
        
        // Implicit return this
        if !c.lastInstructionIs(code.OpReturnValue) {
            c.emit(code.OpGetLocal, 0)  // 'this' is param 0
            c.emit(code.OpReturnValue)
        }
        
        // Create closure
        // ...
        c.emit(code.OpDefineConstructor, fnIndex)
    }
    
    // 4. Compile methods
    for _, method := range node.Methods {
        // Similar to constructor
        // ...
        c.emit(code.OpDefineMethod, methodNameIndex)
    }
    
    // 5. Store class
    classIndex := c.addConstant(class)
    c.emit(code.OpClass, classIndex)
    
    // 6. Define in symbol table
    symbol := c.symbolTable.Define(node.Name.Value)
    if symbol.Scope == GlobalScope {
        c.emit(code.OpSetGlobal, symbol.Index)
    } else {
        c.emit(code.OpSetLocal, symbol.Index)
    }
    
    return nil
}
```

---

### NewExpression

```go
func (c *Compiler) compileNewExpression(node *ast.NewExpression) error {
    // 1. Compile arguments
    for _, arg := range node.Arguments {
        err := c.Compile(arg)
    }
    
    // 2. Load class
    err := c.Compile(node.ClassName)
    
    // 3. Create instance
    c.emit(code.OpNewInstance, len(node.Arguments))
    
    return nil
}
```

**Stack Layout**:
```
Before: [arg1, arg2, ..., argN, class]
After:  [instance]
```

---

### MethodCallExpression

```go
func (c *Compiler) compileMethodCall(node *ast.MethodCallExpression) error {
    // 1. Compile object
    err := c.Compile(node.Object)
    
    // 2. Push method name
    methodNameIndex := c.addConstant(&object.String{Value: node.MethodName.Value})
    c.emit(code.OpConstant, methodNameIndex)
    
    // 3. Compile arguments
    for _, arg := range node.Arguments {
        err := c.Compile(arg)
    }
    
    // 4. Call method
    c.emit(code.OpCallMethod, len(node.Arguments))
    
    return nil
}
```

---

### ThisExpression

```go
case *ast.ThisExpression:
    c.emit(code.OpGetThis)
```

**Inside Methods**: VM resolves `this` to current instance

---

### SuperExpression

```go
case *ast.SuperExpression:
    c.emit(code.OpGetSuper)
```

**Inside Methods**: VM resolves `super` to parent class

---

## Module System

### LoadAndCompileModule

```go
func (c *Compiler) LoadAndCompileModule(modulePath string) error {
    // 1. Load source
    source, err := c.moduleLoader(modulePath)
    
    // 2. Check circular imports
    if c.moduleCache[modulePath] {
        return nil  // Already loaded
    }
    c.moduleCache[modulePath] = true
    
    // 3. Parse module
    l := lexer.New(source)
    p := parser.New(l)
    program := p.ParseProgram()
    
    if len(p.Errors()) > 0 {
        return fmt.Errorf("parser errors: %v", p.Errors())
    }
    
    // 4. Compile module
    return c.Compile(program)
}
```

---

### DefaultModuleLoader

```go
func DefaultModuleLoader(modulePath string) (string, error) {
    extensions := []string{".ভাষা", ".bhasa"}
    searchPaths := []string{
        modulePath,
        "modules/" + modulePath,
    }
    
    // Try all combinations
    for _, basePath := range searchPaths {
        for _, ext := range extensions {
            testPath := basePath
            if !strings.HasSuffix(basePath, ext) {
                testPath = basePath + ext
            }
            if _, statErr := os.Stat(testPath); statErr == nil {
                content, err := os.ReadFile(testPath)
                return string(content), err
            }
        }
    }
    
    return "", fmt.Errorf("module not found: %s", modulePath)
}
```

**Features**:
- Multiple file extensions
- Multiple search paths
- Graceful error messages

---

## Helper Methods

### emit

```go
func (c *Compiler) emit(op code.Opcode, operands ...int) int {
    ins := code.Make(op, operands...)
    pos := c.addInstruction(ins)
    c.setLastInstruction(op, pos)
    return pos
}
```

**Returns**: Position of emitted instruction (for patching)

---

### addConstant

```go
func (c *Compiler) addConstant(obj object.Object) int {
    c.constants = append(c.constants, obj)
    return len(c.constants) - 1
}
```

**Returns**: Index in constant pool

---

### changeOperand

```go
func (c *Compiler) changeOperand(opPos int, operand int) {
    op := code.Opcode(c.currentInstructions()[opPos])
    newInstruction := code.Make(op, operand)
    c.replaceInstruction(opPos, newInstruction)
}
```

**Purpose**: Patch jump offsets

---

### loadSymbol

```go
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
```

**Purpose**: Load variable based on scope

---

### enterScope / leaveScope

```go
func (c *Compiler) enterScope() {
    scope := CompilationScope{
        instructions: code.Instructions{},
        lastInstruction: EmittedInstruction{},
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
```

**Purpose**: Manage nested scopes

---

## Optimization Techniques

### 1. OpPop Removal

**Pattern**:
```
OpSomeExpression
OpPop
OpReturnValue
```

**Optimized To**:
```
OpSomeExpression
OpReturnValue
```

**Implementation**:
```go
if c.lastInstructionIs(code.OpPop) {
    c.replaceLastPopWithReturn()
}
```

---

### 2. Constant Pool Reuse

**Instead of**:
```
"hello" appears 10 times → 10 constants
```

**We do**:
```
"hello" appears 10 times → 1 constant, referenced 10 times
```

---

### 3. Short-Circuit Evaluation (TODO)

**Pattern**:
```bhasa
false && expensive()
```

**Should compile to**:
```
OpFalse
OpJumpNotTruthy end  // Skip expensive() call
<expensive() code>
end:
```

---

### 4. Tail Call Optimization (TODO)

**Pattern**:
```bhasa
ফাংশন() {
    ফেরত myself();  // Tail call
}
```

**Could be optimized to a jump instead of call**

---

## See Also

- [Symbol Table Documentation](./symbol-table-documentation.md)
- [Compilation Examples](./compilation-examples.md)
- [Quick Reference](./quick-reference.md)
- [Bytecode Documentation](../../code/docs/)
- [VM Documentation](../../vm/docs/)

