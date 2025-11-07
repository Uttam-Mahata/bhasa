# Struct and Enum Design for Bhasa

## Overview
This document outlines the design for adding structs (structured data types) and enums (enumerated types) to the Bhasa programming language, paving the way for future OOP features.

---

## 1. Struct Design

### Syntax Options

#### Option A: Using `স্ট্রাক্ট` (struct) keyword
```bhasa
// Define a struct type
ধরি ব্যক্তি = স্ট্রাক্ট {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা,
    ঠিকানা: লেখা
};

// Create instance
ধরি person = ব্যক্তি{
    নাম: "রহিম",
    বয়স: 30,
    ঠিকানা: "ঢাকা"
};

// Access fields
লেখ(person.নাম);
লেখ(person.বয়স);

// Modify fields
person.বয়স = 31;
```

#### Option B: Using `রেকর্ড` (record) keyword (More Bengali)
```bhasa
ধরি ব্যক্তি = রেকর্ড {
    নাম: লেখা,
    বয়স: পূর্ণসংখ্যা
};
```

**Decision: Use `স্ট্রাক্ট` (struct) - More familiar to programmers, clear intent**

### Features

1. **Field Declaration with Types**
   - All fields must have type annotations
   - Fields are ordered (maintain definition order)

2. **Struct Literals**
   ```bhasa
   ধরি p = ব্যক্তি{নাম: "করিম", বয়স: 25, ঠিকানা: "চট্টগ্রাম"};
   ```

3. **Field Access** (dot notation)
   ```bhasa
   person.নাম
   person.বয়স
   ```

4. **Field Assignment**
   ```bhasa
   person.নাম = "নতুন নাম";
   ```

5. **Nested Structs**
   ```bhasa
   ধরি ঠিকানা_ধরন = স্ট্রাক্ট {
       রাস্তা: লেখা,
       শহর: লেখা,
       জিপকোড: পূর্ণসংখ্যা
   };

   ধরি ব্যক্তি = স্ট্রাক্ট {
       নাম: লেখা,
       ঠিকানা: ঠিকানা_ধরন
   };

   ধরি p = ব্যক্তি{
       নাম: "রহিম",
       ঠিকানা: ঠিকানা_ধরন{
           রাস্তা: "মিরপুর",
           শহর: "ঢাকা",
           জিপকোড: 1216
       }
   };

   লেখ(p.ঠিকানা.শহর);  // "ঢাকা"
   ```

6. **Methods on Structs** (Go-style)
   ```bhasa
   // Method definition
   ধরি (p: ব্যক্তি) পরিচয় = ফাংশন(): লেখা {
       ফেরত "আমি " + p.নাম + ", বয়স " + লেখা(p.বয়স);
   };

   // Method call
   ধরি person = ব্যক্তি{নাম: "করিম", বয়স: 25};
   লেখ(person.পরিচয়());  // "আমি করিম, বয়স 25"
   ```

7. **Struct Comparison**
   ```bhasa
   ধরি p1 = ব্যক্তি{নাম: "রহিম", বয়স: 30};
   ধরি p2 = ব্যক্তি{নাম: "রহিম", বয়স: 30};

   যদি (p1 == p2) {
       লেখ("Same values");
   }
   ```

---

## 2. Enum Design

### Syntax

```bhasa
// Simple enum
ধরি দিক = enum {
    উত্তর,
    দক্ষিণ,
    পূর্ব,
    পশ্চিম
};

// Enum with explicit values
ধরি স্ট্যাটাস = enum {
    সফল = 0,
    ব্যর্থ = 1,
    অপেক্ষমান = 2
};

// Enum with associated data (like Rust)
ধরি ফলাফল = enum {
    সফল(মান: পূর্ণসংখ্যা),
    ত্রুটি(বার্তা: লেখা)
};
```

### Usage

```bhasa
// Simple enum usage
ধরি current_direction: দিক = দিক.উত্তর;

যদি (current_direction == দিক.উত্তর) {
    লেখ("Going north");
}

// Pattern matching (future feature)
যদি (current_direction) {
    দিক.উত্তর => লেখ("North"),
    দিক.দক্ষিণ => লেখ("South"),
    দিক.পূর্ব => লেখ("East"),
    দিক.পশ্চিম => লেখ("West")
}

// Enum with associated data
ধরি result = ফলাফল.সফল(মান: 42);

// Pattern matching with data extraction
যদি (result) {
    ফলাফল.সফল(মান) => লেখ("Success:", মান),
    ফলাফল.ত্রুটি(বার্তা) => লেখ("Error:", বার্তা)
}
```

---

## 3. Implementation Requirements

### Tokens (token/token.go)
```go
// Struct tokens
STRUCT = "স্ট্রাক্ট"  // struct keyword
DOT    = "."          // field access (already exists)

// Enum tokens
ENUM   = "enum"       // enum keyword
ARROW  = "=>"         // pattern matching (future)
```

### AST Nodes (ast/ast.go)

#### Struct AST Nodes
```go
// StructDefinition - defines a struct type
type StructDefinition struct {
    Token  token.Token  // the স্ট্রাক্ট token
    Name   *Identifier
    Fields []*StructField
}

type StructField struct {
    Name string
    Type *TypeAnnotation
}

// StructLiteral - creates a struct instance
type StructLiteral struct {
    Token      token.Token  // the { token
    StructType *Identifier  // struct type name
    Fields     map[string]Expression  // field values
}

// MemberAccessExpression - accesses struct field
type MemberAccessExpression struct {
    Token  token.Token  // the . token
    Object Expression   // the struct instance
    Member *Identifier  // the field name
}

// MemberAssignmentStatement - assigns to struct field
type MemberAssignmentStatement struct {
    Token  token.Token
    Object Expression
    Member *Identifier
    Value  Expression
}
```

#### Enum AST Nodes
```go
// EnumDefinition - defines an enum type
type EnumDefinition struct {
    Token    token.Token  // the enum token
    Name     *Identifier
    Variants []*EnumVariant
}

type EnumVariant struct {
    Name   string
    Value  Expression  // optional explicit value
    Fields []*StructField  // optional associated data
}

// EnumLiteral - creates an enum value
type EnumLiteral struct {
    Token    token.Token
    EnumType *Identifier
    Variant  *Identifier
    Fields   map[string]Expression  // if variant has associated data
}
```

### Object System (object/object.go)
```go
// Struct object
type Struct struct {
    TypeName string
    Fields   map[string]Object
    Methods  map[string]*Closure
}

// Enum object
type Enum struct {
    TypeName string
    Variant  string
    Value    Object  // associated data
}

// StructType - stores struct definition
type StructType struct {
    Name       string
    FieldTypes map[string]*ast.TypeAnnotation
}
```

### Bytecode (code/code.go)
```go
OpStruct       // Create struct instance
OpGetField     // Get struct field
OpSetField     // Set struct field
OpEnum         // Create enum value
OpMatchEnum    // Pattern match enum (future)
```

---

## 4. Example Use Cases

### For Self-Hosted Compiler

```bhasa
// Token definition
ধরি টোকেন = স্ট্রাক্ট {
    Type: লেখা,
    Literal: লেখা,
    Line: পূর্ণসংখ্যা,
    Column: পূর্ণসংখ্যা
};

// AST Node
ধরি LetStatement = স্ট্রাক্ট {
    Token: টোকেন,
    Name: Identifier,
    TypeAnnot: TypeAnnotation,
    Value: Expression
};

// Opcode enum
ধরি Opcode = enum {
    OpConstant,
    OpPop,
    OpAdd,
    OpSub,
    OpMul,
    OpDiv,
    OpCall
};
```

### Data Structures

```bhasa
// Linked list node
ধরি Node = স্ট্রাক্ট {
    value: পূর্ণসংখ্যা,
    next: Node  // nullable/option type needed
};

// Result type for error handling
ধরি Result = enum {
    Ok(value: পূর্ণসংখ্যা),
    Err(message: লেখা)
};
```

---

## 5. Implementation Phases

### Phase 1: Basic Structs (Week 1-2)
- [ ] Add STRUCT token
- [ ] Implement struct definition parsing
- [ ] Implement struct literal parsing
- [ ] Implement field access (dot notation)
- [ ] Compile struct definitions
- [ ] Create Struct object type
- [ ] VM support for struct operations

### Phase 2: Struct Methods (Week 3)
- [ ] Parse method definitions
- [ ] Compile method calls
- [ ] VM support for method dispatch

### Phase 3: Basic Enums (Week 4)
- [ ] Add ENUM token
- [ ] Parse simple enum definitions
- [ ] Parse enum value access
- [ ] Compile enum definitions
- [ ] Create Enum object type
- [ ] VM support for enum operations

### Phase 4: Advanced Enums (Week 5)
- [ ] Enums with associated data
- [ ] Pattern matching (basic)

---

## 6. Backward Compatibility

All existing code will continue to work. Structs and enums are purely additive features.

---

## 7. Future Extensions (Post-Struct/Enum)

1. **Interfaces** (after structs)
2. **Classes** (enhance structs with inheritance)
3. **Generics** (parameterized structs/enums)
4. **Pattern matching** (full implementation)
5. **Option/Maybe type** (built on enums)
6. **Result type** (built on enums for error handling)

---

## Decision Summary

✅ **Use `স্ট্রাক্ট` keyword for structs**
✅ **Use `enum` keyword for enums**
✅ **Go-style method syntax**: `ধরি (receiver: Type) method = ফাংশন() { ... }`
✅ **Dot notation for field access**: `object.field`
✅ **Rust-style enums with associated data**
✅ **Pattern matching syntax**: `যদি (value) { variant => ... }`

Let's implement these! 🚀
