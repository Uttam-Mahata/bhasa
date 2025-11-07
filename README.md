# ভাষা (Bhasa) - A Bengali Programming Language

A **compiled** programming language that uses Bengali keywords, built with Go for India. 🇮🇳

## Features

- 🇮🇳 **Bengali keywords and syntax**
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
- 🔧 **Self-hosting capable** - Can write a compiler in Bhasa itself!
- 🧮 **Math functions** (power, sqrt, abs, max, min)
- 📝 **String manipulation** (split, join, uppercase, lowercase, trim, replace)

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

### Type Casting Functions
- **পূর্ণসংখ্যা(value)** - Cast to integer (from string, boolean, or integer)
- **অক্ষর_রূপ(value)** - Cast to character (from integer code or string)
- **ছোট_সংখ্যা(value)** - Cast to short integer (16-bit, -32768 to 32767)
- **বাইট(value)** - Cast to byte (8-bit, 0 to 255)
- **দশমিক(value)** - Cast to float/double (from string, integer, or boolean)
- **বুলিয়ান(value)** - Cast to boolean (from any type)
- **লেখা_রূপ(value)** - Cast to string representation (from any type)

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

## Running the REPL

```bash
./bhasa
```

Then you can type Bengali code interactively!

## Project Structure

```
bhasa/
├── main.go           # Entry point
├── token/            # Token definitions
├── lexer/            # Lexical analyzer
├── ast/              # Abstract Syntax Tree
├── parser/           # Parser implementation
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

