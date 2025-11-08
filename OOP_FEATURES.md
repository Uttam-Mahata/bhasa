# ভাষা (Bhasa) - Object-Oriented Programming Features

## Overview

Bhasa now supports comprehensive Object-Oriented Programming (OOP) features with meaningful Bengali naming conventions. All OOP keywords use proper Bengali words rather than transliterations.

---

## Bengali OOP Keywords

| English | Bengali | Meaning |
|---------|---------|---------|
| class | **শ্রেণী** | Category/Class |
| method | **পদ্ধতি** | Procedure/Method |
| constructor | **নির্মাতা** | Creator/Builder |
| this | **এই** | This |
| new | **নতুন** | New |
| extends | **প্রসারিত** | Extended/Expanded |
| public | **সার্বজনীন** | For all people |
| private | **ব্যক্তিগত** | Personal/Private |
| protected | **সুরক্ষিত** | Protected/Safeguarded |
| static | **স্থির** | Fixed/Static |
| abstract | **বিমূর্ত** | Conceptual/Abstract |
| interface | **চুক্তি** | Contract/Agreement |
| implements | **বাস্তবায়ন** | Implementation |
| super | **উর্ধ্ব** | Upper/Higher |
| override | **পুনর্সংজ্ঞা** | Redefine |
| final | **চূড়ান্ত** | Ultimate/Final |

---

## 1. Basic Class Definition (শ্রেণী)

### Simple Class with Fields and Constructor

```bengali
শ্রেণী ব্যক্তি {
    // Fields (সার্বজনীন = public)
    সার্বজনীন নাম: লেখা;
    সার্বজনীন বয়স: পূর্ণসংখ্যা;
    ব্যক্তিগত আইডি: পূর্ণসংখ্যা;

    // Constructor (নির্মাতা)
    সার্বজনীন নির্মাতা(নাম: লেখা, বয়স: পূর্ণসংখ্যা, আইডি: পূর্ণসংখ্যা) {
        এই.নাম = নাম;
        এই.বয়স = বয়স;
        এই.আইডি = আইডি;
    }

    // Public Method (সার্বজনীন পদ্ধতি)
    সার্বজনীন পদ্ধতি পরিচয়_দাও(): লেখা {
        ফেরত "নাম: " + এই.নাম + ", বয়স: " + এই.বয়স;
    }

    // Private Method (ব্যক্তিগত পদ্ধতি)
    ব্যক্তিগত পদ্ধতি আইডি_পাও(): পূর্ণসংখ্যা {
        ফেরত এই.আইডি;
    }
}

// Creating an instance (নতুন)
ধরি ব্যক্তি১ = নতুন ব্যক্তি("রহিম", ৩০, ১০০১);
লেখ(ব্যক্তি১.পরিচয়_দাও());
```

---

## 2. Inheritance (প্রসারিত)

### Class Inheritance

```bengali
// Parent Class
শ্রেণী প্রাণী {
    সার্বজনীন নাম: লেখা;

    সার্বজনীন নির্মাতা(নাম: লেখা) {
        এই.নাম = নাম;
    }

    সার্বজনীন পদ্ধতি শব্দ_করো(): লেখা {
        ফেরত "প্রাণী শব্দ করছে";
    }
}

// Child Class (প্রসারিত = extends)
শ্রেণী কুকুর প্রসারিত প্রাণী {
    সার্বজনীন জাত: লেখা;

    সার্বজনীন নির্মাতা(নাম: লেখা, জাত: লেখা) {
        উর্ধ্ব.নির্মাতা(নাম);  // Call parent constructor
        এই.জাত = জাত;
    }

    // Override parent method (পুনর্সংজ্ঞা)
    পুনর্সংজ্ঞা সার্বজনীন পদ্ধতি শব্দ_করো(): লেখা {
        ফেরত "ঘেউ ঘেউ!";
    }

    সার্বজনীন পদ্ধতি তথ্য_দাও(): লেখা {
        ফেরত এই.নাম + " একটি " + এই.জাত + " কুকুর";
    }
}

// Usage
ধরি আমার_কুকুর = নতুন কুকুর("টমি", "জার্মান শেফার্ড");
লেখ(আমার_কুকুর.শব্দ_করো());    // Output: ঘেউ ঘেউ!
লেখ(আমার_কুকুর.তথ্য_দাও());     // Output: টমি একটি জার্মান শেফার্ড কুকুর
```

---

## 3. Interfaces (চুক্তি)

### Interface Definition and Implementation

```bengali
// Interface (চুক্তি = contract)
চুক্তি কথাবলার_যোগ্যতা {
    পদ্ধতি বলো(বার্তা: লেখা): লেখা;
    পদ্ধতি শোনো(): লেখা;
}

// Class implementing interface (বাস্তবায়ন = implements)
শ্রেণী মানুষ বাস্তবায়ন কথাবলার_যোগ্যতা {
    সার্বজনীন নাম: লেখা;

    সার্বজনীন নির্মাতা(নাম: লেখা) {
        এই.নাম = নাম;
    }

    সার্বজনীন পদ্ধতি বলো(বার্তা: লেখা): লেখা {
        ফেরত এই.নাম + " বলছে: " + বার্তা;
    }

    সার্বজনীন পদ্ধতি শোনো(): লেখা {
        ফেরত এই.নাম + " শুনছে...";
    }
}

// Usage
ধরি ব্যক্তি = নতুন মানুষ("করিম");
লেখ(ব্যক্তি.বলো("হ্যালো!"));   // Output: করিম বলছে: হ্যালো!
```

---

## 4. Static Members (স্থির)

### Static Fields and Methods

```bengali
শ্রেণী গণক {
    // Static field (স্থির)
    স্থির সার্বজনীন গণনা: পূর্ণসংখ্যা = ০;

    সার্বজনীন নির্মাতা() {
        গণক.গণনা = গণক.গণনা + ১;
    }

    // Static method (স্থির পদ্ধতি)
    স্থির সার্বজনীন পদ্ধতি মোট_সংখ্যা(): পূর্ণসংখ্যা {
        ফেরত গণক.গণনা;
    }
}

// Usage
ধরি গ১ = নতুন গণক();
ধরি গ২ = নতুন গণক();
ধরি গ৩ = নতুন গণক();

লেখ(গণক.মোট_সংখ্যা());  // Output: ৩
```

---

## 5. Abstract Classes (বিমূর্ত শ্রেণী)

### Abstract Class with Abstract Methods

```bengali
// Abstract class (বিমূর্ত)
বিমূর্ত শ্রেণী আকৃতি {
    সুরক্ষিত নাম: লেখা;

    সার্বজনীন নির্মাতা(নাম: লেখা) {
        এই.নাম = নাম;
    }

    // Abstract method (বিমূর্ত পদ্ধতি) - no implementation
    বিমূর্ত সার্বজনীন পদ্ধতি ক্ষেত্রফল(): দশমিক;

    সার্বজনীন পদ্ধতি বর্ণনা(): লেখা {
        ফেরত "এটি একটি " + এই.নাম;
    }
}

// Concrete class
শ্রেণী বৃত্ত প্রসারিত আকৃতি {
    ব্যক্তিগত ব্যাসার্ধ: দশমিক;

    সার্বজনীন নির্মাতা(ব্যাসার্ধ: দশমিক) {
        উর্ধ্ব.নির্মাতা("বৃত্ত");
        এই.ব্যাসার্ধ = ব্যাসার্ধ;
    }

    // Implement abstract method
    পুনর্সংজ্ঞা সার্বজনীন পদ্ধতি ক্ষেত্রফল(): দশমিক {
        ফেরত ৩.১৪১৫৯ * এই.ব্যাসার্ধ * এই.ব্যাসার্ধ;
    }
}

// Usage
ধরি আমার_বৃত্ত = নতুন বৃত্ত(৫.০);
লেখ(আমার_বৃত্ত.বর্ণনা());     // Output: এটি একটি বৃত্ত
লেখ(আমার_বৃত্ত.ক্ষেত্রফল());   // Output: 78.53975
```

---

## 6. Final Classes and Methods (চূড়ান্ত)

### Final Classes (Cannot be extended)

```bengali
// Final class (চূড়ান্ত) - cannot be extended
চূড়ান্ত শ্রেণী স্ট্রিং_সাহায্যকারী {
    স্থির সার্বজনীন পদ্ধতি বড়_হাতের_অক্ষরে(টেক্সট: লেখা): লেখা {
        ফেরত উপরে(টেক্সট);
    }

    স্থির সার্বজনীন পদ্ধতি ছোট_হাতের_অক্ষরে(টেক্সট: লেখা): লেখা {
        ফেরত নিচে(টেক্সট);
    }
}

// Final method (চূড়ান্ত পদ্ধতি) - cannot be overridden
শ্রেণী বেস {
    চূড়ান্ত সার্বজনীন পদ্ধতি গুরুত্বপূর্ণ_পদ্ধতি(): লেখা {
        ফেরত "এটি পরিবর্তন করা যাবে না";
    }
}
```

---

## 7. Complex Example: Bank Account System

```bengali
// Interface for transactions
চুক্তি লেনদেন_যোগ্য {
    পদ্ধতি জমা_করো(পরিমাণ: দশমিক): বুলিয়ান;
    পদ্ধতি তোলো(পরিমাণ: দশমিক): বুলিয়ান;
    পদ্ধতি ব্যালেন্স_দেখাও(): দশমিক;
}

// Abstract base class
বিমূর্ত শ্রেণী ব্যাংক_অ্যাকাউন্ট বাস্তবায়ন লেনদেন_যোগ্য {
    সুরক্ষিত অ্যাকাউন্ট_নম্বর: লেখা;
    সুরক্ষিত মালিক: লেখা;
    সুরক্ষিত ব্যালেন্স: দশমিক;
    স্থির ব্যক্তিগত মোট_অ্যাকাউন্ট: পূর্ণসংখ্যা = ০;

    সার্বজনীন নির্মাতা(অ্যাকাউন্ট_নম্বর: লেখা, মালিক: লেখা) {
        এই.অ্যাকাউন্ট_নম্বর = অ্যাকাউন্ট_নম্বর;
        এই.মালিক = মালিক;
        এই.ব্যালেন্স = ০.০;
        ব্যাংক_অ্যাকাউন্ট.মোট_অ্যাকাউন্ট = ব্যাংক_অ্যাকাউন্ট.মোট_অ্যাকাউন্ট + ১;
    }

    সার্বজনীন পদ্ধতি জমা_করো(পরিমাণ: দশমিক): বুলিয়ান {
        যদি (পরিমাণ > ০.০) {
            এই.ব্যালেন্স = এই.ব্যালেন্স + পরিমাণ;
            ফেরত সত্য;
        }
        ফেরত মিথ্যা;
    }

    বিমূর্ত সার্বজনীন পদ্ধতি তোলো(পরিমাণ: দশমিক): বুলিয়ান;

    সার্বজনীন পদ্ধতি ব্যালেন্স_দেখাও(): দশমিক {
        ফেরত এই.ব্যালেন্স;
    }

    সার্বজনীন পদ্ধতি তথ্য(): লেখা {
        ফেরত "অ্যাকাউন্ট: " + এই.অ্যাকাউন্ট_নম্বর + ", মালিক: " + এই.মালিক;
    }
}

// Savings Account (সঞ্চয় অ্যাকাউন্ট)
শ্রেণী সঞ্চয়_অ্যাকাউন্ট প্রসারিত ব্যাংক_অ্যাকাউন্ট {
    ব্যক্তিগত সুদের_হার: দশমিক;
    ব্যক্তিগত ন্যূনতম_ব্যালেন্স: দশমিক = ১০০০.০;

    সার্বজনীন নির্মাতা(অ্যাকাউন্ট_নম্বর: লেখা, মালিক: লেখা, সুদের_হার: দশমিক) {
        উর্ধ্ব.নির্মাতা(অ্যাকাউন্ট_নম্বর, মালিক);
        এই.সুদের_হার = সুদের_হার;
    }

    পুনর্সংজ্ঞা সার্বজনীন পদ্ধতি তোলো(পরিমাণ: দশমিক): বুলিয়ান {
        যদি (পরিমাণ > ০.০ && এই.ব্যালেন্স - পরিমাণ >= এই.ন্যূনতম_ব্যালেন্স) {
            এই.ব্যালেন্স = এই.ব্যালেন্স - পরিমাণ;
            ফেরত সত্য;
        }
        ফেরত মিথ্যা;
    }

    সার্বজনীন পদ্ধতি সুদ_যোগ_করো(): লেখা {
        ধরি সুদ = এই.ব্যালেন্স * এই.সুদের_হার;
        এই.ব্যালেন্স = এই.ব্যালেন্স + সুদ;
        ফেরত "সুদ যোগ হয়েছে: " + সুদ;
    }
}

// Current Account (চলতি অ্যাকাউন্ট)
শ্রেণী চলতি_অ্যাকাউন্ট প্রসারিত ব্যাংক_অ্যাকাউন্ট {
    ব্যক্তিগত ওভারড্রাফট_সীমা: দশমিক;

    সার্বজনীন নির্মাতা(অ্যাকাউন্ট_নম্বর: লেখা, মালিক: লেখা, ওভারড্রাফট_সীমা: দশমিক) {
        উর্ধ্ব.নির্মাতা(অ্যাকাউন্ট_নম্বর, মালিক);
        এই.ওভারড্রাফট_সীমা = ওভারড্রাফট_সীমা;
    }

    পুনর্সংজ্ঞা সার্বজনীন পদ্ধতি তোলো(পরিমাণ: দশমিক): বুলিয়ান {
        যদি (পরিমাণ > ০.০ && পরিমাণ <= এই.ব্যালেন্স + এই.ওভারড্রাফট_সীমা) {
            এই.ব্যালেন্স = এই.ব্যালেন্স - পরিমাণ;
            ফেরত সত্য;
        }
        ফেরত মিথ্যা;
    }
}

// Usage Example
ধরি সঞ্চয় = নতুন সঞ্চয়_অ্যাকাউন্ট("SAV001", "রহিম আহমেদ", ০.০৫);
সঞ্চয়.জমা_করো(৫০০০.০);
লেখ(সঞ্চয়.তথ্য());
লেখ("ব্যালেন্স: " + সঞ্চয়.ব্যালেন্স_দেখাও());

সঞ্চয়.তোলো(২০০০.০);
লেখ("টাকা তোলার পর ব্যালেন্স: " + সঞ্চয়.ব্যালেন্স_দেখাও());

লেখ(সঞ্চয়.সুদ_যোগ_করো());
লেখ("সুদ যোগের পর ব্যালেন্স: " + সঞ্চয়.ব্যালেন্স_দেখাও());
```

---

## 8. Access Modifiers Summary

| Modifier | Bengali | Description |
|----------|---------|-------------|
| Public | **সার্বজনীন** | Accessible from anywhere |
| Private | **ব্যক্তিগত** | Accessible only within the class |
| Protected | **সুরক্ষিত** | Accessible within class and subclasses |

---

## 9. Key Concepts

### Encapsulation (এনক্যাপসুলেশন)
Use access modifiers to hide internal implementation:
- **সার্বজনীন** (public) for public API
- **ব্যক্তিগত** (private) for internal details
- **সুরক্ষিত** (protected) for subclass access

### Inheritance (উত্তরাধিকার)
- Use **প্রসারিত** to extend a parent class
- Use **উর্ধ্ব** to access parent class members
- Use **পুনর্সংজ্ঞা** to override methods

### Polymorphism (বহুরূপতা)
- Implement **চুক্তি** (interfaces) with **বাস্তবায়ন**
- Override methods in subclasses
- Use abstract classes (**বিমূর্ত শ্রেণী**) for common behavior

### Abstraction (বিমূর্তকরণ)
- Use **বিমূর্ত** keyword for abstract classes and methods
- Define contracts with **চুক্তি** (interfaces)

---

## 10. Best Practices

1. **Use meaningful Bengali names**: Choose proper Bengali words that convey meaning, not just transliterations
2. **Follow naming conventions**:
   - Classes: Noun (ব্যক্তি, প্রাণী, গাড়ি)
   - Methods: Verb (গণনা_করো, দেখাও, বলো)
   - Fields: Noun (নাম, বয়স, ঠিকানা)

3. **Access modifiers**:
   - Default to **ব্যক্তিগত** (private) for fields
   - Use **সার্বজনীন** (public) for API methods
   - Use **সুরক্ষিত** (protected) for inheritance

4. **Prefer composition over inheritance** when appropriate
5. **Use interfaces** (**চুক্তি**) to define contracts
6. **Make immutable classes** **চূড়ান্ত** (final)

---

## Architecture Notes

The OOP implementation in Bhasa includes:

- **Full inheritance hierarchy** with method lookup
- **Interface checking** at runtime
- **Access control** enforcement
- **Static members** support
- **Method overriding** with validation
- **Constructor chaining** with **উর্ধ্ব**

All OOP features are compiled to bytecode and executed on the stack-based VM for optimal performance.
