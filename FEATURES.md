# ভাষা (Bhasa) - Features & Implementation Details

## Overview

Bhasa is a fully-functional programming language that uses Bengali keywords and supports Bengali numerals. It's implemented in Go and includes a complete interpreter with lexer, parser, AST, and evaluator.

## Core Components

### 1. Lexer (`lexer/lexer.go`)
- **UTF-8 Support**: Properly handles multi-byte Unicode characters including Bengali script
- **Bengali Vowel Signs**: Recognizes all Bengali মাত্রা (vowel signs) and diacritics (U+0981 to U+09CD)
- **Numeral Support**: Accepts both Bengali (০-৯) and Arabic (0-9) numerals
- **Comment Support**: Single-line comments with `//`
- **String Literals**: Full support for Bengali text in strings

### 2. Parser (`parser/parser.go`)
- **Pratt Parsing**: Implements operator precedence parsing
- **Statement Types**:
  - Variable declarations (`ধরি`)
  - Variable assignments
  - Return statements (`ফেরত`)
  - While loops (`যতক্ষণ`)
  - Expression statements
  - Block statements

- **Expression Types**:
  - Literals (integers, strings, booleans)
  - Identifiers
  - Prefix expressions (`!`, `-`)
  - Infix expressions (`+`, `-`, `*`, `/`, `%`, `==`, `!=`, `<`, `>`, `<=`, `>=`)
  - If-else expressions (`যদি`/`নাহলে`)
  - Function literals (`ফাংশন`)
  - Call expressions
  - Array literals
  - Index expressions
  - Hash literals

### 3. AST (`ast/ast.go`)
- Complete Abstract Syntax Tree representation
- Supports all statement and expression types
- Pretty-printing for debugging

### 4. Object System (`object/object.go`)
- **Data Types**:
  - Integer
  - String
  - Boolean
  - Array
  - Hash (with hashable keys)
  - Function
  - Null
  - Return values
  - Errors

- **Environment**: Variable scoping with enclosed environments for closures

### 5. Evaluator (`evaluator/evaluator.go`)
- **Tree-walking interpreter**: Evaluates AST nodes recursively
- **First-class functions**: Functions are values that can be passed around
- **Closures**: Functions capture their defining environment
- **Recursion**: Full support for recursive functions
- **Error handling**: Propagates errors through evaluation

### 6. Built-in Functions (`evaluator/builtins.go`)

| Function | Bengali | Purpose |
|----------|---------|---------|
| print | লেখ | Output to console |
| length | দৈর্ঘ্য | Get length of strings/arrays |
| first | প্রথম | First element of array |
| last | শেষ | Last element of array |
| rest | বাকি | All but first element |
| push | যোগ | Add element to array |
| type | টাইপ | Get type of value |

### 7. REPL (`repl/repl.go`)
- Interactive shell for live coding
- Maintains environment across inputs
- Bengali command support (`প্রস্থান` to exit)
- Error display with helpful messages

### 8. Main Program (`main.go`)
- Dual mode: REPL or file execution
- File reading and execution
- Error reporting

## Language Features

### ✅ Variables
- Declaration with `ধরি`
- Reassignment support
- Lexical scoping

### ✅ Data Types
- Integers (Bengali and Arabic numerals)
- Strings (full Unicode support)
- Booleans (`সত্য`, `মিথ্যা`)
- Arrays (dynamic, heterogeneous)
- Hash maps (key-value pairs)

### ✅ Operators
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `!`
- String concatenation with `+`

### ✅ Control Flow
- If-else statements (`যদি`/`নাহলে`)
- While loops (`যতক্ষণ`)
- Early returns with `ফেরত`

### ✅ Functions
- First-class functions
- Higher-order functions
- Closures
- Recursion
- Multiple parameters
- Return values

### ✅ Collections
- Arrays with indexing
- Hash maps with string/number/boolean keys
- Array manipulation functions

## Technical Highlights

### Unicode Handling
The lexer properly handles Bengali script which uses:
- Base consonants (ক, খ, গ, ঘ, etc.)
- Vowel signs/মাত্রা (া, ি, ী, ু, ূ, etc.)
- Diacritics (ং, ঃ, ঁ)
- Hasant (্) for conjuncts

### Number System
Converts Bengali numerals to internal representation:
- ০১২৩৪৫৬৭৮৯ → 0123456789

### Error Messages
Clear error reporting for:
- Lexical errors (unexpected characters)
- Parse errors (syntax mistakes)
- Runtime errors (type mismatches, undefined variables)

## Example Programs

### 1. Hello World
```bengali
লেখ("নমস্কার বিশ্ব!");
```

### 2. Factorial (Recursion)
```bengali
ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
    যদি (n == ০) {
        ফেরত ১;
    } নাহলে {
        ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
    }
};
```

### 3. Higher-Order Functions
```bengali
ধরি দুইগুণ = ফাংশন(f, x) {
    ফেরত f(f(x));
};

ধরি যোগএক = ফাংশন(x) {
    ফেরত x + ১;
};

লেখ(দুইগুণ(যোগএক, ৫));  // Output: 7
```

## Performance Characteristics

- **Lexing**: O(n) where n is input length
- **Parsing**: O(n) with operator precedence
- **Evaluation**: Tree-walking (not optimized for speed)
- **Memory**: Garbage collected by Go

## Future Enhancement Ideas

1. **More Bengali Keywords**:
   - `পর্যন্ত` (for/until)
   - `বিরতি` (break)
   - `চালিয়ে_যাও` (continue)
   - `শ্রেণী` (class)

2. **Additional Built-ins**:
   - File I/O operations
   - Math functions
   - String manipulation
   - Date/time handling

3. **Advanced Features**:
   - List comprehensions
   - Pattern matching
   - Async/await
   - Module system
   - Standard library

4. **Optimizations**:
   - Bytecode compiler
   - Virtual machine
   - JIT compilation
   - Constant folding

5. **Tooling**:
   - Syntax highlighter
   - Formatter
   - Linter
   - Package manager
   - Documentation generator

## Design Philosophy

1. **Accessibility**: Make programming accessible to Bengali speakers
2. **Simplicity**: Clean, intuitive syntax
3. **Completeness**: Full-featured language, not a toy
4. **Correctness**: Proper Unicode handling for Bengali script
5. **Extensibility**: Easy to add new features

## Inspiration

This language draws inspiration from:
- **Monkey Language** (Writing an Interpreter in Go book)
- **JavaScript**: Dynamic typing, first-class functions
- **Python**: Clean syntax, built-in functions
- **Go**: Simple, readable implementation

## Project Statistics

- **Lines of Code**: ~2,500
- **Packages**: 7 (token, lexer, ast, parser, object, evaluator, repl)
- **Built-in Functions**: 7
- **Keywords**: 8
- **Example Programs**: 9
- **Development Time**: Single session implementation

## Testing

All example programs have been tested and verified:
- ✅ hello.bhasa
- ✅ variables.bhasa
- ✅ functions.bhasa
- ✅ conditionals.bhasa
- ✅ loops.bhasa
- ✅ arrays.bhasa
- ✅ hash.bhasa
- ✅ fibonacci.bhasa
- ✅ comprehensive.bhasa

Run all tests with:
```bash
./run_examples.sh
```

## Conclusion

Bhasa is a complete, working programming language that demonstrates:
- Full Bengali script support
- Modern language features
- Clean interpreter implementation
- Practical usability

It serves as both a functional programming language and an educational resource for understanding language implementation.

