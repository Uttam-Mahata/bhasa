# Token Package Documentation

## Overview

The **Token** package defines the fundamental building blocks of the Bhasa language's lexical structure. It contains token types, token structures, keyword mappings, and utility functions for handling Bengali numerals.

## Table of Contents

- [What is a Token?](#what-is-a-token)
- [Token Structure](#token-structure)
- [Token Types](#token-types)
- [Keywords](#keywords)
- [Bengali Numeral Support](#bengali-numeral-support)
- [Utility Functions](#utility-functions)
- [Usage Examples](#usage-examples)

## What is a Token?

A **token** is the smallest meaningful unit in source code. The lexer breaks source code into tokens, which the parser then uses to build an Abstract Syntax Tree (AST).

### Example

**Source Code:**
```bengali
ধরি x = ৫;
```

**Tokens:**
```
LET        "ধরি"
IDENT      "x"
ASSIGN     "="
INT        "5"
SEMICOLON  ";"
```

## Token Structure

### Token Type

```go
type TokenType string
```

A `TokenType` is a string constant that identifies the category of a token.

### Token

```go
type Token struct {
    Type    TokenType  // Token type (LET, IDENT, INT, etc.)
    Literal string     // Actual text from source code
    Line    int        // Line number (1-indexed)
    Column  int        // Column number (starts at 1)
}
```

#### Fields

| Field | Type | Purpose | Example |
|-------|------|---------|---------|
| `Type` | `TokenType` | Category of the token | `token.LET` |
| `Literal` | `string` | Actual source text | `"ধরি"` |
| `Line` | `int` | Source line number | `1` |
| `Column` | `int` | Source column number | `1` |

#### Example Token

```go
tok := token.Token{
    Type:    token.LET,
    Literal: "ধরি",
    Line:    1,
    Column:  1,
}
```

## Token Types

### Special Tokens

```go
ILLEGAL  // Unrecognized character
EOF      // End of file
```

**Usage:**
- `ILLEGAL`: Represents invalid input
- `EOF`: Marks the end of the token stream

### Identifiers and Literals

```go
IDENT   // Variable/function names (e.g., "myVariable", "ফাংশনNaam")
INT     // Integer literals (e.g., "123", "৫")
STRING  // String literals (e.g., "hello", "হ্যালো")
```

### Operators

#### Arithmetic Operators

```go
PLUS     "+"   // Addition
MINUS    "-"   // Subtraction
ASTERISK "*"   // Multiplication
SLASH    "/"   // Division
PERCENT  "%"   // Modulo
```

#### Comparison Operators

```go
LT     "<"    // Less than
GT     ">"    // Greater than
EQ     "=="   // Equal to
NOT_EQ "!="   // Not equal to
LTE    "<="   // Less than or equal
GTE    ">="   // Greater than or equal
```

#### Logical Operators

```go
BANG "!"    // Logical NOT
AND  "&&"   // Logical AND
OR   "||"   // Logical OR
```

#### Bitwise Operators

```go
BIT_AND   "&"    // Bitwise AND
BIT_OR    "|"    // Bitwise OR
BIT_XOR   "^"    // Bitwise XOR
BIT_NOT   "~"    // Bitwise NOT
LSHIFT    "<<"   // Left shift
RSHIFT    ">>"   // Right shift
```

#### Assignment

```go
ASSIGN "="   // Assignment operator
```

### Delimiters

```go
COMMA     ","   // Separator in lists
SEMICOLON ";"   // Statement terminator
COLON     ":"   // Type annotation separator

LPAREN   "("    // Left parenthesis
RPAREN   ")"    // Right parenthesis
LBRACE   "{"    // Left brace
RBRACE   "}"    // Right brace
LBRACKET "["    // Left bracket
RBRACKET "]"    // Right bracket

DOT   "."       // Member access
ARROW "=>"      // Function arrow
```

## Keywords

All keywords in Bhasa are in Bengali script. The language uses meaningful Bengali words rather than transliterations.

### Control Flow Keywords

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `LET` | ধরি | let | Variable declaration |
| `FUNCTION` | ফাংশন | function | Function definition |
| `IF` | যদি | if | Conditional statement |
| `ELSE` | নাহলে | else | Alternative branch |
| `RETURN` | ফেরত | return | Return statement |
| `WHILE` | যতক্ষণ | while | While loop |
| `FOR` | পর্যন্ত | for | For loop |
| `BREAK` | বিরতি | break | Break loop |
| `CONTINUE` | চালিয়ে_যাও | continue | Continue loop |

**Example:**
```bengali
ধরি x = ১০;           // let x = 10;
যদি (x > ৫) {         // if (x > 5) {
    ফেরত সত্য;        //     return true;
}                     // }
```

### Boolean Keywords

| Token | Bengali | English | Value |
|-------|---------|---------|-------|
| `TRUE` | সত্য | true | Boolean true |
| `FALSE` | মিথ্যা | false | Boolean false |

**Example:**
```bengali
ধরি isValid = সত্য;   // let isValid = true;
ধরি isEmpty = মিথ্যা;  // let isEmpty = false;
```

### Type Keywords

| Token | Bengali | English | Type |
|-------|---------|---------|------|
| `TYPE_BYTE` | বাইট | byte | 8-bit integer |
| `TYPE_SHORT` | ছোট_সংখ্যা | short | 16-bit integer |
| `TYPE_INT` | পূর্ণসংখ্যা | int | 32-bit integer |
| `TYPE_LONG` | দীর্ঘ_সংখ্যা | long | 64-bit integer |
| `TYPE_FLOAT` | দশমিক | float | 32-bit float |
| `TYPE_DOUBLE` | দশমিক_দ্বিগুণ | double | 64-bit float |
| `TYPE_CHAR` | অক্ষর | char | Unicode character |
| `TYPE_STRING` | পাঠ্য | string | String type |
| `TYPE_BOOLEAN` | বুলিয়ান | boolean | Boolean type |
| `TYPE_ARRAY` | তালিকা | array | Array type |
| `TYPE_HASH` | ম্যাপ | map | Hash/map type |

**Example:**
```bengali
ধরি x: পূর্ণসংখ্যা = ১০;              // let x: int = 10;
ধরি arr: তালিকা<পূর্ণসংখ্যা>;        // let arr: array<int>;
ধরি map: ম্যাপ<পাঠ্য, পূর্ণসংখ্যা>;  // let map: map<string, int>;
```

### Type Casting

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `AS` | হিসাবে | as | Type casting |

**Example:**
```bengali
ধরি x = ১০ হিসাবে দশমিক;  // let x = 10 as float;
```

### Struct and Enum Keywords

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `STRUCT` | স্ট্রাক্ট | struct | Struct definition |
| `ENUM` | গণনা | enum | Enum definition |

**Example:**
```bengali
ধরি person = স্ট্রাক্ট {
    নাম: "রহিম",
    বয়স: ২৫
};

ধরি color = গণনা {
    লাল,
    সবুজ,
    নীল
};
```

### OOP Keywords

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `CLASS` | শ্রেণী | class | Class definition |
| `METHOD` | পদ্ধতি | method | Method definition |
| `CONSTRUCTOR` | নির্মাতা | constructor | Constructor definition |
| `THIS` | এই | this | Current instance reference |
| `NEW` | নতুন | new | Create instance |
| `EXTENDS` | প্রসারিত | extends | Inheritance |
| `PUBLIC` | সার্বজনীন | public | Public access |
| `PRIVATE` | ব্যক্তিগত | private | Private access |
| `PROTECTED` | সুরক্ষিত | protected | Protected access |
| `STATIC` | স্থির | static | Static member |
| `ABSTRACT` | বিমূর্ত | abstract | Abstract class/method |
| `INTERFACE` | চুক্তি | interface | Interface definition |
| `IMPLEMENTS` | বাস্তবায়ন | implements | Interface implementation |
| `SUPER` | উর্ধ্ব | super | Parent class reference |
| `OVERRIDE` | পুনর্সংজ্ঞা | override | Override method |
| `FINAL` | চূড়ান্ত | final | Final class/method |

**Example:**
```bengali
শ্রেণী Person {
    সার্বজনীন নাম: পাঠ্য;
    ব্যক্তিগত বয়স: পূর্ণসংখ্যা;
    
    নির্মাতা(নাম: পাঠ্য, বয়স: পূর্ণসংখ্যা) {
        এই.নাম = নাম;
        এই.বয়স = বয়স;
    }
}

ধরি p = নতুন Person("রহিম", ২৫);
```

### Module System

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `IMPORT` | অন্তর্ভুক্ত | import | Import module |

**Example:**
```bengali
অন্তর্ভুক্ত "math.bhasa";
```

## Bengali Numeral Support

Bhasa supports both Arabic (0-9) and Bengali (০-৯) numerals natively.

### Bengali Digit Map

```go
var BengaliDigits = map[rune]rune{
    '০': '0',  // U+09E6 → U+0030
    '১': '1',  // U+09E7 → U+0031
    '২': '2',  // U+09E8 → U+0032
    '৩': '3',  // U+09E9 → U+0033
    '৪': '4',  // U+09EA → U+0034
    '৫': '5',  // U+09EB → U+0035
    '৬': '6',  // U+09EC → U+0036
    '৭': '7',  // U+09ED → U+0037
    '৮': '8',  // U+09EE → U+0038
    '৯': '9',  // U+09EF → U+0039
}
```

### Conversion Table

| Bengali | Arabic | Unicode |
|---------|--------|---------|
| ০ | 0 | U+09E6 |
| ১ | 1 | U+09E7 |
| ২ | 2 | U+09E8 |
| ৩ | 3 | U+09E9 |
| ৪ | 4 | U+09EA |
| ৫ | 5 | U+09EB |
| ৬ | 6 | U+09EC |
| ৭ | 7 | U+09ED |
| ৮ | 8 | U+09EE |
| ৯ | 9 | U+09EF |

## Utility Functions

### LookupIdent

Checks if an identifier is a reserved keyword.

```go
func LookupIdent(ident string) TokenType
```

**Parameters:**
- `ident`: The identifier string to check

**Returns:**
- The keyword's `TokenType` if it's a keyword
- `IDENT` if it's a regular identifier

**Example:**
```go
token.LookupIdent("ধরি")        // Returns token.LET
token.LookupIdent("ফাংশন")     // Returns token.FUNCTION
token.LookupIdent("myVariable") // Returns token.IDENT
```

### ConvertBengaliNumber

Converts Bengali numerals to Arabic numerals.

```go
func ConvertBengaliNumber(s string) string
```

**Parameters:**
- `s`: String containing Bengali numerals

**Returns:**
- String with Bengali numerals converted to Arabic

**Example:**
```go
token.ConvertBengaliNumber("১২৩")    // Returns "123"
token.ConvertBengaliNumber("৪৫৬")    // Returns "456"
token.ConvertBengaliNumber("১0৩")    // Returns "103" (mixed)
token.ConvertBengaliNumber("abc")    // Returns "abc" (no digits)
```

**How it works:**
```go
func ConvertBengaliNumber(s string) string {
    result := ""
    for _, ch := range s {
        if digit, ok := BengaliDigits[ch]; ok {
            result += string(digit)  // Convert Bengali digit
        } else {
            result += string(ch)     // Keep other characters
        }
    }
    return result
}
```

## Usage Examples

### Creating Tokens

```go
// Create a LET token
letToken := token.Token{
    Type:    token.LET,
    Literal: "ধরি",
    Line:    1,
    Column:  1,
}

// Create an identifier token
identToken := token.Token{
    Type:    token.IDENT,
    Literal: "myVar",
    Line:    1,
    Column:  5,
}

// Create an integer token
intToken := token.Token{
    Type:    token.INT,
    Literal: "123",
    Line:    1,
    Column:  13,
}
```

### Keyword Detection

```go
import "bhasa/token"

// Check if string is a keyword
literal := "ধরি"
tokenType := token.LookupIdent(literal)

if tokenType == token.LET {
    fmt.Println("It's a LET keyword")
} else {
    fmt.Println("It's an identifier")
}
```

### Bengali Number Conversion

```go
import "bhasa/token"

// Convert Bengali numerals in source code
source := "১২৩"
arabic := token.ConvertBengaliNumber(source)
fmt.Println(arabic)  // Output: "123"

// Parse the number
value, err := strconv.ParseInt(arabic, 10, 64)
if err == nil {
    fmt.Printf("Parsed value: %d\n", value)  // Output: 123
}
```

### Token Comparison

```go
// Compare token types
if tok.Type == token.LET {
    // Handle variable declaration
}

if tok.Type == token.FUNCTION {
    // Handle function definition
}

// Check for operators
isArithmetic := tok.Type == token.PLUS ||
                tok.Type == token.MINUS ||
                tok.Type == token.ASTERISK ||
                tok.Type == token.SLASH
```

### Working with Position Information

```go
// Create error message with position
func reportError(tok token.Token, msg string) {
    fmt.Printf("Error at line %d, column %d: %s\n",
        tok.Line, tok.Column, msg)
    fmt.Printf("Token: %s (%s)\n", tok.Type, tok.Literal)
}

// Example usage
tok := token.Token{
    Type:    token.ILLEGAL,
    Literal: "@",
    Line:    5,
    Column:  12,
}
reportError(tok, "unexpected character")
// Output:
// Error at line 5, column 12: unexpected character
// Token: ILLEGAL (@)
```

## Token Type Categories

### By Purpose

**Keywords:** Define language structure
```
LET, FUNCTION, IF, ELSE, RETURN, WHILE, FOR, etc.
```

**Operators:** Perform operations
```
PLUS, MINUS, EQ, AND, OR, BIT_AND, etc.
```

**Delimiters:** Structure code
```
LPAREN, RPAREN, LBRACE, RBRACE, SEMICOLON, etc.
```

**Literals:** Represent values
```
INT, STRING, TRUE, FALSE
```

**Identifiers:** Name variables/functions
```
IDENT
```

## Design Principles

### 1. **Bengali-First Design**
All keywords use meaningful Bengali words, not transliterations:
- `ধরি` (dhari - "let") instead of "লেট" (let)
- `যদি` (jodi - "if") instead of "ইফ" (if)

### 2. **Unicode-Native**
Full support for Bengali script (UTF-8):
- Keywords: Bengali words
- Numerals: Both Bengali and Arabic
- Identifiers: Mixed Bengali/English

### 3. **Position Tracking**
Every token carries position information:
- Essential for error reporting
- Enables IDE features
- Helps with debugging

### 4. **Type Safety**
Strong typing with `TokenType`:
- Compile-time checking
- Clear intent
- Easy to extend

## Summary

The **Token** package provides:

✅ **50+ Token Types**: Complete coverage of Bhasa syntax  
✅ **Bengali Keywords**: 30+ meaningful Bengali keywords  
✅ **Dual Numeral System**: Arabic and Bengali digits  
✅ **Position Tracking**: Line and column for every token  
✅ **Type Safety**: String-based token types  
✅ **Utility Functions**: Keyword lookup and numeral conversion  
✅ **OOP Support**: Complete class-based programming tokens  
✅ **Type System**: Rich type annotation tokens  

For implementation details and advanced usage, see [token-documentation.md](./token-documentation.md).

---

## Quick Reference

### Common Token Checks

```go
// Is it a keyword?
if token.LookupIdent(lit) != token.IDENT {
    // It's a keyword
}

// Is it an operator?
isOp := tok.Type == token.PLUS ||
        tok.Type == token.MINUS ||
        tok.Type == token.ASTERISK

// Is it a delimiter?
isDelim := tok.Type == token.LPAREN ||
           tok.Type == token.COMMA ||
           tok.Type == token.SEMICOLON

// Is it EOF?
if tok.Type == token.EOF {
    // End of input
}
```

### Bengali vs English

```go
// These are equivalent in Bhasa:
// Bengali: ধরি x = ৫;
// English concept: let x = 5;

// But the language uses Bengali keywords:
✓ ধরি x = ৫;        // Correct
✗ let x = 5;        // Invalid (unless 'let' is variable name)
```

The token package is the foundation of Bhasa's lexical analysis, providing a clean interface between the lexer and parser.

