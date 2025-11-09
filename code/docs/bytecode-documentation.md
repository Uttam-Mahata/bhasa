# Bhasa Bytecode Documentation

## Table of Contents

1. [Overview](#overview)
2. [Core Types](#core-types)
3. [Instruction Format](#instruction-format)
4. [Opcode Reference](#opcode-reference)
5. [Operand Widths](#operand-widths)
6. [Instruction Categories](#instruction-categories)
7. [API Reference](#api-reference)
8. [Examples](#examples)
9. [Debugging](#debugging)

---

## Overview

The Bhasa bytecode system is the foundation of the virtual machine (VM) that executes compiled Bhasa programs. Instead of interpreting the AST directly (which would be slow), the compiler translates the AST into a sequence of low-level bytecode instructions that the VM can execute efficiently.

### Why Bytecode?

**Benefits**:
- ⚡ **Fast Execution**: Direct instruction execution is faster than tree traversal
- 💾 **Compact Representation**: Binary format is more compact than AST
- 🔧 **Optimizable**: Easier to apply optimizations
- 🎯 **Predictable Performance**: Constant-time instruction dispatch

**Architecture**:
```
Source Code → Lexer → Tokens → Parser → AST → Compiler → Bytecode → VM → Execution
```

---

## Core Types

### Instructions

```go
type Instructions []byte
```

A sequence of bytes representing compiled bytecode. Each instruction consists of:
1. **Opcode** (1 byte): The operation to perform
2. **Operands** (0-4 bytes): Arguments to the operation

**Example**:
```
[OpConstant, 0, 5, OpPop]
 │           └──┘  └─── Pop from stack
 │            │
 │            └─ Operand: constant index 5
 └─ Opcode: Load constant
```

---

### Opcode

```go
type Opcode byte
```

A single byte representing an instruction operation. There are 78 different opcodes in Bhasa, ranging from basic arithmetic to object-oriented operations.

**Example**:
```go
OpAdd      // Opcode value: 2
OpConstant // Opcode value: 0
```

---

### Definition

```go
type Definition struct {
    Name          string  // Human-readable name
    OperandWidths []int   // Width of each operand in bytes
}
```

Metadata about an opcode, including its name and the size of its operands.

**Example**:
```go
definitions[OpConstant] = &Definition{
    Name:          "OpConstant",
    OperandWidths: []int{2},  // One 2-byte operand
}
```

---

## Instruction Format

### General Structure

```
[Opcode] [Operand1] [Operand2] ... [OperandN]
   1B      varies     varies        varies
```

### Operand Widths

Operands can be:
- **0 bytes**: No operands (e.g., `OpAdd`, `OpPop`)
- **1 byte**: Small integers (0-255) for local variables, builtin indices
- **2 bytes**: Large integers (0-65535) for constants, global variables, jump offsets

### Examples

**No Operands**:
```
OpAdd: [0x02]
       └─ Just the opcode
```

**One 2-byte Operand**:
```
OpConstant: [0x00, 0x00, 0x05]
            └──┘  └────────┘
             │         │
          Opcode   Constant index 5
```

**Two Operands**:
```
OpClosure: [0x36, 0x00, 0x03, 0x02]
           └──┘  └────────┘  └──┘
            │         │        │
         Opcode    Const #3   2 free vars
```

---

## Opcode Reference

### Arithmetic Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpAdd` | 2 | None | `[a, b]` → `[a+b]` | Add two values |
| `OpSub` | 3 | None | `[a, b]` → `[a-b]` | Subtract b from a |
| `OpMul` | 4 | None | `[a, b]` → `[a*b]` | Multiply two values |
| `OpDiv` | 5 | None | `[a, b]` → `[a/b]` | Divide a by b |
| `OpMod` | 6 | None | `[a, b]` → `[a%b]` | Modulo operation |

**Example**:
```bhasa
5 + 3
```
Compiles to:
```
OpConstant 0  // Push 5
OpConstant 1  // Push 3
OpAdd         // Pop both, push 8
```

---

### Bitwise Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpBitAnd` | 7 | None | `[a, b]` → `[a&b]` | Bitwise AND |
| `OpBitOr` | 8 | None | `[a, b]` → `[a\|b]` | Bitwise OR |
| `OpBitXor` | 9 | None | `[a, b]` → `[a^b]` | Bitwise XOR |
| `OpBitNot` | 10 | None | `[a]` → `[~a]` | Bitwise NOT |
| `OpLeftShift` | 11 | None | `[a, b]` → `[a<<b]` | Left shift |
| `OpRightShift` | 12 | None | `[a, b]` → `[a>>b]` | Right shift |

**Example**:
```bhasa
5 & 3  // Bitwise AND
```
Compiles to:
```
OpConstant 0  // Push 5
OpConstant 1  // Push 3
OpBitAnd      // Pop both, push 1
```

---

### Comparison Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpEqual` | 16 | None | `[a, b]` → `[a==b]` | Equal comparison |
| `OpNotEqual` | 17 | None | `[a, b]` → `[a!=b]` | Not equal |
| `OpGreaterThan` | 18 | None | `[a, b]` → `[a>b]` | Greater than |
| `OpGreaterThanEqual` | 19 | None | `[a, b]` → `[a>=b]` | Greater or equal |

**Note**: Less than is implemented as `b > a` (operand reversal)

**Example**:
```bhasa
x > 10
```
Compiles to:
```
OpGetGlobal 0  // Push x
OpConstant 1   // Push 10
OpGreaterThan  // Pop both, push comparison result
```

---

### Logical Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpBang` | 21 | None | `[a]` → `[!a]` | Logical NOT |
| `OpAnd` | 22 | None | `[a, b]` → `[a&&b]` | Logical AND |
| `OpOr` | 23 | None | `[a, b]` → `[a\|\|b]` | Logical OR |

**Example**:
```bhasa
!সত্য  // NOT true
```
Compiles to:
```
OpTrue  // Push true
OpBang  // Pop, push false
```

---

### Unary Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpMinus` | 20 | None | `[a]` → `[-a]` | Unary negation |

**Example**:
```bhasa
-5
```
Compiles to:
```
OpConstant 0  // Push 5
OpMinus       // Pop, push -5
```

---

### Constant Loading

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpConstant` | 0 | `index: uint16` | `[]` → `[constant]` | Load constant from pool |
| `OpTrue` | 13 | None | `[]` → `[true]` | Push true |
| `OpFalse` | 14 | None | `[]` → `[false]` | Push false |
| `OpNull` | 26 | None | `[]` → `[null]` | Push null |

**Example**:
```bhasa
ধরি x = 42;
```
The constant `42` is stored in the constant pool at index 0:
```
OpConstant 0  // Push constant[0] (42)
OpSetGlobal 0 // Store in global[0] (x)
```

---

### Stack Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpPop` | 1 | None | `[a]` → `[]` | Pop top value |

**Example**:
```bhasa
5 + 3;  // Expression statement, result discarded
```
Compiles to:
```
OpConstant 0  // Push 5
OpConstant 1  // Push 3
OpAdd         // Push 8
OpPop         // Discard result
```

---

### Variable Operations

#### Global Variables

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetGlobal` | 27 | `index: uint16` | `[]` → `[value]` | Load global variable |
| `OpSetGlobal` | 28 | `index: uint16` | `[value]` → `[]` | Store global variable |

**Example**:
```bhasa
ধরি x = 10;  // Global variable
x = 20;      // Assignment
```
Compiles to:
```
OpConstant 0   // Push 10
OpSetGlobal 0  // x = 10
OpConstant 1   // Push 20
OpSetGlobal 0  // x = 20
```

#### Local Variables

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetLocal` | 35 | `index: uint8` | `[]` → `[value]` | Load local variable |
| `OpSetLocal` | 36 | `index: uint8` | `[value]` → `[]` | Store local variable |

**Example**:
```bhasa
ফাংশন() {
    ধরি x = 5;  // Local variable
    x = 10;      // Local assignment
}
```
Compiles to:
```
OpConstant 0   // Push 5
OpSetLocal 0   // x = 5 (local slot 0)
OpConstant 1   // Push 10
OpSetLocal 0   // x = 10
```

#### Free Variables (Closures)

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetFree` | 39 | `index: uint8` | `[]` → `[value]` | Load free variable |

**Example**:
```bhasa
ফাংশন() {
    ধরি x = 5;
    ফাংশন() {
        ফেরত x;  // x is a free variable
    }
}
```

---

### Control Flow

#### Jumps

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpJump` | 25 | `offset: uint16` | `[]` → `[]` | Unconditional jump |
| `OpJumpNotTruthy` | 24 | `offset: uint16` | `[cond]` → `[]` | Jump if condition is falsy |

**Example**:
```bhasa
যদি (x > 5) {
    দেখাও("বড়");
}
```
Compiles to:
```
OpGetGlobal 0      // Push x
OpConstant 1       // Push 5
OpGreaterThan      // Compare
OpJumpNotTruthy 10 // Jump to end if false
OpConstant 2       // Push "বড়"
OpGetBuiltin 0     // Get দেখাও
OpCall 1           // Call with 1 arg
```

---

### Function Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpCall` | 32 | `numArgs: uint8` | `[fn, arg1, ...argN]` → `[result]` | Call function |
| `OpReturn` | 34 | None | `[]` → `[]` | Return without value |
| `OpReturnValue` | 33 | None | `[value]` → `[value]` | Return with value |
| `OpClosure` | 38 | `constIdx: uint16, freeVars: uint8` | `[free1, ...freeN]` → `[closure]` | Create closure |
| `OpCurrentClosure` | 40 | None | `[]` → `[closure]` | Get current closure (for recursion) |

**Example**:
```bhasa
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};
যোগ(5, 3);
```
Compiles to:
```
// Function definition
OpClosure 0 0      // Create closure from constant[0], 0 free vars
OpSetGlobal 0      // Store as যোগ

// Function call
OpGetGlobal 0      // Load যোগ
OpConstant 1       // Push 5
OpConstant 2       // Push 3
OpCall 2           // Call with 2 args
```

---

### Builtin Functions

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetBuiltin` | 37 | `index: uint8` | `[]` → `[builtin]` | Load builtin function |

**Builtin Functions**:
- `len()` - Get length
- `প্রথম()` (first) - Get first element
- `শেষ()` (last) - Get last element
- `বাকি()` (rest) - Get all but first
- `যোগ()` (push) - Add element
- `দেখাও()` (print) - Print value

**Example**:
```bhasa
দেখাও("হ্যালো");
```
Compiles to:
```
OpConstant 0     // Push "হ্যালো"
OpGetBuiltin 5   // Get দেখাও (index 5)
OpCall 1         // Call with 1 arg
```

---

### Array Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpArray` | 29 | `size: uint16` | `[el1, ...elN]` → `[array]` | Create array |
| `OpIndex` | 31 | None | `[array, index]` → `[element]` | Index operation |

**Example**:
```bhasa
ধরি arr = [1, 2, 3];
arr[1]
```
Compiles to:
```
// Array creation
OpConstant 0       // Push 1
OpConstant 1       // Push 2
OpConstant 2       // Push 3
OpArray 3          // Create array with 3 elements
OpSetGlobal 0      // Store as arr

// Array indexing
OpGetGlobal 0      // Load arr
OpConstant 3       // Push 1 (index)
OpIndex            // Get arr[1]
```

---

### Hash Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpHash` | 30 | `pairs: uint16` | `[k1, v1, ...kN, vN]` → `[hash]` | Create hash |

**Example**:
```bhasa
{"নাম": "রহিম", "বয়স": 30}
```
Compiles to:
```
OpConstant 0       // Push "নাম"
OpConstant 1       // Push "রহিম"
OpConstant 2       // Push "বয়স"
OpConstant 3       // Push 30
OpHash 2           // Create hash with 2 pairs
```

---

### Type System Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpTypeCheck` | 41 | `typeIdx: uint16` | `[value]` → `[bool]` | Check if value matches type |
| `OpTypeCast` | 42 | `typeIdx: uint16` | `[value]` → `[casted]` | Cast value to type |
| `OpAssertType` | 43 | `typeIdx: uint16` | `[value]` → `[value]` | Assert type or error |

**Example**:
```bhasa
ধরি x = 5 as দশমিক;  // Type cast
```
Compiles to:
```
OpConstant 0       // Push 5
OpConstant 1       // Push type "দশমিক"
OpTypeCast 1       // Cast to float
OpSetGlobal 0      // Store as x
```

---

### Struct Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpStruct` | 44 | `defIdx: uint16` | `[f1, ...fN]` → `[struct]` | Create struct instance |
| `OpGetStructField` | 45 | None | `[struct, field]` → `[value]` | Get struct field |
| `OpSetStructField` | 46 | None | `[struct, field, value]` → `[]` | Set struct field |

**Example**:
```bhasa
ধরি ব্যক্তি = স্ট্রাক্ট {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা
};

ধরি p = ব্যক্তি{নাম: "রহিম", বয়স: 30};
p.নাম
```
Compiles to:
```
// Struct creation
OpConstant 0         // Push "রহিম"
OpConstant 1         // Push 30
OpStruct 0           // Create struct from definition 0
OpSetGlobal 0        // Store as p

// Field access
OpGetGlobal 0        // Load p
OpConstant 2         // Push field name "নাম"
OpGetStructField     // Get field value
```

---

### Enum Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpEnum` | 47 | `typeIdx: uint16, variantIdx: uint16` | `[]` → `[enum]` | Create enum variant |

**Example**:
```bhasa
ধরি দিক = গণনা { উত্তর, দক্ষিণ };
ধরি dir = দিক.উত্তর;
```
Compiles to:
```
OpEnum 0 0         // Create enum type 0, variant 0 (উত্তর)
OpSetGlobal 0      // Store as dir
```

---

### OOP Operations

#### Class Definition and Instantiation

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpClass` | 48 | `defIdx: uint16` | `[]` → `[class]` | Create class definition |
| `OpNewInstance` | 49 | `numArgs: uint8` | `[class, arg1, ...argN]` → `[instance]` | Create new instance |
| `OpDefineMethod` | 53 | `nameIdx: uint16` | `[class, method]` → `[class]` | Define method |
| `OpDefineConstructor` | 54 | `closureIdx: uint16` | `[class, constructor]` → `[class]` | Define constructor |

**Example**:
```bhasa
শ্রেণী ব্যক্তি {
    সার্বজনীন নাম: লেখা;
    
    সার্বজনীন নির্মাতা(নাম: লেখা) {
        এই.নাম = নাম;
    }
}

ধরি p = নতুন ব্যক্তি("রহিম");
```
Compiles to:
```
// Class definition
OpClass 0              // Create class from definition 0
OpClosure 1 0          // Create constructor closure
OpDefineConstructor 0  // Attach constructor
OpSetGlobal 0          // Store class as ব্যক্তি

// Instantiation
OpGetGlobal 0          // Load class ব্যক্তি
OpConstant 1           // Push "রহিম"
OpNewInstance 1        // Create instance with 1 arg
OpSetGlobal 1          // Store as p
```

#### Method Calls

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpCallMethod` | 50 | `numArgs: uint8` | `[obj, method, arg1, ...argN]` → `[result]` | Call method |

**Example**:
```bhasa
p.বলো("হ্যালো");
```
Compiles to:
```
OpGetGlobal 0        // Load p
OpConstant 1         // Push method name "বলো"
OpConstant 2         // Push "হ্যালো"
OpCallMethod 1       // Call method with 1 arg
```

#### Context References

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetThis` | 51 | None | `[]` → `[this]` | Get current object (এই) |
| `OpGetSuper` | 52 | None | `[]` → `[super]` | Get parent class (উর্ধ্ব) |

**Example**:
```bhasa
পদ্ধতি সেট_নাম(নাম: লেখা) {
    এই.নাম = নাম;
}
```
Compiles to:
```
OpGetThis            // Push this
OpConstant 0         // Push field name "নাম"
OpGetLocal 0         // Push parameter নাম
OpSetInstanceField   // Set field
```

#### Inheritance

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpInherit` | 57 | `parentIdx: uint16` | `[class]` → `[class]` | Set up inheritance |

**Example**:
```bhasa
শ্রেণী ছাত্র প্রসারিত ব্যক্তি {
    // ...
}
```
Compiles to:
```
OpClass 1              // Create ছাত্র class
OpGetGlobal 0          // Load ব্যক্তি (parent)
OpInherit 0            // Set inheritance
OpSetGlobal 1          // Store as ছাত্র
```

#### Interface Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpInterface` | 55 | `defIdx: uint16` | `[]` → `[interface]` | Create interface |
| `OpCheckInterface` | 56 | `ifaceIdx: uint16` | `[obj]` → `[bool]` | Check if implements interface |

**Example**:
```bhasa
চুক্তি যোগাযোগ {
    পদ্ধতি বলো(বার্তা: লেখা): শূন্য;
}
```
Compiles to:
```
OpInterface 0        // Create interface from definition 0
OpSetGlobal 0        // Store as যোগাযোগ
```

#### Instance Field Operations

| Opcode | Value | Operands | Stack Effect | Description |
|--------|-------|----------|--------------|-------------|
| `OpGetInstanceField` | 58 | None | `[instance, field]` → `[value]` | Get instance field |
| `OpSetInstanceField` | 59 | None | `[instance, field, value]` → `[]` | Set instance field |

---

## Operand Widths

### 0 Bytes (No Operands)

Most operations don't need operands because they work on stack values:
- Arithmetic: `OpAdd`, `OpSub`, `OpMul`, `OpDiv`, `OpMod`
- Bitwise: `OpBitAnd`, `OpBitOr`, `OpBitXor`, `OpBitNot`
- Comparison: `OpEqual`, `OpNotEqual`, `OpGreaterThan`
- Logical: `OpBang`, `OpAnd`, `OpOr`
- Unary: `OpMinus`
- Stack: `OpPop`
- Control: `OpReturn`, `OpReturnValue`
- Literals: `OpTrue`, `OpFalse`, `OpNull`
- Access: `OpIndex`, `OpGetStructField`, `OpSetStructField`
- OOP: `OpGetThis`, `OpGetSuper`

### 1 Byte (uint8: 0-255)

Used for small indices:
- `OpCall`: Number of arguments (max 255)
- `OpGetLocal`, `OpSetLocal`: Local variable index (max 255 locals)
- `OpGetBuiltin`: Builtin function index
- `OpGetFree`: Free variable index
- `OpNewInstance`: Constructor arguments
- `OpCallMethod`: Method arguments

### 2 Bytes (uint16: 0-65535)

Used for large indices:
- `OpConstant`: Constant pool index (max 65535 constants)
- `OpGetGlobal`, `OpSetGlobal`: Global variable index
- `OpJump`, `OpJumpNotTruthy`: Jump offset
- `OpArray`: Array size
- `OpHash`: Number of key-value pairs
- `OpTypeCheck`, `OpTypeCast`, `OpAssertType`: Type index
- `OpStruct`: Struct definition index
- `OpEnum`: Enum type and variant indices
- `OpClass`: Class definition index
- `OpDefineMethod`, `OpDefineConstructor`: Name/closure index
- `OpInterface`, `OpCheckInterface`: Interface index
- `OpInherit`: Parent class index

### Mixed Width

- `OpClosure`: 2-byte constant index + 1-byte free variable count

---

## Instruction Categories

### By Purpose

1. **Data Loading**: `OpConstant`, `OpTrue`, `OpFalse`, `OpNull`, `OpGetGlobal`, `OpGetLocal`, `OpGetBuiltin`, `OpGetFree`

2. **Data Storing**: `OpSetGlobal`, `OpSetLocal`, `OpPop`

3. **Arithmetic**: `OpAdd`, `OpSub`, `OpMul`, `OpDiv`, `OpMod`, `OpMinus`

4. **Bitwise**: `OpBitAnd`, `OpBitOr`, `OpBitXor`, `OpBitNot`, `OpLeftShift`, `OpRightShift`

5. **Comparison**: `OpEqual`, `OpNotEqual`, `OpGreaterThan`, `OpGreaterThanEqual`

6. **Logical**: `OpBang`, `OpAnd`, `OpOr`

7. **Control Flow**: `OpJump`, `OpJumpNotTruthy`, `OpReturn`, `OpReturnValue`

8. **Functions**: `OpCall`, `OpClosure`, `OpCurrentClosure`

9. **Collections**: `OpArray`, `OpHash`, `OpIndex`

10. **Types**: `OpTypeCheck`, `OpTypeCast`, `OpAssertType`

11. **Structs**: `OpStruct`, `OpGetStructField`, `OpSetStructField`

12. **Enums**: `OpEnum`

13. **Classes**: `OpClass`, `OpNewInstance`, `OpDefineMethod`, `OpDefineConstructor`, `OpInherit`

14. **Methods**: `OpCallMethod`, `OpGetThis`, `OpGetSuper`, `OpGetInstanceField`, `OpSetInstanceField`

15. **Interfaces**: `OpInterface`, `OpCheckInterface`

---

## API Reference

### Make

```go
func Make(op Opcode, operands ...int) []byte
```

Creates a bytecode instruction from an opcode and its operands.

**Parameters**:
- `op`: The opcode
- `operands`: Variable number of operand values

**Returns**: Byte slice containing the complete instruction

**Example**:
```go
// OpConstant with index 5
instruction := code.Make(code.OpConstant, 5)
// Result: [0x00, 0x00, 0x05]

// OpCall with 2 arguments
instruction := code.Make(code.OpCall, 2)
// Result: [0x20, 0x02]
```

---

### Lookup

```go
func Lookup(op byte) (*Definition, error)
```

Returns the definition for an opcode byte.

**Parameters**:
- `op`: The opcode byte value

**Returns**: 
- Definition pointer or nil
- Error if opcode is undefined

**Example**:
```go
def, err := code.Lookup(0x00)
// Returns: &Definition{Name: "OpConstant", OperandWidths: []int{2}}
```

---

### ReadOperands

```go
func ReadOperands(def *Definition, ins Instructions) ([]int, int)
```

Reads operand values from an instruction.

**Parameters**:
- `def`: The opcode definition
- `ins`: The instruction bytes (without the opcode)

**Returns**:
- Slice of operand values
- Number of bytes read

**Example**:
```go
ins := []byte{0x00, 0x05}  // Operands only
def := definitions[OpConstant]
operands, bytesRead := code.ReadOperands(def, ins)
// operands: [5], bytesRead: 2
```

---

### ReadUint16

```go
func ReadUint16(ins Instructions) uint16
```

Reads a 16-bit unsigned integer from instructions (big-endian).

**Example**:
```go
ins := []byte{0x00, 0x05}
value := code.ReadUint16(ins)
// value: 5
```

---

### ReadUint8

```go
func ReadUint8(ins Instructions) uint8
```

Reads an 8-bit unsigned integer from instructions.

**Example**:
```go
ins := []byte{0x02}
value := code.ReadUint8(ins)
// value: 2
```

---

### Instructions.String

```go
func (ins Instructions) String() string
```

Converts instructions to a human-readable string for debugging.

**Example**:
```go
ins := code.Instructions{
    code.Make(code.OpConstant, 1)[0],
    0x00, 0x01,
    code.Make(code.OpPop)[0],
}
fmt.Println(ins.String())

// Output:
// 0000 OpConstant 1
// 0003 OpPop
```

---

## Examples

### Example 1: Simple Arithmetic

**Bhasa Code**:
```bhasa
5 + 3 * 2
```

**Bytecode**:
```
0000 OpConstant 0    // Push 5
0003 OpConstant 1    // Push 3
0006 OpConstant 2    // Push 2
0009 OpMul           // 3 * 2 = 6
0010 OpAdd           // 5 + 6 = 11
```

**Stack States**:
```
Initial: []
After 5:   [5]
After 3:   [5, 3]
After 2:   [5, 3, 2]
After Mul: [5, 6]
After Add: [11]
```

---

### Example 2: Variable Declaration and Use

**Bhasa Code**:
```bhasa
ধরি x = 10;
ধরি y = x + 5;
```

**Bytecode**:
```
0000 OpConstant 0    // Push 10
0003 OpSetGlobal 0   // x = 10
0006 OpGetGlobal 0   // Push x
0009 OpConstant 1    // Push 5
0012 OpAdd           // x + 5
0013 OpSetGlobal 1   // y = result
```

---

### Example 3: If-Else

**Bhasa Code**:
```bhasa
যদি (x > 5) {
    10
} নাহলে {
    20
}
```

**Bytecode**:
```
0000 OpGetGlobal 0       // Push x
0003 OpConstant 0        // Push 5
0006 OpGreaterThan       // x > 5
0007 OpJumpNotTruthy 14  // Jump to else if false
0010 OpConstant 1        // Push 10
0013 OpJump 17           // Jump over else
0016 OpConstant 2        // Push 20
0019 ...                 // Continue
```

---

### Example 4: Function Definition and Call

**Bhasa Code**:
```bhasa
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
add(5, 3);
```

**Bytecode**:
```
// Function definition
0000 OpClosure 0 0       // Create closure from constant[0]
0004 OpSetGlobal 0       // Store as add

// Function call
0007 OpGetGlobal 0       // Load add
0010 OpConstant 1        // Push 5
0013 OpConstant 2        // Push 3
0016 OpCall 2            // Call with 2 args

// Inside the function (compiled separately)
0000 OpGetLocal 0        // Get parameter a
0002 OpGetLocal 1        // Get parameter b
0004 OpAdd               // a + b
0005 OpReturnValue       // Return result
```

---

### Example 5: Array Operations

**Bhasa Code**:
```bhasa
ধরি arr = [1, 2, 3];
arr[1]
```

**Bytecode**:
```
// Array creation
0000 OpConstant 0        // Push 1
0003 OpConstant 1        // Push 2
0006 OpConstant 2        // Push 3
0009 OpArray 3           // Create array with 3 elements
0012 OpSetGlobal 0       // Store as arr

// Array indexing
0015 OpGetGlobal 0       // Load arr
0018 OpConstant 3        // Push 1 (index)
0021 OpIndex             // Get arr[1]
```

---

### Example 6: Class and Methods

**Bhasa Code**:
```bhasa
শ্রেণী ব্যক্তি {
    সার্বজনীন নাম: লেখা;
    
    সার্বজনীন নির্মাতা(n: লেখা) {
        এই.নাম = n;
    }
    
    সার্বজনীন পদ্ধতি বলো(): শূন্য {
        দেখাও(এই.নাম);
    }
}

ধরি p = নতুন ব্যক্তি("রহিম");
p.বলো();
```

**Bytecode**:
```
// Class definition
0000 OpClass 0              // Create class
0003 OpClosure 1 0          // Constructor closure
0007 OpDefineConstructor 0  // Attach constructor
0010 OpClosure 2 0          // Method closure
0014 OpDefineMethod 0       // Attach method "বলো"
0017 OpSetGlobal 0          // Store as ব্যক্তি

// Instance creation
0020 OpGetGlobal 0          // Load class ব্যক্তি
0023 OpConstant 3           // Push "রহিম"
0026 OpNewInstance 1        // Create instance
0029 OpSetGlobal 1          // Store as p

// Method call
0032 OpGetGlobal 1          // Load p
0035 OpConstant 4           // Method name "বলো"
0038 OpCallMethod 0         // Call with 0 args
```

---

## Debugging

### Disassembling Instructions

Use the `String()` method to see human-readable bytecode:

```go
bytecode := compiler.Bytecode()
fmt.Println(bytecode.Instructions.String())
```

**Output**:
```
0000 OpConstant 0
0003 OpConstant 1
0006 OpAdd
0007 OpPop
```

### Understanding Jump Offsets

Jump instructions use absolute offsets:

```
0007 OpJumpNotTruthy 14  // Jump TO position 14, not BY 14 bytes
```

### Tracing Execution

Add print statements in the VM's fetch-decode-execute loop:

```go
fmt.Printf("IP: %d, Opcode: %s\n", vm.ip, code.Lookup(vm.currentInstruction()))
fmt.Printf("Stack: %v\n", vm.stack)
```

### Common Issues

1. **Wrong operand count**: Ensure `Make()` receives correct number of operands
2. **Wrong operand width**: Check definition's `OperandWidths`
3. **Stack underflow**: Missing OpPop or too many pops
4. **Jump offsets**: Calculate absolute positions, not relative
5. **Constant pool indices**: Ensure constants are added before use

---

## Performance Considerations

### Instruction Size

Smaller instructions = faster decoding:
- Use `uint8` for local variables (max 255 locals per scope)
- Use `uint16` for globals and constants (max 65535)

### Stack Operations

Minimize stack manipulation:
- Avoid unnecessary `OpPop` followed by `OpGet`
- Reorder operations to reduce stack depth

### Jump Optimization

- Use `OpJumpNotTruthy` instead of `OpBang` + `OpJumpNotTruthy`
- Optimize short-circuit evaluation for `&&` and `||`

---

## Future Enhancements

Possible future opcodes:

1. **Optimization Opcodes**:
   - `OpIncrement`: Increment local variable in place
   - `OpDecrement`: Decrement local variable
   - `OpAddConstant`: Add constant to top of stack

2. **Advanced OOP**:
   - `OpProperty`: Define property with getter/setter
   - `OpOperatorOverload`: Define operator overloading

3. **Concurrency**:
   - `OpSpawn`: Spawn goroutine
   - `OpChannel`: Create channel
   - `OpSend`, `OpReceive`: Channel operations

4. **Exception Handling**:
   - `OpTry`, `OpCatch`, `OpFinally`
   - `OpThrow`: Throw exception

---

## See Also

- [Compiler Documentation](../../compiler/docs/) - How AST is compiled to bytecode
- [VM Documentation](../../vm/docs/) - How bytecode is executed
- [AST Documentation](../../ast/docs/) - AST structure
- [Quick Reference](./quick-reference.md) - Opcode lookup tables

