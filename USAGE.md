# ভাষা (Bhasa) - Usage Guide

## Quick Start

### Build the Interpreter

```bash
go build -o bhasa
```

### Run the REPL (Interactive Mode)

```bash
./bhasa
```

This will start an interactive shell where you can type Bengali code:

```
>> ধরি x = ৫;
>> লেখ(x);
5
>> ধরি যোগ = ফাংশন(a, b) { ফেরত a + b; };
>> লেখ(যোগ(১০, ২০));
30
```

### Run a File

```bash
./bhasa examples/hello.bhasa
```

## Language Features

### 1. Variables

Use `ধরি` (let) to declare variables:

```bengali
ধরি x = 10;
ধরি নাম = "বাংলা";
ধরি সত্য_মান = সত্য;
```

Reassign variables without `ধরি`:

```bengali
x = 20;
```

### 2. Data Types

- **Numbers**: `১০`, `২৫`, `100` (supports both Bengali and Arabic numerals)
- **Strings**: `"নমস্কার"`, `"Hello"`
- **Booleans**: `সত্য` (true), `মিথ্যা` (false)
- **Arrays**: `[১, ২, ৩, ৪, ৫]`
- **Hash Maps**: `{"নাম": "রহিম", "বয়স": ২৫}`

### 3. Operators

**Arithmetic:**
- `+` Addition
- `-` Subtraction
- `*` Multiplication
- `/` Division
- `%` Modulo

**Comparison:**
- `==` Equal
- `!=` Not equal
- `<` Less than
- `>` Greater than
- `<=` Less than or equal
- `>=` Greater than or equal

**Logical:**
- `!` Not

### 4. Functions

Declare functions with `ফাংশন`:

```bengali
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};

লেখ(যোগ(৫, ৩));  // Output: 8
```

Functions support recursion:

```bengali
ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
    যদি (n == ০) {
        ফেরত ১;
    } নাহলে {
        ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
    }
};
```

### 5. Conditionals

Use `যদি` (if) and `নাহলে` (else):

```bengali
যদি (x > ৫) {
    লেখ("Greater than 5");
} নাহলে {
    লেখ("Not greater than 5");
}
```

### 6. Loops

Use `যতক্ষণ` (while) for loops:

```bengali
ধরি i = ১;
যতক্ষণ (i <= ১০) {
    লেখ(i);
    i = i + ১;
}
```

### 7. Arrays

Create and manipulate arrays:

```bengali
ধরি তালিকা = [১, ২, ৩, ৪, ৫];
লেখ(তালিকা[০]);        // Access: 1
লেখ(প্রথম(তালিকা));    // First: 1
লেখ(শেষ(তালিকা));      // Last: 5
লেখ(দৈর্ঘ্য(তালিকা));  // Length: 5
```

### 8. Hash Maps

Create key-value pairs:

```bengali
ধরি ব্যক্তি = {
    "নাম": "রহিম",
    "বয়স": ২৫,
    "শহর": "ঢাকা"
};

লেখ(ব্যক্তি["নাম"]);  // Output: রহিম
```

## Built-in Functions

| Function | Bengali | Description | Example |
|----------|---------|-------------|---------|
| print | `লেখ(...)` | Print to console | `লেখ("Hello")` |
| length | `দৈর্ঘ্য(x)` | Get length of string/array | `দৈর্ঘ্য([১,২,৩])` |
| first | `প্রথম(arr)` | Get first element | `প্রথম([১,২,৩])` |
| last | `শেষ(arr)` | Get last element | `শেষ([১,২,৩])` |
| rest | `বাকি(arr)` | Get all but first | `বাকি([১,২,৩])` |
| push | `যোগ(arr, x)` | Add element to array | `যোগ([১,২], ৩)` |
| type | `টাইপ(x)` | Get type of value | `টাইপ(৫)` |

## Comments

Use `//` for single-line comments:

```bengali
// এটি একটি মন্তব্য (This is a comment)
ধরি x = ৫;  // Inline comment
```

## Examples

See the `examples/` directory for more examples:

- `hello.bhasa` - Hello World
- `variables.bhasa` - Variable declarations and math
- `functions.bhasa` - Function definitions including recursion
- `conditionals.bhasa` - If-else statements
- `loops.bhasa` - While loops
- `arrays.bhasa` - Array operations
- `hash.bhasa` - Hash map operations
- `fibonacci.bhasa` - Fibonacci sequence
- `comprehensive.bhasa` - All features together

## Keywords Reference

| English | Bengali | Token |
|---------|---------|-------|
| let | ধরি | Variable declaration |
| function | ফাংশন | Function definition |
| if | যদি | Conditional |
| else | নাহলে | Else clause |
| return | ফেরত | Return statement |
| true | সত্য | Boolean true |
| false | মিথ্যা | Boolean false |
| while | যতক্ষণ | While loop |

## Tips

1. **Bengali Numerals**: You can use Bengali numerals (০-৯) or Arabic numerals (0-9) interchangeably
2. **Semicolons**: Semicolons are optional in most cases but recommended for clarity
3. **String Concatenation**: Use `+` to concatenate strings: `"Hello" + " " + "World"`
4. **REPL Commands**: Type `প্রস্থান` or `exit` to quit the REPL

## Error Messages

The interpreter provides helpful error messages:

```
Parser errors:
    expected next token to be ;, got } instead
```

Check your syntax if you encounter errors!

## Contributing

Feel free to add more Bengali keywords or built-in functions to make the language even more expressive!

