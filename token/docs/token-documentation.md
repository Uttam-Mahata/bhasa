# Token Package - Complete Technical Documentation

## Table of Contents

- [Overview](#overview)
- [Token Structure Deep Dive](#token-structure-deep-dive)
- [Token Types Reference](#token-types-reference)
- [Keywords System](#keywords-system)
- [Bengali Numeral System](#bengali-numeral-system)
- [Implementation Details](#implementation-details)
- [Use Cases](#use-cases)
- [Best Practices](#best-practices)

---

## Overview

The token package is the interface between the lexer and parser. It defines:
1. Token types (what kind of token it is)
2. Token structure (data associated with each token)
3. Keyword mapping (Bengali words to token types)
4. Bengali numeral conversion (০-৯ to 0-9)

### Role in Compilation Pipeline

```
Source Code → [Lexer] → Token Stream → [Parser] → AST
                           ↑
                     Token Package
```

The token package provides:
- **Type definitions**: What tokens exist
- **Structure**: How tokens are represented
- **Utilities**: Helper functions for token processing

---

## Token Structure Deep Dive

### TokenType

```go
type TokenType string
```

**Why `string` instead of `int`?**

Advantages of string-based types:
1. **Readable debugging**: `"LET"` vs `23`
2. **Easy comparison**: `tok.Type == token.LET`
3. **Flexible**: Can add new types without conflicts
4. **Serializable**: Can easily convert to JSON/text

Disadvantages:
1. **Memory**: Strings use more memory than integers
2. **Performance**: String comparison slightly slower

**Design Decision:** Readability and maintainability outweigh minor performance cost.

### Token Structure

```go
type Token struct {
    Type    TokenType  // What kind of token
    Literal string     // The actual text
    Line    int        // Where it appears (line)
    Column  int        // Where it appears (column)
}
```

#### Field Details

**Type** (`TokenType`):
- Category of the token
- Used by parser to understand meaning
- Examples: `LET`, `IDENT`, `PLUS`, `INT`

**Literal** (`string`):
- Exact text from source code
- Preserved for error messages
- Used for identifiers and values
- Examples: `"ধরি"`, `"myVar"`, `"123"`

**Line** (`int`):
- Line number in source (1-indexed)
- Used for error reporting
- Updated by lexer on newlines

**Column** (`int`):
- Column number in source (1-indexed after first char)
- Pinpoints exact position
- Used for precise error messages

#### Memory Layout

```
Token (32 bytes on 64-bit system):
├─ Type:    16 bytes (string header: ptr + len)
├─ Literal: 16 bytes (string header: ptr + len)
├─ Line:    8 bytes  (int)
└─ Column:  8 bytes  (int)
Total: 48 bytes per token
```

**Note:** Actual string data stored separately on heap.

---

## Token Types Reference

### Special Tokens

#### ILLEGAL

```go
const ILLEGAL = "ILLEGAL"
```

**Purpose:** Represents unrecognized characters

**When created:**
- Lexer encounters invalid character
- Character not in Bhasa grammar

**Example:**
```go
// Source: ধরি x @ 5;
//                ^ invalid character
Token{Type: ILLEGAL, Literal: "@", Line: 1, Column: 8}
```

**Handling:**
```go
if tok.Type == token.ILLEGAL {
    return fmt.Errorf("illegal character '%s' at %d:%d",
        tok.Literal, tok.Line, tok.Column)
}
```

#### EOF

```go
const EOF = "EOF"
```

**Purpose:** Marks end of input

**When created:**
- Lexer reaches end of file
- No more characters to read

**Example:**
```go
Token{Type: EOF, Literal: "", Line: 10, Column: 25}
```

**Handling:**
```go
for tok.Type != token.EOF {
    // Process tokens
    tok = lexer.NextToken()
}
```

### Identifier and Literals

#### IDENT

```go
const IDENT = "IDENT"
```

**Purpose:** Variable and function names

**Valid identifiers:**
- Start with letter or underscore
- Contain letters, digits, underscores
- Bengali and English characters allowed

**Examples:**
```
myVariable     ✓
ফাংশন_নাম      ✓
_private       ✓
var123         ✓
মিশ্র_mixed    ✓
123abc         ✗ (starts with digit)
my-var         ✗ (hyphen not allowed)
```

#### INT

```go
const INT = "INT"
```

**Purpose:** Integer literals

**Supported formats:**
- Arabic: `0`, `123`, `999`
- Bengali: `০`, `১২৩`, `৯৯৯`
- Mixed: `১2৩` (converted to `123`)

**Note:** Lexer converts Bengali to Arabic internally.

#### STRING

```go
const STRING = "STRING"
```

**Purpose:** String literals

**Format:** Enclosed in double quotes

**Examples:**
```bengali
"hello"          ✓
"হ্যালো"         ✓
"mixed মিশ্র"     ✓
""               ✓ (empty string)
```

**Current limitations:**
- No escape sequences (yet)
- Only double quotes supported
- No multiline strings

### Operators

#### Arithmetic Operators

```go
PLUS     "+"   // Addition: a + b
MINUS    "-"   // Subtraction: a - b
ASTERISK "*"   // Multiplication: a * b
SLASH    "/"   // Division: a / b
PERCENT  "%"   // Modulo: a % b
```

**Precedence:** Handled by parser, not token

**Example tokens:**
```bengali
// Source: ৫ + ৩ * ২
Token{INT, "5", 1, 1}
Token{PLUS, "+", 1, 3}
Token{INT, "3", 1, 5}
Token{ASTERISK, "*", 1, 7}
Token{INT, "2", 1, 9}
```

#### Comparison Operators

```go
LT     "<"    // Less than
GT     ">"    // Greater than
EQ     "=="   // Equal
NOT_EQ "!="   // Not equal
LTE    "<="   // Less than or equal
GTE    ">="   // Greater than or equal
```

**Two-character operators:**
- Lexer uses peek-ahead to recognize
- `=` vs `==`: Check next character
- `<` vs `<=` vs `<<`: Similar logic

**Example:**
```bengali
// Source: x >= 10
Token{IDENT, "x", 1, 1}
Token{GTE, ">=", 1, 3}
Token{INT, "10", 1, 6}
```

#### Logical Operators

```go
BANG "!"    // NOT: !condition
AND  "&&"   // AND: a && b
OR   "||"   // OR: a || b
```

**Truth table:**
```
NOT (!):
  !true  = false
  !false = true

AND (&&):
  true  && true  = true
  true  && false = false
  false && false = false

OR (||):
  true  || true  = true
  true  || false = true
  false || false = false
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

**Distinction from logical:**
- `&` vs `&&`: Bitwise vs logical AND
- `|` vs `||`: Bitwise vs logical OR

**Examples:**
```bengali
৫ & ৩     // 5 & 3 = 1 (0101 & 0011 = 0001)
৫ | ৩     // 5 | 3 = 7 (0101 | 0011 = 0111)
৫ ^ ৩     // 5 ^ 3 = 6 (0101 ^ 0011 = 0110)
~৫        // ~5 = -6 (NOT 0101 = 1010 in two's complement)
৫ << ১    // 5 << 1 = 10 (0101 << 1 = 1010)
৫ >> ১    // 5 >> 1 = 2 (0101 >> 1 = 0010)
```

### Delimiters

#### Grouping Delimiters

```go
LPAREN   "("    // Function calls, grouping
RPAREN   ")"    // Close parenthesis
LBRACE   "{"    // Blocks, hashes, structs
RBRACE   "}"    // Close brace
LBRACKET "["    // Array literals, indexing
RBRACKET "]"    // Close bracket
```

**Usage:**
```bengali
(x + y)         // Grouping
func(a, b)      // Function call
{x: 1, y: 2}    // Hash literal
[1, 2, 3]       // Array literal
arr[0]          // Index access
```

#### Separator Delimiters

```go
COMMA     ","   // Separate items in lists
SEMICOLON ";"   // End statements
COLON     ":"   // Type annotations, hash pairs
```

**Usage:**
```bengali
func(a, b, c)          // COMMA: separate parameters
ধরি x = 5;             // SEMICOLON: end statement
ধরি x: পূর্ণসংখ্যা;     // COLON: type annotation
{name: "রহিম"}         // COLON: key-value pair
```

#### Special Delimiters

```go
DOT   "."     // Member access
ARROW "=>"    // Function arrow (future use)
```

**Usage:**
```bengali
obj.field          // DOT: access field
arr.length         // DOT: access property
x => x * 2         // ARROW: lambda (potential future syntax)
```

---

## Keywords System

### Keyword Map

```go
var keywords = map[string]TokenType{
    "ধরি":         LET,
    "ফাংশন":       FUNCTION,
    "যদি":         IF,
    // ... more keywords
}
```

**Implementation:**
- Hash map for O(1) lookup
- String keys (Bengali words)
- TokenType values

### LookupIdent Function

```go
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok  // It's a keyword
    }
    return IDENT    // It's an identifier
}
```

**Flow:**
```
Input: "ধরি"
  ↓
Check keywords map
  ↓
Found: "ধরি" → LET
  ↓
Return: LET

Input: "myVariable"
  ↓
Check keywords map
  ↓
Not found
  ↓
Return: IDENT
```

**Usage in Lexer:**
```go
// After reading identifier
literal := l.readIdentifier()
tok := token.Token{
    Type:    token.LookupIdent(literal),  // Determine if keyword
    Literal: literal,
    Line:    l.line,
    Column:  l.column,
}
```

### Why Bengali Keywords?

**Design Philosophy:**
1. **Cultural Identity**: Bangla is the national language
2. **Accessibility**: Native speakers can learn faster
3. **Uniqueness**: Distinguishes Bhasa from other languages
4. **Meaning Over Transliteration**: Uses actual Bengali words

**Examples:**
```
English → Poor Transliteration → Good Bengali
let     → লেট               → ধরি (dhari - to hold/let)
if      → ইফ                → যদি (jodi - if)
return  → রিটার্ন            → ফেরত (ferot - return/give back)
```

### Keyword Categories

**Control Flow:**
```go
ধরি যদি নাহলে ফেরত যতক্ষণ পর্যন্ত বিরতি চালিয়ে_যাও
```

**Functions:**
```go
ফাংশন
```

**Boolean:**
```go
সত্য মিথ্যা
```

**Type System:**
```go
বাইট ছোট_সংখ্যা পূর্ণসংখ্যা দীর্ঘ_সংখ্যা দশমিক দশমিক_দ্বিগুণ
অক্ষর পাঠ্য বুলিয়ান তালিকা ম্যাপ
```

**OOP:**
```go
শ্রেণী পদ্ধতি নির্মাতা এই নতুন প্রসারিত
সার্বজনীন ব্যক্তিগত সুরক্ষিত স্থির বিমূর্ত
চুক্তি বাস্তবায়ন উর্ধ্ব পুনর্সংজ্ঞা চূড়ান্ত
```

---

## Bengali Numeral System

### Unicode Ranges

**Bengali Digits:** U+09E6 to U+09EF

```
০ U+09E6  BENGALI DIGIT ZERO
১ U+09E7  BENGALI DIGIT ONE
২ U+09E8  BENGALI DIGIT TWO
৩ U+09E9  BENGALI DIGIT THREE
৪ U+09EA  BENGALI DIGIT FOUR
৫ U+09EB  BENGALI DIGIT FIVE
৬ U+09EC  BENGALI DIGIT SIX
৭ U+09ED  BENGALI DIGIT SEVEN
৮ U+09EE  BENGALI DIGIT EIGHT
৯ U+09EF  BENGALI DIGIT NINE
```

### BengaliDigits Map

```go
var BengaliDigits = map[rune]rune{
    '০': '0',
    '১': '1',
    '২': '2',
    '৩': '3',
    '৪': '4',
    '৫': '5',
    '৬': '6',
    '৭': '7',
    '৮': '8',
    '৯': '9',
}
```

**Type:** `map[rune]rune`
- Key: Bengali digit (rune)
- Value: Arabic digit (rune)

**Why rune?**
- Unicode code points
- Bengali characters are multi-byte
- Rune represents single character

### ConvertBengaliNumber Function

```go
func ConvertBengaliNumber(s string) string {
    result := ""
    for _, ch := range s {
        if digit, ok := BengaliDigits[ch]; ok {
            result += string(digit)
        } else {
            result += string(ch)
        }
    }
    return result
}
```

**Algorithm:**
1. Iterate through each rune in input
2. Check if rune is Bengali digit
3. If yes: convert to Arabic
4. If no: keep as-is
5. Build result string

**Examples:**

Example 1: Pure Bengali
```go
Input:  "১২৩"
Runes:  ['১', '২', '৩']
Step 1: '১' → found → '1'
Step 2: '২' → found → '2'
Step 3: '৩' → found → '3'
Output: "123"
```

Example 2: Mixed
```go
Input:  "১0৩"
Runes:  ['১', '0', '৩']
Step 1: '১' → found → '1'
Step 2: '0' → not found → '0'
Step 3: '৩' → found → '3'
Output: "103"
```

Example 3: With non-digits
```go
Input:  "১২৩abc"
Runes:  ['১', '২', '৩', 'a', 'b', 'c']
Step 1: '১' → found → '1'
Step 2: '২' → found → '2'
Step 3: '৩' → found → '3'
Step 4: 'a' → not found → 'a'
Step 5: 'b' → not found → 'b'
Step 6: 'c' → not found → 'c'
Output: "123abc"
```

### Performance Considerations

**Time Complexity:** O(n) where n = string length
- Single pass through string
- Map lookup is O(1)

**Space Complexity:** O(n)
- New string allocated
- Size proportional to input

**Optimization Opportunity:**
Could use `strings.Builder` for better performance:
```go
func ConvertBengaliNumber(s string) string {
    var result strings.Builder
    result.Grow(len(s))  // Pre-allocate
    for _, ch := range s {
        if digit, ok := BengaliDigits[ch]; ok {
            result.WriteRune(digit)
        } else {
            result.WriteRune(ch)
        }
    }
    return result.String()
}
```

---

## Implementation Details

### Token Creation Patterns

#### Pattern 1: Simple Token

```go
tok := token.Token{
    Type:    token.PLUS,
    Literal: "+",
    Line:    1,
    Column:  5,
}
```

#### Pattern 2: With Keyword Lookup

```go
literal := "ধরি"
tok := token.Token{
    Type:    token.LookupIdent(literal),
    Literal: literal,
    Line:    1,
    Column:  1,
}
```

#### Pattern 3: With Number Conversion

```go
bengaliNum := "১২৩"
arabicNum := token.ConvertBengaliNumber(bengaliNum)
tok := token.Token{
    Type:    token.INT,
    Literal: arabicNum,  // "123"
    Line:    1,
    Column:  5,
}
```

### Token Comparison

#### By Type

```go
if tok.Type == token.LET {
    // Handle let statement
}
```

#### By Category

```go
isOperator := tok.Type == token.PLUS ||
              tok.Type == token.MINUS ||
              tok.Type == token.ASTERISK ||
              tok.Type == token.SLASH

isComparison := tok.Type == token.LT ||
                tok.Type == token.GT ||
                tok.Type == token.EQ ||
                tok.Type == token.NOT_EQ
```

#### Helper Functions (Not in package, but useful)

```go
func isKeyword(tok token.Token) bool {
    return token.LookupIdent(tok.Literal) != token.IDENT
}

func isLiteral(tok token.Token) bool {
    return tok.Type == token.INT ||
           tok.Type == token.STRING ||
           tok.Type == token.TRUE ||
           tok.Type == token.FALSE
}

func isDelimiter(tok token.Token) bool {
    delims := []token.TokenType{
        token.COMMA, token.SEMICOLON, token.COLON,
        token.LPAREN, token.RPAREN,
        token.LBRACE, token.RBRACE,
        token.LBRACKET, token.RBRACKET,
    }
    for _, d := range delims {
        if tok.Type == d {
            return true
        }
    }
    return false
}
```

---

## Use Cases

### Use Case 1: Lexer Token Generation

```go
func (l *Lexer) NextToken() token.Token {
    // Skip whitespace
    l.skipWhitespace()
    
    // Record position
    tokLine := l.line
    tokCol := l.column
    
    // Determine token type
    switch l.ch {
    case '+':
        return token.Token{
            Type:    token.PLUS,
            Literal: string(l.ch),
            Line:    tokLine,
            Column:  tokCol,
        }
    
    default:
        if isLetter(l.ch) {
            lit := l.readIdentifier()
            return token.Token{
                Type:    token.LookupIdent(lit),
                Literal: lit,
                Line:    tokLine,
                Column:  tokCol,
            }
        }
    }
}
```

### Use Case 2: Parser Token Consumption

```go
func (p *Parser) parseLetStatement() *ast.LetStatement {
    // Current token should be LET
    if p.curToken.Type != token.LET {
        return nil
    }
    
    stmt := &ast.LetStatement{Token: p.curToken}
    
    // Next token should be IDENT
    if !p.expectPeek(token.IDENT) {
        return nil
    }
    
    stmt.Name = &ast.Identifier{
        Token: p.curToken,
        Value: p.curToken.Literal,
    }
    
    // Expect ASSIGN
    if !p.expectPeek(token.ASSIGN) {
        return nil
    }
    
    // Parse value...
    return stmt
}
```

### Use Case 3: Error Reporting

```go
func (p *Parser) peekError(expected token.TokenType) {
    msg := fmt.Sprintf(
        "[Line %d, Col %d] expected %s, got %s instead",
        p.peekToken.Line,
        p.peekToken.Column,
        expected,
        p.peekToken.Type,
    )
    p.errors = append(p.errors, msg)
}
```

### Use Case 4: Number Parsing with Bengali Support

```go
func parseInt(tok token.Token) (int64, error) {
    // Token literal might be Bengali
    arabicNum := token.ConvertBengaliNumber(tok.Literal)
    
    // Parse Arabic numeral
    value, err := strconv.ParseInt(arabicNum, 10, 64)
    if err != nil {
        return 0, fmt.Errorf(
            "cannot parse '%s' as integer at %d:%d",
            tok.Literal, tok.Line, tok.Column,
        )
    }
    
    return value, nil
}
```

---

## Best Practices

### 1. Always Include Position Information

```go
// ✓ Good
tok := token.Token{
    Type:    token.IDENT,
    Literal: "x",
    Line:    5,
    Column:  12,
}

// ✗ Bad
tok := token.Token{
    Type:    token.IDENT,
    Literal: "x",
    // Missing line and column
}
```

**Why:** Position info is crucial for error reporting and debugging.

### 2. Use LookupIdent for All Identifiers

```go
// ✓ Good
literal := l.readIdentifier()
tokType := token.LookupIdent(literal)

// ✗ Bad
literal := l.readIdentifier()
tokType := token.IDENT  // Might be a keyword!
```

**Why:** Ensures keywords are properly recognized.

### 3. Convert Bengali Numbers Early

```go
// ✓ Good (in lexer)
bengaliNum := l.readNumber()
arabicNum := token.ConvertBengaliNumber(bengaliNum)
return token.Token{Type: token.INT, Literal: arabicNum}

// ✗ Bad (convert later)
bengaliNum := l.readNumber()
return token.Token{Type: token.INT, Literal: bengaliNum}
// Parser has to deal with Bengali digits
```

**Why:** Centralize conversion, simplify downstream processing.

### 4. Check Token Type Before Accessing Literal

```go
// ✓ Good
if tok.Type == token.INT {
    value, _ := strconv.ParseInt(tok.Literal, 10, 64)
    // Safe to parse
}

// ✗ Bad
value, err := strconv.ParseInt(tok.Literal, 10, 64)
// Might fail if tok is not INT
```

**Why:** Type safety, avoid runtime errors.

### 5. Use Constants, Not Strings

```go
// ✓ Good
if tok.Type == token.LET {
    // ...
}

// ✗ Bad
if tok.Type == "ধরি" {  // Wrong! token.Type is not "ধরি"
    // ...
}
```

**Why:** Type constants are guaranteed correct.

---

## Summary

The **token** package provides the foundational types for Bhasa's lexical structure:

### Key Components
- **Token**: Structure with type, literal, line, column
- **TokenType**: String-based type identifiers
- **Keywords**: Map of Bengali words to token types
- **Utilities**: LookupIdent, ConvertBengaliNumber

### Design Highlights
✅ Bengali-first keyword system  
✅ Dual numeral support (Bengali & Arabic)  
✅ Position tracking for every token  
✅ String-based types for readability  
✅ Simple, extensible design  

### Usage Pattern
```
Source Code → Lexer (uses token package) → Token Stream → Parser
```

The token package is lean, focused, and essential—providing exactly what's needed for lexical analysis without unnecessary complexity.

