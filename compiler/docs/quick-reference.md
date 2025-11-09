# Compiler Quick Reference

## Symbol Table Quick Lookup

### Scope Types

| Scope | Description | Index Range | Example |
|-------|-------------|-------------|---------|
| **GLOBAL** | Module-level | 0-65535 | `ধরি x = 5` |
| **LOCAL** | Function params/locals | 0-255 | `ফাংশন(x) { ধরি y = 10; }` |
| **BUILTIN** | Built-in functions | 0-255 | `len`, `দেখাও` |
| **FREE** | Captured variables | 0-255 | Closure variables |
| **FUNCTION** | Self-reference | 0 | Recursive functions |

### Symbol Table Operations

```go
// Create tables
global := NewSymbolTable()
local := NewEnclosedSymbolTable(global)

// Define symbols
s := table.Define("x")                    // Without type
s := table.DefineWithType("x", typeAnnot) // With type
s := table.DefineBuiltin(idx, "len")      // Builtin
s := table.DefineFunctionName("fn")       // For recursion

// Resolve symbols
symbol, ok := table.Resolve("x")
if !ok {
    // Variable not found
}
```

---

## Compilation Patterns

### Variable Declaration

```go
// Define symbol
symbol := c.symbolTable.Define(name)

// Compile value
c.Compile(value)

// Store
if symbol.Scope == GlobalScope {
    c.emit(code.OpSetGlobal, symbol.Index)
} else {
    c.emit(code.OpSetLocal, symbol.Index)
}
```

### Variable Access

```go
// Resolve
symbol, ok := c.symbolTable.Resolve(name)
if !ok {
    return fmt.Errorf("undefined: %s", name)
}

// Load
c.loadSymbol(symbol)
```

### Binary Operation

```go
c.Compile(left)     // Compile left
c.Compile(right)    // Compile right
c.emit(opcode)      // Emit operation
```

### If-Else

```go
c.Compile(condition)
jumpPos := c.emit(code.OpJumpNotTruthy, 9999)
c.Compile(consequence)
endPos := c.emit(code.OpJump, 9999)
c.changeOperand(jumpPos, len(c.currentInstructions()))
c.Compile(alternative)
c.changeOperand(endPos, len(c.currentInstructions()))
```

### While Loop

```go
start := len(c.currentInstructions())
c.loopStack = append(c.loopStack, LoopContext{loopStart: start})

c.Compile(condition)
exitPos := c.emit(code.OpJumpNotTruthy, 9999)
c.Compile(body)
c.emit(code.OpJump, start)

c.changeOperand(exitPos, len(c.currentInstructions()))

// Patch breaks/continues
ctx := c.loopStack[len(c.loopStack)-1]
for _, pos := range ctx.breakPositions {
    c.changeOperand(pos, len(c.currentInstructions()))
}
for _, pos := range ctx.contPositions {
    c.changeOperand(pos, start)
}

c.loopStack = c.loopStack[:len(c.loopStack)-1]
```

### Function

```go
c.enterScope()

// Define parameters
for _, p := range parameters {
    c.symbolTable.Define(p.Value)
}

// Compile body
c.Compile(body)

// Handle return
if !c.lastInstructionIs(code.OpReturnValue) {
    c.emit(code.OpReturn)
}

// Collect free vars and leave scope
freeSymbols := c.symbolTable.FreeSymbols
instructions := c.leaveScope()

// Load free vars
for _, s := range freeSymbols {
    c.loadSymbol(s)
}

// Create function and emit closure
fn := &object.CompiledFunction{...}
fnIndex := c.addConstant(fn)
c.emit(code.OpClosure, fnIndex, len(freeSymbols))
```

### Array Literal

```go
for _, element := range elements {
    c.Compile(element)
}
c.emit(code.OpArray, len(elements))
```

### Hash Literal

```go
// Sort keys for determinism
keys := sortKeys(pairs)

for _, key := range keys {
    c.Compile(key)
    c.Compile(pairs[key])
}
c.emit(code.OpHash, len(pairs)*2)
```

---

## Common Operations

### Add Constant

```go
index := c.addConstant(obj)
c.emit(code.OpConstant, index)
```

### Emit Instruction

```go
pos := c.emit(opcode, operands...)
// Returns position for patching
```

### Patch Jump

```go
// Emit jump with placeholder
jumpPos := c.emit(code.OpJump, 9999)

// ... compile more code

// Patch jump to current position
targetPos := len(c.currentInstructions())
c.changeOperand(jumpPos, targetPos)
```

### Load Symbol

```go
switch symbol.Scope {
case GlobalScope:
    c.emit(code.OpGetGlobal, symbol.Index)
case LocalScope:
    c.emit(code.OpGetLocal, symbol.Index)
case BuiltinScope:
    c.emit(code.OpGetBuiltin, symbol.Index)
case FreeScope:
    c.emit(code.OpGetFree, symbol.Index)
case FunctionScope:
    c.emit(code.OpCurrentClosure)
}
```

---

## Scope Management

### Enter Scope

```go
c.enterScope()
// Creates new compilation scope
// Creates new symbol table linked to current
```

### Leave Scope

```go
instructions := c.leaveScope()
// Pops compilation scope
// Restores previous symbol table
// Returns instructions from scope
```

---

## Loop Management

### Push Loop Context

```go
loopCtx := LoopContext{
    loopStart: len(c.currentInstructions()),
}
c.loopStack = append(c.loopStack, loopCtx)
```

### Record Break

```go
pos := c.emit(code.OpJump, 9999)
ctx := &c.loopStack[len(c.loopStack)-1]
ctx.breakPositions = append(ctx.breakPositions, pos)
```

### Record Continue

```go
pos := c.emit(code.OpJump, 9999)
ctx := &c.loopStack[len(c.loopStack)-1]
ctx.contPositions = append(ctx.contPositions, pos)
```

### Patch and Pop

```go
ctx := c.loopStack[len(c.loopStack)-1]

// Patch breaks to loop end
for _, pos := range ctx.breakPositions {
    c.changeOperand(pos, afterLoop)
}

// Patch continues to loop start
for _, pos := range ctx.contPositions {
    c.changeOperand(pos, loopStart)
}

// Pop context
c.loopStack = c.loopStack[:len(c.loopStack)-1]
```

---

## Optimization Helpers

### Check Last Instruction

```go
if c.lastInstructionIs(code.OpPop) {
    // Last instruction is OpPop
}
```

### Remove Last Pop

```go
c.removeLastPop()
// Removes OpPop from end of instructions
```

### Replace Pop with Return

```go
c.replaceLastPopWithReturn()
// Changes OpPop to OpReturnValue
```

---

## Error Handling

```go
// Undefined variable
symbol, ok := c.symbolTable.Resolve(name)
if !ok {
    return fmt.Errorf("undefined variable %s", name)
}

// Break outside loop
if len(c.loopStack) == 0 {
    return fmt.Errorf("break statement outside loop")
}

// Continue outside loop
if len(c.loopStack) == 0 {
    return fmt.Errorf("continue statement outside loop")
}

// Compilation error
err := c.Compile(node)
if err != nil {
    return err
}
```

---

## Module System

### Load Module

```go
err := c.LoadAndCompileModule(modulePath)
if err != nil {
    return fmt.Errorf("import error: %v", err)
}
```

### Check Circular Import

```go
if c.moduleCache[modulePath] {
    return nil  // Already loaded
}
c.moduleCache[modulePath] = true
```

---

## OOP Patterns

### Class Definition

```go
class := &object.Class{...}

// Compile constructor
c.enterScope()
c.symbolTable.Define("এই")  // this
// ... compile constructor
c.emit(code.OpDefineConstructor, idx)

// Compile methods
for _, method := range methods {
    c.enterScope()
    c.symbolTable.Define("এই")  // this
    // ... compile method
    c.emit(code.OpDefineMethod, nameIdx)
}

classIdx := c.addConstant(class)
c.emit(code.OpClass, classIdx)
```

### New Instance

```go
// Compile arguments
for _, arg := range args {
    c.Compile(arg)
}

// Load class
c.Compile(className)

// Create instance
c.emit(code.OpNewInstance, len(args))
```

### Method Call

```go
// Compile object
c.Compile(object)

// Push method name
nameIdx := c.addConstant(&object.String{Value: methodName})
c.emit(code.OpConstant, nameIdx)

// Compile arguments
for _, arg := range args {
    c.Compile(arg)
}

// Call method
c.emit(code.OpCallMethod, len(args))
```

---

## Constant Pool Management

```go
// Add constant
index := c.addConstant(obj)

// Access constants
constants := c.constants

// Get constant
obj := c.constants[index]
```

---

## Instruction Helpers

```go
// Get current instructions
ins := c.currentInstructions()

// Add instruction
pos := c.addInstruction(instruction)

// Replace instruction
c.replaceInstruction(pos, newInstruction)

// Change operand
c.changeOperand(pos, newOperand)

// Set last instruction tracking
c.setLastInstruction(opcode, pos)
```

---

## Debugging Tips

### Print Instructions

```go
bytecode := compiler.Bytecode()
fmt.Println(bytecode.Instructions.String())
```

### Print Symbol Table

```go
// Walk symbol table
for name, symbol := range table.store {
    fmt.Printf("%s: %+v\n", name, symbol)
}

// Print free symbols
fmt.Printf("Free: %+v\n", table.FreeSymbols)
```

### Print Constants

```go
for i, constant := range compiler.constants {
    fmt.Printf("%d: %s\n", i, constant.Inspect())
}
```

### Trace Compilation

```go
func (c *Compiler) Compile(node ast.Node) error {
    fmt.Printf("Compiling: %T\n", node)
    
    switch node := node.(type) {
    // ... cases
    }
}
```

---

## Common Mistakes

### ❌ Forgetting to Check Resolution

```go
symbol, _ := table.Resolve("x")  // Might be undefined!
c.loadSymbol(symbol)              // CRASH if undefined
```

✅ **Correct**:
```go
symbol, ok := table.Resolve("x")
if !ok {
    return fmt.Errorf("undefined: x")
}
c.loadSymbol(symbol)
```

---

### ❌ Not Patching Jumps

```go
jumpPos := c.emit(code.OpJump, 9999)
// ... more code
// FORGOT to patch!
```

✅ **Correct**:
```go
jumpPos := c.emit(code.OpJump, 9999)
// ... more code
c.changeOperand(jumpPos, targetPos)  // Patch it!
```

---

### ❌ Wrong Scope Management

```go
c.enterScope()
// ... forget to leave
```

✅ **Correct**:
```go
c.enterScope()
// ... compile
instructions := c.leaveScope()  // Always leave!
```

---

### ❌ Not Loading Free Variables

```go
freeSymbols := c.symbolTable.FreeSymbols
// ... create function
c.emit(code.OpClosure, fnIdx, len(freeSymbols))  // Wrong order!
```

✅ **Correct**:
```go
freeSymbols := c.symbolTable.FreeSymbols
for _, s := range freeSymbols {
    c.loadSymbol(s)  // Load BEFORE closure
}
c.emit(code.OpClosure, fnIdx, len(freeSymbols))
```

---

## Quick Checklist

### Before Compiling

- [ ] Create compiler with `New()`
- [ ] Builtins registered automatically

### During Compilation

- [ ] Enter/leave scopes properly
- [ ] Resolve variables before use
- [ ] Patch all jump placeholders
- [ ] Track loops for break/continue
- [ ] Load free vars before closures

### After Compilation

- [ ] Call `Bytecode()` to extract
- [ ] Check for compilation errors
- [ ] Verify bytecode with `String()`

---

## Performance Tips

1. **Reuse constants**: Same value → same index
2. **Minimize scopes**: Only when necessary
3. **Optimize jumps**: Short-circuit when possible
4. **Cache symbols**: Don't resolve repeatedly

---

## See Also

- [Compiler Documentation](./compiler-documentation.md)
- [Symbol Table Documentation](./symbol-table-documentation.md)
- [Compilation Examples](./compilation-examples.md)
- [Bytecode Documentation](../../code/docs/)

