package parser

import (
	"bhasa/ast"
	"bhasa/lexer"
	"bhasa/token"
	"fmt"
	"strconv"
)

// Operator precedence
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
	PREFIX       // -X or !X or ~X
	CALL         // myFunction(X)
	INDEX        // array[index]
)

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
	token.DOT:      INDEX, // Member access has same precedence as index
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Parser represents a parser
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New creates a new Parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Register prefix parse functions
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BIT_NOT, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	p.registerPrefix(token.STRUCT, p.parseStructDefinition)
	p.registerPrefix(token.ENUM, p.parseEnumDefinition)
	// OOP prefix parsers
	p.registerPrefix(token.NEW, p.parseNewExpression)
	p.registerPrefix(token.THIS, p.parseThisExpression)
	p.registerPrefix(token.SUPER, p.parseSuperExpression)

	// Register infix parse functions
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.BIT_OR, p.parseInfixExpression)
	p.registerInfix(token.BIT_XOR, p.parseInfixExpression)
	p.registerInfix(token.BIT_AND, p.parseInfixExpression)
	p.registerInfix(token.LSHIFT, p.parseInfixExpression)
	p.registerInfix(token.RSHIFT, p.parseInfixExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.PERCENT, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.AS, p.parseTypeCastExpression)
	p.registerInfix(token.DOT, p.parseMemberAccess)

	// Read two tokens to initialize curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns the parser errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("[Line %d, Col %d] expected next token to be %s, got %s instead",
		p.peekToken.Line, p.peekToken.Column, t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) error(msg string) {
	fullMsg := fmt.Sprintf("[Line %d, Col %d] %s",
		p.curToken.Line, p.curToken.Column, msg)
	p.errors = append(p.errors, fullMsg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the program
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.FOR:
		return p.parseForStatement()
	case token.BREAK:
		return p.parseBreakStatement()
	case token.CONTINUE:
		return p.parseContinueStatement()
	case token.IMPORT:
		return p.parseImportStatement()
	case token.CLASS, token.ABSTRACT:
		return p.parseClassDefinition()
	case token.INTERFACE:
		return p.parseInterfaceDefinition()
	case token.IDENT:
		// Check if this is a member assignment (identifier.member = value)
		if p.peekTokenIs(token.DOT) {
			// Parse the member access expression
			left := p.parseExpressionStatement()
			// Check if it's actually an assignment
			if memberAccess, ok := left.Expression.(*ast.MemberAccessExpression); ok {
				if p.peekTokenIs(token.ASSIGN) {
					return p.parseMemberAssignmentStatement(memberAccess)
				}
			}
			return left
		}
		// Check if this is an assignment (identifier followed by =)
		if p.peekTokenIs(token.ASSIGN) {
			return p.parseAssignmentStatement()
		}
		return p.parseExpressionStatement()
	case token.THIS, token.SUPER:
		// Check if this is a member assignment (this.member = value or super.member = value)
		if p.peekTokenIs(token.DOT) {
			// Parse the member access expression
			left := p.parseExpressionStatement()
			// Check if it's actually an assignment
			if memberAccess, ok := left.Expression.(*ast.MemberAccessExpression); ok {
				if p.peekTokenIs(token.ASSIGN) {
					return p.parseMemberAssignmentStatement(memberAccess)
				}
			}
			return left
		}
		return p.parseExpressionStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Check for optional type annotation: ধরি x: পূর্ণসংখ্যা = 10
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume :
		p.nextToken() // move to type token
		stmt.TypeAnnot = p.parseTypeAnnotation()
		if stmt.TypeAnnot == nil {
			return nil
		}
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseAssignmentStatement() *ast.AssignmentStatement {
	stmt := &ast.AssignmentStatement{Token: p.curToken}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseMemberAssignmentStatement(memberAccess *ast.MemberAccessExpression) *ast.MemberAssignmentStatement {
	stmt := &ast.MemberAssignmentStatement{
		Token:  memberAccess.Token,
		Object: memberAccess.Object,
		Member: memberAccess.Member,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

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

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse initialization
	p.nextToken()
	if p.curToken.Type == token.LET {
		stmt.Init = p.parseLetStatement()
	} else if p.curToken.Type == token.IDENT {
		stmt.Init = p.parseAssignmentStatement()
	}
	// If semicolon, skip it
	if p.curToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	// Parse condition
	if p.curToken.Type != token.SEMICOLON {
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}

	// Parse increment
	p.nextToken()
	if p.curToken.Type != token.RPAREN {
		if p.curToken.Type == token.IDENT && p.peekTokenIs(token.ASSIGN) {
			stmt.Increment = p.parseAssignmentStatement()
		} else {
			// Could be an expression statement
			exprStmt := &ast.ExpressionStatement{Token: p.curToken}
			exprStmt.Expression = p.parseExpression(LOWEST)
			stmt.Increment = exprStmt
		}
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	stmt := &ast.BreakStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseContinueStatement() *ast.ContinueStatement {
	stmt := &ast.ContinueStatement{Token: p.curToken}

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.curToken}

	p.nextToken()

	// Parse the module path (should be a string literal)
	stmt.Path = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

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

func (p *Parser) parseIdentifier() ast.Expression {
	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Check if this is a struct literal (Identifier followed by {)
	if p.peekTokenIs(token.LBRACE) {
		return p.parseStructLiteral(ident)
	}

	return ident
}

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

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

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

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

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

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	lit.Parameters, lit.ParameterTypes = p.parseFunctionParameters()

	// Check for optional return type annotation: ফাংশন(x: পূর্ণসংখ্যা): দশমিক { ... }
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume :
		p.nextToken() // move to type token
		lit.ReturnType = p.parseTypeAnnotation()
		if lit.ReturnType == nil {
			return nil
		}
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, []*ast.TypeAnnotation) {
	identifiers := []*ast.Identifier{}
	types := []*ast.TypeAnnotation{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers, types
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	// Check for optional type annotation for parameter
	var typeAnnot *ast.TypeAnnotation
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // consume :
		p.nextToken() // move to type token
		typeAnnot = p.parseTypeAnnotation()
		if typeAnnot == nil {
			return nil, nil
		}
	}
	types = append(types, typeAnnot)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)

		// Check for optional type annotation for parameter
		var typeAnnot *ast.TypeAnnotation
		if p.peekTokenIs(token.COLON) {
			p.nextToken() // consume :
			p.nextToken() // move to type token
			typeAnnot = p.parseTypeAnnotation()
			if typeAnnot == nil {
				return nil, nil
			}
		}
		types = append(types, typeAnnot)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil, nil
	}

	return identifiers, types
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(token.RBRACKET)
	return array
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.curToken, Left: left}

	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}

	return exp
}

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

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

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

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// Type annotation parsing functions

func (p *Parser) isTypeToken(t token.TokenType) bool {
	return t == token.TYPE_BYTE ||
		t == token.TYPE_SHORT ||
		t == token.TYPE_INT ||
		t == token.TYPE_LONG ||
		t == token.TYPE_FLOAT ||
		t == token.TYPE_DOUBLE ||
		t == token.TYPE_CHAR ||
		t == token.TYPE_STRING ||
		t == token.TYPE_BOOLEAN ||
		t == token.TYPE_ARRAY ||
		t == token.TYPE_HASH
}

func (p *Parser) parseTypeAnnotation() *ast.TypeAnnotation {
	if !p.isTypeToken(p.curToken.Type) {
		p.error(fmt.Sprintf("expected type annotation, got %s", p.curToken.Type))
		return nil
	}

	typeAnnot := &ast.TypeAnnotation{
		Token:    p.curToken,
		TypeName: p.curToken.Literal,
	}

	// Check for generic type parameters (e.g., তালিকা<পূর্ণসংখ্যা> or ম্যাপ<লেখা, পূর্ণসংখ্যা>)
	if p.peekTokenIs(token.LT) {
		p.nextToken() // consume <

		// For hash types (ম্যাপ), we expect: <keyType, valueType>
		// For array types (তালিকা), we expect: <elementType>
		if p.curToken.Type == token.TYPE_HASH {
			p.nextToken() // move to first type
			if !p.isTypeToken(p.curToken.Type) {
				p.error(fmt.Sprintf("expected type for hash key, got %s", p.curToken.Type))
				return nil
			}
			typeAnnot.KeyType = p.parseTypeAnnotation()

			if !p.expectPeek(token.COMMA) {
				return nil
			}

			p.nextToken() // move to value type
			if !p.isTypeToken(p.curToken.Type) {
				p.error(fmt.Sprintf("expected type for hash value, got %s", p.curToken.Type))
				return nil
			}
			typeAnnot.ElementType = p.parseTypeAnnotation()
		} else if p.curToken.Type == token.TYPE_ARRAY {
			p.nextToken() // move to element type
			if !p.isTypeToken(p.curToken.Type) {
				p.error(fmt.Sprintf("expected type for array element, got %s", p.curToken.Type))
				return nil
			}
			typeAnnot.ElementType = p.parseTypeAnnotation()
		}

		if !p.expectPeek(token.GT) {
			return nil
		}
	}

	return typeAnnot
}

func (p *Parser) parseTypeCastExpression(left ast.Expression) ast.Expression {
	exp := &ast.TypeCastExpression{
		Token:      p.curToken, // the 'as' token
		Expression: left,
	}

	p.nextToken() // move to type token

	exp.TargetType = p.parseTypeAnnotation()
	if exp.TargetType == nil {
		return nil
	}

	return exp
}

// Struct parsing functions

func (p *Parser) parseStructDefinition() ast.Expression {
	// For now, parse struct literals directly: স্ট্রাক্ট {field: value, ...}
	structToken := p.curToken

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	// Parse as struct literal
	lit := &ast.StructLiteral{
		Token:      structToken,
		Fields:     make(map[string]ast.Expression),
		FieldOrder: []string{},
	}

	// Empty struct
	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return lit
	}

	p.nextToken() // move to first field name

	// Parse first field: name: value
	if !p.curTokenIs(token.IDENT) {
		p.error("expected field name")
		return nil
	}

	fieldName := p.curToken.Literal

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // move to value
	value := p.parseExpression(LOWEST)
	if value == nil {
		return nil
	}

	lit.Fields[fieldName] = value
	lit.FieldOrder = append(lit.FieldOrder, fieldName)

	// Parse remaining fields
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		p.nextToken() // move to field name

		if !p.curTokenIs(token.IDENT) {
			p.error("expected field name")
			return nil
		}

		fieldName := p.curToken.Literal

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // move to value
		value := p.parseExpression(LOWEST)
		if value == nil {
			return nil
		}

		lit.Fields[fieldName] = value
		lit.FieldOrder = append(lit.FieldOrder, fieldName)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return lit
}

func (p *Parser) parseStructFields() []*ast.StructField {
	fields := []*ast.StructField{}

	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return fields
	}

	p.nextToken()

	// Parse first field: name: type
	if !p.curTokenIs(token.IDENT) {
		p.error("expected field name")
		return nil
	}

	field := &ast.StructField{Name: p.curToken.Literal}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // move to type
	field.TypeAnnot = p.parseTypeAnnotation()
	if field.TypeAnnot == nil {
		return nil
	}

	fields = append(fields, field)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		p.nextToken() // move to field name

		if !p.curTokenIs(token.IDENT) {
			p.error("expected field name")
			return nil
		}

		field := &ast.StructField{Name: p.curToken.Literal}

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // move to type
		field.TypeAnnot = p.parseTypeAnnotation()
		if field.TypeAnnot == nil {
			return nil
		}

		fields = append(fields, field)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return fields
}

// Enum parsing functions

func (p *Parser) parseEnumDefinition() ast.Expression {
	// Parse enum definition: গণনা { variant1, variant2, variant3 }
	enumToken := p.curToken

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	enumDef := &ast.EnumDefinition{
		Token:    enumToken,
		Variants: []*ast.EnumVariant{},
	}

	// Empty enum
	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return enumDef
	}

	p.nextToken() // move to first variant name

	// Parse first variant
	if !p.curTokenIs(token.IDENT) {
		p.error("expected variant name")
		return nil
	}

	variant := &ast.EnumVariant{Name: p.curToken.Literal}

	// Check for explicit value: variant = 0
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken() // consume =
		p.nextToken() // move to value

		if !p.curTokenIs(token.INT) {
			p.error("expected integer value for enum variant")
			return nil
		}

		// Parse the integer value
		value, err := strconv.Atoi(p.curToken.Literal)
		if err != nil {
			p.error(fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
			return nil
		}
		variant.Value = &value
	}

	enumDef.Variants = append(enumDef.Variants, variant)

	// Parse remaining variants
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		p.nextToken() // move to variant name

		if !p.curTokenIs(token.IDENT) {
			p.error("expected variant name")
			return nil
		}

		variant := &ast.EnumVariant{Name: p.curToken.Literal}

		// Check for explicit value
		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken() // consume =
			p.nextToken() // move to value

			if !p.curTokenIs(token.INT) {
				p.error("expected integer value for enum variant")
				return nil
			}

			value, err := strconv.Atoi(p.curToken.Literal)
			if err != nil {
				p.error(fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
				return nil
			}
			variant.Value = &value
		}

		enumDef.Variants = append(enumDef.Variants, variant)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return enumDef
}

func (p *Parser) parseMemberAccess(left ast.Expression) ast.Expression {
	exp := &ast.MemberAccessExpression{
		Token:  p.curToken, // the . token
		Object: left,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	exp.Member = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	return exp
}

func (p *Parser) parseStructLiteral(structType *ast.Identifier) ast.Expression {
	lit := &ast.StructLiteral{
		Token:      p.peekToken, // will be the { token
		StructType: structType,
		Fields:     make(map[string]ast.Expression),
		FieldOrder: []string{},
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	if p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		return lit
	}

	p.nextToken()

	// Parse first field: name: value
	if !p.curTokenIs(token.IDENT) {
		p.error("expected field name in struct literal")
		return nil
	}

	fieldName := p.curToken.Literal

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // move to value
	fieldValue := p.parseExpression(LOWEST)
	if fieldValue == nil {
		return nil
	}

	lit.Fields[fieldName] = fieldValue
	lit.FieldOrder = append(lit.FieldOrder, fieldName)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // consume comma
		p.nextToken() // move to field name

		if !p.curTokenIs(token.IDENT) {
			p.error("expected field name in struct literal")
			return nil
		}

		fieldName := p.curToken.Literal

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // move to value
		fieldValue := p.parseExpression(LOWEST)
		if fieldValue == nil {
			return nil
		}

		lit.Fields[fieldName] = fieldValue
		lit.FieldOrder = append(lit.FieldOrder, fieldName)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return lit
}

// ============================================================================
// OOP Parsing Functions
// ============================================================================

// parseClassDefinition parses a class definition
// Syntax: [বিমূর্ত] [চূড়ান্ত] শ্রেণী ClassName [প্রসারিত ParentClass] [বাস্তবায়ন Interface1, Interface2] { ... }
func (p *Parser) parseClassDefinition() *ast.ClassDefinition {
	classDef := &ast.ClassDefinition{
		Token:        p.curToken,
		Fields:       []*ast.ClassField{},
		Methods:      []*ast.MethodDefinition{},
		Constructors: []*ast.ConstructorDefinition{},
		Interfaces:   []*ast.Identifier{},
	}

	// Check for abstract modifier
	if p.curTokenIs(token.ABSTRACT) {
		classDef.IsAbstract = true
		p.nextToken() // move to CLASS or FINAL
	}

	// Check for final modifier
	if p.curTokenIs(token.FINAL) {
		classDef.IsFinal = true
		p.nextToken() // move to CLASS
	}

	if !p.curTokenIs(token.CLASS) {
		p.error("expected শ্রেণী keyword")
		return nil
	}

	// Get class name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	classDef.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Check for extends (প্রসারিত)
	if p.peekTokenIs(token.EXTENDS) {
		p.nextToken() // move to EXTENDS
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		classDef.SuperClass = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	// Check for implements (বাস্তবায়ন)
	if p.peekTokenIs(token.IMPLEMENTS) {
		p.nextToken() // move to IMPLEMENTS
		for {
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			classDef.Interfaces = append(classDef.Interfaces, &ast.Identifier{
				Token: p.curToken,
				Value: p.curToken.Literal,
			})

			if !p.peekTokenIs(token.COMMA) {
				break
			}
			p.nextToken() // skip comma
		}
	}

	// Expect class body
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	p.nextToken() // move past {

	// Parse class body (fields, constructors, methods)
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		// Check for access modifiers
		access := ast.PUBLIC // default
		isStatic := false
		isFinal := false
		isAbstract := false
		isOverride := false

		// Parse modifiers
		for p.curTokenIs(token.PUBLIC) || p.curTokenIs(token.PRIVATE) || p.curTokenIs(token.PROTECTED) ||
			p.curTokenIs(token.STATIC) || p.curTokenIs(token.FINAL) || p.curTokenIs(token.ABSTRACT) ||
			p.curTokenIs(token.OVERRIDE) {

			switch p.curToken.Type {
			case token.PUBLIC:
				access = ast.PUBLIC
			case token.PRIVATE:
				access = ast.PRIVATE
			case token.PROTECTED:
				access = ast.PROTECTED
			case token.STATIC:
				isStatic = true
			case token.FINAL:
				isFinal = true
			case token.ABSTRACT:
				isAbstract = true
			case token.OVERRIDE:
				isOverride = true
			}
			p.nextToken()
		}

		// Check what follows
		if p.curTokenIs(token.CONSTRUCTOR) {
			// Parse constructor
			constructor := p.parseConstructorDefinition()
			if constructor != nil {
				constructor.Access = access
				classDef.Constructors = append(classDef.Constructors, constructor)
			}
			p.nextToken() // Move to next token after constructor
		} else if p.curTokenIs(token.METHOD) {
			// Parse method
			method := p.parseMethodDefinition()
			if method != nil {
				method.Access = access
				method.IsStatic = isStatic
				method.IsFinal = isFinal
				method.IsAbstract = isAbstract
				method.IsOverride = isOverride
				classDef.Methods = append(classDef.Methods, method)
			}
			p.nextToken() // Move to next token after method
		} else if p.curTokenIs(token.IDENT) {
			// Parse field
			field := p.parseClassField()
			if field != nil {
				field.Access = access
				field.IsStatic = isStatic
				field.IsFinal = isFinal
				classDef.Fields = append(classDef.Fields, field)
			}
			p.nextToken() // Move to next token after field
		} else {
			p.error(fmt.Sprintf("unexpected token in class body: %s", p.curToken.Literal))
			p.nextToken()
		}
	}

	if !p.curTokenIs(token.RBRACE) {
		p.error("expected } at end of class definition")
		return nil
	}

	return classDef
}

// parseClassField parses a class field declaration
// Syntax: fieldName: TypeAnnotation;
func (p *Parser) parseClassField() *ast.ClassField {
	field := &ast.ClassField{}

	field.Name = p.curToken.Literal

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken() // move to type
	field.TypeAnnot = p.parseTypeAnnotation()

	// Expect semicolon
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return field
}

// parseMethodDefinition parses a method definition
// Syntax: পদ্ধতি methodName(param1: Type1, param2: Type2): ReturnType { ... }
func (p *Parser) parseMethodDefinition() *ast.MethodDefinition {
	method := &ast.MethodDefinition{
		Token:          p.curToken,
		Parameters:     []*ast.Identifier{},
		ParameterTypes: []*ast.TypeAnnotation{},
	}

	// Expect method name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	method.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Parse parameters
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken() // move past (

	// Parse parameter list
	for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) {
		if !p.curTokenIs(token.IDENT) {
			p.error("expected parameter name")
			return nil
		}

		param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		method.Parameters = append(method.Parameters, param)

		// Expect type annotation
		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // move to type
		paramType := p.parseTypeAnnotation()
		method.ParameterTypes = append(method.ParameterTypes, paramType)

		if p.peekTokenIs(token.COMMA) {
			p.nextToken() // skip comma
			p.nextToken() // move to next parameter
		} else {
			p.nextToken() // move to )
			break
		}
	}

	if !p.curTokenIs(token.RPAREN) {
		p.error("expected ) after parameters")
		return nil
	}

	// Check for return type
	if p.peekTokenIs(token.COLON) {
		p.nextToken() // move to :
		p.nextToken() // move to return type
		method.ReturnType = p.parseTypeAnnotation()
	}

	// For abstract methods, no body
	if method.IsAbstract {
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}
		return method
	}

	// Parse method body
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	method.Body = p.parseBlockStatement()

	return method
}

// parseConstructorDefinition parses a constructor definition
// Syntax: নির্মাতা(param1: Type1, param2: Type2) { ... }
func (p *Parser) parseConstructorDefinition() *ast.ConstructorDefinition {
	constructor := &ast.ConstructorDefinition{
		Token:          p.curToken,
		Parameters:     []*ast.Identifier{},
		ParameterTypes: []*ast.TypeAnnotation{},
	}

	// Parse parameters
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken() // move past (

	// Parse parameter list
	for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) {
		if !p.curTokenIs(token.IDENT) {
			p.error("expected parameter name")
			return nil
		}

		param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		constructor.Parameters = append(constructor.Parameters, param)

		// Expect type annotation
		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // move to type
		paramType := p.parseTypeAnnotation()
		constructor.ParameterTypes = append(constructor.ParameterTypes, paramType)

		if p.peekTokenIs(token.COMMA) {
			p.nextToken() // skip comma
			p.nextToken() // move to next parameter
		} else {
			p.nextToken() // move to )
			break
		}
	}

	if !p.curTokenIs(token.RPAREN) {
		p.error("expected ) after parameters")
		return nil
	}

	// Parse constructor body
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	constructor.Body = p.parseBlockStatement()

	return constructor
}

// parseInterfaceDefinition parses an interface definition
// Syntax: চুক্তি InterfaceName { পদ্ধতি method1(Type1): ReturnType; ... }
func (p *Parser) parseInterfaceDefinition() *ast.InterfaceDefinition {
	interfaceDef := &ast.InterfaceDefinition{
		Token:   p.curToken,
		Methods: []*ast.InterfaceMethod{},
	}

	// Get interface name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	interfaceDef.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect interface body
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	p.nextToken() // move past {

	// Parse interface methods
	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		if !p.curTokenIs(token.METHOD) {
			p.error("expected পদ্ধতি in interface")
			p.nextToken()
			continue
		}

		method := &ast.InterfaceMethod{
			Parameters:     []*ast.Identifier{},
			ParameterTypes: []*ast.TypeAnnotation{},
		}

		// Expect method name
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		method.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

		// Parse parameters
		if !p.expectPeek(token.LPAREN) {
			return nil
		}

		p.nextToken() // move past (

		// Parse parameter list
		for !p.curTokenIs(token.RPAREN) && !p.curTokenIs(token.EOF) {
			if !p.curTokenIs(token.IDENT) {
				p.error("expected parameter name")
				return nil
			}

			param := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			method.Parameters = append(method.Parameters, param)

			// Expect type annotation
			if !p.expectPeek(token.COLON) {
				return nil
			}

			p.nextToken() // move to type
			paramType := p.parseTypeAnnotation()
			method.ParameterTypes = append(method.ParameterTypes, paramType)

			if p.peekTokenIs(token.COMMA) {
				p.nextToken() // skip comma
				p.nextToken() // move to next parameter
			} else {
				p.nextToken() // move to )
				break
			}
		}

		if !p.curTokenIs(token.RPAREN) {
			p.error("expected ) after parameters")
			return nil
		}

		// Check for return type
		if p.peekTokenIs(token.COLON) {
			p.nextToken() // move to :
			p.nextToken() // move to return type
			method.ReturnType = p.parseTypeAnnotation()
		}

		// Expect semicolon
		if p.peekTokenIs(token.SEMICOLON) {
			p.nextToken()
		}

		interfaceDef.Methods = append(interfaceDef.Methods, method)
		p.nextToken()
	}

	if !p.curTokenIs(token.RBRACE) {
		p.error("expected } at end of interface definition")
		return nil
	}

	return interfaceDef
}

// parseNewExpression parses a new instance expression
// Syntax: নতুন ClassName(arg1, arg2, ...)
func (p *Parser) parseNewExpression() ast.Expression {
	newExpr := &ast.NewExpression{
		Token:     p.curToken,
		Arguments: []ast.Expression{},
	}

	// Expect class name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	newExpr.ClassName = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// Expect (
	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	// Parse arguments
	newExpr.Arguments = p.parseExpressionList(token.RPAREN)

	return newExpr
}

// parseThisExpression parses the 'this' keyword (এই)
func (p *Parser) parseThisExpression() ast.Expression {
	return &ast.ThisExpression{Token: p.curToken}
}

// parseSuperExpression parses the 'super' keyword (উর্ধ্ব)
func (p *Parser) parseSuperExpression() ast.Expression {
	return &ast.SuperExpression{Token: p.curToken}
}
