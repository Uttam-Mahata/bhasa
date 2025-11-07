# ভাষা (Bhasa) - A Bengali Programming Language

A **compiled** programming language that uses Bengali keywords, built with Go for India. 🇮🇳

## Features

- 🇮🇳 **Bengali keywords and syntax**
- 🔤 **Bengali variable names** - Full Unicode support for identifiers
- ⚡ **Bytecode compiler** (3-10x faster than interpretation!)
- 🖥️ **Stack-based virtual machine**
- 📝 Variables and functions with closures
- 🔢 Numbers, strings, booleans, arrays, and hash maps
- 🔄 Control flow (if-else, while loops)
- 🚀 Interactive REPL
- 📚 **20+ Built-in functions** (string methods, math functions, array operations)
- 🎯 Recursion and higher-order functions
- 🔗 **Logical operators** (&&, ||, !)
- 🧮 **Math functions** (power, sqrt, abs, max, min)
- 📝 **String manipulation** (split, join, uppercase, lowercase, trim, replace)
- 🎨 **Object-Oriented Programming** (classes, objects, methods)

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
| class | শ্রেণী | Class declaration |
| new | নতুন | Object instantiation |
| this | এই | Current instance |

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

### Object-Oriented Programming
```bengali
// Define a class
শ্রেণী গাড়ি {
    শুরু = ফাংশন() {
        লেখ("গাড়ি চলছে!");
    };
    
    থামো = ফাংশন() {
        লেখ("গাড়ি থেমেছে!");
    };
}

// Create an instance
ধরি আমার_গাড়ি = নতুন গাড়ি();
```

See [OOP.md](OOP.md) for detailed OOP documentation.

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

