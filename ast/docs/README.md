# AST Documentation

Welcome to the Bhasa Abstract Syntax Tree (AST) documentation!

## 📚 Documentation Files

### [AST Documentation](./ast-documentation.md)
Comprehensive documentation covering:
- Complete overview of the AST structure
- Detailed explanation of every node type
- Bengali keyword reference
- Code examples for all features
- Best practices for AST manipulation

**Recommended for**: Deep understanding of the AST, implementing new features, working with the parser/evaluator

### [Quick Reference](./quick-reference.md)
Concise reference guide with:
- Node hierarchy diagram
- Quick lookup tables
- Common patterns and examples
- Testing helpers
- Common operations

**Recommended for**: Quick lookups, day-to-day development, cheat sheet

---

## 🎯 What is the AST?

The Abstract Syntax Tree (AST) is the core data structure that represents parsed Bhasa programs. It's a tree representation where:

- **Nodes** represent language constructs (variables, functions, classes, etc.)
- **Edges** represent relationships (a function contains statements, an expression has operands)
- **Tokens** are preserved for error reporting and source location tracking

### Flow

```
Source Code → Lexer → Tokens → Parser → AST → Evaluator/Compiler → Execution
```

---

## 🏗️ AST Structure

### Three Main Components

1. **Node Interface** - Base interface for all AST nodes
2. **Statement Interface** - Executable statements (declarations, loops, etc.)
3. **Expression Interface** - Evaluable expressions (literals, operations, etc.)

### Example

**Bhasa Code**:
```bhasa
ধরি x = 5 + 3;
```

**AST Representation**:
```
Program
└── LetStatement
    ├── Name: Identifier("x")
    └── Value: InfixExpression
        ├── Left: IntegerLiteral(5)
        ├── Operator: "+"
        └── Right: IntegerLiteral(3)
```

---

## 🔑 Key Features

### ✅ Comprehensive Language Support

- **Variables & Functions**: Let statements, function literals, calls
- **Control Flow**: If-else, while, for, break, continue
- **Data Structures**: Arrays, hash maps, structs, enums
- **Type System**: Type annotations, type casting, generic collections
- **OOP**: Classes, interfaces, inheritance, polymorphism, access control

### ✅ Bengali-First Design

All keywords and built-in types use Bengali terms:
- `ধরি` (let), `ফেরত` (return), `যদি` (if)
- `পূর্ণসংখ্যা` (integer), `লেখা` (string)
- `শ্রেণী` (class), `পদ্ধতি` (method)

### ✅ Rich Metadata

Every node includes:
- Original token (for error messages)
- Source location (line and column)
- String representation (for debugging)

---

## 📖 Getting Started

### For New Developers

1. Start with the [Quick Reference](./quick-reference.md) to get familiar with node types
2. Read the [AST Documentation](./ast-documentation.md) sections relevant to your task
3. Look at examples in the test files (`ast_test.go`)

### For Feature Development

1. Read the relevant section in [AST Documentation](./ast-documentation.md)
2. Check if similar node types exist that you can model after
3. Implement `Node`, `statementNode()` or `expressionNode()`, and `String()` methods
4. Add tests for your new node type

### For Parser Development

1. Understand the node types you need to create in [AST Documentation](./ast-documentation.md)
2. Reference the `String()` output to verify your parser produces correct AST
3. Use token preservation for accurate error reporting

### For Evaluator Development

1. Study the node structure in [AST Documentation](./ast-documentation.md)
2. Implement evaluation logic for each node type
3. Handle nil cases for optional fields (type annotations, else clauses)

---

## 📋 Node Type Categories

### Basic Statements
- `LetStatement` - Variable declarations
- `ReturnStatement` - Return values
- `AssignmentStatement` - Variable updates
- `ExpressionStatement` - Expression as statement

### Control Flow
- `IfExpression` - Conditionals
- `WhileStatement` - While loops
- `ForStatement` - For loops
- `BreakStatement` - Break from loops
- `ContinueStatement` - Skip iteration

### Expressions
- **Literals**: Integer, String, Boolean, Array, Hash
- **Operations**: Prefix, Infix
- **Functions**: FunctionLiteral, CallExpression
- **Access**: Identifier, IndexExpression, MemberAccessExpression

### Data Types
- `StructDefinition` - Define struct types
- `StructLiteral` - Create struct instances
- `EnumDefinition` - Define enum types
- `EnumValue` - Access enum variants
- `TypeAnnotation` - Type information

### Object-Oriented
- `ClassDefinition` - Define classes
- `InterfaceDefinition` - Define interfaces
- `MethodDefinition` - Define methods
- `ConstructorDefinition` - Define constructors
- `NewExpression` - Create instances
- `ThisExpression` - Current object reference
- `SuperExpression` - Parent class reference
- `MethodCallExpression` - Call methods

---

## 🧪 Testing

### Unit Tests

Tests are located in `ast_test.go`. They verify:
- Node creation
- String representation
- Token literal values

### Example Test

```go
func TestLetStatement(t *testing.T) {
    stmt := &ast.LetStatement{
        Token: token.Token{Type: token.LET, Literal: "ধরি"},
        Name: &ast.Identifier{
            Token: token.Token{Type: token.IDENT, Literal: "x"},
            Value: "x",
        },
        Value: &ast.IntegerLiteral{
            Token: token.Token{Type: token.INT, Literal: "5"},
            Value: 5,
        },
    }
    
    if stmt.String() != "ধরি x = 5;" {
        t.Errorf("stmt.String() wrong. got=%q", stmt.String())
    }
}
```

---

## 💡 Common Tasks

### Adding a New Statement Type

1. Define the struct with required fields
2. Implement `statementNode()`, `TokenLiteral()`, `String()`
3. Add to parser logic
4. Add evaluator logic
5. Write tests

### Adding a New Expression Type

1. Define the struct with required fields
2. Implement `expressionNode()`, `TokenLiteral()`, `String()`
3. Add to parser logic (with appropriate precedence)
4. Add evaluator logic
5. Write tests

### Modifying Existing Nodes

1. Update the struct definition
2. Update `String()` method if needed
3. Update parser logic
4. Update evaluator logic
5. Update and add tests
6. Update documentation

---

## 🔍 Debugging Tips

### Print AST Structure

```go
fmt.Println(program.String())
```

### Inspect Specific Nodes

```go
for i, stmt := range program.Statements {
    fmt.Printf("Statement %d: %T - %s\n", i, stmt, stmt.String())
}
```

### Validate Node Types

```go
if letStmt, ok := stmt.(*ast.LetStatement); ok {
    // Work with LetStatement
}
```

### Check for Nil Fields

```go
if node.TypeAnnot != nil {
    fmt.Println("Has type annotation:", node.TypeAnnot.String())
}
```

---

## 📚 Related Documentation

- **Lexer**: How tokens are generated from source code
- **Parser**: How AST is built from tokens  
- **Evaluator**: How AST is interpreted/executed
- **Token Package**: Token types and structure

---

## 🤝 Contributing

When adding features to the AST:

1. ✅ Follow existing naming conventions
2. ✅ Always implement all interface methods
3. ✅ Preserve tokens for error reporting
4. ✅ Write comprehensive tests
5. ✅ Update this documentation
6. ✅ Add examples for new features

---

## 📞 Need Help?

- Check [AST Documentation](./ast-documentation.md) for detailed explanations
- Check [Quick Reference](./quick-reference.md) for quick lookups
- Look at test files for usage examples
- Review existing similar node types

---

## 🗺️ AST Visual Overview

```
                         Node
                          │
        ┌─────────────────┴─────────────────┐
        │                                   │
    Statement                          Expression
        │                                   │
        ├─ LetStatement                     ├─ Identifier
        ├─ ReturnStatement                  ├─ Literals
        ├─ AssignmentStatement              │  ├─ Integer
        ├─ ImportStatement                  │  ├─ String
        ├─ BlockStatement                   │  ├─ Boolean
        ├─ WhileStatement                   │  ├─ Array
        ├─ ForStatement                     │  └─ Hash
        ├─ BreakStatement                   │
        ├─ ContinueStatement                ├─ Operations
        ├─ MemberAssignmentStatement        │  ├─ Prefix
        │                                   │  └─ Infix
        ├─ OOP Statements                   │
        │  ├─ ClassDefinition                ├─ Control Flow
        │  ├─ InterfaceDefinition            │  └─ IfExpression
        │  ├─ MethodDefinition               │
        │  └─ ConstructorDefinition          ├─ Functions
        │                                   │  ├─ FunctionLiteral
        └─ ExpressionStatement               │  └─ CallExpression
                                            │
                                            ├─ Structs/Enums
                                            │  ├─ StructDefinition
                                            │  ├─ StructLiteral
                                            │  ├─ EnumDefinition
                                            │  └─ EnumValue
                                            │
                                            ├─ Access
                                            │  ├─ IndexExpression
                                            │  └─ MemberAccessExpression
                                            │
                                            └─ OOP Expressions
                                               ├─ NewExpression
                                               ├─ ThisExpression
                                               ├─ SuperExpression
                                               └─ MethodCallExpression
```

---

**Happy Coding! 🚀**

