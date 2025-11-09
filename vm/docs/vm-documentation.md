# VM Implementation - Complete Technical Documentation

## Table of Contents

- [VM Structure](#vm-structure)
- [Initialization](#initialization)
- [Execution Loop](#execution-loop)
- [Opcode Handlers](#opcode-handlers)
- [Arithmetic Operations](#arithmetic-operations)
- [Comparison Operations](#comparison-operations)
- [Control Flow](#control-flow)
- [Function Calls](#function-calls)
- [Data Structures](#data-structures)
- [OOP Operations](#oop-operations)
- [Type System](#type-system)
- [Helper Functions](#helper-functions)

---

## VM Structure

### Complete VM Definition

```go
type VM struct {
    constants []object.Object     // Constant pool from compiler
    
    stack []object.Object         // Value stack (2048 elements)
    sp    int                     // Stack pointer (next free slot)
    
    globals []object.Object       // Global variables (65536 slots)
    
    frames      []*Frame          // Call frame stack
    framesIndex int               // Current frame index
    
    // OOP support
    pendingConstructor *object.Closure          // Constructor for next class
    pendingMethods     map[string]*object.Closure  // Methods for next class
}
```

### Field Details

**constants** (`[]object.Object`)
- Compiled-in literals and functions
- Indexed by OpConstant instructions
- Immutable during execution

**stack** (`[]object.Object`)
- Size: 2048
- Stores intermediate computation values
- Grows upward (sp increases)

**sp** (`int`)
- Stack pointer
- Points to **next free slot**
- Top of stack is at `stack[sp-1]`

**globals** (`[]object.Object`)
- Size: 65,536
- Stores global variables
- Indexed by compiler-assigned indices
- Shared across all frames

**frames** (`[]*Frame`)
- Maximum 1024 frames
- Each frame represents a function call
- Enables recursion and local variables

**framesIndex** (`int`)
- Points to next free frame slot
- Current frame at `frames[framesIndex-1]`

**pendingConstructor** (`*object.Closure`)
- Temporary storage during class definition
- Set by OpDefineConstructor
- Used by OpClass

**pendingMethods** (`map[string]*object.Closure`)
- Temporary storage for methods
- Set by OpDefineMethod
- Used by OpClass
- Cleared after class creation

---

## Initialization

### New() - Create VM

```go
func New(bytecode *compiler.Bytecode) *VM {
    // Create main function from bytecode
    mainFn := &object.CompiledFunction{
        Instructions: bytecode.Instructions,
    }
    mainClosure := &object.Closure{Fn: mainFn}
    mainFrame := NewFrame(mainClosure, 0)
    
    // Initialize frame stack
    frames := make([]*Frame, MaxFrames)
    frames[0] = mainFrame
    
    return &VM{
        constants:   bytecode.Constants,
        stack:       make([]object.Object, StackSize),
        sp:          0,
        globals:     make([]object.Object, GlobalsSize),
        frames:      frames,
        framesIndex: 1,
        pendingConstructor: nil,
        pendingMethods:     make(map[string]*object.Closure),
    }
}
```

**Process:**
1. Wrap main bytecode in CompiledFunction
2. Wrap in Closure (no free variables)
3. Create frame 0 (main frame)
4. Allocate stacks and arrays
5. Return initialized VM

### NewWithGlobalsStore() - Reuse Globals

```go
func NewWithGlobalsStore(bytecode *compiler.Bytecode, s []object.Object) *VM {
    vm := New(bytecode)
    vm.globals = s
    return vm
}
```

**Use Case:** REPL - preserve variables between evaluations

---

## Execution Loop

### Run() - Main Execution

```go
func (vm *VM) Run() error {
    var ip int
    var ins code.Instructions
    var op code.Opcode
    
    for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
        // Advance instruction pointer
        vm.currentFrame().ip++
        
        // Fetch current instruction
        ip = vm.currentFrame().ip
        ins = vm.currentFrame().Instructions()
        op = code.Opcode(ins[ip])
        
        // Dispatch on opcode
        switch op {
        case code.OpConstant:
            // Handle OpConstant
        case code.OpAdd:
            // Handle OpAdd
        // ... more opcodes
        }
    }
    
    return nil
}
```

**Fetch-Decode-Execute Cycle:**
1. **Fetch**: Read opcode at IP
2. **Decode**: Determine operation
3. **Execute**: Perform operation
4. **Repeat**: Loop until program ends

### Instruction Pointer Management

**Increment IP:**
```go
vm.currentFrame().ip++  // Advance to next instruction
```

**Skip Operands:**
```go
vm.currentFrame().ip += 2  // Skip 2-byte operand
vm.currentFrame().ip += 1  // Skip 1-byte operand
```

**Jump:**
```go
pos := int(code.ReadUint16(ins[ip+1:]))
vm.currentFrame().ip = pos - 1  // -1 because loop increments
```

---

## Opcode Handlers

### OpConstant - Load Constant

```go
case code.OpConstant:
    constIndex := code.ReadUint16(ins[ip+1:])
    vm.currentFrame().ip += 2  // Skip operand
    
    if constIndex >= uint16(len(vm.constants)) {
        return fmt.Errorf("constant index out of range")
    }
    
    err := vm.push(vm.constants[constIndex])
    if err != nil {
        return err
    }
```

**Format:** `OpConstant <index:uint16>`

**Operation:**
1. Read 2-byte index
2. Validate index
3. Push constant onto stack
4. Advance IP by 2

**Example:**
```
Bytecode: OpConstant 0x0005
Constants[5] = Integer(42)
Result: Push 42 onto stack
```

### OpPop - Remove Top

```go
case code.OpPop:
    // Safety check: prevent stack underflow
    if vm.sp > vm.currentFrame().basePointer {
        vm.pop()
    }
```

**Format:** `OpPop`

**Operation:**
1. Check stack has values
2. Decrement SP
3. Value discarded

**Safety:** Prevents popping below frame base

### OpTrue/OpFalse - Push Boolean

```go
case code.OpTrue:
    err := vm.push(True)
    if err != nil {
        return err
    }

case code.OpFalse:
    err := vm.push(False)
    if err != nil {
        return err
    }
```

**Optimization:** Use singleton objects

### OpNull - Push Null

```go
case code.OpNull:
    err := vm.push(Null)
    if err != nil {
        return err
    }
```

---

## Arithmetic Operations

### Binary Operations Dispatcher

```go
case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv, code.OpMod,
     code.OpBitAnd, code.OpBitOr, code.OpBitXor, code.OpLeftShift, code.OpRightShift:
    err := vm.executeBinaryOperation(op)
    if err != nil {
        return err
    }
```

### executeBinaryOperation

```go
func (vm *VM) executeBinaryOperation(op code.Opcode) error {
    right := vm.pop()
    left := vm.pop()
    
    leftType := left.Type()
    rightType := right.Type()
    
    // String concatenation
    if leftType == object.STRING_OBJ && rightType == object.STRING_OBJ {
        return vm.executeBinaryStringOperation(op, left, right)
    }
    
    // Numeric operations
    if vm.isNumericType(leftType) && vm.isNumericType(rightType) {
        return vm.executeBinaryNumericOperation(op, left, right)
    }
    
    return fmt.Errorf("unsupported types for binary operation")
}
```

**Type Checking:**
1. String + String → String concatenation
2. Numeric + Numeric → Numeric operation
3. Other → Error

### Integer Operations

```go
func (vm *VM) executeBinaryIntegerOperation(
    op code.Opcode,
    left, right object.Object,
) error {
    leftValue := vm.toInt64(left)
    rightValue := vm.toInt64(right)
    
    var result int64
    
    switch op {
    case code.OpAdd:
        result = leftValue + rightValue
    case code.OpSub:
        result = leftValue - rightValue
    case code.OpMul:
        result = leftValue * rightValue
    case code.OpDiv:
        if rightValue == 0 {
            return fmt.Errorf("division by zero")
        }
        result = leftValue / rightValue
    case code.OpMod:
        if rightValue == 0 {
            return fmt.Errorf("modulo by zero")
        }
        result = leftValue % rightValue
    // ... bitwise operations
    }
    
    return vm.push(vm.promoteIntegerResult(result, left, right))
}
```

**Process:**
1. Convert operands to int64
2. Perform operation
3. Check for errors (division by zero)
4. Promote result to appropriate type
5. Push result

### Float Operations

```go
func (vm *VM) executeBinaryFloatOperation(
    op code.Opcode,
    left, right object.Object,
) error {
    leftValue := vm.toFloat64(left)
    rightValue := vm.toFloat64(right)
    
    var result float64
    
    switch op {
    case code.OpAdd:
        result = leftValue + rightValue
    case code.OpSub:
        result = leftValue - rightValue
    case code.OpMul:
        result = leftValue * rightValue
    case code.OpDiv:
        if rightValue == 0 {
            return fmt.Errorf("division by zero")
        }
        result = leftValue / rightValue
    case code.OpMod:
        // Floating-point modulo
        result = leftValue - rightValue*float64(int64(leftValue/rightValue))
    }
    
    // Promote to Double if either operand is Double
    if left.Type() == object.DOUBLE_OBJ || right.Type() == object.DOUBLE_OBJ {
        return vm.push(&object.Double{Value: result})
    }
    return vm.push(&object.Float{Value: float32(result)})
}
```

### Type Promotion

```go
func (vm *VM) promoteIntegerResult(result int64, left, right object.Object) object.Object {
    leftType := left.Type()
    rightType := right.Type()
    
    // Promotion hierarchy: Long > Int > Short > Byte
    if leftType == object.LONG_OBJ || rightType == object.LONG_OBJ {
        return &object.Long{Value: result}
    }
    if leftType == object.INT_OBJ || rightType == object.INT_OBJ {
        return &object.Int{Value: int32(result)}
    }
    if leftType == object.SHORT_OBJ || rightType == object.SHORT_OBJ {
        return &object.Short{Value: int16(result)}
    }
    return &object.Integer{Value: result}
}
```

**Rules:**
- Use larger type of the two operands
- Default to Integer for backward compatibility

### Bitwise Operations

```go
case code.OpBitAnd:
    result = leftValue & rightValue
case code.OpBitOr:
    result = leftValue | rightValue
case code.OpBitXor:
    result = leftValue ^ rightValue
case code.OpLeftShift:
    if rightValue < 0 {
        return fmt.Errorf("negative shift amount")
    }
    if rightValue >= 64 {
        return fmt.Errorf("shift amount too large")
    }
    result = leftValue << uint(rightValue)
case code.OpRightShift:
    if rightValue < 0 {
        return fmt.Errorf("negative shift amount")
    }
    if rightValue >= 64 {
        return fmt.Errorf("shift amount too large")
    }
    result = leftValue >> uint(rightValue)
```

**Validations:**
- No floating-point bitwise operations
- Shift amount must be non-negative
- Shift amount must be < 64

---

## Comparison Operations

### executeComparison

```go
func (vm *VM) executeComparison(op code.Opcode) error {
    right := vm.pop()
    left := vm.pop()
    
    // Handle NULL comparisons
    if left.Type() == object.NULL_OBJ || right.Type() == object.NULL_OBJ {
        switch op {
        case code.OpEqual:
            return vm.push(nativeBoolToBooleanObject(
                left.Type() == object.NULL_OBJ && right.Type() == object.NULL_OBJ))
        case code.OpNotEqual:
            return vm.push(nativeBoolToBooleanObject(
                !(left.Type() == object.NULL_OBJ && right.Type() == object.NULL_OBJ)))
        case code.OpGreaterThan:
            return vm.push(False)  // NULL is never greater
        case code.OpGreaterThanEqual:
            return vm.push(nativeBoolToBooleanObject(
                left.Type() == object.NULL_OBJ && right.Type() == object.NULL_OBJ))
        }
    }
    
    // Numeric comparisons
    if vm.isNumericType(left.Type()) && vm.isNumericType(right.Type()) {
        return vm.executeNumericComparison(op, left, right)
    }
    
    // Equality for non-numeric types
    switch op {
    case code.OpEqual:
        return vm.push(nativeBoolToBooleanObject(right == left))
    case code.OpNotEqual:
        return vm.push(nativeBoolToBooleanObject(right != left))
    }
}
```

**Comparison Rules:**
1. NULL == NULL → true
2. NULL compared to non-NULL → false
3. Numeric types → numeric comparison
4. Other types → pointer equality

### Numeric Comparison

```go
func (vm *VM) executeNumericComparison(
    op code.Opcode,
    left, right object.Object,
) error {
    // Use float comparison if either is floating-point
    if vm.isFloatingType(left.Type()) || vm.isFloatingType(right.Type()) {
        leftValue := vm.toFloat64(left)
        rightValue := vm.toFloat64(right)
        
        switch op {
        case code.OpEqual:
            return vm.push(nativeBoolToBooleanObject(leftValue == rightValue))
        case code.OpNotEqual:
            return vm.push(nativeBoolToBooleanObject(leftValue != rightValue))
        case code.OpGreaterThan:
            return vm.push(nativeBoolToBooleanObject(leftValue > rightValue))
        case code.OpGreaterThanEqual:
            return vm.push(nativeBoolToBooleanObject(leftValue >= rightValue))
        }
    }
    
    // Integer comparison
    leftValue := vm.toInt64(left)
    rightValue := vm.toInt64(right)
    
    switch op {
    case code.OpEqual:
        return vm.push(nativeBoolToBooleanObject(leftValue == rightValue))
    case code.OpNotEqual:
        return vm.push(nativeBoolToBooleanObject(leftValue != rightValue))
    case code.OpGreaterThan:
        return vm.push(nativeBoolToBooleanObject(leftValue > rightValue))
    case code.OpGreaterThanEqual:
        return vm.push(nativeBoolToBooleanObject(leftValue >= rightValue))
    }
}
```

---

## Control Flow

### OpJump - Unconditional Jump

```go
case code.OpJump:
    pos := int(code.ReadUint16(ins[ip+1:]))
    vm.currentFrame().ip = pos - 1  // -1 because loop increments
```

**Format:** `OpJump <position:uint16>`

**Operation:**
- Set IP to target position
- Minus 1 because loop will increment

### OpJumpNotTruthy - Conditional Jump

```go
case code.OpJumpNotTruthy:
    pos := int(code.ReadUint16(ins[ip+1:]))
    vm.currentFrame().ip += 2  // Skip operand
    
    condition := vm.pop()
    if !isTruthy(condition) {
        vm.currentFrame().ip = pos - 1
    }
```

**Format:** `OpJumpNotTruthy <position:uint16>`

**Operation:**
1. Read target position
2. Skip operand bytes
3. Pop condition from stack
4. If false, jump to target
5. Otherwise continue

**Truthiness:**
```go
func isTruthy(obj object.Object) bool {
    switch obj := obj.(type) {
    case *object.Boolean:
        return obj.Value
    case *object.Null:
        return false
    default:
        return true  // Everything else is truthy
    }
}
```

---

## Function Calls

### OpCall - Execute Function

```go
case code.OpCall:
    numArgs := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    err := vm.executeCall(int(numArgs))
    if err != nil {
        return err
    }
```

**Format:** `OpCall <numArgs:uint8>`

### executeCall

```go
func (vm *VM) executeCall(numArgs int) error {
    callee := vm.stack[vm.sp-1-numArgs]
    
    switch callee := callee.(type) {
    case *object.Closure:
        return vm.callClosure(callee, numArgs)
    case *object.Builtin:
        return vm.callBuiltin(callee, numArgs)
    case *object.BoundMethod:
        return vm.callBoundMethod(callee, numArgs)
    default:
        return fmt.Errorf("calling non-function")
    }
}
```

**Stack Layout Before Call:**
```
[..., callee, arg1, arg2, ..., argN]
           ^                        ^
           sp-numArgs-1             sp-1
```

### callClosure - Call User Function

```go
func (vm *VM) callClosure(cl *object.Closure, numArgs int) error {
    // Validate argument count
    if numArgs != cl.Fn.NumParameters {
        return fmt.Errorf("wrong number of arguments: want=%d, got=%d",
            cl.Fn.NumParameters, numArgs)
    }
    
    // Create new frame
    frame := NewFrame(cl, vm.sp-numArgs)
    vm.pushFrame(frame)
    
    // Allocate space for locals
    vm.sp = frame.basePointer + cl.Fn.NumLocals
    
    return nil
}
```

**Process:**
1. Validate argument count
2. Create new frame (base pointer = sp - numArgs)
3. Push frame onto frame stack
4. Allocate local variables
5. Execution continues in new frame

**Stack After Call:**
```
Frame 0: [..., callee, arg1, arg2]
                       ^BP
Frame 1: [arg1, arg2, local1, local2, ...]
          ^BP         ^locals start
```

### callBuiltin - Call Native Function

```go
func (vm *VM) callBuiltin(builtin *object.Builtin, numArgs int) error {
    // Get arguments from stack
    args := vm.stack[vm.sp-numArgs : vm.sp]
    
    // Call builtin function
    result := builtin.Fn(args...)
    
    // Clean up stack
    vm.sp = vm.sp - numArgs - 1
    
    // Push result
    if result != nil {
        vm.push(result)
    } else {
        vm.push(Null)
    }
    
    return nil
}
```

**Process:**
1. Extract arguments from stack
2. Call Go function
3. Remove callee and arguments
4. Push result (or null)

**No Frame:** Builtins execute in current frame

### OpReturnValue - Return With Value

```go
case code.OpReturnValue:
    returnValue := vm.pop()
    
    // Pop frame
    frame := vm.popFrame()
    
    // Reset stack to before call
    vm.sp = frame.basePointer - 1
    
    // Push return value
    err := vm.push(returnValue)
    if err != nil {
        return err
    }
```

**Stack Transformation:**
```
Before: [caller_stack, callee, args, locals, return_value]
                       ^BP
After:  [caller_stack, return_value]
```

### OpReturn - Return Null

```go
case code.OpReturn:
    frame := vm.popFrame()
    vm.sp = frame.basePointer - 1
    
    err := vm.push(Null)
    if err != nil {
        return err
    }
```

**Same as OpReturnValue but pushes null**

---

## Data Structures

### OpArray - Create Array

```go
case code.OpArray:
    numElements := int(code.ReadUint16(ins[ip+1:]))
    vm.currentFrame().ip += 2
    
    // Build array from stack
    array := vm.buildArray(vm.sp-numElements, vm.sp)
    vm.sp = vm.sp - numElements
    
    err := vm.push(array)
    if err != nil {
        return err
    }
```

**buildArray:**
```go
func (vm *VM) buildArray(startIndex, endIndex int) object.Object {
    elements := make([]object.Object, endIndex-startIndex)
    
    for i := startIndex; i < endIndex; i++ {
        elements[i-startIndex] = vm.stack[i]
    }
    
    return &object.Array{Elements: elements}
}
```

**Stack:**
```
Before: [elem1, elem2, elem3]
After:  [[elem1, elem2, elem3]]
```

### OpHash - Create Hash

```go
case code.OpHash:
    numElements := int(code.ReadUint16(ins[ip+1:]))
    vm.currentFrame().ip += 2
    
    hash, err := vm.buildHash(vm.sp-numElements, vm.sp)
    if err != nil {
        return err
    }
    vm.sp = vm.sp - numElements
    
    err = vm.push(hash)
    if err != nil {
        return err
    }
```

**buildHash:**
```go
func (vm *VM) buildHash(startIndex, endIndex int) (object.Object, error) {
    hashedPairs := make(map[object.HashKey]object.HashPair)
    
    // Stack has alternating keys and values
    for i := startIndex; i < endIndex; i += 2 {
        key := vm.stack[i]
        value := vm.stack[i+1]
        
        // Check if key is hashable
        hashKey, ok := key.(object.Hashable)
        if !ok {
            return nil, fmt.Errorf("unusable as hash key: %s", key.Type())
        }
        
        hashedPairs[hashKey.HashKey()] = object.HashPair{
            Key:   key,
            Value: value,
        }
    }
    
    return &object.Hash{Pairs: hashedPairs}, nil
}
```

**Stack:**
```
Before: [key1, val1, key2, val2]
After:  [{key1: val1, key2: val2}]
```

### OpIndex - Index Access

```go
case code.OpIndex:
    index := vm.pop()
    left := vm.pop()
    
    err := vm.executeIndexExpression(left, index)
    if err != nil {
        return err
    }
```

**executeIndexExpression:**
```go
func (vm *VM) executeIndexExpression(left, index object.Object) error {
    switch {
    case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
        return vm.executeArrayIndex(left, index)
    case left.Type() == object.HASH_OBJ:
        return vm.executeHashIndex(left, index)
    default:
        return fmt.Errorf("index operator not supported: %s", left.Type())
    }
}
```

**Array Index:**
```go
func (vm *VM) executeArrayIndex(array, index object.Object) error {
    arrayObject := array.(*object.Array)
    i := index.(*object.Integer).Value
    max := int64(len(arrayObject.Elements) - 1)
    
    if i < 0 || i > max {
        return vm.push(Null)  // Out of bounds returns null
    }
    
    return vm.push(arrayObject.Elements[i])
}
```

**Hash Index:**
```go
func (vm *VM) executeHashIndex(hash, index object.Object) error {
    hashObject := hash.(*object.Hash)
    
    key, ok := index.(object.Hashable)
    if !ok {
        return fmt.Errorf("unusable as hash key: %s", index.Type())
    }
    
    pair, ok := hashObject.Pairs[key.HashKey()]
    if !ok {
        return vm.push(Null)  // Missing key returns null
    }
    
    return vm.push(pair.Value)
}
```

---

## OOP Operations

### OpClass - Define Class

```go
case code.OpClass:
    constIndex := code.ReadUint16(ins[ip+1:])
    vm.currentFrame().ip += 2
    
    // Get class template from constants
    classTemplate := vm.constants[constIndex].(*object.Class)
    
    // Create runtime class with pending constructor/methods
    class := &object.Class{
        Name:         classTemplate.Name,
        SuperClass:   classTemplate.SuperClass,
        Fields:       classTemplate.Fields,
        Methods:      make(map[string]*object.Method),
        Constructor:  vm.pendingConstructor,
        StaticFields: classTemplate.StaticFields,
        IsAbstract:   classTemplate.IsAbstract,
        IsFinal:      classTemplate.IsFinal,
        FieldAccess:  classTemplate.FieldAccess,
        FieldOrder:   classTemplate.FieldOrder,
    }
    
    // Attach runtime methods
    for name, method := range classTemplate.Methods {
        methodCopy := &object.Method{
            Name:       method.Name,
            Access:     method.Access,
            IsStatic:   method.IsStatic,
            IsFinal:    method.IsFinal,
            IsAbstract: method.IsAbstract,
            Closure:    method.Closure,
        }
        if runtimeClosure, exists := vm.pendingMethods[name]; exists {
            methodCopy.Closure = runtimeClosure
        }
        class.Methods[name] = methodCopy
    }
    
    // Clear pending data
    vm.pendingConstructor = nil
    vm.pendingMethods = make(map[string]*object.Closure)
    
    err := vm.push(class)
```

**Process:**
1. Load class template from constants
2. Create runtime copy
3. Attach pending constructor
4. Attach pending methods
5. Clear pending data
6. Push class

### OpNewInstance - Create Instance

```go
case code.OpNewInstance:
    numArgs := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    // Get class from stack
    classObj := vm.stack[vm.sp-1]
    class := classObj.(*object.Class)
    vm.sp--  // Remove class
    
    // Create instance
    instance := &object.ClassInstance{
        Class:  class,
        Fields: make(map[string]object.Object),
        This:   nil,
    }
    instance.This = instance
    
    // Initialize fields to null
    for _, fieldName := range class.FieldOrder {
        instance.Fields[fieldName] = Null
    }
    
    // Call constructor if exists
    if class.Constructor != nil {
        // Complex stack manipulation to call constructor
        // ... (see code for details)
    } else {
        // No constructor, just push instance
        err := vm.push(instance)
    }
```

**Constructor Calling:**
1. Push placeholder (class itself)
2. Push instance as first argument (this)
3. Constructor arguments already on stack
4. Reorder stack: [class, instance, arg1, arg2, ...]
5. Call constructor closure
6. Constructor returns instance

### OpCallMethod - Call Method

```go
case code.OpCallMethod:
    numArgs := code.ReadUint8(ins[ip+1:])
    vm.currentFrame().ip += 1
    
    // Pop arguments
    args := make([]object.Object, numArgs)
    for i := int(numArgs) - 1; i >= 0; i-- {
        args[i] = vm.pop()
    }
    
    // Get method name and object
    methodName := vm.pop().(*object.String).Value
    obj := vm.pop()
    instance := obj.(*object.ClassInstance)
    
    // Find method in class hierarchy
    method := instance.Class.GetMethod(methodName)
    if method == nil {
        return fmt.Errorf("method '%s' not found", methodName)
    }
    
    // Prepare arguments: [this, arg1, arg2, ...]
    allArgs := append([]object.Object{instance}, args...)
    
    // Create frame and push arguments
    frame := NewFrame(method.Closure, vm.sp-len(allArgs))
    vm.pushFrame(frame)
    
    for _, arg := range allArgs {
        err := vm.push(arg)
    }
    
    vm.sp = frame.basePointer + method.Closure.Fn.NumLocals
```

### OpGetThis - Get 'this'

```go
case code.OpGetThis:
    // 'this' is always first parameter (index 0)
    basePointer := vm.currentFrame().basePointer
    err := vm.push(vm.stack[basePointer])
```

**'this' is stored at base pointer of method frame**

### OpGetInstanceField - Get Field

```go
case code.OpGetInstanceField:
    fieldName := vm.pop().(*object.String).Value
    obj := vm.pop()
    instance := obj.(*object.ClassInstance)
    
    // Get field value
    value, exists := instance.GetField(fieldName)
    if !exists {
        value = Null
    }
    
    err := vm.push(value)
```

### OpSetInstanceField - Set Field

```go
case code.OpSetInstanceField:
    value := vm.pop()
    fieldName := vm.pop().(*object.String).Value
    obj := vm.pop()
    instance := obj.(*object.ClassInstance)
    
    // Set field
    instance.SetField(fieldName, value)
    
    // Push value back (for chaining)
    err := vm.push(value)
```

---

## Type System

### Type Checking

```go
func (vm *VM) checkType(obj object.Object, expectedType string) bool {
    actualType := vm.getTypeName(obj)
    return actualType == expectedType
}

func (vm *VM) getTypeName(obj object.Object) string {
    switch obj.Type() {
    case object.BYTE_OBJ:
        return "বাইট"
    case object.SHORT_OBJ:
        return "ছোট_সংখ্যা"
    case object.INT_OBJ:
        return "পূর্ণসংখ্যা"
    case object.LONG_OBJ, object.INTEGER_OBJ:
        return "দীর্ঘ_সংখ্যা"
    case object.FLOAT_OBJ:
        return "দশমিক"
    case object.DOUBLE_OBJ:
        return "দশমিক_দ্বিগুণ"
    case object.STRING_OBJ:
        return "লেখা"
    case object.BOOLEAN_OBJ:
        return "বুলিয়ান"
    // ... more types
    }
}
```

### Type Casting

```go
func (vm *VM) castType(obj object.Object, targetType string) (object.Object, error) {
    if vm.checkType(obj, targetType) {
        return obj, nil  // Already correct type
    }
    
    switch targetType {
    case "বাইট":
        val := vm.toInt64(obj)
        if val < 0 || val > 255 {
            return nil, fmt.Errorf("value out of range for byte")
        }
        return &object.Byte{Value: int8(val)}, nil
        
    case "পূর্ণসংখ্যা":
        val := vm.toInt64(obj)
        if val < -2147483648 || val > 2147483647 {
            return nil, fmt.Errorf("value out of range for int")
        }
        return &object.Int{Value: int32(val)}, nil
        
    case "দশমিক":
        val := vm.toFloat64(obj)
        return &object.Float{Value: float32(val)}, nil
        
    case "লেখা":
        return &object.String{Value: obj.Inspect()}, nil
        
    // ... more conversions
    }
}
```

### OpTypeCast - Explicit Cast

```go
case code.OpTypeCast:
    constIndex := code.ReadUint16(ins[ip+1:])
    vm.currentFrame().ip += 2
    
    targetType := vm.constants[constIndex].(*object.String).Value
    value := vm.pop()
    
    castedValue, err := vm.castType(value, targetType)
    if err != nil {
        return err
    }
    
    err = vm.push(castedValue)
```

---

## Helper Functions

### Stack Helpers

```go
func (vm *VM) push(o object.Object) error {
    if vm.sp >= StackSize {
        return fmt.Errorf("stack overflow")
    }
    vm.stack[vm.sp] = o
    vm.sp++
    return nil
}

func (vm *VM) pop() object.Object {
    o := vm.stack[vm.sp-1]
    vm.sp--
    return o
}

func (vm *VM) StackTop() object.Object {
    if vm.sp == 0 {
        return nil
    }
    return vm.stack[vm.sp-1]
}

func (vm *VM) LastPoppedStackElem() object.Object {
    return vm.stack[vm.sp]
}
```

### Frame Helpers

```go
func (vm *VM) currentFrame() *Frame {
    return vm.frames[vm.framesIndex-1]
}

func (vm *VM) pushFrame(f *Frame) {
    vm.frames[vm.framesIndex] = f
    vm.framesIndex++
}

func (vm *VM) popFrame() *Frame {
    vm.framesIndex--
    return vm.frames[vm.framesIndex]
}
```

### Type Conversion

```go
func (vm *VM) toInt64(obj object.Object) int64 {
    switch v := obj.(type) {
    case *object.Integer:
        return v.Value
    case *object.Byte:
        return int64(v.Value)
    case *object.Short:
        return int64(v.Value)
    case *object.Int:
        return int64(v.Value)
    case *object.Long:
        return v.Value
    case *object.Char:
        return int64(v.Value)
    case *object.Float:
        return int64(v.Value)
    case *object.Double:
        return int64(v.Value)
    default:
        return 0
    }
}

func (vm *VM) toFloat64(obj object.Object) float64 {
    switch v := obj.(type) {
    case *object.Integer:
        return float64(v.Value)
    case *object.Byte:
        return float64(v.Value)
    case *object.Short:
        return float64(v.Value)
    case *object.Int:
        return float64(v.Value)
    case *object.Long:
        return float64(v.Value)
    case *object.Char:
        return float64(v.Value)
    case *object.Float:
        return float64(v.Value)
    case *object.Double:
        return v.Value
    default:
        return 0.0
    }
}
```

### Type Checking

```go
func (vm *VM) isNumericType(t object.ObjectType) bool {
    return t == object.INTEGER_OBJ || t == object.BYTE_OBJ || 
           t == object.SHORT_OBJ || t == object.INT_OBJ ||
           t == object.LONG_OBJ || t == object.FLOAT_OBJ ||
           t == object.DOUBLE_OBJ || t == object.CHAR_OBJ
}

func (vm *VM) isFloatingType(t object.ObjectType) bool {
    return t == object.FLOAT_OBJ || t == object.DOUBLE_OBJ
}
```

---

## Summary

The **VM** implementation provides:

✅ **Efficient Execution**: Stack-based bytecode interpreter  
✅ **Type System**: Rich numeric types with automatic promotion  
✅ **OOP Support**: Classes, methods, inheritance  
✅ **Error Handling**: Comprehensive validation  
✅ **Call Frames**: Proper function call management  
✅ **Closures**: Free variable capture  
✅ **Built-ins**: Native function integration  

**Key Design Decisions:**
- Stack-based for cache efficiency
- Type-specific operations for performance
- Singleton objects for common values
- Frame-based local variables
- Separate integer/float arithmetic

For frame system details, see [frame-documentation.md](./frame-documentation.md).

