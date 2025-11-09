# Compilation Examples

This document provides detailed step-by-step examples of how Bhasa code is compiled to bytecode.

## Table of Contents

1. [Simple Expressions](#simple-expressions)
2. [Variables](#variables)
3. [Control Flow](#control-flow)
4. [Functions](#functions)
5. [Closures](#closures)
6. [Loops](#loops)
7. [Data Structures](#data-structures)
8. [Object-Oriented Programming](#object-oriented-programming)

---

## Simple Expressions

### Example 1: Arithmetic

**Bhasa Code**:
```bhasa
5 + 3
```

**AST**:
```
InfixExpression
├── Left: IntegerLiteral(5)
├── Operator: "+"
└── Right: IntegerLiteral(3)
```

**Compilation Steps**:

| Step | Action | Constants | Bytecode |
|------|--------|-----------|----------|
| 1 | Compile left (5) | [5] | OpConstant 0 |
| 2 | Compile right (3) | [5, 3] | OpConstant 1 |
| 3 | Emit OpAdd | [5, 3] | OpAdd |
| 4 | Pop result (stmt) | [5, 3] | OpPop |

**Final Bytecode**:
```
0000 OpConstant 0    // Push 5
0003 OpConstant 1    // Push 3
0006 OpAdd           // 5 + 3
0007 OpPop           // Discard (expression statement)
```

**Constants**: `[5, 3]`

---

### Example 2: Complex Expression

**Bhasa Code**:
```bhasa
(10 + 5) * 2
```

**AST**:
```
InfixExpression(*)
├── Left: InfixExpression(+)
│   ├── Left: IntegerLiteral(10)
│   ├── Operator: "+"
│   └── Right: IntegerLiteral(5)
├── Operator: "*"
└── Right: IntegerLiteral(2)
```

**Compilation Steps**:

| Step | Action | Stack | Bytecode |
|------|--------|-------|----------|
| 1 | Compile 10 | [10] | OpConstant 0 |
| 2 | Compile 5 | [10, 5] | OpConstant 1 |
| 3 | OpAdd | [15] | OpAdd |
| 4 | Compile 2 | [15, 2] | OpConstant 2 |
| 5 | OpMul | [30] | OpMul |

**Final Bytecode**:
```
OpConstant 0
OpConstant 1
OpAdd
OpConstant 2
OpMul
OpPop
```

---

## Variables

### Example 3: Global Variable

**Bhasa Code**:
```bhasa
ধরি x = 10;
ধরি y = 20;
x + y;
```

**Symbol Table Evolution**:

```
After "ধরি x = 10;":
┌─────────────────┐
│ Global          │
├─────────────────┤
│ x → GLOBAL:0    │
└─────────────────┘

After "ধরি y = 20;":
┌─────────────────┐
│ Global          │
├─────────────────┤
│ x → GLOBAL:0    │
│ y → GLOBAL:1    │
└─────────────────┘
```

**Bytecode**:
```
// ধরি x = 10;
0000 OpConstant 0        // Push 10
0003 OpSetGlobal 0       // x = 10

// ধরি y = 20;
0006 OpConstant 1        // Push 20
0009 OpSetGlobal 1       // y = 20

// x + y;
0012 OpGetGlobal 0       // Load x
0015 OpGetGlobal 1       // Load y
0018 OpAdd               // Add
0019 OpPop               // Discard
```

**Constants**: `[10, 20]`

---

### Example 4: Variable Assignment

**Bhasa Code**:
```bhasa
ধরি x = 5;
x = x + 10;
```

**Bytecode**:
```
// ধরি x = 5;
0000 OpConstant 0        // Push 5
0003 OpSetGlobal 0       // x = 5

// x = x + 10;
0006 OpGetGlobal 0       // Load x
0009 OpConstant 1        // Push 10
0012 OpAdd               // x + 10
0013 OpSetGlobal 0       // x = result
```

---

### Example 5: Type Annotation

**Bhasa Code**:
```bhasa
ধরি x: পূর্ণসংখ্যা = 5;
```

**Symbol Table**:
```
┌─────────────────────────────────┐
│ Global                          │
├─────────────────────────────────┤
│ x → GLOBAL:0                    │
│     TypeAnnot: পূর্ণসংখ্যা      │
└─────────────────────────────────┘
```

**Bytecode**:
```
0000 OpConstant 0        // Push 5
0003 OpAssertType 1      // Check type (পূর্ণসংখ্যা)
0006 OpSetGlobal 0       // x = 5
```

**Constants**: `[5, "পূর্ণসংখ্যা"]`

---

## Control Flow

### Example 6: If-Else Expression

**Bhasa Code**:
```bhasa
যদি (x > 5) {
    10
} নাহলে {
    20
}
```

**Compilation Steps**:

| Step | Action | Bytecode Pos |
|------|--------|--------------|
| 1 | Compile condition | 0-6 |
| 2 | Emit jump (placeholder) | 7 |
| 3 | Compile consequence | 10 |
| 4 | Emit jump over alternative | 13 |
| 5 | **Patch first jump** → 16 | - |
| 6 | Compile alternative | 16 |
| 7 | **Patch second jump** → 19 | - |

**Bytecode**:
```
0000 OpGetGlobal 0       // Load x
0003 OpConstant 0        // Push 5
0006 OpGreaterThan       // x > 5
0007 OpJumpNotTruthy 16  // Jump to else if false [PATCHED]
0010 OpConstant 1        // Push 10
0013 OpJump 19           // Skip else [PATCHED]
0016 OpConstant 2        // Push 20
0019 ...                 // Continue
```

**Jump Patching**:
```
Initially:
0007 OpJumpNotTruthy 9999  // Placeholder

After patching:
0007 OpJumpNotTruthy 16    // Points to alternative
```

---

### Example 7: Nested If

**Bhasa Code**:
```bhasa
যদি (x > 5) {
    যদি (y > 10) {
        1
    } নাহলে {
        2
    }
} নাহলে {
    3
}
```

**Bytecode**:
```
// Outer condition
0000 OpGetGlobal 0       // x
0003 OpConstant 0        // 5
0006 OpGreaterThan
0007 OpJumpNotTruthy 28  // To outer else

// Inner if (consequence of outer)
0010 OpGetGlobal 1       // y
0013 OpConstant 1        // 10
0016 OpGreaterThan
0017 OpJumpNotTruthy 24  // To inner else
0020 OpConstant 2        // 1
0023 OpJump 27           // Skip inner else
0026 OpConstant 3        // 2
0029 OpJump 32           // Skip outer else

// Outer else
0032 OpConstant 4        // 3
```

---

## Functions

### Example 8: Simple Function

**Bhasa Code**:
```bhasa
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
```

**Symbol Tables**:

```
Global:
┌─────────────────┐
│ add → GLOBAL:0  │
└─────────────────┘

Function Scope (during compilation):
┌─────────────────┐
│ Outer: Global   │
├─────────────────┤
│ a → LOCAL:0     │
│ b → LOCAL:1     │
└─────────────────┘
```

**Function Bytecode** (stored in constants[0]):
```
0000 OpGetLocal 0        // Get a
0002 OpGetLocal 1        // Get b
0004 OpAdd               // a + b
0005 OpReturnValue       // Return result
```

**Main Bytecode**:
```
0000 OpClosure 0 0       // Create closure (fn idx 0, 0 free vars)
0004 OpSetGlobal 0       // Store as add
```

**Constants**: `[CompiledFunction{...}]`

---

### Example 9: Function with Local Variables

**Bhasa Code**:
```bhasa
ফাংশন(x) {
    ধরি y = x * 2;
    ধরি z = y + 10;
    z
}
```

**Symbol Table**:
```
Function Scope:
┌─────────────────┐
│ Outer: Global   │
├─────────────────┤
│ x → LOCAL:0     │  (parameter)
│ y → LOCAL:1     │  (local)
│ z → LOCAL:2     │  (local)
└─────────────────┘
```

**Function Bytecode**:
```
// ধরি y = x * 2;
0000 OpGetLocal 0        // Get x
0002 OpConstant 0        // Push 2
0005 OpMul               // x * 2
0006 OpSetLocal 1        // y = result

// ধরি z = y + 10;
0008 OpGetLocal 1        // Get y
0010 OpConstant 1        // Push 10
0013 OpAdd               // y + 10
0014 OpSetLocal 2        // z = result

// z (implicit return)
0016 OpGetLocal 2        // Get z
0018 OpReturnValue       // Return z
```

---

## Closures

### Example 10: Simple Closure

**Bhasa Code**:
```bhasa
ধরি makeAdder = ফাংশন(x) {
    ফাংশন(y) {
        ফেরত x + y;
    }
};

ধরি add5 = makeAdder(5);
add5(3);
```

**Symbol Tables**:

```
Outer Function (makeAdder):
┌─────────────────┐
│ Outer: Global   │
├─────────────────┤
│ x → LOCAL:0     │
└─────────────────┘

Inner Function:
┌─────────────────┐
│ Outer: makeAdder│
├─────────────────┤
│ y → LOCAL:0     │
│ x → FREE:0      │  ← Captured from outer
└─────────────────┘
FreeSymbols: [x]
```

**Outer Function Bytecode**:
```
0000 OpGetLocal 0        // Load x (for closure)
0002 OpClosure 0 1       // Create inner function with 1 free var
0006 OpReturnValue       // Return inner function
```

**Inner Function Bytecode** (constants[0]):
```
0000 OpGetFree 0         // Get x (free variable)
0002 OpGetLocal 0        // Get y (parameter)
0004 OpAdd               // x + y
0005 OpReturnValue       // Return result
```

**Main Bytecode**:
```
// ধরি makeAdder = ...
0000 OpClosure 1 0       // Create makeAdder (no free vars)
0004 OpSetGlobal 0       // Store as makeAdder

// ধরি add5 = makeAdder(5);
0007 OpGetGlobal 0       // Load makeAdder
0010 OpConstant 0        // Push 5
0013 OpCall 1            // Call makeAdder(5)
0015 OpSetGlobal 1       // Store as add5

// add5(3);
0018 OpGetGlobal 1       // Load add5
0021 OpConstant 1        // Push 3
0024 OpCall 1            // Call add5(3)
0026 OpPop               // Discard result
```

---

### Example 11: Multiple Free Variables

**Bhasa Code**:
```bhasa
ফাংশন(a, b) {
    ফাংশন(c) {
        a + b + c
    }
}
```

**Symbol Tables**:

```
Inner Function:
┌─────────────────┐
│ c → LOCAL:0     │
│ a → FREE:0      │  ← Captured
│ b → FREE:1      │  ← Captured
└─────────────────┘
FreeSymbols: [a, b]
```

**Inner Function Bytecode**:
```
0000 OpGetFree 0         // Get a
0002 OpGetFree 1         // Get b
0004 OpAdd               // a + b
0005 OpGetLocal 0        // Get c
0007 OpAdd               // (a + b) + c
0008 OpReturnValue
```

**Outer Function Bytecode**:
```
0000 OpGetLocal 0        // Load a for closure
0002 OpGetLocal 1        // Load b for closure
0004 OpClosure 0 2       // Create closure with 2 free vars
0008 OpReturnValue
```

---

## Loops

### Example 12: While Loop

**Bhasa Code**:
```bhasa
ধরি i = 0;
যতক্ষণ (i < 10) {
    i = i + 1;
}
```

**Compilation Steps**:

| Step | Action | Position |
|------|--------|----------|
| 1 | Mark loop start | 6 |
| 2 | Push loop context | - |
| 3 | Compile condition | 6-11 |
| 4 | Emit exit jump | 12 |
| 5 | Compile body | 15-24 |
| 6 | Emit jump to start | 25 |
| 7 | Patch exit jump | - |
| 8 | Pop loop context | - |

**Bytecode**:
```
// ধরি i = 0;
0000 OpConstant 0        // Push 0
0003 OpSetGlobal 0       // i = 0

// Loop start (position 6)
0006 OpGetGlobal 0       // Load i
0009 OpConstant 1        // Push 10
0012 OpGreaterThan       // Reversed: i < 10
0013 OpJumpNotTruthy 28  // Exit if false [PATCHED]

// Loop body
0016 OpGetGlobal 0       // Load i
0019 OpConstant 2        // Push 1
0022 OpAdd               // i + 1
0023 OpSetGlobal 0       // i = result
0026 OpJump 6            // Jump to loop start

// After loop (position 28)
0029 OpNull              // Loop result
```

---

### Example 13: For Loop with Break

**Bhasa Code**:
```bhasa
পর্যন্ত (ধরি i = 0; i < 10; i = i + 1) {
    যদি (i == 5) {
        বিরতি;
    }
}
```

**Loop Context**:
```go
LoopContext{
    loopStart: 3,
    breakPositions: [23],  // Position of break jump
    contPositions: [],
}
```

**Bytecode**:
```
// Initialization
0000 OpConstant 0        // Push 0
0003 OpSetLocal 0        // i = 0

// Loop start (position 3)
0005 OpGetLocal 0        // Load i
0007 OpConstant 1        // Push 10
0010 OpGreaterThan       // i < 10
0011 OpJumpNotTruthy 35  // Exit if false [PATCHED]

// Body
0014 OpGetLocal 0        // Load i
0016 OpConstant 2        // Push 5
0019 OpEqual             // i == 5
0020 OpJumpNotTruthy 26  // Skip break if false

// Break
0023 OpJump 9999         // Will be patched to 35 [PATCHED]

// After if
0026 // Continue body...

// Increment
0029 OpGetLocal 0        // Load i
0031 OpConstant 3        // Push 1
0034 OpAdd               // i + 1
0035 OpSetLocal 0        // i = result

0037 OpJump 3            // Jump to loop start

// After loop (position 35) ← Break jumps here
0040 OpNull
```

---

## Data Structures

### Example 14: Array Literal

**Bhasa Code**:
```bhasa
[1, 2, 3, 4, 5]
```

**Bytecode**:
```
0000 OpConstant 0        // Push 1
0003 OpConstant 1        // Push 2
0006 OpConstant 2        // Push 3
0009 OpConstant 3        // Push 4
0012 OpConstant 4        // Push 5
0015 OpArray 5           // Create array with 5 elements
0018 OpPop               // Discard (expression statement)
```

**Stack Evolution**:
```
[]
[1]
[1, 2]
[1, 2, 3]
[1, 2, 3, 4]
[1, 2, 3, 4, 5]
[[1, 2, 3, 4, 5]]  ← After OpArray
```

---

### Example 15: Hash Literal

**Bhasa Code**:
```bhasa
{"নাম": "রহিম", "বয়স": 30}
```

**Note**: Keys are sorted for deterministic compilation

**Bytecode**:
```
0000 OpConstant 0        // Push "নাম"
0003 OpConstant 1        // Push "রহিম"
0006 OpConstant 2        // Push "বয়স"
0009 OpConstant 3        // Push 30
0012 OpHash 4            // Create hash (4 = 2 pairs * 2)
0015 OpPop
```

**Constants**: `["নাম", "রহিম", "বয়স", 30]`

---

### Example 16: Nested Structures

**Bhasa Code**:
```bhasa
[[1, 2], [3, 4]]
```

**Bytecode**:
```
// Inner array [1, 2]
0000 OpConstant 0        // 1
0003 OpConstant 1        // 2
0006 OpArray 2           // Create [1, 2]

// Inner array [3, 4]
0009 OpConstant 2        // 3
0012 OpConstant 3        // 4
0015 OpArray 2           // Create [3, 4]

// Outer array
0018 OpArray 2           // Create [[1,2], [3,4]]
0021 OpPop
```

---

## Object-Oriented Programming

### Example 17: Class Definition

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
```

**Symbol Tables**:

```
Constructor Scope:
┌─────────────────┐
│ এই → LOCAL:0    │  (this)
│ n → LOCAL:1     │  (parameter)
└─────────────────┘

Method Scope:
┌─────────────────┐
│ এই → LOCAL:0    │  (this)
└─────────────────┘
```

**Constructor Bytecode**:
```
0000 OpGetLocal 0        // Get this (এই)
0002 OpConstant 0        // Push field name "নাম"
0005 OpGetLocal 1        // Get parameter n
0007 OpSetInstanceField  // this.নাম = n
0008 OpGetLocal 0        // Load this for return
0010 OpReturnValue       // Return this
```

**Method Bytecode**:
```
0000 OpGetLocal 0        // Get this (এই)
0002 OpConstant 0        // Push field name "নাম"
0005 OpGetInstanceField  // Get this.নাম
0006 OpGetBuiltin 5      // Get দেখাও (print)
0008 OpCall 1            // Call দেখাও(this.নাম)
0010 OpNull              // Return null
0011 OpReturnValue
```

**Main Bytecode**:
```
0000 OpClass 0           // Create class from definition
0003 OpClosure 1 0       // Constructor closure
0007 OpDefineConstructor 0
0010 OpClosure 2 0       // Method closure
0014 OpDefineMethod 1    // Define method "বলো"
0017 OpSetGlobal 0       // Store class as ব্যক্তি
```

---

### Example 18: Instance Creation

**Bhasa Code**:
```bhasa
ধরি p = নতুন ব্যক্তি("রহিম");
p.বলো();
```

**Bytecode**:
```
// Instance creation
0000 OpConstant 0        // Push "রহিম"
0003 OpGetGlobal 0       // Load class ব্যক্তি
0006 OpNewInstance 1     // Create instance with 1 arg
0008 OpSetGlobal 1       // Store as p

// Method call
0011 OpGetGlobal 1       // Load p
0014 OpConstant 1        // Push method name "বলো"
0017 OpCallMethod 0      // Call method with 0 args
0019 OpPop
```

**Stack Evolution**:
```
Instance creation:
["রহিম"]
["রহিম", <class>]
[<instance>]
[]

Method call:
[<instance>]
[<instance>, "বলো"]
[<result>]
[]
```

---

### Example 19: Inheritance

**Bhasa Code**:
```bhasa
শ্রেণী প্রাণী {
    সার্বজনীন পদ্ধতি বলো(): শূন্য {
        দেখাও("প্রাণী");
    }
}

শ্রেণী কুকুর প্রসারিত প্রাণী {
    পুনর্সংজ্ঞা পদ্ধতি বলো(): শূন্য {
        উর্ধ্ব.বলো();
        দেখাও("ঘেউ ঘেউ");
    }
}
```

**Override Method Bytecode**:
```
0000 OpGetSuper          // Get parent class (উর্ধ্ব)
0001 OpConstant 0        // Push method name "বলো"
0004 OpCallMethod 0      // Call parent.বলো()
0006 OpPop               // Discard result
0007 OpConstant 1        // Push "ঘেউ ঘেউ"
0010 OpGetBuiltin 5      // Get দেখাও
0012 OpCall 1            // Call দেখাও("ঘেউ ঘেউ")
0014 OpNull
0015 OpReturnValue
```

---

## Summary

This document showed:

✅ **Expression compilation**: Stack-based evaluation  
✅ **Variable management**: Symbol tables and scopes  
✅ **Control flow**: Jump patching  
✅ **Function compilation**: Scope management  
✅ **Closures**: Free variable capture  
✅ **Loops**: Break/continue handling  
✅ **Data structures**: Arrays and hashes  
✅ **OOP**: Classes, methods, inheritance  

## See Also

- [Compiler Documentation](./compiler-documentation.md)
- [Symbol Table Documentation](./symbol-table-documentation.md)
- [Quick Reference](./quick-reference.md)
- [Bytecode Documentation](../../code/docs/)

