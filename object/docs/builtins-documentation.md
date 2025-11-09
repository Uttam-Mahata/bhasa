# Built-in Functions Documentation

Complete reference for all 40+ built-in functions in Bhasa.

## Table of Contents

- [Overview](#overview)
- [Basic I/O](#basic-io)
- [Array Operations](#array-operations)
- [String Operations](#string-operations)
- [Math Operations](#math-operations)
- [Array Advanced](#array-advanced)
- [File I/O](#file-io)
- [JSON Operations](#json-operations)
- [Hash Operations](#hash-operations)
- [Character Operations](#character-operations)
- [Type Conversion](#type-conversion)
- [Type System](#type-system)

---

## Overview

Built-in functions are **native Go functions** callable from Bhasa code. They provide essential functionality without requiring external libraries.

### Built-in Structure

```go
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
    Fn BuiltinFunction
}
```

### Registration

```go
var Builtins = []BuiltinDef{
    {"লেখ", &Builtin{Fn: printFunc}},
    {"দৈর্ঘ্য", &Builtin{Fn: lenFunc}},
    // ... more builtins
}
```

---

## Basic I/O

### লেখ (Print)

**Signature:** `লেখ(args...)`

**Purpose:** Print values to console

**Parameters:**
- `args...`: Any number of objects to print

**Returns:** `NULL`

**Implementation:**
```go
func(args ...Object) Object {
    for _, arg := range args {
        fmt.Println(arg.Inspect())
    }
    return &Null{}
}
```

**Examples:**
```bengali
লেখ("হ্যালো");              // Output: হ্যালো
লেখ(১০);                   // Output: 10
লেখ("বয়স:", ২৫);          // Output: বয়স:
                           //         25
```

**Behavior:**
- Each argument printed on separate line
- Calls `Inspect()` on each object
- Always returns `NULL`

---

## Array Operations

### দৈর্ঘ্য (Length)

**Signature:** `দৈর্ঘ্য(value)`

**Purpose:** Get length of string or array

**Parameters:**
- `value`: String or Array

**Returns:** Integer (length)

**Examples:**
```bengali
দৈর্ঘ্য([১, ২, ৩])          // Returns: 3
দৈর্ঘ্য("হ্যালো")          // Returns: 5 (counts runes, not bytes)
দৈর্ঘ্য("")               // Returns: 0
```

**Implementation:**
```go
func(args ...Object) Object {
    if len(args) != 1 {
        return &Error{Message: "wrong number of arguments"}
    }
    switch arg := args[0].(type) {
    case *String:
        return &Integer{Value: int64(len([]rune(arg.Value)))}
    case *Array:
        return &Integer{Value: int64(len(arg.Elements))}
    default:
        return &Error{Message: "argument not supported"}
    }
}
```

**Note:** For strings, counts Unicode characters (runes), not bytes.

### প্রথম (First)

**Signature:** `প্রথম(array)`

**Purpose:** Get first element of array

**Parameters:**
- `array`: Array object

**Returns:** First element or `NULL` if empty

**Examples:**
```bengali
প্রথম([১, ২, ৩])           // Returns: 1
প্রথম([])                 // Returns: null
```

### শেষ (Last)

**Signature:** `শেষ(array)`

**Purpose:** Get last element of array

**Parameters:**
- `array`: Array object

**Returns:** Last element or `NULL` if empty

**Examples:**
```bengali
শেষ([১, ২, ৩])            // Returns: 3
শেষ([])                  // Returns: null
```

### বাকি (Rest)

**Signature:** `বাকি(array)`

**Purpose:** Get all elements except first

**Parameters:**
- `array`: Array object

**Returns:** New array without first element, or `NULL` if empty

**Examples:**
```bengali
বাকি([১, ২, ৩, ৪])        // Returns: [2, 3, 4]
বাকি([১])                // Returns: []
বাকি([])                 // Returns: null
```

**Implementation:**
```go
func(args ...Object) Object {
    arr := args[0].(*Array)
    length := len(arr.Elements)
    if length > 0 {
        newElements := make([]Object, length-1)
        copy(newElements, arr.Elements[1:length])
        return &Array{Elements: newElements}
    }
    return &Null{}
}
```

### যোগ (Push)

**Signature:** `যোগ(array, element)`

**Purpose:** Add element to end of array

**Parameters:**
- `array`: Array object
- `element`: Element to add

**Returns:** New array with element added

**Examples:**
```bengali
যোগ([১, ২], ৩)            // Returns: [1, 2, 3]
যোগ([], "first")         // Returns: ["first"]
```

**Note:** Does not modify original array (functional approach).

---

## String Operations

### বিভক্ত (Split)

**Signature:** `বিভক্ত(string, delimiter)`

**Purpose:** Split string by delimiter

**Parameters:**
- `string`: String to split
- `delimiter`: Separator string

**Returns:** Array of strings

**Examples:**
```bengali
বিভক্ত("এক,দুই,তিন", ",")      // Returns: ["এক", "দুই", "তিন"]
বিভক্ত("hello world", " ")   // Returns: ["hello", "world"]
বিভক্ত("abc", "")            // Returns: ["a", "b", "c"]
```

### যুক্ত (Join)

**Signature:** `যুক্ত(array, delimiter)`

**Purpose:** Join array elements with delimiter

**Parameters:**
- `array`: Array of elements
- `delimiter`: Separator string

**Returns:** Joined string

**Examples:**
```bengali
যুক্ত(["এক", "দুই", "তিন"], ",")   // Returns: "এক,দুই,তিন"
যুক্ত([১, ২, ৩], "-")             // Returns: "1-2-3"
```

**Note:** Calls `Inspect()` on each element.

### উপরে (Uppercase)

**Signature:** `উপরে(string)`

**Purpose:** Convert string to uppercase

**Parameters:**
- `string`: String to convert

**Returns:** Uppercase string

**Examples:**
```bengali
উপরে("hello")              // Returns: "HELLO"
উপরে("হ্যালো")             // Returns: "হ্যালো" (Bengali unchanged)
```

### নিচে (Lowercase)

**Signature:** `নিচে(string)`

**Purpose:** Convert string to lowercase

**Parameters:**
- `string`: String to convert

**Returns:** Lowercase string

**Examples:**
```bengali
নিচে("HELLO")              // Returns: "hello"
```

### ছাঁটো (Trim)

**Signature:** `ছাঁটো(string)`

**Purpose:** Remove leading/trailing whitespace

**Parameters:**
- `string`: String to trim

**Returns:** Trimmed string

**Examples:**
```bengali
ছাঁটো("  hello  ")         // Returns: "hello"
ছাঁটো("\n\tworld\t\n")     // Returns: "world"
```

### প্রতিস্থাপন (Replace)

**Signature:** `প্রতিস্থাপন(string, old, new)`

**Purpose:** Replace all occurrences

**Parameters:**
- `string`: Original string
- `old`: String to replace
- `new`: Replacement string

**Returns:** Modified string

**Examples:**
```bengali
প্রতিস্থাপন("hello world", "world", "বিশ্ব")
// Returns: "hello বিশ্ব"

প্রতিস্থাপন("aaa", "a", "b")
// Returns: "bbb"
```

### খুঁজুন (Find/IndexOf)

**Signature:** `খুঁজুন(string, substring)`

**Purpose:** Find index of substring

**Parameters:**
- `string`: String to search in
- `substring`: String to find

**Returns:** Integer index (or -1 if not found)

**Examples:**
```bengali
খুঁজুন("hello world", "world")    // Returns: 6
খুঁজুন("hello", "x")              // Returns: -1
```

---

## Math Operations

### শক্তি (Power)

**Signature:** `শক্তি(base, exponent)`

**Purpose:** Calculate power (base^exponent)

**Parameters:**
- `base`: Integer base
- `exponent`: Integer exponent

**Returns:** Integer result

**Examples:**
```bengali
শক্তি(২, ৩)                // Returns: 8
শক্তি(৫, ২)                // Returns: 25
শক্তি(১০, ০)               // Returns: 1
```

### বর্গমূল (Square Root)

**Signature:** `বর্গমূল(number)`

**Purpose:** Calculate square root

**Parameters:**
- `number`: Integer (must be non-negative)

**Returns:** Integer (truncated)

**Examples:**
```bengali
বর্গমূল(৯)                 // Returns: 3
বর্গমূল(১৬)                // Returns: 4
বর্গমূল(২)                 // Returns: 1 (truncated)
```

**Error:** Negative numbers return error.

### পরম (Absolute Value)

**Signature:** `পরম(number)`

**Purpose:** Get absolute value

**Parameters:**
- `number`: Integer

**Returns:** Non-negative integer

**Examples:**
```bengali
পরম(-৫)                    // Returns: 5
পরম(৫)                     // Returns: 5
পরম(০)                     // Returns: 0
```

### সর্বোচ্চ (Max)

**Signature:** `সর্বোচ্চ(a, b)`

**Purpose:** Get maximum of two numbers

**Parameters:**
- `a`: First integer
- `b`: Second integer

**Returns:** Larger integer

**Examples:**
```bengali
সর্বোচ্চ(৫, ১০)             // Returns: 10
সর্বোচ্চ(-৫, -১০)           // Returns: -5
```

### সর্বনিম্ন (Min)

**Signature:** `সর্বনিম্ন(a, b)`

**Purpose:** Get minimum of two numbers

**Parameters:**
- `a`: First integer
- `b`: Second integer

**Returns:** Smaller integer

**Examples:**
```bengali
সর্বনিম্ন(৫, ১০)            // Returns: 5
সর্বনিম্ন(-৫, -১০)          // Returns: -10
```

---

## Array Advanced

### উল্টাও (Reverse)

**Signature:** `উল্টাও(array)`

**Purpose:** Reverse array

**Parameters:**
- `array`: Array to reverse

**Returns:** New reversed array

**Examples:**
```bengali
উল্টাও([১, ২, ৩])           // Returns: [3, 2, 1]
উল্টাও(["a", "b"])        // Returns: ["b", "a"]
```

### সাজাও (Sort)

**Signature:** `সাজাও(array)`

**Purpose:** Sort array (integers only)

**Parameters:**
- `array`: Array of integers

**Returns:** New sorted array (ascending)

**Examples:**
```bengali
সাজাও([৩, ১, ২])            // Returns: [1, 2, 3]
সাজাও([৫, -২, ০])          // Returns: [-2, 0, 5]
```

**Implementation:** Bubble sort

**Note:** Only works with integer arrays.

---

## File I/O

### ফাইল_পড়ো (Read File)

**Signature:** `ফাইল_পড়ো(filename)`

**Purpose:** Read entire file

**Parameters:**
- `filename`: String path to file

**Returns:** String content or Error

**Examples:**
```bengali
ধরি content = ফাইল_পড়ো("data.txt");
লেখ(content);
```

**Errors:**
- File not found
- Permission denied
- I/O errors

### ফাইল_লেখো (Write File)

**Signature:** `ফাইল_লেখো(filename, content)`

**Purpose:** Write/overwrite file

**Parameters:**
- `filename`: String path to file
- `content`: String to write

**Returns:** NULL or Error

**Examples:**
```bengali
ফাইল_লেখো("output.txt", "হ্যালো বিশ্ব");
```

**Behavior:**
- Creates file if doesn't exist
- Overwrites if exists
- Permissions: 0644

### ফাইল_যোগ (Append to File)

**Signature:** `ফাইল_যোগ(filename, content)`

**Purpose:** Append to file

**Parameters:**
- `filename`: String path to file
- `content`: String to append

**Returns:** NULL or Error

**Examples:**
```bengali
ফাইল_যোগ("log.txt", "নতুন লাইন\n");
```

**Behavior:**
- Creates file if doesn't exist
- Appends to end if exists

### ফাইল_আছে (File Exists)

**Signature:** `ফাইল_আছে(filename)`

**Purpose:** Check if file exists

**Parameters:**
- `filename`: String path to file

**Returns:** Boolean

**Examples:**
```bengali
যদি (ফাইল_আছে("config.json")) {
    ধরি config = ফাইল_পড়ো("config.json");
}
```

---

## JSON Operations

### JSON_পার্স (Parse JSON)

**Signature:** `JSON_পার্স(jsonString)`

**Purpose:** Parse JSON string to Bhasa objects

**Parameters:**
- `jsonString`: Valid JSON string

**Returns:** Bhasa object (Hash, Array, String, Integer, Boolean, or Null)

**Examples:**
```bengali
ধরি data = JSON_পার্স('{"name": "রহিম", "age": 25}');
লেখ(data["name"]);        // Output: রহিম
```

**Mapping:**
```
JSON → Bhasa
null → NULL
true/false → Boolean
number → Integer
string → String
array → Array
object → Hash
```

### JSON_স্ট্রিং (Stringify)

**Signature:** `JSON_স্ট্রিং(object)`

**Purpose:** Convert Bhasa object to JSON string

**Parameters:**
- `object`: Bhasa object

**Returns:** JSON string

**Examples:**
```bengali
ধরি obj = {"নাম": "করিম", "বয়স": ৩০};
ধরি json = JSON_স্ট্রিং(obj);
লেখ(json);                // Output: {"নাম":"করিম","বয়স":30}
```

**Supported Types:**
- NULL → `null`
- Boolean → `true`/`false`
- Integer → number
- String → string
- Array → array
- Hash → object

---

## Hash Operations

### চাবিগুলো (Keys)

**Signature:** `চাবিগুলো(hash)`

**Purpose:** Get array of hash keys

**Parameters:**
- `hash`: Hash object

**Returns:** Array of keys

**Examples:**
```bengali
ধরি h = {"a": ১, "b": ২};
লেখ(চাবিগুলো(h));         // Output: ["a", "b"] (order may vary)
```

### মানগুলো (Values)

**Signature:** `মানগুলো(hash)`

**Purpose:** Get array of hash values

**Parameters:**
- `hash`: Hash object

**Returns:** Array of values

**Examples:**
```bengali
ধরি h = {"a": ১, "b": ২};
লেখ(মানগুলো(h));          // Output: [1, 2] (order may vary)
```

### চাবি_আছে (Has Key)

**Signature:** `চাবি_আছে(hash, key)`

**Purpose:** Check if hash has key

**Parameters:**
- `hash`: Hash object
- `key`: Key to check (must be hashable)

**Returns:** Boolean

**Examples:**
```bengali
ধরি h = {"name": "রহিম"};
লেখ(চাবি_আছে(h, "name"));    // Output: true
লেখ(চাবি_আছে(h, "age"));     // Output: false
```

### একত্রিত (Merge)

**Signature:** `একত্রিত(hash1, hash2)`

**Purpose:** Merge two hashes

**Parameters:**
- `hash1`: First hash
- `hash2`: Second hash

**Returns:** New merged hash

**Examples:**
```bengali
ধরি h1 = {"a": ১, "b": ২};
ধরি h2 = {"b": ৩, "c": ৪};
ধরি merged = একত্রিত(h1, h2);
// merged = {"a": 1, "b": 3, "c": 4}
// hash2 values overwrite hash1
```

---

## Character Operations

### অক্ষর (CharAt)

**Signature:** `অক্ষর(string, index)`

**Purpose:** Get character at index

**Parameters:**
- `string`: String
- `index`: Integer index (0-based)

**Returns:** Single-character string

**Examples:**
```bengali
অক্ষর("হ্যালো", ০)          // Returns: "হ"
অক্ষর("hello", ১)         // Returns: "e"
```

**Error:** Index out of bounds

### কোড (CharCode)

**Signature:** `কোড(string)`

**Purpose:** Get Unicode code point of first character

**Parameters:**
- `string`: Non-empty string

**Returns:** Integer code point

**Examples:**
```bengali
কোড("A")                   // Returns: 65
কোড("অ")                   // Returns: 2437 (U+0985)
```

**Error:** Empty string

### অক্ষর_থেকে_কোড (FromCharCode)

**Signature:** `অক্ষর_থেকে_কোড(code)`

**Purpose:** Create string from code point

**Parameters:**
- `code`: Integer Unicode code point (0 to 0x10FFFF)

**Returns:** Single-character string

**Examples:**
```bengali
অক্ষর_থেকে_কোড(৬৫)        // Returns: "A"
অক্ষর_থেকে_কোড(২৪৩৭)      // Returns: "অ"
```

**Error:** Invalid code point

---

## Type Conversion

### সংখ্যা (ParseInt)

**Signature:** `সংখ্যা(string)`

**Purpose:** Parse string to integer

**Parameters:**
- `string`: Numeric string (Arabic or Bengali digits)

**Returns:** Integer or Error

**Examples:**
```bengali
সংখ্যা("১২৩")               // Returns: 123
সংখ্যা("123")              // Returns: 123
সংখ্যা("  42  ")           // Returns: 42 (whitespace trimmed)
```

**Error:** Non-numeric string

### লেখা (ToString)

**Signature:** `লেখা(integer)`

**Purpose:** Convert integer to string

**Parameters:**
- `integer`: Integer value

**Returns:** String representation

**Examples:**
```bengali
লেখা(১২৩)                  // Returns: "123"
লেখা(-৪৫)                  // Returns: "-45"
```

### বাইট (ToByte)

**Signature:** `বাইট(value)`

**Purpose:** Convert to byte (0-255)

**Parameters:**
- `value`: Integer, String, or numeric type

**Returns:** Byte or Error

**Examples:**
```bengali
বাইট(২৫৫)                  // Returns: Byte(255)
বাইট("১২৮")                // Returns: Byte(128)
```

**Error:** Out of range (0-255)

### পূর্ণসংখ্যা (ToInt)

**Signature:** `পূর্ণসংখ্যা(value)`

**Purpose:** Convert to 32-bit int

**Parameters:**
- `value`: Numeric value or string

**Returns:** Int or Error

**Examples:**
```bengali
পূর্ণসংখ্যা(১০০০০)         // Returns: Int(10000)
```

**Error:** Out of range (-2147483648 to 2147483647)

### দশমিক (ToFloat)

**Signature:** `দশমিক(value)`

**Purpose:** Convert to float

**Parameters:**
- `value`: Numeric value or string

**Returns:** Float

**Examples:**
```bengali
দশমিক(১০)                  // Returns: Float(10.0)
দশমিক("৩.১৪")              // Returns: Float(3.14)
```

---

## Type System

### টাইপ (Type)

**Signature:** `টাইপ(value)`

**Purpose:** Get type name of value

**Parameters:**
- `value`: Any object

**Returns:** String type name

**Examples:**
```bengali
টাইপ(৫)                    // Returns: "INTEGER"
টাইপ("hello")             // Returns: "STRING"
টাইপ([১, ২])              // Returns: "ARRAY"
টাইপ(সত্য)                 // Returns: "BOOLEAN"
```

**Type Names:**
```
INTEGER, STRING, BOOLEAN, NULL,
ARRAY, HASH, FUNCTION, BUILTIN,
STRUCT, ENUM, CLASS, CLASS_INSTANCE
```

---

## Summary

Bhasa provides **40+ built-in functions** covering:

- ✅ **I/O**: Console and file operations
- ✅ **Arrays**: Manipulation and iteration
- ✅ **Strings**: Parsing, formatting, searching
- ✅ **Math**: Arithmetic and comparison
- ✅ **JSON**: Serialization and deserialization
- ✅ **Hashes**: Key-value operations
- ✅ **Types**: Conversion and introspection

All built-ins have **Bengali names** matching the language's design philosophy.

For usage examples in context, see the main [README.md](./README.md).

