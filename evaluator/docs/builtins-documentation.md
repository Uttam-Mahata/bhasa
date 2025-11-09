# Bhasa Built-in Functions Documentation

## Table of Contents

1. [Overview](#overview)
2. [Function Reference](#function-reference)
3. [Implementation Details](#implementation-details)
4. [Usage Examples](#usage-examples)
5. [Error Handling](#error-handling)
6. [Adding New Builtins](#adding-new-builtins)

---

## Overview

Bhasa provides a set of built-in functions with Bengali names that are automatically available in the global scope. These functions are implemented in Go and provide essential functionality for I/O, array manipulation, and type inspection.

### Built-in Functions List

| Bengali | Transliteration | English | Purpose |
|---------|----------------|---------|---------|
| `লেখ` | lekha | write/print | Print to console |
| `দৈর্ঘ্য` | doirggho | length | Get length of string/array |
| `প্রথম` | prothom | first | Get first element of array |
| `শেষ` | shesh | last | Get last element of array |
| `বাকি` | baki | rest | Get all but first element |
| `যোগ` | jog | push | Add element to array |
| `টাইপ` | type | type | Get type of value |

---

## Function Reference

### লেখ (Print)

**Signature**:
```bhasa
লেখ(arg1, arg2, ..., argN)
```

**Purpose**: Prints values to console (stdout)

**Parameters**:
- Accepts any number of arguments
- Each argument can be any type

**Returns**: `NULL`

**Behavior**:
- Prints each argument on a new line
- Uses object's `Inspect()` method for string representation

**Implementation**:
```go
"লেখ": {
    Fn: func(args ...object.Object) object.Object {
        for _, arg := range args {
            fmt.Println(arg.Inspect())
        }
        return NULL
    },
},
```

**Examples**:
```bhasa
লেখ("হ্যালো বিশ্ব");
// Output: হ্যালো বিশ্ব

লেখ(5, 10, 15);
// Output:
// 5
// 10
// 15

লেখ([1, 2, 3]);
// Output: [1, 2, 3]

লেখ({"নাম": "রহিম"});
// Output: {নাম: রহিম}
```

**No Arguments**:
```bhasa
লেখ();
// Output: (empty - nothing printed)
```

---

### দৈর্ঘ্য (Length)

**Signature**:
```bhasa
দৈর্ঘ্য(value)
```

**Purpose**: Returns the length of a string or array

**Parameters**:
- `value`: Must be a string or array

**Returns**: 
- `INTEGER` - Length of the string/array
- `ERROR` - If wrong number of arguments or unsupported type

**Behavior**:
- For strings: Returns number of **Unicode characters** (runes), not bytes
- For arrays: Returns number of elements

**Implementation**:
```go
"দৈর্ঘ্য": {
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
```

**Examples**:
```bhasa
// Strings
দৈর্ঘ্য("hello");      // → 5
দৈর্ঘ্য("হ্যালো");      // → 5 (5 Unicode characters)
দৈর্ঘ্য("");           // → 0

// Arrays
দৈর্ঘ্য([1, 2, 3]);    // → 3
দৈর্ঘ্য([]);           // → 0
দৈর্ঘ্য([[1], [2]]);   // → 2

// Errors
দৈর্ঘ্য(5);            // ERROR: argument to 'দৈর্ঘ্য' not supported, got INTEGER
দৈর্ঘ্য();             // ERROR: wrong number of arguments. got=0, want=1
দৈর্ঘ্য("a", "b");     // ERROR: wrong number of arguments. got=2, want=1
```

**Unicode Handling**:
```bhasa
// Bengali text
ধরি text = "বাংলা";
দৈর্ঘ্য(text);  // → 5 (not byte count)

// Emoji
দৈর্ঘ্য("😀😁😂");  // → 3
```

---

### প্রথম (First)

**Signature**:
```bhasa
প্রথম(array)
```

**Purpose**: Returns the first element of an array

**Parameters**:
- `array`: Must be an array

**Returns**:
- The first element of the array
- `NULL` - If array is empty
- `ERROR` - If wrong number of arguments or not an array

**Implementation**:
```go
"প্রথম": {
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
```

**Examples**:
```bhasa
// Normal usage
প্রথম([1, 2, 3]);          // → 1
প্রথম(["a", "b", "c"]);    // → "a"
প্রথম([[1, 2], [3, 4]]);   // → [1, 2]

// Empty array
প্রথম([]);                 // → null

// Errors
প্রথম("hello");            // ERROR: argument to 'প্রথম' must be ARRAY
প্রথম(5);                  // ERROR: argument to 'প্রথম' must be ARRAY
প্রথম([1], [2]);           // ERROR: wrong number of arguments
```

**Use Case**:
```bhasa
ধরি numbers = [10, 20, 30];
যদি (দৈর্ঘ্য(numbers) > 0) {
    ধরি first = প্রথম(numbers);
    লেখ("First:", first);
}
```

---

### শেষ (Last)

**Signature**:
```bhasa
শেষ(array)
```

**Purpose**: Returns the last element of an array

**Parameters**:
- `array`: Must be an array

**Returns**:
- The last element of the array
- `NULL` - If array is empty
- `ERROR` - If wrong number of arguments or not an array

**Implementation**:
```go
"শেষ": {
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
```

**Examples**:
```bhasa
// Normal usage
শেষ([1, 2, 3]);           // → 3
শেষ(["a", "b", "c"]);     // → "c"
শেষ([[1, 2], [3, 4]]);    // → [3, 4]

// Single element
শেষ([42]);                // → 42

// Empty array
শেষ([]);                  // → null

// Errors
শেষ("hello");             // ERROR: argument to 'শেষ' must be ARRAY
শেষ();                    // ERROR: wrong number of arguments
```

**Use Case**:
```bhasa
ধরি items = [1, 2, 3, 4, 5];
ধরি last = শেষ(items);
লেখ("Last item:", last);
```

---

### বাকি (Rest)

**Signature**:
```bhasa
বাকি(array)
```

**Purpose**: Returns all elements except the first (tail of the list)

**Parameters**:
- `array`: Must be an array

**Returns**:
- New array with all elements except the first
- `NULL` - If array is empty
- `ERROR` - If wrong number of arguments or not an array

**Behavior**:
- Creates a **new array** (does not modify original)
- Original array is unchanged

**Implementation**:
```go
"বাকি": {
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
```

**Examples**:
```bhasa
// Normal usage
বাকি([1, 2, 3]);          // → [2, 3]
বাকি(["a", "b", "c"]);    // → ["b", "c"]

// Single element
বাকি([42]);               // → []

// Empty array
বাকি([]);                 // → null

// Original unchanged
ধরি arr = [1, 2, 3];
ধরি rest = বাকি(arr);
লেখ(arr);                 // → [1, 2, 3] (unchanged)
লেখ(rest);                // → [2, 3]

// Errors
বাকি("hello");            // ERROR: argument to 'বাকি' must be ARRAY
বাকি([1], [2]);           // ERROR: wrong number of arguments
```

**Use Case - Recursive List Processing**:
```bhasa
ধরি sum = ফাংশন(arr) {
    যদি (দৈর্ঘ্য(arr) == 0) {
        ফেরত 0;
    }
    প্রথম(arr) + sum(বাকি(arr))
};

sum([1, 2, 3, 4, 5]);  // → 15
```

---

### যোগ (Push)

**Signature**:
```bhasa
যোগ(array, element)
```

**Purpose**: Adds an element to the end of an array

**Parameters**:
- `array`: Must be an array
- `element`: Any value to add

**Returns**:
- New array with element added at the end
- `ERROR` - If wrong number of arguments or first arg is not an array

**Behavior**:
- Creates a **new array** (does not modify original)
- Original array is unchanged
- Element can be of any type

**Implementation**:
```go
"যোগ": {
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
```

**Examples**:
```bhasa
// Normal usage
যোগ([1, 2], 3);           // → [1, 2, 3]
যোগ(["a"], "b");          // → ["a", "b"]
যোগ([], 1);               // → [1]

// Add any type
যোগ([1, 2], "hello");     // → [1, 2, "hello"]
যোগ([[1]], [2]);          // → [[1], [2]]

// Original unchanged
ধরি arr = [1, 2];
ধরি newArr = যোগ(arr, 3);
লেখ(arr);                 // → [1, 2] (unchanged)
লেখ(newArr);              // → [1, 2, 3]

// Chain operations
যোগ(যোগ([1], 2), 3);      // → [1, 2, 3]

// Errors
যোগ([1]);                 // ERROR: wrong number of arguments. got=1, want=2
যোগ("hello", "x");        // ERROR: argument to 'যোগ' must be ARRAY
যোগ([1], 2, 3);           // ERROR: wrong number of arguments. got=3, want=2
```

**Use Case - Building Arrays**:
```bhasa
ধরি buildArray = ফাংশন(n) {
    ধরি helper = ফাংশন(arr, i) {
        যদি (i > n) {
            ফেরত arr;
        }
        helper(যোগ(arr, i), i + 1)
    };
    helper([], 1)
};

buildArray(5);  // → [1, 2, 3, 4, 5]
```

---

### টাইপ (Type)

**Signature**:
```bhasa
টাইপ(value)
```

**Purpose**: Returns the type of a value as a string

**Parameters**:
- `value`: Any value

**Returns**:
- `STRING` - Type name
- `ERROR` - If wrong number of arguments

**Type Names**:
- `"INTEGER"` - Numbers
- `"STRING"` - Strings
- `"BOOLEAN"` - Booleans
- `"ARRAY"` - Arrays
- `"HASH"` - Hash maps
- `"FUNCTION"` - Functions
- `"BUILTIN"` - Built-in functions
- `"NULL"` - Null value
- `"ERROR"` - Error objects

**Implementation**:
```go
"টাইপ": {
    Fn: func(args ...object.Object) object.Object {
        if len(args) != 1 {
            return newError("wrong number of arguments. got=%d, want=1", len(args))
        }
        return &object.String{Value: string(args[0].Type())}
    },
},
```

**Examples**:
```bhasa
// Basic types
টাইপ(5);                  // → "INTEGER"
টাইপ("hello");            // → "STRING"
টাইপ(সত্য);               // → "BOOLEAN"
টাইপ(মিথ্যা);              // → "BOOLEAN"

// Collections
টাইপ([1, 2, 3]);          // → "ARRAY"
টাইপ({"key": "value"});   // → "HASH"

// Functions
টাইপ(ফাংশন(x) { x });    // → "FUNCTION"
টাইপ(লেখ);                // → "BUILTIN"

// Special values
টাইপ(null);               // → "NULL"

// Errors
টাইপ();                   // ERROR: wrong number of arguments
টাইপ(1, 2);               // ERROR: wrong number of arguments
```

**Use Case - Type Checking**:
```bhasa
ধরি process = ফাংশন(val) {
    ধরি t = টাইপ(val);
    যদি (t == "INTEGER") {
        লেখ("Number:", val);
    } নাহলে যদি (t == "STRING") {
        লেখ("Text:", val);
    } নাহলে যদি (t == "ARRAY") {
        লেখ("Array length:", দৈর্ঘ্য(val));
    } নাহলে {
        লেখ("Unknown type:", t);
    }
};

process(42);           // Number: 42
process("hello");      // Text: hello
process([1, 2, 3]);    // Array length: 3
```

---

## Implementation Details

### Builtin Function Structure

```go
type Builtin struct {
    Fn func(args ...object.Object) object.Object
}
```

**Storage**:
```go
var builtins = map[string]*object.Builtin{
    "function_name": {
        Fn: func(args ...object.Object) object.Object {
            // Implementation
        },
    },
}
```

---

### Lookup Process

When an identifier is evaluated:

1. **Check environment** (user-defined variables)
2. **Check builtins** (built-in functions)
3. **Return error** if not found

```go
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
    // 1. Check environment
    if val, ok := env.Get(node.Value); ok {
        return val
    }
    
    // 2. Check builtins
    if builtin, ok := builtins[node.Value]; ok {
        return builtin
    }
    
    // 3. Not found
    return newError("identifier not found: " + node.Value)
}
```

---

### Builtin vs User Function

**Builtin Functions**:
- Implemented in Go
- Cannot be reassigned
- No environment (no closure)
- Direct function call

**User Functions**:
- Implemented in Bhasa
- Stored in environment
- Have closure (captured environment)
- Evaluated by interpreter

---

## Usage Examples

### Example 1: Array Processing

```bhasa
ধরি numbers = [1, 2, 3, 4, 5];

// Get info
লেখ("Length:", দৈর্ঘ্য(numbers));
লেখ("First:", প্রথম(numbers));
লেখ("Last:", শেষ(numbers));

// Process
ধরি rest = বাকি(numbers);
লেখ("Rest:", rest);

ধরি extended = যোগ(numbers, 6);
লেখ("Extended:", extended);

// Output:
// Length: 5
// First: 1
// Last: 5
// Rest: [2, 3, 4, 5]
// Extended: [1, 2, 3, 4, 5, 6]
```

---

### Example 2: Type-Safe Function

```bhasa
ধরি safeAdd = ফাংশন(a, b) {
    যদি (টাইপ(a) != "INTEGER") {
        ফেরত "Error: first arg must be INTEGER";
    }
    যদি (টাইপ(b) != "INTEGER") {
        ফেরত "Error: second arg must be INTEGER";
    }
    a + b
};

লেখ(safeAdd(5, 3));          // 8
লেখ(safeAdd("5", 3));        // Error: first arg must be INTEGER
লেখ(safeAdd(5, "3"));        // Error: second arg must be INTEGER
```

---

### Example 3: Recursive List Sum

```bhasa
ধরি sum = ফাংশন(arr) {
    যদি (টাইপ(arr) != "ARRAY") {
        ফেরত 0;
    }
    যদি (দৈর্ঘ্য(arr) == 0) {
        ফেরত 0;
    }
    প্রথম(arr) + sum(বাকি(arr))
};

লেখ(sum([1, 2, 3, 4, 5]));  // 15
```

---

### Example 4: Map Function

```bhasa
ধরি map = ফাংশন(arr, fn) {
    যদি (দৈর্ঘ্য(arr) == 0) {
        ফেরত [];
    }
    ধরি first = প্রথম(arr);
    ধরি rest = বাকি(arr);
    যোগ(map(rest, fn), fn(first))
};

ধরি double = ফাংশন(x) { x * 2 };
লেখ(map([1, 2, 3], double));  // [2, 4, 6]
```

---

### Example 5: Filter Function

```bhasa
ধরি filter = ফাংশন(arr, predicate) {
    যদি (দৈর্ঘ্য(arr) == 0) {
        ফেরত [];
    }
    
    ধরি first = প্রথম(arr);
    ধরি rest = বাকি(arr);
    ধরি filtered = filter(rest, predicate);
    
    যদি (predicate(first)) {
        যোগ(filtered, first)
    } নাহলে {
        filtered
    }
};

ধরি isEven = ফাংশন(x) { x % 2 == 0 };
লেখ(filter([1, 2, 3, 4, 5, 6], isEven));  // [2, 4, 6]
```

---

## Error Handling

### Common Errors

**Wrong Argument Count**:
```bhasa
দৈর্ঘ্য();          // ERROR: wrong number of arguments. got=0, want=1
দৈর্ঘ্য([1], [2]);   // ERROR: wrong number of arguments. got=2, want=1
```

**Wrong Argument Type**:
```bhasa
দৈর্ঘ্য(5);         // ERROR: argument to 'দৈর্ঘ্য' not supported, got INTEGER
প্রথম("hello");     // ERROR: argument to 'প্রথম' must be ARRAY, got STRING
```

**Type-Specific Errors**:
```bhasa
// Hash key not hashable
{"ফাংশন(x){x}": "value"}  // ERROR: unusable as hash key
```

---

## Adding New Builtins

### Step 1: Add to builtins map

```go
var builtins = map[string]*object.Builtin{
    // ... existing builtins
    
    "নতুন_ফাংশন": {  // New function
        Fn: func(args ...object.Object) object.Object {
            // 1. Validate argument count
            if len(args) != 1 {
                return newError("wrong number of arguments. got=%d, want=1", len(args))
            }
            
            // 2. Validate argument types
            if args[0].Type() != object.INTEGER_OBJ {
                return newError("argument must be INTEGER, got %s", args[0].Type())
            }
            
            // 3. Extract values
            val := args[0].(*object.Integer).Value
            
            // 4. Implement logic
            result := val * 2
            
            // 5. Return object
            return &object.Integer{Value: result}
        },
    },
}
```

### Step 2: Add tests

```go
func TestNewBuiltin(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {`নতুন_ফাংশন(5)`, 10},
        {`নতুন_ফাংশন(0)`, 0},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}
```

### Step 3: Update documentation

Add entry to this file and update README.

---

## Best Practices

### 1. Argument Validation

Always validate:
- Argument count
- Argument types
- Argument values (if applicable)

### 2. Error Messages

Make error messages clear and specific:
```go
return newError("wrong number of arguments. got=%d, want=1", len(args))
return newError("argument to 'দৈর্ঘ্য' not supported, got %s", args[0].Type())
```

### 3. Return Consistent Types

- Return same type for success cases
- Return ERROR object for failures
- Document return types

### 4. Don't Modify Arguments

Create new objects instead of modifying:
```go
// Good
newElements := make([]object.Object, length+1)
copy(newElements, arr.Elements)
newElements[length] = newElement
return &object.Array{Elements: newElements}

// Bad (modifies original)
arr.Elements = append(arr.Elements, newElement)
return arr
```

---

## See Also

- [Evaluator Documentation](./evaluator-documentation.md)
- [Evaluation Examples](./evaluation-examples.md)
- [Quick Reference](./quick-reference.md)
- [Object System](../../object/docs/)

