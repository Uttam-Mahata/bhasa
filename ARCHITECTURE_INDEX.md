# Bhasa Codebase Documentation Index

## Quick Navigation

This directory contains comprehensive documentation about the Bhasa language architecture and codebase structure. Use this index to find what you're looking for.

---

## 1. START HERE

### For Quick Understanding
- **[CODEBASE_MAP.txt](CODEBASE_MAP.txt)** - ASCII diagram of entire architecture with all components
- **[ARCHITECTURE_SUMMARY.md](ARCHITECTURE_SUMMARY.md)** - Quick reference guide (15 sections)

### For Deep Dive
- **[ARCHITECTURE.md](ARCHITECTURE.md)** - Comprehensive 929-line guide covering everything (14 sections)

---

## 2. UNDERSTANDING THE SYSTEM

### How It Works
1. **[Lexer](token/)** - Tokenizes Bengali source code (UTF-8 support)
2. **[Parser](parser/)** - Builds Abstract Syntax Tree using Pratt parsing
3. **[Compiler](compiler/)** - Converts AST to bytecode instructions
4. **[Virtual Machine](vm/)** - Executes bytecode on stack-based architecture

### Key Concepts
- **Tokens** - 41 token types including Bengali keywords
- **AST** - 20+ node types for all language constructs
- **Bytecode** - 41 opcodes for compilation target
- **Symbol Table** - Variable scoping with 5 scope types
- **Stack-Based VM** - 2,048 element stack with 65,536 global slots

---

## 3. LANGUAGE FEATURES

### Current Features
- Bengali keywords (ধরি, ফাংশন, যদি, etc.)
- Variables, functions, closures, recursion
- Arrays, hash maps, multiple numeric types
- Control flow (if-else, while, for, break, continue)
- 30+ built-in functions
- File I/O operations
- Bitwise operators
- Basic struct and enum support

### Documentation
- **[FEATURES.md](FEATURES.md)** - Feature overview and examples
- **[USAGE.md](USAGE.md)** - Language usage guide
- **[QUICKSTART.md](QUICKSTART.md)** - 2-minute quick start

---

## 4. OOP FEATURES

### Current OOP Status
- Structs: Fully working (definition, creation, field access)
- Enums: Partially working (simple enums, basic variants)
- Methods: Not yet implemented
- Inheritance: Not yet implemented
- Interfaces: Not yet implemented

### OOP Design Documents
- **[STRUCT_ENUM_DESIGN.md](STRUCT_ENUM_DESIGN.md)** - Detailed OOP roadmap
  - Struct design with examples
  - Enum design with variants
  - Future extensions (interfaces, inheritance)
  - Implementation phases and timeline

---

## 5. COMPILER & VM DETAILS

### Technical Deep Dives
- **[COMPILER.md](COMPILER.md)** - Compiler architecture and bytecode design
- **[COMPILER_API.md](COMPILER_API.md)** - Self-hosted compiler API documentation

### Self-Hosted Compiler
The Bhasa compiler is written in Bhasa itself!

Self-hosting modules (in `/modules/`):
- লেক্সার.ভাষা - Lexer in Bhasa
- পার্সার.ভাষা - Parser in Bhasa  
- কম্পাইলার.ভাษা - Compiler in Bhasa
- এবং আরও অনেক কিছু...

---

## 6. CODEBASE STRUCTURE

### By File Size
1. **vm/vm.go** (1,172 lines) - Stack-based bytecode executor
2. **parser/parser.go** (1,169 lines) - Pratt parser with operator precedence
3. **compiler/compiler.go** (868 lines) - AST to bytecode compiler
4. **ast/ast.go** (654 lines) - 20+ AST node types
5. **object/object.go** (370 lines) - Runtime object system
6. **code/code.go** (200 lines) - 41 bytecode instructions
7. **lexer/lexer.go** (210 lines) - UTF-8 lexical analyzer
8. **token/token.go** (164 lines) - Token type definitions
9. **Others** (~90 lines) - REPL, main, symbol table

**Total: 3,860 lines of Go code**

### By Responsibility

#### Phase 1: Tokenization
```
token/token.go        - Token type definitions
lexer/lexer.go        - Lexical analyzer (tokenizer)
```

#### Phase 2: Parsing  
```
parser/parser.go      - Pratt parser (1,169 lines)
ast/ast.go            - AST node definitions (654 lines)
```

#### Phase 3: Compilation
```
compiler/compiler.go       - AST → Bytecode (868 lines)
compiler/symbol_table.go   - Variable scoping
code/code.go               - Instruction definitions
```

#### Phase 4: Execution
```
vm/vm.go              - Stack-based VM (1,172 lines)
vm/frame.go           - Function call frames
object/object.go      - Runtime objects (370 lines)
```

---

## 7. EXTENDING THE LANGUAGE

### To Add Methods to Structs
See **ARCHITECTURE_SUMMARY.md Section 11** for required changes:
1. Token definitions
2. Parser modifications
3. AST node additions
4. Compiler changes
5. Object system updates
6. VM instruction handlers

### To Enhance Enums
Similar 5-step process focusing on:
- Associated data storage
- Pattern matching
- Data extraction

### To Add Interfaces
Follows established pattern with focus on:
- Type contracts
- Interface checking
- Polymorphism support

---

## 8. QUICK REFERENCE

### Keywords (20 total)
**Control Flow:** ধরি, ফাংশন, যদি, নাহলে, ফেরত, যতক্ষণ, পর্যন্ত, বিরতি, চালিয়ে_যাও
**Literals:** সত্য, মিথ্যা, নাল
**Types:** বাইট, ছোট_সংখ্যা, পূর্ণসংখ্যা, দীর্ঘ_সংখ্যা, দশমিক, দশমিক_দ্বিগুণ, অক্ষর, লেখা, বুলিয়ান, তালিকা, ম্যাপ
**OOP:** স্ট্রাক্ট, গণনা, হিসাবে

### Operators (20+)
**Arithmetic:** + - * / %
**Comparison:** == != < > <= >=
**Logical:** ! && ||
**Bitwise:** & | ^ ~ << >>
**Assignment:** =
**Access:** .

### Built-in Functions (30+)
See **ARCHITECTURE_SUMMARY.md Section 10** for complete list:
- I/O: লেখ
- String: বিভক্ত, যুক্ত, উপরে, নিচে, ছাঁটো, প্রতিস্থাপন, খুঁজুন
- Array: দৈর্ঘ্য, প্রথম, শেষ, বাকি, যোগ, উল্টাও
- Math: শক্তি, বর্গমূল, পরম, সর্বোচ্চ, সর্বনিম্ন, গোলাকার
- Type: টাইপ, বাইট, ছোট_সংখ্যা, পূর্ণসংখ্যা, দীর্ঘ_সংখ্যা, দশমিক, দশমিক_দ্বিগুণ
- File I/O: ফাইল_পড়ো, ফাইল_লেখো, ফাইল_যোগ, ফাইল_আছে
- Character: অক্ষর, কোড, অক্ষর_থেকে_কোড, সংখ্যা, লেখা

---

## 9. PERFORMANCE

### Execution Speed
- Arithmetic operations: **10x faster** than tree-walking
- Variable access: **5x faster**
- Function calls: **3x faster**
- **Overall: 3-10x faster** than interpretation

### Memory Layout
- Stack: 2,048 elements
- Global storage: 65,536 slots
- Call frames: 1,024 maximum
- Constants pool: 65,536 maximum

### Key Optimizations
- O(1) variable access via array indexing
- O(1) stack operations
- Single-pass compilation
- Frame-based function calls
- Closure support with free variable capture

---

## 10. TESTING

### Tests Included
- `tests/lexer_test.bhasa` - Lexer tests (in Bhasa!)
- `tests/parser_test.bhasa` - Parser tests
- `tests/compiler_test.bhasa` - Compiler tests
- `tests/bootstrap_test.bhasa` - Self-hosting tests

### Example Programs (10+)
Located in `examples/`:
- hello.bhasa - Hello World
- variables.bhasa - Variable declarations
- functions.bhasa - Function definitions and recursion
- fibonacci.bhasa - Recursive fibonacci
- arrays.bhasa - Array operations
- hash.bhasa - Hash map operations
- conditionals.bhasa - If-else statements
- loops.bhasa - For and while loops
- bitwise_comprehensive.bhasa - Bitwise operators
- bengali_variable_names.bhasa - Bengali identifiers

---

## 11. IMPORTANT FILES TO KNOW

### Core Components
| File | Lines | Purpose |
|------|-------|---------|
| parser/parser.go | 1,169 | Pratt parser |
| vm/vm.go | 1,172 | Stack-based VM |
| compiler/compiler.go | 868 | AST to bytecode |
| ast/ast.go | 654 | AST node types |
| object/object.go | 370 | Runtime objects |
| code/code.go | 200 | Bytecode instructions |

### Supporting Files
| File | Purpose |
|------|---------|
| compiler/symbol_table.go | Variable scoping |
| vm/frame.go | Call frames |
| lexer/lexer.go | Tokenization |
| token/token.go | Token types |
| main.go | Entry point |
| repl/repl.go | Interactive shell |

---

## 12. ARCHITECTURE OVERVIEW

### Compilation Pipeline
```
Bengali Source Code
        ↓
    [LEXER]          → Tokenize
        ↓
    [PARSER]         → Build AST
        ↓
    [COMPILER]       → Generate Bytecode
        ↓
    [VM]             → Execute
        ↓
   OUTPUT
```

### Design Patterns
- **Pratt Parsing** - Operator precedence parsing
- **Visitor Pattern** - AST traversal
- **Symbol Table** - Scope chain hierarchy
- **Jump Patching** - Control flow compilation
- **Free Variables** - Closure implementation
- **Stack Machine** - Efficient execution

---

## 13. DEVELOPMENT ROADMAP

### Recommended OOP Implementation Order

#### Phase 1: Enhanced Structs (1-2 weeks)
- Struct type definitions
- Struct comparison operators
- Nested struct definitions
- Field default values

#### Phase 2: Struct Methods (1-2 weeks)
- Method definitions
- Method dispatch
- Receiver binding
- Method chaining

#### Phase 3: Enhanced Enums (1 week)
- Associated data
- Data storage
- Pattern matching (basic)

#### Phase 4: Interfaces (2 weeks)
- Interface syntax
- Interface checking
- Polymorphism

#### Phase 5: Type System (2 weeks)
- Null coalescing
- Option types
- Result types

---

## 14. WHERE TO START

### If You Want to...

**Understand the architecture:**
1. Start with **CODEBASE_MAP.txt**
2. Read **ARCHITECTURE_SUMMARY.md**
3. Dive into **ARCHITECTURE.md**

**Add OOP features:**
1. Read **STRUCT_ENUM_DESIGN.md**
2. Study **ARCHITECTURE_SUMMARY.md Section 11**
3. Review relevant source files

**Fix a bug:**
1. Check **ARCHITECTURE_SUMMARY.md Section 1** for pipeline
2. Identify which phase the bug is in
3. Look at corresponding file

**Optimize code:**
1. Check **ARCHITECTURE_SUMMARY.md Section 14**
2. Review **vm/vm.go** and **compiler/compiler.go**
3. Consider bytecode optimizations

**Add a builtin function:**
1. Check **object/object.go** for existing builtins
2. Add Bengali keyword to **token/token.go** if needed
3. Implement in **evaluator/builtins.go** and **object/object.go**

**Extend the type system:**
1. Add token to **token/token.go**
2. Add AST node to **ast/ast.go**
3. Update **parser/parser.go**
4. Update **compiler/compiler.go**
5. Add object type to **object/object.go**
6. Add VM handler to **vm/vm.go**

---

## 15. KEY INSIGHTS

### The Architecture is Well-Designed Because:
1. **Clear phase separation** - Lexer → Parser → Compiler → VM
2. **Proper abstraction levels** - Each phase handles one concern
3. **Good data structures** - Symbol table, bytecode, stack frames
4. **Extensibility** - Easy to add new features
5. **Performance** - 3-10x faster than tree-walking interpreter
6. **Maintainability** - Modular, focused files

### Foundation for OOP:
- Struct and enum support already in place
- Symbol table ready for method resolution
- Object system has room for methods and inheritance
- Bytecode has OOP instruction placeholders
- VM has partial OOP support

---

## 16. FILES CREATED BY THIS EXPLORATION

New documentation files created:
1. **ARCHITECTURE.md** (929 lines) - Complete architectural reference
2. **ARCHITECTURE_SUMMARY.md** - Quick reference guide (15 sections)
3. **CODEBASE_MAP.txt** - ASCII diagram of entire system
4. **ARCHITECTURE_INDEX.md** - This file

These provide comprehensive documentation for understanding and extending Bhasa.

---

## Quick Links to Documentation

| Document | Purpose | Read Time |
|----------|---------|-----------|
| CODEBASE_MAP.txt | Visual architecture | 5 min |
| ARCHITECTURE_SUMMARY.md | Quick reference | 10 min |
| ARCHITECTURE.md | Complete guide | 30 min |
| STRUCT_ENUM_DESIGN.md | OOP roadmap | 15 min |
| COMPILER.md | Compiler details | 20 min |
| FEATURES.md | Language features | 15 min |

---

**Happy exploring! The Bhasa codebase is well-structured and ready for OOP enhancements.**
