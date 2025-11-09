# Object Package Documentation

## Overview

The **Object** package is the runtime value system of the Bhasa language. It defines all value types, the environment (variable scope), and built-in functions. Every value in Bhasa—from integers to functions to class instances—is represented as an `Object`.

## Table of Contents

- [What is an Object?](#what-is-an-object)
- [Object System Architecture](#object-system-architecture)
- [Value Types](#value-types)
- [Environment System](#environment-system)
- [Built-in Functions](#built-in-functions)
- [Hash System](#hash-system)
- [OOP System](#oop-system)
- [Quick Reference](#quick-reference)

## What is an Object?

An **Object** in Bhasa is any runtime value. It's an interface that all values implement:

```go
type Object interface {
    Type() ObjectType
    Inspect() string
}
```

### Why an Interface?

**Benefits:**
1. **Polymorphism**: Functions can accept any object type
2. **Type Safety**: Runtime type checking with `Type()`
3. **Debugging**: `Inspect()` provides string representation
4. **Extensibility**: Easy to add new types

### Object Types

```go
type ObjectType string

const (
    INTEGER_OBJ    = "INTEGER"
    BOOLEAN_OBJ    = "BOOLEAN"
    STRING_OBJ     = "STRING"
    NULL_OBJ       = "NULL"
    FUNCTION_OBJ   = "FUNCTION"
    ARRAY_OBJ      = "ARRAY"
    HASH_OBJ       = "HASH"
    // ... and many more
)
```

## Object System Architecture

### Type Hierarchy

```
Object (interface)
├─ Value Types
│  ├─ Numeric
│  │  ├─ Integer (legacy int64)
│  │  ├─ Byte (8-bit)
│  │  ├─ Short (16-bit)
│  │  ├─ Int (32-bit)
│  │  ├─ Long (64-bit)
│  │  ├─ Float (32-bit)
│  │  ├─ Double (64-bit)
│  │  └─ Char (rune)
│  ├─ Boolean
│  ├─ String
│  └─ Null
├─ Composite Types
│  ├─ Array
│  ├─ Hash
│  ├─ Struct
│  └─ Enum
├─ Functions
│  ├─ Function (interpreted)
│  ├─ Builtin (native Go)
│  ├─ CompiledFunction (bytecode)
│  └─ Closure
├─ Control Flow
│  ├─ ReturnValue
│  └─ Error
└─ OOP Types
   ├─ Class
   ├─ ClassInstance
   ├─ Method
   ├─ BoundMethod
   └─ Interface
```

### Design Principles

1. **Everything is an Object**: Even functions and errors
2. **Immutability**: Most objects are immutable
3. **Type Safety**: Runtime type checking
4. **Garbage Collected**: Go's GC manages memory

## Value Types

### Numeric Types

Bhasa provides a rich numeric type system:

#### Integer (Legacy)

```go
type Integer struct {
    Value int64
}
```

- **Range**: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
- **Use**: General-purpose integer
- **Example**: `&Integer{Value: 42}`

#### Byte

```go
type Byte struct {
    Value int8
}
```

- **Range**: 0 to 255
- **Use**: Small integers, byte operations
- **Example**: `&Byte{Value: 255}`

#### Short

```go
type Short struct {
    Value int16
}
```

- **Range**: -32,768 to 32,767
- **Use**: Medium integers
- **Example**: `&Short{Value: 1000}`

#### Int

```go
type Int struct {
    Value int32
}
```

- **Range**: -2,147,483,648 to 2,147,483,647
- **Use**: Standard 32-bit integer
- **Example**: `&Int{Value: 1000000}`

#### Long

```go
type Long struct {
    Value int64
}
```

- **Range**: Same as Integer (int64)
- **Use**: Large integers
- **Example**: `&Long{Value: 9999999999}`

#### Float

```go
type Float struct {
    Value float32
}
```

- **Precision**: ~7 decimal digits
- **Use**: Single-precision floating point
- **Example**: `&Float{Value: 3.14159}`

#### Double

```go
type Double struct {
    Value float64
}
```

- **Precision**: ~15 decimal digits
- **Use**: Double-precision floating point
- **Example**: `&Double{Value: 3.141592653589793}`

#### Char

```go
type Char struct {
    Value rune
}
```

- **Range**: Any Unicode code point
- **Use**: Single character
- **Example**: `&Char{Value: 'অ'}` (Bengali letter)

### Boolean

```go
type Boolean struct {
    Value bool
}
```

**Values:**
- `&Boolean{Value: true}`  → সত্য
- `&Boolean{Value: false}` → মিথ্যা

**Singleton Pattern** (recommended):
```go
var (
    TRUE  = &Boolean{Value: true}
    FALSE = &Boolean{Value: false}
)
```

### String

```go
type String struct {
    Value string
}
```

- **Encoding**: UTF-8
- **Immutable**: Once created, never changes
- **Example**: `&String{Value: "হ্যালো বিশ্ব"}`

### Null

```go
type Null struct{}
```

- **Singleton**: Only one null value exists
- **Represents**: Absence of value
- **Example**: `&Null{}`

**Singleton Pattern:**
```go
var NULL = &Null{}
```

## Composite Types

### Array

```go
type Array struct {
    Elements []Object
}
```

**Characteristics:**
- Heterogeneous: Can hold different types
- Dynamic size
- Zero-indexed

**Example:**
```go
arr := &Array{
    Elements: []Object{
        &Integer{Value: 1},
        &String{Value: "hello"},
        &Boolean{Value: true},
    },
}
// Inspect: [1, hello, true]
```

### Hash

```go
type Hash struct {
    Pairs map[HashKey]HashPair
}

type HashPair struct {
    Key   Object
    Value Object
}

type HashKey struct {
    Type  ObjectType
    Value uint64
}
```

**Key Types:**
- Integer, Byte, Short, Int, Long
- Float, Double
- String
- Boolean
- Char

**Example:**
```go
hash := &Hash{
    Pairs: map[HashKey]HashPair{
        (&String{Value: "name"}).HashKey(): {
            Key:   &String{Value: "name"},
            Value: &String{Value: "রহিম"},
        },
        (&String{Value: "age"}).HashKey(): {
            Key:   &String{Value: "age"},
            Value: &Integer{Value: 25},
        },
    },
}
// Inspect: {name: রহিম, age: 25}
```

### Struct

```go
type Struct struct {
    Fields     map[string]Object
    FieldOrder []string
}
```

**Features:**
- Named fields
- Maintains insertion order
- Heterogeneous values

**Example:**
```go
person := &Struct{
    Fields: map[string]Object{
        "নাম":  &String{Value: "করিম"},
        "বয়স": &Integer{Value: 30},
    },
    FieldOrder: []string{"নাম", "বয়স"},
}
// Inspect: {নাম: করিম, বয়স: 30}
```

### Enum

```go
type EnumType struct {
    Name     string
    Variants map[string]int
}

type Enum struct {
    EnumType    string
    VariantName string
    Value       int
}
```

**Example:**
```go
// Type definition
colorType := &EnumType{
    Name: "Color",
    Variants: map[string]int{
        "লাল":  0,
        "সবুজ": 1,
        "নীল":  2,
    },
}

// Value
red := &Enum{
    EnumType:    "Color",
    VariantName: "লাল",
    Value:       0,
}
// Inspect: Color.লাল
```

## Function Types

### Function

```go
type Function struct {
    Parameters []*ast.Identifier
    Body       *ast.BlockStatement
    Env        *Environment
}
```

**Interpreted function** with:
- Parameters: List of parameter names
- Body: AST of function body
- Env: Captured environment (closure)

**Example:**
```bengali
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
```

### Builtin

```go
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
    Fn BuiltinFunction
}
```

**Native Go function** callable from Bhasa

**Example:**
```go
&Builtin{
    Fn: func(args ...Object) Object {
        // Implementation in Go
        return &Integer{Value: 42}
    },
}
```

### CompiledFunction

```go
type CompiledFunction struct {
    Instructions  []byte
    NumLocals     int
    NumParameters int
}
```

**Bytecode function** for VM execution

### Closure

```go
type Closure struct {
    Fn   *CompiledFunction
    Free []Object
}
```

**Compiled function with captured free variables**

## Control Flow Objects

### ReturnValue

```go
type ReturnValue struct {
    Value Object
}
```

**Wraps a return value** to propagate through nested evaluations

**Example:**
```go
// ফেরত ৫;
&ReturnValue{Value: &Integer{Value: 5}}
```

### Error

```go
type Error struct {
    Message string
}
```

**Runtime error** with message

**Example:**
```go
&Error{Message: "division by zero"}
```

## Environment System

### Environment Structure

```go
type Environment struct {
    store map[string]Object
    outer *Environment
}
```

**Variable scope manager** with:
- `store`: Current scope's variables
- `outer`: Parent scope (if any)

### Scope Chain

```
Global Environment
  ↓ outer
Function Environment
  ↓ outer
Block Environment
```

**Example:**
```bengali
ধরি x = ১০;              // Global scope

ফাংশন outer() {
    ধরি y = ২০;          // outer's scope
    
    ফাংশন inner() {
        ধরি z = ৩০;      // inner's scope
        ফেরত x + y + z;  // Can access x, y, z
    }
    
    ফেরত inner();
}
```

### Functions

#### NewEnvironment

```go
func NewEnvironment() *Environment
```

Creates a **global** environment with no parent.

```go
env := object.NewEnvironment()
env.Set("x", &object.Integer{Value: 10})
```

#### NewEnclosedEnvironment

```go
func NewEnclosedEnvironment(outer *Environment) *Environment
```

Creates a **nested** environment with parent.

```go
globalEnv := object.NewEnvironment()
funcEnv := object.NewEnclosedEnvironment(globalEnv)
```

#### Get

```go
func (e *Environment) Get(name string) (Object, bool)
```

Retrieves variable, searching up scope chain.

```go
if val, ok := env.Get("x"); ok {
    // Variable found
} else {
    // Variable not found
}
```

#### Set

```go
func (e *Environment) Set(name string, val Object) Object
```

Sets variable in **current** scope.

```go
env.Set("x", &object.Integer{Value: 42})
```

## Built-in Functions

Bhasa provides **40+ built-in functions** with Bengali names.

### Categories

1. **Basic I/O**: লেখ (print)
2. **Array Operations**: দৈর্ঘ্য, প্রথম, শেষ, বাকি, যোগ
3. **Type Introspection**: টাইপ
4. **String Operations**: বিভক্ত, যুক্ত, উপরে, নিচে, ছাঁটো
5. **Math Operations**: শক্তি, বর্গমূল, পরম, সর্বোচ্চ, সর্বনিম্ন
6. **File I/O**: ফাইল_পড়ো, ফাইল_লেখো, ফাইল_যোগ, ফাইল_আছে
7. **JSON**: JSON_পার্স, JSON_স্ট্রিং
8. **Hash Operations**: চাবিগুলো, মানগুলো, চাবি_আছে, একত্রিত
9. **Type Conversion**: বাইট, ছোট_সংখ্যা, পূর্ণসংখ্যা, দীর্ঘ_সংখ্যা, দশমিক

### Examples

#### লেখ (Print)

```bengali
লেখ("হ্যালো");          // Output: হ্যালো
লেখ(১০ + ২০);          // Output: 30
```

#### দৈর্ঘ্য (Length)

```bengali
ধরি arr = [১, ২, ৩];
লেখ(দৈর্ঘ্য(arr));     // Output: 3

ধরি str = "হ্যালো";
লেখ(দৈর্ঘ্য(str));     // Output: 5
```

#### যোগ (Push)

```bengali
ধরি arr = [১, ২];
ধরি newArr = যোগ(arr, ৩);
লেখ(newArr);          // Output: [1, 2, 3]
```

#### বিভক্ত (Split)

```bengali
ধরি text = "এক,দুই,তিন";
ধরি parts = বিভক্ত(text, ",");
লেখ(parts);           // Output: [এক, দুই, তিন]
```

#### শক্তি (Power)

```bengali
লেখ(শক্তি(২, ৩));     // Output: 8 (2^3)
```

## Hash System

### Hashable Interface

```go
type Hashable interface {
    HashKey() HashKey
}
```

Types that can be hash keys must implement `Hashable`.

### HashKey Generation

Different types generate hash keys differently:

#### Integer

```go
func (i *Integer) HashKey() HashKey {
    return HashKey{
        Type:  INTEGER_OBJ,
        Value: uint64(i.Value),
    }
}
```

#### String

```go
func (s *String) HashKey() HashKey {
    h := fnv.New64a()
    h.Write([]byte(s.Value))
    return HashKey{
        Type:  STRING_OBJ,
        Value: h.Sum64(),
    }
}
```

#### Boolean

```go
func (b *Boolean) HashKey() HashKey {
    var value uint64
    if b.Value {
        value = 1
    } else {
        value = 0
    }
    return HashKey{
        Type:  BOOLEAN_OBJ,
        Value: value,
    }
}
```

### Why HashKey?

**Problem**: Go maps require comparable keys
**Solution**: Convert objects to HashKey struct

```go
// Can't do this:
map[Object]Object

// Do this instead:
map[HashKey]HashPair
```

## OOP System

### Class

```go
type Class struct {
    Name         string
    SuperClass   *Class
    Interfaces   []*Interface
    Fields       map[string]string
    Methods      map[string]*Method
    Constructor  *Closure
    StaticFields map[string]Object
    IsAbstract   bool
    IsFinal      bool
    FieldAccess  map[string]string
    FieldOrder   []string
}
```

**Features:**
- Inheritance (single)
- Interfaces (multiple)
- Access modifiers
- Static members
- Abstract/Final modifiers

**Example:**
```bengali
শ্রেণী Person {
    সার্বজনীন নাম: পাঠ্য;
    ব্যক্তিগত বয়স: পূর্ণসংখ্যা;
    
    নির্মাতা(নাম: পাঠ্য, বয়স: পূর্ণসংখ্যা) {
        এই.নাম = নাম;
        এই.বয়স = বয়স;
    }
    
    সার্বজনীন পদ্ধতি greet(): পাঠ্য {
        ফেরত "হ্যালো, আমি " + এই.নাম;
    }
}
```

### ClassInstance

```go
type ClassInstance struct {
    Class  *Class
    Fields map[string]Object
    This   Object
}
```

**Instance of a class** with field values.

**Example:**
```bengali
ধরি p = নতুন Person("রহিম", ২৫);
লেখ(p.নাম);          // Output: রহিম
লেখ(p.greet());     // Output: হ্যালো, আমি রহিম
```

### Method

```go
type Method struct {
    Name       string
    Access     string
    IsStatic   bool
    IsFinal    bool
    IsAbstract bool
    Closure    *Closure
}
```

**Method definition** with modifiers.

### BoundMethod

```go
type BoundMethod struct {
    Receiver Object
    Method   *Closure
}
```

**Method bound to instance** with `this` reference.

### Interface

```go
type Interface struct {
    Name             string
    MethodSignatures map[string][]string
}
```

**Interface definition** (চুক্তি in Bengali).

## Quick Reference

### Creating Objects

```go
// Integer
i := &object.Integer{Value: 42}

// String
s := &object.String{Value: "হ্যালো"}

// Boolean (use singletons)
t := &object.Boolean{Value: true}

// Array
arr := &object.Array{
    Elements: []object.Object{i, s, t},
}

// Hash
h := &object.Hash{
    Pairs: make(map[object.HashKey]object.HashPair),
}
```

### Type Checking

```go
switch obj := value.(type) {
case *object.Integer:
    fmt.Println("Integer:", obj.Value)
case *object.String:
    fmt.Println("String:", obj.Value)
case *object.Boolean:
    fmt.Println("Boolean:", obj.Value)
default:
    fmt.Println("Unknown type")
}
```

### Environment Usage

```go
// Create global environment
env := object.NewEnvironment()

// Set variable
env.Set("x", &object.Integer{Value: 10})

// Get variable
if val, ok := env.Get("x"); ok {
    fmt.Println("Found:", val.Inspect())
}

// Create nested environment
funcEnv := object.NewEnclosedEnvironment(env)
funcEnv.Set("y", &object.Integer{Value: 20})
// funcEnv can access both x and y
// env can only access x
```

### Using Builtins

```go
// Get builtin by name
builtin := object.GetBuiltinByName("লেখ")

// Call builtin
result := builtin.Fn(
    &object.String{Value: "হ্যালো"},
)
```

## Summary

The **Object** package provides:

✅ **20+ Value Types**: Complete type system  
✅ **Environment**: Scoped variable management  
✅ **40+ Builtins**: Rich standard library  
✅ **Hash System**: Efficient key-value storage  
✅ **OOP Support**: Classes, inheritance, interfaces  
✅ **Type Safety**: Runtime type checking  
✅ **Closures**: First-class functions with capture  

For detailed implementation information, see [object-documentation.md](./object-documentation.md) and [builtins-documentation.md](./builtins-documentation.md).

---

The object package is the heart of Bhasa's runtime system, providing everything needed to represent and manipulate values during program execution.

