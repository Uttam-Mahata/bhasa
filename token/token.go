package token

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int // Line number where token appears
	Column  int // Column number where token appears
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

	AND = "&&" // Logical AND
	OR  = "||" // Logical OR

	// Bitwise operators
	BIT_AND   = "&"  // Bitwise AND
	BIT_OR    = "|"  // Bitwise OR
	BIT_XOR   = "^"  // Bitwise XOR
	BIT_NOT   = "~"  // Bitwise NOT
	LSHIFT    = "<<" // Left shift
	RSHIFT    = ">>" // Right shift

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
	LET      = "ধরি"         // let (variable declaration)
	FUNCTION = "ফাংশন"       // function
	IF       = "যদি"         // if
	ELSE     = "নাহলে"       // else
	RETURN   = "ফেরত"        // return
	TRUE     = "সত্য"        // true
	FALSE    = "মিথ্যা"      // false
	WHILE    = "যতক্ষণ"      // while
	FOR      = "পর্যন্ত"     // for
	BREAK    = "বিরতি"       // break
	CONTINUE = "চালিয়ে_যাও"  // continue
	IMPORT   = "অন্তর্ভুক্ত"  // import/include

	// Type keywords (Bengali)
	TYPE_BYTE    = "বাইট"           // byte type
	TYPE_SHORT   = "ছোট_সংখ্যা"     // short type
	TYPE_INT     = "পূর্ণসংখ্যা"    // int type
	TYPE_LONG    = "দীর্ঘ_সংখ্যা"   // long type
	TYPE_FLOAT   = "দশমিক"          // float type
	TYPE_DOUBLE  = "দশমিক_দ্বিগুণ"  // double type
	TYPE_CHAR    = "অক্ষর"          // char type
	TYPE_STRING  = "লেখা"           // string type
	TYPE_BOOLEAN = "বুলিয়ান"       // boolean type
	TYPE_ARRAY   = "তালিকা"         // array type
	TYPE_HASH    = "ম্যাপ"          // hash/map type
	AS           = "as"            // type casting keyword

	// Struct and Enum keywords
	STRUCT = "স্ট্রাক্ট" // struct keyword
	ENUM   = "enum"     // enum keyword
	DOT    = "."        // dot for field access
	ARROW  = "=>"       // arrow for pattern matching
)

var keywords = map[string]TokenType{
	"ধরি":         LET,
	"ফাংশন":       FUNCTION,
	"যদি":         IF,
	"নাহলে":       ELSE,
	"ফেরত":        RETURN,
	"সত্য":        TRUE,
	"মিথ্যা":      FALSE,
	"যতক্ষণ":      WHILE,
	"পর্যন্ত":     FOR,
	"বিরতি":       BREAK,
	"চালিয়ে_যাও":  CONTINUE,
	"অন্তর্ভুক্ত": IMPORT,
	// Type keywords
	"বাইট":           TYPE_BYTE,
	"ছোট_সংখ্যা":     TYPE_SHORT,
	"পূর্ণসংখ্যা":    TYPE_INT,
	"দীর্ঘ_সংখ্যা":   TYPE_LONG,
	"দশমিক":          TYPE_FLOAT,
	"দশমিক_দ্বিগুণ":  TYPE_DOUBLE,
	"অক্ষর":          TYPE_CHAR,
	"লেখা":           TYPE_STRING,
	"বুলিয়ান":       TYPE_BOOLEAN,
	"তালিকা":         TYPE_ARRAY,
	"ম্যাপ":          TYPE_HASH,
	"as":            AS,
	// Struct and Enum keywords
	"স্ট্রাক্ট": STRUCT,
	"enum":     ENUM,
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

