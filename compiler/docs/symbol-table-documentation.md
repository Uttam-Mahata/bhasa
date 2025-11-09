# Symbol Table Documentation

## Table of Contents

1. [Overview](#overview)
2. [Core Types](#core-types)
3. [Scope Types](#scope-types)
4. [API Reference](#api-reference)
5. [Scope Resolution](#scope-resolution)
6. [Closure Implementation](#closure-implementation)
7. [Type Annotations](#type-annotations)
8. [Examples](#examples)
9. [Best Practices](#best-practices)

---

## Overview

The Symbol Table is a critical component of the Bhasa compiler that tracks variable names, their scopes, and storage locations. It implements lexical scoping with support for closures through **free variables**.

### Key Responsibilities

- **Variable Tracking**: Map names to storage locations
- **Scope Management**: Handle nested scopes (global, local, etc.)
- **Closure Support**: Track free variables from outer scopes
- **Type Tracking**: Store optional type annotations
- **Builtin Registration**: Track built-in functions

---

## Core Types

### SymbolScope

```go
type SymbolScope string

const (
    GlobalScope   SymbolScope = "GLOBAL"    // Module-level variables
    LocalScope    SymbolScope = "LOCAL"     // Function parameters and locals
    BuiltinScope  SymbolScope = "BUILTIN"   // Built-in functions
    FreeScope     SymbolScope = "FREE"      // Captured variables (closures)
    FunctionScope SymbolScope = "FUNCTION"  // Function name (recursion)
)
```

**Purpose**: Identifies where a variable is stored

**Usage**:
- Compiler uses scope to emit correct OpCode (OpGetGlobal vs OpGetLocal)
- VM uses scope to know where to look for variable

---

### Symbol

```go
type Symbol struct {
    Name       string              // Variable name
    Scope      SymbolScope         // Scope type
    Index      int                 // Storage index
    TypeAnnot  *ast.TypeAnnotation // Optional type
}
```

**Fields**:

- **Name**: The identifier (e.g., "x", "myFunc", "নাম")
- **Scope**: Where the variable is stored
- **Index**: Position in storage array
  - Global scope: index in global array
  - Local scope: index in stack frame
  - Builtin scope: index in builtins array
  - Free scope: index in free variables array
- **TypeAnnot**: Optional type annotation (e.g., `পূর্ণসংখ্যা`, `লেখা`)

**Example**:
```go
Symbol{
    Name:  "x",
    Scope: LocalScope,
    Index: 0,
    TypeAnnot: &ast.TypeAnnotation{TypeName: "পূর্ণসংখ্যা"},
}
```

---

### SymbolTable

```go
type SymbolTable struct {
    Outer          *SymbolTable      // Parent scope
    store          map[string]Symbol // Symbol storage
    numDefinitions int               // Number of symbols defined
    FreeSymbols    []Symbol          // Captured free variables
}
```

**Fields**:

- **Outer**: Pointer to parent scope (nil for global scope)
  - Forms a **linked list** of scopes
  - Used for variable resolution

- **store**: Map of name → Symbol
  - Fast O(1) lookup
  - Stores all symbols in this scope

- **numDefinitions**: Counter for local variable indices
  - Incremented on each Define()
  - Used as index for next defined variable

- **FreeSymbols**: List of captured variables
  - Only populated in nested scopes
  - Used for closure creation

---

## Scope Types

### 1. Global Scope

**Definition**: Top-level module scope

**Characteristics**:
- `Outer == nil`
- Variables defined at program level
- Persistent across function calls

**Example**:
```bhasa
ধরি x = 5;        // Global
ধরি y = 10;       // Global

ফাংশন() {
    x + y         // Access globals
}
```

**Symbol Table**:
```
┌─────────────────┐
│  Global Scope   │
│  Outer: nil     │
├─────────────────┤
│  x → Index: 0   │
│  y → Index: 1   │
└─────────────────┘
```

---

### 2. Local Scope

**Definition**: Function parameters and local variables

**Characteristics**:
- `Outer != nil` (points to enclosing scope)
- Variables local to function
- Destroyed after function returns

**Example**:
```bhasa
ফাংশন(a, b) {     // a, b are local
    ধরি c = a + b; // c is local
    ফেরত c;
}
```

**Symbol Table**:
```
┌─────────────────┐
│  Global Scope   │
└────────┬────────┘
         │
         ↓
┌─────────────────┐
│  Local Scope    │
│  Outer: Global  │
├─────────────────┤
│  a → Index: 0   │  (parameter)
│  b → Index: 1   │  (parameter)
│  c → Index: 2   │  (local)
└─────────────────┘
```

---

### 3. Builtin Scope

**Definition**: Built-in functions

**Characteristics**:
- Predefined by compiler
- Cannot be redefined
- Always accessible

**Example**:
```bhasa
len([1, 2, 3])     // len is builtin
দেখাও("hello")     // দেখাও is builtin
```

**Built-in Functions**:
- `len` - Get length
- `প্রথম` - First element
- `শেষ` - Last element
- `বাকি` - Rest (all but first)
- `যোগ` - Push to array
- `দেখাও` - Print

---

### 4. Free Scope

**Definition**: Variables captured from outer scopes (closures)

**Characteristics**:
- Created automatically during resolution
- Stored in closure's free variables array
- Allows nested functions to access outer variables

**Example**:
```bhasa
ফাংশন(x) {        // x is local in outer
    ফাংশন(y) {    // y is local in inner
        x + y      // x is FREE in inner (captured from outer)
    }
}
```

**Symbol Tables**:
```
Outer Function:
┌─────────────────┐
│  Local Scope    │
├─────────────────┤
│  x → Index: 0   │
└─────────────────┘

Inner Function:
┌─────────────────┐
│  Local Scope    │
├─────────────────┤
│  y → Index: 0   │  (parameter)
│  x → FREE: 0    │  (captured from outer)
└─────────────────┘
FreeSymbols: [x]
```

---

### 5. Function Scope

**Definition**: Function name for recursion

**Characteristics**:
- Allows function to call itself
- Special scope type
- Index is always 0

**Example**:
```bhasa
ধরি factorial = ফাংশন(n) {
    যদি (n <= 1) {
        ফেরত 1;
    }
    ফেরত n * factorial(n - 1);  // Recursive call
};
```

**Symbol Table**:
```
┌─────────────────────┐
│  Function Scope     │
├─────────────────────┤
│  n → Local: 0       │
│  factorial → FUNC:0 │  (self-reference)
└─────────────────────┘
```

---

## API Reference

### NewSymbolTable

```go
func NewSymbolTable() *SymbolTable
```

**Creates a new symbol table**:
- Empty store
- No outer scope (global)
- Empty free symbols list

**Usage**:
```go
global := NewSymbolTable()
```

---

### NewEnclosedSymbolTable

```go
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable
```

**Creates a nested symbol table**:
- New empty store
- Points to outer scope
- Used for function scopes

**Usage**:
```go
global := NewSymbolTable()
local := NewEnclosedSymbolTable(global)
```

**Scope Chain**:
```
global ← local ← nested
```

---

### Define

```go
func (s *SymbolTable) Define(name string) Symbol
```

**Defines a new symbol**:
- Adds to store
- Assigns index (numDefinitions)
- Determines scope (global if no outer, local otherwise)
- Increments numDefinitions

**Returns**: The defined symbol

**Example**:
```go
symbol := table.Define("x")
// symbol.Name: "x"
// symbol.Scope: GlobalScope (if table.Outer == nil)
// symbol.Index: 0 (first definition)
```

---

### DefineWithType

```go
func (s *SymbolTable) DefineWithType(name string, typeAnnot *ast.TypeAnnotation) Symbol
```

**Same as Define** but includes type annotation

**Example**:
```go
typeAnnot := &ast.TypeAnnotation{TypeName: "পূর্ণসংখ্যা"}
symbol := table.DefineWithType("x", typeAnnot)
// symbol.TypeAnnot: পূর্ণসংখ্যা
```

---

### DefineBuiltin

```go
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol
```

**Defines a builtin function**:
- Uses provided index (from builtins array)
- Sets scope to BUILTIN
- Does NOT increment numDefinitions

**Usage**:
```go
for i, builtin := range object.Builtins {
    table.DefineBuiltin(i, builtin.Name)
}
```

---

### DefineFunctionName

```go
func (s *SymbolTable) DefineFunctionName(name string) Symbol
```

**Defines function name for recursion**:
- Uses FUNCTION scope
- Index is always 0
- Allows function to reference itself

**Usage** (typically done by compiler):
```go
symbol := table.DefineFunctionName("factorial")
```

---

### Resolve

```go
func (s *SymbolTable) Resolve(name string) (Symbol, bool)
```

**Resolves a symbol by name**:
- Looks in current scope's store
- If not found, looks in outer scope
- If found in outer local scope, creates free variable
- Returns (symbol, true) if found, (empty, false) otherwise

**Resolution Algorithm**:

1. **Check current scope**:
   - If found → return symbol

2. **Check outer scope** (if exists):
   - Recursively resolve in outer
   - If not found → return (empty, false)
   - If scope is Global/Builtin → return as-is
   - If scope is Local → **define as free variable**

**Example**:
```go
symbol, ok := table.Resolve("x")
if !ok {
    return fmt.Errorf("undefined variable: x")
}
```

---

### defineFree (internal)

```go
func (s *SymbolTable) defineFree(original Symbol) Symbol
```

**Defines a free variable**:
- Adds original symbol to FreeSymbols list
- Creates new symbol with FREE scope
- Index is position in FreeSymbols array
- Stores in current scope

**Usage**: Called automatically by Resolve()

---

## Scope Resolution

### Resolution Examples

#### Example 1: Simple Global

```bhasa
ধরি x = 5;
x + 10;
```

**Symbol Tables**:
```
Global:
  Define("x") → {Name: "x", Scope: GLOBAL, Index: 0}
  Resolve("x") → {Name: "x", Scope: GLOBAL, Index: 0} ✓
```

---

#### Example 2: Local Variable

```bhasa
ফাংশন(x) {
    ধরি y = x + 5;
    y;
}
```

**Symbol Tables**:
```
Function Scope (Outer: Global):
  Define("x") → {Name: "x", Scope: LOCAL, Index: 0}   (parameter)
  Define("y") → {Name: "y", Scope: LOCAL, Index: 1}   (local)
  Resolve("x") → {Name: "x", Scope: LOCAL, Index: 0} ✓
  Resolve("y") → {Name: "y", Scope: LOCAL, Index: 1} ✓
```

---

#### Example 3: Free Variable (Closure)

```bhasa
ধরি makeAdder = ফাংশন(x) {
    ফাংশন(y) {
        x + y
    }
};
```

**Symbol Tables**:

**Outer Function Scope**:
```
Outer: Global
  Define("x") → {Name: "x", Scope: LOCAL, Index: 0}
```

**Inner Function Scope**:
```
Outer: Outer Function Scope
  Define("y") → {Name: "y", Scope: LOCAL, Index: 0}
  Resolve("x"):
    1. Not in current store
    2. Check outer → Found: {Name: "x", Scope: LOCAL, Index: 0}
    3. It's LOCAL in outer, so define as FREE
    4. Return: {Name: "x", Scope: FREE, Index: 0}
  
  FreeSymbols: [{Name: "x", Scope: LOCAL, Index: 0}]
```

**Result**:
- Inner function captures `x` as free variable
- `x` stored in closure's free variables array
- Compiler emits `OpGetFree 0` to access it

---

#### Example 4: Multiple Nesting Levels

```bhasa
ধরি a = 1;
ফাংশন() {
    ধরি b = 2;
    ফাংশন() {
        ধরি c = 3;
        a + b + c;
    }
}
```

**Symbol Tables**:

**Global**:
```
Define("a") → {Name: "a", Scope: GLOBAL, Index: 0}
```

**Level 1 Function**:
```
Outer: Global
Define("b") → {Name: "b", Scope: LOCAL, Index: 0}
```

**Level 2 Function**:
```
Outer: Level 1 Function
Define("c") → {Name: "c", Scope: LOCAL, Index: 0}

Resolve("a"):
  → Not in current
  → Not in Level 1
  → Found in Global: GLOBAL, return as-is
  
Resolve("b"):
  → Not in current
  → Found in Level 1: LOCAL, define as FREE
  → {Name: "b", Scope: FREE, Index: 0}
  
Resolve("c"):
  → Found in current: LOCAL
  → {Name: "c", Scope: LOCAL, Index: 0}

FreeSymbols: [{Name: "b", Scope: LOCAL, Index: 0}]
```

---

## Closure Implementation

### How Closures Work

**1. Compiler Phase**:
- Inner function resolved references to outer variables
- Symbol table tracks these as **free variables**
- Compiler emits instructions to load free variables

**2. Runtime Phase**:
- VM creates closure with free variables
- Free variables captured at closure creation time
- Closure carries its environment

### Example: Counter Closure

**Bhasa Code**:
```bhasa
ধরি makeCounter = ফাংশন() {
    ধরি count = 0;
    ফাংশন() {
        count = count + 1;
        count
    }
};

ধরি counter = makeCounter();
counter();  // 1
counter();  // 2
counter();  // 3
```

**Compilation**:

**Outer Function**:
```
OpConstant 0         // 0
OpSetLocal 0         // count = 0
OpGetLocal 0         // Load count for closure
OpClosure 1 1        // Create inner function with 1 free var
OpReturnValue
```

**Inner Function** (constant[1]):
```
OpGetFree 0          // Load count (free)
OpConstant 0         // 1
OpAdd                // count + 1
OpSetFree 0          // Store back to count
OpGetFree 0          // Load count for return
OpReturnValue
```

**Key Point**: Each call to `makeCounter()` creates a **new** closure with its own `count`

---

## Type Annotations

### Type Tracking

Symbol table optionally tracks type annotations:

```bhasa
ধরি x: পূর্ণসংখ্যা = 5;
```

**Symbol**:
```go
{
    Name: "x",
    Scope: GlobalScope,
    Index: 0,
    TypeAnnot: &ast.TypeAnnotation{TypeName: "পূর্ণসংখ্যা"}
}
```

### Type Checking

**Compiler Phase**:
- Type annotations stored in symbol table
- Compiler emits `OpAssertType` for runtime checks

**Runtime Phase**:
- VM checks value type against annotation
- Throws error on type mismatch

---

## Examples

### Example 1: Simple Program

**Code**:
```bhasa
ধরি x = 5;
ধরি y = 10;
x + y;
```

**Symbol Table Evolution**:

```
Initial:
┌─────────────┐
│  Global     │
│  store: {}  │
└─────────────┘

After "ধরি x = 5;":
┌─────────────────────┐
│  Global             │
│  store: {           │
│    x → GLOBAL:0     │
│  }                  │
│  numDefinitions: 1  │
└─────────────────────┘

After "ধরি y = 10;":
┌─────────────────────┐
│  Global             │
│  store: {           │
│    x → GLOBAL:0     │
│    y → GLOBAL:1     │
│  }                  │
│  numDefinitions: 2  │
└─────────────────────┘
```

---

### Example 2: Function with Parameters

**Code**:
```bhasa
ফাংশন(a, b) {
    ধরি sum = a + b;
    sum;
}
```

**Symbol Table**:

```
Function Scope:
┌─────────────────────┐
│  Outer: Global      │
│  store: {           │
│    a → LOCAL:0      │  (parameter)
│    b → LOCAL:1      │  (parameter)
│    sum → LOCAL:2    │  (local variable)
│  }                  │
│  numDefinitions: 3  │
└─────────────────────┘
```

---

### Example 3: Closure

**Code**:
```bhasa
ফাংশন(x) {
    ফাংশন(y) {
        x + y
    }
}
```

**Symbol Tables**:

```
Outer Function:
┌─────────────────────┐
│  Outer: Global      │
│  store: {           │
│    x → LOCAL:0      │
│  }                  │
│  numDefinitions: 1  │
│  FreeSymbols: []    │
└─────────────────────┘

Inner Function:
┌─────────────────────┐
│  Outer: Outer Func  │
│  store: {           │
│    y → LOCAL:0      │
│    x → FREE:0       │
│  }                  │
│  numDefinitions: 1  │
│  FreeSymbols: [     │
│    {x, LOCAL, 0}    │
│  ]                  │
└─────────────────────┘
```

---

### Example 4: Recursive Function

**Code**:
```bhasa
ধরি factorial = ফাংশন(n) {
    যদি (n <= 1) {
        ফেরত 1;
    }
    ফেরত n * factorial(n - 1);
};
```

**Symbol Table** (inside function):

```
Function Scope:
┌───────────────────────┐
│  Outer: Global        │
│  store: {             │
│    n → LOCAL:0        │
│    factorial → FUNC:0 │
│  }                    │
│  numDefinitions: 1    │
└───────────────────────┘
```

**Note**: `factorial` defined with `DefineFunctionName()` for recursion

---

## Best Practices

### 1. Always Check Resolution Result

```go
// ✅ Good
symbol, ok := table.Resolve("x")
if !ok {
    return fmt.Errorf("undefined variable: x")
}

// ❌ Bad
symbol, _ := table.Resolve("x")
// Using symbol without checking ok
```

---

### 2. Use Proper Scope Management

```go
// ✅ Good - Enter and leave scopes properly
compiler.enterScope()
// ... compile function body
instructions := compiler.leaveScope()

// ❌ Bad - Manual scope manipulation
compiler.scopeIndex++
// ... 
compiler.scopeIndex--
```

---

### 3. Define Before Resolve

```go
// ✅ Good
table.Define("x")
// ... later
symbol, _ := table.Resolve("x")  // Found

// ❌ Bad
symbol, ok := table.Resolve("x")  // Not found!
if ok {
    // This block won't execute
}
```

---

### 4. Handle Free Variables Correctly

```go
// ✅ Good - Compiler handles this automatically
freeSymbols := symbolTable.FreeSymbols
for _, s := range freeSymbols {
    compiler.loadSymbol(s)
}

// ❌ Bad - Don't manually manipulate FreeSymbols
symbolTable.FreeSymbols = append(...)  // Don't do this
```

---

## Common Patterns

### Pattern 1: Define and Store

```go
// Define symbol
symbol := c.symbolTable.Define(name)

// Compile value
c.Compile(value)

// Store based on scope
if symbol.Scope == GlobalScope {
    c.emit(code.OpSetGlobal, symbol.Index)
} else {
    c.emit(code.OpSetLocal, symbol.Index)
}
```

---

### Pattern 2: Resolve and Load

```go
// Resolve symbol
symbol, ok := c.symbolTable.Resolve(name)
if !ok {
    return fmt.Errorf("undefined: %s", name)
}

// Load based on scope
c.loadSymbol(symbol)
```

---

### Pattern 3: Function Scope

```go
// Enter new scope
c.enterScope()

// Define parameters
for _, param := range parameters {
    c.symbolTable.Define(param.Value)
}

// Compile body
c.Compile(body)

// Collect free variables
freeSymbols := c.symbolTable.FreeSymbols

// Leave scope
instructions := c.leaveScope()
```

---

## Troubleshooting

### Issue: Undefined Variable Error

**Symptom**: Compiler says variable is undefined

**Causes**:
1. Variable not defined before use
2. Variable defined in wrong scope
3. Typo in variable name

**Solution**:
```go
// Check if variable is defined
symbol, ok := table.Resolve("varName")
if !ok {
    // Variable not in any scope
}

// Check which scope it's in
fmt.Println(symbol.Scope)
```

---

### Issue: Closure Not Working

**Symptom**: Closure doesn't capture variables correctly

**Causes**:
1. Free variables not loaded before closure creation
2. Symbol table not properly nested

**Solution**:
```go
// Ensure free variables are loaded
freeSymbols := c.symbolTable.FreeSymbols
for _, s := range freeSymbols {
    c.loadSymbol(s)  // Load before OpClosure
}
c.emit(code.OpClosure, fnIndex, len(freeSymbols))
```

---

### Issue: Wrong Variable Value

**Symptom**: Variable has unexpected value

**Causes**:
1. Wrong index used
2. Wrong scope used
3. Variable shadowing

**Solution**:
```go
// Print symbol info for debugging
fmt.Printf("Symbol: %+v\n", symbol)

// Check scope and index
fmt.Printf("Scope: %s, Index: %d\n", symbol.Scope, symbol.Index)
```

---

## Performance Considerations

### O(1) Lookup

Symbol table uses a map for O(1) average lookup time:

```go
store map[string]Symbol  // Fast lookup
```

---

### Nested Scope Chain

Worst case: O(n) where n is nesting depth

```
Global → Function1 → Function2 → Function3
```

For most programs, nesting depth < 10, so very fast

---

### Memory Usage

Each symbol table stores:
- Map of symbols: O(m) where m = number of symbols
- Free symbols: O(f) where f = captured variables

Total: O(m + f) per scope

---

## See Also

- [Compiler Documentation](./compiler-documentation.md) - How symbol table is used
- [Compilation Examples](./compilation-examples.md) - Examples with symbol tables
- [Quick Reference](./quick-reference.md) - Quick lookup
- [VM Documentation](../../vm/docs/) - How symbols are used at runtime

