# Bhasa Programming Language - AI Agent Instructions

## Project Overview
Bhasa (ভাষা) is a **compiled Bengali programming language** featuring a full toolchain: lexer → parser → compiler → stack-based VM. The language compiles to bytecode (3-10x faster than interpretation) and supports closures, recursion, and a self-hosted compiler written entirely in Bhasa.

## Core Architecture

### Compilation Pipeline
```
.bhasa/.ভাষা source → Lexer → Parser → AST → Compiler → Bytecode → VM → Execution
```

**Key packages** (9 total):
- `token/` - Token type definitions (Bengali keywords)
- `lexer/` - UTF-8 tokenizer with Bengali numeral support
- `parser/` - Pratt parser with operator precedence
- `ast/` - Abstract syntax tree node definitions
- `compiler/` - Single-pass AST to bytecode compiler
- `code/` - 41 bytecode opcodes (including bitwise ops)
- `vm/` - Stack-based VM with 2048-element stack
- `object/` - Runtime object system (Integer, String, Array, Hash, Closure, Struct, Enum)
- `evaluator/` - Legacy tree-walking interpreter (kept for reference)

### Critical Implementation Files
- `compiler/compiler.go` (1147 lines) - Bytecode generation, symbol table, loop tracking
- `vm/vm.go` (1172 lines) - Stack execution with call frames
- `parser/parser.go` (1169 lines) - Pratt parsing with prefix/infix functions
- `ast/ast.go` (654 lines) - 20+ AST node types

## Development Workflows

### Building
```bash
make build              # Current platform
make all                # All platforms (Linux/Windows/macOS ARM64/AMD64)
make linux-amd64        # Specific platform
make clean              # Remove artifacts
```

Binaries output to `bin/` with format: `bhasa-{OS}-{ARCH}[.exe]`

### Running Programs
```bash
./bhasa                          # Start REPL
./bhasa program.bhasa            # Run source file
./bhasa -c program.bhasa         # Compile to .compiled bytecode
./bhasa program.compiled         # Execute pre-compiled bytecode
./bhasa -h                       # Show help
```

**File extensions**: `.bhasa` or `.ভাষা` (source), `.compiled` or `.সংকলিত` (bytecode)

### Testing
```bash
./run_examples.sh               # Run all 26+ example programs
go test -v ./...                # Go unit tests
./bhasa tests/lexer_test.ভাষা   # Self-hosted tests (written in Bhasa!)
```

**Test files** in `tests/`: lexer_test, parser_test, compiler_test, bootstrap_test, struct_test, enum_test, type_system_test

**Example programs** in `examples/`: hello, variables, functions, fibonacci, arrays, hash, loops, conditionals, bitwise_comprehensive, bengali_variable_names, file_io, json_support

## Language-Specific Conventions

### Bengali Keywords & Syntax
The language uses Bengali keywords exclusively. Key mappings:

| Bengali | English | Usage |
|---------|---------|-------|
| `ধরি` | let | Variable declaration |
| `ফাংশন` | function | Function literal |
| `যদি`/`নাহলে` | if/else | Conditionals |
| `ফেরত` | return | Return statement |
| `যতক্ষণ` | while | While loop |
| `পর্যন্ত` | for | C-style for loop |
| `বিরতি`/`চালিয়ে_যাও` | break/continue | Loop control |
| `অন্তর্ভুক্ত` | import | Module import |
| `স্ট্রাক্ট`/`গণনা` | struct/enum | OOP features |
| `লেখ()` | print | Built-in output function |

**Bengali numerals**: ০-৯ automatically converted to 0-9 during lexing

**Variable names**: Full Unicode support - identifiers like `ব্যক্তি১`, `যোগফল`, `বেতন` are valid

### Built-in Functions (30+)
- **I/O**: `লেখ()`
- **String**: `বিভক্ত(), যুক্ত(), উপরে(), নিচে(), ছাঁটো(), প্রতিস্থাপন(), খুঁজুন()`
- **Array**: `দৈর্ঘ্য(), প্রথম(), শেষ(), বাকি(), যোগ(), উল্টাও()`
- **Math**: `শক্তি(), বর্গমূল(), পরম(), সর্বোচ্চ(), সর্বনিম্ন(), গোলাকার()`
- **File I/O**: `ফাইল_পড়ো(), ফাইল_লেখো(), ফাইল_যোগ(), ফাইল_আছে()`
- **Character ops**: `অক্ষর(), কোড(), অক্ষর_থেকে_কোড(), সংখ্যা(), লেখা()`
- **Type casting**: `বাইট(), ছোট_সংখ্যা(), পূর্ণসংখ্যা(), দীর্ঘ_সংখ্যা(), দশমিক(), দশমিক_দ্বিগুণ()`

## Key Patterns & Conventions

### 1. Lexer UTF-8 Handling
The lexer (`lexer/lexer.go`) operates on **runes**, not bytes:
```go
type Lexer struct {
    input []rune  // NOT []byte - handles multi-byte Bengali characters
    position int
    readPosition int
    ch rune       // Current character
}
```
- Bengali vowel signs (মাত্রা) U+0981-U+09CD are part of identifiers
- `isBengaliDigit()` helper allows ০-৯ in identifiers (e.g., `ব্যক্তি১`)
- Line/column tracking for error reporting

### 2. Parser Operator Precedence
Uses **Pratt parsing** with precedence levels (lowest to highest):
```
LOGICAL_OR (||) < LOGICAL_AND (&&) < EQUALS (==, !=) < 
LESSGREATER (<, >, <=, >=) < BIT_OR (|) < BIT_XOR (^) < 
BIT_AND (&) < SHIFT (<<, >>) < SUM (+, -) < PRODUCT (*, /, %) < 
PREFIX (!, -, ~) < CALL (function()) < INDEX ([i], .field)
```
Register new operators by adding to `infixParseFns` or `prefixParseFns` maps.

### 3. Compiler Symbol Resolution
Variables resolved in this order:
1. **Local scope** (current function) - `OpGetLocal`
2. **Enclosing scopes** (captured) - `OpGetFree` (for closures)
3. **Global scope** (module-level) - `OpGetGlobal`
4. **Built-in functions** - `OpGetBuiltin`

**Important**: Symbol table is hierarchical - call `c.enterScope()` before compiling functions, `c.leaveScope()` after.

### 4. Loop Context Management
When compiling loops (`যতক্ষণ`, `পর্যন্ত`), push/pop `LoopContext`:
```go
c.loopStack = append(c.loopStack, LoopContext{
    loopStart: startPos,
    breakPositions: []int{},
    contPositions: []int{},
})
defer func() { c.loopStack = c.loopStack[:len(c.loopStack)-1] }()
```
Used for jump patching in `বিরতি` (break) and `চালিয়ে_যাও` (continue) statements.

### 5. Module System
Modules use `অন্তর্ভুক্ত "path"` syntax:
- **Module loader**: `compiler/compiler.go` has `LoadAndCompileModule()` function
- **Circular import prevention**: `moduleCache` map tracks loaded modules
- **Module files**: Located in `modules/` with `.ভাষা` extension
- **Self-hosted compiler modules**: `টোকেন.ভাষা`, `লেক্সার.ভাষা`, `পার্সার.ভাষা`, `কম্পাইলার.ভাষা`, etc.

### 6. Struct & Enum (OOP Features)
**Status**: Partially implemented (parser/compiler ready, VM has bugs)

**Struct syntax**:
```bengali
ধরি Person = স্ট্রাক্ট {নাম: "রহিম", বয়স: 25};
লেখ(Person.নাম);  // Field access
Person.বয়স = 30;  // Field assignment
```

**Enum syntax**:
```bengali
ধরি Direction = গণনা { North, South, East, West };
ধরি dir = Direction.North;
```

**Opcodes**: `OpStruct`, `OpGetStructField`, `OpSetStructField`, `OpEnum`
**Known issue**: VM crashes on `OpNewInstance` (line 419-473 in `vm/vm.go`)

### 7. Error Handling
Error messages in Bengali (`errors/bengali_errors.go`):
```go
ErrUnexpectedToken = "অপ্রত্যাশিত টোকেন"
ErrUndefinedVariable = "অসংজ্ঞায়িত ভেরিয়েবল: %s"
```
Parser errors tracked in `p.errors` slice - check `len(p.Errors())` after parsing.

## Documentation Structure

- `docs/ARCHITECTURE.md` - Complete 800+ line architecture doc
- `docs/COMPILER.md` - Bytecode compilation details
- `docs/FEATURES.md` - Language features and implementation status
- `docs/SELF_HOSTING.md` - Self-hosted compiler guide
- `docs/OOP_STATUS.md` - Current OOP implementation bugs
- `RESERVED_BENGALI_KEYWORDS.txt` - List of reserved keywords
- Package-specific docs in `{package}/docs/README.md`

## Common Tasks

### Adding a New Bytecode Instruction
1. Define opcode in `code/code.go`: `const OpMyOp Opcode = 0x??`
2. Add to `definitions` map with operand widths
3. Compile in `compiler/compiler.go`: emit via `c.emit(code.OpMyOp, operands...)`
4. Execute in `vm/vm.go`: add case in switch statement
5. Test in `tests/` with Bhasa test file

### Adding a New Built-in Function
1. Add to `object.Builtins` slice in `object/object.go`
2. Implement function body with argument validation
3. Bengali name convention: use complete Bengali words (e.g., `যোগফল`, not `যোগ` if `যোগ` already exists)
4. Test in `examples/` directory

### Debugging Compilation Issues
```bash
# 1. Check lexer output (manual inspection)
# 2. Verify parser AST with String() method
# 3. Inspect bytecode: see compiler/serializer.go for disassembly helpers
# 4. Trace VM execution: add debug prints in vm/vm.go Run() loop
```

**Bytecode inspection**: Compiled bytecode stored in `Bytecode.Instructions` ([]byte) and `Bytecode.Constants` ([]object.Object)

## Anti-Patterns to Avoid

1. **Don't use byte slicing on source code** - Always work with runes for Bengali text
2. **Don't add Bengali keywords to Go identifiers** - Keep Go code in English
3. **Don't forget symbol scope management** - Always match `enterScope()` with `leaveScope()`
4. **Don't shadow built-in functions** - Check `object.Builtins` before naming variables
5. **Don't modify evaluator/** - Legacy interpreter; all new features go to compiler/VM
6. **Don't use `লেখা` as TYPE_STRING** - Changed to `পাঠ্য` (use `লেখা` for toString function only)

## Performance Characteristics

- **Compilation**: O(n) single-pass
- **VM execution**: 3-10x faster than tree-walking interpreter
- **Variable access**: O(1) via array indexing (locals) or array lookup (globals)
- **Function calls**: O(1) frame push/pop
- **Stack depth**: 2048 elements (configurable in `vm/vm.go`)
- **Global variables**: 65,536 slots

## Repository Context

- **Branch**: `development` (main development branch)
- **Owner**: Uttam-Mahata
- **Go version**: See `go.mod`
- **Binary naming**: `bhasa-{os}-{arch}` (e.g., `bhasa-linux-amd64`)
- **License**: See LICENSE file

## Quick Reference

**Most-edited files for feature additions**:
1. `parser/parser.go` - Add new syntax parsing
2. `compiler/compiler.go` - Add bytecode generation
3. `vm/vm.go` - Add opcode execution
4. `token/token.go` - Add new keywords/operators
5. `ast/ast.go` - Add new AST nodes

**When adding Bengali keywords**: Update lexer keyword map (`token.LookupIdent()`)

**When debugging**: Check line/column info in tokens - lexer tracks position for error messages

---

*This is a production-ready compiled language with ~5000+ lines of Go code, 30+ built-in functions, 41 opcodes, and a self-hosted compiler. Treat it as a serious language implementation, not a toy project.*
