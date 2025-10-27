package evaluator

import (
	"bhasa/object"
	"fmt"
	"math"
	"strings"
)

var builtins = map[string]*object.Builtin{
	"লেখ": { // "write" - prints to console
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"দৈর্ঘ্য": { // "length" - returns length of string/array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len([]rune(arg.Value)))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to 'দৈর্ঘ্য' not supported, got %s", args[0].Type())
			}
		},
	},
	"প্রথম": { // "first" - returns first element of array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'প্রথম' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},
	"শেষ": { // "last" - returns last element of array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'শেষ' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"বাকি": { // "rest" - returns all but first element
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'বাকি' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}
			return NULL
		},
	},
	"যোগ": { // "push" - adds element to end of array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'যোগ' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"টাইপ": { // "type" - returns type of object
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: string(args[0].Type())}
		},
	},
	// String methods
	"বিভক্ত": { // "split" - splits string by delimiter
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to 'বিভক্ত' must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to 'বিভক্ত' must be STRING, got %s", args[1].Type())
			}
			str := args[0].(*object.String).Value
			delimiter := args[1].(*object.String).Value
			parts := strings.Split(str, delimiter)
			elements := make([]object.Object, len(parts))
			for i, part := range parts {
				elements[i] = &object.String{Value: part}
			}
			return &object.Array{Elements: elements}
		},
	},
	"যুক্ত": { // "join" - joins array elements with delimiter
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to 'যুক্ত' must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to 'যুক্ত' must be STRING, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			delimiter := args[1].(*object.String).Value
			strs := make([]string, len(arr.Elements))
			for i, elem := range arr.Elements {
				strs[i] = elem.Inspect()
			}
			return &object.String{Value: strings.Join(strs, delimiter)}
		},
	},
	"উপরে": { // "uppercase" - converts string to uppercase
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to 'উপরে' must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToUpper(str)}
		},
	},
	"নিচে": { // "lowercase" - converts string to lowercase
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to 'নিচে' must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToLower(str)}
		},
	},
	"ছাঁটো": { // "trim" - trims whitespace from string
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to 'ছাঁটো' must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.TrimSpace(str)}
		},
	},
	"প্রতিস্থাপন": { // "replace" - replaces substring in string
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to 'প্রতিস্থাপন' must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to 'প্রতিস্থাপন' must be STRING, got %s", args[1].Type())
			}
			if args[2].Type() != object.STRING_OBJ {
				return newError("third argument to 'প্রতিস্থাপন' must be STRING, got %s", args[2].Type())
			}
			str := args[0].(*object.String).Value
			old := args[1].(*object.String).Value
			new := args[2].(*object.String).Value
			return &object.String{Value: strings.ReplaceAll(str, old, new)}
		},
	},
	"খুঁজুন": { // "find" - finds index of substring in string
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to 'খুঁজুন' must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to 'খুঁজুন' must be STRING, got %s", args[1].Type())
			}
			str := args[0].(*object.String).Value
			substr := args[1].(*object.String).Value
			index := strings.Index(str, substr)
			return &object.Integer{Value: int64(index)}
		},
	},
	// Math functions
	"শক্তি": { // "power" - power function
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("first argument to 'শক্তি' must be INTEGER, got %s", args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to 'শক্তি' must be INTEGER, got %s", args[1].Type())
			}
			base := float64(args[0].(*object.Integer).Value)
			exp := float64(args[1].(*object.Integer).Value)
			result := math.Pow(base, exp)
			return &object.Integer{Value: int64(result)}
		},
	},
	"বর্গমূল": { // "square root" - square root function
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to 'বর্গমূল' must be INTEGER, got %s", args[0].Type())
			}
			n := float64(args[0].(*object.Integer).Value)
			if n < 0 {
				return newError("cannot take square root of negative number")
			}
			result := math.Sqrt(n)
			return &object.Integer{Value: int64(result)}
		},
	},
	"পরম": { // "absolute" - absolute value
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to 'পরম' must be INTEGER, got %s", args[0].Type())
			}
			n := args[0].(*object.Integer).Value
			if n < 0 {
				n = -n
			}
			return &object.Integer{Value: n}
		},
	},
	"সর্বোচ্চ": { // "max" - maximum of two numbers
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("first argument to 'সর্বোচ্চ' must be INTEGER, got %s", args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to 'সর্বোচ্চ' must be INTEGER, got %s", args[1].Type())
			}
			a := args[0].(*object.Integer).Value
			b := args[1].(*object.Integer).Value
			if a > b {
				return &object.Integer{Value: a}
			}
			return &object.Integer{Value: b}
		},
	},
	"সর্বনিম্ন": { // "min" - minimum of two numbers
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("first argument to 'সর্বনিম্ন' must be INTEGER, got %s", args[0].Type())
			}
			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to 'সর্বনিম্ন' must be INTEGER, got %s", args[1].Type())
			}
			a := args[0].(*object.Integer).Value
			b := args[1].(*object.Integer).Value
			if a < b {
				return &object.Integer{Value: a}
			}
			return &object.Integer{Value: b}
		},
	},
	"গোলাকার": { // "round" - round to nearest integer
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.INTEGER_OBJ {
				return newError("argument to 'গোলাকার' must be INTEGER, got %s", args[0].Type())
			}
			n := float64(args[0].(*object.Integer).Value)
			result := math.Round(n)
			return &object.Integer{Value: int64(result)}
		},
	},
	// Array methods
	"সাজাও": { // "sort" - sort array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'সাজাও' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			
			// Create a copy of the array
			sorted := make([]object.Object, length)
			copy(sorted, arr.Elements)
			
			// Simple bubble sort for integers
			for i := 0; i < length; i++ {
				for j := i + 1; j < length; j++ {
					if sorted[i].Type() != object.INTEGER_OBJ || sorted[j].Type() != object.INTEGER_OBJ {
						return newError("সাজাও can only sort arrays of integers")
					}
					vi := sorted[i].(*object.Integer).Value
					vj := sorted[j].(*object.Integer).Value
					if vi > vj {
						sorted[i], sorted[j] = sorted[j], sorted[i]
					}
				}
			}
			return &object.Array{Elements: sorted}
		},
	},
	"উল্টাও": { // "reverse" - reverse array
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'উল্টাও' must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			reversed := make([]object.Object, length)
			for i := 0; i < length; i++ {
				reversed[i] = arr.Elements[length-1-i]
			}
			return &object.Array{Elements: reversed}
		},
	},
}

