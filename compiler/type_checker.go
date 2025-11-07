package compiler

import (
	"bhasa/ast"
	"bhasa/object"
	"fmt"
)

// TypeChecker performs static type checking
type TypeChecker struct {
	errors []string
}

// NewTypeChecker creates a new TypeChecker
func NewTypeChecker() *TypeChecker {
	return &TypeChecker{
		errors: []string{},
	}
}

// Errors returns the type checking errors
func (tc *TypeChecker) Errors() []string {
	return tc.errors
}

// addError adds a type checking error
func (tc *TypeChecker) addError(msg string) {
	tc.errors = append(tc.errors, msg)
}

// InferType infers the type of an expression.
// Returns the Bengali type name if type can be inferred, or empty string if unknown.
// Supported expressions: IntegerLiteral, StringLiteral, Boolean, ArrayLiteral,
// HashLiteral, FunctionLiteral, InfixExpression, PrefixExpression.
func (tc *TypeChecker) InferType(expr ast.Expression) string {
	switch e := expr.(type) {
	case *ast.IntegerLiteral:
		return "পূর্ণসংখ্যা"
	case *ast.StringLiteral:
		return "লেখা"
	case *ast.Boolean:
		return "বুলিয়ান"
	case *ast.ArrayLiteral:
		return "তালিকা"
	case *ast.HashLiteral:
		return "হ্যাশ"
	case *ast.FunctionLiteral:
		return "ফাংশন_টাইপ"
	case *ast.InfixExpression:
		// Arithmetic and comparison operators
		switch e.Operator {
		case "+", "-", "*", "/", "%", "&", "|", "^", "<<", ">>":
			return "পূর্ণসংখ্যা"
		case "==", "!=", "<", ">", "<=", ">=", "&&", "||":
			return "বুলিয়ান"
		default:
			return ""
		}
	case *ast.PrefixExpression:
		switch e.Operator {
		case "!", "~":
			return "বুলিয়ান"
		case "-":
			return "পূর্ণসংখ্যা"
		default:
			return ""
		}
	default:
		return "" // unknown type
	}
}

// CheckLetStatement checks type consistency in let statements
func (tc *TypeChecker) CheckLetStatement(stmt *ast.LetStatement) bool {
	if stmt.TypeAnnotation == nil {
		// No type annotation, no checking needed
		return true
	}

	expectedType := stmt.TypeAnnotation.Type
	inferredType := tc.InferType(stmt.Value)

	if inferredType == "" {
		// Cannot infer type, skip checking
		return true
	}

	if expectedType != inferredType {
		tc.addError(fmt.Sprintf(
			"type mismatch: variable '%s' declared as %s but assigned %s",
			stmt.Name.Value, expectedType, inferredType))
		return false
	}

	return true
}

// CheckFunctionReturn checks if function returns match declared return type
func (tc *TypeChecker) CheckFunctionReturn(fn *ast.FunctionLiteral) bool {
	if fn.ReturnType == nil {
		// No return type annotation, no checking needed
		return true
	}

	expectedReturnType := fn.ReturnType.Type
	
	// Walk through the function body to find return statements
	return tc.checkBlockForReturns(fn.Body, expectedReturnType)
}

// checkBlockForReturns recursively checks return statements in a block
func (tc *TypeChecker) checkBlockForReturns(block *ast.BlockStatement, expectedType string) bool {
	if block == nil {
		return true
	}

	valid := true
	for _, stmt := range block.Statements {
		switch s := stmt.(type) {
		case *ast.ReturnStatement:
			if s.ReturnValue != nil {
				inferredType := tc.InferType(s.ReturnValue)
				if inferredType != "" && inferredType != expectedType {
					tc.addError(fmt.Sprintf(
						"return type mismatch: expected %s but got %s",
						expectedType, inferredType))
					valid = false
				}
			}
		case *ast.BlockStatement:
			if !tc.checkBlockForReturns(s, expectedType) {
				valid = false
			}
		}
	}

	return valid
}

// GetObjectType returns the type name of an object
func GetObjectType(obj object.Object) string {
	switch obj.Type() {
	case object.INTEGER_OBJ:
		return "পূর্ণসংখ্যা"
	case object.STRING_OBJ:
		return "লেখা"
	case object.BOOLEAN_OBJ:
		return "বুলিয়ান"
	case object.ARRAY_OBJ:
		return "তালিকা"
	case object.HASH_OBJ:
		return "হ্যাশ"
	case object.FUNCTION_OBJ, object.COMPILED_FUNCTION_OBJ:
		return "ফাংশন_টাইপ"
	default:
		return ""
	}
}

// CheckAssignment checks type consistency in assignments
func (tc *TypeChecker) CheckAssignment(varType string, value ast.Expression) bool {
	if varType == "" {
		// No type annotation, no checking needed
		return true
	}

	inferredType := tc.InferType(value)
	if inferredType == "" {
		// Cannot infer type, skip checking
		return true
	}

	if varType != inferredType {
		tc.addError(fmt.Sprintf(
			"type mismatch: expected %s but assigned %s",
			varType, inferredType))
		return false
	}

	return true
}
