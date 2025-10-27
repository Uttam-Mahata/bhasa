# ভাষা (Bhasa) - A Bengali Programming Language

A **compiled** programming language that uses Bengali keywords, built with Go for India. 🇮🇳

## Features

- 🇮🇳 Bengali keywords and syntax
- ⚡ **Bytecode compiler** (3-10x faster than interpretation!)
- 🖥️ **Stack-based virtual machine**
- 📝 Variables and functions with closures
- 🔢 Numbers, strings, booleans, arrays, and hash maps
- 🔄 Control flow (if-else, while/for loops, break/continue)
- 🚀 Interactive REPL
- 📚 Built-in functions (20+ functions)
- 🎯 Recursion and higher-order functions
- 🔤 String manipulation methods
- 🧮 Math functions
- ⚡ Logical operators (&&, ||)

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

