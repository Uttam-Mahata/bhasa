# Parser Package Documentation

## Overview

The **Parser** package transforms a stream of tokens from the lexer into an **Abstract Syntax Tree (AST)**. It uses **Pratt parsing** (top-down operator precedence parsing) to handle operator precedence and associativity elegantly.

## Table of Contents

- [What is a Parser?](#what-is-a-parser)
- [Pratt Parsing](#pratt-parsing)
- [Parser Structure](#parser-structure)
- [Precedence Levels](#precedence-levels)
- [Parse Functions](#parse-functions)
- [Statement Parsing](#statement-parsing)
- [Expression Parsing](#expression-parsing)
- [Error Handling](#error-handling)
- [Usage Examples](#usage-examples)

## What is a Parser?

A **parser** performs **syntax analysis** - converting a flat sequence of tokens into a hierarchical tree structure (AST) that represents the program's structure.

### Compilation Pipeline

```
Source Code → [Lexer] → Tokens → [Parser] → AST → [Evaluator/Compiler]
                                      ↑
                               Parser Package
```

### Example

**Input (Tokens):**
```
LET, IDENT("x"), ASSIGN, INT(5), SEMICOLON
```

**Output (AST):**
```
*ast.LetStatement{
    Name: &ast.Identifier{Value: "x"},
    Value: &ast.IntegerLiteral{Value: 5},
}
```

## Pratt Parsing

The parser uses **Pratt parsing**, invented by Vaughan Pratt in 1973.

### Why Pratt Parsing?

**Traditional recursive descent:**
- One function per precedence level
- Verbose and rigid
- Hard to modify

**Pratt parsing:**
- Single recursive function
- Table-driven precedence
- Easy to extend
- Elegant handling of operators

### Key Concepts

#### 1. Prefix Parse Functions

Handle tokens that **start** an expression:
- Identifiers: `x`
- Literals: `5`, `"hello"`, `true`
- Prefix operators: `-5`, `!flag`
- Grouping: `(expression)`

#### 2. Infix Parse Functions

Handle operators that appear **between** expressions:
- Binary operators: `a + b`, `x * y`
- Comparisons: `x == y`, `a < b`
- Function calls: `func(args)`
- Index access: `arr[0]`

#### 3. Precedence

Determines evaluation order:

```bengali
১ + ২ * ৩        // 1 + (2 * 3) = 7
                // Not (1 + 2) * 3 = 9
```

Precedence levels (lowest to highest):
1. Logical OR (`||`)
2. Logical AND (`&&`)
3. Bitwise OR (`|`)
4. Bitwise XOR (`^`)
5. Bitwise AND (`&`)
6. Equality (`==`, `!=`)
7. Comparison (`<`, `>`, `<=`, `>=`)
8. Shift (`<<`, `>>`)
9. Addition/Subtraction (`+`, `-`)
10. Multiplication/Division (`*`, `/`, `%`)
11. Prefix (unary `-`, `!`, `~`)
12. Call/Index (`func()`, `arr[]`)

## Parser Structure

### Parser Type

```go
type Parser struct {
    l      *lexer.Lexer           // Lexer for token stream
    errors []string               // Parse errors
    
    curToken  token.Token         // Current token
    peekToken token.Token         // Next token (lookahead)
    
    prefixParseFns map[token.TokenType]prefixParseFn
    infixParseFns  map[token.TokenType]infixParseFn
}
```

### Parse Function Types

```go
type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression
```

**Prefix function:**
- Takes no arguments
- Returns expression

**Infix function:**
- Takes left expression
- Returns combined expression

### Initialization

```go
func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        l:      l,
        errors: []string{},
    }
    
    // Register prefix parse functions
    p.registerPrefix(token.IDENT, p.parseIdentifier)
    p.registerPrefix(token.INT, p.parseIntegerLiteral)
    p.registerPrefix(token.BANG, p.parsePrefixExpression)
    // ... more registrations
    
    // Register infix parse functions
    p.registerInfix(token.PLUS, p.parseInfixExpression)
    p.registerInfix(token.ASTERISK, p.parseInfixExpression)
    p.registerInfix(token.LPAREN, p.parseCallExpression)
    // ... more registrations
    
    // Read two tokens to initialize curToken and peekToken
    p.nextToken()
    p.nextToken()
    
    return p
}
```

## Precedence Levels

### Precedence Constants

```go
const (
    _ int = iota
    LOWEST
    LOGICAL_OR   // ||
    LOGICAL_AND  // &&
    BIT_OR       // |
    BIT_XOR      // ^
    BIT_AND      // &
    EQUALS       // ==
    LESSGREATER  // > or <
    SHIFT        // << >>
    SUM          // +
    PRODUCT      // *
    PREFIX       // -X or !X
    CALL         // myFunction(X)
    INDEX        // array[index]
)
```

### Precedence Map

```go
var precedences = map[token.TokenType]int{
    token.OR:       LOGICAL_OR,
    token.AND:      LOGICAL_AND,
    token.BIT_OR:   BIT_OR,
    token.BIT_XOR:  BIT_XOR,
    token.BIT_AND:  BIT_AND,
    token.EQ:       EQUALS,
    token.NOT_EQ:   EQUALS,
    token.LT:       LESSGREATER,
    token.GT:       LESSGREATER,
    token.LTE:      LESSGREATER,
    token.GTE:      LESSGREATER,
    token.LSHIFT:   SHIFT,
    token.RSHIFT:   SHIFT,
    token.PLUS:     SUM,
    token.MINUS:    SUM,
    token.SLASH:    PRODUCT,
    token.ASTERISK: PRODUCT,
    token.PERCENT:  PRODUCT,
    token.LPAREN:   CALL,
    token.LBRACKET: INDEX,
    token.DOT:      INDEX,
}
```

### Precedence Helpers

```go
func (p *Parser) peekPrecedence() int {
    if p, ok := precedences[p.peekToken.Type]; ok {
        return p
    }
    return LOWEST
}

func (p *Parser) curPrecedence() int {
    if p, ok := precedences[p.curToken.Type]; ok {
        return p
    }
    return LOWEST
}
```

## Parse Functions

### Main Entry Point

```go
func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}
    
    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    
    return program
}
```

**Process:**
1. Create empty program node
2. Parse statements until EOF
3. Add each statement to program
4. Return complete AST

### Expression Parsing (Core Algorithm)

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
    // Get prefix parse function
    prefix := p.prefixParseFns[p.curToken.Type]
    if prefix == nil {
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
    }
    
    // Parse left side
    leftExp := prefix()
    
    // Parse infix operators while precedence allows
    for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        
        p.nextToken()
        leftExp = infix(leftExp)
    }
    
    return leftExp
}
```

**Algorithm:**
1. Get prefix function for current token
2. Parse left expression
3. While next operator has higher precedence:
   - Get infix function
   - Parse right expression
   - Combine into new expression
4. Return final expression

### Example: Parsing `১ + ২ * ৩`

```
Tokens: INT(1) PLUS INT(2) ASTERISK INT(3)

Step 1: parseExpression(LOWEST)
  prefix = parseIntegerLiteral
  leftExp = IntegerLiteral(1)
  
Step 2: peek PLUS (precedence SUM)
  SUM > LOWEST → continue
  p.nextToken() → curToken = PLUS
  infix = parseInfixExpression
  
Step 3: parseInfixExpression(IntegerLiteral(1))
  operator = "+"
  precedence = SUM
  p.nextToken() → curToken = INT(2)
  right = parseExpression(SUM)
  
Step 4: parseExpression(SUM)
  prefix = parseIntegerLiteral
  leftExp = IntegerLiteral(2)
  peek ASTERISK (precedence PRODUCT)
  PRODUCT > SUM → continue
  
Step 5: parseInfixExpression(IntegerLiteral(2))
  operator = "*"
  right = IntegerLiteral(3)
  return InfixExpression(2 * 3)
  
Step 6: Back to step 3
  right = InfixExpression(2 * 3)
  return InfixExpression(1 + (2 * 3))

Result:
InfixExpression{
  Left: IntegerLiteral(1),
  Operator: "+",
  Right: InfixExpression{
    Left: IntegerLiteral(2),
    Operator: "*",
    Right: IntegerLiteral(3),
  },
}
```

## Statement Parsing

### Let Statement

```bengali
ধরি x = ৫;
ধরি y: পূর্ণসংখ্যা = ১০;
```

```go
func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken}
    
    // Expect identifier
    if !p.expectPeek(token.IDENT) {
        return nil
    }
    stmt.Name = &ast.Identifier{
        Token: p.curToken,
        Value: p.curToken.Literal,
    }
    
    // Optional type annotation
    if p.peekTokenIs(token.COLON) {
        p.nextToken() // consume :
        p.nextToken() // move to type
        stmt.TypeAnnot = p.parseTypeAnnotation()
    }
    
    // Expect =
    if !p.expectPeek(token.ASSIGN) {
        return nil
    }
    
    p.nextToken()
    stmt.Value = p.parseExpression(LOWEST)
    
    // Optional semicolon
    if p.peekTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    
    return stmt
}
```

### Return Statement

```bengali
ফেরত x + y;
```

```go
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}
    
    p.nextToken()
    stmt.ReturnValue = p.parseExpression(LOWEST)
    
    if p.peekTokenIs(token.SEMICOLON) {
        p.nextToken()
    }
    
    return stmt
}
```

### While Statement

```bengali
যতক্ষণ (x < ১০) {
    x = x + ১;
}
```

```go
func (p *Parser) parseWhileStatement() *ast.WhileStatement {
    stmt := &ast.WhileStatement{Token: p.curToken}
    
    if !p.expectPeek(token.LPAREN) {
        return nil
    }
    
    p.nextToken()
    stmt.Condition = p.parseExpression(LOWEST)
    
    if !p.expectPeek(token.RPAREN) {
        return nil
    }
    
    if !p.expectPeek(token.LBRACE) {
        return nil
    }
    
    stmt.Body = p.parseBlockStatement()
    
    return stmt
}
```

### If Expression

```bengali
যদি (x > ৫) {
    লেখ("বড়");
} নাহলে {
    লেখ("ছোট");
}
```

```go
func (p *Parser) parseIfExpression() ast.Expression {
    expression := &ast.IfExpression{Token: p.curToken}
    
    if !p.expectPeek(token.LPAREN) {
        return nil
    }
    
    p.nextToken()
    expression.Condition = p.parseExpression(LOWEST)
    
    if !p.expectPeek(token.RPAREN) {
        return nil
    }
    
    if !p.expectPeek(token.LBRACE) {
        return nil
    }
    
    expression.Consequence = p.parseBlockStatement()
    
    // Optional else clause
    if p.peekTokenIs(token.ELSE) {
        p.nextToken()
        
        if !p.expectPeek(token.LBRACE) {
            return nil
        }
        
        expression.Alternative = p.parseBlockStatement()
    }
    
    return expression
}
```

## Expression Parsing

### Literals

#### Integer Literal

```bengali
৫
123
```

```go
func (p *Parser) parseIntegerLiteral() ast.Expression {
    lit := &ast.IntegerLiteral{Token: p.curToken}
    
    value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
    if err != nil {
        msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
        p.errors = append(p.errors, msg)
        return nil
    }
    
    lit.Value = value
    return lit
}
```

#### String Literal

```bengali
"হ্যালো বিশ্ব"
```

```go
func (p *Parser) parseStringLiteral() ast.Expression {
    return &ast.StringLiteral{
        Token: p.curToken,
        Value: p.curToken.Literal,
    }
}
```

### Prefix Expressions

```bengali
-৫
!সত্য
~০xFF
```

```go
func (p *Parser) parsePrefixExpression() ast.Expression {
    expression := &ast.PrefixExpression{
        Token:    p.curToken,
        Operator: p.curToken.Literal,
    }
    
    p.nextToken()
    expression.Right = p.parseExpression(PREFIX)
    
    return expression
}
```

### Infix Expressions

```bengali
x + y
a * b
n == ৫
```

```go
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    expression := &ast.InfixExpression{
        Token:    p.curToken,
        Operator: p.curToken.Literal,
        Left:     left,
    }
    
    precedence := p.curPrecedence()
    p.nextToken()
    expression.Right = p.parseExpression(precedence)
    
    return expression
}
```

### Function Literals

```bengali
ফাংশন(x, y) {
    ফেরত x + y;
}
```

```go
func (p *Parser) parseFunctionLiteral() ast.Expression {
    lit := &ast.FunctionLiteral{Token: p.curToken}
    
    if !p.expectPeek(token.LPAREN) {
        return nil
    }
    
    lit.Parameters, lit.ParameterTypes = p.parseFunctionParameters()
    
    // Optional return type
    if p.peekTokenIs(token.COLON) {
        p.nextToken()
        p.nextToken()
        lit.ReturnType = p.parseTypeAnnotation()
    }
    
    if !p.expectPeek(token.LBRACE) {
        return nil
    }
    
    lit.Body = p.parseBlockStatement()
    
    return lit
}
```

### Call Expressions

```bengali
add(১, ২)
লেখ("হ্যালো")
```

```go
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
    exp := &ast.CallExpression{
        Token:    p.curToken,
        Function: function,
    }
    exp.Arguments = p.parseExpressionList(token.RPAREN)
    return exp
}
```

### Array Literals

```bengali
[১, ২, ৩]
["a", "b", "c"]
[সত্য, মিথ্যা]
```

```go
func (p *Parser) parseArrayLiteral() ast.Expression {
    array := &ast.ArrayLiteral{Token: p.curToken}
    array.Elements = p.parseExpressionList(token.RBRACKET)
    return array
}
```

### Hash Literals

```bengali
{"নাম": "রহিম", "বয়স": ২৫}
```

```go
func (p *Parser) parseHashLiteral() ast.Expression {
    hash := &ast.HashLiteral{Token: p.curToken}
    hash.Pairs = make(map[ast.Expression]ast.Expression)
    
    for !p.peekTokenIs(token.RBRACE) {
        p.nextToken()
        key := p.parseExpression(LOWEST)
        
        if !p.expectPeek(token.COLON) {
            return nil
        }
        
        p.nextToken()
        value := p.parseExpression(LOWEST)
        
        hash.Pairs[key] = value
        
        if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
            return nil
        }
    }
    
    if !p.expectPeek(token.RBRACE) {
        return nil
    }
    
    return hash
}
```

## Error Handling

### Error Collection

```go
type Parser struct {
    // ...
    errors []string
}

func (p *Parser) Errors() []string {
    return p.errors
}
```

### Error Types

#### Peek Error

```go
func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf(
        "[Line %d, Col %d] expected next token to be %s, got %s instead",
        p.peekToken.Line,
        p.peekToken.Column,
        t,
        p.peekToken.Type,
    )
    p.errors = append(p.errors, msg)
}
```

#### No Prefix Parse Function

```go
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
    msg := fmt.Sprintf("no prefix parse function for %s found", t)
    p.errors = append(p.errors, msg)
}
```

### Error Recovery

Parser continues after errors to find multiple issues in one pass.

## Usage Examples

### Basic Usage

```go
import (
    "bhasa/lexer"
    "bhasa/parser"
    "fmt"
)

input := `ধরি x = ৫;`

l := lexer.New(input)
p := parser.New(l)

program := p.ParseProgram()

// Check for errors
if len(p.Errors()) != 0 {
    for _, err := range p.Errors() {
        fmt.Println(err)
    }
    return
}

// Use AST
fmt.Println(program.String())
```

### Parsing Expressions

```go
input := `১ + ২ * ৩`

l := lexer.New(input)
p := parser.New(l)

program := p.ParseProgram()
// program.Statements[0] is *ast.ExpressionStatement
// statement.Expression is *ast.InfixExpression
```

### Parsing Functions

```go
input := `
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};
`

l := lexer.New(input)
p := parser.New(l)

program := p.ParseProgram()
// program.Statements[0] is *ast.LetStatement
// statement.Value is *ast.FunctionLiteral
```

## Summary

The **Parser** package provides:

✅ **Pratt Parsing**: Elegant operator precedence handling  
✅ **Complete Syntax**: All Bhasa language constructs  
✅ **Error Recovery**: Multiple errors reported  
✅ **Type Annotations**: Optional type system support  
✅ **OOP Support**: Class, method, interface parsing  
✅ **Extensible**: Easy to add new syntax  

For detailed parsing algorithms and examples, see [parser-documentation.md](./parser-documentation.md).

---

The parser transforms flat token streams into meaningful hierarchical structures, enabling subsequent compilation or interpretation phases.

