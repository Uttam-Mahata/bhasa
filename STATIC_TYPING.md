# Static Typing Feature

## Overview

Bhasa now supports optional static typing! You can add type annotations to variables and functions to catch type errors during compilation. This feature is completely optional - existing code without type annotations will continue to work exactly as before.

## Type Keywords

The following Bengali type keywords are available:

| Type Keyword | English | Description |
|-------------|---------|-------------|
| `পূর্ণসংখ্যা` | Integer | Whole numbers |
| `লেখা` | String | Text strings |
| `বুলিয়ান` | Boolean | True/False values |
| `তালিকা` | Array | Arrays/Lists |
| `হ্যাশ` | Hash | Hash maps/Dictionaries |
| `ফাংশন_টাইপ` | Function | Function types |

## Variable Type Annotations

You can annotate variables with their expected types using the colon `:` syntax:

```bengali
// Basic integer type
ধরি x: পূর্ণসংখ্যা = ১০;

// String type
ধরি name: লেখা = "বাংলা";

// Boolean type
ধরি flag: বুলিয়ান = সত্য;

// Array type
ধরি numbers: তালিকা = [১, ২, ৩, ৪, ৫];

// Hash type
ধরি person: হ্যাশ = {"নাম": "রহিম", "বয়স": ২৫};
```

### Type Checking

When you add type annotations, the compiler will check that the assigned value matches the declared type:

```bengali
// ✓ Correct - this compiles without warnings
ধরি x: পূর্ণসংখ্যা = ৫;

// ✗ Type mismatch - compiler warning
ধরি y: পূর্ণসংখ্যা = "not a number";
// Warning: type mismatch: variable 'y' declared as পূর্ণসংখ্যা but assigned লেখা
```

## Function Type Annotations

### Parameter Type Annotations

You can annotate function parameters with their expected types:

```bengali
ধরি add = ফাংশন(a: পূর্ণসংখ্যা, b: পূর্ণসংখ্যা) {
    ফেরত a + b;
};
```

### Return Type Annotations

You can also specify the return type of a function:

```bengali
ধরি multiply = ফাংশন(a: পূর্ণসংখ্যা, b: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত a * b;
};
```

### Full Function Type Annotations

Combining parameter and return type annotations:

```bengali
ধরি isPositive = ফাংশন(n: পূর্ণসংখ্যা): বুলিয়ান {
    ফেরত n > ০;
};

ধরি greet = ফাংশন(name: লেখা): লেখা {
    ফেরত "নমস্কার, " + name;
};
```

### Function Variable Types

You can also declare the type of a function variable:

```bengali
ধরি calculator: ফাংশন_টাইপ = ফাংশন(x: পূর্ণসংখ্যা, y: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত x + y;
};
```

## Dynamic Typing (Backward Compatible)

The type system is completely optional. You can mix typed and untyped code freely:

```bengali
// Without type annotations (dynamic typing)
ধরি x = ৫;
ধরি name = "বাংলা";

// With type annotations (static typing)
ধরি y: পূর্ণসংখ্যা = ১০;
ধরি title: লেখা = "শিরোনাম";

// Mixed in same program
ধরি add = ফাংশন(a, b) {
    ফেরত a + b;
};

ধরি multiply = ফাংশন(a: পূর্ণসংখ্যা, b: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত a * b;
};
```

## Type Checking Behavior

### Compile-Time Warnings

Type mismatches generate **warnings** during compilation, not errors. Your program will still run:

```
Type checking warnings:
	type mismatch: variable 'num' declared as পূর্ণসংখ্যা but assigned লেখা
	type mismatch: variable 'text' declared as লেখা but assigned পূর্ণসংখ্যা
```

This design choice maintains backward compatibility and allows for gradual adoption of types.

### Type Inference

The type checker can infer types from expressions:

- Literal values: `৫` → `পূর্ণসংখ্যা`, `"text"` → `লেখা`
- Arithmetic operations: `x + y` → `পূর্ণসংখ্যা`
- Comparison operations: `x > y` → `বুলিয়ান`
- Boolean operations: `!flag` → `বুলিয়ান`

## Examples

### Example 1: Basic Type Annotations

```bengali
// Variable type annotations
ধরি age: পূর্ণসংখ্যা = ২৫;
ধরি name: লেখা = "রহিম";
ধরি isStudent: বুলিয়ান = সত্য;

লেখ("নাম: " + name);
লেখ("বয়স: ");
লেখ(age);
লেখ("ছাত্র: ");
লেখ(isStudent);
```

### Example 2: Typed Functions

```bengali
// Function with full type annotations
ধরি calculateArea = ফাংশন(length: পূর্ণসংখ্যা, width: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত length * width;
};

ধরি area = calculateArea(১০, ২০);
লেখ("এলাকা: ");
লেখ(area);
```

### Example 3: Higher-Order Functions

```bengali
// Higher-order function with types
ধরি applyTwice = ফাংশন(f: ফাংশন_টাইপ, x: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত f(f(x));
};

ধরি increment = ফাংশন(n: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত n + ১;
};

ধরি result = applyTwice(increment, ৫);
লেখ("ফলাফল: ");
লেখ(result); // Outputs: 7
```

### Example 4: Mixed Typing

```bengali
// Mixing typed and untyped code
ধরি typedAdd = ফাংশন(a: পূর্ণসংখ্যা, b: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত a + b;
};

ধরি untypedAdd = ফাংশন(a, b) {
    ফেরত a + b;
};

// Both work fine
লেখ(typedAdd(১০, ২০));    // 30
লেখ(untypedAdd(৫, ১৫));   // 20
```

## Benefits of Static Typing

1. **Early Error Detection**: Catch type-related bugs during compilation
2. **Better Documentation**: Type annotations serve as inline documentation
3. **IDE Support**: Enables better autocomplete and tooling (future)
4. **Refactoring Safety**: Types help ensure correctness when changing code
5. **Optional**: Use types where they help, skip them where they don't

## Migration Guide

### For Existing Code

No changes needed! All existing code continues to work:

```bengali
// This still works exactly as before
ধরি x = ১০;
ধরি f = ফাংশন(a, b) { ফেরত a + b; };
```

### Gradually Adding Types

You can add types incrementally:

```bengali
// Step 1: Start with dynamic code
ধরি calculate = ফাংশন(x, y) {
    ফেরত x * y + ১০;
};

// Step 2: Add parameter types
ধরি calculate = ফাংশন(x: পূর্ণসংখ্যা, y: পূর্ণসংখ্যা) {
    ফেরত x * y + ১০;
};

// Step 3: Add return type
ধরি calculate = ফাংশন(x: পূর্ণসংখ্যা, y: পূর্ণসংখ্যা): পূর্ণসংখ্যা {
    ফেরত x * y + ১০;
};
```

## Future Enhancements

Potential future improvements to the type system:

1. **Strict Mode**: Optional flag to treat type warnings as errors
2. **Generic Types**: Support for `তালিকা<পূর্ণসংখ্যা>` (Array of integers)
3. **Union Types**: Support for values that can be multiple types
4. **Type Aliases**: Custom type names
5. **Interface Types**: Structural typing for objects
6. **Type Inference**: Automatically infer types without annotations

## Technical Details

### Implementation

- **Lexer**: Recognizes type keywords as tokens
- **Parser**: Parses type annotations in syntax tree
- **AST**: Stores type information in nodes
- **Type Checker**: Validates type consistency
- **Compiler**: Integrates type checking during compilation
- **Symbol Table**: Tracks types for variables

### Performance

Type checking adds minimal overhead:
- O(n) complexity where n is the number of statements
- Performed only once during compilation
- No runtime overhead - types are not checked during execution

## Conclusion

Static typing in Bhasa provides the best of both worlds:
- **Flexibility**: Optional types mean you choose when to use them
- **Safety**: Type checking catches errors early
- **Compatibility**: Existing code continues to work
- **Gradual Adoption**: Add types incrementally

Start using types today to make your Bhasa code more robust and maintainable!
