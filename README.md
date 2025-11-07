# ভাষা (Bhasa) - A Bengali Programming Language

A **compiled** programming language that uses Bengali keywords, built with Go as a hobby project.

## Features

- 📝 **Bengali keywords and syntax**
- 🔤 **Bengali variable names** - Full Unicode support for identifiers
- ⚡ **Bytecode compiler** (3-10x faster than interpretation!)
- 🖥️ **Stack-based virtual machine**
- 📝 Variables and functions with closures
- 🔢 Numbers, strings, booleans, arrays, and hash maps
- 🔄 Control flow (if-else, while, for loops)
- 🚀 Interactive REPL
- 📚 **30+ Built-in functions** (file I/O, string methods, math functions, array operations)
- 🎯 Recursion and higher-order functions
- 🔗 **Logical operators** (&&, ||, !)
- 🔢 **Bitwise operators** (&, |, ^, ~, <<, >>)
- 📁 **File I/O support** - Read and write files
- 🔧 **Self-hosting capable** - Full compiler written in Bhasa itself!
- 🧮 **Math functions** (power, sqrt, abs, max, min)
- 📝 **String manipulation** (split, join, uppercase, lowercase, trim, replace)
- 🔢 **Multiple numeric types** (Byte, Short, Int, Long, Float, Double) with type casting
- 📦 **Module system** with `অন্তর্ভুক্ত` (import) support
- 🎨 **Optional static typing** - Add type annotations for better code safety!

## Bengali Keywords

| English | Bengali | Usage |
|---------|---------|-------|
| let | ধরি | Variable declaration |
| function | ফাংশন | Function declaration |
| if | যদি | Conditional |
| else | নাহলে | Else clause |
| return | ফেরত | Return statement |
| true | সত্য | Boolean true |
| false | মিথ্যা | Boolean false |
| while | যতক্ষণ | While loop |
| for | পর্যন্ত | For loop |
| break | বিরতি | Break statement |
| continue | চালিয়ে_যাও | Continue statement |
| null | নাল | Null value |
| import | অন্তর্ভুক্ত | Import module |

## Type Keywords (Optional Static Typing)

| Type | Bengali | Usage |
|------|---------|-------|
| integer | পূর্ণসংখ্যা | Integer type annotation |
| string | লেখা | String type annotation |
| boolean | বুলিয়ান | Boolean type annotation |
| array | তালিকা | Array type annotation |
| hash | হ্যাশ | Hash type annotation |
| function | ফাংশন_টাইপ | Function type annotation |

## Quick Links

- 📖 [**Self-Hosting Compiler Guide**](SELF_HOSTING.md) - Learn how the Bhasa compiler is written in Bhasa
- 🔧 [**Compiler API Documentation**](COMPILER_API.md) - Complete API reference for the self-hosted compiler
- 🎨 [**Static Typing Guide**](STATIC_TYPING.md) - Learn about optional type annotations
- 🧪 [**Test Suite**](tests/) - Comprehensive tests for lexer, parser, compiler, and bootstrap

## Installation

```bash
go build -o bhasa
./bhasa
```

## Example Programs

### Hello World
```bengali
লেখ("নমস্কার বিশ্ব!")
```

### Variables and Math
```bengali
ধরি x = ৫;
ধরি y = ১০;
ধরি যোগফল = x + y;
লেখ(যোগফল);
```

### Bengali Variable Names
```bengali
// Variables can use Bengali names
ধরি নাম = "রহিম";
ধরি বয়স = ২৫;
ধরি বেতন = ৫০০০০;

// Functions with Bengali names
ধরি যোগফল_বের_করো = ফাংশন(ক, খ) {
    ফেরত ক + খ;
};

লেখ(যোগফল_বের_করো(১০, ২০));  // Output: 30
```

### Functions
```bengali
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};

লেখ(যোগ(৫, ৩));
```

### Conditionals
```bengali
ধরি x = ১০;
যদি (x > ৫) {
    লেখ("x is greater than 5");
} নাহলে {
    লেখ("x is not greater than 5");
}
```

### Logical Operators
```bengali
// AND operator
যদি (x > 5 && y < 10) {
    লেখ("Both conditions are true");
}

// OR operator
যদি (score < 40 || score > 90) {
    লেখ("Special attention needed");
}
```

### Bitwise Operators
```bengali
// Bitwise AND
ধরি a = ১২ & ১০;  // 8

// Bitwise OR
ধরি b = ১২ | ১০;  // 14

// Bitwise XOR
ধরি c = ১২ ^ ১০;  // 6

// Bitwise NOT
ধরি d = ~৫;  // -6

// Left Shift
ধরি e = ৫ << ২;  // 20

// Right Shift
ধরি f = ২০ >> ২;  // 5
```

### For Loops
```bengali
// C-style for loop
পর্যন্ত (ধরি i = ০; i < ১০; i = i + ১) {
    লেখ(i);
}
```

## Self-Hosting Capability

Bhasa now has all the features needed to write a compiler for itself! See `examples/simple_lexer_demo.ভাষা` for a working lexer written entirely in Bhasa.

**Key self-hosting features:**
- Character access and manipulation (`অক্ষর`, `কোড`)
- String parsing (`সংখ্যা`)
- File I/O for reading/writing source files
- For loops for iteration
- Arrays and hashes for data structures

## Built-in Functions

### Basic Functions
- **লেখ()** - Print to console
- **দৈর্ঘ্য()** - Length of string/array
- **টাইপ()** - Get type of value

### String Methods
- **বিভক্ত(str, delimiter)** - Split string
- **যুক্ত(arr, delimiter)** - Join array elements
- **উপরে(str)** - Convert to uppercase
- **নিচে(str)** - Convert to lowercase
- **ছাঁটো(str)** - Trim whitespace
- **প্রতিস্থাপন(str, old, new)** - Replace text
- **খুঁজুন(str, substr)** - Find substring index

### Character/Conversion Functions (Self-Hosting Support)
- **অক্ষর(str, index)** - Get character at index
- **কোড(char)** - Get Unicode code point
- **অক্ষর_থেকে_কোড(code)** - Create character from code
- **সংখ্যা(str)** - Parse string to integer
- **লেখা(num)** - Convert integer to string

### File I/O Functions
- **ফাইল_পড়ো(path)** - Read file contents
- **ফাইল_লেখো(path, content)** - Write to file
- **ফাইল_যোগ(path, content)** - Append to file
- **ফাইল_আছে(path)** - Check if file exists

### Math Functions
- **শক্তি(base, exp)** - Power
- **বর্গমূল(n)** - Square root
- **পরম(n)** - Absolute value
- **সর্বোচ্চ(a, b)** - Maximum
- **সর্বনিম্ন(a, b)** - Minimum
- **গোলাকার(n)** - Round number

### Array Functions
- **প্রথম(arr)** - First element
- **শেষ(arr)** - Last element
- **বাকি(arr)** - All but first
- **যোগ(arr, element)** - Add element
- **উল্টাও(arr)** - Reverse array

## Type Casting Functions

Bhasa supports multiple numeric types with explicit casting:

```bengali
// Numeric type conversions
ধরি x = ১০০;
ধরি b = বাইট(x);                    // Convert to Byte (0-255)
ধরি s = ছোট_সংখ্যা(x);              // Convert to Short (-32768 to 32767)
ধরি i = পূর্ণসংখ্যা(x);              // Convert to Int
ধরি l = দীর্ঘ_সংখ্যা(x);              // Convert to Long
ধরি f = দশমিক(x);                   // Convert to Float
ধরি d = দশমিক_দ্বিগুণ(x);            // Convert to Double

// Character conversion
ধরি ch = অক্ষর_রূপান্তর("A");       // String to Char
```

### Supported Numeric Types
- **বাইট (Byte)**: 0 to 255
- **ছোট_সংখ্যা (Short)**: -32,768 to 32,767
- **পূর্ণসংখ্যা (Int)**: -2,147,483,648 to 2,147,483,647
- **দীর্ঘ_সংখ্যা (Long)**: Full 64-bit integer
- **দশমিক (Float)**: 32-bit floating point
- **দশমিক_দ্বিগুণ (Double)**: 64-bit floating point

## Self-Hosting Compiler

Bhasa includes a **complete self-hosted compiler** written entirely in Bhasa itself! This means you can compile Bhasa programs using a compiler written in Bhasa.

### Self-Hosted Modules (in `modules/`)

All compiler components are implemented in `.ভাষা` files:

- **টোকেন.ভাষা** - Token type definitions and utilities
- **লেক্সার.ভাষা** - Lexical analyzer (tokenizer)
- **এএসটি.ভাষা** - Abstract Syntax Tree node structures
- **পার্সার.ভাষা** - Pratt parser with operator precedence
- **প্রতীক_টেবিল.ভাষা** - Symbol table for scoping
- **কোড.ভাষা** - Bytecode instruction encoding/decoding
- **কম্পাইলার.ভাষা** - AST to bytecode compiler
- **মডিউল_লোডার.ভাষা** - Module import system
- **ভাষা_কম্পাইলার.ভাষা** - Main compiler driver

### Using the Self-Hosted Compiler

```bengali
// Import compiler modules
অন্তর্ভুক্ত "modules/ভাষা_কম্পাইলার";

// Compile a file
ধরি ফলাফল = ফাইল_কম্পাইল_করো("my_program.ভাষা");
যদি (ফলাফল["সফল"]) {
    লেখ("কম্পাইল সফল!");
} নাহলে {
    লেখ("ত্রুটি: " + ফলাফল["ত্রুটি"]);
}
```

For complete documentation, see [SELF_HOSTING.md](SELF_HOSTING.md) and [COMPILER_API.md](COMPILER_API.md).

## Running the REPL

```bash
./bhasa
```

Then you can type Bengali code interactively!

## Project Structure

```
bhasa/
├── main.go                    # Entry point
├── token/                     # Token definitions (Go)
├── lexer/                     # Lexical analyzer (Go)
├── ast/                       # Abstract Syntax Tree (Go)
├── parser/                    # Parser implementation (Go)
├── modules/                   # Self-hosted compiler modules (.ভাষা)
│   ├── টোকেন.ভাষা             # Token module
│   ├── লেক্সার.ভাষা            # Lexer module
│   ├── এএসটি.ভাষা              # AST module
│   ├── পার্সার.ভাষা            # Parser module
│   ├── প্রতীক_টেবিল.ভাষা       # Symbol table
│   ├── কোড.ভাষা                # Bytecode instructions
│   ├── কম্পাইলার.ভাষা          # Compiler module
│   ├── মডিউল_লোডার.ভাষা        # Module loader
│   └── ভাষা_কম্পাইলার.ভাষা     # Main compiler driver
├── tests/                     # Test files (.ভাষা)
│   ├── lexer_test.ভাষা
│   ├── parser_test.ভাষা
│   ├── compiler_test.ভাষা
│   └── bootstrap_test.ভাষা
├── compiler/         # Bytecode compiler
│   ├── compiler.go   # AST → Bytecode
│   └── symbol_table.go # Variable scoping
├── code/             # Bytecode instruction set
├── vm/               # Virtual machine
│   ├── vm.go         # Stack-based VM
│   └── frame.go      # Call frames
├── object/           # Object system
├── repl/             # Interactive REPL
└── examples/         # Example programs
```

## Architecture

**Compilation Pipeline:**
```
Bengali Source → Lexer → Parser → AST → Compiler → Bytecode → VM → Execution
```

**Key Components:**
- **Compiler**: Translates AST to bytecode (35+ opcodes)
- **Virtual Machine**: Stack-based execution engine
- **Symbol Table**: Manages variable scopes (global, local, free, builtin)
- **Closures**: Full support for lexical scoping

See [COMPILER.md](COMPILER.md) for detailed architecture documentation.

