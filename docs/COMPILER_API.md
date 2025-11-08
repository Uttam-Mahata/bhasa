# ভাষা কম্পাইলার API ডকুমেন্টেশন

## সারসংক্ষেপ

এই ডকুমেন্টে ভাষা স্ব-হোস্টিং কম্পাইলারের সম্পূর্ণ API বর্ণনা করা হয়েছে। প্রতিটি মডিউলের ফাংশন, ডেটা স্ট্রাকচার এবং ব্যবহার উদাহরণ সহ বিস্তারিত তথ্য দেওয়া আছে।

---

## টোকেন মডিউল (`modules/টোকেন.ভাষা`)

### ডেটা স্ট্রাকচার

#### টোকেন
```bengali
{
    "টাইপ": স্ট্রিং,      // Token type
    "লিটারেল": স্ট্রিং,    // Literal value
    "লাইন": সংখ্যা,        // Line number
    "কলাম": সংখ্যা         // Column number
}
```

### প্রধান ফাংশন

#### `নতুন_টোকেন(টাইপ, লিটারেল, লাইন, কলাম)`
নতুন টোকেন তৈরি করে।

**Parameters:**
- `টাইপ`: Token type string
- `লিটারেল`: Literal value
- `লাইন`: Line number
- `কলাম`: Column number

**Returns:** Token object

**Example:**
```bengali
ধরি টোকেন = নতুন_টোকেন(পরিচয়, "x", ১, ০);
```

#### `শনাক্তকারী_খুঁজো(নাম)`
শনাক্তকারী keyword কিনা পরীক্ষা করে।

**Parameters:**
- `নাম`: Identifier name

**Returns:** Token type (keyword বা identifier)

#### `বাংলা_সংখ্যা_রূপান্তর(লেখা)`
বাংলা সংখ্যাকে ইংরেজি সংখ্যায় রূপান্তর করে।

---

## লেক্সার মডিউল (`modules/লেক্সার.ভাষা`)

### ডেটা স্ট্রাকচার

#### লেক্সার
```bengali
{
    "ইনপুট": স্ট্রিং,
    "অবস্থান": সংখ্যা,
    "পড়ার_অবস্থান": সংখ্যা,
    "বর্তমান_অক্ষর": স্ট্রিং,
    "লাইন": সংখ্যা,
    "কলাম": সংখ্যা
}
```

### প্রধান ফাংশন

#### `নতুন_লেক্সার(ইনপুট)`
নতুন lexer তৈরি করে।

**Parameters:**
- `ইনপুট`: Source code string

**Returns:** Lexer object

#### `পরবর্তী_টোকেন(লেক্সার)`
পরবর্তী token পায়।

**Returns:** 
```bengali
{
    "টোকেন": Token object,
    "লেক্সার": Updated lexer
}
```

#### `সব_টোকেন_করো(লেক্সার)`
সব tokens তৈরি করে।

**Returns:**
```bengali
{
    "টোকেন": Array of tokens,
    "লেক্সার": Updated lexer
}
```

---

## পার্সার মডিউল (`modules/পার্সার.ভাষা`)

### ডেটা স্ট্রাকচার

#### পার্সার
```bengali
{
    "লেক্সার": Lexer object,
    "বর্তমান_টোকেন": Token,
    "পরবর্তী_টোকেন": Token,
    "ত্রুটি": Array of strings
}
```

### প্রধান ফাংশন

#### `নতুন_পার্সার(লেক্সার)`
নতুন parser তৈরি করে।

#### `পার্স_প্রোগ্রাম(পার্সার)`
সম্পূর্ণ program parse করে।

**Returns:**
```bengali
{
    "প্রোগ্রাম": AST Program node,
    "পার্সার": Updated parser
}
```

#### `পার্স_বিবৃতি(পার্সার)`
একটি statement parse করে।

#### `পার্স_অভিব্যক্তি(পার্সার, অগ্রাধিকার)`
Pratt parsing দিয়ে expression parse করে।

---

## কম্পাইলার মডিউল (`modules/কম্পাইলার.ভাষা`)

### ডেটা স্ট্রাকচার

#### কম্পাইলার
```bengali
{
    "ধ্রুবক": Array,              // Constants pool
    "প্রতীক_টেবিল": SymbolTable,  // Symbol table
    "স্কোপ_তালিকা": Array,        // Compilation scopes
    "স্কোপ_সূচক": সংখ্যা,          // Current scope index
    "লুপ_স্ট্যাক": Array          // Loop stack for break/continue
}
```

### প্রধান ফাংশন

#### `নতুন_কম্পাইলার()`
নতুন compiler তৈরি করে।

**Returns:** Compiler object

#### `কম্পাইল(কম্পাইলার, নোড)`
AST node compile করে bytecode এ।

**Parameters:**
- `কম্পাইলার`: Compiler object
- `নোড`: AST node

**Returns:**
```bengali
{
    "সফল": Boolean,
    "কম্পাইলার": Updated compiler,
    "ত্রুটি": Error message (if any)
}
```

#### `বাইটকোড_পাও(কম্পাইলার)`
Final bytecode পায়।

**Returns:**
```bengali
{
    "নির্দেশ": Bytecode instructions array,
    "ধ্রুবক": Constants array
}
```

---

## কোড মডিউল (`modules/কোড.ভাষা`)

### অপকোড তালিকা

| Opcode | Value | Description |
|--------|-------|-------------|
| অপ_ধ্রুবক | 0 | Load constant |
| অপ_পপ | 1 | Pop from stack |
| অপ_যোগ | 2 | Addition |
| অপ_বিয়োগ | 3 | Subtraction |
| অপ_গুণ | 4 | Multiplication |
| অপ_ভাগ | 5 | Division |
| অপ_ভাগশেষ | 6 | Modulo |
| অপ_বিট_এবং | 7 | Bitwise AND |
| অপ_বিট_অথবা | 8 | Bitwise OR |
| অপ_বিট_এক্সঅর | 9 | Bitwise XOR |
| অপ_বিট_না | 10 | Bitwise NOT |
| অপ_বাম_সরান | 11 | Left shift |
| অপ_ডান_সরান | 12 | Right shift |
| অপ_সত্য | 13 | Push true |
| অপ_মিথ্যা | 14 | Push false |
| অপ_সমান | 15 | Equality |
| অপ_অসমান | 16 | Not equal |
| অপ_বড় | 17 | Greater than |
| অপ_বড়_সমান | 18 | Greater or equal |
| অপ_ঋণাত্মক | 19 | Unary minus |
| অপ_না | 20 | Logical NOT |
| অপ_যৌক্তিক_এবং | 21 | Logical AND |
| অপ_যৌক্তিক_অথবা | 22 | Logical OR |
| অপ_লাফ_মিথ্যা | 23 | Jump if falsy |
| অপ_লাফ | 24 | Unconditional jump |
| অপ_নাল | 25 | Push null |
| অপ_বৈশ্বিক_পাও | 26 | Get global var |
| অপ_বৈশ্বিক_সেট | 27 | Set global var |
| অপ_অ্যারে | 28 | Create array |
| অপ_হ্যাশ | 29 | Create hash |
| অপ_সূচক | 30 | Index operation |
| অপ_কল | 31 | Function call |
| অপ_মান_ফেরত | 32 | Return value |
| অপ_ফেরত | 33 | Return |
| অপ_স্থানীয়_পাও | 34 | Get local var |
| অপ_স্থানীয়_সেট | 35 | Set local var |
| অপ_বিল্টইন_পাও | 36 | Get builtin |
| অপ_ক্লোজার | 37 | Create closure |
| অপ_মুক্ত_পাও | 38 | Get free var |
| অপ_বর্তমান_ক্লোজার | 39 | Get current closure |

### প্রধান ফাংশন

#### `নির্দেশ_তৈরি(অপ, অপারেন্ড)`
Bytecode instruction তৈরি করে।

**Parameters:**
- `অপ`: Opcode number
- `অপারেন্ড`: Array of operands

**Returns:** Instruction bytes array

#### `নির্দেশ_থেকে_স্ট্রিং(নির্দেশ)`
Instructions কে readable string এ রূপান্তর করে।

---

## প্রতীক টেবিল মডিউল (`modules/প্রতীক_টেবিল.ভাষা`)

### ডেটা স্ট্রাকচার

#### প্রতীক টেবিল
```bengali
{
    "বাইরের": Parent symbol table,
    "ভাণ্ডার": HashMap of symbols,
    "সংজ্ঞা_সংখ্যা": Number of definitions,
    "মুক্ত_প্রতীক": Array of free variables
}
```

#### প্রতীক
```bengali
{
    "নাম": Variable name,
    "স্কোপ": Scope type,
    "সূচক": Variable index
}
```

### স্কোপ টাইপ
- `বৈশ্বিক_স্কোপ`: Global scope
- `স্থানীয়_স্কোপ`: Local scope
- `বিল্টইন_স্কোপ`: Builtin functions
- `মুক্ত_স্কোপ`: Free variables (closures)
- `ফাংশন_স্কোপ`: Function name (recursion)

### প্রধান ফাংশন

#### `নতুন_প্রতীক_টেবিল()`
নতুন symbol table তৈরি করে।

#### `প্রতীক_সংজ্ঞা(টেবিল, নাম)`
নতুন variable define করে।

#### `প্রতীক_খোঁজো(টেবিল, নাম)`
Variable lookup করে।

**Returns:**
```bengali
{
    "পাওয়া_গেছে": Boolean,
    "প্রতীক": Symbol object,
    "টেবিল": Updated table
}
```

---

## মডিউল লোডার (`modules/মডিউল_লোডার.ভাষা`)

### প্রধান ফাংশন

#### `নতুন_মডিউল_লোডার()`
নতুন module loader তৈরি করে।

#### `মডিউল_লোড(লোডার, পথ)`
Module load করে।

**Parameters:**
- `লোডার`: Module loader object
- `পথ`: Module file path

**Returns:**
```bengali
{
    "সফল": Boolean,
    "মডিউল": Module object,
    "লোডার": Updated loader,
    "ত্রুটি": Error message
}
```

---

## মূল কম্পাইলার ড্রাইভার (`modules/ভাষা_কম্পাইলার.ভাষা`)

### প্রধান ফাংশন

#### `কম্পাইল_ফাইল(ফাইল_পথ)`
File compile করে।

**Parameters:**
- `ফাইল_পথ`: Path to `.ভাষা` file

**Returns:**
```bengali
{
    "সফল": Boolean,
    "বাইটকোড": Bytecode object,
    "টোকেন_সংখ্যা": Number,
    "বিবৃতি_সংখ্যা": Number,
    "নির্দেশ_সংখ্যা": Number,
    "ধ্রুবক_সংখ্যা": Number
}
```

#### `কম্পাইল_সোর্স(সোর্স_কোড)`
Source string compile করে।

#### `বাইটকোড_দেখাও(বাইটকোড)`
Bytecode details print করে।

#### `কম্পাইলার_পরীক্ষা()`
Built-in tests চালায়।

#### `মূল(আর্গুমেন্ট)`
Main entry point.

---

## সম্পূর্ণ উদাহরণ

```bengali
// example.ভাষা

// ১. সাধারণ arithm etic
ধরি x = ৫ + ৩ * ২;
লেখ(x);  // ১১

// ২. ফাংশন
ধরি ফ্যাক্টোরিয়াল = ফাংশন(n) {
    যদি (n <= ১) {
        ফেরত ১;
    }
    ফেরত n * ফ্যাক্টোরিয়াল(n - ১);
};

লেখ(ফ্যাক্টোরিয়াল(৫));  // ১২০

// ৩. Array
ধরি তালিকা = [১, ২, ৩, ৪, ৫];
লেখ(দৈর্ঘ্য(তালিকা));  // ৫

// ৪. Hash
ধরি ব্যক্তি = {
    "নাম": "রহিম",
    "বয়স": ২৫
};
লেখ(ব্যক্তি["নাম"]);  // "রহিম"

// ৫. Loop
পর্যন্ত (ধরি i = ০; i < ৫; i = i + ১) {
    লেখ(i);
}
```

---

## Error Handling

### Parser Errors
Parser errors `ত্রুটি` array এ store হয়:

```bengali
যদি (দৈর্ঘ্য(পার্সার["ত্রুটি"]) > ০) {
    // Handle errors
}
```

### Compiler Errors
Compilation results indicate success:

```bengali
ধরি ফলাফল = কম্পাইল(কম্পাইলার, নোড);
যদি (!ফলাফল["সফল"]) {
    লেখ("ত্রুটি: " + ফলাফল["ত্রুটি"]);
}
```

---

## Performance Considerations

1. **Module Caching**: লোড করা modules automatically cache হয়
2. **Scope Management**: Symbol tables efficiently manage scoping
3. **Bytecode Size**: Compact bytecode representation
4. **Constant Pooling**: Repeated constants শুধু একবার store হয়

---

## Contributing

API তে পরিবর্তনের সময় এই ডকুমেন্ট update করুন।

---

**Last Updated**: 2025  
**Version**: 1.0.0

