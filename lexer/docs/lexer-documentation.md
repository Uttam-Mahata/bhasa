# Lexer Implementation Documentation

## Table of Contents

- [Overview](#overview)
- [Data Structures](#data-structures)
- [Initialization](#initialization)
- [Core Functions](#core-functions)
- [Character Reading](#character-reading)
- [Token Generation](#token-generation)
- [Identifier Processing](#identifier-processing)
- [Number Processing](#number-processing)
- [String Processing](#string-processing)
- [Whitespace and Comments](#whitespace-and-comments)
- [Unicode Handling](#unicode-handling)
- [Operator Recognition](#operator-recognition)
- [Position Tracking](#position-tracking)
- [Complete Examples](#complete-examples)
- [Testing Strategies](#testing-strategies)
- [Performance Considerations](#performance-considerations)

---

## Overview

The **Lexer** (lexical analyzer) is the first phase of the Bhasa language interpreter. It transforms a stream of characters (source code) into a stream of tokens that can be consumed by the parser.

### Lexical Analysis Process

```
Source Code → Lexer → Token Stream → Parser → AST
```

**Example:**
```
Input:  "ধরি x = ৫;"
        ↓
Lexer:  [LET, IDENT, ASSIGN, INT, SEMICOLON, EOF]
        ↓
Parser: AST (Let Statement Node)
```

### Design Principles

1. **Single Responsibility**: Only converts characters to tokens
2. **Single Pass**: Reads input once, left-to-right
3. **No Backtracking**: Peek-ahead for lookahead
4. **Unicode First**: Full Bengali script support
5. **Position Aware**: Track line/column for errors

---

## Data Structures

### Lexer Structure

```go
type Lexer struct {
    input        []rune  // Input as rune slice for Unicode support
    position     int     // Current position in input (points to current char)
    readPosition int     // Next reading position (after current char)
    ch           rune    // Current character under examination
    line         int     // Current line number (1-indexed)
    column       int     // Current column number (0-indexed initially)
}
```

#### Field Descriptions

| Field | Type | Purpose | Example State |
|-------|------|---------|---------------|
| `input` | `[]rune` | Source code as Unicode code points | `['ধ', 'র', 'ি', ' ', 'x']` |
| `position` | `int` | Index of current character | `0` |
| `readPosition` | `int` | Index of next character to read | `1` |
| `ch` | `rune` | Current character being examined | `'ধ'` |
| `line` | `int` | Current line (starts at 1) | `1` |
| `column` | `int` | Current column (increments with each char) | `1` |

#### Why []rune Instead of string?

**Problem with strings:**
```go
s := "ধরি"
len(s)     // Returns 9 (bytes), not 3 (characters)
s[0]       // Returns 224 (byte), not 'ধ' (character)
```

**Solution with []rune:**
```go
r := []rune("ধরি")
len(r)     // Returns 3 (characters)
r[0]       // Returns 'ধ' (Unicode code point U+09A7)
```

### Token Structure

```go
type Token struct {
    Type    TokenType  // Token type (LET, IDENT, INT, etc.)
    Literal string     // Actual text from source
    Line    int        // Line number where token starts
    Column  int        // Column number where token starts
}
```

**Example Token:**
```go
token.Token{
    Type:    token.LET,
    Literal: "ধরি",
    Line:    1,
    Column:  1,
}
```

---

## Initialization

### New() Constructor

```go
func New(input string) *Lexer {
    l := &Lexer{
        input:  []rune(input),  // Convert string to rune slice
        line:   1,              // Start at line 1
        column: 0,              // Column will be 1 after first readChar()
    }
    l.readChar()  // Initialize by reading first character
    return l
}
```

#### Initialization Process

**Step-by-step:**
```
1. Input: "ধরি x"
2. Convert to runes: ['ধ', 'র', 'ি', ' ', 'x']
3. Initial state:
   - position = 0
   - readPosition = 0
   - ch = 0 (null)
   - line = 1
   - column = 0
4. Call readChar():
   - ch = 'ধ'
   - position = 0
   - readPosition = 1
   - column = 1
5. Ready to tokenize!
```

**Why read first character?**
- Lexer is always "one character ahead"
- Makes NextToken() logic simpler
- Current character is always available in `l.ch`

---

## Core Functions

### NextToken() - Main Entry Point

```go
func (l *Lexer) NextToken() token.Token {
    var tok token.Token
    
    l.skipWhitespace()  // Ignore whitespace
    
    // Record position at start of token
    tokLine := l.line
    tokCol := l.column
    
    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            // Two-character operator: ==
            ch := l.ch
            l.readChar()
            tok = l.newTokenWithPos(token.EQ, string(ch)+string(l.ch))
        } else if l.peekChar() == '>' {
            // Two-character operator: =>
            ch := l.ch
            l.readChar()
            tok = l.newTokenWithPos(token.ARROW, string(ch)+string(l.ch))
        } else {
            // Single-character operator: =
            tok = l.newTokenWithPos(token.ASSIGN, string(l.ch))
        }
    // ... more cases ...
    default:
        if isLetter(l.ch) {
            // Identifier or keyword
            literal := l.readIdentifier()
            tok = token.Token{
                Type:    token.LookupIdent(literal),
                Literal: literal,
                Line:    tokLine,
                Column:  tokCol,
            }
            return tok  // Early return (readIdentifier advances position)
        } else if isDigit(l.ch) || isBengaliDigit(l.ch) {
            // Number literal
            literal := l.readNumber()
            tok = token.Token{
                Type:    token.INT,
                Literal: literal,
                Line:    tokLine,
                Column:  tokCol,
            }
            return tok  // Early return
        } else {
            tok = l.newTokenWithPos(token.ILLEGAL, string(l.ch))
        }
    }
    
    l.readChar()  // Advance to next character
    return tok
}
```

#### Token Recognition Strategy

**1. Skip Non-Tokens:**
- Whitespace (spaces, tabs, newlines)
- Comments (line comments starting with `//`)

**2. Record Position:**
- Save line and column before consuming characters
- Tokens are labeled with their starting position

**3. Classify Character:**
- **Single-char operators:** `+`, `-`, `*`, `(`, `)`, etc.
- **Multi-char operators:** `==`, `!=`, `<=`, `>=`, `&&`, `||`, `<<`, `>>`
- **Identifiers/Keywords:** Start with letter or underscore
- **Numbers:** Start with digit
- **Strings:** Start with `"`
- **EOF:** Null character (`0`)

**4. Handle Multi-Character Tokens:**
- Use `peekChar()` for lookahead
- Consume characters as needed
- Return early to avoid extra `readChar()`

---

## Character Reading

### readChar() - Consume Next Character

```go
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0  // EOF represented as null
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
    l.column++
    
    if l.ch == '\n' {
        l.line++
        l.column = 0
    }
}
```

#### Character Reading Process

**Example: Reading "ধরি"**
```
Initial State:
  input = ['ধ', 'র', 'ি']
  position = 0, readPosition = 1
  ch = 'ধ', line = 1, column = 1

Call readChar():
  ch = input[1] = 'র'
  position = 1, readPosition = 2
  column = 2

Call readChar():
  ch = input[2] = 'ি'
  position = 2, readPosition = 3
  column = 3

Call readChar():
  readPosition (3) >= len(input) (3)
  ch = 0 (EOF)
  position = 3, readPosition = 4
  column = 4
```

#### Newline Handling

```go
if l.ch == '\n' {
    l.line++     // Increment line number
    l.column = 0 // Reset column to 0 (will be 1 after next readChar())
}
```

**Example: Multi-line Input**
```
Input: "ধরি x = ৫;\nধরি y = ১০;"

After first line:
  line = 1, column varies

After '\n':
  line = 2, column = 0

After 'ধ' in second line:
  line = 2, column = 1
```

### peekChar() - Look Ahead Without Consuming

```go
func (l *Lexer) peekChar() rune {
    if l.readPosition >= len(l.input) {
        return 0  // EOF
    }
    return l.input[l.readPosition]
}
```

#### Use Cases for Peek

**1. Two-Character Operators:**
```go
// Current: '='
// Peek: '='
// Result: Token type EQ ("==")

if l.ch == '=' && l.peekChar() == '=' {
    // It's "=="
}
```

**2. Disambiguation:**
```go
// '<' can be:
//   - LT ("<")
//   - LTE ("<=")
//   - LSHIFT ("<<")

case '<':
    if l.peekChar() == '=' {
        // "<="
    } else if l.peekChar() == '<' {
        // "<<"
    } else {
        // "<"
    }
```

**3. Comment Detection:**
```go
// '/' can be:
//   - SLASH ("/")
//   - Comment ("//")

case '/':
    if l.peekChar() == '/' {
        l.skipComment()
        return l.NextToken()  // Get next token after comment
    } else {
        tok = l.newTokenWithPos(token.SLASH, string(l.ch))
    }
```

---

## Token Generation

### newTokenWithPos() - Create Token With Position

```go
func (l *Lexer) newTokenWithPos(tokenType token.TokenType, literal string) token.Token {
    return token.Token{
        Type:    tokenType,
        Literal: literal,
        Line:    l.line,
        Column:  l.column - len([]rune(literal)) + 1,
    }
}
```

#### Column Calculation

The column calculation accounts for the fact that we've already consumed the token's characters:

```
Current position: column = 5
Token literal: "==" (length 2)
Token started at: column - len("==") + 1 = 5 - 2 + 1 = 4
```

**Example:**
```
Input: "ধরি x == ৫"
       ^^^^^^^
       123456789

When at '=' (second one):
  l.column = 9
  literal = "=="
  Token column = 9 - 2 + 1 = 8 (correct start position)
```

**Why use []rune(literal) for length?**
- `len(literal)` returns byte count, not character count
- Bengali characters are multi-byte
- `len([]rune(literal))` returns character count

**Example:**
```go
literal := "ধরি"
len(literal)           // 9 bytes
len([]rune(literal))   // 3 characters
```

---

## Identifier Processing

### readIdentifier() - Read Variable/Keyword Name

```go
func (l *Lexer) readIdentifier() string {
    startPos := l.position
    // Read first character (must be letter or underscore)
    for isLetter(l.ch) || isDigit(l.ch) || isBengaliDigit(l.ch) {
        l.readChar()
    }
    return string(l.input[startPos:l.position])
}
```

#### Identifier Rules

**Valid identifiers:**
- Must start with: letter or underscore
- Can contain: letters, digits, underscores
- Bengali script fully supported

**Examples:**
```
✓ ফ্যাক্টোরিয়াল    (Bengali word)
✓ factorial      (English word)
✓ _private       (underscore prefix)
✓ var123         (letters + digits)
✓ ফাংশন_১        (Bengali + digit)
✓ নাম_name       (mixed Bengali + English)

✗ 123abc         (starts with digit)
✗ my-var         (hyphen not allowed)
✗ my.var         (dot not allowed)
```

#### Identifier Reading Process

**Example: Reading "ফ্যাক্টোরিয়াল"**
```
Initial: position = 0, ch = 'ফ'

Iteration 1: isLetter('ফ') = true, readChar(), ch = '্'
Iteration 2: isLetter('্') = true (vowel sign), readChar(), ch = 'য'
Iteration 3: isLetter('য') = true, readChar(), ch = 'া'
Iteration 4: isLetter('া') = true (vowel sign), readChar(), ch = 'ক'
Iteration 5: isLetter('ক') = true, readChar(), ch = '্'
Iteration 6: isLetter('্') = true, readChar(), ch = 'ট'
... and so on ...

Stop when: ch = ' ' (space, not a letter)
Return: input[0:current_position]
```

#### Keyword Detection

After reading an identifier, check if it's a keyword:

```go
literal := l.readIdentifier()
tok = token.Token{
    Type:    token.LookupIdent(literal),  // Returns keyword type or IDENT
    Literal: literal,
    Line:    tokLine,
    Column:  tokCol,
}
```

**LookupIdent() Logic:**
```go
func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok  // It's a keyword
    }
    return IDENT  // It's an identifier
}
```

**Example:**
```
"ধরি"        → LET (keyword)
"ফাংশন"      → FUNCTION (keyword)
"myVariable" → IDENT (identifier)
```

---

## Number Processing

### readNumber() - Read Integer Literal

```go
func (l *Lexer) readNumber() string {
    startPos := l.position
    for isDigit(l.ch) || isBengaliDigit(l.ch) {
        l.readChar()
    }
    result := string(l.input[startPos:l.position])
    // Convert Bengali digits to Arabic
    return token.ConvertBengaliNumber(result)
}
```

#### Number Reading Process

**Example 1: Arabic Digits "123"**
```
Initial: position = 0, ch = '1'

Iteration 1: isDigit('1') = true, readChar(), ch = '2'
Iteration 2: isDigit('2') = true, readChar(), ch = '3'
Iteration 3: isDigit('3') = true, readChar(), ch = ';'
Stop: isDigit(';') = false

Extract: input[0:3] = "123"
Convert: ConvertBengaliNumber("123") = "123"
Return: "123"
```

**Example 2: Bengali Digits "১২৩"**
```
Initial: position = 0, ch = '১'

Iteration 1: isBengaliDigit('১') = true, readChar(), ch = '২'
Iteration 2: isBengaliDigit('২') = true, readChar(), ch = '৩'
Iteration 3: isBengaliDigit('৩') = true, readChar(), ch = ';'
Stop: isBengaliDigit(';') = false

Extract: input[0:3] = "১২৩"
Convert: ConvertBengaliNumber("১২৩") = "123"
Return: "123"
```

**Example 3: Mixed Digits "১2৩" (unusual but supported)**
```
Initial: position = 0, ch = '১'

Read all digits: "১2৩"
Convert:
  '১' → '1'
  '2' → '2' (already Arabic)
  '৩' → '3'
Return: "123"
```

#### Bengali to Arabic Conversion

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

**Conversion Map:**
```go
var BengaliDigits = map[rune]rune{
    '০': '0', '১': '1', '২': '2', '৩': '3', '৪': '4',
    '৫': '5', '৬': '6', '৭': '7', '৮': '8', '৯': '9',
}
```

**Why Convert?**
- Internal representation uses integers
- Go's `strconv.Atoi()` expects Arabic digits
- Simpler for evaluation phase

---

## String Processing

### readString() - Read String Literal

```go
func (l *Lexer) readString() string {
    startPos := l.position + 1  // Skip opening quote
    for {
        l.readChar()
        if l.ch == '"' || l.ch == 0 {  // Closing quote or EOF
            break
        }
    }
    return string(l.input[startPos:l.position])
}
```

#### String Reading Process

**Example: Reading "হ্যালো"**
```
Input: "হ্যালো"
       ^      ^
       start  end

Initial: position = 0, ch = '"'
startPos = 0 + 1 = 1 (skip opening quote)

Iteration 1: readChar(), ch = 'হ'
Iteration 2: readChar(), ch = '্'
Iteration 3: readChar(), ch = 'য'
Iteration 4: readChar(), ch = 'া'
Iteration 5: readChar(), ch = 'ল'
Iteration 6: readChar(), ch = 'ো'
Iteration 7: readChar(), ch = '"'
Stop: ch == '"'

Extract: input[1:7] = "হ্যালো"
Return: "হ্যালো"
```

#### String Features

**1. Unicode Support:**
```
"হ্যালো"        ✓ Bengali
"Hello"         ✓ English
"مرحبا"         ✓ Arabic
"你好"           ✓ Chinese
"🔥🚀"           ✓ Emojis
```

**2. Empty Strings:**
```
""   → returns ""
```

**3. Unterminated Strings:**
```
"hello
(no closing quote)

Stops at: ch == 0 (EOF)
Returns partial: "hello\n" (including newline)
```

**4. No Escape Sequences (Currently):**
```
"line1\nline2"  → literal string, \n not interpreted
```

**Future Enhancement:**
Could add escape sequence handling:
```go
for {
    l.readChar()
    if l.ch == '\\' {
        l.readChar()  // Skip escape and next char
    } else if l.ch == '"' || l.ch == 0 {
        break
    }
}
```

---

## Whitespace and Comments

### skipWhitespace() - Ignore Whitespace

```go
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}
```

#### Whitespace Characters

- `' '` - Space
- `'\t'` - Tab
- `'\n'` - Newline (Line Feed)
- `'\r'` - Carriage Return

**Example:**
```
Input: "ধরি    x\t=\n৫;"
        ^^^^  ^^^ (whitespace)

After skipWhitespace():
  Tokens: LET, IDENT, ASSIGN, INT, SEMICOLON
  (Whitespace not in token stream)
```

### skipComment() - Ignore Line Comments

```go
func (l *Lexer) skipComment() {
    for l.ch != '\n' && l.ch != 0 {
        l.readChar()
    }
    l.skipWhitespace()
}
```

#### Comment Processing

**Example:**
```
Input: "ধরি x = ৫; // এটি একটি কমেন্ট\nধরি y = ১০;"

Position: 'ধ' 'র' 'ি' ' ' 'x' ' ' '=' ' ' '৫' ';' ' ' '/' '/'
                                                             ^
When at '/':
  peekChar() == '/' → it's a comment
  skipComment() reads until '\n' or EOF
  skipWhitespace() skips the newline
  NextToken() called recursively
  
Result: Comment ignored, continues with next line
```

**Why Call skipWhitespace() After Comment?**
- Comments end at newline
- Newline is whitespace
- Remove newline to continue cleanly

---

## Unicode Handling

### isLetter() - Check if Character is Letter

```go
func isLetter(ch rune) bool {
    // Check for letters, Bengali vowel signs (মাত্রা), and other combining marks
    return unicode.IsLetter(ch) || ch == '_' || isBengaliVowelSign(ch) || unicode.Is(unicode.Mn, ch)
}
```

#### Why So Complex?

**Problem: Bengali Script is Complex**

Bengali words are not just sequences of letters:

```
Word: কাজ (kaj - work)
  ক (k) - base letter
  া (a) - vowel sign (not a separate letter)
  জ (j) - base letter
```

If we only checked `unicode.IsLetter()`:
- `ক` → true (letter)
- `া` → false (not a letter, it's a combining mark)
- Result: Would split at vowel sign!

**Solution: Include Combining Marks**

```go
return unicode.IsLetter(ch) ||        // Regular letters
       ch == '_' ||                   // Underscore (for identifiers)
       isBengaliVowelSign(ch) ||      // Bengali vowel signs
       unicode.Is(unicode.Mn, ch)     // Nonspacing marks (diacritics)
```

### isBengaliVowelSign() - Check Bengali Vowel Sign

```go
func isBengaliVowelSign(ch rune) bool {
    // Bengali vowel signs (মাত্রা) and diacritics: U+0981 to U+09CD
    // This includes: ঁ ং ঃ, vowel signs, and hasant
    return (ch >= 0x0981 && ch <= 0x09CD) || ch == 0x09D7
}
```

#### Bengali Unicode Ranges

**U+0981 - U+09CD:**
```
U+0981   ঁ   Candrabindu
U+0982   ং   Anusvara
U+0983   ঃ   Visarga
U+0985-U+098C   Vowels (অ আ ই ঈ উ ঊ ঋ ৠ)
U+098F-U+0990   Vowels (এ ঐ)
U+0993-U+0994   Vowels (ও ঔ)
U+0995-U+09B9   Consonants (ক-হ)
U+09BC   ়   Nukta
U+09BE-U+09C4   Vowel signs (া ি ী ু ূ ৃ ে ৈ)
U+09C7-U+09C8   Vowel signs (ে ৈ)
U+09CB-U+09CC   Vowel signs (ো ৌ)
U+09CD   ্   Hasant (virama)
```

**U+09D7:**
```
U+09D7   ৗ   AU length mark
```

#### Example: Complex Bengali Word

```
Word: স্ট্রাক্ট (struct)

Character breakdown:
  স   U+09B8  - Base letter
  ্   U+09CD  - Hasant (conjunct former)
  ট   U+099F  - Base letter
  ্   U+09CD  - Hasant
  র   U+09B0  - Base letter
  া   U+09BE  - Vowel sign
  ক   U+0995  - Base letter
  ্   U+09CD  - Hasant
  ট   U+099F  - Base letter

All treated as single identifier by isLetter()
```

### isDigit() and isBengaliDigit()

```go
func isDigit(ch rune) bool {
    return '0' <= ch && ch <= '9'
}

func isBengaliDigit(ch rune) bool {
    return '০' <= ch && ch <= '৯'
}
```

#### Digit Ranges

**Arabic Numerals:** U+0030 - U+0039
```
0 (U+0030) to 9 (U+0039)
```

**Bengali Numerals:** U+09E6 - U+09EF
```
০ (U+09E6) to ৯ (U+09EF)
```

---

## Operator Recognition

### Single-Character Operators

Simple operators are recognized directly:

```go
case '+':
    tok = l.newTokenWithPos(token.PLUS, string(l.ch))
case '-':
    tok = l.newTokenWithPos(token.MINUS, string(l.ch))
case '*':
    tok = l.newTokenWithPos(token.ASTERISK, string(l.ch))
case '(':
    tok = l.newTokenWithPos(token.LPAREN, string(l.ch))
// ... etc
```

### Multi-Character Operators

Multi-character operators require lookahead:

#### Equality Operators

```go
case '=':
    if l.peekChar() == '=' {
        // "=="
        ch := l.ch
        l.readChar()  // Consume second '='
        tok = l.newTokenWithPos(token.EQ, string(ch)+string(l.ch))
    } else if l.peekChar() == '>' {
        // "=>"
        ch := l.ch
        l.readChar()  // Consume '>'
        tok = l.newTokenWithPos(token.ARROW, string(ch)+string(l.ch))
    } else {
        // "="
        tok = l.newTokenWithPos(token.ASSIGN, string(l.ch))
    }
```

**Decision Tree:**
```
Current: '='
  ├─ Peek: '=' → Token: EQ ("==")
  ├─ Peek: '>' → Token: ARROW ("=>")
  └─ Peek: other → Token: ASSIGN ("=")
```

#### Comparison Operators

```go
case '<':
    if l.peekChar() == '=' {
        // "<="
        ch := l.ch
        l.readChar()
        tok = l.newTokenWithPos(token.LTE, string(ch)+string(l.ch))
    } else if l.peekChar() == '<' {
        // "<<"
        ch := l.ch
        l.readChar()
        tok = l.newTokenWithPos(token.LSHIFT, string(ch)+string(l.ch))
    } else {
        // "<"
        tok = l.newTokenWithPos(token.LT, string(l.ch))
    }
```

**Decision Tree:**
```
Current: '<'
  ├─ Peek: '=' → Token: LTE ("<=")
  ├─ Peek: '<' → Token: LSHIFT ("<<")
  └─ Peek: other → Token: LT ("<")
```

#### Logical and Bitwise Operators

```go
case '&':
    if l.peekChar() == '&' {
        // "&&" (logical AND)
        ch := l.ch
        l.readChar()
        tok = l.newTokenWithPos(token.AND, string(ch)+string(l.ch))
    } else {
        // "&" (bitwise AND)
        tok = l.newTokenWithPos(token.BIT_AND, string(l.ch))
    }
```

**Disambiguation:**
```
'&'  → BIT_AND (bitwise)
'&&' → AND (logical)

'|'  → BIT_OR (bitwise)
'||' → OR (logical)
```

### Complete Operator Set

#### Arithmetic Operators
```
+   PLUS
-   MINUS
*   ASTERISK
/   SLASH
%   PERCENT
```

#### Comparison Operators
```
<   LT
>   GT
==  EQ
!=  NOT_EQ
<=  LTE
>=  GTE
```

#### Logical Operators
```
!   BANG (NOT)
&&  AND
||  OR
```

#### Bitwise Operators
```
&   BIT_AND
|   BIT_OR
^   BIT_XOR
~   BIT_NOT
<<  LSHIFT
>>  RSHIFT
```

#### Other Operators
```
=   ASSIGN
=>  ARROW (function arrow)
```

---

## Position Tracking

### Why Track Positions?

Position information enables:
1. **Error Messages**: Show where errors occur
2. **Debugging**: Help developers locate issues
3. **IDE Features**: Hover information, go-to-definition
4. **Syntax Highlighting**: Associate tokens with locations

### Position Fields

```go
type Token struct {
    Type    TokenType
    Literal string
    Line    int     // 1-indexed
    Column  int     // 1-indexed after first character
}
```

### Line Tracking

Lines increment on newline:

```go
if l.ch == '\n' {
    l.line++
    l.column = 0
}
```

**Example:**
```
Input:
Line 1: ধরি x = ৫;
Line 2: ধরি y = ১০;

After "৫;":  line = 1, column = 11
After '\n': line = 2, column = 0
After 'ধ':  line = 2, column = 1
```

### Column Tracking

Columns increment on each character:

```go
l.column++  // In readChar()
```

**Column Reset:**
- Starts at 0 (becomes 1 after first readChar())
- Resets to 0 on newline

### Token Position Recording

```go
// Record position BEFORE consuming token
tokLine := l.line
tokCol := l.column

// ... read token ...

tok = token.Token{
    Type:    tokenType,
    Literal: literal,
    Line:    tokLine,   // Start line
    Column:  tokCol,    // Start column
}
```

**Why record before?**
- After reading token, position is at next character
- Token should be labeled with its starting position

### Position Calculation for Multi-Character Tokens

For tokens created with `newTokenWithPos()`:

```go
Column: l.column - len([]rune(literal)) + 1
```

**Example:**
```
Current: "==" at position column = 8
Token started at: 8 - 2 + 1 = 7
```

### Error Reporting Example

```go
// In parser or evaluator
if err != nil {
    fmt.Printf("Error at line %d, column %d: %s\n",
        token.Line, token.Column, err.Error())
    
    // Show source line with error marker
    showErrorLine(source, token.Line, token.Column)
}

func showErrorLine(source string, line, col int) {
    lines := strings.Split(source, "\n")
    if line <= len(lines) {
        fmt.Println(lines[line-1])
        fmt.Println(strings.Repeat(" ", col-1) + "^")
    }
}
```

**Output:**
```
Error at line 3, column 12: unexpected token '!'
    ধরি x = ১০!
               ^
```

---

## Complete Examples

### Example 1: Simple Variable Declaration

**Input:**
```bengali
ধরি x = ৫;
```

**Tokenization Trace:**

```
Initial State:
  input = ['ধ', 'র', 'ি', ' ', 'x', ' ', '=', ' ', '৫', ';']
  position = 0, readPosition = 1
  ch = 'ধ', line = 1, column = 1

Call NextToken():
  skipWhitespace() - ch = 'ধ' (not whitespace, no change)
  tokLine = 1, tokCol = 1
  isLetter('ধ') = true
  readIdentifier():
    startPos = 0
    Loop: 'ধ', 'র', 'ি' (all letters)
    Stop at ' ' (space)
    position = 3, ch = ' '
    return "ধরি"
  LookupIdent("ধরি") = LET
  Token{Type: LET, Literal: "ধরি", Line: 1, Column: 1}

Call NextToken():
  skipWhitespace() - skip ' ', ch = 'x', position = 4, column = 5
  tokLine = 1, tokCol = 5
  isLetter('x') = true
  readIdentifier():
    startPos = 4
    Loop: 'x'
    Stop at ' '
    position = 5, ch = ' '
    return "x"
  LookupIdent("x") = IDENT
  Token{Type: IDENT, Literal: "x", Line: 1, Column: 5}

Call NextToken():
  skipWhitespace() - skip ' ', ch = '=', position = 6, column = 7
  tokLine = 1, tokCol = 7
  case '=':
    peekChar() = ' ' (not '=' or '>')
    tok = ASSIGN
    readChar() - ch = ' '
  Token{Type: ASSIGN, Literal: "=", Line: 1, Column: 7}

Call NextToken():
  skipWhitespace() - skip ' ', ch = '৫', position = 8, column = 9
  tokLine = 1, tokCol = 9
  isBengaliDigit('৫') = true
  readNumber():
    startPos = 8
    Loop: '৫'
    Stop at ';'
    position = 9, ch = ';'
    result = "৫"
    ConvertBengaliNumber("৫") = "5"
    return "5"
  Token{Type: INT, Literal: "5", Line: 1, Column: 9}

Call NextToken():
  skipWhitespace() - ch = ';' (not whitespace)
  tokLine = 1, tokCol = 10
  case ';':
    tok = SEMICOLON
    readChar() - ch = 0 (EOF)
  Token{Type: SEMICOLON, Literal: ";", Line: 1, Column: 10}

Call NextToken():
  skipWhitespace() - ch = 0 (EOF, no change)
  tokLine = 1, tokCol = 11
  case 0:
    tok = EOF
  Token{Type: EOF, Literal: "", Line: 1, Column: 11}
```

**Final Token Stream:**
```
[
  Token{LET, "ধরি", 1, 1},
  Token{IDENT, "x", 1, 5},
  Token{ASSIGN, "=", 1, 7},
  Token{INT, "5", 1, 9},
  Token{SEMICOLON, ";", 1, 10},
  Token{EOF, "", 1, 11}
]
```

### Example 2: Function Definition

**Input:**
```bengali
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};
```

**Token Stream:**
```
Token{LET, "ধরি", 1, 1}
Token{IDENT, "যোগ", 1, 5}
Token{ASSIGN, "=", 1, 9}
Token{FUNCTION, "ফাংশন", 1, 11}
Token{LPAREN, "(", 1, 17}
Token{IDENT, "a", 1, 18}
Token{COMMA, ",", 1, 19}
Token{IDENT, "b", 1, 21}
Token{RPAREN, ")", 1, 22}
Token{LBRACE, "{", 1, 24}
Token{RETURN, "ফেরত", 2, 5}
Token{IDENT, "a", 2, 10}
Token{PLUS, "+", 2, 12}
Token{IDENT, "b", 2, 14}
Token{SEMICOLON, ";", 2, 15}
Token{RBRACE, "}", 3, 1}
Token{SEMICOLON, ";", 3, 2}
Token{EOF, "", 3, 3}
```

### Example 3: Conditional with Comparison

**Input:**
```bengali
যদি (x >= ১০) {
    লেখ("বড়");
} নাহলে {
    লেখ("ছোট");
}
```

**Token Stream:**
```
Token{IF, "যদি", 1, 1}
Token{LPAREN, "(", 1, 5}
Token{IDENT, "x", 1, 6}
Token{GTE, ">=", 1, 8}
Token{INT, "10", 1, 11}
Token{RPAREN, ")", 1, 13}
Token{LBRACE, "{", 1, 15}
Token{IDENT, "লেখ", 2, 5}
Token{LPAREN, "(", 2, 8}
Token{STRING, "বড়", 2, 9}
Token{RPAREN, ")", 2, 13}
Token{SEMICOLON, ";", 2, 14}
Token{RBRACE, "}", 3, 1}
Token{ELSE, "নাহলে", 3, 3}
Token{LBRACE, "{", 3, 9}
Token{IDENT, "লেখ", 4, 5}
Token{LPAREN, "(", 4, 8}
Token{STRING, "ছোট", 4, 9}
Token{RPAREN, ")", 4, 13}
Token{SEMICOLON, ";", 4, 14}
Token{RBRACE, "}", 5, 1}
Token{EOF, "", 5, 2}
```

### Example 4: Comments and Bitwise Operations

**Input:**
```bengali
// বিটওয়াইজ অপারেশন
ধরি a = ৫ & ৩;   // AND
ধরি b = ৫ | ৩;   // OR
ধরি c = ৫ << ২;  // Left shift
```

**Token Stream:**
```
// First comment skipped
Token{LET, "ধরি", 2, 1}
Token{IDENT, "a", 2, 5}
Token{ASSIGN, "=", 2, 7}
Token{INT, "5", 2, 9}
Token{BIT_AND, "&", 2, 11}
Token{INT, "3", 2, 13}
Token{SEMICOLON, ";", 2, 14}
// Second comment skipped
Token{LET, "ধরি", 3, 1}
Token{IDENT, "b", 3, 5}
Token{ASSIGN, "=", 3, 7}
Token{INT, "5", 3, 9}
Token{BIT_OR, "|", 3, 11}
Token{INT, "3", 3, 13}
Token{SEMICOLON, ";", 3, 14}
// Third comment skipped
Token{LET, "ধরি", 4, 1}
Token{IDENT, "c", 4, 5}
Token{ASSIGN, "=", 4, 7}
Token{INT, "5", 4, 9}
Token{LSHIFT, "<<", 4, 11}
Token{INT, "2", 4, 14}
Token{SEMICOLON, ";", 4, 15}
Token{EOF, "", 4, 16}
```

---

## Testing Strategies

### Unit Tests

Test individual functions:

```go
func TestNew(t *testing.T) {
    input := "ধরি"
    l := lexer.New(input)
    
    if l.ch != 'ধ' {
        t.Errorf("Expected first char 'ধ', got %c", l.ch)
    }
    if l.line != 1 {
        t.Errorf("Expected line 1, got %d", l.line)
    }
}

func TestReadIdentifier(t *testing.T) {
    input := "ফ্যাক্টোরিয়াল = "
    l := lexer.New(input)
    
    ident := l.readIdentifier()
    expected := "ফ্যাক্টোরিয়াল"
    
    if ident != expected {
        t.Errorf("Expected %s, got %s", expected, ident)
    }
}
```

### Integration Tests

Test complete tokenization:

```go
func TestNextToken(t *testing.T) {
    input := `ধরি x = ৫;`
    
    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
        expectedLine    int
        expectedColumn  int
    }{
        {token.LET, "ধরি", 1, 1},
        {token.IDENT, "x", 1, 5},
        {token.ASSIGN, "=", 1, 7},
        {token.INT, "5", 1, 9},
        {token.SEMICOLON, ";", 1, 10},
        {token.EOF, "", 1, 11},
    }
    
    l := lexer.New(input)
    
    for i, tt := range tests {
        tok := l.NextToken()
        
        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - wrong type. expected=%q, got=%q",
                i, tt.expectedType, tok.Type)
        }
        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("tests[%d] - wrong literal. expected=%q, got=%q",
                i, tt.expectedLiteral, tok.Literal)
        }
        if tok.Line != tt.expectedLine {
            t.Fatalf("tests[%d] - wrong line. expected=%d, got=%d",
                i, tt.expectedLine, tok.Line)
        }
        if tok.Column != tt.expectedColumn {
            t.Fatalf("tests[%d] - wrong column. expected=%d, got=%d",
                i, tt.expectedColumn, tok.Column)
        }
    }
}
```

### Edge Cases to Test

1. **Empty Input:**
```go
input := ""
// Should return EOF immediately
```

2. **Only Whitespace:**
```go
input := "   \t\n  "
// Should skip all whitespace, return EOF
```

3. **Unterminated String:**
```go
input := `"hello`
// Should handle gracefully (stops at EOF)
```

4. **Long Identifiers:**
```go
input := strings.Repeat("a", 1000)
// Should handle without issues
```

5. **All Bengali Digits:**
```go
input := "০১২৩৪৫৬৭৮৯"
// Should convert correctly
```

6. **Mixed Scripts:**
```go
input := "variable_নাম"
// Should accept mixed English/Bengali
```

7. **Complex Conjuncts:**
```go
input := "স্ট্রাক্ট"
// Should handle as single identifier
```

8. **All Operators:**
```go
input := "+ - * / % < > <= >= == != && || & | ^ ~ << >> ! ="
// Should recognize each operator
```

---

## Performance Considerations

### Time Complexity

**Overall: O(n)** where n = input length
- Single pass through input
- Each character read once
- Constant time per character

**Per-Function:**
- `readChar()`: O(1)
- `peekChar()`: O(1)
- `readIdentifier()`: O(m) where m = identifier length
- `readNumber()`: O(m) where m = number length
- `readString()`: O(m) where m = string length
- `skipWhitespace()`: O(m) where m = whitespace count
- `NextToken()`: O(1) amortized

### Space Complexity

**O(n)** where n = input length
- `input` slice: O(n)
- Other fields: O(1)
- No additional data structures

### Optimizations

1. **Rune Slice:**
   - Convert once at initialization
   - Random access in O(1)
   - No repeated conversions

2. **Single Pass:**
   - No backtracking
   - No re-reading
   - Efficient for large files

3. **Peek Without Consume:**
   - Lookahead doesn't advance
   - Avoids buffering

4. **Minimal String Operations:**
   - Slice directly from rune array
   - Avoid repeated string concatenation

### Benchmarking

```go
func BenchmarkNextToken(b *testing.B) {
    input := `ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
        যদি (n <= ১) {
            ফেরত ১;
        }
        ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
    };`
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        l := lexer.New(input)
        for {
            tok := l.NextToken()
            if tok.Type == token.EOF {
                break
            }
        }
    }
}
```

**Expected Results:**
- Thousands of tokenizations per millisecond
- Linear scaling with input size
- Consistent performance across different character types

---

## Summary

The **Bhasa Lexer** is a robust, Unicode-aware lexical analyzer that:

### Key Features ✅
- **Unicode Support**: Full Bengali script support using runes
- **Dual Numeral Systems**: Arabic (0-9) and Bengali (০-৯) digits
- **Position Tracking**: Line and column numbers for error reporting
- **Peek-Ahead**: Lookahead for multi-character operators
- **Comment Support**: Line comments with `//`
- **Comprehensive Operators**: Arithmetic, logical, bitwise, comparison
- **Bengali Keywords**: 30+ keywords in Bengali

### Design Principles ✅
- **Single Pass**: O(n) time complexity
- **No Backtracking**: Efficient character consumption
- **Clear Separation**: Only tokenization, no parsing
- **Maintainable**: Well-structured, easy to extend

### Token Categories ✅
1. Keywords (Bengali): ধরি, ফাংশন, যদি, etc.
2. Identifiers: Variable and function names
3. Literals: Integers and strings
4. Operators: +, -, *, /, ==, !=, &&, ||, etc.
5. Delimiters: (), {}, [], ,, ;, :
6. Special: EOF, ILLEGAL

### Use Cases ✅
- **Parser Input**: Provides token stream for syntax analysis
- **Error Detection**: Identifies invalid characters
- **Position Tracking**: Enables precise error messages
- **Syntax Highlighting**: Provides token types for IDE features

The lexer forms the foundation of the Bhasa interpreter, converting raw text into a structured token stream ready for parsing and evaluation.

---

**Next Steps:**
- Parser uses token stream to build AST
- Evaluator interprets AST to execute code
- Compiler (optional) generates bytecode

For parser documentation, see `../parser/docs/`.

