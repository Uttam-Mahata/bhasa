# Bhasa Bytecode Instruction Examples

This document provides detailed visual examples of how bytecode instructions work, including stack states and step-by-step execution.

## 📖 Table of Contents

1. [Arithmetic Operations](#arithmetic-operations)
2. [Variable Operations](#variable-operations)
3. [Control Flow](#control-flow)
4. [Functions](#functions)
5. [Collections](#collections)
6. [Type System](#type-system)
7. [Structs](#structs)
8. [Object-Oriented Programming](#object-oriented-programming)
9. [Complex Examples](#complex-examples)

---

## Arithmetic Operations

### Example 1: Simple Addition

**Bhasa Code**:
```bhasa
5 + 3
```

**Bytecode**:
```
0000 OpConstant 0    // Load constant 5
0003 OpConstant 1    // Load constant 3
0006 OpAdd           // Add them
```

**Execution Steps**:

| Step | Instruction | Stack State | Description |
|------|-------------|-------------|-------------|
| 0 | *Initial* | `[]` | Empty stack |
| 1 | `OpConstant 0` | `[5]` | Push 5 |
| 2 | `OpConstant 1` | `[5, 3]` | Push 3 |
| 3 | `OpAdd` | `[8]` | Pop 5 and 3, push 8 |

**Stack Visualization**:
```
Step 0:          Step 1:          Step 2:          Step 3:
                                  [3]              [8]
                 [5]              [5]              
-------          -------          -------          -------
Empty            After OpConstant After OpConstant After OpAdd
                 5                3                
```

---

### Example 2: Complex Expression

**Bhasa Code**:
```bhasa
(5 + 3) * 2
```

**Bytecode**:
```
0000 OpConstant 0    // Load 5
0003 OpConstant 1    // Load 3
0006 OpAdd           // 5 + 3
0007 OpConstant 2    // Load 2
0010 OpMul           // (5+3) * 2
```

**Execution Steps**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 0 | *Initial* | `[]` | |
| 1 | `OpConstant 0` | `[5]` | Push 5 |
| 2 | `OpConstant 1` | `[5, 3]` | Push 3 |
| 3 | `OpAdd` | `[8]` | Pop both, push 8 |
| 4 | `OpConstant 2` | `[8, 2]` | Push 2 |
| 5 | `OpMul` | `[16]` | Pop both, push 16 |

---

### Example 3: Unary Negation

**Bhasa Code**:
```bhasa
-5 + 3
```

**Bytecode**:
```
0000 OpConstant 0    // Load 5
0003 OpMinus         // Negate it
0004 OpConstant 1    // Load 3
0007 OpAdd           // Add
```

**Execution Steps**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpConstant 0` | `[5]` | Push 5 |
| 2 | `OpMinus` | `[-5]` | Pop 5, push -5 |
| 3 | `OpConstant 1` | `[-5, 3]` | Push 3 |
| 4 | `OpAdd` | `[-2]` | Pop -5 and 3, push -2 |

---

## Variable Operations

### Example 4: Global Variable Declaration

**Bhasa Code**:
```bhasa
ধরি x = 10;
```

**Bytecode**:
```
0000 OpConstant 0    // Load 10
0003 OpSetGlobal 0   // Store in global[0] (x)
```

**Execution Steps**:

| Step | Instruction | Stack | Globals | Description |
|------|-------------|-------|---------|-------------|
| 0 | *Initial* | `[]` | `[]` | |
| 1 | `OpConstant 0` | `[10]` | `[]` | Push 10 |
| 2 | `OpSetGlobal 0` | `[]` | `[10]` | Pop, store as x |

---

### Example 5: Variable Access and Update

**Bhasa Code**:
```bhasa
ধরি x = 10;
x = x + 5;
```

**Bytecode**:
```
0000 OpConstant 0    // Load 10
0003 OpSetGlobal 0   // x = 10
0006 OpGetGlobal 0   // Load x
0009 OpConstant 1    // Load 5
0012 OpAdd           // x + 5
0013 OpSetGlobal 0   // x = result
```

**Execution Steps**:

| Step | Instruction | Stack | Globals | Description |
|------|-------------|-------|---------|-------------|
| 1 | `OpConstant 0` | `[10]` | `[]` | Push 10 |
| 2 | `OpSetGlobal 0` | `[]` | `[x=10]` | Store x |
| 3 | `OpGetGlobal 0` | `[10]` | `[x=10]` | Load x |
| 4 | `OpConstant 1` | `[10, 5]` | `[x=10]` | Push 5 |
| 5 | `OpAdd` | `[15]` | `[x=10]` | Add |
| 6 | `OpSetGlobal 0` | `[]` | `[x=15]` | Update x |

---

### Example 6: Local Variables in Function

**Bhasa Code**:
```bhasa
ফাংশন() {
    ধরি x = 5;
    ধরি y = 10;
    x + y
}
```

**Bytecode** (function body):
```
0000 OpConstant 0    // Load 5
0003 OpSetLocal 0    // x = 5 (local slot 0)
0005 OpConstant 1    // Load 10
0008 OpSetLocal 1    // y = 10 (local slot 1)
0010 OpGetLocal 0    // Load x
0012 OpGetLocal 1    // Load y
0014 OpAdd           // x + y
0015 OpReturnValue   // Return result
```

**Execution Steps**:

| Step | Instruction | Stack | Locals | Description |
|------|-------------|-------|--------|-------------|
| 1 | `OpConstant 0` | `[5]` | `[]` | Push 5 |
| 2 | `OpSetLocal 0` | `[]` | `[5]` | x = 5 |
| 3 | `OpConstant 1` | `[10]` | `[5]` | Push 10 |
| 4 | `OpSetLocal 1` | `[]` | `[5, 10]` | y = 10 |
| 5 | `OpGetLocal 0` | `[5]` | `[5, 10]` | Load x |
| 6 | `OpGetLocal 1` | `[5, 10]` | `[5, 10]` | Load y |
| 7 | `OpAdd` | `[15]` | `[5, 10]` | Add |
| 8 | `OpReturnValue` | `[15]` | - | Return |

---

## Control Flow

### Example 7: Simple If Statement

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
0000 OpGetGlobal 0       // Load x
0003 OpConstant 0        // Load 5
0006 OpGreaterThan       // x > 5
0007 OpJumpNotTruthy 14  // Jump to else if false
0010 OpConstant 1        // Load 10 (consequence)
0013 OpJump 17           // Jump over else
0016 OpConstant 2        // Load 20 (alternative)
0019 ...                 // Continue
```

**Execution Steps (x = 10)**:

| Step | Instruction | Stack | PC | Description |
|------|-------------|-------|----|----|
| 1 | `OpGetGlobal 0` | `[10]` | 0 | Load x |
| 2 | `OpConstant 0` | `[10, 5]` | 3 | Load 5 |
| 3 | `OpGreaterThan` | `[true]` | 6 | 10 > 5 |
| 4 | `OpJumpNotTruthy 14` | `[]` | 7 | Don't jump (true) |
| 5 | `OpConstant 1` | `[10]` | 10 | Load 10 |
| 6 | `OpJump 17` | `[10]` | 13 | Jump over else |
| 7 | *Skipped* | - | 19 | Continue |

**Execution Steps (x = 3)**:

| Step | Instruction | Stack | PC | Description |
|------|-------------|-------|----|----|
| 1 | `OpGetGlobal 0` | `[3]` | 0 | Load x |
| 2 | `OpConstant 0` | `[3, 5]` | 3 | Load 5 |
| 3 | `OpGreaterThan` | `[false]` | 6 | 3 > 5 |
| 4 | `OpJumpNotTruthy 14` | `[]` | 7 | **Jump to 14** |
| 5 | *Skipped* | - | 14 | At else |
| 6 | `OpConstant 2` | `[20]` | 16 | Load 20 |
| 7 | ... | `[20]` | 19 | Continue |

---

### Example 8: While Loop

**Bhasa Code**:
```bhasa
ধরি i = 0;
যতক্ষণ (i < 3) {
    i = i + 1;
}
```

**Bytecode**:
```
0000 OpConstant 0        // Load 0
0003 OpSetGlobal 0       // i = 0
0006 OpGetGlobal 0       // Load i (loop start)
0009 OpConstant 1        // Load 3
0012 OpGreaterThan       // Reverse: i < 3 becomes 3 > i
0013 OpJumpNotTruthy 28  // Exit loop if false
0016 OpGetGlobal 0       // Load i
0019 OpConstant 2        // Load 1
0022 OpAdd               // i + 1
0023 OpSetGlobal 0       // i = result
0026 OpJump 6            // Jump back to loop start
0029 ...                 // After loop
```

**Execution (3 iterations)**:

| Iter | After Condition | Stack | i | Action |
|------|----------------|-------|---|--------|
| 0 | - | `[]` | 0 | Initial |
| 1 | `[true]` | - | 0 | Enter loop |
| 1 | - | `[]` | 1 | i incremented |
| 2 | `[true]` | - | 1 | Continue |
| 2 | - | `[]` | 2 | i incremented |
| 3 | `[true]` | - | 2 | Continue |
| 3 | - | `[]` | 3 | i incremented |
| 4 | `[false]` | `[]` | 3 | Exit loop |

---

## Functions

### Example 9: Function Definition and Call

**Bhasa Code**:
```bhasa
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
add(5, 3);
```

**Main Bytecode**:
```
0000 OpClosure 0 0       // Create closure from constant[0]
0004 OpSetGlobal 0       // Store as add
0007 OpGetGlobal 0       // Load add
0010 OpConstant 1        // Push 5
0013 OpConstant 2        // Push 3
0016 OpCall 2            // Call with 2 args
```

**Function Bytecode** (compiled into constant[0]):
```
0000 OpGetLocal 0        // Load parameter a
0002 OpGetLocal 1        // Load parameter b
0004 OpAdd               // a + b
0005 OpReturnValue       // Return result
```

**Execution Steps**:

| Step | Instruction | Stack | Globals | Description |
|------|-------------|-------|---------|-------------|
| 1 | `OpClosure 0 0` | `[<fn>]` | `[]` | Create function |
| 2 | `OpSetGlobal 0` | `[]` | `[add=<fn>]` | Store as add |
| 3 | `OpGetGlobal 0` | `[<fn>]` | `[add=<fn>]` | Load add |
| 4 | `OpConstant 1` | `[<fn>, 5]` | `[add=<fn>]` | Push arg 5 |
| 5 | `OpConstant 2` | `[<fn>, 5, 3]` | `[add=<fn>]` | Push arg 3 |
| 6 | `OpCall 2` | *Call* | - | Execute function |

**Inside Function Call**:

| Step | Instruction | Stack | Locals | Description |
|------|-------------|-------|--------|-------------|
| 1 | *Setup* | `[]` | `[5, 3]` | Parameters loaded |
| 2 | `OpGetLocal 0` | `[5]` | `[5, 3]` | Load a |
| 3 | `OpGetLocal 1` | `[5, 3]` | `[5, 3]` | Load b |
| 4 | `OpAdd` | `[8]` | `[5, 3]` | Add |
| 5 | `OpReturnValue` | `[8]` | - | Return to caller |

**After Return**:

| Step | Stack | Description |
|------|-------|-------------|
| 7 | `[8]` | Result on stack |

---

### Example 10: Closure with Free Variables

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

**Outer Function** (makeAdder):
```
0000 OpGetLocal 0        // Get parameter x
0002 OpClosure 0 1       // Create closure with 1 free var (x)
0006 OpReturnValue       // Return inner function
```

**Inner Function** (anonymous):
```
0000 OpGetFree 0         // Get free variable x
0002 OpGetLocal 0        // Get parameter y
0004 OpAdd               // x + y
0005 OpReturnValue       // Return result
```

**Execution**:

| Call | Free Vars | Locals | Stack | Result |
|------|-----------|--------|-------|--------|
| `makeAdder(5)` | `[]` | `[5]` | `[<closure>]` | Closure with x=5 |
| `add5(3)` | `[x=5]` | `[y=3]` | `[8]` | 5 + 3 = 8 |

---

## Collections

### Example 11: Array Creation and Access

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
0018 OpConstant 3        // Push index 1
0021 OpIndex             // Get arr[1]
```

**Execution Steps (Array Creation)**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpConstant 0` | `[1]` | Push 1 |
| 2 | `OpConstant 1` | `[1, 2]` | Push 2 |
| 3 | `OpConstant 2` | `[1, 2, 3]` | Push 3 |
| 4 | `OpArray 3` | `[[1,2,3]]` | Pop 3, create array |
| 5 | `OpSetGlobal 0` | `[]` | Store as arr |

**Execution Steps (Array Access)**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpGetGlobal 0` | `[[1,2,3]]` | Load arr |
| 2 | `OpConstant 3` | `[[1,2,3], 1]` | Push index |
| 3 | `OpIndex` | `[2]` | Get arr[1] = 2 |

---

### Example 12: Hash Creation and Access

**Bhasa Code**:
```bhasa
ধরি dict = {"নাম": "রহিম", "বয়স": 30};
dict["নাম"]
```

**Bytecode**:
```
// Hash creation
0000 OpConstant 0        // Push "নাম"
0003 OpConstant 1        // Push "রহিম"
0006 OpConstant 2        // Push "বয়স"
0009 OpConstant 3        // Push 30
0012 OpHash 2            // Create hash with 2 pairs
0015 OpSetGlobal 0       // Store as dict

// Hash access
0018 OpGetGlobal 0       // Load dict
0021 OpConstant 0        // Push "নাম"
0024 OpIndex             // Get dict["নাম"]
```

**Execution Steps (Hash Creation)**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpConstant 0` | `["নাম"]` | Push key |
| 2 | `OpConstant 1` | `["নাম", "রহিম"]` | Push value |
| 3 | `OpConstant 2` | `["নাম", "রহিম", "বয়স"]` | Push key |
| 4 | `OpConstant 3` | `["নাম", "রহিম", "বয়স", 30]` | Push value |
| 5 | `OpHash 2` | `[{...}]` | Pop 4, create hash |
| 6 | `OpSetGlobal 0` | `[]` | Store |

---

## Type System

### Example 13: Type Casting

**Bhasa Code**:
```bhasa
ধরি x = 5;
ধরি y = x as দশমিক;
```

**Bytecode**:
```
0000 OpConstant 0        // Push 5
0003 OpSetGlobal 0       // x = 5
0006 OpGetGlobal 0       // Load x
0009 OpTypeCast 1        // Cast to type[1] (দশমিক)
0012 OpSetGlobal 1       // y = result
```

**Execution Steps**:

| Step | Instruction | Stack | Globals | Description |
|------|-------------|-------|---------|-------------|
| 1 | `OpConstant 0` | `[5]` | `[]` | Push 5 (int) |
| 2 | `OpSetGlobal 0` | `[]` | `[x=5]` | Store x |
| 3 | `OpGetGlobal 0` | `[5]` | `[x=5]` | Load x |
| 4 | `OpTypeCast 1` | `[5.0]` | `[x=5]` | Cast to float |
| 5 | `OpSetGlobal 1` | `[]` | `[x=5, y=5.0]` | Store y |

---

## Structs

### Example 14: Struct Creation and Access

**Bhasa Code**:
```bhasa
ধরি ব্যক্তি = স্ট্রাক্ট {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা
};

ধরি p = ব্যক্তি{নাম: "রহিম", বয়স: 30};
p.নাম
```

**Bytecode**:
```
// Struct instance creation
0000 OpConstant 0        // Push "রহিম"
0003 OpConstant 1        // Push 30
0006 OpStruct 0          // Create struct from def[0]
0009 OpSetGlobal 0       // Store as p

// Field access
0012 OpGetGlobal 0       // Load p
0015 OpConstant 2        // Push field name "নাম"
0018 OpGetStructField    // Get field
```

**Execution Steps**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpConstant 0` | `["রহিম"]` | Push value for নাম |
| 2 | `OpConstant 1` | `["রহিম", 30]` | Push value for বয়স |
| 3 | `OpStruct 0` | `[<struct>]` | Create struct |
| 4 | `OpSetGlobal 0` | `[]` | Store as p |
| 5 | `OpGetGlobal 0` | `[<struct>]` | Load p |
| 6 | `OpConstant 2` | `[<struct>, "নাম"]` | Push field name |
| 7 | `OpGetStructField` | `["রহিম"]` | Get p.নাম |

---

## Object-Oriented Programming

### Example 15: Class Definition and Instantiation

**Bhasa Code**:
```bhasa
শ্রেণী ব্যক্তি {
    সার্বজনীন নাম: লেখা;
    
    সার্বজনীন নির্মাতা(n: লেখা) {
        এই.নাম = n;
    }
}

ধরি p = নতুন ব্যক্তি("রহিম");
```

**Bytecode**:
```
// Class definition
0000 OpClass 0              // Create class from def[0]
0003 OpClosure 1 0          // Create constructor closure
0007 OpDefineConstructor 0  // Attach constructor to class
0010 OpSetGlobal 0          // Store class as ব্যক্তি

// Instantiation
0013 OpGetGlobal 0          // Load class ব্যক্তি
0016 OpConstant 2           // Push "রহিম"
0019 OpNewInstance 1        // Create instance with 1 arg
0022 OpSetGlobal 1          // Store as p
```

**Constructor Bytecode**:
```
0000 OpGetThis              // Get this
0001 OpConstant 0           // Push field name "নাম"
0004 OpGetLocal 0           // Get parameter n
0006 OpSetInstanceField     // Set this.নাম = n
0007 OpReturn               // Return
```

**Execution Steps**:

| Phase | Step | Stack | Globals | Description |
|-------|------|-------|---------|-------------|
| Define | 1 | `[<class>]` | `[]` | Create class |
| Define | 2 | `[<class>, <ctor>]` | `[]` | Constructor |
| Define | 3 | `[<class>]` | `[]` | Attach constructor |
| Define | 4 | `[]` | `[ব্যক্তি=<cls>]` | Store class |
| Create | 1 | `[<class>]` | - | Load class |
| Create | 2 | `[<class>, "রহিম"]` | - | Push arg |
| Create | 3 | `[<obj>]` | - | Create instance |
| Create | 4 | `[]` | `[..., p=<obj>]` | Store instance |

**Inside Constructor**:

| Step | Instruction | Stack | This | Description |
|------|-------------|-------|------|-------------|
| 1 | `OpGetThis` | `[<obj>]` | `<obj>` | Get this |
| 2 | `OpConstant 0` | `[<obj>, "নাম"]` | `<obj>` | Field name |
| 3 | `OpGetLocal 0` | `[<obj>, "নাম", "রহিম"]` | `<obj>` | Get param |
| 4 | `OpSetInstanceField` | `[]` | `<obj>` | Set field |
| 5 | `OpReturn` | - | - | Return |

---

### Example 16: Method Call

**Bhasa Code**:
```bhasa
p.বলো("হ্যালো");
```

**Bytecode**:
```
0000 OpGetGlobal 0       // Load p
0003 OpConstant 1        // Push method name "বলো"
0006 OpConstant 2        // Push argument "হ্যালো"
0009 OpCallMethod 1      // Call method with 1 arg
```

**Method Bytecode** (বলো):
```
0000 OpGetThis           // Get this
0001 OpConstant 0        // Push field name "নাম"
0004 OpGetInstanceField  // Get this.নাম
0005 OpGetLocal 0        // Get parameter (বার্তা)
0007 OpAdd               // Concatenate (assuming + for strings)
0008 OpGetBuiltin 5      // Get দেখাও (print)
0010 OpCall 1            // Call print
0012 OpReturn            // Return
```

**Execution**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpGetGlobal 0` | `[<obj>]` | Load p |
| 2 | `OpConstant 1` | `[<obj>, "বলো"]` | Method name |
| 3 | `OpConstant 2` | `[<obj>, "বলো", "হ্যালো"]` | Argument |
| 4 | `OpCallMethod 1` | *Execute* | Call method |

**Inside Method**:

| Step | Instruction | Stack | This | Description |
|------|-------------|-------|------|-------------|
| 1 | `OpGetThis` | `[<obj>]` | `<obj>` | Get this |
| 2 | `OpConstant 0` | `[<obj>, "নাম"]` | `<obj>` | Field name |
| 3 | `OpGetInstanceField` | `["রহিম"]` | `<obj>` | Get name |
| 4 | `OpGetLocal 0` | `["রহিম", "হ্যালো"]` | `<obj>` | Get param |
| 5 | `OpAdd` | `["রহিম হ্যালো"]` | `<obj>` | Concat |
| 6 | `OpGetBuiltin 5` | `["রহিম হ্যালো", <print>]` | `<obj>` | Get print |
| 7 | `OpCall 1` | `[null]` | `<obj>` | Print |
| 8 | `OpReturn` | - | - | Return |

---

### Example 17: Inheritance and Super Call

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

**Bytecode** (কুকুর.বলো):
```
0000 OpGetSuper          // Get parent class
0001 OpConstant 0        // Push method name "বলো"
0004 OpCallMethod 0      // Call parent method
0007 OpPop               // Discard return
0008 OpConstant 1        // Push "ঘেউ ঘেউ"
0011 OpGetBuiltin 5      // Get print
0013 OpCall 1            // Call print
0015 OpReturn            // Return
```

**Execution**:

| Step | Instruction | Stack | Description |
|------|-------------|-------|-------------|
| 1 | `OpGetSuper` | `[<parent>]` | Get parent (প্রাণী) |
| 2 | `OpConstant 0` | `[<parent>, "বলো"]` | Method name |
| 3 | `OpCallMethod 0` | *Call* | Call parent.বলো() |
| - | *Inside parent* | - | Prints "প্রাণী" |
| 4 | `OpPop` | `[]` | Discard result |
| 5 | `OpConstant 1` | `["ঘেউ ঘেউ"]` | Push string |
| 6 | `OpGetBuiltin 5` | `["ঘেউ ঘেউ", <print>]` | Get print |
| 7 | `OpCall 1` | `[null]` | Print "ঘেউ ঘেউ" |
| 8 | `OpReturn` | - | Done |

**Output**:
```
প্রাণী
ঘেউ ঘেউ
```

---

## Complex Examples

### Example 18: Fibonacci Function (Recursive)

**Bhasa Code**:
```bhasa
ধরি fib = ফাংশন(n) {
    যদি (n < 2) {
        ফেরত n;
    } নাহলে {
        ফেরত fib(n - 1) + fib(n - 2);
    }
};
fib(5);
```

**Function Bytecode**:
```
0000 OpCurrentClosure    // Get self for recursion
0001 OpSetLocal 1        // Store in local[1]
0003 OpGetLocal 0        // Get n
0005 OpConstant 0        // Push 2
0008 OpGreaterThan       // Reverse: n < 2 becomes 2 > n
0009 OpJumpNotTruthy 15  // Jump if n >= 2
0012 OpGetLocal 0        // Load n
0014 OpReturnValue       // Return n
0015 OpGetLocal 1        // Load fib
0017 OpGetLocal 0        // Load n
0019 OpConstant 1        // Push 1
0022 OpSub               // n - 1
0023 OpCall 1            // fib(n-1)
0025 OpGetLocal 1        // Load fib again
0027 OpGetLocal 0        // Load n
0029 OpConstant 0        // Push 2
0032 OpSub               // n - 2
0033 OpCall 1            // fib(n-2)
0035 OpAdd               // fib(n-1) + fib(n-2)
0036 OpReturnValue       // Return result
```

**Call Tree for fib(5)**:

```
fib(5)
├── fib(4)
│   ├── fib(3)
│   │   ├── fib(2)
│   │   │   ├── fib(1) → 1
│   │   │   └── fib(0) → 0
│   │   │   → 1
│   │   └── fib(1) → 1
│   │   → 2
│   └── fib(2)
│       ├── fib(1) → 1
│       └── fib(0) → 0
│       → 1
│   → 3
└── fib(3)
    ├── fib(2)
    │   ├── fib(1) → 1
    │   └── fib(0) → 0
    │   → 1
    └── fib(1) → 1
    → 2
→ 5
```

---

### Example 19: Higher-Order Function (Map)

**Bhasa Code**:
```bhasa
ধরি map = ফাংশন(arr, fn) {
    ধরি result = [];
    পর্যন্ত (ধরি i = 0; i < len(arr); i = i + 1) {
        result = যোগ(result, fn(arr[i]));
    }
    ফেরত result;
};

ধরি double = ফাংশন(x) { ফেরত x * 2; };
map([1, 2, 3], double);
```

**Simplified Bytecode Flow**:

1. **Define map function** → Store in global
2. **Define double function** → Store in global
3. **Call map([1,2,3], double)**:
   - Create array [1, 2, 3]
   - Load double function
   - Call map with 2 args
   
4. **Inside map**:
   - Create empty result array
   - Loop through arr:
     - Get arr[i]
     - Call fn on element
     - Push to result
   - Return result

**Result**: `[2, 4, 6]`

---

### Example 20: Complete OOP Example

**Bhasa Code**:
```bhasa
বিমূর্ত শ্রেণী আকার {
    বিমূর্ত পদ্ধতি ক্ষেত্রফল(): দশমিক;
}

শ্রেণী বৃত্ত প্রসারিত আকার {
    সার্বজনীন ব্যাসার্ধ: দশমিক;
    
    সার্বজনীন নির্মাতা(r: দশমিক) {
        এই.ব্যাসার্ধ = r;
    }
    
    পুনর্সংজ্ঞা পদ্ধতি ক্ষেত্রফল(): দশমিক {
        ফেরত 3.14159 * এই.ব্যাসার্ধ * এই.ব্যাসার্ধ;
    }
}

ধরি c = নতুন বৃত্ত(5.0);
দেখাও(c.ক্ষেত্রফল());
```

**Execution Flow**:

1. **Define abstract class আকার**
   - Store class definition
   - No constructor (abstract)

2. **Define class বৃত্ত**
   - Create class
   - Set parent to আকার
   - Define constructor
   - Define ক্ষেত্রফল method

3. **Create instance c = new বৃত্ত(5.0)**
   - Load class
   - Push 5.0
   - Call OpNewInstance
   - Constructor sets ব্যাসার্ধ = 5.0

4. **Call c.ক্ষেত্রফল()**
   - Load c
   - Push method name
   - Call OpCallMethod
   - Inside method:
     - OpGetThis
     - Get ব্যাসার্ধ field (5.0)
     - Multiply: 3.14159 * 5.0 * 5.0
     - Return 78.53975

5. **Print result**
   - Call দেখাও(78.53975)

**Output**: `78.53975`

---

## Summary

This document demonstrates how bytecode instructions work in practice:

- **Stack-based execution**: All operations manipulate a stack
- **Instruction sequencing**: PC advances through bytecode
- **Function calls**: Stack frames for locals and parameters
- **Closures**: Free variables captured from outer scopes
- **OOP**: Classes, instances, methods, inheritance
- **Control flow**: Jumps for conditionals and loops

For more information, see:
- [Complete Bytecode Documentation](./bytecode-documentation.md)
- [Quick Reference](./quick-reference.md)
- [README](./README.md)

