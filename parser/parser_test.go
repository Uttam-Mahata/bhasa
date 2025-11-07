package parser

import (
	"bhasa/ast"
	"bhasa/lexer"
	"testing"
)

func TestTypeAnnotations(t *testing.T) {
	tests := []struct {
		input              string
		expectedVarName    string
		expectedType       string
		expectedValue      interface{}
	}{
		{"ধরি x: পূর্ণসংখ্যা = ৫;", "x", "পূর্ণসংখ্যা", int64(5)},
		{"ধরি name: লেখা = \"test\";", "name", "লেখা", "test"},
		{"ধরি flag: বুলিয়ান = সত্য;", "flag", "বুলিয়ান", true},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statement. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		letStmt, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Fatalf("stmt not *ast.LetStatement. got=%T", stmt)
		}

		if letStmt.Name.Value != tt.expectedVarName {
			t.Errorf("letStmt.Name.Value not '%s'. got=%s",
				tt.expectedVarName, letStmt.Name.Value)
		}

		if letStmt.TypeAnnotation == nil {
			t.Fatalf("letStmt.TypeAnnotation is nil")
		}

		if letStmt.TypeAnnotation.Type != tt.expectedType {
			t.Errorf("letStmt.TypeAnnotation.Type not '%s'. got=%s",
				tt.expectedType, letStmt.TypeAnnotation.Type)
		}
	}
}

func TestFunctionTypeAnnotations(t *testing.T) {
	input := `ধরি f = ফাংশন(x: পূর্ণসংখ্যা, y: লেখা): বুলিয়ান { ফেরত সত্য; };`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("program.Statements[0] not *ast.LetStatement. got=%T",
			program.Statements[0])
	}

	function, ok := stmt.Value.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt.Value not *ast.FunctionLiteral. got=%T", stmt.Value)
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d",
			len(function.Parameters))
	}

	// Check first parameter type annotation
	if function.Parameters[0].TypeAnnotation == nil {
		t.Fatalf("function.Parameters[0].TypeAnnotation is nil")
	}
	if function.Parameters[0].TypeAnnotation.Type != "পূর্ণসংখ্যা" {
		t.Errorf("function.Parameters[0].TypeAnnotation.Type wrong. got=%s",
			function.Parameters[0].TypeAnnotation.Type)
	}

	// Check second parameter type annotation
	if function.Parameters[1].TypeAnnotation == nil {
		t.Fatalf("function.Parameters[1].TypeAnnotation is nil")
	}
	if function.Parameters[1].TypeAnnotation.Type != "লেখা" {
		t.Errorf("function.Parameters[1].TypeAnnotation.Type wrong. got=%s",
			function.Parameters[1].TypeAnnotation.Type)
	}

	// Check return type annotation
	if function.ReturnType == nil {
		t.Fatalf("function.ReturnType is nil")
	}
	if function.ReturnType.Type != "বুলিয়ান" {
		t.Errorf("function.ReturnType.Type wrong. got=%s",
			function.ReturnType.Type)
	}
}

func TestOptionalTypeAnnotations(t *testing.T) {
	// Test that code without type annotations still works
	input := `ধরি x = ৫;
	ধরি f = ফাংশন(a, b) { ফেরত a + b; };`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}

	// Check first statement (variable without type)
	letStmt := program.Statements[0].(*ast.LetStatement)
	if letStmt.TypeAnnotation != nil {
		t.Errorf("letStmt.TypeAnnotation should be nil for untyped variable")
	}

	// Check second statement (function without types)
	letStmt2 := program.Statements[1].(*ast.LetStatement)
	function := letStmt2.Value.(*ast.FunctionLiteral)
	
	if function.Parameters[0].TypeAnnotation != nil {
		t.Errorf("function parameter should have no type annotation")
	}
	
	if function.ReturnType != nil {
		t.Errorf("function should have no return type annotation")
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
