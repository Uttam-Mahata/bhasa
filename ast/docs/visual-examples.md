# AST Visual Examples

This document provides visual representations of how Bhasa code maps to AST structures.

## 📖 Table of Contents

1. [Basic Statements](#basic-statements)
2. [Expressions](#expressions)
3. [Functions](#functions)
4. [Control Flow](#control-flow)
5. [Data Structures](#data-structures)
6. [Structs and Enums](#structs-and-enums)
7. [Object-Oriented Programming](#object-oriented-programming)
8. [Complex Examples](#complex-examples)

---

## Basic Statements

### Variable Declaration

**Code**:
```bhasa
ধরি x = 5;
```

**AST**:
```
Program
└── LetStatement
    ├── Token: {Type: LET, Literal: "ধরি"}
    ├── Name: Identifier
    │   ├── Token: {Type: IDENT, Literal: "x"}
    │   └── Value: "x"
    ├── TypeAnnot: nil
    └── Value: IntegerLiteral
        ├── Token: {Type: INT, Literal: "5"}
        └── Value: 5

String(): "ধরি x = 5;"
```

---

### Variable with Type Annotation

**Code**:
```bhasa
ধরি নাম: লেখা = "বাংলা";
```

**AST**:
```
Program
└── LetStatement
    ├── Token: {Type: LET, Literal: "ধরি"}
    ├── Name: Identifier
    │   ├── Token: {Type: IDENT, Literal: "নাম"}
    │   └── Value: "নাম"
    ├── TypeAnnot: TypeAnnotation
    │   ├── Token: {Type: TYPE_STRING, Literal: "লেখা"}
    │   ├── TypeName: "লেখা"
    │   ├── ElementType: nil
    │   └── KeyType: nil
    └── Value: StringLiteral
        ├── Token: {Type: STRING, Literal: "বাংলা"}
        └── Value: "বাংলা"

String(): "ধরি নাম: লেখা = \"বাংলা\";"
```

---

### Assignment Statement

**Code**:
```bhasa
x = 10;
```

**AST**:
```
Program
└── AssignmentStatement
    ├── Token: {Type: IDENT, Literal: "x"}
    ├── Name: Identifier
    │   ├── Token: {Type: IDENT, Literal: "x"}
    │   └── Value: "x"
    └── Value: IntegerLiteral
        ├── Token: {Type: INT, Literal: "10"}
        └── Value: 10

String(): "x = 10;"
```

---

### Return Statement

**Code**:
```bhasa
ফেরত x + 5;
```

**AST**:
```
Program
└── ReturnStatement
    ├── Token: {Type: RETURN, Literal: "ফেরত"}
    └── ReturnValue: InfixExpression
        ├── Token: {Type: PLUS, Literal: "+"}
        ├── Left: Identifier
        │   └── Value: "x"
        ├── Operator: "+"
        └── Right: IntegerLiteral
            └── Value: 5

String(): "ফেরত (x + 5);"
```

---

## Expressions

### Arithmetic Expression

**Code**:
```bhasa
(5 + 3) * 2
```

**AST**:
```
InfixExpression
├── Token: {Type: ASTERISK, Literal: "*"}
├── Left: InfixExpression
│   ├── Token: {Type: PLUS, Literal: "+"}
│   ├── Left: IntegerLiteral
│   │   └── Value: 5
│   ├── Operator: "+"
│   └── Right: IntegerLiteral
│       └── Value: 3
├── Operator: "*"
└── Right: IntegerLiteral
    └── Value: 2

String(): "((5 + 3) * 2)"
```

---

### Comparison Expression

**Code**:
```bhasa
x > 10
```

**AST**:
```
InfixExpression
├── Token: {Type: GT, Literal: ">"}
├── Left: Identifier
│   └── Value: "x"
├── Operator: ">"
└── Right: IntegerLiteral
    └── Value: 10

String(): "(x > 10)"
```

---

### Prefix Expression

**Code**:
```bhasa
!সত্য
```

**AST**:
```
PrefixExpression
├── Token: {Type: BANG, Literal: "!"}
├── Operator: "!"
└── Right: Boolean
    ├── Token: {Type: TRUE, Literal: "সত্য"}
    └── Value: true

String(): "(!সত্য)"
```

---

## Functions

### Simple Function

**Code**:
```bhasa
ফাংশন(x, y) {
    ফেরত x + y;
}
```

**AST**:
```
FunctionLiteral
├── Token: {Type: FUNCTION, Literal: "ফাংশন"}
├── Parameters: []
│   ├── Identifier {Value: "x"}
│   └── Identifier {Value: "y"}
├── ParameterTypes: nil
├── ReturnType: nil
└── Body: BlockStatement
    └── Statements: []
        └── ReturnStatement
            └── ReturnValue: InfixExpression
                ├── Left: Identifier {Value: "x"}
                ├── Operator: "+"
                └── Right: Identifier {Value: "y"}

String(): "ফাংশন(x, y) (x + y)"
```

---

### Typed Function

**Code**:
```bhasa
ফাংশন(x: পূর্ণসংখ্যা, y: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত x * y;
}
```

**AST**:
```
FunctionLiteral
├── Token: {Type: FUNCTION, Literal: "ফাংশন"}
├── Parameters: []
│   ├── Identifier {Value: "x"}
│   └── Identifier {Value: "y"}
├── ParameterTypes: []
│   ├── TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
│   └── TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
├── ReturnType: TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
└── Body: BlockStatement
    └── Statements: []
        └── ReturnStatement
            └── ReturnValue: InfixExpression
                ├── Left: Identifier {Value: "x"}
                ├── Operator: "*"
                └── Right: Identifier {Value: "y"}

String(): "ফাংশন(x: পূর্ণসংখ্যা, y: পূর্ণসংখ্যা): পূর্ণসংখ্যা (x * y)"
```

---

### Function Call

**Code**:
```bhasa
যোগ(5, 10)
```

**AST**:
```
CallExpression
├── Token: {Type: LPAREN, Literal: "("}
├── Function: Identifier
│   └── Value: "যোগ"
└── Arguments: []
    ├── IntegerLiteral {Value: 5}
    └── IntegerLiteral {Value: 10}

String(): "যোগ(5, 10)"
```

---

## Control Flow

### If-Else Expression

**Code**:
```bhasa
যদি (x > 5) {
    দেখাও("বড়");
} নাহলে {
    দেখাও("ছোট");
}
```

**AST**:
```
IfExpression
├── Token: {Type: IF, Literal: "যদি"}
├── Condition: InfixExpression
│   ├── Left: Identifier {Value: "x"}
│   ├── Operator: ">"
│   └── Right: IntegerLiteral {Value: 5}
├── Consequence: BlockStatement
│   └── Statements: []
│       └── ExpressionStatement
│           └── Expression: CallExpression
│               ├── Function: Identifier {Value: "দেখাও"}
│               └── Arguments: []
│                   └── StringLiteral {Value: "বড়"}
└── Alternative: BlockStatement
    └── Statements: []
        └── ExpressionStatement
            └── Expression: CallExpression
                ├── Function: Identifier {Value: "দেখাও"}
                └── Arguments: []
                    └── StringLiteral {Value: "ছোট"}

String(): "if(x > 5) দেখাও(\"বড়\")else দেখাও(\"ছোট\")"
```

---

### While Loop

**Code**:
```bhasa
যতক্ষণ (i < 10) {
    i = i + 1;
}
```

**AST**:
```
WhileStatement
├── Token: {Type: WHILE, Literal: "যতক্ষণ"}
├── Condition: InfixExpression
│   ├── Left: Identifier {Value: "i"}
│   ├── Operator: "<"
│   └── Right: IntegerLiteral {Value: 10}
└── Body: BlockStatement
    └── Statements: []
        └── AssignmentStatement
            ├── Name: Identifier {Value: "i"}
            └── Value: InfixExpression
                ├── Left: Identifier {Value: "i"}
                ├── Operator: "+"
                └── Right: IntegerLiteral {Value: 1}

String(): "while (i < 10) i = (i + 1);"
```

---

### For Loop

**Code**:
```bhasa
পর্যন্ত (ধরি i = 0; i < 5; i = i + 1) {
    দেখাও(i);
}
```

**AST**:
```
ForStatement
├── Token: {Type: FOR, Literal: "পর্যন্ত"}
├── Init: LetStatement
│   ├── Name: Identifier {Value: "i"}
│   └── Value: IntegerLiteral {Value: 0}
├── Condition: InfixExpression
│   ├── Left: Identifier {Value: "i"}
│   ├── Operator: "<"
│   └── Right: IntegerLiteral {Value: 5}
├── Increment: AssignmentStatement
│   ├── Name: Identifier {Value: "i"}
│   └── Value: InfixExpression
│       ├── Left: Identifier {Value: "i"}
│       ├── Operator: "+"
│       └── Right: IntegerLiteral {Value: 1}
└── Body: BlockStatement
    └── Statements: []
        └── ExpressionStatement
            └── Expression: CallExpression
                ├── Function: Identifier {Value: "দেখাও"}
                └── Arguments: []
                    └── Identifier {Value: "i"}

String(): "for (ধরি i = 0;; (i < 5); i = (i + 1);) দেখাও(i)"
```

---

## Data Structures

### Array Literal

**Code**:
```bhasa
[1, 2, 3, 4, 5]
```

**AST**:
```
ArrayLiteral
├── Token: {Type: LBRACKET, Literal: "["}
└── Elements: []
    ├── IntegerLiteral {Value: 1}
    ├── IntegerLiteral {Value: 2}
    ├── IntegerLiteral {Value: 3}
    ├── IntegerLiteral {Value: 4}
    └── IntegerLiteral {Value: 5}

String(): "[1, 2, 3, 4, 5]"
```

---

### Array with Type Annotation

**Code**:
```bhasa
ধরি numbers: তালিকা<পূর্ণসংখ্যা> = [1, 2, 3];
```

**AST**:
```
LetStatement
├── Token: {Type: LET, Literal: "ধরি"}
├── Name: Identifier {Value: "numbers"}
├── TypeAnnot: TypeAnnotation
│   ├── TypeName: "তালিকা"
│   ├── ElementType: TypeAnnotation
│   │   ├── TypeName: "পূর্ণসংখ্যা"
│   │   ├── ElementType: nil
│   │   └── KeyType: nil
│   └── KeyType: nil
└── Value: ArrayLiteral
    └── Elements: [1, 2, 3]

String(): "ধরি numbers: তালিকা<পূর্ণসংখ্যা> = [1, 2, 3];"
```

---

### Array Indexing

**Code**:
```bhasa
array[2]
```

**AST**:
```
IndexExpression
├── Token: {Type: LBRACKET, Literal: "["}
├── Left: Identifier {Value: "array"}
└── Index: IntegerLiteral {Value: 2}

String(): "(array[2])"
```

---

### Hash Literal

**Code**:
```bhasa
{"নাম": "রহিম", "বয়স": 30}
```

**AST**:
```
HashLiteral
├── Token: {Type: LBRACE, Literal: "{"}
└── Pairs: map[Expression]Expression
    ├── StringLiteral{Value: "নাম"} => StringLiteral{Value: "রহিম"}
    └── StringLiteral{Value: "বয়স"} => IntegerLiteral{Value: 30}

String(): "{\"নাম\":\"রহিম\", \"বয়স\":30}"
```

---

### Hash with Type Annotation

**Code**:
```bhasa
ধরি person: ম্যাপ<লেখা, লেখা> = {"নাম": "করিম"};
```

**AST**:
```
LetStatement
├── Token: {Type: LET, Literal: "ধরি"}
├── Name: Identifier {Value: "person"}
├── TypeAnnot: TypeAnnotation
│   ├── TypeName: "ম্যাপ"
│   ├── KeyType: TypeAnnotation
│   │   └── TypeName: "লেখা"
│   └── ElementType: TypeAnnotation
│       └── TypeName: "লেখা"
└── Value: HashLiteral
    └── Pairs: {"নাম": "করিম"}

String(): "ধরি person: ম্যাপ<লেখা, লেখা> = {\"নাম\":\"করিম\"};"
```

---

## Structs and Enums

### Struct Definition

**Code**:
```bhasa
ধরি ব্যক্তি = স্ট্রাক্ট {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা
};
```

**AST**:
```
LetStatement
├── Token: {Type: LET, Literal: "ধরি"}
├── Name: Identifier {Value: "ব্যক্তি"}
└── Value: StructDefinition
    ├── Token: {Type: STRUCT, Literal: "স্ট্রাক্ট"}
    ├── Name: Identifier {Value: "ব্যক্তি"}
    └── Fields: []
        ├── StructField
        │   ├── Name: "নাম"
        │   └── TypeAnnot: TypeAnnotation {TypeName: "লেখা"}
        └── StructField
            ├── Name: "বয়স"
            └── TypeAnnot: TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}

String(): "ধরি ব্যক্তি = স্ট্রাক্ট {নাম: লেখা, বয়স: পূর্ণসংখ্যা};"
```

---

### Struct Literal

**Code**:
```bhasa
ব্যক্তি{নাম: "রহিম", বয়স: 30}
```

**AST**:
```
StructLiteral
├── Token: {Type: LBRACE, Literal: "{"}
├── StructType: Identifier {Value: "ব্যক্তি"}
├── Fields: map[string]Expression
│   ├── "নাম" => StringLiteral {Value: "রহিম"}
│   └── "বয়স" => IntegerLiteral {Value: 30}
└── FieldOrder: ["নাম", "বয়স"]

String(): "ব্যক্তি{নাম: \"রহিম\", বয়স: 30}"
```

---

### Member Access

**Code**:
```bhasa
person.নাম
```

**AST**:
```
MemberAccessExpression
├── Token: {Type: DOT, Literal: "."}
├── Object: Identifier {Value: "person"}
└── Member: Identifier {Value: "নাম"}

String(): "(person.নাম)"
```

---

### Member Assignment

**Code**:
```bhasa
person.বয়স = 31;
```

**AST**:
```
MemberAssignmentStatement
├── Token: {Type: IDENT, Literal: "person"}
├── Object: Identifier {Value: "person"}
├── Member: Identifier {Value: "বয়স"}
└── Value: IntegerLiteral {Value: 31}

String(): "person.বয়স = 31;"
```

---

### Enum Definition

**Code**:
```bhasa
ধরি দিক = গণনা {
    উত্তর,
    দক্ষিণ,
    পূর্ব,
    পশ্চিম
};
```

**AST**:
```
LetStatement
├── Token: {Type: LET, Literal: "ধরি"}
├── Name: Identifier {Value: "দিক"}
└── Value: EnumDefinition
    ├── Token: {Type: ENUM, Literal: "গণনা"}
    ├── Name: Identifier {Value: "দিক"}
    └── Variants: []
        ├── EnumVariant {Name: "উত্তর", Value: nil}
        ├── EnumVariant {Name: "দক্ষিণ", Value: nil}
        ├── EnumVariant {Name: "পূর্ব", Value: nil}
        └── EnumVariant {Name: "পশ্চিম", Value: nil}

String(): "ধরি দিক = গণনা {উত্তর, দক্ষিণ, পূর্ব, পশ্চিম};"
```

---

### Enum Value

**Code**:
```bhasa
দিক.উত্তর
```

**AST**:
```
EnumValue
├── Token: {Type: IDENT, Literal: "দিক"}
├── EnumType: Identifier {Value: "দিক"}
└── VariantName: Identifier {Value: "উত্তর"}

String(): "দিক.উত্তর"
```

---

## Object-Oriented Programming

### Simple Class

**Code**:
```bhasa
শ্রেণী ব্যক্তি {
    সার্বজনীন নাম: লেখা;
    ব্যক্তিগত বয়স: পূর্ণসংখ্যা;
}
```

**AST**:
```
ClassDefinition
├── Token: {Type: CLASS, Literal: "শ্রেণী"}
├── Name: Identifier {Value: "ব্যক্তি"}
├── IsAbstract: false
├── IsFinal: false
├── SuperClass: nil
├── Interfaces: []
├── Fields: []
│   ├── ClassField
│   │   ├── Name: "নাম"
│   │   ├── TypeAnnot: TypeAnnotation {TypeName: "লেখা"}
│   │   ├── Access: "সার্বজনীন"
│   │   ├── IsStatic: false
│   │   └── IsFinal: false
│   └── ClassField
│       ├── Name: "বয়স"
│       ├── TypeAnnot: TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
│       ├── Access: "ব্যক্তিগত"
│       ├── IsStatic: false
│       └── IsFinal: false
├── Constructors: []
└── Methods: []

String(): "শ্রেণী ব্যক্তি { ... }"
```

---

### Constructor

**Code**:
```bhasa
সার্বজনীন নির্মাতা(নাম: লেখা, বয়স: পূর্ণসংখ্যা) {
    এই.নাম = নাম;
    এই.বয়স = বয়স;
}
```

**AST**:
```
ConstructorDefinition
├── Token: {Type: CONSTRUCTOR, Literal: "নির্মাতা"}
├── Access: "সার্বজনীন"
├── Parameters: []
│   ├── Identifier {Value: "নাম"}
│   └── Identifier {Value: "বয়স"}
├── ParameterTypes: []
│   ├── TypeAnnotation {TypeName: "লেখা"}
│   └── TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
└── Body: BlockStatement
    └── Statements: []
        ├── MemberAssignmentStatement
        │   ├── Object: ThisExpression
        │   ├── Member: Identifier {Value: "নাম"}
        │   └── Value: Identifier {Value: "নাম"}
        └── MemberAssignmentStatement
            ├── Object: ThisExpression
            ├── Member: Identifier {Value: "বয়স"}
            └── Value: Identifier {Value: "বয়স"}

String(): "সার্বজনীন নির্মাতা(নাম: লেখা, বয়স: পূর্ণসংখ্যা) ..."
```

---

### Method Definition

**Code**:
```bhasa
সার্বজনীন পদ্ধতি বলো(): শূন্য {
    দেখাও("হ্যালো");
}
```

**AST**:
```
MethodDefinition
├── Token: {Type: METHOD, Literal: "পদ্ধতি"}
├── Name: Identifier {Value: "বলো"}
├── Access: "সার্বজনীন"
├── IsStatic: false
├── IsFinal: false
├── IsAbstract: false
├── IsOverride: false
├── Parameters: []
├── ParameterTypes: []
├── ReturnType: TypeAnnotation {TypeName: "শূন্য"}
└── Body: BlockStatement
    └── Statements: []
        └── ExpressionStatement
            └── Expression: CallExpression
                ├── Function: Identifier {Value: "দেখাও"}
                └── Arguments: []
                    └── StringLiteral {Value: "হ্যালো"}

String(): "সার্বজনীন পদ্ধতি বলো(): শূন্য দেখাও(\"হ্যালো\")"
```

---

### Class with Inheritance

**Code**:
```bhasa
শ্রেণী ছাত্র প্রসারিত ব্যক্তি {
    সার্বজনীন রোল: পূর্ণসংখ্যা;
}
```

**AST**:
```
ClassDefinition
├── Token: {Type: CLASS, Literal: "শ্রেণী"}
├── Name: Identifier {Value: "ছাত্র"}
├── IsAbstract: false
├── IsFinal: false
├── SuperClass: Identifier {Value: "ব্যক্তি"}
├── Interfaces: []
├── Fields: []
│   └── ClassField
│       ├── Name: "রোল"
│       ├── TypeAnnot: TypeAnnotation {TypeName: "পূর্ণসংখ্যা"}
│       ├── Access: "সার্বজনীন"
│       ├── IsStatic: false
│       └── IsFinal: false
├── Constructors: []
└── Methods: []

String(): "শ্রেণী ছাত্র প্রসারিত ব্যক্তি { ... }"
```

---

### Interface Definition

**Code**:
```bhasa
চুক্তি যোগাযোগ {
    পদ্ধতি বলো(বার্তা: লেখা): শূন্য;
}
```

**AST**:
```
InterfaceDefinition
├── Token: {Type: INTERFACE, Literal: "চুক্তি"}
├── Name: Identifier {Value: "যোগাযোগ"}
└── Methods: []
    └── InterfaceMethod
        ├── Name: Identifier {Value: "বলো"}
        ├── Parameters: []
        │   └── Identifier {Value: "বার্তা"}
        ├── ParameterTypes: []
        │   └── TypeAnnotation {TypeName: "লেখা"}
        └── ReturnType: TypeAnnotation {TypeName: "শূন্য"}

String(): "চুক্তি যোগাযোগ { ... }"
```

---

### New Expression

**Code**:
```bhasa
নতুন ব্যক্তি("রহিম", 30)
```

**AST**:
```
NewExpression
├── Token: {Type: NEW, Literal: "নতুন"}
├── ClassName: Identifier {Value: "ব্যক্তি"}
└── Arguments: []
    ├── StringLiteral {Value: "রহিম"}
    └── IntegerLiteral {Value: 30}

String(): "নতুন ব্যক্তি(\"রহিম\", 30)"
```

---

### Method Call

**Code**:
```bhasa
person.বলো("হ্যালো")
```

**AST**:
```
MethodCallExpression
├── Token: {Type: DOT, Literal: "."}
├── Object: Identifier {Value: "person"}
├── MethodName: Identifier {Value: "বলো"}
└── Arguments: []
    └── StringLiteral {Value: "হ্যালো"}

String(): "person.বলো(\"হ্যালো\")"
```

---

### This Expression

**Code**:
```bhasa
এই.নাম
```

**AST**:
```
MemberAccessExpression
├── Token: {Type: DOT, Literal: "."}
├── Object: ThisExpression
│   └── Token: {Type: THIS, Literal: "এই"}
└── Member: Identifier {Value: "নাম"}

String(): "(এই.নাম)"
```

---

### Super Method Call

**Code**:
```bhasa
উর্ধ্ব.বলো()
```

**AST**:
```
MethodCallExpression
├── Token: {Type: DOT, Literal: "."}
├── Object: SuperExpression
│   └── Token: {Type: SUPER, Literal: "উর্ধ্ব"}
├── MethodName: Identifier {Value: "বলো"}
└── Arguments: []

String(): "উর্ধ্ব.বলো()"
```

---

## Complex Examples

### Fibonacci Function

**Code**:
```bhasa
ধরি ফিবোনাচি = ফাংশন(n: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    যদি (n <= 1) {
        ফেরত n;
    } নাহলে {
        ফেরত ফিবোনাচি(n - 1) + ফিবোনাচি(n - 2);
    }
};
```

**AST**:
```
Program
└── LetStatement
    ├── Name: Identifier {Value: "ফিবোনাচি"}
    └── Value: FunctionLiteral
        ├── Parameters: [Identifier("n")]
        ├── ParameterTypes: [TypeAnnotation("পূর্ণসংখ্যা")]
        ├── ReturnType: TypeAnnotation("পূর্ণসংখ্যা")
        └── Body: BlockStatement
            └── Statements: [
                IfExpression
                ├── Condition: InfixExpression
                │   ├── Left: Identifier("n")
                │   ├── Operator: "<="
                │   └── Right: IntegerLiteral(1)
                ├── Consequence: BlockStatement
                │   └── ReturnStatement
                │       └── ReturnValue: Identifier("n")
                └── Alternative: BlockStatement
                    └── ReturnStatement
                        └── ReturnValue: InfixExpression
                            ├── Left: CallExpression
                            │   ├── Function: Identifier("ফিবোনাচি")
                            │   └── Arguments: [
                            │       InfixExpression(n - 1)
                            │   ]
                            ├── Operator: "+"
                            └── Right: CallExpression
                                ├── Function: Identifier("ফিবোনাচি")
                                └── Arguments: [
                                    InfixExpression(n - 2)
                                ]
            ]
```

---

### Complete OOP Example

**Code**:
```bhasa
বিমূর্ত শ্রেণী আকার {
    বিমূর্ত পদ্ধতি ক্ষেত্রফল(): দশমিক;
}

শ্রেণী বৃত্ত প্রসারিত আকার {
    সার্বজনীন ব্যাসার্ধ: দশমিক;
    
    সার্বজনীন নির্মাতা(r: দশমিক) {
        এই.ব্যাসার্ধ = r;
    }
    
    পুনর্সংজ্ঞা পদ্ধতি ক্ষেত্রফল(): দশমিক {
        ফেরত 3.14159 * এই.ব্যাসার্ধ * এই.ব্যাসার্ধ;
    }
}

ধরি c = নতুন বৃত্ত(5.0);
দেখাও(c.ক্ষেত্রফল());
```

**AST**:
```
Program
├── ClassDefinition (আকার)
│   ├── Name: Identifier("আকার")
│   ├── IsAbstract: true
│   └── Methods: [
│       MethodDefinition
│       ├── Name: Identifier("ক্ষেত্রফল")
│       ├── IsAbstract: true
│       └── ReturnType: TypeAnnotation("দশমিক")
│   ]
│
├── ClassDefinition (বৃত্ত)
│   ├── Name: Identifier("বৃত্ত")
│   ├── SuperClass: Identifier("আকার")
│   ├── Fields: [
│   │   ClassField
│   │   ├── Name: "ব্যাসার্ধ"
│   │   ├── TypeAnnot: TypeAnnotation("দশমিক")
│   │   └── Access: "সার্বজনীন"
│   ]
│   ├── Constructors: [
│   │   ConstructorDefinition
│   │   ├── Parameters: [Identifier("r")]
│   │   ├── ParameterTypes: [TypeAnnotation("দশমিক")]
│   │   └── Body: BlockStatement
│   │       └── MemberAssignmentStatement
│   │           ├── Object: ThisExpression
│   │           ├── Member: Identifier("ব্যাসার্ধ")
│   │           └── Value: Identifier("r")
│   ]
│   └── Methods: [
│       MethodDefinition
│       ├── Name: Identifier("ক্ষেত্রফল")
│       ├── IsOverride: true
│       ├── ReturnType: TypeAnnotation("দশমিক")
│       └── Body: BlockStatement
│           └── ReturnStatement
│               └── ReturnValue: InfixExpression
│                   ├── Left: InfixExpression
│                   │   ├── Left: FloatLiteral(3.14159)
│                   │   ├── Operator: "*"
│                   │   └── Right: MemberAccessExpression
│                   │       ├── Object: ThisExpression
│                   │       └── Member: Identifier("ব্যাসার্ধ")
│                   ├── Operator: "*"
│                   └── Right: MemberAccessExpression
│                       ├── Object: ThisExpression
│                       └── Member: Identifier("ব্যাসার্ধ")
│   ]
│
├── LetStatement (c)
│   └── Value: NewExpression
│       ├── ClassName: Identifier("বৃত্ত")
│       └── Arguments: [FloatLiteral(5.0)]
│
└── ExpressionStatement
    └── Expression: CallExpression
        ├── Function: Identifier("দেখাও")
        └── Arguments: [
            MethodCallExpression
            ├── Object: Identifier("c")
            ├── MethodName: Identifier("ক্ষেত্রফল")
            └── Arguments: []
        ]
```

---

## Summary

This document demonstrates how Bhasa source code is parsed into AST structures. Key takeaways:

1. **Every construct has a corresponding AST node type**
2. **Nodes preserve tokens for error reporting**
3. **String() method reconstructs source code**
4. **Tree structure reflects code structure**
5. **Type annotations are optional but fully supported**
6. **OOP features are comprehensively represented**

For more information, see:
- [AST Documentation](./ast-documentation.md) - Complete reference
- [Quick Reference](./quick-reference.md) - Lookup tables
- [README](./README.md) - Overview and guides

