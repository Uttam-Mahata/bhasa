# Bhasa AST Quick Reference

## Node Hierarchy

```
Node (interface)
├── Statement (interface)
│   ├── LetStatement
│   ├── ReturnStatement
│   ├── ExpressionStatement
│   ├── AssignmentStatement
│   ├── ImportStatement
│   ├── BlockStatement
│   ├── WhileStatement
│   ├── ForStatement
│   ├── BreakStatement
│   ├── ContinueStatement
│   ├── MemberAssignmentStatement
│   ├── MethodDefinition
│   ├── ConstructorDefinition
│   ├── ClassDefinition
│   └── InterfaceDefinition
│
└── Expression (interface)
    ├── Identifier
    ├── IntegerLiteral
    ├── StringLiteral
    ├── Boolean
    ├── PrefixExpression
    ├── InfixExpression
    ├── IfExpression
    ├── FunctionLiteral
    ├── CallExpression
    ├── ArrayLiteral
    ├── IndexExpression
    ├── HashLiteral
    ├── TypeAnnotation
    ├── TypedIdentifier
    ├── TypeCastExpression
    ├── StructDefinition
    ├── StructLiteral
    ├── MemberAccessExpression
    ├── EnumDefinition
    ├── EnumValue
    ├── NewExpression
    ├── ThisExpression
    ├── SuperExpression
    └── MethodCallExpression
```

## Statement Types

| Type | Bengali Keyword | Description | Example |
|------|-----------------|-------------|---------|
| LetStatement | `ধরি` | Variable declaration | `ধরি x = 5;` |
| ReturnStatement | `ফেরত` | Return from function | `ফেরত x + 1;` |
| AssignmentStatement | - | Variable reassignment | `x = 10;` |
| ImportStatement | `অন্তর্ভুক্ত` | Import module | `অন্তর্ভুক্ত "গণিত";` |
| WhileStatement | `যতক্ষণ` | While loop | `যতক্ষণ (x < 10) { }` |
| ForStatement | `পর্যন্ত` | For loop | `পর্যন্ত (ধরি i=0; i<10; i=i+1) { }` |
| BreakStatement | `বিরতি` | Break from loop | `বিরতি;` |
| ContinueStatement | `চালিয়ে_যাও` | Continue loop | `চালিয়ে_যাও;` |

## Expression Types

| Type | Description | Example |
|------|-------------|---------|
| Identifier | Variable/function name | `x`, `myVar` |
| IntegerLiteral | Integer number | `42`, `-15` |
| StringLiteral | String value | `"হ্যালো"` |
| Boolean | Boolean value | `সত্য`, `মিথ্যা` |
| PrefixExpression | Prefix operator | `!true`, `-5` |
| InfixExpression | Binary operator | `5 + 3`, `x == y` |
| IfExpression | Conditional | `যদি (x > 5) { } নাহলে { }` |
| FunctionLiteral | Function definition | `ফাংশন(x, y) { ফেরত x + y; }` |
| CallExpression | Function call | `myFunc(1, 2)` |
| ArrayLiteral | Array | `[1, 2, 3]` |
| IndexExpression | Array/hash access | `arr[0]`, `hash["key"]` |
| HashLiteral | Hash map | `{"key": "value"}` |

## Type Annotations

| Bengali | English | Description |
|---------|---------|-------------|
| `পূর্ণসংখ্যা` | integer | Integer type |
| `দশমিক` | float | Floating-point type |
| `লেখা` | string | String type |
| `বুলিয়ান` | boolean | Boolean type |
| `বাইট` | byte | Byte type |
| `শূন্য` | void | Void/null type |
| `তালিকা<T>` | array<T> | Array type |
| `ম্যাপ<K,V>` | map<K,V> | Hash map type |

## Struct & Enum

### Struct Definition
```bhasa
ধরি ব্যক্তি = স্ট্রাক্ট {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা
};
```

### Struct Literal
```bhasa
ধরি p = ব্যক্তি{নাম: "রহিম", বয়স: 30};
```

### Member Access
```bhasa
p.নাম  // Access field
p.বয়স = 31;  // Assign field
```

### Enum Definition
```bhasa
ধরি দিক = গণনা {
    উত্তর,
    দক্ষিণ,
    পূর্ব,
    পশ্চিম
};
```

### Enum Value
```bhasa
ধরি dir = দিক.উত্তর;
```

## OOP Features

### Access Modifiers

| Bengali | English | Scope |
|---------|---------|-------|
| `সার্বজনীন` | public | Accessible everywhere |
| `ব্যক্তিগত` | private | Only within class |
| `সুরক্ষিত` | protected | Class and subclasses |

### Class Modifiers

| Bengali | English | Description |
|---------|---------|-------------|
| `বিমূর্ত` | abstract | Cannot be instantiated |
| `চূড়ান্ত` | final | Cannot be extended/overridden |
| `স্থির` | static | Class-level member |
| `পুনর্সংজ্ঞা` | override | Overrides parent method |

### Class Definition

```bhasa
শ্রেণী ব্যক্তি {
    সার্বজনীন নাম: লেখা;
    ব্যক্তিগত বয়স: পূর্ণসংখ্যা;
    
    সার্বজনীন নির্মাতা(নাম: লেখা, বয়স: পূর্ণসংখ্যা) {
        এই.নাম = নাম;
        এই.বয়স = বয়স;
    }
    
    সার্বজনীন পদ্ধতি বলো(): শূন্য {
        দেখাও("হ্যালো");
    }
}
```

### Inheritance

```bhasa
শ্রেণী ছাত্র প্রসারিত ব্যক্তি {
    সার্বজনীন রোল: পূর্ণসংখ্যা;
    
    পুনর্সংজ্ঞা পদ্ধতি বলো(): শূন্য {
        উর্ধ্ব.বলো();  // Call parent method
        দেখাও("আমি ছাত্র");
    }
}
```

### Interface

```bhasa
চুক্তি যোগাযোগ {
    পদ্ধতি বলো(বার্তা: লেখা): শূন্য;
    পদ্ধতি শুনো(): লেখা;
}

শ্রেণী বার্তাপ্রেরক বাস্তবায়ন যোগাযোগ {
    সার্বজনীন পদ্ধতি বলো(বার্তা: লেখা): শূন্য {
        দেখাও(বার্তা);
    }
    
    সার্বজনীন পদ্ধতি শুনো(): লেখা {
        ফেরত "হ্যালো";
    }
}
```

### Creating Instances

```bhasa
ধরি p = নতুন ব্যক্তি("রহিম", 30);
p.বলো();
```

### This and Super

```bhasa
এই.নাম          // Access current object
এই.পদ্ধতি()      // Call method on current object
উর্ধ্ব.পদ্ধতি()   // Call parent class method
```

## Common Patterns

### Function with Type Annotations
```bhasa
ধরি যোগ = ফাংশন(a: পূর্ণসংখ্যা, b: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত a + b;
};
```

### Conditional Expression
```bhasa
ধরি result = যদি (x > 5) { "বড়" } নাহলে { "ছোট" };
```

### Array Iteration
```bhasa
ধরি numbers = [1, 2, 3, 4, 5];
পর্যন্ত (ধরি i = 0; i < len(numbers); i = i + 1) {
    দেখাও(numbers[i]);
}
```

### Hash Map
```bhasa
ধরি person: ম্যাপ<লেখা, লেখা> = {
    "নাম": "রহিম",
    "শহর": "ঢাকা"
};
```

### Type Casting
```bhasa
ধরি x = 5;
ধরি y = x as দশমিক;
```

## AST Node Structure

### Every Node Has:
1. **Token**: Original source token
2. **TokenLiteral()**: Returns token's literal value
3. **String()**: Converts back to source code

### Statements Have:
- **statementNode()**: Marker method

### Expressions Have:
- **expressionNode()**: Marker method

## Key AST Fields

### LetStatement
- `Name`: Variable identifier
- `TypeAnnot`: Optional type
- `Value`: Initial value expression

### FunctionLiteral
- `Parameters`: Parameter names
- `ParameterTypes`: Parameter types
- `ReturnType`: Return type
- `Body`: Function body block

### ClassDefinition
- `Name`: Class name
- `IsAbstract`: Abstract flag
- `IsFinal`: Final flag
- `SuperClass`: Parent class
- `Interfaces`: Implemented interfaces
- `Fields`: Class fields
- `Constructors`: Constructor definitions
- `Methods`: Method definitions

### MethodDefinition
- `Name`: Method name
- `Access`: Access modifier
- `IsStatic`: Static flag
- `IsFinal`: Final flag
- `IsAbstract`: Abstract flag
- `IsOverride`: Override flag
- `Parameters`: Parameter names
- `ParameterTypes`: Parameter types
- `ReturnType`: Return type
- `Body`: Method body

## Testing Helpers

### Create Identifier
```go
&ast.Identifier{
    Token: token.Token{Type: token.IDENT, Literal: "x"},
    Value: "x",
}
```

### Create Integer
```go
&ast.IntegerLiteral{
    Token: token.Token{Type: token.INT, Literal: "5"},
    Value: 5,
}
```

### Create Let Statement
```go
&ast.LetStatement{
    Token: token.Token{Type: token.LET, Literal: "ধরি"},
    Name: &ast.Identifier{...},
    Value: &ast.IntegerLiteral{...},
}
```

## Common Operations

### Walk AST
```go
func Walk(node ast.Node) {
    switch n := node.(type) {
    case *ast.Program:
        for _, stmt := range n.Statements {
            Walk(stmt)
        }
    case *ast.LetStatement:
        Walk(n.Value)
    // ... handle other types
    }
}
```

### Find All Identifiers
```go
func FindIdentifiers(node ast.Node) []*ast.Identifier {
    var identifiers []*ast.Identifier
    // Traverse AST and collect identifiers
    return identifiers
}
```

### Type Check
```go
func TypeCheck(node ast.Expression, expected string) error {
    // Validate expression type matches expected type
}
```

## See Also

- [Full AST Documentation](./ast-documentation.md)
- [Parser Documentation](../../parser/docs/)
- [Evaluator Documentation](../../evaluator/docs/)

