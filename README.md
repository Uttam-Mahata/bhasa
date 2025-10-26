# ভাষা (Bhasa) - A Bengali Programming Language

A programming language that uses Bengali keywords, built with Go.

## Features

- 🇧🇩 Bengali keywords and syntax
- 📝 Variables and functions
- 🔢 Numbers, strings, booleans, arrays, and hash maps
- 🔄 Control flow (if-else, while loops)
- 🚀 Interactive REPL
- 📚 Built-in functions

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
├── main.go           # Entry point and REPL
├── token/            # Token definitions
├── lexer/            # Lexical analyzer
├── ast/              # Abstract Syntax Tree
├── parser/           # Parser implementation
├── evaluator/        # Expression evaluator
├── object/           # Object system
└── examples/         # Example programs
```

