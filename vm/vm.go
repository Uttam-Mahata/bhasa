package vm

import (
	"bhasa/code"
	"bhasa/compiler"
	"bhasa/object"
	"fmt"
)

const StackSize = 2048
const GlobalsSize = 65536
const MaxFrames = 1024

var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}
var Null = &object.Null{}

// VM is a virtual machine
type VM struct {
	constants []object.Object

	stack []object.Object
	sp    int // Always points to the next value. Top of stack is stack[sp-1]

	globals []object.Object

	frames      []*Frame
	framesIndex int
}

// New creates a new VM
func New(bytecode *compiler.Bytecode) *VM {
	mainFn := &object.CompiledFunction{Instructions: bytecode.Instructions}
	mainClosure := &object.Closure{Fn: mainFn}
	mainFrame := NewFrame(mainClosure, 0)

	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame

	return &VM{
		constants: bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,

		globals: make([]object.Object, GlobalsSize),

		frames:      frames,
		framesIndex: 1,
	}
}

// NewWithGlobalsStore creates a VM with existing globals
func NewWithGlobalsStore(bytecode *compiler.Bytecode, s []object.Object) *VM {
	vm := New(bytecode)
	vm.globals = s
	return vm
}

// StackTop returns the top element of the stack
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

// Run executes the bytecode
func (vm *VM) Run() error {
	var ip int
	var ins code.Instructions
	var op code.Opcode

	for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
		vm.currentFrame().ip++

		ip = vm.currentFrame().ip
		ins = vm.currentFrame().Instructions()
		op = code.Opcode(ins[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case code.OpPop:
			vm.pop()

		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv, code.OpMod,
			code.OpBitAnd, code.OpBitOr, code.OpBitXor, code.OpLeftShift, code.OpRightShift:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}

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

		case code.OpEqual, code.OpNotEqual, code.OpGreaterThan, code.OpGreaterThanEqual:
			err := vm.executeComparison(op)
			if err != nil {
				return err
			}

		case code.OpBang:
			err := vm.executeBangOperator()
			if err != nil {
				return err
			}

		case code.OpAnd:
			err := vm.executeAndOperator()
			if err != nil {
				return err
			}

		case code.OpOr:
			err := vm.executeOrOperator()
			if err != nil {
				return err
			}

		case code.OpMinus:
			err := vm.executeMinusOperator()
			if err != nil {
				return err
			}

		case code.OpBitNot:
			err := vm.executeBitNotOperator()
			if err != nil {
				return err
			}

		case code.OpJump:
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip = pos - 1

		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				vm.currentFrame().ip = pos - 1
			}

		case code.OpNull:
			err := vm.push(Null)
			if err != nil {
				return err
			}

		case code.OpSetGlobal:
			globalIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			vm.globals[globalIndex] = vm.pop()

		case code.OpGetGlobal:
			globalIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			err := vm.push(vm.globals[globalIndex])
			if err != nil {
				return err
			}

		case code.OpArray:
			numElements := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			array := vm.buildArray(vm.sp-numElements, vm.sp)
			vm.sp = vm.sp - numElements

			err := vm.push(array)
			if err != nil {
				return err
			}

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

		case code.OpIndex:
			index := vm.pop()
			left := vm.pop()

			err := vm.executeIndexExpression(left, index)
			if err != nil {
				return err
			}

		case code.OpCall:
			numArgs := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			err := vm.executeCall(int(numArgs))
			if err != nil {
				return err
			}

		case code.OpReturnValue:
			returnValue := vm.pop()

			frame := vm.popFrame()
			vm.sp = frame.basePointer - 1

			err := vm.push(returnValue)
			if err != nil {
				return err
			}

		case code.OpReturn:
			frame := vm.popFrame()
			vm.sp = frame.basePointer - 1

			err := vm.push(Null)
			if err != nil {
				return err
			}

		case code.OpSetLocal:
			localIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			frame := vm.currentFrame()

			vm.stack[frame.basePointer+int(localIndex)] = vm.pop()

		case code.OpGetLocal:
			localIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			frame := vm.currentFrame()

			err := vm.push(vm.stack[frame.basePointer+int(localIndex)])
			if err != nil {
				return err
			}

		case code.OpGetBuiltin:
			builtinIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			definition := object.Builtins[builtinIndex]

			err := vm.push(definition.Builtin)
			if err != nil {
				return err
			}

		case code.OpClosure:
			constIndex := code.ReadUint16(ins[ip+1:])
			numFree := code.ReadUint8(ins[ip+3:])
			vm.currentFrame().ip += 3

			err := vm.pushClosure(int(constIndex), int(numFree))
			if err != nil {
				return err
			}

		case code.OpGetFree:
			freeIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			currentClosure := vm.currentFrame().cl

			err := vm.push(currentClosure.Free[freeIndex])
			if err != nil {
				return err
			}

		case code.OpCurrentClosure:
			currentClosure := vm.currentFrame().cl

			err := vm.push(currentClosure)
			if err != nil {
				return err
			}

		case code.OpAssertType:
			constIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			expectedType := vm.constants[constIndex].(*object.String).Value
			value := vm.pop()

			if !vm.checkType(value, expectedType) {
				return fmt.Errorf("type error: expected %s, got %s", expectedType, vm.getTypeName(value))
			}

			// Push value back on stack after type check
			err := vm.push(value)
			if err != nil {
				return err
			}

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
			if err != nil {
				return err
			}

		case code.OpTypeCheck:
			constIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			expectedType := vm.constants[constIndex].(*object.String).Value
			value := vm.pop()

			// Push boolean result of type check
			result := vm.checkType(value, expectedType)
			if result {
				err := vm.push(True)
				if err != nil {
					return err
				}
			} else {
				err := vm.push(False)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

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

// LastPoppedStackElem returns the last popped stack element
func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	leftType := left.Type()
	rightType := right.Type()

	// String operations
	if leftType == object.STRING_OBJ && rightType == object.STRING_OBJ {
		return vm.executeBinaryStringOperation(op, left, right)
	}

	// Check if both operands are numeric types
	if vm.isNumericType(leftType) && vm.isNumericType(rightType) {
		return vm.executeBinaryNumericOperation(op, left, right)
	}

	return fmt.Errorf("unsupported types for binary operation: %s %s", leftType, rightType)
}

// isNumericType checks if a type is numeric
func (vm *VM) isNumericType(t object.ObjectType) bool {
	return t == object.INTEGER_OBJ || t == object.BYTE_OBJ || t == object.SHORT_OBJ ||
		t == object.INT_OBJ || t == object.LONG_OBJ || t == object.FLOAT_OBJ ||
		t == object.DOUBLE_OBJ || t == object.CHAR_OBJ
}

// isFloatingType checks if a type is floating-point
func (vm *VM) isFloatingType(t object.ObjectType) bool {
	return t == object.FLOAT_OBJ || t == object.DOUBLE_OBJ
}

// executeBinaryNumericOperation handles operations between any numeric types with type promotion
func (vm *VM) executeBinaryNumericOperation(op code.Opcode, left, right object.Object) error {
	leftType := left.Type()
	rightType := right.Type()

	// Determine if we need floating-point arithmetic
	isFloat := vm.isFloatingType(leftType) || vm.isFloatingType(rightType)

	// Bitwise operations require integer types
	if op == code.OpBitAnd || op == code.OpBitOr || op == code.OpBitXor ||
		op == code.OpLeftShift || op == code.OpRightShift {
		if isFloat {
			return fmt.Errorf("bitwise operations not supported for floating-point types")
		}
		return vm.executeBinaryIntegerOperation(op, left, right)
	}

	// For floating-point operations
	if isFloat {
		return vm.executeBinaryFloatOperation(op, left, right)
	}

	// For integer operations
	return vm.executeBinaryIntegerOperation(op, left, right)
}

// toInt64 extracts int64 value from any numeric type
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

// toFloat64 extracts float64 value from any numeric type
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

// promoteIntegerResult promotes integer result to appropriate type
func (vm *VM) promoteIntegerResult(result int64, left, right object.Object) object.Object {
	leftType := left.Type()
	rightType := right.Type()

	// Promotion rules: use the larger type
	// Long > Int > Short > Byte/Char
	if leftType == object.LONG_OBJ || rightType == object.LONG_OBJ {
		return &object.Long{Value: result}
	}
	if leftType == object.INT_OBJ || rightType == object.INT_OBJ {
		return &object.Int{Value: int32(result)}
	}
	if leftType == object.SHORT_OBJ || rightType == object.SHORT_OBJ {
		return &object.Short{Value: int16(result)}
	}
	// Default to Integer (legacy int64) for backward compatibility
	return &object.Integer{Value: result}
}

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
	case code.OpBitAnd:
		result = leftValue & rightValue
	case code.OpBitOr:
		result = leftValue | rightValue
	case code.OpBitXor:
		result = leftValue ^ rightValue
	case code.OpLeftShift:
		if rightValue < 0 {
			return fmt.Errorf("negative shift amount: %d", rightValue)
		}
		if rightValue >= 64 {
			return fmt.Errorf("shift amount too large: %d (must be less than 64)", rightValue)
		}
		result = leftValue << uint(rightValue)
	case code.OpRightShift:
		if rightValue < 0 {
			return fmt.Errorf("negative shift amount: %d", rightValue)
		}
		if rightValue >= 64 {
			return fmt.Errorf("shift amount too large: %d (must be less than 64)", rightValue)
		}
		result = leftValue >> uint(rightValue)
	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}

	return vm.push(vm.promoteIntegerResult(result, left, right))
}

// executeBinaryFloatOperation handles floating-point arithmetic
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
		// Floating-point modulo using fmod equivalent
		if rightValue == 0 {
			return fmt.Errorf("modulo by zero")
		}
		result = leftValue - rightValue*float64(int64(leftValue/rightValue))
	default:
		return fmt.Errorf("unknown float operator: %d", op)
	}

	// Promote to Double if either operand is Double, otherwise Float
	if left.Type() == object.DOUBLE_OBJ || right.Type() == object.DOUBLE_OBJ {
		return vm.push(&object.Double{Value: result})
	}
	return vm.push(&object.Float{Value: float32(result)})
}

func (vm *VM) executeBinaryStringOperation(
	op code.Opcode,
	left, right object.Object,
) error {
	if op != code.OpAdd {
		return fmt.Errorf("unknown string operator: %d", op)
	}

	leftValue := left.(*object.String).Value
	rightValue := right.(*object.String).Value

	return vm.push(&object.String{Value: leftValue + rightValue})
}

func (vm *VM) executeComparison(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	// Handle numeric comparisons
	if vm.isNumericType(left.Type()) && vm.isNumericType(right.Type()) {
		return vm.executeNumericComparison(op, left, right)
	}

	// Handle equality for non-numeric types
	switch op {
	case code.OpEqual:
		return vm.push(nativeBoolToBooleanObject(right == left))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObject(right != left))
	default:
		return fmt.Errorf("unknown operator: %d (%s %s)", op, left.Type(), right.Type())
	}
}

func (vm *VM) executeNumericComparison(
	op code.Opcode,
	left, right object.Object,
) error {
	// Use float comparison if either operand is floating-point
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
		default:
			return fmt.Errorf("unknown operator: %d", op)
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
	default:
		return fmt.Errorf("unknown operator: %d", op)
	}
}

// Legacy function - kept for compatibility (redirects to new function)
func (vm *VM) executeIntegerComparison(
	op code.Opcode,
	left, right object.Object,
) error {
	return vm.executeNumericComparison(op, left, right)
}

func (vm *VM) executeBangOperator() error {
	operand := vm.pop()

	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	case Null:
		return vm.push(True)
	default:
		return vm.push(False)
	}
}

func (vm *VM) executeAndOperator() error {
	right := vm.pop()
	left := vm.pop()

	// Both operands are already on stack, return true if both are truthy
	if isTruthy(left) && isTruthy(right) {
		return vm.push(True)
	}
	return vm.push(False)
}

func (vm *VM) executeOrOperator() error {
	right := vm.pop()
	left := vm.pop()

	// Return true if either operand is truthy
	if isTruthy(left) || isTruthy(right) {
		return vm.push(True)
	}
	return vm.push(False)
}

func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()

	if !vm.isNumericType(operand.Type()) {
		return fmt.Errorf("unsupported type for negation: %s", operand.Type())
	}

	// Handle floating-point negation
	if vm.isFloatingType(operand.Type()) {
		value := vm.toFloat64(operand)
		if operand.Type() == object.DOUBLE_OBJ {
			return vm.push(&object.Double{Value: -value})
		}
		return vm.push(&object.Float{Value: -float32(value)})
	}

	// Handle integer negation
	value := vm.toInt64(operand)
	switch operand.Type() {
	case object.BYTE_OBJ:
		return vm.push(&object.Byte{Value: -int8(value)})
	case object.SHORT_OBJ:
		return vm.push(&object.Short{Value: -int16(value)})
	case object.INT_OBJ:
		return vm.push(&object.Int{Value: -int32(value)})
	case object.LONG_OBJ:
		return vm.push(&object.Long{Value: -value})
	case object.CHAR_OBJ:
		return vm.push(&object.Char{Value: -rune(value)})
	default:
		return vm.push(&object.Integer{Value: -value})
	}
}

func (vm *VM) executeBitNotOperator() error {
	operand := vm.pop()

	if !vm.isNumericType(operand.Type()) {
		return fmt.Errorf("unsupported type for bitwise NOT: %s", operand.Type())
	}

	if vm.isFloatingType(operand.Type()) {
		return fmt.Errorf("bitwise NOT not supported for floating-point types")
	}

	value := vm.toInt64(operand)
	switch operand.Type() {
	case object.BYTE_OBJ:
		return vm.push(&object.Byte{Value: int8(^value)})
	case object.SHORT_OBJ:
		return vm.push(&object.Short{Value: int16(^value)})
	case object.INT_OBJ:
		return vm.push(&object.Int{Value: int32(^value)})
	case object.LONG_OBJ:
		return vm.push(&object.Long{Value: ^value})
	case object.CHAR_OBJ:
		return vm.push(&object.Char{Value: rune(^value)})
	default:
		return vm.push(&object.Integer{Value: ^value})
	}
}

func (vm *VM) buildArray(startIndex, endIndex int) object.Object {
	elements := make([]object.Object, endIndex-startIndex)

	for i := startIndex; i < endIndex; i++ {
		elements[i-startIndex] = vm.stack[i]
	}

	return &object.Array{Elements: elements}
}

func (vm *VM) buildHash(startIndex, endIndex int) (object.Object, error) {
	hashedPairs := make(map[object.HashKey]object.HashPair)

	for i := startIndex; i < endIndex; i += 2 {
		key := vm.stack[i]
		value := vm.stack[i+1]

		pair := object.HashPair{Key: key, Value: value}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", key.Type())
		}

		hashedPairs[hashKey.HashKey()] = pair
	}

	return &object.Hash{Pairs: hashedPairs}, nil
}

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

func (vm *VM) executeArrayIndex(array, index object.Object) error {
	arrayObject := array.(*object.Array)
	i := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if i < 0 || i > max {
		return vm.push(Null)
	}

	return vm.push(arrayObject.Elements[i])
}

func (vm *VM) executeHashIndex(hash, index object.Object) error {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return fmt.Errorf("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return vm.push(Null)
	}

	return vm.push(pair.Value)
}

func (vm *VM) executeCall(numArgs int) error {
	callee := vm.stack[vm.sp-1-numArgs]

	switch callee := callee.(type) {
	case *object.Closure:
		return vm.callClosure(callee, numArgs)
	case *object.Builtin:
		return vm.callBuiltin(callee, numArgs)
	default:
		return fmt.Errorf("calling non-function and non-builtin")
	}
}

func (vm *VM) callClosure(cl *object.Closure, numArgs int) error {
	if numArgs != cl.Fn.NumParameters {
		return fmt.Errorf("wrong number of arguments: want=%d, got=%d",
			cl.Fn.NumParameters, numArgs)
	}

	frame := NewFrame(cl, vm.sp-numArgs)
	vm.pushFrame(frame)

	vm.sp = frame.basePointer + cl.Fn.NumLocals

	return nil
}

func (vm *VM) callBuiltin(builtin *object.Builtin, numArgs int) error {
	args := vm.stack[vm.sp-numArgs : vm.sp]

	result := builtin.Fn(args...)
	vm.sp = vm.sp - numArgs - 1

	if result != nil {
		vm.push(result)
	} else {
		vm.push(Null)
	}

	return nil
}

func (vm *VM) pushClosure(constIndex int, numFree int) error {
	constant := vm.constants[constIndex]
	function, ok := constant.(*object.CompiledFunction)
	if !ok {
		return fmt.Errorf("not a function: %+v", constant)
	}

	free := make([]object.Object, numFree)
	for i := 0; i < numFree; i++ {
		free[i] = vm.stack[vm.sp-numFree+i]
	}
	vm.sp = vm.sp - numFree

	closure := &object.Closure{Fn: function, Free: free}
	return vm.push(closure)
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.Null:
		return false
	default:
		return true
	}
}

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

// Type checking and casting functions

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
	case object.CHAR_OBJ:
		return "অক্ষর"
	case object.STRING_OBJ:
		return "লেখা"
	case object.BOOLEAN_OBJ:
		return "বুলিয়ান"
	case object.ARRAY_OBJ:
		return "তালিকা"
	case object.HASH_OBJ:
		return "ম্যাপ"
	default:
		return string(obj.Type())
	}
}

func (vm *VM) checkType(obj object.Object, expectedType string) bool {
	actualType := vm.getTypeName(obj)
	return actualType == expectedType
}

func (vm *VM) castType(obj object.Object, targetType string) (object.Object, error) {
	// If already the correct type, return as-is
	if vm.checkType(obj, targetType) {
		return obj, nil
	}

	switch targetType {
	case "বাইট":
		val := vm.toInt64(obj)
		if val < 0 || val > 255 {
			return nil, fmt.Errorf("value %d out of range for byte (0-255)", val)
		}
		return &object.Byte{Value: int8(val)}, nil

	case "ছোট_সংখ্যা":
		val := vm.toInt64(obj)
		if val < -32768 || val > 32767 {
			return nil, fmt.Errorf("value %d out of range for short (-32768 to 32767)", val)
		}
		return &object.Short{Value: int16(val)}, nil

	case "পূর্ণসংখ্যা":
		val := vm.toInt64(obj)
		if val < -2147483648 || val > 2147483647 {
			return nil, fmt.Errorf("value %d out of range for int", val)
		}
		return &object.Int{Value: int32(val)}, nil

	case "দীর্ঘ_সংখ্যা":
		val := vm.toInt64(obj)
		return &object.Long{Value: val}, nil

	case "দশমিক":
		val := vm.toFloat64(obj)
		return &object.Float{Value: float32(val)}, nil

	case "দশমিক_দ্বিগুণ":
		val := vm.toFloat64(obj)
		return &object.Double{Value: val}, nil

	case "লেখা":
		return &object.String{Value: obj.Inspect()}, nil

	case "অক্ষর":
		str, ok := obj.(*object.String)
		if !ok {
			return nil, fmt.Errorf("cannot cast %s to char", obj.Type())
		}
		if len([]rune(str.Value)) != 1 {
			return nil, fmt.Errorf("string must be exactly one character to cast to char")
		}
		return &object.Char{Value: []rune(str.Value)[0]}, nil

	default:
		return nil, fmt.Errorf("cannot cast to type %s", targetType)
	}
}
