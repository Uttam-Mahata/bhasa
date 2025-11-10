# Bhasa Language Support for VS Code

Language support extension for the Bhasa programming language - a Bengali-focused programming language that supports both Bengali and English syntax.

## Features

This extension provides comprehensive language support for Bhasa programming language:

* **Syntax Highlighting** - Full syntax highlighting for both `.bhasa` and `.ভাষা` file extensions
* **Code Completion** - IntelliSense support for Bengali and English keywords
* **Hover Information** - Hover over keywords to see their descriptions
* **Auto-closing Pairs** - Automatic closing of brackets, quotes, and parentheses
* **Comment Support** - Line (`//`) and block (`/* */`) comment support
* **Code Folding** - Region-based code folding support

### Supported Keywords

#### Bengali Keywords
- `ধরি` (let) - Variable declaration
- `ফাংশন` (function) - Function declaration
- `যদি` (if) - Conditional statement
- `নাহলে` (else) - Else clause
- `ফেরত` (return) - Return statement
- `যতক্ষণ` (while) - While loop
- `পর্যন্ত` (for) - For loop
- `বিরতি` (break) - Break statement
- `চালিয়ে_যাও` (continue) - Continue statement
- `সত্য` (true), `মিথ্যা` (false) - Boolean values
- `লেখ` (print) - Print function

#### English Keywords
The extension also supports equivalent English keywords for all Bengali keywords.

## Requirements

To run Bhasa programs, you need the Bhasa interpreter installed on your system. Visit the [Bhasa repository](https://github.com/Uttam-Mahata/bhasa) for installation instructions.

## Usage

1. Create a new file with `.bhasa` or `.ভাষা` extension
2. Start writing your Bhasa code with syntax highlighting
3. Use IntelliSense (Ctrl+Space) for code completion
4. Hover over keywords to see their descriptions

### Example

```bhasa
// নমস্কার বিশ্ব (Hello World)
লেখ("নমস্কার বিশ্ব!");

// Function example
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};

লেখ(যোগ(১০, ২০));
```

## Extension Settings

This extension contributes the following settings:

* Automatic activation for `.bhasa` and `.ভাষা` files
* Syntax highlighting configuration
* Language configuration for auto-closing pairs and comments

## Known Issues

This is an initial release. Please report any issues on the [GitHub repository](https://github.com/Uttam-Mahata/bhasa).

## Release Notes

### 0.0.1

Initial release of Bhasa Language Support:
- Syntax highlighting for Bhasa language
- IntelliSense support for keywords
- Auto-closing pairs and bracket matching
- Hover information for keywords
- Support for both `.bhasa` and `.ভাষা` file extensions

## Contributing

Contributions are welcome! Please visit the [Bhasa repository](https://github.com/Uttam-Mahata/bhasa) to contribute.

## License

This extension follows the same license as the Bhasa programming language.

**Enjoy coding in Bhasa! নমস্কার! 🇧🇩**
