# Bhasa Architecture - Quick Reference

## 1. COMPILATION PIPELINE AT A GLANCE

```
Bengali Source Code
       ↓
  [LEXER]        → Tokenize (token/token.go)
       ↓
  [PARSER]       → Build AST (parser/parser.go)
       ↓
  [COMPILER]     → Generate Bytecode (compiler/compiler.go)
       ↓
  [VM]           → Execute (vm/vm.go)
       ↓
   OUTPUT
```

---

## 2. KEY FILES BY RESPONSIBILITY

### Phase 1: Tokenization
- **token/token.go** - Token types (41 keywords, 20+ operators)
- **lexer/lexer.go** - Tokenizer with Bengali support

### Phase 2: Parsing
- **parser/parser.go** - Pratt parser (1,169 lines)
- **ast/ast.go** - AST node definitions (654 lines)

### Phase 3: Compilation
- **compiler/compiler.go** - AST to bytecode (868 lines)
- **compiler/symbol_table.go** - Scoping & variable tracking
- **code/code.go** - Instruction definitions

### Phase 4: Execution
- **vm/vm.go** - Stack-based VM (1,172 lines)
- **vm/frame.go** - Function call frames
- **object/object.go** - Runtime objects (370 lines)

### Utilities
- **main.go** - Entry point
- **repl/repl.go** - Interactive shell
- **evaluator/** - Legacy tree-walking interpreter

---

## 3. COMPILATION CONCEPTS

### Symbol Table (Variable Scoping)

**5 Scope Types:**
```
GlobalScope   → Module-level variables
LocalScope    → Function-local variables
BuiltinScope  → Built-in functions
FreeScope     → Variables captured in closures
FunctionScope → For recursion support
```

**Lookup Chain:**
```
identifier x
  ↓
Check local scope
  ↓ (not found)
Check enclosing scopes
  ↓ (not found)
Check global scope
  ↓ (not found)
Check built-in functions
  ↓ (not found)
ERROR: undefined variable
```

### Bytecode Instruction Set (41 opcodes)

**Major Categories:**
- **Data**: OpConstant, OpNull, OpTrue, OpFalse
- **Arithmetic**: OpAdd, OpSub, OpMul, OpDiv, OpMod
- **Bitwise**: OpBitAnd, OpBitOr, OpBitXor, OpBitNot, OpLeftShift, OpRightShift
- **Comparison**: OpEqual, OpNotEqual, OpGreaterThan, OpGreaterThanEqual
- **Logical**: OpAnd, OpOr, OpBang
- **Variables**: OpGetGlobal, OpSetGlobal, OpGetLocal, OpSetLocal, OpGetBuiltin, OpGetFree
- **Control Flow**: OpJump, OpJumpNotTruthy, OpReturn, OpReturnValue
- **Functions**: OpClosure, OpCall, OpCurrentClosure
- **Collections**: OpArray, OpHash, OpIndex
- **OOP**: OpStruct, OpGetStructField, OpSetStructField, OpEnum

### Virtual Machine Stack

```
VM Memory Layout:

┌─────────────────────────┐
│   STACK (2048)          │ ← sp points here
├─────────────────────────┤
│   (free space)          │
├─────────────────────────┤
│   GLOBALS (65536)       │ ← Global variables
├─────────────────────────┤
│   CONSTANTS             │ ← From bytecode
└─────────────────────────┘
```

**Execution for `2 + 3`:**
```
OpConstant 0  → push 2 onto stack
OpConstant 1  → push 3 onto stack
OpAdd         → pop 3, pop 2, push 5
```

---

## 4. PARSER OPERATOR PRECEDENCE

```
Priority 1: INDEX (arr[i], obj.field)
Priority 2: CALL (fn())
Priority 3: PREFIX (!, -, ~)
Priority 4: PRODUCT (*, /, %)
Priority 5: SUM (+, -)
Priority 6: SHIFT (<<, >>)
Priority 7: BIT_AND (&)
Priority 8: BIT_XOR (^)
Priority 9: BIT_OR (|)
Priority 10: COMPARISON (<, >, <=, >=)
Priority 11: EQUALITY (==, !=)
Priority 12: LOGICAL_AND (&&)
Priority 13: LOGICAL_OR (||)
Priority 14: LOWEST
```

---

## 5. STATEMENT VS EXPRESSION

### Statements (Execute, produce side effects)
- LetStatement (ধরি)
- ReturnStatement (ফেরত)
- AssignmentStatement
- BlockStatement
- WhileStatement (যতক্ষণ)
- ForStatement (পর্যন্ত)
- BreakStatement (বিরতি)
- ContinueStatement (চালিয়ে_যাও)
- MemberAssignmentStatement
- ImportStatement (অন্তর্ভুক্ত)

### Expressions (Evaluate, produce values)
- Literals: Integer, String, Boolean, Array, Hash
- Identifiers
- PrefixExpression (!x, -x, ~x)
- InfixExpression (x + y)
- IfExpression (যদি/নাহলে)
- FunctionLiteral (ফাংশন)
- CallExpression (fn())
- IndexExpression (arr[i])
- TypeCastExpression (as)
- StructDefinition (স্ট্রাক্ট)
- StructLiteral
- MemberAccessExpression (obj.field)
- EnumDefinition (গণনা)
- EnumValue

---

## 6. RUNTIME OBJECT TYPES

```
Primitives: Integer, Byte, Short, Int, Long, Float, Double, Char, Boolean, String

Collections: Array, Hash

Functions: Function, Builtin, CompiledFunction, Closure

Control: ReturnValue, Error

OOP: Struct, Enum, EnumType

Special: Null
```

---

## 7. CLOSURE IMPLEMENTATION

**Example Code:**
```bengali
ধরি makeCounter = ফাংশন() {
    ধরি count = 0;
    ফেরত ফাংশন() {
        count = count + 1;
        ফেরত count;
    };
};
```

**Compilation Process:**
1. Outer function compiled, count is local variable
2. Inner function references count (from parent scope)
3. count marked as "FREE" symbol
4. Inner function compiled with OpGetFree instruction
5. At runtime, closure captures count in Free array

**Execution:**
```
makeCounter()     → Creates CompiledFunction
   ↓
Call OpClosure    → Wraps function with Free vars
   ↓
Push to stack     → Closure object
   ↓
Call returned fn  → Uses captured count
```

---

## 8. STRUCT/ENUM CURRENT STATE

### What Works ✓

**Structs:**
- Definition: `ধরি Person = স্ট্রাক্ট { name: লেখা, age: পূর্ণসংখ্যা }`
- Creation: `Person{name: "রহিম", age: 30}`
- Access: `person.name`
- Assignment: `person.age = 31`

**Enums:**
- Definition: `ধরি Direction = গণনা { North, South, East, West }`
- Value: `Direction.North`
- Comparison: `if (dir == Direction.North) { ... }`

### What's Missing ✗

- Methods on structs
- Inheritance
- Interfaces
- Enum associated data
- Pattern matching
- Visibility modifiers
- Type constraints

---

## 9. KEY DESIGN PATTERNS

### Pratt Parsing

**Strategy:** Parse by operator precedence, not recursion
- Register prefix handlers (!, -, literals)
- Register infix handlers (+, -, *, /, etc.)
- Use precedence levels to control parsing

**Example:**
```
Input: 2 + 3 * 4

Parse 2              → IntegerLiteral(2)
See +, check precedence
  Parse 3            → IntegerLiteral(3)
  See *, precedence > +
    Parse 4          → IntegerLiteral(4)
  Return 3 * 4
Return 2 + (3 * 4)
```

### Jump Patching

**For if-else statements:**
```
Compile condition
Emit OpJumpNotTruthy with placeholder position
Compile consequence block
Patch OpJumpNotTruthy to point to alternative
Compile alternative block
```

---

## 10. BUILTIN FUNCTIONS REFERENCE

### I/O
- লেখ(value) - Print

### String (7)
- বিভক্ত, যুক্ত, উপরে, নিচে, ছাঁটো, প্রতিস্থাপন, খুঁজুন

### Array (6)
- দৈর্ঘ্য, প্রথম, শেষ, বাকি, যোগ, উল্টাও

### Math (6)
- শক্তি, বর্গমূল, পরম, সর্বোচ্চ, সর্বনিম্ন, গোলাকার

### Type (10+)
- টাইপ, বাইট, ছোট_সংখ্যা, পূর্ণসংখ্যা, দীর্ঘ_সংখ্যা, দশমিক, দশমিক_দ্বিগুণ

### File I/O (4)
- ফাইল_পড়ো, ফাইল_লেখো, ফাইল_যোগ, ফাইল_আছে

### Character (3)
- অক্ষর, কোড, অক্ষর_থেকে_কোড, সংখ্যা, লেখা

---

## 11. WHERE TO MAKE CHANGES

### To Add Method Support
1. **token/token.go** - Add METHOD keyword
2. **parser/parser.go** - Parse method syntax `(receiver: Type) name = ফাংশন() { ... }`
3. **ast/ast.go** - Add MethodDefinition node
4. **compiler/compiler.go** - Compile methods as closures
5. **object/object.go** - Add Methods map to Struct
6. **vm/vm.go** - Add method dispatch logic

### To Enhance Enums
1. **parser/parser.go** - Parse associated data `Variant(field: Type)`
2. **ast/ast.go** - Extend EnumVariant with fields
3. **compiler/compiler.go** - Compile variant construction
4. **object/object.go** - Store data in Enum
5. **vm/vm.go** - Support pattern matching opcodes

### To Add Interfaces
1. **token/token.go** - Add INTERFACE keyword
2. **parser/parser.go** - Parse interface definitions
3. **ast/ast.go** - Add InterfaceDefinition node
4. **compiler/compiler.go** - Compile interface checks
5. **object/object.go** - Add Interface type
6. **vm/vm.go** - Implement interface checking

---

## 12. IMPORTANT INVARIANTS

1. **Instruction Pointer** always points to next instruction to execute
2. **Stack Pointer** points to next free slot (not to top element)
3. **Constants Pool** indices are uint16 (65536 max constants)
4. **Global Storage** 65536 slots available (needs DefineGlobal)
5. **Closures** must capture all free variables at compile time
6. **Type Checking** is optional (dynamic typing)

---

## 13. TESTING STRATEGY

### Tests Written in Bhasa
- tests/lexer_test.bhasa - Lexer functionality
- tests/parser_test.bhasa - Parser correctness
- tests/compiler_test.bhasa - Bytecode generation
- tests/bootstrap_test.bhasa - Self-hosting validation

### Example Programs
- Basic: hello.bhasa, variables.bhasa
- Functions: functions.bhasa, fibonacci.bhasa
- Collections: arrays.bhasa, hash.bhasa
- Control: conditionals.bhasa, loops.bhasa
- Advanced: bitwise_comprehensive.bhasa, bengali_variable_names.bhasa

---

## 14. PERFORMANCE CHARACTERISTICS

| Operation | Complexity | Notes |
|-----------|-----------|-------|
| Variable access | O(1) | Array indexing |
| Function call | O(1) | Frame push/pop |
| Method lookup | O(1) | Hash map lookup |
| Closure creation | O(n) | n = free variables |
| Array indexing | O(1) | Stack access |
| Bytecode dispatch | O(1) | Switch statement |

**Performance vs Tree-Walking:**
- Arithmetic: 10x faster
- Variables: 5x faster
- Functions: 3x faster
- **Overall: 3-10x faster**

---

## 15. QUICK LINKS

- Full Architecture: `/ARCHITECTURE.md`
- Compiler Design: `/COMPILER.md`
- Compiler API: `/COMPILER_API.md`
- Features Guide: `/FEATURES.md`
- Struct/Enum Design: `/STRUCT_ENUM_DESIGN.md`
- Quick Start: `/QUICKSTART.md`
- Usage Guide: `/USAGE.md`

