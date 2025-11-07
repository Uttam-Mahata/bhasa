package ast

import (
	"bhasa/token"
	"bytes"
	"strings"
)

// Node represents a node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement node
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement represents a variable declaration (ধরি)
type LetStatement struct {
	Token      token.Token     // the ধরি token
	Name       *Identifier
	TypeAnnot  *TypeAnnotation // optional type annotation (can be nil)
	Value      Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	if ls.TypeAnnot != nil {
		out.WriteString(": ")
		out.WriteString(ls.TypeAnnot.String())
	}
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ReturnStatement represents a return statement (ফেরত)
type ReturnStatement struct {
	Token       token.Token // the ফেরত token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// AssignmentStatement represents a variable assignment (not declaration)
type AssignmentStatement struct {
	Token token.Token // the identifier token
	Name  *Identifier
	Value Expression
}

func (as *AssignmentStatement) statementNode()       {}
func (as *AssignmentStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.Name.String())
	out.WriteString(" = ")
	if as.Value != nil {
		out.WriteString(as.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// ImportStatement represents an import/include statement (অন্তর্ভুক্ত)
type ImportStatement struct {
	Token token.Token // the অন্তর্ভুক্ত token
	Path  Expression  // the module path (string literal)
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString(is.TokenLiteral() + " ")
	if is.Path != nil {
		out.WriteString(is.Path.String())
	}
	out.WriteString(";")
	return out.String()
}

// BlockStatement represents a block of statements
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// WhileStatement represents a while loop (যতক্ষণ)
type WhileStatement struct {
	Token     token.Token // the যতক্ষণ token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out bytes.Buffer
	out.WriteString("while ")
	out.WriteString(ws.Condition.String())
	out.WriteString(" ")
	out.WriteString(ws.Body.String())
	return out.String()
}

// Identifier represents an identifier
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// IntegerLiteral represents an integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// StringLiteral represents a string literal
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// Boolean represents a boolean value
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

// PrefixExpression represents a prefix expression (e.g., !true, -5)
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression represents an infix expression (e.g., 5 + 5)
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

// IfExpression represents an if-else expression
type IfExpression struct {
	Token       token.Token // The যদি token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// FunctionLiteral represents a function literal
type FunctionLiteral struct {
	Token          token.Token       // The ফাংশন token
	Parameters     []*Identifier     // Parameter names (for backward compatibility)
	ParameterTypes []*TypeAnnotation // Optional parameter type annotations (parallel to Parameters)
	ReturnType     *TypeAnnotation   // Optional return type annotation
	Body           *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for i, p := range fl.Parameters {
		paramStr := p.String()
		if fl.ParameterTypes != nil && i < len(fl.ParameterTypes) && fl.ParameterTypes[i] != nil {
			paramStr += ": " + fl.ParameterTypes[i].String()
		}
		params = append(params, paramStr)
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	if fl.ReturnType != nil {
		out.WriteString(": ")
		out.WriteString(fl.ReturnType.String())
	}
	out.WriteString(" ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// CallExpression represents a function call
type CallExpression struct {
	Token     token.Token // The ( token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// ArrayLiteral represents an array literal
type ArrayLiteral struct {
	Token    token.Token // the [ token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// IndexExpression represents an index expression (e.g., array[0])
type IndexExpression struct {
	Token token.Token // The [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

// HashLiteral represents a hash map literal
type HashLiteral struct {
	Token token.Token // the { token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// ForStatement represents a for loop (পর্যন্ত)
type ForStatement struct {
	Token       token.Token // the পর্যন্ত token
	Init        Statement   // initialization (can be LetStatement or AssignStatement)
	Condition   Expression
	Increment   Statement
	Body        *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("for (")
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")
	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")
	if fs.Increment != nil {
		out.WriteString(fs.Increment.String())
	}
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// BreakStatement represents a break statement (বিরতি)
type BreakStatement struct {
	Token token.Token // the বিরতি token
}

func (bs *BreakStatement) statementNode()       {}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string {
	return bs.TokenLiteral() + ";"
}

// ContinueStatement represents a continue statement (চালিয়ে_যাও)
type ContinueStatement struct {
	Token token.Token // the চালিয়ে_যাও token
}

func (cs *ContinueStatement) statementNode()       {}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string {
	return cs.TokenLiteral() + ";"
}

// TypeAnnotation represents a type annotation
type TypeAnnotation struct {
	Token      token.Token // The type token (e.g., পূর্ণসংখ্যা, বাইট, etc.)
	TypeName   string      // The name of the type
	ElementType *TypeAnnotation // For array/hash element types (e.g., তালিকা<পূর্ণসংখ্যা>)
	KeyType     *TypeAnnotation // For hash key types (e.g., ম্যাপ<লেখা, পূর্ণসংখ্যা>)
}

func (ta *TypeAnnotation) expressionNode()      {}
func (ta *TypeAnnotation) TokenLiteral() string { return ta.Token.Literal }
func (ta *TypeAnnotation) String() string {
	if ta.ElementType != nil && ta.KeyType != nil {
		// Hash type: ম্যাপ<লেখা, পূর্ণসংখ্যা>
		return ta.TypeName + "<" + ta.KeyType.String() + ", " + ta.ElementType.String() + ">"
	} else if ta.ElementType != nil {
		// Array type: তালিকা<পূর্ণসংখ্যা>
		return ta.TypeName + "<" + ta.ElementType.String() + ">"
	}
	return ta.TypeName
}

// TypedIdentifier represents an identifier with a type annotation
type TypedIdentifier struct {
	Token      token.Token     // The identifier token
	Name       string          // The identifier name
	TypeAnnot  *TypeAnnotation // The type annotation (can be nil for untyped)
}

func (ti *TypedIdentifier) expressionNode()      {}
func (ti *TypedIdentifier) TokenLiteral() string { return ti.Token.Literal }
func (ti *TypedIdentifier) String() string {
	if ti.TypeAnnot != nil {
		return ti.Name + ": " + ti.TypeAnnot.String()
	}
	return ti.Name
}

// TypeCastExpression represents a type cast (e.g., x as দশমিক)
type TypeCastExpression struct {
	Token      token.Token     // The 'as' token
	Expression Expression      // The expression to cast
	TargetType *TypeAnnotation // The target type
}

func (tce *TypeCastExpression) expressionNode()      {}
func (tce *TypeCastExpression) TokenLiteral() string { return tce.Token.Literal }
func (tce *TypeCastExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(tce.Expression.String())
	out.WriteString(" as ")
	out.WriteString(tce.TargetType.String())
	out.WriteString(")")
	return out.String()
}

