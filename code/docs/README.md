# Bytecode Documentation

Welcome to the Bhasa bytecode and instruction set documentation!

## 📚 Documentation Files

### [Bytecode Documentation](./bytecode-documentation.md)
Comprehensive documentation covering:
- Complete bytecode system overview
- Detailed reference for all 78 opcodes
- Instruction format and encoding
- Operand widths and types
- API reference for bytecode manipulation
- Performance considerations and best practices

**Recommended for**: Deep understanding of the bytecode system, implementing new opcodes, compiler/VM development

### [Quick Reference](./quick-reference.md)
Concise reference guide with:
- Complete opcode lookup table
- Opcodes organized by category
- Operand width reference
- Common bytecode patterns
- Stack behavior guide
- Debugging helpers

**Recommended for**: Quick lookups, day-to-day development, cheat sheet

### [Instruction Examples](./instruction-examples.md)
Visual learning guide with:
- Step-by-step execution examples
- Stack state visualizations
- Examples from simple to complex
- Complete execution traces
- Real-world patterns

**Recommended for**: Learning how bytecode works, understanding execution flow, visual learners

---

## 🎯 What is Bytecode?

Bytecode is a low-level, platform-independent instruction set that the Bhasa Virtual Machine (VM) executes. It's the compiled form of Bhasa programs, sitting between the high-level AST and machine code.

### Compilation Pipeline

```
Bhasa Source Code
      ↓
   Lexer (tokenization)
      ↓
   Parser (AST generation)
      ↓
   Compiler (bytecode generation)  ← This module
      ↓
   Bytecode Instructions
      ↓
   Virtual Machine (execution)
      ↓
   Program Output
```

---

## 🏗️ Architecture

### Stack-Based Virtual Machine

Bhasa uses a **stack-based VM architecture** where:

- All operations work on a **stack** of values
- Instructions **push** values onto the stack
- Instructions **pop** values from the stack
- **No registers** - simpler design, more portable

**Example**:
```bhasa
5 + 3
```

Compiles to:
```
OpConstant 0    // Push 5
OpConstant 1    // Push 3
OpAdd           // Pop 5 and 3, push 8
```

Stack states:
```
[]  →  [5]  →  [5,3]  →  [8]
```

---

## 📊 Key Components

### 1. Instructions

```go
type Instructions []byte
```

A byte sequence representing compiled code. Can be:
- Disassembled to human-readable format
- Executed by the VM
- Serialized/deserialized for caching

### 2. Opcodes

```go
type Opcode byte
```

Single-byte operation codes. There are **78 opcodes** in Bhasa:

| Category | Count | Examples |
|----------|-------|----------|
| Arithmetic | 5 | `OpAdd`, `OpSub`, `OpMul` |
| Bitwise | 6 | `OpBitAnd`, `OpBitOr` |
| Comparison | 4 | `OpEqual`, `OpGreaterThan` |
| Logical | 3 | `OpBang`, `OpAnd`, `OpOr` |
| Control Flow | 2 | `OpJump`, `OpJumpNotTruthy` |
| Variables | 6 | `OpGetGlobal`, `OpSetLocal` |
| Functions | 6 | `OpCall`, `OpClosure` |
| Collections | 3 | `OpArray`, `OpHash`, `OpIndex` |
| Types | 3 | `OpTypeCheck`, `OpTypeCast` |
| Structs | 3 | `OpStruct`, `OpGetStructField` |
| Enums | 1 | `OpEnum` |
| OOP | 12 | `OpClass`, `OpNewInstance` |

### 3. Definitions

```go
type Definition struct {
    Name          string
    OperandWidths []int
}
```

Metadata about each opcode:
- **Name**: Human-readable name for debugging
- **OperandWidths**: Size of each operand in bytes

**Example**:
```go
OpConstant: {Name: "OpConstant", OperandWidths: []int{2}}
//                                                    └─ 2-byte operand
```

---

## 🔢 Instruction Format

### Basic Structure

Every instruction consists of:

```
[Opcode (1 byte)] [Operand 1] [Operand 2] ...
```

### Operand Sizes

- **0 bytes**: Operations on stack values (most common)
- **1 byte**: Small indices (locals, builtins)
- **2 bytes**: Large indices (constants, globals, jumps)

### Examples

**No operands** (1 byte total):
```
OpAdd: [0x02]
```

**2-byte operand** (3 bytes total):
```
OpConstant 5: [0x00, 0x00, 0x05]
              └──┘  └─────────┘
             opcode  operand
```

**Multiple operands** (4 bytes total):
```
OpClosure 3 2: [0x26, 0x00, 0x03, 0x02]
               └──┘  └─────────┘  └──┘
              opcode  constant   free vars
                      index (2B)  count (1B)
```

---

## 🎓 Getting Started

### For New Developers

1. **Start with basics**: Read [Quick Reference](./quick-reference.md) to understand the instruction set
2. **Learn by example**: Go through [Instruction Examples](./instruction-examples.md)
3. **Deep dive**: Read specific sections in [Bytecode Documentation](./bytecode-documentation.md)
4. **Practice**: Try disassembling simple programs

### For Compiler Development

1. Understand the **instruction format** and **operand widths**
2. Study how **control flow** is compiled (jumps, loops)
3. Learn **function calling conventions** and **closure handling**
4. Implement **compilation for new language features**

### For VM Development

1. Understand the **stack-based execution model**
2. Learn **instruction decoding** and **operand reading**
3. Study **function call frames** and **closure environments**
4. Implement **execution for each opcode**

### For Optimization

1. Study **instruction sequences** and identify patterns
2. Learn about **peephole optimization** opportunities
3. Understand **constant folding** and **dead code elimination**
4. Profile bytecode execution to find hotspots

---

## 📖 Instruction Categories

### Data Loading & Storage

**Load values**:
```
OpConstant      - Load from constant pool
OpTrue/OpFalse  - Load boolean literals
OpNull          - Load null value
OpGetGlobal     - Load global variable
OpGetLocal      - Load local variable
OpGetBuiltin    - Load builtin function
OpGetFree       - Load free variable (closure)
```

**Store values**:
```
OpSetGlobal     - Store global variable
OpSetLocal      - Store local variable
OpPop           - Discard top value
```

### Arithmetic & Bitwise

**Arithmetic**:
```
OpAdd, OpSub, OpMul, OpDiv, OpMod
OpMinus (unary negation)
```

**Bitwise**:
```
OpBitAnd, OpBitOr, OpBitXor, OpBitNot
OpLeftShift, OpRightShift
```

### Comparison & Logical

**Comparison**:
```
OpEqual, OpNotEqual
OpGreaterThan, OpGreaterThanEqual
```

**Logical**:
```
OpBang (NOT), OpAnd (AND), OpOr (OR)
```

### Control Flow

```
OpJump             - Unconditional jump
OpJumpNotTruthy    - Conditional jump
OpCall             - Function call
OpReturn           - Return void
OpReturnValue      - Return with value
```

### Collections

```
OpArray            - Create array
OpHash             - Create hash map
OpIndex            - Array/hash access
```

### Advanced Features

**Type System**:
```
OpTypeCheck, OpTypeCast, OpAssertType
```

**Structs**:
```
OpStruct, OpGetStructField, OpSetStructField
```

**Enums**:
```
OpEnum
```

**OOP**:
```
OpClass, OpNewInstance, OpCallMethod
OpGetThis, OpGetSuper
OpDefineMethod, OpDefineConstructor
OpInherit, OpInterface, OpCheckInterface
OpGetInstanceField, OpSetInstanceField
```

---

## 💡 Common Patterns

### Variable Declaration

```bhasa
ধরি x = 5;
```
```
OpConstant 0
OpSetGlobal 0
```

### Function Call

```bhasa
add(5, 3)
```
```
OpGetGlobal 0    // load function
OpConstant 1     // arg 1
OpConstant 2     // arg 2
OpCall 2         // call with 2 args
```

### If-Else

```bhasa
যদি (condition) { a } নাহলে { b }
```
```
<compile condition>
OpJumpNotTruthy <else_offset>
<compile a>
OpJump <end_offset>
<compile b>
```

### While Loop

```bhasa
যতক্ষণ (condition) { body }
```
```
<start>
<compile condition>
OpJumpNotTruthy <end>
<compile body>
OpJump <start>
<end>
```

### Array Literal

```bhasa
[1, 2, 3]
```
```
OpConstant 0    // 1
OpConstant 1    // 2
OpConstant 2    // 3
OpArray 3       // create array
```

---

## 🔍 Debugging Bytecode

### Disassemble Instructions

```go
instructions := compiler.Bytecode().Instructions
fmt.Println(instructions.String())
```

**Output**:
```
0000 OpConstant 0
0003 OpConstant 1
0006 OpAdd
0007 OpPop
```

### Understanding Offsets

- **Offset** is the byte position in the instruction sequence
- Jumps use **absolute offsets**, not relative
- Each instruction has a different size based on operands

**Example**:
```
0000 OpConstant 0    (3 bytes: 0000-0002)
0003 OpConstant 1    (3 bytes: 0003-0005)
0006 OpAdd           (1 byte:  0006)
0007 OpPop           (1 byte:  0007)
```

### Tracing Execution

Add debug output in the VM:

```go
fmt.Printf("PC: %04d | OP: %s | Stack: %v\n",
    vm.ip,
    code.Lookup(vm.currentInstruction()),
    vm.stack[:vm.sp])
```

**Output**:
```
PC: 0000 | OP: OpConstant 0 | Stack: []
PC: 0003 | OP: OpConstant 1 | Stack: [5]
PC: 0006 | OP: OpAdd        | Stack: [5, 3]
PC: 0007 | OP: OpPop        | Stack: [8]
```

---

## 🚀 API Usage

### Creating Instructions

```go
// Make a simple instruction
ins := code.Make(code.OpAdd)
// Result: [0x02]

// Make instruction with operand
ins := code.Make(code.OpConstant, 5)
// Result: [0x00, 0x00, 0x05]

// Make instruction with multiple operands
ins := code.Make(code.OpClosure, 3, 2)
// Result: [0x26, 0x00, 0x03, 0x02]
```

### Reading Instructions

```go
// Get opcode at position
opcode := instructions[ip]

// Look up opcode definition
def, err := code.Lookup(opcode)
if err != nil {
    log.Fatal(err)
}

// Read operands
operands, bytesRead := code.ReadOperands(def, instructions[ip+1:])

// Advance instruction pointer
ip += 1 + bytesRead
```

### Concatenating Instructions

```go
instructions := []byte{}
instructions = append(instructions, code.Make(code.OpConstant, 0)...)
instructions = append(instructions, code.Make(code.OpConstant, 1)...)
instructions = append(instructions, code.Make(code.OpAdd)...)
instructions = append(instructions, code.Make(code.OpPop)...)
```

---

## 🎯 Design Decisions

### Why Stack-Based?

**Pros**:
- ✅ Simpler implementation
- ✅ More compact bytecode
- ✅ No register allocation needed
- ✅ More portable

**Cons**:
- ❌ More stack operations
- ❌ Can't optimize register usage

**Conclusion**: For an interpreted language, simplicity wins.

### Why 1-Byte Opcodes?

- **256 possible opcodes** is more than enough
- **Compact encoding** saves memory
- **Fast lookup** in definition table

### Why Variable-Length Instructions?

- **No-operand instructions** = 1 byte only
- **Compact for common operations**
- **Flexible for different operand sizes**

### Operand Width Choices

- **1 byte (uint8)**: For small counts (locals, builtin indices)
  - Max 255 locals per scope is reasonable
  - Max 255 builtins is plenty
  
- **2 bytes (uint16)**: For large pools (constants, globals, jumps)
  - Max 65,535 constants is sufficient
  - Max 65,535 byte offset for jumps covers large programs

---

## 📈 Performance Considerations

### Instruction Compactness

Smaller bytecode = better cache utilization:

```
OpAdd (1 byte)   vs   Register Add (3-4 bytes)
OpPop (1 byte)   vs   Register Move (3-4 bytes)
```

### Jump Optimization

- Minimize jumps in hot loops
- Use fall-through for common case
- Short-circuit logical operators

### Constant Pool

- Reuse constants (don't duplicate)
- Order frequently-used constants first
- Consider constant folding at compile time

### Local Variables

Prefer locals over globals:
```
OpGetLocal  (2 bytes)  vs  OpGetGlobal (3 bytes)
OpSetLocal  (2 bytes)  vs  OpSetGlobal (3 bytes)
```

---

## 🔧 Extending the Bytecode

### Adding a New Opcode

1. **Define the opcode**:
```go
const (
    // ... existing opcodes
    OpMyNewOp Opcode = iota
)
```

2. **Add definition**:
```go
var definitions = map[Opcode]*Definition{
    // ... existing definitions
    OpMyNewOp: {"OpMyNewOp", []int{2}}, // 2-byte operand
}
```

3. **Implement in compiler**:
```go
func (c *Compiler) compile(node ast.Node) error {
    switch node := node.(type) {
    // ... existing cases
    case *ast.MyNewNode:
        c.emit(code.OpMyNewOp, operand)
    }
}
```

4. **Implement in VM**:
```go
func (vm *VM) Run() error {
    for vm.ip < len(vm.instructions) {
        switch op {
        // ... existing cases
        case code.OpMyNewOp:
            operand := code.ReadUint16(vm.instructions[vm.ip+1:])
            // Execute operation
            vm.ip += 3
        }
    }
}
```

5. **Add tests**:
```go
func TestOpMyNewOp(t *testing.T) {
    // Test compilation and execution
}
```

---

## 📚 Related Documentation

- **[AST Documentation](../../ast/docs/)** - How source code is parsed
- **[Compiler Documentation](../../compiler/docs/)** - How AST is compiled to bytecode
- **[VM Documentation](../../vm/docs/)** - How bytecode is executed
- **[Object System](../../object/docs/)** - Runtime value representation

---

## 🧪 Testing

### Unit Tests

Tests are in `code_test.go`:

```go
func TestMake(t *testing.T) {
    tests := []struct {
        op       code.Opcode
        operands []int
        expected []byte
    }{
        {code.OpConstant, []int{65534}, []byte{byte(code.OpConstant), 255, 254}},
        {code.OpAdd, []int{}, []byte{byte(code.OpAdd)}},
    }
    // ...
}
```

### Integration Tests

Test full compilation and execution:

```go
func TestCompileAndRun(t *testing.T) {
    input := "5 + 3"
    expected := 8
    
    program := parse(input)
    compiler := NewCompiler()
    compiler.Compile(program)
    
    vm := NewVM(compiler.Bytecode())
    vm.Run()
    
    result := vm.StackTop()
    if result.(*object.Integer).Value != expected {
        t.Errorf("wrong result")
    }
}
```

---

## 💭 Future Enhancements

### Potential Optimizations

1. **Peephole Optimization**:
   - `OpConstant x, OpConstant y, OpAdd` → `OpAddConstant x y`
   - `OpGetLocal x, OpGetLocal x` → `OpGetLocal x, OpDup`

2. **Instruction Fusion**:
   - `OpGetGlobal x, OpPush` → `OpPushGlobal x`
   - Common patterns as single instructions

3. **Jump Tables**:
   - For switch/match statements
   - More efficient than cascading if-else

4. **Register Optimization**:
   - Hybrid stack-register VM
   - Keep hot values in registers

### New Features

1. **Async/Await**:
   - `OpAwait`, `OpAsync`
   - Coroutine support

2. **Exception Handling**:
   - `OpTry`, `OpCatch`, `OpFinally`, `OpThrow`
   - Structured exception handling

3. **Generators**:
   - `OpYield`, `OpResume`
   - Iterator protocol

4. **Pattern Matching**:
   - `OpMatch`, `OpMatchCase`
   - Efficient pattern matching

---

## 🤝 Contributing

When modifying bytecode:

1. ✅ **Maintain backward compatibility** when possible
2. ✅ **Document all changes** in this documentation
3. ✅ **Add comprehensive tests** for new opcodes
4. ✅ **Consider performance implications**
5. ✅ **Update compiler AND VM** together
6. ✅ **Add examples** to instruction-examples.md

---

## 📞 Need Help?

- Check [Bytecode Documentation](./bytecode-documentation.md) for detailed explanations
- Check [Quick Reference](./quick-reference.md) for quick lookups
- Check [Instruction Examples](./instruction-examples.md) for visual examples
- Review test files for usage patterns

---

## 🗺️ Bytecode System Overview

```
┌──────────────────────────────────────────────────────┐
│                   Bhasa Program                      │
└─────────────────┬────────────────────────────────────┘
                  │
                  ↓
         ┌────────────────┐
         │     Lexer      │
         └────────┬───────┘
                  │ Tokens
                  ↓
         ┌────────────────┐
         │     Parser     │
         └────────┬───────┘
                  │ AST
                  ↓
         ┌────────────────┐
         │   Compiler     │ ← code.go defines instructions
         └────────┬───────┘
                  │ Bytecode
                  ↓
    ┌─────────────────────────────┐
    │      Instructions           │
    │  [OpCode][Operands]...      │ ← Compact binary format
    └─────────────┬───────────────┘
                  │
                  ↓
         ┌────────────────┐
         │   Virtual      │
         │   Machine      │ ← Executes bytecode
         └────────┬───────┘
                  │
                  ↓
         ┌────────────────┐
         │  Stack-Based   │
         │  Execution     │ ← All operations on stack
         └────────┬───────┘
                  │
                  ↓
         ┌────────────────┐
         │    Output      │
         └────────────────┘
```

---

**Happy Coding! 🚀**

