package object

import (
	"bhasa/ast"
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"os"
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
		"বিভক্ত", // split
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ || args[1].Type() != STRING_OBJ {
				return &Error{Message: "arguments to 'বিভক্ত' must be STRING"}
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
		"যুক্ত", // join
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ || args[1].Type() != STRING_OBJ {
				return &Error{Message: "first argument must be ARRAY, second must be STRING"}
			}
			arr := args[0].(*Array)
			delimiter := args[1].(*String).Value
			parts := make([]string, len(arr.Elements))
			for i, elem := range arr.Elements {
				parts[i] = elem.Inspect()
			}
			return &String{Value: strings.Join(parts, delimiter)}
		}},
	},
	{
		"উপরে", // uppercase
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'উপরে' must be STRING"}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.ToUpper(str)}
		}},
	},
	{
		"নিচে", // lowercase
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'নিচে' must be STRING"}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.ToLower(str)}
		}},
	},
	{
		"ছাঁটো", // trim
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'ছাঁটো' must be STRING"}
			}
			str := args[0].(*String).Value
			return &String{Value: strings.TrimSpace(str)}
		}},
	},
	{
		"প্রতিস্থাপন", // replace
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 3 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=3", len(args))}
			}
			if args[0].Type() != STRING_OBJ || args[1].Type() != STRING_OBJ || args[2].Type() != STRING_OBJ {
				return &Error{Message: "all arguments to 'প্রতিস্থাপন' must be STRING"}
			}
			str := args[0].(*String).Value
			old := args[1].(*String).Value
			new := args[2].(*String).Value
			return &String{Value: strings.ReplaceAll(str, old, new)}
		}},
	},
	{
		"খুঁজুন", // find/indexOf
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ || args[1].Type() != STRING_OBJ {
				return &Error{Message: "arguments to 'খুঁজুন' must be STRING"}
			}
			str := args[0].(*String).Value
			substr := args[1].(*String).Value
			return &Integer{Value: int64(strings.Index(str, substr))}
		}},
	},
	// Math functions
	{
		"শক্তি", // power
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ || args[1].Type() != INTEGER_OBJ {
				return &Error{Message: "arguments to 'শক্তি' must be INTEGER"}
			}
			base := float64(args[0].(*Integer).Value)
			exp := float64(args[1].(*Integer).Value)
			result := math.Pow(base, exp)
			return &Integer{Value: int64(result)}
		}},
	},
	{
		"বর্গমূল", // square root
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: "argument to 'বর্গমূল' must be INTEGER"}
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
		"পরম", // absolute value
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: "argument to 'পরম' must be INTEGER"}
			}
			n := args[0].(*Integer).Value
			if n < 0 {
				return &Integer{Value: -n}
			}
			return &Integer{Value: n}
		}},
	},
	{
		"সর্বোচ্চ", // max
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ || args[1].Type() != INTEGER_OBJ {
				return &Error{Message: "arguments to 'সর্বোচ্চ' must be INTEGER"}
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
		"সর্বনিম্ন", // min
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ || args[1].Type() != INTEGER_OBJ {
				return &Error{Message: "arguments to 'সর্বনিম্ন' must be INTEGER"}
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
		"গোলাকার", // round
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != INTEGER_OBJ {
				return &Error{Message: "argument to 'গোলাকার' must be INTEGER"}
			}
			// For integers, round returns the same value
			return args[0]
		}},
	},
	// Array methods
	{
		"উল্টাও", // reverse
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "argument to 'উল্টাও' must be ARRAY"}
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
	{
		"সাজাও", // sort - sorts integers in ascending order
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "argument to 'সাজাও' must be ARRAY"}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			
			// Create a copy to avoid modifying original
			sorted := make([]Object, length)
			copy(sorted, arr.Elements)
			
			// Simple bubble sort for integers
			for i := 0; i < length-1; i++ {
				for j := 0; j < length-i-1; j++ {
					// Only sort integers
					if sorted[j].Type() == INTEGER_OBJ && sorted[j+1].Type() == INTEGER_OBJ {
						v1 := sorted[j].(*Integer).Value
						v2 := sorted[j+1].(*Integer).Value
						if v1 > v2 {
							sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
						}
					}
				}
			}
			return &Array{Elements: sorted}
		}},
	},
	{
		"ফিল্টার", // filter - filters array based on function
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "first argument to 'ফিল্টার' must be ARRAY"}
			}
			if args[1].Type() != CLOSURE_OBJ {
				return &Error{Message: "second argument to 'ফিল্টার' must be FUNCTION"}
			}
			
			// Note: We can't actually execute the closure here since we don't have VM access
			// This is a limitation - filter needs to be implemented differently
			// For now, return error
			return &Error{Message: "ফিল্টার: function execution not yet supported in built-ins"}
		}},
	},
	{
		"ম্যাপ", // map - applies function to each element
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "first argument to 'ম্যাপ' must be ARRAY"}
			}
			if args[1].Type() != CLOSURE_OBJ {
				return &Error{Message: "second argument to 'ম্যাপ' must be FUNCTION"}
			}
			
			// Note: Same limitation as filter - needs VM access to execute closures
			return &Error{Message: "ম্যাপ: function execution not yet supported in built-ins"}
		}},
	},
	// File I/O functions
	{
		"ফাইল_পড়ো", // read file
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'ফাইল_পড়ো' must be STRING"}
			}
			
			filename := args[0].(*String).Value
			content, err := os.ReadFile(filename)
			if err != nil {
				return &Error{Message: fmt.Sprintf("error reading file: %s", err)}
			}
			
			return &String{Value: string(content)}
		}},
	},
	{
		"ফাইল_লেখো", // write file
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ || args[1].Type() != STRING_OBJ {
				return &Error{Message: "both arguments to 'ফাইল_লেখো' must be STRING"}
			}
			
			filename := args[0].(*String).Value
			content := args[1].(*String).Value
			
			err := os.WriteFile(filename, []byte(content), 0644)
			if err != nil {
				return &Error{Message: fmt.Sprintf("error writing file: %s", err)}
			}
			
			return &Null{}
		}},
	},
	{
		"ফাইল_যোগ", // append to file
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args))}
			}
			if args[0].Type() != STRING_OBJ || args[1].Type() != STRING_OBJ {
				return &Error{Message: "both arguments to 'ফাইল_যোগ' must be STRING"}
			}
			
			filename := args[0].(*String).Value
			content := args[1].(*String).Value
			
			f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return &Error{Message: fmt.Sprintf("error opening file: %s", err)}
			}
			defer f.Close()
			
			if _, err := f.WriteString(content); err != nil {
				return &Error{Message: fmt.Sprintf("error appending to file: %s", err)}
			}
			
			return &Null{}
		}},
	},
	{
		"ফাইল_আছে", // check if file exists
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'ফাইল_আছে' must be STRING"}
			}
			
			filename := args[0].(*String).Value
			_, err := os.Stat(filename)
			
			if os.IsNotExist(err) {
				return &Boolean{Value: false}
			}
			return &Boolean{Value: true}
		}},
	},
	// JSON functions
	{
		"JSON_পার্স", // parse JSON string
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argument to 'JSON_পার্স' must be STRING"}
			}
			
			jsonStr := args[0].(*String).Value
			var data interface{}
			
			err := json.Unmarshal([]byte(jsonStr), &data)
			if err != nil {
				return &Error{Message: fmt.Sprintf("error parsing JSON: %s", err)}
			}
			
			// Convert JSON data to Bhasa objects
			return jsonToObject(data)
		}},
	},
	{
		"JSON_স্ট্রিং", // stringify object to JSON
		&Builtin{Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
			}
			
			// Convert Bhasa object to JSON-compatible structure
			data := objectToJSON(args[0])
			
			jsonBytes, err := json.Marshal(data)
			if err != nil {
				return &Error{Message: fmt.Sprintf("error creating JSON: %s", err)}
			}
			
			return &String{Value: string(jsonBytes)}
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

// jsonToObject converts JSON data to Bhasa objects
func jsonToObject(data interface{}) Object {
	switch v := data.(type) {
	case nil:
		return &Null{}
	case bool:
		return &Boolean{Value: v}
	case float64:
		return &Integer{Value: int64(v)}
	case string:
		return &String{Value: v}
	case []interface{}:
		elements := make([]Object, len(v))
		for i, item := range v {
			elements[i] = jsonToObject(item)
		}
		return &Array{Elements: elements}
	case map[string]interface{}:
		pairs := make(map[HashKey]HashPair)
		for key, value := range v {
			keyObj := &String{Value: key}
			valueObj := jsonToObject(value)
			pairs[keyObj.HashKey()] = HashPair{Key: keyObj, Value: valueObj}
		}
		return &Hash{Pairs: pairs}
	default:
		return &Null{}
	}
}

// objectToJSON converts Bhasa objects to JSON-compatible structure
func objectToJSON(obj Object) interface{} {
	switch o := obj.(type) {
	case *Null:
		return nil
	case *Boolean:
		return o.Value
	case *Integer:
		return o.Value
	case *String:
		return o.Value
	case *Array:
		result := make([]interface{}, len(o.Elements))
		for i, elem := range o.Elements {
			result[i] = objectToJSON(elem)
		}
		return result
	case *Hash:
		result := make(map[string]interface{})
		for _, pair := range o.Pairs {
			if key, ok := pair.Key.(*String); ok {
				result[key.Value] = objectToJSON(pair.Value)
			}
		}
		return result
	default:
		return nil
	}
}

