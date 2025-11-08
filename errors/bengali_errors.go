package errors

import "fmt"

// Parser Error Messages (পার্সার ত্রুটি বার্তা)
const (
	// General parsing errors
	ErrUnexpectedToken     = "অপ্রত্যাশিত টোকেন"                                          // Unexpected token
	ErrExpectedToken       = "প্রত্যাশিত টোকেন '%s', কিন্তু পেয়েছি '%s'"                  // Expected token '%s', but got '%s'
	ErrNoPrefixParseFn     = "'%s' এর জন্য কোনো প্রিফিক্স পার্স ফাংশন পাওয়া যায়নি"        // No prefix parse function for '%s'
	ErrExpectedIdentifier  = "প্রত্যাশিত ছিল শনাক্তকারী"                                  // Expected identifier
	ErrExpectedExpression  = "প্রত্যাশিত ছিল এক্সপ্রেশন"                                  // Expected expression

	// Class/OOP related errors
	ErrExpectedClassName   = "শ্রেণীর নাম প্রত্যাশিত"                                     // Expected class name
	ErrExpectedClassKeyword = "'শ্রেণী' কীওয়ার্ড প্রত্যাশিত"                            // Expected 'শ্রেণী' keyword
	ErrExpectedMethodName  = "পদ্ধতির নাম প্রত্যাশিত"                                     // Expected method name
	ErrExpectedParamName   = "প্যারামিটারের নাম প্রত্যাশিত"                               // Expected parameter name
	ErrUnexpectedClassToken = "শ্রেণীর বডিতে অপ্রত্যাশিত টোকেন: %s"                       // Unexpected token in class body: %s
	ErrExpectedClosingBrace = "শ্রেণী সংজ্ঞার শেষে '}' প্রত্যাশিত"                         // Expected '}' at end of class definition
	ErrExpectedClosingParen = "প্যারামিটারের পরে ')' প্রত্যাশিত"                           // Expected ')' after parameters

	// Type annotation errors
	ErrInvalidTypeAnnotation = "অবৈধ টাইপ অ্যানোটেশন"                                    // Invalid type annotation

	// Function/Statement errors
	ErrExpectedLBrace      = "'{' প্রত্যাশিত"                                           // Expected '{'
	ErrExpectedRBrace      = "'}' প্রত্যাশিত"                                           // Expected '}'
	ErrExpectedLParen      = "'(' প্রত্যাশিত"                                           // Expected '('
	ErrExpectedRParen      = "')' প্রত্যাশিত"                                           // Expected ')'
	ErrExpectedSemicolon   = "';' প্রত্যাশিত"                                           // Expected ';'
	ErrExpectedColon       = "':' প্রত্যাশিত"                                           // Expected ':'
)

// Compiler Error Messages (কম্পাইলার ত্রুটি বার্তা)
const (
	ErrUndefinedVariable   = "অসংজ্ঞায়িত ভেরিয়েবল: %s"                                  // Undefined variable: %s
	ErrUnknownOperator     = "অজানা অপারেটর: %s"                                       // Unknown operator: %s
	ErrTypeMismatch        = "টাইপ মিসম্যাচ: %s এবং %s"                                 // Type mismatch: %s and %s
	ErrInvalidOperation    = "অবৈধ অপারেশন: %s"                                        // Invalid operation: %s
	ErrTooManyConstants    = "অত্যধিক ধ্রুবক: %d (সর্বোচ্চ %d)"                           // Too many constants: %d (max %d)
	ErrConstantIndexOutOfRange = "ধ্রুবক সূচক সীমার বাইরে: %d (দৈর্ঘ্য %d)"              // Constant index out of range: %d (length %d)
)

// VM/Runtime Error Messages (ভিএম/রানটাইম ত্রুটি বার্তা)
const (
	ErrStackOverflow       = "স্ট্যাক ওভারফ্লো"                                         // Stack overflow
	ErrStackUnderflow      = "স্ট্যাক আন্ডারফ্লো"                                       // Stack underflow
	ErrDivisionByZero      = "শূন্য দ্বারা ভাগ"                                         // Division by zero
	ErrIndexOutOfBounds    = "সূচক সীমার বাইরে: %d"                                     // Index out of bounds: %d
	ErrInvalidArrayIndex   = "অবৈধ অ্যারে সূচক: %s"                                    // Invalid array index: %s
	ErrInvalidHashKey      = "অবৈধ হ্যাশ কী: %s"                                       // Invalid hash key: %s
	ErrNotAFunction        = "ফাংশন নয়: %s"                                           // Not a function: %s
	ErrNotAClass           = "শ্রেণী নয়: %s"                                           // Not a class: %s
	ErrNotAnObject         = "অবজেক্ট নয়: %s"                                         // Not an object: %s
	ErrWrongNumberOfArgs   = "ভুল সংখ্যক আর্গুমেন্ট: প্রত্যাশিত %d, পেয়েছি %d"           // Wrong number of arguments: expected %d, got %d
	ErrUnsupportedOperation = "অসমর্থিত অপারেশন: %s এবং %s এর জন্য"                      // Unsupported operation for %s and %s
	ErrNullPointer         = "নাল পয়েন্টার ত্রুটি"                                      // Null pointer error
	ErrPropertyNotFound    = "প্রপার্টি পাওয়া যায়নি: %s"                               // Property not found: %s
	ErrMethodNotFound      = "পদ্ধতি পাওয়া যায়নি: %s"                                  // Method not found: %s
	ErrInvalidThis         = "'এই' শুধুমাত্র পদ্ধতির মধ্যে ব্যবহার করা যায়"              // 'this' can only be used in methods
	ErrInvalidSuper        = "'উর্ধ্ব' শুধুমাত্র চাইল্ড ক্লাসে ব্যবহার করা যায়"           // 'super' can only be used in child classes
)

// Helper functions for formatted error messages
func UnexpectedToken(token string) string {
	return fmt.Sprintf("%s: %s", ErrUnexpectedToken, token)
}

func ExpectedToken(expected, got string) string {
	return fmt.Sprintf(ErrExpectedToken, expected, got)
}

func NoPrefixParseFn(tokenType string) string {
	return fmt.Sprintf(ErrNoPrefixParseFn, tokenType)
}

func UnexpectedClassToken(token string) string {
	return fmt.Sprintf(ErrUnexpectedClassToken, token)
}

func UndefinedVariable(name string) string {
	return fmt.Sprintf(ErrUndefinedVariable, name)
}

func UnknownOperator(op string) string {
	return fmt.Sprintf(ErrUnknownOperator, op)
}

func TypeMismatch(t1, t2 string) string {
	return fmt.Sprintf(ErrTypeMismatch, t1, t2)
}

func InvalidOperation(op string) string {
	return fmt.Sprintf(ErrInvalidOperation, op)
}

func TooManyConstants(count, max int) string {
	return fmt.Sprintf(ErrTooManyConstants, count, max)
}

func ConstantIndexOutOfRange(index, length int) string {
	return fmt.Sprintf(ErrConstantIndexOutOfRange, index, length)
}

func IndexOutOfBounds(index int) string {
	return fmt.Sprintf(ErrIndexOutOfBounds, index)
}

func InvalidArrayIndex(indexType string) string {
	return fmt.Sprintf(ErrInvalidArrayIndex, indexType)
}

func InvalidHashKey(keyType string) string {
	return fmt.Sprintf(ErrInvalidHashKey, keyType)
}

func NotAFunction(objType string) string {
	return fmt.Sprintf(ErrNotAFunction, objType)
}

func NotAClass(objType string) string {
	return fmt.Sprintf(ErrNotAClass, objType)
}

func NotAnObject(objType string) string {
	return fmt.Sprintf(ErrNotAnObject, objType)
}

func WrongNumberOfArgs(expected, got int) string {
	return fmt.Sprintf(ErrWrongNumberOfArgs, expected, got)
}

func UnsupportedOperation(t1, t2 string) string {
	return fmt.Sprintf(ErrUnsupportedOperation, t1, t2)
}

func PropertyNotFound(prop string) string {
	return fmt.Sprintf(ErrPropertyNotFound, prop)
}

func MethodNotFound(method string) string {
	return fmt.Sprintf(ErrMethodNotFound, method)
}
