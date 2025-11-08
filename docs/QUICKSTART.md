# ভাষা (Bhasa) Quick Start Guide

## 🚀 Get Started in 2 Minutes

### 1. Build the Compiler

```bash
go build -o bhasa
```

**Note**: Bhasa is now a compiled language! Your code is compiled to bytecode and executed on a fast virtual machine (3-10x faster than interpretation).

### 2. Try the REPL

```bash
./bhasa
```

Now type some Bengali code:

```bengali
>> ধরি x = ১০;
>> ধরি y = ২০;
>> লেখ(x + y);
30
>> ধরি যোগ = ফাংশন(a, b) { ফেরত a + b; };
>> লেখ(যোগ(৫, ৩));
8
```

### 3. Run Your First Program

Create a file `my_program.bhasa`:

```bengali
// আমার প্রথম প্রোগ্রাম
লেখ("নমস্কার!");

ধরি নাম = "বাংলা";
লেখ("আমার নাম: " + নাম);

ধরি সংখ্যা = ৫;
ধরি দ্বিগুণ = সংখ্যা * ২;
লেখ("দ্বিগুণ: ");
লেখ(দ্বিগুণ);
```

Run it:

```bash
./bhasa my_program.bhasa
```

### 4. Explore Examples

```bash
./bhasa examples/hello.bhasa        # Hello World
./bhasa examples/fibonacci.bhasa    # Fibonacci sequence
./bhasa examples/comprehensive.bhasa # All features
```

Or run all examples at once:

```bash
./run_examples.sh
```

## 📚 Learn More

- **README.md** - Project overview
- **USAGE.md** - Complete language guide
- **COMPILER.md** - Compiler architecture and bytecode
- **FEATURES.md** - Technical details
- **examples/** - 9 example programs

## ⚡ Performance

Bhasa uses a **bytecode compiler** and **stack-based VM** for fast execution:
- **3-10x faster** than tree-walking interpretation
- Efficient variable access (array indexing)
- Optimized function calls (call frames)
- Full closure support with free variables

## 🎯 Key Concepts

### Variables
```bengali
ধরি x = ১০;    // Declare
x = ২০;        // Reassign
```

### Functions
```bengali
ধরি ফাংশন_নাম = ফাংশন(param1, param2) {
    ফেরত param1 + param2;
};
```

### Conditionals
```bengali
যদি (শর্ত) {
    // code
} নাহলে {
    // code
}
```

### Loops
```bengali
যতক্ষণ (শর্ত) {
    // code
}
```

### Arrays
```bengali
ধরি তালিকা = [১, ২, ৩];
লেখ(তালিকা[০]);
```

## 🔑 Essential Functions

- `লেখ(x)` - Print
- `দৈর্ঘ্য(x)` - Length
- `প্রথম(arr)` - First element
- `শেষ(arr)` - Last element

## 💡 Tips

1. You can use both Bengali (০-৯) and Arabic (0-9) numerals
2. Semicolons are optional but recommended
3. Type `প্রস্থান` or `exit` to quit REPL
4. Comments start with `//`

## 🎉 Happy Coding!

Start creating amazing programs in Bengali!

```bengali
ধরি শুভেচ্ছা = "শুভ প্রোগ্রামিং!";
লেখ(শুভেচ্ছা);
```

