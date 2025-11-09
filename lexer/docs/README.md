# Lexer Documentation

## Overview

The **Lexer** (lexical analyzer/tokenizer) is the first stage of the Bhasa language interpreter. It converts raw source code text into a stream of **tokens** that can be parsed by the parser.

## Table of Contents

- [What is a Lexer?](#what-is-a-lexer)
- [Key Features](#key-features)
- [Architecture](#architecture)
- [Token Types](#token-types)
- [Usage](#usage)
- [Unicode Support](#unicode-support)
- [Position Tracking](#position-tracking)
- [For More Details](#for-more-details)

## What is a Lexer?

A **lexer** (also called a **scanner** or **tokenizer**) performs **lexical analysis** - the process of converting a sequence of characters into a sequence of tokens. It's the first phase of compilation/interpretation.

### Example

**Input (Source Code):**
```bengali
ধরি x = ৫;
```

**Output (Token Stream):**
```
Token{Type: LET, Literal: "ধরি", Line: 1, Column: 1}
Token{Type: IDENT, Literal: "x", Line: 1, Column: 5}
Token{Type: ASSIGN, Literal: "=", Line: 1, Column: 7}
Token{Type: INT, Literal: "5", Line: 1, Column: 9}
Token{Type: SEMICOLON, Literal: ";", Line: 1, Column: 10}
Token{Type: EOF, Literal: "", Line: 1, Column: 11}
```

## Key Features

### 1. **Unicode Support** 🌐
- Full support for Bengali script (UTF-8)
- Uses `rune` type (int32) instead of `byte` to handle multi-byte characters
- Supports Bengali vowel signs (মাত্রা) and diacritics

### 2. **Bengali & Arabic Numerals** 🔢
- Accepts both Arabic (0-9) and Bengali (০-৯) digits
- Automatically converts Bengali digits to Arabic internally
- Examples: `১২৩` → `123`, `৫০` → `50`

### 3. **Position Tracking** 📍
- Tracks line and column numbers for every token
- Essential for error reporting with accurate locations
- Automatically handles newlines and column counting

### 4. **Peek-Ahead** 👀
- Can look at the next character without consuming it
- Necessary for multi-character operators like `==`, `<=`, `>>`
- Enables single-pass tokenization

### 5. **Comment Support** 💬
- Line comments start with `//`
- Comments are automatically skipped during tokenization
- Whitespace and comments are ignored

### 6. **Comprehensive Operators** ⚙️
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `&&`, `||`, `!`
- Bitwise: `&`, `|`, `^`, `~`, `<<`, `>>`
- Assignment: `=`
- Function arrow: `=>`

## Architecture

### Lexer Structure

```go
type Lexer struct {
    input        []rune  // Input as rune slice (Unicode support)
    position     int     // Current position (current char)
    readPosition int     // Next reading position (next char)
    ch           rune    // Current character under examination
    line         int     // Current line number (starts at 1)
    column       int     // Current column number (starts at 0)
}
```

### State Machine

The lexer operates as a **state machine** with the following states:

1. **Start**: Begin processing from position 0
2. **Read Character**: Consume next character
3. **Classify**: Determine token type based on character
4. **Multi-char Check**: Use peek-ahead for operators like `==`, `&&`
5. **Build Token**: Create token with type, literal, position
6. **Repeat**: Continue until EOF

### Single-Pass Design

The lexer processes input in a **single pass** from left to right:
- No backtracking required
- Efficient memory usage
- O(n) time complexity where n = input length

## Token Types

The lexer recognizes several categories of tokens:

### 1. **Keywords** (Bengali)
```
ধরি        (let)          - Variable declaration
ফাংশন      (function)     - Function definition
যদি        (if)           - Conditional
নাহলে      (else)         - Else clause
ফেরত       (return)       - Return statement
সত্য       (true)         - Boolean true
মিথ্যা     (false)        - Boolean false
যতক্ষণ     (while)        - While loop
```

### 2. **Identifiers**
- Variable names and function names
- Must start with letter or underscore
- Can contain letters, digits, underscores
- Supports Bengali characters and vowel signs

### 3. **Literals**
- **Integers**: `123`, `৪৫৬` (both Arabic and Bengali)
- **Strings**: `"হ্যালো"`, `"Hello"`

### 4. **Operators**
- Arithmetic, comparison, logical, bitwise (see Key Features)

### 5. **Delimiters**
- `(` `)` `{` `}` `[` `]` - Grouping
- `,` `;` `:` `.` - Separators

### 6. **Special**
- `EOF` - End of file
- `ILLEGAL` - Unrecognized character

## Usage

### Basic Usage

```go
import (
    "bhasa/lexer"
    "bhasa/token"
)

// Create a new lexer
input := `ধরি x = ১০;`
l := lexer.New(input)

// Get tokens one by one
for {
    tok := l.NextToken()
    
    fmt.Printf("Type: %s, Literal: %s, Line: %d, Col: %d\n",
        tok.Type, tok.Literal, tok.Line, tok.Column)
    
    if tok.Type == token.EOF {
        break
    }
}
```

### Tokenizing a Complete Program

```go
input := `
ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
};
`

l := lexer.New(input)
tokens := []token.Token{}

for {
    tok := l.NextToken()
    tokens = append(tokens, tok)
    if tok.Type == token.EOF {
        break
    }
}
```

## Unicode Support

### Why Runes?

Go's `string` type is a sequence of bytes, but Unicode characters can be multiple bytes. Bengali characters often use 2-3 bytes per character.

**Solution**: Use `[]rune` instead of `string`:
- `rune` is an alias for `int32`
- Represents a single Unicode code point
- Bengali character "ক" = single rune (multiple bytes)

### Bengali Script Features

#### 1. **Vowel Signs (মাত্রা)**
```
ক  → single character
কা → ক + া (base + vowel sign)
কি → ক + ি (base + vowel sign)
```

The lexer recognizes vowel signs (U+0981 to U+09CD) as part of identifiers.

#### 2. **Combining Marks**
```
কং → ক + ং (anusvara)
কঃ → ক + ঃ (visarga)
```

These are handled using `unicode.Mn` (nonspacing marks) category.

#### 3. **Conjuncts (যুক্তাক্ষর)**
```
ক্ষ → ক + ্ + ষ
স্ত → স + ্ + ত
```

The hasant (্) and following consonant form conjuncts.

### Bengali Digit Conversion

```go
// Input: "১২৩"
// Lexer sees: rune '১', rune '২', rune '৩'
// Reads as: "১২৩"
// Converts to: "123"
// Returns: Token{Type: INT, Literal: "123"}
```

**Conversion Table:**
```
০ → 0    ৫ → 5
১ → 1    ৬ → 6
২ → 2    ৭ → 7
৩ → 3    ৮ → 8
৪ → 4    ৯ → 9
```

## Position Tracking

### Line and Column Tracking

The lexer maintains accurate position information:

```
Line 1: ধরি x = ৫;
        ^   ^   ^  ^
        1   5   9  11 (columns)

Line 2: ধরি y = ১০;
        ^   ^   ^   ^
        1   5   9   12 (columns)
```

### Newline Handling

When `\n` is encountered:
1. Increment `line` counter
2. Reset `column` to 0
3. Continue reading

### Token Position

Each token records its starting position:

```go
tok := token.Token{
    Type:    token.LET,
    Literal: "ধরি",
    Line:    1,      // Starting line
    Column:  1,      // Starting column
}
```

### Use in Error Reporting

Position information enables helpful error messages:

```
Error at line 5, column 12: unexpected token '!'
    ধরি x = ১০!
               ^
```

## For More Details

For comprehensive documentation with implementation details, examples, and advanced topics, see:

- **[Lexer Implementation Guide](./lexer-documentation.md)** - Complete technical reference

## Quick Reference

### Core Functions

| Function | Purpose |
|----------|---------|
| `New(input string)` | Create new lexer |
| `NextToken()` | Get next token |
| `readChar()` | Consume next character |
| `peekChar()` | Look ahead without consuming |
| `readIdentifier()` | Read variable/keyword name |
| `readNumber()` | Read integer literal |
| `readString()` | Read string literal |
| `skipWhitespace()` | Skip spaces, tabs, newlines |
| `skipComment()` | Skip line comments |

### Helper Functions

| Function | Purpose |
|----------|---------|
| `isLetter(ch)` | Check if character is letter |
| `isDigit(ch)` | Check if character is Arabic digit |
| `isBengaliDigit(ch)` | Check if character is Bengali digit |
| `isBengaliVowelSign(ch)` | Check if character is vowel sign |

## Example: Complete Tokenization

**Input:**
```bengali
// ফ্যাক্টোরিয়াল গণনা
ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
};
```

**Token Stream:**
```
// Comment skipped
LET      "ধরি"              line 2, col 1
IDENT    "ফ্যাক্টোরিয়াল"    line 2, col 5
ASSIGN   "="                line 2, col 18
FUNCTION "ফাংশন"            line 2, col 20
LPAREN   "("                line 2, col 26
IDENT    "n"                line 2, col 27
RPAREN   ")"                line 2, col 28
LBRACE   "{"                line 2, col 30
IF       "যদি"              line 3, col 5
LPAREN   "("                line 3, col 9
IDENT    "n"                line 3, col 10
LTE      "<="               line 3, col 12
INT      "1"                line 3, col 15
RPAREN   ")"                line 3, col 16
LBRACE   "{"                line 3, col 18
RETURN   "ফেরত"             line 4, col 9
INT      "1"                line 4, col 14
SEMICOLON ";"               line 4, col 15
RBRACE   "}"                line 5, col 5
RETURN   "ফেরত"             line 6, col 5
IDENT    "n"                line 6, col 10
ASTERISK "*"                line 6, col 12
IDENT    "ফ্যাক্টোরিয়াল"    line 6, col 14
LPAREN   "("                line 6, col 27
IDENT    "n"                line 6, col 28
MINUS    "-"                line 6, col 30
INT      "1"                line 6, col 32
RPAREN   ")"                line 6, col 33
SEMICOLON ";"               line 6, col 34
RBRACE   "}"                line 7, col 1
SEMICOLON ";"               line 7, col 2
EOF      ""                 line 7, col 3
```

## Benefits of Lexer Design

### 1. **Separation of Concerns**
- Lexer: Character → Token
- Parser: Token → AST
- Clear boundaries and responsibilities

### 2. **Error Detection**
- Invalid characters detected immediately
- Position tracking for error messages
- Fail fast on illegal input

### 3. **Flexibility**
- Easy to add new operators
- Easy to add new keywords
- Minimal changes to parser

### 4. **Performance**
- Single-pass algorithm
- No backtracking
- Efficient memory usage

### 5. **Maintainability**
- Clear, readable code
- Easy to test individual functions
- Well-defined token types

---

## Summary

The **Bhasa Lexer** is a robust, Unicode-aware tokenizer that:
- ✅ Handles Bengali script and vowel signs
- ✅ Supports both Arabic and Bengali numerals
- ✅ Tracks positions for error reporting
- ✅ Recognizes 30+ Bengali keywords
- ✅ Supports comprehensive operator set
- ✅ Implements efficient single-pass algorithm
- ✅ Provides clean interface for parser

For implementation details, see [lexer-documentation.md](./lexer-documentation.md).

