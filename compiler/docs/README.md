# Compiler Documentation

Welcome to the Bhasa compiler documentation!

## 📚 Documentation Files

### [Compiler Documentation](./compiler-documentation.md)
Comprehensive documentation for `compiler.go`:
- Complete compiler architecture
- Detailed compilation process for all AST node types
- Scope management and function compilation
- Control flow compilation (if-else, loops, break/continue)
- OOP compilation (classes, methods, inheritance)
- Module system and imports
- Optimization techniques

**Recommended for**: Understanding the compilation pipeline, implementing new language features, debugging compilation issues

### [Symbol Table Documentation](./symbol-table-documentation.md)
Complete reference for `symbol_table.go`:
- Symbol table architecture
- Scope resolution (global, local, free, builtin, function)
- Variable binding and lookups
- Closure implementation
- Type annotation tracking
- Nested scope handling

**Recommended for**: Understanding variable scoping, implementing new scope types, working with closures

### [Quick Reference](./quick-reference.md)
Concise reference guide:
- Compilation patterns for common constructs
- Symbol scope quick lookup
- Common compiler operations
- Troubleshooting guide

**Recommended for**: Quick lookups, day-to-day development

### [Compilation Examples](./compilation-examples.md)
Visual learning guide:
- Step-by-step compilation examples
- AST to bytecode transformation
- Scope management examples
- Real-world patterns

**Recommended for**: Learning the compilation process, visual learners

---

## 🎯 What is the Compiler?

The Bhasa compiler translates Abstract Syntax Tree (AST) nodes into bytecode instructions that the Virtual Machine (VM) can execute efficiently.

### Compilation Pipeline

```
Source Code
    ↓
  Lexer
    ↓
  Tokens
    ↓
  Parser
    ↓
   AST
    ↓
Compiler ← You are here
    ↓
Bytecode
    ↓
   VM
    ↓
 Output
```

---

## 🏗️ Architecture Overview

### Core Components

```
┌─────────────────────────────────────────────────┐
│                   Compiler                      │
├─────────────────────────────────────────────────┤
│  • Constants Pool    - Stores literals         │
│  • Symbol Table      - Tracks variables        │
│  • Scopes Stack      - Manages nested scopes   │
│  • Loop Stack        - Handles break/continue  │
│  • Module Cache      - Prevents circular deps  │
└─────────────────────────────────────────────────┘
         │
         ├── Symbol Table
         │   • Global Scope
         │   • Local Scope
         │   • Builtin Scope
         │   • Free Scope (closures)
         │   • Function Scope (recursion)
         │
         └── Compilation Scopes
             • Instructions buffer
             • Last instruction tracking
             • Previous instruction tracking
```

---

## 📊 Key Concepts

### 1. Two-Pass Compilation

The compiler works in two main passes:

**Pass 1: Compilation** - Converts AST to bytecode
```go
compiler := New()
err := compiler.Compile(ast)
```

**Pass 2: Bytecode Extraction** - Extract compiled bytecode
```go
bytecode := compiler.Bytecode()
// bytecode.Instructions: []byte
// bytecode.Constants: []object.Object
```

### 2. Stack-Based Compilation

All expressions compile to stack operations:

```bhasa
5 + 3
```

Compiles to:
```
OpConstant 0    // Push 5
OpConstant 1    // Push 3
OpAdd           // Pop both, push result
```

### 3. Symbol Resolution

Variables are resolved through the symbol table:

```bhasa
ধরি x = 10;     // Define in symbol table
x + 5           // Resolve from symbol table
```

### 4. Scope Management

Nested scopes are tracked with a stack:

```bhasa
ধরি x = 5;              // Global scope
ফাংশন() {
    ধরি y = 10;         // Local scope
    ফাংশন() {
        ধরি z = x + y;  // Nested local scope
                        // x = free variable
                        // y = free variable
    }
}
```

---

## 🔑 Key Types

### Compiler

```go
type Compiler struct {
    constants    []object.Object   // Constant pool
    symbolTable  *SymbolTable      // Variable tracking
    scopes       []CompilationScope // Scope stack
    scopeIndex   int               // Current scope
    loopStack    []LoopContext     // Loop tracking
    moduleCache  map[string]bool   // Module tracking
    moduleLoader ModuleLoader      // Module loading
}
```

**Purpose**: Main compilation state machine

### CompilationScope

```go
type CompilationScope struct {
    instructions        code.Instructions  // Bytecode buffer
    lastInstruction     EmittedInstruction // Last opcode
    previousInstruction EmittedInstruction // Previous opcode
}
```

**Purpose**: Tracks compilation state for each scope (global, function, etc.)

### SymbolTable

```go
type SymbolTable struct {
    Outer          *SymbolTable         // Parent scope
    store          map[string]Symbol    // Symbol storage
    numDefinitions int                  // Number of symbols
    FreeSymbols    []Symbol             // Captured variables
}
```

**Purpose**: Manages variable names and their scopes

### Symbol

```go
type Symbol struct {
    Name      string                // Variable name
    Scope     SymbolScope          // Scope type
    Index     int                  // Storage index
    TypeAnnot *ast.TypeAnnotation  // Optional type
}
```

**Purpose**: Represents a single variable binding

---

## 📖 Compilation Process

### Example: Variable Declaration

**Bhasa Code**:
```bhasa
ধরি x = 5 + 3;
```

**Compilation Steps**:

1. **Compile LetStatement**:
   - Define symbol `x` in symbol table
   - Compile the value expression

2. **Compile InfixExpression (5 + 3)**:
   - Compile left operand (5)
   - Compile right operand (3)
   - Emit OpAdd

3. **Store Result**:
   - Emit OpSetGlobal or OpSetLocal based on scope

**Generated Bytecode**:
```
OpConstant 0    // Push 5
OpConstant 1    // Push 3
OpAdd           // Add them
OpSetGlobal 0   // Store in x
```

---

### Example: Function Definition

**Bhasa Code**:
```bhasa
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
```

**Compilation Steps**:

1. **Enter new scope** (function scope)
2. **Define parameters** (a, b) in symbol table
3. **Compile function body**
4. **Add implicit return** if needed
5. **Collect free variables** (closure)
6. **Leave scope** and extract instructions
7. **Create CompiledFunction** object
8. **Emit OpClosure** instruction

**Generated Bytecode** (function body):
```
OpGetLocal 0    // Get a
OpGetLocal 1    // Get b
OpAdd           // Add them
OpReturnValue   // Return result
```

**Generated Bytecode** (main):
```
OpClosure 0 0   // Create closure (fn index 0, 0 free vars)
OpSetGlobal 0   // Store as add
```

---

### Example: Control Flow

**Bhasa Code**:
```bhasa
যদি (x > 5) {
    10
} নাহলে {
    20
}
```

**Compilation Steps**:

1. **Compile condition** (x > 5)
2. **Emit OpJumpNotTruthy** with placeholder offset
3. **Compile consequence** (10)
4. **Emit OpJump** to skip alternative
5. **Patch first jump** to point to alternative
6. **Compile alternative** (20)
7. **Patch second jump** to point to end

**Generated Bytecode**:
```
OpGetGlobal 0       // Load x
OpConstant 0        // Push 5
OpGreaterThan       // Compare
OpJumpNotTruthy 10  // Jump to alternative if false
OpConstant 1        // Push 10
OpJump 13           // Jump over alternative
OpConstant 2        // Push 20
```

---

## 🎓 Getting Started

### For New Developers

1. **Understand the pipeline**: Read the architecture overview
2. **Study symbol tables**: Learn about scopes in [Symbol Table Documentation](./symbol-table-documentation.md)
3. **Follow examples**: Work through [Compilation Examples](./compilation-examples.md)
4. **Read compiler code**: Study [Compiler Documentation](./compiler-documentation.md)

### For Feature Development

1. **Identify AST node type** to compile
2. **Add case to Compile()** method
3. **Emit appropriate bytecode** instructions
4. **Update symbol table** if needed
5. **Handle scoping** correctly
6. **Add tests** for the new feature

### For Debugging

1. **Use String() method** on instructions to see bytecode
2. **Check symbol table** state
3. **Trace compilation** with print statements
4. **Verify bytecode** matches expectations
5. **Test with simple examples** first

---

## 💡 Common Patterns

### Variable Definition

```go
// Define symbol
symbol := c.symbolTable.Define(name)

// Compile value
c.Compile(valueExpression)

// Store based on scope
if symbol.Scope == GlobalScope {
    c.emit(code.OpSetGlobal, symbol.Index)
} else {
    c.emit(code.OpSetLocal, symbol.Index)
}
```

### Variable Access

```go
// Resolve symbol
symbol, ok := c.symbolTable.Resolve(name)
if !ok {
    return fmt.Errorf("undefined variable %s", name)
}

// Load based on scope
c.loadSymbol(symbol)
```

### Binary Operation

```go
// Compile left operand
c.Compile(node.Left)

// Compile right operand
c.Compile(node.Right)

// Emit operation
c.emit(operatorOpcode)
```

### Control Flow

```go
// Compile condition
c.Compile(condition)

// Emit jump with placeholder
jumpPos := c.emit(code.OpJumpNotTruthy, 9999)

// Compile consequence
c.Compile(consequence)

// Patch jump to point here
afterPos := len(c.currentInstructions())
c.changeOperand(jumpPos, afterPos)
```

### Function Compilation

```go
// Enter new scope
c.enterScope()

// Define parameters
for _, param := range parameters {
    c.symbolTable.Define(param.Value)
}

// Compile body
c.Compile(body)

// Add return if needed
if !c.lastInstructionIs(code.OpReturnValue) {
    c.emit(code.OpReturn)
}

// Collect free variables and exit scope
freeSymbols := c.symbolTable.FreeSymbols
instructions := c.leaveScope()

// Create function object and emit closure
compiledFn := &object.CompiledFunction{...}
fnIndex := c.addConstant(compiledFn)
c.emit(code.OpClosure, fnIndex, len(freeSymbols))
```

---

## 🔍 Scope Management

### Scope Types

| Scope | Description | Example |
|-------|-------------|---------|
| **Global** | Top-level variables | `ধরি x = 5` (at program level) |
| **Local** | Function parameters and locals | `ফাংশন(x) { ধরি y = 10; }` |
| **Builtin** | Built-in functions | `len`, `দেখাও`, `প্রথম` |
| **Free** | Captured from outer scope | Closure variables |
| **Function** | Function name for recursion | Named recursive functions |

### Scope Resolution

```
Current Scope
     │
     ├─ Check local store
     │  └─ Found? → Return symbol
     │
     └─ Not found?
        └─ Check Outer scope
           ├─ Global/Builtin? → Return as-is
           └─ Local? → Define as Free variable
```

---

## 🚀 Optimization Techniques

### 1. Constant Folding (Future)

```bhasa
5 + 3  // Could be compiled to: OpConstant 8
```

### 2. Dead Code Elimination

Removes unreachable code after returns/breaks

### 3. Jump Optimization

Short-circuit evaluation for `&&` and `||`:

```bhasa
false && expensive()  // Don't call expensive()
```

### 4. Tail Call Optimization (Future)

Optimize recursive tail calls

---

## 🔧 Extending the Compiler

### Adding a New Language Feature

1. **Define AST node** in `ast/ast.go`
2. **Add parser support** in `parser/parser.go`
3. **Add compiler case**:

```go
func (c *Compiler) Compile(node ast.Node) error {
    switch node := node.(type) {
    // ... existing cases ...
    
    case *ast.MyNewNode:
        return c.compileMyNewFeature(node)
    }
}
```

4. **Implement compilation**:

```go
func (c *Compiler) compileMyNewFeature(node *ast.MyNewNode) error {
    // Compile sub-expressions
    err := c.Compile(node.SubExpression)
    if err != nil {
        return err
    }
    
    // Emit bytecode
    c.emit(code.OpMyNewOp, operand)
    
    return nil
}
```

5. **Add VM support** in `vm/vm.go`
6. **Write tests**

---

## 📚 Related Documentation

- **[AST Documentation](../../ast/docs/)** - Input to the compiler
- **[Bytecode Documentation](../../code/docs/)** - Output from the compiler
- **[VM Documentation](../../vm/docs/)** - Executes the bytecode
- **[Object System](../../object/docs/)** - Runtime values

---

## 🧪 Testing

### Unit Tests

Located in `compiler_test.go`:

```go
func TestIntegerArithmetic(t *testing.T) {
    tests := []compilerTestCase{
        {
            input:    "1 + 2",
            expected: []code.Instructions{
                code.Make(code.OpConstant, 0),
                code.Make(code.OpConstant, 1),
                code.Make(code.OpAdd),
                code.Make(code.OpPop),
            },
        },
    }
    runCompilerTests(t, tests)
}
```

### Integration Tests

Test full pipeline: source → AST → bytecode → execution

```go
func TestCompileAndRun(t *testing.T) {
    input := "5 + 3"
    program := parse(input)
    
    compiler := New()
    compiler.Compile(program)
    
    vm := vm.New(compiler.Bytecode())
    vm.Run()
    
    result := vm.StackTop()
    // Assert result is 8
}
```

---

## 💭 Design Decisions

### Why Stack-Based?

- ✅ Simpler implementation
- ✅ No register allocation
- ✅ Portable bytecode
- ✅ Easy to reason about

### Why Symbol Tables?

- ✅ O(1) variable lookup
- ✅ Efficient scope nesting
- ✅ Support for closures
- ✅ Type annotation tracking

### Why Constant Pool?

- ✅ Reuse identical constants
- ✅ Reference by index (small)
- ✅ Separate code from data

---

## 🤝 Contributing

When modifying the compiler:

1. ✅ **Maintain AST compatibility**
2. ✅ **Update symbol table** correctly
3. ✅ **Emit correct bytecode** sequence
4. ✅ **Handle all scopes** properly
5. ✅ **Add comprehensive tests**
6. ✅ **Update documentation**
7. ✅ **Consider optimization** opportunities

---

## 📞 Need Help?

- Check [Compiler Documentation](./compiler-documentation.md) for detailed explanations
- Check [Symbol Table Documentation](./symbol-table-documentation.md) for scoping questions
- Check [Quick Reference](./quick-reference.md) for quick lookups
- Check [Compilation Examples](./compilation-examples.md) for visual examples
- Review test files for usage patterns

---

## 🗺️ Compiler System Overview

```
                    ┌──────────────┐
                    │ Source Code  │
                    └──────┬───────┘
                           │
                           ↓
                    ┌──────────────┐
                    │   Parser     │
                    └──────┬───────┘
                           │ AST
                           ↓
        ┌──────────────────────────────────────┐
        │           Compiler                   │
        ├──────────────────────────────────────┤
        │                                      │
        │  ┌────────────────┐  ┌────────────┐ │
        │  │ Symbol Table   │  │  Scopes    │ │
        │  ├────────────────┤  ├────────────┤ │
        │  │ • Global       │  │ • Global   │ │
        │  │ • Local        │  │ • Function │ │
        │  │ • Builtin      │  │ • Block    │ │
        │  │ • Free         │  └────────────┘ │
        │  │ • Function     │                 │
        │  └────────────────┘                 │
        │                                      │
        │  ┌────────────────┐                 │
        │  │ Constants Pool │                 │
        │  ├────────────────┤                 │
        │  │ • Integers     │                 │
        │  │ • Strings      │                 │
        │  │ • Functions    │                 │
        │  │ • Classes      │                 │
        │  └────────────────┘                 │
        │                                      │
        └──────────────┬───────────────────────┘
                       │
                       ↓
            ┌──────────────────┐
            │    Bytecode      │
            ├──────────────────┤
            │ • Instructions   │
            │ • Constants      │
            └──────┬───────────┘
                   │
                   ↓
            ┌──────────────────┐
            │  Virtual Machine │
            └──────────────────┘
```

---

**Happy Compiling! 🚀**

