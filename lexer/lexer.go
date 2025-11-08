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
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{
		input:  []rune(input),
		line:   1,
		column: 0,
	}
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
	l.column++
	
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
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
	
	// Record position at start of token
	tokLine := l.line
	tokCol := l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.EQ, string(ch)+string(l.ch))
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.ARROW, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.ASSIGN, string(l.ch))
		}
	case '+':
		tok = l.newTokenWithPos(token.PLUS, string(l.ch))
	case '-':
		tok = l.newTokenWithPos(token.MINUS, string(l.ch))
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.NOT_EQ, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.BANG, string(l.ch))
		}
	case '*':
		tok = l.newTokenWithPos(token.ASTERISK, string(l.ch))
	case '/':
		if l.peekChar() == '/' {
			l.skipComment()
			return l.NextToken()
		} else {
			tok = l.newTokenWithPos(token.SLASH, string(l.ch))
		}
	case '%':
		tok = l.newTokenWithPos(token.PERCENT, string(l.ch))
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.LTE, string(ch)+string(l.ch))
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.LSHIFT, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.LT, string(l.ch))
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.GTE, string(ch)+string(l.ch))
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.RSHIFT, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.GT, string(l.ch))
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.AND, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.BIT_AND, string(l.ch))
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = l.newTokenWithPos(token.OR, string(ch)+string(l.ch))
		} else {
			tok = l.newTokenWithPos(token.BIT_OR, string(l.ch))
		}
	case '^':
		tok = l.newTokenWithPos(token.BIT_XOR, string(l.ch))
	case '~':
		tok = l.newTokenWithPos(token.BIT_NOT, string(l.ch))
	case ',':
		tok = l.newTokenWithPos(token.COMMA, string(l.ch))
	case ';':
		tok = l.newTokenWithPos(token.SEMICOLON, string(l.ch))
	case ':':
		tok = l.newTokenWithPos(token.COLON, string(l.ch))
	case '.':
		tok = l.newTokenWithPos(token.DOT, string(l.ch))
	case '(':
		tok = l.newTokenWithPos(token.LPAREN, string(l.ch))
	case ')':
		tok = l.newTokenWithPos(token.RPAREN, string(l.ch))
	case '{':
		tok = l.newTokenWithPos(token.LBRACE, string(l.ch))
	case '}':
		tok = l.newTokenWithPos(token.RBRACE, string(l.ch))
	case '[':
		tok = l.newTokenWithPos(token.LBRACKET, string(l.ch))
	case ']':
		tok = l.newTokenWithPos(token.RBRACKET, string(l.ch))
	case '"':
		literal := l.readString()
		tok = token.Token{
			Type:    token.STRING,
			Literal: literal,
			Line:    tokLine,
			Column:  tokCol,
		}
	case 0:
		tok = token.Token{
			Type:    token.EOF,
			Literal: "",
			Line:    l.line,
			Column:  l.column,
		}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok = token.Token{
				Type:    token.LookupIdent(literal),
				Literal: literal,
				Line:    tokLine,
				Column:  tokCol,
			}
			return tok
		} else if isDigit(l.ch) || isBengaliDigit(l.ch) {
			literal := l.readNumber()
			tok = token.Token{
				Type:    token.INT,
				Literal: literal,
				Line:    tokLine,
				Column:  tokCol,
			}
			return tok
		} else {
			tok = l.newTokenWithPos(token.ILLEGAL, string(l.ch))
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads an identifier (variable name or keyword)
// Identifiers can contain letters, underscores, and digits (but must start with a letter or underscore)
func (l *Lexer) readIdentifier() string {
	startPos := l.position
	// Read first character (must be letter or underscore)
	for isLetter(l.ch) || isDigit(l.ch) || isBengaliDigit(l.ch) {
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

// newTokenWithPos creates a new token with position information
func (l *Lexer) newTokenWithPos(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.line,
		Column:  l.column - len([]rune(literal)) + 1,
	}
}
