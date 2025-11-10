# ✅ Bhasa Extension - Successfully Installed!

## 🎉 Installation Complete

The Bhasa language extension is now installed in your VS Code!

**Location**: `/home/uttam/bhasa/vscode-bhasa/bhasa/bhasa-0.0.1.vsix`

---

## 🚀 How to Use the Extension

### Method 1: Open Any Bhasa File

Simply open or create a file with `.bhasa` or `.ভাষা` extension:

```bash
# Open an existing example
code /home/uttam/bhasa/examples/hello.bhasa

# Or create a new file
code ~/my-program.bhasa
```

The extension will automatically activate and provide:
- ✅ Syntax highlighting
- ✅ Auto-completion (Ctrl+Space)
- ✅ Hover documentation
- ✅ Auto-closing brackets

### Method 2: Test with the Sample File

```bash
# Open the test file
code /home/uttam/bhasa/vscode-bhasa/bhasa/test.bhasa
```

---

## 🧪 Testing the Features

Once you open a `.bhasa` file:

1. **Syntax Highlighting**: Keywords like `ধরি`, `ফাংশন`, `যদি` should be colored

2. **Auto-Completion**: 
   - Type `ধ` and press `Ctrl+Space`
   - You should see suggestions like `ধরি`

3. **Hover Information**:
   - Hover your mouse over `ফাংশন`
   - You should see: "Function declaration keyword"

4. **Auto-Closing**:
   - Type `{` and it will automatically add `}`
   - Type `"` and it will automatically add the closing `"`

---

## 📝 Create Your First Bhasa Program

```bash
# Create a new file
cd ~
cat > hello.bhasa << 'EOF'
// নমস্কার বিশ্ব (Hello World)
লেখ("নমস্কার বিশ্ব!");

// Variable declaration
ধরি নাম = "উত্তম";
লেখ(নাম);

// Function example
ধরি যোগ = ফাংশন(a, b) {
    ফেরত a + b;
};

লেখ(যোগ(১০, ২০));
EOF

# Open in VS Code
code hello.bhasa
```

---

## 🔄 Updating the Extension

If you make changes to the extension:

```bash
cd /home/uttam/bhasa/vscode-bhasa/bhasa

# 1. Make your changes to src/extension.ts or syntaxes/bhasa.tmLanguage.json

# 2. Recompile
npm run compile

# 3. Repackage
vsce package

# 4. Reinstall
code --install-extension bhasa-0.0.1.vsix --force
```

---

## 🗑️ Uninstalling the Extension

If you need to remove the extension:

```bash
code --uninstall-extension Uttam-Mahata.bhasa
```

---

## 🔍 Verify Installation

Check if the extension is installed:

```bash
code --list-extensions | grep bhasa
```

You should see: `uttam-mahata.bhasa`

---

## 🎨 What's Working Now

✅ **File Types**: `.bhasa` and `.ভাষা` files
✅ **Keywords**: 40+ Bengali and English keywords
✅ **Functions**: লেখ, দৈর্ঘ্য, ধাক্কা, etc.
✅ **Operators**: +, -, *, /, ==, !=, &&, ||, etc.
✅ **Comments**: // and /* */
✅ **Numbers**: Both 0-9 and ০-৯
✅ **Strings**: Single and double quotes

---

## 🐛 Troubleshooting

**Extension not working?**
1. Reload VS Code: `Ctrl+Shift+P` → "Reload Window"
2. Check file extension is `.bhasa` or `.ভাষা`
3. Check the extension is enabled: Extensions panel (Ctrl+Shift+X)

**No syntax highlighting?**
1. Close and reopen the file
2. Check language mode (bottom right): should say "Bhasa"
3. Manually set language: `Ctrl+Shift+P` → "Change Language Mode" → "Bhasa"

---

## 📦 Extension Files

The packaged extension is here:
```
/home/uttam/bhasa/vscode-bhasa/bhasa/bhasa-0.0.1.vsix
```

You can share this `.vsix` file with others to install the extension!

---

## 🎯 Quick Commands

```bash
# Open VS Code with a Bhasa file
code /home/uttam/bhasa/examples/hello.bhasa

# List all Bhasa examples
ls /home/uttam/bhasa/examples/*.bhasa

# Run a Bhasa program
/home/uttam/bhasa/bin/bhasa-linux-amd64 hello.bhasa
```

---

## ✨ Success!

Your Bhasa extension is now active! Open any `.bhasa` file in VS Code to see it in action.

**Test it now**: `code /home/uttam/bhasa/examples/hello.bhasa`

🎊 নমস্কার! Happy coding in Bhasa! 🇧🇩
