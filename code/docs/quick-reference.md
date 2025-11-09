# Bhasa Bytecode Quick Reference

## Opcode Summary Table

| # | Opcode | Operands | Stack Effect | Category |
|---|--------|----------|--------------|----------|
| 0 | OpConstant | uint16 | `[]` → `[value]` | Load |
| 1 | OpPop | - | `[a]` → `[]` | Stack |
| 2 | OpAdd | - | `[a,b]` → `[a+b]` | Arithmetic |
| 3 | OpSub | - | `[a,b]` → `[a-b]` | Arithmetic |
| 4 | OpMul | - | `[a,b]` → `[a*b]` | Arithmetic |
| 5 | OpDiv | - | `[a,b]` → `[a/b]` | Arithmetic |
| 6 | OpMod | - | `[a,b]` → `[a%b]` | Arithmetic |
| 7 | OpBitAnd | - | `[a,b]` → `[a&b]` | Bitwise |
| 8 | OpBitOr | - | `[a,b]` → `[a\|b]` | Bitwise |
| 9 | OpBitXor | - | `[a,b]` → `[a^b]` | Bitwise |
| 10 | OpBitNot | - | `[a]` → `[~a]` | Bitwise |
| 11 | OpLeftShift | - | `[a,b]` → `[a<<b]` | Bitwise |
| 12 | OpRightShift | - | `[a,b]` → `[a>>b]` | Bitwise |
| 13 | OpTrue | - | `[]` → `[true]` | Load |
| 14 | OpFalse | - | `[]` → `[false]` | Load |
| 15 | - | - | - | - |
| 16 | OpEqual | - | `[a,b]` → `[a==b]` | Comparison |
| 17 | OpNotEqual | - | `[a,b]` → `[a!=b]` | Comparison |
| 18 | OpGreaterThan | - | `[a,b]` → `[a>b]` | Comparison |
| 19 | OpGreaterThanEqual | - | `[a,b]` → `[a>=b]` | Comparison |
| 20 | OpMinus | - | `[a]` → `[-a]` | Unary |
| 21 | OpBang | - | `[a]` → `[!a]` | Logical |
| 22 | OpAnd | - | `[a,b]` → `[a&&b]` | Logical |
| 23 | OpOr | - | `[a,b]` → `[a\|\|b]` | Logical |
| 24 | OpJumpNotTruthy | uint16 | `[a]` → `[]` | Control |
| 25 | OpJump | uint16 | `[]` → `[]` | Control |
| 26 | OpNull | - | `[]` → `[null]` | Load |
| 27 | OpGetGlobal | uint16 | `[]` → `[value]` | Variable |
| 28 | OpSetGlobal | uint16 | `[a]` → `[]` | Variable |
| 29 | OpArray | uint16 | `[e1..eN]` → `[arr]` | Collection |
| 30 | OpHash | uint16 | `[k1,v1..kN,vN]` → `[hash]` | Collection |
| 31 | OpIndex | - | `[arr,idx]` → `[val]` | Collection |
| 32 | OpCall | uint8 | `[fn,args...]` → `[ret]` | Function |
| 33 | OpReturnValue | - | `[a]` → `[a]` | Function |
| 34 | OpReturn | - | `[]` → `[]` | Function |
| 35 | OpGetLocal | uint8 | `[]` → `[value]` | Variable |
| 36 | OpSetLocal | uint8 | `[a]` → `[]` | Variable |
| 37 | OpGetBuiltin | uint8 | `[]` → `[builtin]` | Function |
| 38 | OpClosure | uint16,uint8 | `[free...]` → `[closure]` | Function |
| 39 | OpGetFree | uint8 | `[]` → `[value]` | Variable |
| 40 | OpCurrentClosure | - | `[]` → `[closure]` | Function |
| 41 | OpTypeCheck | uint16 | `[a]` → `[bool]` | Type |
| 42 | OpTypeCast | uint16 | `[a]` → `[casted]` | Type |
| 43 | OpAssertType | uint16 | `[a]` → `[a]` | Type |
| 44 | OpStruct | uint16 | `[fields...]` → `[struct]` | Struct |
| 45 | OpGetStructField | - | `[obj,field]` → `[val]` | Struct |
| 46 | OpSetStructField | - | `[obj,fld,val]` → `[]` | Struct |
| 47 | OpEnum | uint16,uint16 | `[]` → `[enum]` | Enum |
| 48 | OpClass | uint16 | `[]` → `[class]` | OOP |
| 49 | OpNewInstance | uint8 | `[cls,args...]` → `[obj]` | OOP |
| 50 | OpCallMethod | uint8 | `[obj,mth,args...]` → `[ret]` | OOP |
| 51 | OpGetThis | - | `[]` → `[this]` | OOP |
| 52 | OpGetSuper | - | `[]` → `[super]` | OOP |
| 53 | OpDefineMethod | uint16 | `[cls,mth]` → `[cls]` | OOP |
| 54 | OpDefineConstructor | uint16 | `[cls,ctor]` → `[cls]` | OOP |
| 55 | OpInterface | uint16 | `[]` → `[iface]` | OOP |
| 56 | OpCheckInterface | uint16 | `[obj]` → `[bool]` | OOP |
| 57 | OpInherit | uint16 | `[cls]` → `[cls]` | OOP |
| 58 | OpGetInstanceField | - | `[obj,field]` → `[val]` | OOP |
| 59 | OpSetInstanceField | - | `[obj,fld,val]` → `[]` | OOP |

---

## Opcodes by Category

### Arithmetic Operations
```
OpAdd, OpSub, OpMul, OpDiv, OpMod
OpMinus (unary)
```

### Bitwise Operations
```
OpBitAnd, OpBitOr, OpBitXor, OpBitNot
OpLeftShift, OpRightShift
```

### Comparison Operations
```
OpEqual, OpNotEqual
OpGreaterThan, OpGreaterThanEqual
```

### Logical Operations
```
OpBang (NOT), OpAnd (AND), OpOr (OR)
```

### Stack Operations
```
OpPop
```

### Constant Loading
```
OpConstant (any value)
OpTrue, OpFalse, OpNull (literals)
```

### Variable Operations
```
OpGetGlobal, OpSetGlobal  (global scope)
OpGetLocal, OpSetLocal    (local scope)
OpGetFree                 (closure variables)
OpGetBuiltin              (builtin functions)
```

### Control Flow
```
OpJump              (unconditional)
OpJumpNotTruthy     (conditional)
```

### Function Operations
```
OpCall              (call function)
OpReturn            (return void)
OpReturnValue       (return value)
OpClosure           (create closure)
OpCurrentClosure    (get current for recursion)
```

### Collection Operations
```
OpArray             (create array)
OpHash              (create hash)
OpIndex             (array/hash access)
```

### Type System
```
OpTypeCheck         (check type)
OpTypeCast          (cast type)
OpAssertType        (assert type)
```

### Struct Operations
```
OpStruct            (create struct)
OpGetStructField    (get field)
OpSetStructField    (set field)
```

### Enum Operations
```
OpEnum              (create enum variant)
```

### OOP Operations
```
OpClass             (define class)
OpNewInstance       (create instance)
OpCallMethod        (call method)
OpGetThis           (get this)
OpGetSuper          (get super)
OpDefineMethod      (define method)
OpDefineConstructor (define constructor)
OpInterface         (define interface)
OpCheckInterface    (check implementation)
OpInherit           (set inheritance)
OpGetInstanceField  (get field)
OpSetInstanceField  (set field)
```

---

## Operand Width Reference

### No Operands (0 bytes)
```
OpPop, OpAdd, OpSub, OpMul, OpDiv, OpMod
OpBitAnd, OpBitOr, OpBitXor, OpBitNot
OpLeftShift, OpRightShift
OpTrue, OpFalse, OpNull
OpEqual, OpNotEqual, OpGreaterThan, OpGreaterThanEqual
OpMinus, OpBang, OpAnd, OpOr
OpReturn, OpReturnValue
OpIndex
OpGetStructField, OpSetStructField
OpGetThis, OpGetSuper
OpGetInstanceField, OpSetInstanceField
OpCurrentClosure
```

### 1-Byte Operands (uint8: 0-255)
```
OpCall           numArgs
OpGetLocal       localIndex
OpSetLocal       localIndex
OpGetBuiltin     builtinIndex
OpGetFree        freeVarIndex
OpNewInstance    numArgs
OpCallMethod     numArgs
```

### 2-Byte Operands (uint16: 0-65535)
```
OpConstant         constantIndex
OpGetGlobal        globalIndex
OpSetGlobal        globalIndex
OpJump             offset
OpJumpNotTruthy    offset
OpArray            numElements
OpHash             numPairs
OpTypeCheck        typeIndex
OpTypeCast         typeIndex
OpAssertType       typeIndex
OpStruct           defIndex
OpClass            defIndex
OpDefineMethod     nameIndex
OpDefineConstructor closureIndex
OpInterface        defIndex
OpCheckInterface   ifaceIndex
OpInherit          parentIndex
```

### Mixed Operands
```
OpClosure        constantIndex(uint16), numFree(uint8)
OpEnum           typeIndex(uint16), variantIndex(uint16)
```

---

## Common Patterns

### Variable Declaration
```
OpConstant <idx>     // value
OpSetGlobal <idx>    // store
```

### Variable Access
```
OpGetGlobal <idx>    // load
```

### Arithmetic Expression (a + b)
```
OpGetGlobal <a>
OpGetGlobal <b>
OpAdd
```

### If-Else
```
<condition>
OpJumpNotTruthy <else_offset>
<consequence>
OpJump <end_offset>
<alternative>
```

### While Loop
```
<condition>
OpJumpNotTruthy <end>
<body>
OpJump <start>
```

### Function Call
```
OpGetGlobal <fn>     // function
OpConstant <arg1>    // argument 1
OpConstant <arg2>    // argument 2
OpCall 2             // call with 2 args
```

### Array Creation
```
OpConstant <el1>
OpConstant <el2>
OpConstant <el3>
OpArray 3
```

### Array Access
```
OpGetGlobal <arr>
OpConstant <idx>
OpIndex
```

### Hash Creation
```
OpConstant <key1>
OpConstant <val1>
OpConstant <key2>
OpConstant <val2>
OpHash 2
```

### Function Definition
```
OpClosure <fnIdx> <numFree>
OpSetGlobal <nameIdx>
```

### Class Definition
```
OpClass <defIdx>
OpClosure <ctorIdx> 0
OpDefineConstructor <idx>
OpClosure <methIdx> 0
OpDefineMethod <nameIdx>
OpSetGlobal <classNameIdx>
```

### Object Instantiation
```
OpGetGlobal <classIdx>
OpConstant <arg1>
OpConstant <arg2>
OpNewInstance 2
```

### Method Call
```
OpGetGlobal <objIdx>
OpConstant <methodName>
OpConstant <arg>
OpCallMethod 1
```

### This Access
```
OpGetThis
OpConstant <fieldName>
OpGetInstanceField
```

---

## Instruction Size Reference

| Operands | Bytes | Example |
|----------|-------|---------|
| None | 1 | `OpAdd` = 1 byte |
| uint8 | 2 | `OpCall 2` = 2 bytes |
| uint16 | 3 | `OpConstant 5` = 3 bytes |
| uint16 + uint8 | 4 | `OpClosure 0 2` = 4 bytes |
| uint16 + uint16 | 5 | `OpEnum 0 1` = 5 bytes |

---

## Stack Behavior

### Stack-Based VM

All operations work on a stack:

```
Stack grows up ↑

[...]  ← Top of stack (TOS)
[val2]
[val1]
[val0]
```

### Binary Operations

```
Before:  [..., a, b]
OpAdd:   [..., a+b]
```

### Unary Operations

```
Before:  [..., a]
OpMinus: [..., -a]
```

### Comparison

```
Before:     [..., a, b]
OpEqual:    [..., bool]
```

### Jumps

```
OpJumpNotTruthy:
  Before:  [..., condition]
  After:   [...]  (condition popped)
  Effect:  Jump if condition is falsy
```

---

## Function Call Convention

### Calling

```
Stack before OpCall:
[function, arg1, arg2, argN]

Stack after OpCall:
[returnValue]
```

### Returning

```
OpReturnValue:
  Keeps top value on stack
  
OpReturn:
  Pushes null on stack
```

---

## Debugging Helpers

### Read Instruction at Offset

```go
opcode := instructions[offset]
def, _ := code.Lookup(opcode)
operands, width := code.ReadOperands(def, instructions[offset+1:])
```

### Format Single Instruction

```go
def, _ := code.Lookup(ins[i])
operands, read := code.ReadOperands(def, ins[i+1:])
fmt.Printf("%04d %s\n", i, fmtInstruction(def, operands))
i += 1 + read
```

### Check Stack Depth

Most operations have predictable stack effects:
- Binary ops: 2 → 1
- Unary ops: 1 → 1
- Loads: 0 → 1
- Stores: 1 → 0

---

## Common Values

### Builtin Function Indices

```
0: len      (length)
1: প্রথম     (first)
2: শেষ      (last)
3: বাকি     (rest)
4: যোগ      (push)
5: দেখাও    (print)
```

### Boolean Values

```
OpTrue  → true  (সত্য)
OpFalse → false (মিথ্যা)
```

### Null Value

```
OpNull → null/nil (শূন্য)
```

---

## Performance Tips

### Prefer Local Variables

```
OpGetLocal  (1 byte operand)
  vs
OpGetGlobal (2 byte operand)
```

### Minimize Stack Depth

- Reorder operations to keep stack shallow
- Use locals instead of leaving values on stack

### Optimize Jumps

- Calculate jump offsets at compile time
- Use short-circuit evaluation for && and ||

### Batch Operations

- Avoid redundant loads/stores
- Combine operations when possible

---

## Error Checking

### Common Errors

1. **Undefined Opcode**: Check opcode value is valid
2. **Wrong Operand Count**: Match definition's OperandWidths
3. **Stack Underflow**: Ensure values on stack before pop
4. **Invalid Jump**: Jump offset must be within bytecode
5. **Bad Constant Index**: Index must be < constant pool size

### Validation

```go
// Check opcode is defined
def, err := code.Lookup(opcode)
if err != nil {
    // Handle error
}

// Check operand count
if len(operands) != len(def.OperandWidths) {
    // Handle error
}

// Check stack has enough values
if stackPointer < requiredDepth {
    // Handle error
}
```

---

## Instruction Examples

### Load Constant 5

```
Bytes: [0x00, 0x00, 0x05]
Disassembly: OpConstant 5
```

### Add Two Numbers

```
Bytes: [0x02]
Disassembly: OpAdd
```

### Jump to Offset 20

```
Bytes: [0x19, 0x00, 0x14]
Disassembly: OpJump 20
```

### Call Function with 3 Args

```
Bytes: [0x20, 0x03]
Disassembly: OpCall 3
```

### Get Local Variable 0

```
Bytes: [0x23, 0x00]
Disassembly: OpGetLocal 0
```

### Create Array with 5 Elements

```
Bytes: [0x1D, 0x00, 0x05]
Disassembly: OpArray 5
```

---

## Quick Lookup: Opcode to Number

```
OpConstant:          0     OpGetLocal:          35
OpPop:               1     OpSetLocal:          36
OpAdd:               2     OpGetBuiltin:        37
OpSub:               3     OpClosure:           38
OpMul:               4     OpGetFree:           39
OpDiv:               5     OpCurrentClosure:    40
OpMod:               6     OpTypeCheck:         41
OpBitAnd:            7     OpTypeCast:          42
OpBitOr:             8     OpAssertType:        43
OpBitXor:            9     OpStruct:            44
OpBitNot:            10    OpGetStructField:    45
OpLeftShift:         11    OpSetStructField:    46
OpRightShift:        12    OpEnum:              47
OpTrue:              13    OpClass:             48
OpFalse:             14    OpNewInstance:       49
OpEqual:             16    OpCallMethod:        50
OpNotEqual:          17    OpGetThis:           51
OpGreaterThan:       18    OpGetSuper:          52
OpGreaterThanEqual:  19    OpDefineMethod:      53
OpMinus:             20    OpDefineConstructor: 54
OpBang:              21    OpInterface:         55
OpAnd:               22    OpCheckInterface:    56
OpOr:                23    OpInherit:           57
OpJumpNotTruthy:     24    OpGetInstanceField:  58
OpJump:              25    OpSetInstanceField:  59
OpNull:              26
OpGetGlobal:         27
OpSetGlobal:         28
OpArray:             29
OpHash:              30
OpIndex:             31
OpCall:              32
OpReturnValue:       33
OpReturn:            34
```

---

## See Also

- [Complete Bytecode Documentation](./bytecode-documentation.md)
- [Instruction Examples](./instruction-examples.md)
- [Compiler Documentation](../../compiler/docs/)
- [VM Documentation](../../vm/docs/)

