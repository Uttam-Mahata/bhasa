# Bhasa VS Code Extension - Development Guide

## Project Structure

```
vscode-bhasa/bhasa/
├── src/
│   └── extension.ts          # Main extension code
├── syntaxes/
│   └── bhasa.tmLanguage.json # Syntax highlighting definitions
├── language-configuration.json # Language configuration
├── package.json              # Extension manifest
├── tsconfig.json            # TypeScript configuration
└── test.bhasa               # Test file for verification
```

## Testing the Extension

### Method 1: Run in Development Mode

1. Open the extension folder in VS Code:
   ```bash
   cd /home/uttam/bhasa/vscode-bhasa/bhasa
   code .
   ```

2. Press `F5` to start debugging
   - This will open a new VS Code window with the extension loaded
   - The extension will be active for `.bhasa` and `.ভাষা` files

3. Test the features:
   - Open `test.bhasa` in the new window
   - Check syntax highlighting
   - Try auto-completion (Ctrl+Space)
   - Hover over keywords to see descriptions

### Method 2: Install Locally

1. Package the extension:
   ```bash
   npm install -g @vscode/vsce
   cd /home/uttam/bhasa/vscode-bhasa/bhasa
   vsce package
   ```

2. Install the `.vsix` file:
   ```bash
   code --install-extension bhasa-0.0.1.vsix
   ```

## Building the Extension

### Compile TypeScript
```bash
npm run compile
```

### Watch Mode (for development)
```bash
npm run watch
```

### Lint Code
```bash
npm run lint
```

## Features Implemented

✅ **Syntax Highlighting**
- Bengali keywords (ধরি, ফাংশন, যদি, etc.)
- English keywords (let, function, if, etc.)
- Comments (// and /* */)
- Strings (single and double quotes)
- Numbers (including Bengali numerals ০-৯)
- Operators (arithmetic, logical, comparison, bitwise)

✅ **IntelliSense**
- Auto-completion for all Bhasa keywords
- Context-aware suggestions
- Detail information for each keyword

✅ **Hover Support**
- Hover over keywords to see descriptions
- Works for both Bengali and English keywords

✅ **Language Configuration**
- Auto-closing brackets, parentheses, quotes
- Comment toggling (Ctrl+/)
- Code folding support

## Publishing the Extension

### Prerequisites
```bash
npm install -g @vscode/vsce
```

### Create Personal Access Token
1. Go to https://dev.azure.com/
2. Create a new Personal Access Token with "Marketplace" scope
3. Save the token securely

### Publish to VS Code Marketplace
```bash
# Login (first time only)
vsce login <publisher-name>

# Package the extension
vsce package

# Publish
vsce publish
```

### Alternative: Publish to GitHub
You can also distribute the `.vsix` file via GitHub releases.

## Updating the Extension

1. Make your changes to the code
2. Update version in `package.json`
3. Update `CHANGELOG.md`
4. Compile: `npm run compile`
5. Test with F5
6. Package: `vsce package`
7. Publish: `vsce publish`

## Adding More Features

### Add New Keywords
Edit `syntaxes/bhasa.tmLanguage.json`:
```json
{
  "name": "keyword.control.bhasa",
  "match": "\\b(new_keyword|নতুন_কীওয়ার্ড)\\b"
}
```

### Add Snippets
Create `snippets/bhasa.json`:
```json
{
  "Function": {
    "prefix": "func",
    "body": [
      "ধরি ${1:name} = ফাংশন($2) {",
      "    $0",
      "};"
    ],
    "description": "Function declaration"
  }
}
```

### Add Commands
Edit `src/extension.ts` and `package.json` to register new commands.

## Troubleshooting

### Extension Not Loading
- Check `package.json` activationEvents
- Ensure file extensions are registered correctly
- Run `npm run compile` to rebuild

### Syntax Highlighting Not Working
- Verify `syntaxes/bhasa.tmLanguage.json` is valid JSON
- Check scopeName matches in package.json
- Reload VS Code window (Ctrl+Shift+P > "Reload Window")

### IntelliSense Not Working
- Ensure completion provider is registered in extension.ts
- Check language ID matches ("bhasa")
- Verify the extension is activated

## Resources

- [VS Code Extension API](https://code.visualstudio.com/api)
- [TextMate Grammars](https://macromates.com/manual/en/language_grammars)
- [Publishing Extensions](https://code.visualstudio.com/api/working-with-extensions/publishing-extension)
- [Extension Guidelines](https://code.visualstudio.com/api/references/extension-guidelines)

## Next Steps

Potential enhancements:
- [ ] Add code snippets for common patterns
- [ ] Implement bracket matching
- [ ] Add definition provider (Go to Definition)
- [ ] Add reference provider (Find All References)
- [ ] Add formatting provider
- [ ] Add debugging support
- [ ] Add integration with Bhasa compiler
- [ ] Add error diagnostics
- [ ] Add run/compile commands
- [ ] Add Bengali numeral support in IntelliSense

## Contact

For issues or contributions, visit: https://github.com/Uttam-Mahata/bhasa
