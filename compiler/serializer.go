package compiler

import (
	"bhasa/code"
	"bhasa/object"
	"encoding/binary"
	"fmt"
	"io"
)

// Magic number for Bhasa bytecode files: "BHASA" in hex
const (
	MagicNumber uint32 = 0x42484153 // "BHAS"
	Version     uint32 = 1
)

// Serialize writes the bytecode to a writer in binary format
func (b *Bytecode) Serialize(w io.Writer) error {
	// Write magic number
	if err := binary.Write(w, binary.BigEndian, MagicNumber); err != nil {
		return fmt.Errorf("failed to write magic number: %w", err)
	}

	// Write version
	if err := binary.Write(w, binary.BigEndian, Version); err != nil {
		return fmt.Errorf("failed to write version: %w", err)
	}

	// Write instructions length
	instructionsLen := uint32(len(b.Instructions))
	if err := binary.Write(w, binary.BigEndian, instructionsLen); err != nil {
		return fmt.Errorf("failed to write instructions length: %w", err)
	}

	// Write instructions
	if _, err := w.Write(b.Instructions); err != nil {
		return fmt.Errorf("failed to write instructions: %w", err)
	}

	// Write constants count
	constantsCount := uint32(len(b.Constants))
	if err := binary.Write(w, binary.BigEndian, constantsCount); err != nil {
		return fmt.Errorf("failed to write constants count: %w", err)
	}

	// Write each constant
	for i, constant := range b.Constants {
		if err := serializeObject(w, constant); err != nil {
			return fmt.Errorf("failed to serialize constant %d: %w", i, err)
		}
	}

	return nil
}

// Deserialize reads bytecode from a reader
func Deserialize(r io.Reader) (*Bytecode, error) {
	// Read and verify magic number
	var magic uint32
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return nil, fmt.Errorf("failed to read magic number: %w", err)
	}
	if magic != MagicNumber {
		return nil, fmt.Errorf("invalid magic number: expected 0x%X, got 0x%X", MagicNumber, magic)
	}

	// Read and verify version
	var version uint32
	if err := binary.Read(r, binary.BigEndian, &version); err != nil {
		return nil, fmt.Errorf("failed to read version: %w", err)
	}
	if version != Version {
		return nil, fmt.Errorf("unsupported bytecode version: expected %d, got %d", Version, version)
	}

	// Read instructions length
	var instructionsLen uint32
	if err := binary.Read(r, binary.BigEndian, &instructionsLen); err != nil {
		return nil, fmt.Errorf("failed to read instructions length: %w", err)
	}

	// Read instructions
	instructions := make(code.Instructions, instructionsLen)
	if _, err := io.ReadFull(r, instructions); err != nil {
		return nil, fmt.Errorf("failed to read instructions: %w", err)
	}

	// Read constants count
	var constantsCount uint32
	if err := binary.Read(r, binary.BigEndian, &constantsCount); err != nil {
		return nil, fmt.Errorf("failed to read constants count: %w", err)
	}

	// Read each constant
	constants := make([]object.Object, constantsCount)
	for i := uint32(0); i < constantsCount; i++ {
		constant, err := deserializeObject(r)
		if err != nil {
			return nil, fmt.Errorf("failed to deserialize constant %d: %w", i, err)
		}
		constants[i] = constant
	}

	return &Bytecode{
		Instructions: instructions,
		Constants:    constants,
	}, nil
}

// Object type identifiers for serialization
const (
	objTypeInteger         byte = 1
	objTypeByte            byte = 2
	objTypeShort           byte = 3
	objTypeInt             byte = 4
	objTypeLong            byte = 5
	objTypeFloat           byte = 6
	objTypeDouble          byte = 7
	objTypeChar            byte = 8
	objTypeBoolean         byte = 9
	objTypeString          byte = 10
	objTypeNull            byte = 11
	objTypeCompiledFunc    byte = 12
	objTypeArray           byte = 13
	objTypeHash            byte = 14
)

// serializeObject writes an object to the writer
func serializeObject(w io.Writer, obj object.Object) error {
	switch o := obj.(type) {
	case *object.Integer:
		if err := binary.Write(w, binary.BigEndian, objTypeInteger); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Byte:
		if err := binary.Write(w, binary.BigEndian, objTypeByte); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Short:
		if err := binary.Write(w, binary.BigEndian, objTypeShort); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Int:
		if err := binary.Write(w, binary.BigEndian, objTypeInt); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Long:
		if err := binary.Write(w, binary.BigEndian, objTypeLong); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Float:
		if err := binary.Write(w, binary.BigEndian, objTypeFloat); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Double:
		if err := binary.Write(w, binary.BigEndian, objTypeDouble); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Char:
		if err := binary.Write(w, binary.BigEndian, objTypeChar); err != nil {
			return err
		}
		return binary.Write(w, binary.BigEndian, o.Value)

	case *object.Boolean:
		if err := binary.Write(w, binary.BigEndian, objTypeBoolean); err != nil {
			return err
		}
		var val byte
		if o.Value {
			val = 1
		}
		return binary.Write(w, binary.BigEndian, val)

	case *object.String:
		if err := binary.Write(w, binary.BigEndian, objTypeString); err != nil {
			return err
		}
		// Write string length
		strLen := uint32(len(o.Value))
		if err := binary.Write(w, binary.BigEndian, strLen); err != nil {
			return err
		}
		// Write string bytes
		_, err := w.Write([]byte(o.Value))
		return err

	case *object.Null:
		return binary.Write(w, binary.BigEndian, objTypeNull)

	case *object.CompiledFunction:
		if err := binary.Write(w, binary.BigEndian, objTypeCompiledFunc); err != nil {
			return err
		}
		// Write instructions length
		insLen := uint32(len(o.Instructions))
		if err := binary.Write(w, binary.BigEndian, insLen); err != nil {
			return err
		}
		// Write instructions
		if _, err := w.Write(o.Instructions); err != nil {
			return err
		}
		// Write NumLocals
		if err := binary.Write(w, binary.BigEndian, uint32(o.NumLocals)); err != nil {
			return err
		}
		// Write NumParameters
		return binary.Write(w, binary.BigEndian, uint32(o.NumParameters))

	case *object.Array:
		if err := binary.Write(w, binary.BigEndian, objTypeArray); err != nil {
			return err
		}
		// Write array length
		arrLen := uint32(len(o.Elements))
		if err := binary.Write(w, binary.BigEndian, arrLen); err != nil {
			return err
		}
		// Write each element
		for _, elem := range o.Elements {
			if err := serializeObject(w, elem); err != nil {
				return err
			}
		}
		return nil

	case *object.Hash:
		if err := binary.Write(w, binary.BigEndian, objTypeHash); err != nil {
			return err
		}
		// Write hash length
		hashLen := uint32(len(o.Pairs))
		if err := binary.Write(w, binary.BigEndian, hashLen); err != nil {
			return err
		}
		// Write each key-value pair
		for _, pair := range o.Pairs {
			if err := serializeObject(w, pair.Key); err != nil {
				return err
			}
			if err := serializeObject(w, pair.Value); err != nil {
				return err
			}
		}
		return nil

	default:
		return fmt.Errorf("unsupported object type for serialization: %s", obj.Type())
	}
}

// deserializeObject reads an object from the reader
func deserializeObject(r io.Reader) (object.Object, error) {
	var objType byte
	if err := binary.Read(r, binary.BigEndian, &objType); err != nil {
		return nil, err
	}

	switch objType {
	case objTypeInteger:
		var value int64
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Integer{Value: value}, nil

	case objTypeByte:
		var value int8
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Byte{Value: value}, nil

	case objTypeShort:
		var value int16
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Short{Value: value}, nil

	case objTypeInt:
		var value int32
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Int{Value: value}, nil

	case objTypeLong:
		var value int64
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Long{Value: value}, nil

	case objTypeFloat:
		var value float32
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Float{Value: value}, nil

	case objTypeDouble:
		var value float64
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Double{Value: value}, nil

	case objTypeChar:
		var value rune
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Char{Value: value}, nil

	case objTypeBoolean:
		var value byte
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return nil, err
		}
		return &object.Boolean{Value: value == 1}, nil

	case objTypeString:
		// Read string length
		var strLen uint32
		if err := binary.Read(r, binary.BigEndian, &strLen); err != nil {
			return nil, err
		}
		// Read string bytes
		strBytes := make([]byte, strLen)
		if _, err := io.ReadFull(r, strBytes); err != nil {
			return nil, err
		}
		return &object.String{Value: string(strBytes)}, nil

	case objTypeNull:
		return &object.Null{}, nil

	case objTypeCompiledFunc:
		// Read instructions length
		var insLen uint32
		if err := binary.Read(r, binary.BigEndian, &insLen); err != nil {
			return nil, err
		}
		// Read instructions
		instructions := make([]byte, insLen)
		if _, err := io.ReadFull(r, instructions); err != nil {
			return nil, err
		}
		// Read NumLocals
		var numLocals uint32
		if err := binary.Read(r, binary.BigEndian, &numLocals); err != nil {
			return nil, err
		}
		// Read NumParameters
		var numParameters uint32
		if err := binary.Read(r, binary.BigEndian, &numParameters); err != nil {
			return nil, err
		}
		return &object.CompiledFunction{
			Instructions:  instructions,
			NumLocals:     int(numLocals),
			NumParameters: int(numParameters),
		}, nil

	case objTypeArray:
		// Read array length
		var arrLen uint32
		if err := binary.Read(r, binary.BigEndian, &arrLen); err != nil {
			return nil, err
		}
		// Read each element
		elements := make([]object.Object, arrLen)
		for i := uint32(0); i < arrLen; i++ {
			elem, err := deserializeObject(r)
			if err != nil {
				return nil, err
			}
			elements[i] = elem
		}
		return &object.Array{Elements: elements}, nil

	case objTypeHash:
		// Read hash length
		var hashLen uint32
		if err := binary.Read(r, binary.BigEndian, &hashLen); err != nil {
			return nil, err
		}
		// Read each key-value pair
		pairs := make(map[object.HashKey]object.HashPair)
		for i := uint32(0); i < hashLen; i++ {
			key, err := deserializeObject(r)
			if err != nil {
				return nil, err
			}
			value, err := deserializeObject(r)
			if err != nil {
				return nil, err
			}
			// Get hash key
			hashable, ok := key.(object.Hashable)
			if !ok {
				return nil, fmt.Errorf("key is not hashable: %s", key.Type())
			}
			pairs[hashable.HashKey()] = object.HashPair{Key: key, Value: value}
		}
		return &object.Hash{Pairs: pairs}, nil

	default:
		return nil, fmt.Errorf("unknown object type in bytecode: %d", objType)
	}
}
