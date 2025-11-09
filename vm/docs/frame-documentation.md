# Frame System - Call Frame Documentation

## Overview

The **Frame** system manages function calls in the Bhasa VM. Each frame represents a single function invocation, tracking execution state and isolating local variables.

## Table of Contents

- [Frame Structure](#frame-structure)
- [Frame Lifecycle](#frame-lifecycle)
- [Stack Management](#stack-management)
- [Recursion Support](#recursion-support)
- [Closure Integration](#closure-integration)
- [Examples](#examples)

---

## Frame Structure

### Frame Definition

```go
type Frame struct {
    cl          *object.Closure  // Function being executed
    ip          int              // Instruction pointer (-1 to start)
    basePointer int              // Stack base for this frame
}
```

### Field Descriptions

#### cl (*object.Closure)

**Purpose:** The function being executed

**Contains:**
- `Fn`: CompiledFunction with bytecode
- `Free`: Captured free variables (for closures)

**Example:**
```go
closure := &object.Closure{
    Fn: &object.CompiledFunction{
        Instructions:  bytecode,
        NumLocals:     3,
        NumParameters: 2,
    },
    Free: []object.Object{capturedVar},
}
```

#### ip (int)

**Purpose:** Instruction pointer

**Initial Value:** -1 (before first instruction)

**Points To:** Current instruction being executed

**Updates:**
- Incremented after most instructions
- Set directly for jumps
- Advanced by operand size for multi-byte instructions

**Example:**
```go
// Start
frame.ip = -1

// After first increment
frame.ip = 0  // Executing first instruction

// After OpConstant (2-byte operand)
frame.ip = 3  // Skip instruction + operand
```

#### basePointer (int)

**Purpose:** Base of this frame's stack segment

**Points To:** First parameter/local variable

**Used For:**
- Accessing local variables
- Accessing parameters
- Resetting stack on return

**Calculation:** `basePointer = sp - numArgs` when calling

**Example:**
```
Global Stack:
[global1, global2, func, arg1, arg2, local1, local2]
                        ^BP=3
```

---

## Frame Lifecycle

### 1. Frame Creation

```go
func NewFrame(cl *object.Closure, basePointer int) *Frame {
    return &Frame{
        cl:          cl,
        ip:          -1,           // Will be incremented to 0
        basePointer: basePointer,
    }
}
```

**When Created:**
- Function call (OpCall)
- Method call (OpCallMethod)
- Constructor call

**Initial State:**
- IP = -1 (before first instruction)
- Base pointer set to current stack position

### 2. Frame Activation

**In callClosure:**
```go
func (vm *VM) callClosure(cl *object.Closure, numArgs int) error {
    // Validate arguments
    if numArgs != cl.Fn.NumParameters {
        return fmt.Errorf("wrong number of arguments")
    }
    
    // Create frame
    frame := NewFrame(cl, vm.sp-numArgs)
    
    // Push onto frame stack
    vm.pushFrame(frame)
    
    // Allocate space for locals
    vm.sp = frame.basePointer + cl.Fn.NumLocals
    
    return nil
}
```

**Steps:**
1. Validate argument count
2. Create frame with base pointer
3. Push frame onto frame stack
4. Allocate local variables

### 3. Frame Execution

**Main Loop:**
```go
for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
    vm.currentFrame().ip++
    
    ip := vm.currentFrame().ip
    ins := vm.currentFrame().Instructions()
    op := code.Opcode(ins[ip])
    
    // Execute instruction
    switch op {
        // ... opcode handlers
    }
}
```

**Process:**
- Increment IP
- Fetch instruction
- Decode opcode
- Execute operation
- Repeat until function ends

### 4. Frame Cleanup

**OpReturnValue:**
```go
case code.OpReturnValue:
    returnValue := vm.pop()
    
    // Pop frame
    frame := vm.popFrame()
    
    // Reset stack to before call
    vm.sp = frame.basePointer - 1
    
    // Push return value
    err := vm.push(returnValue)
```

**Steps:**
1. Pop return value from stack
2. Pop frame from frame stack
3. Reset stack pointer (discard parameters and locals)
4. Push return value

**Stack Transformation:**
```
Before: [caller_data, func, arg1, arg2, local1, return_val]
                       ^BP
After:  [caller_data, return_val]
```

---

## Stack Management

### Stack Layout

```
Global View:
┌─────────────────────────────────────┐
│ Global Variables                    │
├─────────────────────────────────────┤
│ Frame 0 (main):                     │
│   [instruction_results]             │
│   BP=0, SP varies                   │
├─────────────────────────────────────┤
│ Frame 1 (function call):            │
│   [callee, arg1, arg2, local1, ...]│
│   ^BP       ^params   ^locals       │
├─────────────────────────────────────┤
│ Frame 2 (nested call):              │
│   [callee, arg1, local1, ...]       │
│   ^BP                                │
└─────────────────────────────────────┘
                                    ^SP
```

### Local Variable Access

**OpGetLocal:**
```go
case code.OpGetLocal:
    localIndex := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    frame := vm.currentFrame()
    
    // Access: stack[basePointer + index]
    err := vm.push(vm.stack[frame.basePointer+int(localIndex)])
```

**Indexing:**
- Local 0 = First parameter (or first local if no params)
- Local 1 = Second parameter (or second local)
- ...
- Locals are consecutive starting at base pointer

**Example:**
```bengali
ফাংশন add(a, b) {
    ধরি sum = a + b;
    ফেরত sum;
}
```

**Locals:**
- Local 0 = a (parameter)
- Local 1 = b (parameter)
- Local 2 = sum (variable)

**Stack:**
```
[..., add_closure, 5, 3, 8]
      ^BP
      Local 0=5, Local 1=3, Local 2=8
```

### OpSetLocal

```go
case code.OpSetLocal:
    localIndex := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    frame := vm.currentFrame()
    
    // Store: stack[basePointer + index] = value
    vm.stack[frame.basePointer+int(localIndex)] = vm.pop()
```

---

## Recursion Support

### Recursive Call Example

```bengali
ধরি factorial = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * factorial(n - ১);
};

factorial(৩);
```

### Frame Stack During Execution

**Initial Call: factorial(3)**
```
Frame 0: main
  IP=20, BP=0

Frame 1: factorial(3)
  IP=0, BP=5
  Stack: [..., factorial, 3]
```

**Recursive Call: factorial(2)**
```
Frame 0: main
  IP=20, BP=0

Frame 1: factorial(3)
  IP=15, BP=5
  Stack: [..., factorial, 3]

Frame 2: factorial(2)
  IP=0, BP=10
  Stack: [..., factorial, 3, factorial, 2]
```

**Recursive Call: factorial(1)**
```
Frame 0: main
  IP=20, BP=0

Frame 1: factorial(3)
  IP=15, BP=5
  Stack: [..., factorial, 3]

Frame 2: factorial(2)
  IP=15, BP=10
  Stack: [..., factorial, 3, factorial, 2]

Frame 3: factorial(1)
  IP=0, BP=15
  Stack: [..., factorial, 3, factorial, 2, factorial, 1]
```

**Base Case Returns: 1**
```
Frame 3 returns 1

Frame 2: factorial(2)
  IP=15, BP=10
  Stack: [..., factorial, 3, factorial, 2, 1]
  Computes: 2 * 1 = 2
```

**Frame 2 Returns: 2**
```
Frame 1: factorial(3)
  IP=15, BP=5
  Stack: [..., factorial, 3, 2]
  Computes: 3 * 2 = 6
```

**Frame 1 Returns: 6**
```
Frame 0: main
  Stack: [..., 6]
  Result: 6
```

### Maximum Recursion Depth

```go
const MaxFrames = 1024
```

**Exceeding Limit:**
- Frame stack has fixed size
- Deep recursion will overflow
- No automatic tail-call optimization (yet)

---

## Closure Integration

### Free Variables

Closures capture variables from outer scopes:

```bengali
ধরি makeAdder = ফাংশন(x) {
    ধরি adder = ফাংশন(y) {
        ফেরত x + y;  // 'x' is free variable
    };
    ফেরত adder;
};

ধরি add5 = makeAdder(৫);
লেখ(add5(৩));  // Output: 8
```

### Closure Structure

```go
type Closure struct {
    Fn   *object.CompiledFunction
    Free []object.Object  // Captured free variables
}
```

**Free Variable Storage:**
- Stored in closure, not on stack
- Accessible via OpGetFree
- Persist after outer function returns

### OpGetFree

```go
case code.OpGetFree:
    freeIndex := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    currentClosure := vm.currentFrame().cl
    
    err := vm.push(currentClosure.Free[freeIndex])
```

**Process:**
1. Get index of free variable
2. Access current frame's closure
3. Get free variable from closure
4. Push onto stack

### Free Variable Example

**Compiled Code for `adder`:**
```
OpGetFree 0      // Get 'x' from closure
OpGetLocal 0     // Get 'y' from parameters
OpAdd            // x + y
OpReturnValue
```

**Frame State:**
```
Frame for adder(3):
  cl.Free = [5]           // 'x' captured
  stack[BP+0] = 3         // 'y' parameter
  
OpGetFree 0 → Push 5
OpGetLocal 0 → Push 3
OpAdd → 5 + 3 = 8
OpReturnValue → Return 8
```

---

## Examples

### Example 1: Simple Function

**Code:**
```bengali
ধরি double = ফাংশন(x) {
    ফেরত x * ২;
};

double(৫);
```

**Frame Sequence:**

**1. Call double(5):**
```
Before Call:
  Stack: [..., double_closure, 5]
  SP: 7

Create Frame:
  Frame 1:
    cl: double_closure
    ip: -1
    basePointer: 6

After Frame Creation:
  Stack: [..., double_closure, 5]
         ^BP=6
  SP: 7 (parameters already on stack)
```

**2. Execute Function:**
```
OpGetLocal 0:
  Stack: [..., double_closure, 5, 5]
  
OpConstant 1 (value=2):
  Stack: [..., double_closure, 5, 5, 2]
  
OpMul:
  Stack: [..., double_closure, 5, 10]
  
OpReturnValue:
  Pop return value: 10
  Pop frame
  Reset SP to BP-1: SP = 5
  Stack: [..., 10]
```

### Example 2: Nested Calls

**Code:**
```bengali
ধরি add = ফাংশন(a, b) { ফেরত a + b; };
ধরি multiply = ফাংশন(x, y) { ফেরত x * y; };
ধরি compute = ফাংশন(n) {
    ফেরত multiply(add(n, ২), ৩);
};

compute(৫);
```

**Frame Sequence:**

**1. compute(5):**
```
Frame 1: compute
  BP: 3
  Locals: [5]
```

**2. Call add(n, 2) = add(5, 2):**
```
Frame 1: compute
  BP: 3
  
Frame 2: add
  BP: 5
  Locals: [5, 2]
  
Returns: 7
```

**3. Call multiply(7, 3):**
```
Frame 1: compute
  BP: 3
  
Frame 2: multiply
  BP: 5
  Locals: [7, 3]
  
Returns: 21
```

**4. compute returns:**
```
Frame 0: main
  Stack: [21]
```

### Example 3: Closure with Free Variables

**Code:**
```bengali
ধরি counter = ফাংশন() {
    ধরি count = ০;
    ফেরত ফাংশন() {
        count = count + ১;
        ফেরত count;
    };
};

ধরি increment = counter();
লেখ(increment());  // 1
লেখ(increment());  // 2
```

**Execution:**

**1. Call counter():**
```
Frame 1: counter
  BP: 3
  Locals: [0]  // count = 0
  
Create inner closure:
  Closure {
    Fn: increment_function,
    Free: [Integer(0)]  // Captured 'count'
  }
  
Returns: closure
```

**2. First increment() call:**
```
Frame 1: increment
  cl.Free: [Integer(0)]
  
OpGetFree 0 → Push 0
OpConstant 1 → Push 1
OpAdd → 0 + 1 = 1
OpSetFree 0 → Free[0] = 1
OpGetFree 0 → Push 1
OpReturnValue → Return 1
```

**3. Second increment() call:**
```
Frame 1: increment
  cl.Free: [Integer(1)]  // Updated from last call
  
OpGetFree 0 → Push 1
OpConstant 1 → Push 1
OpAdd → 1 + 1 = 2
OpSetFree 0 → Free[0] = 2
OpGetFree 0 → Push 2
OpReturnValue → Return 2
```

---

## Frame Stack vs Value Stack

### Two Separate Stacks

**Frame Stack:**
- Stores Frame objects
- Max depth: 1024
- Tracks function call hierarchy
- Contains IP, BP, closure

**Value Stack:**
- Stores Object values
- Max size: 2048
- Holds computation results
- Shared across all frames

### Relationship

```
Frame Stack          Value Stack
┌─────────────┐     ┌──────────────┐
│ Frame 2     │     │ [values...]  │
│ BP=15, IP=5 │────→│ ^BP=15       │
├─────────────┤     │              │
│ Frame 1     │     │ [values...]  │
│ BP=5, IP=12 │────→│ ^BP=5        │
├─────────────┤     │              │
│ Frame 0     │     │ [values...]  │
│ BP=0, IP=20 │────→│ ^BP=0        │
└─────────────┘     └──────────────┘
                                 ^SP
```

Each frame's BP points into the value stack.

---

## Summary

The **Frame** system provides:

✅ **Function Isolation**: Separate execution contexts  
✅ **Local Variables**: Stack-based storage  
✅ **Recursion Support**: Unlimited depth (up to MaxFrames)  
✅ **Closure Integration**: Free variable access  
✅ **Efficient Calls**: Minimal overhead  
✅ **Clean Returns**: Automatic cleanup  

**Key Design Points:**
- Base pointer for local access
- Instruction pointer for execution
- Closure reference for free variables
- Simple push/pop for call/return

The frame system is the foundation of function calls in the Bhasa VM, enabling efficient execution of user-defined functions with proper scope isolation.

