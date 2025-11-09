# Lexer Quick Reference Guide

## Function Reference

### Core Functions

#### New(input string) *Lexer
Creates a new lexer instance.

```go
l := lexer.New("ধরি x = ৫;")
```

#### NextToken() token.Token
Returns the next token from input.

```go
tok := l.NextToken()
// tok = Token{Type: LET, Literal: "ধরি", Line: 1, Column: 1}
```

### Character Reading

#### readChar()
Consumes the next character.
- Advances `position` and `readPosition`
- Updates `ch` with next character
- Increments `column`
- Handles newlines (increments `line`, resets `column`)

#### peekChar() rune
Looks at next character without consuming it.
- Returns next character
- Does not advance position
- Returns `0` if at EOF

### Token Reading

#### readIdentifier() string
Reads an identifier (variable name or keyword).
- Starts with letter or underscore
- Can contain letters, digits, underscores
- Supports Bengali characters

```go
// Input: "ফ্যাক্টোরিয়াল = "
literal := l.readIdentifier()
// Returns: "ফ্যাক্টোরিয়াল"
```

#### readNumber() string
Reads a number literal.
- Accepts Arabic (0-9) and Bengali (০-৯) digits
- Converts Bengali to Arabic automatically
- Returns Arabic numeral string

```go
// Input: "১২৩"
literal := l.readNumber()
// Returns: "123"
```

#### readString() string
Reads a string literal.
- Starts and ends with `"`
- Returns content without quotes
- Stops at closing `"` or EOF

```go
// Input: "হ্যালো"
literal := l.readString()
// Returns: "হ্যালো"
```

### Whitespace and Comments

#### skipWhitespace()
Skips space, tab, newline, carriage return.

```go
// Input: "   \t\n  x"
l.skipWhitespace()
// Now at: 'x'
```

#### skipComment()
Skips line comment (from `//` to end of line).

```go
// Input: "// comment\nx"
l.skipComment()
// Now at: 'x'
```

### Helper Functions

#### isLetter(ch rune) bool
Checks if character can be part of an identifier.
- Letters (Unicode)
- Underscore `_`
- Bengali vowel signs
- Combining marks

#### isDigit(ch rune) bool
Checks if character is Arabic digit (0-9).

#### isBengaliDigit(ch rune) bool
Checks if character is Bengali digit (০-৯).

#### isBengaliVowelSign(ch rune) bool
Checks if character is Bengali vowel sign or diacritic.
- Range: U+0981 to U+09CD
- Also: U+09D7

#### newTokenWithPos(tokenType, literal) token.Token
Creates token with position information.

```go
tok := l.newTokenWithPos(token.PLUS, "+")
// Returns: Token{Type: PLUS, Literal: "+", Line: l.line, Column: calculated}
```

---

## Token Types Reference

### Keywords (Bengali)

| Token | Bengali | English | Usage |
|-------|---------|---------|-------|
| `LET` | ধরি | let | Variable declaration |
| `FUNCTION` | ফাংশন | function | Function definition |
| `IF` | যদি | if | Conditional |
| `ELSE` | নাহলে | else | Else clause |
| `RETURN` | ফেরত | return | Return statement |
| `TRUE` | সত্য | true | Boolean true |
| `FALSE` | মিথ্যা | false | Boolean false |
| `WHILE` | যতক্ষণ | while | While loop |
| `FOR` | পর্যন্ত | for | For loop |
| `BREAK` | বিরতি | break | Break statement |
| `CONTINUE` | চালিয়ে_যাও | continue | Continue statement |
| `IMPORT` | অন্তর্ভুক্ত | import | Import statement |

### Type Keywords

| Token | Bengali | English |
|-------|---------|---------|
| `TYPE_BYTE` | বাইট | byte |
| `TYPE_SHORT` | ছোট_সংখ্যা | short |
| `TYPE_INT` | পূর্ণসংখ্যা | int |
| `TYPE_LONG` | দীর্ঘ_সংখ্যা | long |
| `TYPE_FLOAT` | দশমিক | float |
| `TYPE_DOUBLE` | দশমিক_দ্বিগুণ | double |
| `TYPE_CHAR` | অক্ষর | char |
| `TYPE_STRING` | পাঠ্য | string |
| `TYPE_BOOLEAN` | বুলিয়ান | boolean |
| `TYPE_ARRAY` | তালিকা | array |
| `TYPE_HASH` | ম্যাপ | map |

### OOP Keywords

| Token | Bengali | English |
|-------|---------|---------|
| `CLASS` | শ্রেণী | class |
| `METHOD` | পদ্ধতি | method |
| `CONSTRUCTOR` | নির্মাতা | constructor |
| `THIS` | এই | this |
| `NEW` | নতুন | new |
| `EXTENDS` | প্রসারিত | extends |
| `PUBLIC` | সার্বজনীন | public |
| `PRIVATE` | ব্যক্তিগত | private |
| `PROTECTED` | সুরক্ষিত | protected |
| `STATIC` | স্থির | static |
| `ABSTRACT` | বিমূর্ত | abstract |
| `INTERFACE` | চুক্তি | interface |
| `IMPLEMENTS` | বাস্তবায়ন | implements |
| `SUPER` | উর্ধ্ব | super |
| `OVERRIDE` | পুনর্সংজ্ঞা | override |
| `FINAL` | চূড়ান্ত | final |

### Operators

#### Arithmetic
```
+    PLUS        Addition
-    MINUS       Subtraction
*    ASTERISK    Multiplication
/    SLASH       Division
%    PERCENT     Modulo
```

#### Comparison
```
<    LT          Less than
>    GT          Greater than
==   EQ          Equal to
!=   NOT_EQ      Not equal to
<=   LTE         Less than or equal
>=   GTE         Greater than or equal
```

#### Logical
```
!    BANG        Logical NOT
&&   AND         Logical AND
||   OR          Logical OR
```

#### Bitwise
```
&    BIT_AND     Bitwise AND
|    BIT_OR      Bitwise OR
^    BIT_XOR     Bitwise XOR
~    BIT_NOT     Bitwise NOT
<<   LSHIFT      Left shift
>>   RSHIFT      Right shift
```

#### Other
```
=    ASSIGN      Assignment
=>   ARROW       Function arrow
```

### Delimiters

```
,    COMMA       Separator
;    SEMICOLON   Statement terminator
:    COLON       Key-value separator
.    DOT         Member access
(    LPAREN      Left parenthesis
)    RPAREN      Right parenthesis
{    LBRACE      Left brace
}    RBRACE      Right brace
[    LBRACKET    Left bracket
]    RBRACKET    Right bracket
```

### Literals

```
IDENT    Identifier (variable/function name)
INT      Integer literal
STRING   String literal
```

### Special

```
EOF      End of file
ILLEGAL  Illegal/unrecognized character
```

---

## Usage Patterns

### Basic Tokenization Loop

```go
l := lexer.New(input)

for {
    tok := l.NextToken()
    
    // Process token
    fmt.Printf("%s\n", tok.Type)
    
    if tok.Type == token.EOF {
        break
    }
}
```

### Collecting All Tokens

```go
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

### Error Handling

```go
tok := l.NextToken()

if tok.Type == token.ILLEGAL {
    fmt.Printf("Illegal character '%s' at line %d, column %d\n",
        tok.Literal, tok.Line, tok.Column)
}
```

---

## Character Classification

### Letter Characters
- Unicode letters (Bengali, English, etc.)
- Underscore `_`
- Bengali vowel signs (U+0981-U+09CD, U+09D7)
- Nonspacing marks (Unicode category Mn)

### Digit Characters
- Arabic: `0` to `9` (U+0030-U+0039)
- Bengali: `০` to `৯` (U+09E6-U+09EF)

### Whitespace Characters
- Space: ` `
- Tab: `\t`
- Newline: `\n`
- Carriage return: `\r`

---

## Bengali Digit Conversion

| Bengali | Arabic | Hex |
|---------|--------|-----|
| ০ | 0 | U+09E6 → U+0030 |
| ১ | 1 | U+09E7 → U+0031 |
| ২ | 2 | U+09E8 → U+0032 |
| ৩ | 3 | U+09E9 → U+0033 |
| ৪ | 4 | U+09EA → U+0034 |
| ৫ | 5 | U+09EB → U+0035 |
| ৬ | 6 | U+09EC → U+0036 |
| ৭ | 7 | U+09ED → U+0037 |
| ৮ | 8 | U+09EE → U+0038 |
| ৯ | 9 | U+09EF → U+0039 |

---

## Common Token Sequences

### Variable Declaration
```
Input:  ধরি x = ৫;
Tokens: LET IDENT ASSIGN INT SEMICOLON EOF
```

### Function Definition
```
Input:  ফাংশন(x) { ফেরত x; }
Tokens: FUNCTION LPAREN IDENT RPAREN LBRACE RETURN IDENT SEMICOLON RBRACE EOF
```

### If Statement
```
Input:  যদি (x > ০) { }
Tokens: IF LPAREN IDENT GT INT RPAREN LBRACE RBRACE EOF
```

### While Loop
```
Input:  যতক্ষণ (x < ১০) { }
Tokens: WHILE LPAREN IDENT LT INT RPAREN LBRACE RBRACE EOF
```

### Function Call
```
Input:  লেখ("হ্যালো");
Tokens: IDENT LPAREN STRING RPAREN SEMICOLON EOF
```

### Array Literal
```
Input:  [১, ২, ৩]
Tokens: LBRACKET INT COMMA INT COMMA INT RBRACKET EOF
```

### Binary Expression
```
Input:  x + y * z
Tokens: IDENT PLUS IDENT ASTERISK IDENT EOF
```

---

## Position Tracking

### Line Numbering
- Starts at: 1
- Increments: On newline (`\n`)
- Never resets

### Column Numbering
- Starts at: 0 (becomes 1 after first `readChar()`)
- Increments: On each character
- Resets: On newline (to 0)

### Token Position
Each token records:
- `Line`: Line where token starts
- `Column`: Column where token starts

### Example
```
Line 1: ধরি x = ৫;
        ^   ^   ^  ^
Col:    1   5   9  10

Line 2: লেখ("test");
        ^   ^^    ^
Col:    1   45    11
```

---

## Error Detection

### Illegal Characters

Characters not recognized by lexer:

```go
Token{Type: ILLEGAL, Literal: "?", Line: 1, Column: 5}
```

**Common cases:**
- `@` (not used in Bhasa)
- `#` (not used in Bhasa)
- `$` (not used in Bhasa)
- Unmatched quotes

### Unterminated Strings

String without closing quote:

```
Input: "hello
Token: STRING with partial content
```

**Handled gracefully**: Stops at EOF

---

## Performance Characteristics

### Time Complexity
- **Per token**: O(m) where m = token length
- **Complete input**: O(n) where n = input length
- **Single pass**: No backtracking

### Space Complexity
- **O(n)**: Rune slice of input
- **O(1)**: Other fields

### Optimizations
- ✅ Rune slice for O(1) random access
- ✅ Single pass (no re-reading)
- ✅ Peek without consume
- ✅ Minimal string operations

---

## Testing Tips

### Test Categories

1. **Keywords**: All Bengali keywords recognized
2. **Identifiers**: Bengali, English, mixed
3. **Numbers**: Arabic, Bengali, large numbers
4. **Strings**: Empty, Bengali, English, unterminated
5. **Operators**: Single-char, multi-char, all types
6. **Whitespace**: Spaces, tabs, newlines, combinations
7. **Comments**: Line comments, end of line
8. **Positions**: Line/column tracking accuracy
9. **Edge cases**: Empty input, only whitespace, EOF

### Sample Test

```go
func TestBengaliNumber(t *testing.T) {
    input := "১২৩"
    expected := "123"
    
    l := lexer.New(input)
    tok := l.NextToken()
    
    if tok.Type != token.INT {
        t.Errorf("Expected INT, got %s", tok.Type)
    }
    if tok.Literal != expected {
        t.Errorf("Expected %s, got %s", expected, tok.Literal)
    }
}
```

---

## Troubleshooting

### Issue: Identifier Split at Vowel Sign

**Problem**: "কাজ" tokenized as separate tokens

**Solution**: Ensure `isBengaliVowelSign()` is called in `isLetter()`

### Issue: Column Number Off by One

**Problem**: Token column doesn't match actual position

**Solution**: Check column calculation in `newTokenWithPos()`:
```go
Column: l.column - len([]rune(literal)) + 1
```

### Issue: Bengali Digits Not Recognized

**Problem**: "১২৩" returns ILLEGAL

**Solution**: Check `isBengaliDigit()` range (U+09E6 to U+09EF)

### Issue: Comments Not Skipped

**Problem**: "//" appears in token stream

**Solution**: Ensure `skipComment()` called when `//` detected:
```go
if l.ch == '/' && l.peekChar() == '/' {
    l.skipComment()
    return l.NextToken()  // Recursive call
}
```

---

## Extending the Lexer

### Adding a New Keyword

1. Add token type in `token/token.go`:
```go
const (
    MATCH = "মিল"  // match keyword
)
```

2. Add to keywords map:
```go
var keywords = map[string]TokenType{
    "মিল": MATCH,
}
```

3. Use in parser (no lexer changes needed)

### Adding a New Operator

1. Add token type:
```go
const (
    POWER = "**"  // exponentiation
)
```

2. Add case in `NextToken()`:
```go
case '*':
    if l.peekChar() == '*' {
        ch := l.ch
        l.readChar()
        tok = l.newTokenWithPos(token.POWER, string(ch)+string(l.ch))
    } else {
        tok = l.newTokenWithPos(token.ASTERISK, string(l.ch))
    }
```

### Adding Float Support

1. Modify `readNumber()`:
```go
func (l *Lexer) readNumber() (string, bool) {
    startPos := l.position
    isFloat := false
    
    for isDigit(l.ch) || isBengaliDigit(l.ch) {
        l.readChar()
    }
    
    if l.ch == '.' && (isDigit(l.peekChar()) || isBengaliDigit(l.peekChar())) {
        isFloat = true
        l.readChar()  // consume '.'
        
        for isDigit(l.ch) || isBengaliDigit(l.ch) {
            l.readChar()
        }
    }
    
    result := string(l.input[startPos:l.position])
    return token.ConvertBengaliNumber(result), isFloat
}
```

2. Update `NextToken()`:
```go
literal, isFloat := l.readNumber()
tokType := token.INT
if isFloat {
    tokType = token.FLOAT
}
```

---

## Best Practices

### ✅ Do

- Use `NewToken()` for simple cases
- Use `newTokenWithPos()` for consistent position tracking
- Call `skipWhitespace()` at start of `NextToken()`
- Use `peekChar()` for multi-character operators
- Test with both Bengali and English input
- Test edge cases (empty, EOF, large input)

### ❌ Don't

- Don't manually calculate positions (use helper functions)
- Don't modify position without calling `readChar()`
- Don't assume byte == character (use runes)
- Don't forget to handle EOF (null character)
- Don't skip position recording for tokens

---

## Quick Debugging

### Print Lexer State

```go
func (l *Lexer) DebugState() {
    fmt.Printf("Position: %d, ReadPosition: %d\n", l.position, l.readPosition)
    fmt.Printf("Char: '%c' (U+%04X)\n", l.ch, l.ch)
    fmt.Printf("Line: %d, Column: %d\n", l.line, l.column)
}
```

### Print All Tokens

```go
func DebugTokens(input string) {
    l := lexer.New(input)
    
    for {
        tok := l.NextToken()
        fmt.Printf("%s\t%q\tL%d:C%d\n",
            tok.Type, tok.Literal, tok.Line, tok.Column)
        
        if tok.Type == token.EOF {
            break
        }
    }
}
```

---

## Summary

The Bhasa Lexer provides:
- ✅ Complete Bengali language support
- ✅ Dual numeral system (Arabic + Bengali)
- ✅ Precise position tracking
- ✅ Comprehensive operator support
- ✅ Efficient single-pass algorithm
- ✅ Clean, testable API

**Main Entry Point**: `NextToken()`

**Key Features**: Unicode-aware, position-tracked, peek-ahead capable

For detailed documentation, see [lexer-documentation.md](./lexer-documentation.md).

