package object

import (
	"bhasa/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"math"
	"strings"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	INTEGER_OBJ           = "INTEGER"
	BOOLEAN_OBJ           = "BOOLEAN"
	STRING_OBJ            = "STRING"
	NULL_OBJ              = "NULL"
	RETURN_VALUE_OBJ      = "RETURN_VALUE"
	ERROR_OBJ             = "ERROR"
	FUNCTION_OBJ          = "FUNCTION"
	BUILTIN_OBJ           = "BUILTIN"
	ARRAY_OBJ             = "ARRAY"
	HASH_OBJ              = "HASH"
	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION"
	CLOSURE_OBJ           = "CLOSURE"
)

// Object represents a value in the language
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represents an integer value
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Null represents a null value
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

// ReturnValue wraps a return value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error represents an error
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

// Function represents a function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("ফাংশন")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// BuiltinFunction represents a built-in function
type BuiltinFunction func(args ...Object) Object

// Builtin represents a built-in function
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Array represents an array
type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// HashKey represents a hashable key
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashPair represents a key-value pair in a hash
type HashPair struct {
	Key   Object
	Value Object
}

// Hash represents a hash map
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// Hashable interface for objects that can be hashed
type Hashable interface {
	HashKey() HashKey
}

// HashKey methods for hashable types
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Environment represents a variable environment
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new enclosed environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a variable from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// CompiledFunction represents a compiled function
type CompiledFunction struct {
	Instructions  []byte
	NumLocals     int
	NumParameters int
}

func (cf *CompiledFunction) Type() ObjectType { return COMPILED_FUNCTION_OBJ }
func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

// Closure represents a closure
type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() ObjectType { return CLOSURE_OBJ }
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}

// BuiltinDef represents a builtin definition
type BuiltinDef struct {
	Name    string
	Builtin *Builtin
}

// Builtins is the list of builtin functions
var Builtins = []BuiltinDef{
	{
		"লেখ",
		&Builtin{Fn: func(args ...Object) Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return &Null{}
		}},
	},
	{
		"দৈর্ঘ্য",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			switch arg := args[0].(type) {
			case *String:
				return &Integer{Value: int64(len([]rune(arg.Value)))}
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			default:
				return &Error{Message: fmt.Sprintf("argument to 'দৈর্ঘ্য' not supported, got %s", args[0].Type())}
			}
		}},
	},
	{
		"প্রথম",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'প্রথম' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return &Null{}
		}},
	},
	{
		"শেষ",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'শেষ' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return &Null{}
		}},
	},
	{
		"বাকি",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'বাকি' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &Array{Elements: newElements}
			}
			return &Null{}
		}},
	},
	{
		"যোগ",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'যোগ' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			newElements := make([]Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &Array{Elements: newElements}
		}},
	},
	{
		"টাইপ",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			return &String{Value: string(args[0].Type())}
		}},
	},
	// String methods
	{
		"বিভক্ত",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'বিভক্ত' must be STRING, got %s", args[0].Type())}
			}
			if args[1].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'বিভক্ত' must be STRING, got %s", args[1].Type())}
			}
			str := args[0].(*String).Value
			delimiter := args[1].(*String).Value
			parts := strings.Split(str, delimiter)
			elements := make([]Object, len(parts))
			for i, part := range parts {
				elements[i] = &String{Value: part}
			}
			return &Array{Elements: elements}
		}},
	},
	{
		"যুক্ত",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'যুক্ত' must be ARRAY, got %s", args[0].Type())}
			}
			if args[1].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'যুক্ত' must be STRING, got %s", args[1].Type())}
			}
			arr := args[0].(*Array)
			delimiter := args[1].(*String).Value
			strs := make([]string, len(arr.Elements))
			for i, elem := range arr.Elements {
				strs[i] = elem.Inspect()
			}
			return &String{Value: strings.Join(strs, delimiter)}
		}},
	},
	{
		"উপরে",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'উপরে' must be STRING, got %s", args[0].Type())}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.ToUpper(str)}
		}},
	},
	{
		"নিচে",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'নিচে' must be STRING, got %s", args[0].Type())}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.ToLower(str)}
		}},
	},
	{
		"ছাঁটো",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'ছাঁটো' must be STRING, got %s", args[0].Type())}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.TrimSpace(str)}
		}},
	},
	{
		"প্রতিস্থাপন",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 3 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=3", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'প্রতিস্থাপন' must be STRING, got %s", args[0].Type())}
			}
			if args[1].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'প্রতিস্থাপন' must be STRING, got %s", args[1].Type())}
			}
			if args[2].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("third argument to 'প্রতিস্থাপন' must be STRING, got %s", args[2].Type())}
			}
			str := args[0].(*String).Value
			old := args[1].(*String).Value
			new := args[2].(*String).Value
			return &String{Value: strings.ReplaceAll(str, old, new)}
		}},
	},
	{
		"খুঁজুন",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'খুঁজুন' must be STRING, got %s", args[0].Type())}
			}
			if args[1].Type() != STRING_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'খুঁজুন' must be STRING, got %s", args[1].Type())}
			}
			str := args[0].(*String).Value
			substr := args[1].(*String).Value
			index := strings.Index(str, substr)
			return &Integer{Value: int64(index)}
		}},
	},
	// Math functions
	{
		"শক্তি",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'শক্তি' must be INTEGER, got %s", args[0].Type())}
			}
			if args[1].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'শক্তি' must be INTEGER, got %s", args[1].Type())}
			}
			base := float64(args[0].(*Integer).Value)
			exp := float64(args[1].(*Integer).Value)
			result := math.Pow(base, exp)
			return &Integer{Value: int64(result)}
		}},
	},
	{
		"বর্গমূল",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'বর্গমূল' must be INTEGER, got %s", args[0].Type())}
			}
			n := float64(args[0].(*Integer).Value)
			if n < 0 {
				return &Error{Message: "cannot take square root of negative number"}
			}
			result := math.Sqrt(n)
			return &Integer{Value: int64(result)}
		}},
	},
	{
		"পরম",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'পরম' must be INTEGER, got %s", args[0].Type())}
			}
			n := args[0].(*Integer).Value
			if n < 0 {
				n = -n
			}
			return &Integer{Value: n}
		}},
	},
	{
		"সর্বোচ্চ",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'সর্বোচ্চ' must be INTEGER, got %s", args[0].Type())}
			}
			if args[1].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'সর্বোচ্চ' must be INTEGER, got %s", args[1].Type())}
			}
			a := args[0].(*Integer).Value
			b := args[1].(*Integer).Value
			if a > b {
				return &Integer{Value: a}
			}
			return &Integer{Value: b}
		}},
	},
	{
		"সর্বনিম্ন",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("first argument to 'সর্বনিম্ন' must be INTEGER, got %s", args[0].Type())}
			}
			if args[1].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("second argument to 'সর্বনিম্ন' must be INTEGER, got %s", args[1].Type())}
			}
			a := args[0].(*Integer).Value
			b := args[1].(*Integer).Value
			if a < b {
				return &Integer{Value: a}
			}
			return &Integer{Value: b}
		}},
	},
	{
		"গোলাকার",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'গোলাকার' must be INTEGER, got %s", args[0].Type())}
			}
			n := float64(args[0].(*Integer).Value)
			result := math.Round(n)
			return &Integer{Value: int64(result)}
		}},
	},
	// Array methods
	{
		"সাজাও",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'সাজাও' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			
			// Create a copy of the array
			sorted := make([]Object, length)
			copy(sorted, arr.Elements)
			
			// Simple bubble sort for integers
			for i := 0; i < length; i++ {
				for j := i + 1; j < length; j++ {
					if sorted[i].Type() != INTEGER_OBJ || sorted[j].Type() != INTEGER_OBJ {
						return &Error{Message: "সাজাও can only sort arrays of integers"}
					}
					vi := sorted[i].(*Integer).Value
					vj := sorted[j].(*Integer).Value
					if vi > vj {
						sorted[i], sorted[j] = sorted[j], sorted[i]
					}
				}
			}
			return &Array{Elements: sorted}
		}},
	},
	{
		"উল্টাও",
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: fmt.Sprintf("argument to 'উল্টাও' must be ARRAY, got %s", args[0].Type())}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			reversed := make([]Object, length)
			for i := 0; i < length; i++ {
				reversed[i] = arr.Elements[length-1-i]
			}
			return &Array{Elements: reversed}
		}},
	},
}

// GetBuiltinByName returns a builtin by name
func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

