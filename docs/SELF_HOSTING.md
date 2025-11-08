# ভাষা স্ব-হোস্টিং কম্পাইলার (Bhasa Self-Hosting Compiler)

## সারসংক্ষেপ (Overview)

এই প্রকল্পে ভাষা প্রোগ্রামিং ল্যাঙ্গুয়েজের একটি **সম্পূর্ণ স্ব-হোস্টিং কম্পাইলার** বাস্তবায়ন করা হয়েছে। এই কম্পাইলার সম্পূর্ণভাবে **বাংলা ভাষায়** লেখা এবং `.ভাষা` এক্সটেনশন ফাইল ব্যবহার করে।

##  স্ব-হোস্টিং (Self-Hosting) কী?

**স্ব-হোস্টিং** মানে হল একটি প্রোগ্রামিং ভাষার কম্পাইলার সেই একই ভাষায় লেখা। উদাহরণস্বরূপ:
- Go compiler Go তে লেখা
- Rust compiler Rust এ লেখা
- এখন **ভাষা compiler ভাষা তে লেখা**!

## আর্কিটেকচার

### ফেজ ১: Go বেসলাইন কম্পাইলার
বিদ্যমান Go তে লেখা কম্পাইলার যা `.ভাষা` ফাইল কম্পাইল করে বাইটকোড তৈরি করে।

### ফেজ ২: স্ব-হোস্টিং কম্পাইলার (`.ভাষা` তে লেখা)

#### মডিউল স্ট্রাকচার

```
modules/
├── টোকেন.ভাষা           # Token definitions & lexical utilities (169 lines)
├── লেক্সার.ভাষা          # Lexical analyzer (348 lines)
├── এএসটি.ভাষা            # AST node structures (358 lines)
├── পার্সার.ভাষা          # Pratt parser with precedence (879 lines)
├── প্রতীক_টেবিল.ভাষা     # Symbol table for scoping (166 lines)
├── কোড.ভাষা              # Bytecode operations (366 lines)
├── কম্পাইলার.ভাষা        # AST to bytecode compiler (946 lines)
├── মডিউল_লোডার.ভাষা      # Module loading & caching (202 lines)
└── ভাষা_কম্পাইলার.ভাষা   # Main compiler driver (326 lines)

Total: 3,760+ lines of pure Bengali code!
```

## কম্পাইলেশন পাইপলাইন

```
.ভাষা Source Code
       ↓
   [Lexer] → Tokens
       ↓
   [Parser] → AST
       ↓
  [Compiler] → Bytecode
       ↓
     [VM] → Execution
```

## প্রধান বৈশিষ্ট্য

### ১. সম্পূর্ণ বাংলা সিনট্যাক্স
- **Keywords**: `ধরি`, `ফাংশন`, `যদি`, `নাহলে`, `ফেরত`, `যতক্ষণ`, `পর্যন্ত`
- **Operators**: যোগ, বিয়োগ, গুণ, ভাগ সহ সব অপারেটর
- **Functions**: সব ফাংশন নাম বাংলায়
- **Comments**: বাংলা মন্তব্য

### ২. টাইপ সিস্টেম
- **Integer Types**: `Byte` (8-bit), `Short` (16-bit), `Int` (32-bit), `Long` (64-bit)
- **Float Types**: `Float` (32-bit), `Double` (64-bit)
- **Other Types**: `Char` (rune), `String`, `Boolean`, `Array`, `Hash`
- **Type Casting**: `বাইট()`, `পূর্ণসংখ্যা()`, `দশমিক()`, etc.

### ৩. লেক্সিক্যাল অ্যানালাইসিস
- Unicode support for Bengali characters
- Number literal recognition (both Bengali & English numerals)
- String literals with escape sequences
- Comment handling
- Position tracking (line & column)

### ৪. সিনট্যাক্স অ্যানালাইসিস
- **Pratt Parser** with operator precedence
- Statement parsing: let, return, expression, block, loops
- Expression parsing: infix, prefix, call, index
- Error reporting with line numbers

### ৫. কোড জেনারেশন
- Stack-based bytecode generation
- 40 opcodes covering all operations
- Scope management (global, local, builtin, free)
- Function closures support
- Loop constructs (break/continue)

### ৬. মডিউল সিস্টেম
- File I/O for module loading
- Circular dependency detection
- Module caching
- Relative & absolute path support

## ব্যবহার উদাহরণ

### উদাহরণ ১: সাধারণ প্রোগ্রাম

```bengali
// simple.ভাষা
ধরি x = ৫;
ধরি y = ১০;
ধরি যোগফল = x + y;
লেখ(যোগফল);  // Output: ১৫
```

### উদাহরণ ২: ফাংশন সংজ্ঞা

```bengali
// function.ভাষা
ধরি গুণ = ফাংশন(a, b) {
    ফেরত a * b;
};

ধরি ফলাফল = গুণ(৫, ৩);
লেখ(ফলাফল);  // Output: ১৫
```

### উদাহরণ ৩: কন্ডিশনাল লজিক

```bengali
// condition.ভাষা
ধরি স্কোর = ৮৫;

যদি (স্কোর >= ৮০) {
    লেখ("চমৎকার!");
} নাহলে যদি (স্কোর >= ৬০) {
    লেখ("ভালো");
} নাহলে {
    লেখ("আরো চেষ্টা করো");
}
```

### উদাহরণ ৪: লুপ

```bengali
// loop.ভাষা
পর্যন্ত (ধরি i = ০; i < ১০; i = i + ১) {
    লেখ(i);
}
```

## কম্পাইলার চালানো

### পদ্ধতি ১: Go VM থেকে
```bash
# Run the self-hosted compiler via Go VM
go run cmd/main.go modules/ভাষা_কম্পাইলার.ভাষা
```

### পদ্ধতি ২: টেস্ট চালানো
```bash
# Run lexer tests
go run cmd/main.go tests/lexer_test.ভাষা

# Run parser tests
go run cmd/main.go tests/parser_test.ভাষা

# Run compiler tests
go run cmd/main.go tests/compiler_test.ভাষা

# Run bootstrap test
go run cmd/main.go tests/bootstrap_test.ভাষা
```

## বুটস্ট্র্যাপিং প্রসেস

### ধাপ ১: বেসলাইন কম্পাইলার (Go)
Go তে লেখা কম্পাইলার `.ভাষা` ফাইল কম্পাইল করতে পারে।

### ধাপ ২: স্ব-হোস্টিং কম্পাইলার লেখা
`.ভাষা` তে সম্পূর্ণ কম্পাইলার টুলচেইন লেখা হয়েছে।

### ধাপ ৩: বুটস্ট্র্যাপিং
স্ব-হোস্টিং কম্পাইলার নিজেকে কম্পাইল করতে পারে:

```
ভাষা_কম্পাইলার.ভাষা → [Go Compiler] → Bytecode
Bytecode → [Go VM] → Running Self-Hosted Compiler
Self-Hosted Compiler → Compiles any .ভাষা file!
```

## পরিসংখ্যান

- **Total Code Lines**: 3,760+ lines in Bengali
- **Modules**: 9 major modules
- **Opcodes**: 40 bytecode operations
- **Test Files**: 4 comprehensive test suites
- **Language Features**: Full Bengali syntax support

## প্রযুক্তিগত বিস্তারিত

### টোকেন টাইপ
- Keywords: 20+
- Operators: 25+
- Literals: Integer, Float, String
- Delimiters: `()`, `{}`, `[]`, `;`, `,`

### অপারেটর অগ্রাধিকার (14 স্তর)
1. Lowest
2. Logical OR (`||`)
3. Logical AND (`&&`)
4. Bitwise OR (`|`)
5. Bitwise XOR (`^`)
6. Bitwise AND (`&`)
7. Equality (`==`, `!=`)
8. Comparison (`<`, `>`, `<=`, `>=`)
9. Shift (`<<`, `>>`)
10. Addition (`+`, `-`)
11. Multiplication (`*`, `/`, `%`)
12. Prefix (`-x`, `!x`, `~x`)
13. Call (`func()`)
14. Index (`array[i]`)

### বাইটকোড নির্দেশ
- **Stack Operations**: Push, Pop
- **Arithmetic**: Add, Sub, Mul, Div, Mod
- **Bitwise**: And, Or, Xor, Not, Shift
- **Logic**: And, Or, Not
- **Comparison**: Equal, NotEqual, GreaterThan, etc.
- **Control Flow**: Jump, JumpNotTruthy
- **Variables**: GetGlobal, SetGlobal, GetLocal, SetLocal
- **Functions**: Call, Return, Closure
- **Data Structures**: Array, Hash, Index

## সীমাবদ্ধতা

1. **File I/O**: Module loader এর ফাইল পড়া Go দ্বারা বাস্তবায়িত হতে হবে
2. **VM Execution**: Bytecode execution এখনো Go VM এ হয়
3. **Standard Library**: Limited builtin functions
4. **Error Messages**: Basic error reporting

## ভবিষ্যৎ উন্নতি

- [ ] Full VM in `.ভাষা`
- [ ] Advanced error reporting
- [ ] Optimization passes
- [ ] JIT compilation support
- [ ] Standard library expansion
- [ ] IDE integration (LSP)
- [ ] Debugging tools




---


