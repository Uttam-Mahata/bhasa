# Encoding Support in Bhasa (ভাষা)

## UTF-8 Support

Bhasa has **full UTF-8 support** for Bengali script. All components handle Unicode correctly:

### Verified Components:

✅ **Lexer** - Uses `rune` type for multi-byte character support
✅ **Parser** - Handles Bengali identifiers and keywords
✅ **Strings** - Proper Unicode string handling
✅ **Output** - Go's `fmt` package natively supports UTF-8
✅ **File I/O** - Reads and writes UTF-8 files correctly

## Terminal Requirements

To display Bengali text properly, ensure your terminal:

1. **Supports UTF-8 encoding**
2. **Has Bengali fonts installed** (e.g., Noto Sans Bengali)
3. **Locale is set correctly**

### Check Your Setup:

```bash
# Check current locale
locale

# Should show UTF-8, like:
# LANG=en_US.UTF-8
# or
# LANG=bn_BD.UTF-8
```

### Set Locale (if needed):

```bash
# For English with UTF-8
export LANG=en_US.UTF-8

# For Bengali locale
export LANG=bn_BD.UTF-8
```

## How It Works

### 1. Source Files
- Bhasa source files (`.bhasa` or `.ভাষা`) are saved as **UTF-8**
- Comments, strings, and identifiers can all use Bengali script

### 2. Lexical Analysis
```go
// lexer/lexer.go uses rune (int32) not byte (uint8)
type Lexer struct {
    input  []rune  // Handles multi-byte UTF-8 characters
    // ...
}
```

### 3. String Length
Bengali characters are counted correctly:
```go
// "হ্যালো" has 5 Bengali characters (including vowel sign)
// len([]rune("হ্যালো")) = 5  ✅ Correct
// len("হ্যালো")           = 15 ❌ Byte count (wrong for display)
```

### 4. Output
Go's standard library automatically handles UTF-8 output:
```go
fmt.Println("নমস্কার বিশ্ব!")  // Works perfectly
```

## Testing UTF-8 Support

Run this test to verify Bengali support:

```bash
# Create test file
cat > test_bengali.ভাষা << 'EOF'
লেখ("পরীক্ষা ১২৩");
ধরি নাম = "রহিম";
লেখ(নাম);
ধরি সংখ্যা = ৫০;
লেখ(সংখ্যা);
EOF

# Run it
./bhasa test_bengali.ভাষা

# Expected output:
# পরীক্ষা ১২৩
# রহিম
# 50
```

## Common Issues & Solutions

### Issue: Boxes (□) instead of Bengali characters
**Solution**: Install Bengali fonts
```bash
# Ubuntu/Debian
sudo apt install fonts-noto-bengali

# Fedora
sudo dnf install google-noto-sans-bengali-fonts
```

### Issue: Question marks (???) instead of Bengali
**Solution**: Set UTF-8 locale
```bash
export LANG=en_US.UTF-8
```

### Issue: Works in terminal but not in editor
**Solution**: Configure your editor for UTF-8
- VS Code: Already UTF-8 by default
- Vim: Add `set encoding=utf-8` to `.vimrc`
- Emacs: Add `(set-language-environment "UTF-8")` to init file

## File Encoding

All Bhasa files should be saved as **UTF-8**:

```bash
# Check file encoding
file -bi examples/hello.ভাষা
# Output: text/plain; charset=utf-8

# Convert file to UTF-8 (if needed)
iconv -f ISO-8859-1 -t UTF-8 input.txt -o output.txt
```

## Summary

**Your Bhasa interpreter fully supports Bengali (UTF-8) encoding!**

- ✅ Source code can use Bengali keywords
- ✅ Variable names can be in Bengali  
- ✅ String literals support Bengali text
- ✅ Comments can be in Bengali
- ✅ Output displays Bengali correctly
- ✅ File extensions can be Bengali (`.ভাষা`)

**No encoding configuration needed** - it works out of the box! 🎉

