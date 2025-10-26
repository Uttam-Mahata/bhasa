package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Instructions is a bytecode instruction sequence
type Instructions []byte

// Opcode represents a bytecode instruction
type Opcode byte

// Bytecode instructions
const (
	OpConstant Opcode = iota // Load constant onto stack
	OpPop                    // Pop value from stack
	OpAdd                    // Add two values
	OpSub                    // Subtract
	OpMul                    // Multiply
	OpDiv                    // Divide
	OpMod                    // Modulo
	OpTrue                   // Push true
	OpFalse                  // Push false
	OpEqual                  // Equal comparison
	OpNotEqual               // Not equal comparison
	OpGreaterThan            // Greater than
	OpGreaterThanEqual       // Greater than or equal
	OpMinus                  // Unary minus
	OpBang                   // Logical not
	OpJumpNotTruthy          // Conditional jump
	OpJump                   // Unconditional jump
	OpNull                   // Push null
	OpGetGlobal              // Get global variable
	OpSetGlobal              // Set global variable
	OpArray                  // Create array
	OpHash                   // Create hash
	OpIndex                  // Index operation
	OpCall                   // Function call
	OpReturnValue            // Return with value
	OpReturn                 // Return without value
	OpGetLocal               // Get local variable
	OpSetLocal               // Set local variable
	OpGetBuiltin             // Get builtin function
	OpClosure                // Create closure
	OpGetFree                // Get free variable
	OpCurrentClosure         // Get current closure
)

// Definition holds information about an opcode
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:         {"OpConstant", []int{2}},
	OpPop:              {"OpPop", []int{}},
	OpAdd:              {"OpAdd", []int{}},
	OpSub:              {"OpSub", []int{}},
	OpMul:              {"OpMul", []int{}},
	OpDiv:              {"OpDiv", []int{}},
	OpMod:              {"OpMod", []int{}},
	OpTrue:             {"OpTrue", []int{}},
	OpFalse:            {"OpFalse", []int{}},
	OpEqual:            {"OpEqual", []int{}},
	OpNotEqual:         {"OpNotEqual", []int{}},
	OpGreaterThan:      {"OpGreaterThan", []int{}},
	OpGreaterThanEqual: {"OpGreaterThanEqual", []int{}},
	OpMinus:            {"OpMinus", []int{}},
	OpBang:             {"OpBang", []int{}},
	OpJumpNotTruthy:    {"OpJumpNotTruthy", []int{2}},
	OpJump:             {"OpJump", []int{2}},
	OpNull:             {"OpNull", []int{}},
	OpGetGlobal:        {"OpGetGlobal", []int{2}},
	OpSetGlobal:        {"OpSetGlobal", []int{2}},
	OpArray:            {"OpArray", []int{2}},
	OpHash:             {"OpHash", []int{2}},
	OpIndex:            {"OpIndex", []int{}},
	OpCall:             {"OpCall", []int{1}},
	OpReturnValue:      {"OpReturnValue", []int{}},
	OpReturn:           {"OpReturn", []int{}},
	OpGetLocal:         {"OpGetLocal", []int{1}},
	OpSetLocal:         {"OpSetLocal", []int{1}},
	OpGetBuiltin:       {"OpGetBuiltin", []int{1}},
	OpClosure:          {"OpClosure", []int{2, 1}},
	OpGetFree:          {"OpGetFree", []int{1}},
	OpCurrentClosure:   {"OpCurrentClosure", []int{}},
}

// Lookup returns the definition for an opcode
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make creates a bytecode instruction
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1
	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 1:
			instruction[offset] = byte(o)
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction
}

// ReadOperands reads operands from an instruction
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 1:
			operands[i] = int(ins[offset])
		case 2:
			operands[i] = int(binary.BigEndian.Uint16(ins[offset:]))
		}
		offset += width
	}

	return operands, offset
}

// String converts instructions to string for debugging
func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

// ReadUint16 reads a uint16 from instructions
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// ReadUint8 reads a uint8 from instructions
func ReadUint8(ins Instructions) uint8 {
	return uint8(ins[0])
}

