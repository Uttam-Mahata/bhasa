package lexer

import (
	"bhasa/token"
	"unicode"
)

// Lexer represents a lexical analyzer
type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readChar()
	return l
}

// readChar reads the next character and advances position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar looks ahead at the next character without advancing
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '%':
		tok = newToken(token.PERCENT, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) || isBengaliDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads an identifier (variable name or keyword)
func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return string(l.input[startPos:l.position])
}

// readNumber reads a number (supports both Arabic and Bengali numerals)
func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) || isBengaliDigit(l.ch) {
		l.readChar()
	}
	result := string(l.input[startPos:l.position])
	// Convert Bengali digits to Arabic
	return token.ConvertBengaliNumber(result)
}

// readString reads a string literal
func (l *Lexer) readString() string {
	startPos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return string(l.input[startPos:l.position])
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips comments until end of line
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

// isLetter checks if a character is a letter (including Bengali)
func isLetter(ch rune) bool {
	// Check for letters, Bengali vowel signs (মাত্রা), and other combining marks
	return unicode.IsLetter(ch) || ch == '_' || isBengaliVowelSign(ch) || unicode.Is(unicode.Mn, ch)
}

// isBengaliVowelSign checks if a character is a Bengali vowel sign or combining mark
func isBengaliVowelSign(ch rune) bool {
	// Bengali vowel signs (মাত্রা) and diacritics: U+0981 to U+09CD
	// This includes: ঁ ং ঃ, vowel signs, and hasant
	return (ch >= 0x0981 && ch <= 0x09CD) || ch == 0x09D7
}

// isDigit checks if a character is an Arabic digit
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

// isBengaliDigit checks if a character is a Bengali digit
func isBengaliDigit(ch rune) bool {
	return '০' <= ch && ch <= '৯'
}

// newToken creates a new token
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
