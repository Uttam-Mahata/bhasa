# Research Paper: OOP Implementation in Bhasa

## Overview

This directory contains a comprehensive research paper documenting the Object-Oriented Programming (OOP) implementation in the Bhasa programming language.

**Paper Title:** *Object-Oriented Programming in Bhasa: A Bengali-Native Programming Language with Meaningful Semantic Naming*

**File:** `research_paper.tex`

## Contents

The paper includes:

1. **Abstract** - Summary of the OOP implementation and contributions
2. **Introduction** - Context and motivation for Bengali-native OOP
3. **Background** - Overview of Bhasa language architecture
4. **Design Principles** - Meaningful Bengali naming philosophy
5. **Implementation** - Detailed technical implementation across all components:
   - Lexical Analysis (token/token.go)
   - Abstract Syntax Tree (ast/ast.go)
   - Parsing (parser/parser.go)
   - Bytecode Generation (compiler/compiler.go)
   - Virtual Machine Execution (vm/vm.go)
   - Runtime Object Model (object/object.go)
6. **Examples and Verification** - Working code examples
7. **Results and Discussion** - Metrics and evaluation
8. **Future Work** - Planned enhancements
9. **Related Work** - Comparison with other regional languages
10. **Conclusion** - Summary and impact

## Key Contributions

- **16 Meaningful Bengali OOP Keywords** (not transliterations)
- **1,415 Lines of Implementation Code**
- **12 New Bytecode Operations**
- **Complete Compilation Pipeline Extension**
- **Semantic Authenticity** over phonetic transliteration

## Compilation Instructions

### Requirements

You need a LaTeX distribution with XeLaTeX support and Bengali fonts:

**Ubuntu/Debian:**
```bash
sudo apt-get install texlive-full texlive-xetex fonts-noto-core
```

**macOS (via MacTeX):**
```bash
brew install --cask mactex
sudo tlmgr install polyglossia fontspec
```

**Windows (via MiKTeX):**
Download from: https://miktex.org/download

### Compiling to PDF

```bash
cd /home/user/bhasa
xelatex research_paper.tex
xelatex research_paper.tex  # Run twice for references
```

This will generate `research_paper.pdf`.

### Alternative: Online LaTeX Editors

If you don't have LaTeX installed locally, you can compile online:

1. **Overleaf** (https://www.overleaf.com)
   - Create new project
   - Upload `research_paper.tex`
   - Set compiler to XeLaTeX (in Menu > Settings)
   - Compile

2. **Papeeria** (https://papeeria.com)
   - Upload the .tex file
   - Compile with XeLaTeX

### Fonts Note

The paper uses **Noto Sans Bengali** font for Bengali text. If you don't have this font:

- Download from: https://fonts.google.com/noto/specimen/Noto+Sans+Bengali
- Or modify the `\newfontfamily\bengalifont` line to use a different Bengali font

## Paper Structure

### Sections and Page Count (Estimated)

- Abstract: 1 page
- Introduction: 1-2 pages
- Background: 1 page
- Design Principles: 2 pages
- Implementation: 8-10 pages (largest section with code examples)
- Examples and Verification: 2-3 pages
- Results and Discussion: 2 pages
- Future Work: 1 page
- Related Work: 1 page
- Conclusion: 1 page
- References: 1 page

**Total: ~20-25 pages**

## Key Tables and Figures

1. **Table 1:** OOP Keywords with Semantic Mappings
2. **Table 2:** OOP Bytecode Operations
3. **Table 3:** Implementation Size by Component
4. **Table 4:** Keyword Approach Comparison
5. **Figure 1:** Bhasa Compilation Pipeline

## Code Examples in Paper

The paper includes extensive code examples demonstrating:

- Simple class with constructor and methods
- Inheritance with method overriding
- Interface implementation
- Parser implementation snippets
- Compiler implementation snippets
- VM execution logic

All code uses Bengali keywords to demonstrate the natural programming experience.

## Citation

If you use this work, please cite as:

```
Bhasa Development Team (2025).
Object-Oriented Programming in Bhasa: A Bengali-Native Programming Language
with Meaningful Semantic Naming.
Open Source Project, India.
```

## Related Files

- `OOP_FEATURES.md` - Complete OOP feature documentation
- `examples/test_oop_simple.bhasa` - Working OOP infrastructure test
- `examples/test_oop_class.bhasa` - Full class syntax example
- `examples/oop_meaningful_names.bhasa` - Meaningful naming demonstration

## Implementation Metrics

From the paper:

| Component | Lines of Code |
|-----------|--------------|
| Token definitions | 35 |
| AST nodes | 280 |
| Opcode definitions | 50 |
| Runtime objects | 140 |
| Parser functions | 460 |
| Compiler functions | 250 |
| VM execution | 200 |
| **Total** | **1,415** |

## Keywords Implemented

All 16 OOP keywords with meaningful Bengali semantics:

- শ্রেণী (class - category)
- পদ্ধতি (method - procedure)
- নির্মাতা (constructor - creator)
- এই (this - current)
- নতুন (new - fresh)
- প্রসারিত (extends - expanded)
- সার্বজনীন (public - for all people)
- ব্যক্তিগত (private - personal)
- সুরক্ষিত (protected - safeguarded)
- স্থির (static - fixed)
- বিমূর্ত (abstract - conceptual)
- চুক্তি (interface - contract)
- বাস্তবায়ন (implements - implementation)
- উর্ধ্ব (super - upper)
- পুনর্সংজ্ঞা (override - redefinition)
- চূড়ান্ত (final - conclusive)

## License

This research paper documents open-source work on the Bhasa programming language.

## Contact

For questions or contributions, please visit:
https://github.com/Uttam-Mahata/bhasa
