# Object-Oriented Programming (OOP) in Bhasa

## Current OOP Features

Bhasa now supports basic Object-Oriented Programming features including:

### 1. Class Declarations (শ্রেণী)

You can define classes with methods using the `শ্রেণী` keyword:

```bengali
শ্রেণী গাড়ি {
    শুরু = ফাংশন() {
        লেখ("গাড়ি চলছে!");
    };
    
    থামো = ফাংশন() {
        লেখ("গাড়ি থেমেছে!");
    };
}
```

### 2. Object Instantiation (নতুন)

Create new instances of classes using the `নতুন` keyword:

```bengali
ধরি আমার_গাড়ি = নতুন গাড়ি();
```

### 3. Member Access (.)

Access properties and methods using dot notation:

```bengali
ধরি মান = obj.property;
```

### 4. This Keyword (এই)

Reference the current instance using the `এই` keyword:

```bengali
শ্রেণী ব্যক্তি {
    পরিচয় = ফাংশন() {
        লেখ(এই.নাম);
    };
}
```

## OOP Keywords

| Bengali | English | Usage |
|---------|---------|-------|
| শ্রেণী | class | Class declaration |
| নতুন | new | Object instantiation |
| এই | this | Current instance reference |

## Implementation Status

✅ **Implemented:**
- Class declarations with method definitions
- Object instantiation
- Member access syntax parsing
- Instance type in object system
- Basic OOP infrastructure (AST, compiler, VM opcodes)

🚧 **Partially Implemented:**
- Method calls on instances (infrastructure in place, needs refinement)
- Property access and modification
- Constructor methods
- `this` keyword context binding

📋 **Future Enhancements:**
- Full `this` context in method calls
- Constructor parameters
- Instance property initialization
- Inheritance (class extension)
- Method overriding
- Static methods and properties
- Private/public access modifiers
- Getters and setters

## Basic Example

```bengali
// Define a class
শ্রেণী গাড়ি {
    শুরু = ফাংশন() {
        লেখ("গাড়ি চলছে!");
    };
    
    থামো = ফাংশন() {
        লেখ("গাড়ি থেমেছে!");
    };
}

// Create an instance
ধরি আমার_গাড়ি = নতুন গাড়ি();
লেখ("গাড়ির ইনস্ট্যান্স তৈরি হয়েছে");
```

## Architecture

### Components Added

1. **Tokens:** CLASS, NEW, THIS, DOT
2. **AST Nodes:** ClassStatement, NewExpression, MemberAccessExpression, ThisExpression
3. **Object Types:** Class, Instance
4. **Opcodes:** OpClass, OpNewInstance, OpGetProperty, OpSetProperty, OpThis, OpCallMethod
5. **Compiler:** Class compilation support
6. **VM:** Instance creation and member access handlers

The OOP implementation follows the same bytecode compilation and VM execution model as the rest of the language, ensuring consistency and performance.
