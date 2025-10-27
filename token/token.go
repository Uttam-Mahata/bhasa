package token

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
}

// Token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT  = "IDENT"  // variable names
	INT    = "INT"    // integers
	STRING = "STRING" // strings

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LTE    = "<="
	GTE    = ">="
	
	AND    = "&&"
	OR     = "||"
	
	// Bitwise operators
	BITWISE_AND = "&"
	BITWISE_OR  = "|"
	BITWISE_XOR = "^"
	BITWISE_NOT = "~"
	LEFT_SHIFT  = "<<"
	RIGHT_SHIFT = ">>"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords (Bengali)
	LET      = "ধরি"       // let (variable declaration)
	FUNCTION = "ফাংশন"     // function
	IF       = "যদি"       // if
	ELSE     = "নাহলে"     // else
	RETURN   = "ফেরত"      // return
	TRUE     = "সত্য"      // true
	FALSE    = "মিথ্যা"    // false
	WHILE    = "যতক্ষণ"    // while
	FOR      = "পর্যন্ত"   // for
	BREAK    = "বিরতি"     // break
	CONTINUE = "চালিয়ে_যাও" // continue
)

var keywords = map[string]TokenType{
	"ধরি":       LET,
	"ফাংশন":     FUNCTION,
	"যদি":       IF,
	"নাহলে":     ELSE,
	"ফেরত":      RETURN,
	"সত্য":      TRUE,
	"মিথ্যা":    FALSE,
	"যতক্ষণ":    WHILE,
	"পর্যন্ত":   FOR,
	"বিরতি":     BREAK,
	"চালিয়ে_যাও": CONTINUE,
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// BengaliDigits maps Bengali numerals to Arabic numerals
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

// ConvertBengaliNumber converts Bengali numerals to Arabic numerals
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

