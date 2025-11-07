package compiler

import (
	"bhasa/ast"
	"bhasa/lexer"
	"bhasa/parser"
	"strings"
	"testing"
)

func TestTypeChecking(t *testing.T) {
	tests := []struct {
		input         string
		expectedError bool
		errorContains string
	}{
		{
			// Correct type
			`ধরি x: পূর্ণসংখ্যা = ৫;`,
			false,
			"",
		},
		{
			// Type mismatch
			`ধরি x: পূর্ণসংখ্যা = "string";`,
			true,
			"type mismatch",
		},
		{
			// Correct string type
			`ধরি name: লেখা = "test";`,
			false,
			"",
		},
		{
			// Type mismatch for string
			`ধরি name: লেখা = ১২৩;`,
			true,
			"type mismatch",
		},
		{
			// Correct boolean type
			`ধরি flag: বুলিয়ান = সত্য;`,
			false,
			"",
		},
		{
			// Type mismatch for boolean
			`ধরি flag: বুলিয়ান = ৫;`,
			true,
			"type mismatch",
		},
		{
			// No type annotation (should pass)
			`ধরি x = "anything";`,
			false,
			"",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Fatalf("parser errors: %v", p.Errors())
		}

		comp := New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compilation error: %v", err)
		}

		typeErrors := comp.TypeErrors()
		hasError := len(typeErrors) > 0

		if hasError != tt.expectedError {
			t.Errorf("input %q: expected error=%v, got=%v (errors: %v)",
				tt.input, tt.expectedError, hasError, typeErrors)
		}

		if tt.expectedError && len(typeErrors) > 0 {
			found := false
			for _, errMsg := range typeErrors {
				if strings.Contains(errMsg, tt.errorContains) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("input %q: expected error containing %q, got %v",
					tt.input, tt.errorContains, typeErrors)
			}
		}
	}
}

func TestFunctionReturnTypeChecking(t *testing.T) {
	tests := []struct {
		input         string
		expectedError bool
	}{
		{
			// Correct return type
			`ধরি f = ফাংশন(x: পূর্ণসংখ্যা): পূর্ণসংখ্যা { ফেরত x + ১; };`,
			false,
		},
		{
			// Return type mismatch
			`ধরি f = ফাংশন(x: পূর্ণসংখ্যা): পূর্ণসংখ্যা { ফেরত সত্য; };`,
			true,
		},
		{
			// Correct boolean return
			`ধরি f = ফাংশন(x: পূর্ণসংখ্যা): বুলিয়ান { ফেরত x > ০; };`,
			false,
		},
		{
			// No return type (should pass)
			`ধরি f = ফাংশন(x) { ফেরত x; };`,
			false,
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			t.Fatalf("parser errors: %v", p.Errors())
		}

		comp := New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compilation error: %v", err)
		}

		typeErrors := comp.TypeErrors()
		hasError := len(typeErrors) > 0

		if hasError != tt.expectedError {
			t.Errorf("input %q: expected error=%v, got=%v (errors: %v)",
				tt.input, tt.expectedError, hasError, typeErrors)
		}
	}
}

func TestTypeInference(t *testing.T) {
	tc := NewTypeChecker()

	tests := []struct {
		expr         ast.Expression
		expectedType string
	}{
		{
			&ast.IntegerLiteral{Value: 5},
			"পূর্ণসংখ্যা",
		},
		{
			&ast.StringLiteral{Value: "test"},
			"লেখা",
		},
		{
			&ast.Boolean{Value: true},
			"বুলিয়ান",
		},
		{
			&ast.ArrayLiteral{},
			"তালিকা",
		},
		{
			&ast.HashLiteral{},
			"হ্যাশ",
		},
	}

	for _, tt := range tests {
		inferredType := tc.InferType(tt.expr)
		if inferredType != tt.expectedType {
			t.Errorf("InferType(%T) = %q, want %q",
				tt.expr, inferredType, tt.expectedType)
		}
	}
}
