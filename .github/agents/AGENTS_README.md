# Bhasa Custom Copilot Agents

This directory contains specialized GitHub Copilot agents for different aspects of Bhasa programming language development. Each agent has tailored expertise for specific tasks.

## Available Agents

### 🔧 compiler-expert
**Specialization**: Bytecode compilation, VM execution, symbol tables

Use this agent for:
- Adding new bytecode opcodes
- Debugging VM execution issues
- Implementing control flow constructs
- Closure and free variable handling
- Symbol table and scoping problems
- Jump patching for loops and conditionals

**Key files**: `compiler/compiler.go`, `vm/vm.go`, `code/code.go`, `compiler/symbol_table.go`

### 🐹 go-expert
**Specialization**: Go idioms, UTF-8/Unicode handling, testing

Use this agent for:
- Refactoring Go code for better idioms
- Bengali text processing with runes (not bytes)
- Performance optimization and profiling
- Writing table-driven tests
- Interface design and polymorphism
- Build system and cross-compilation

**Key files**: All `.go` files, `Makefile`, test files

### 🎨 vscode-extension-expert
**Specialization**: VS Code extension development, syntax highlighting

Use this agent for:
- Creating/updating syntax highlighting (TextMate grammars)
- Implementing IntelliSense features (autocomplete, hover, definitions)
- Setting up debugging support
- Creating code snippets for Bengali keywords
- Publishing extension to VS Code Marketplace
- Language Server Protocol (LSP) implementation

**Key files**: `vscode-bhasa/` directory, extension manifest, syntaxes

### 🇧🇩 bengali-language-expert
**Specialization**: Bengali script, Unicode, cultural terminology

Use this agent for:
- Choosing appropriate Bengali keywords for new features
- Translating error messages naturally
- Resolving naming conflicts (like `যোগ` ambiguity)
- Handling Bengali script in lexer/parser (vowel signs, conjuncts)
- Unicode normalization and text processing
- Validating cultural appropriateness of terminology

**Key files**: `token/token.go`, `lexer/lexer.go`, `errors/bengali_errors.go`, documentation

### 🎯 language-designer
**Specialization**: Feature design, syntax, semantics, type systems

Use this agent for:
- Designing new language features (syntax and semantics)
- Planning standard library additions
- Type system design decisions
- Error handling strategy (Result types vs try/catch)
- Backward compatibility planning
- Performance vs expressiveness trade-offs

**Key files**: All language design docs, feature proposals, AST definitions

## How to Use These Agents

### On GitHub.com

1. **In the Agents Panel**:
   - Navigate to https://github.com/copilot/agents
   - Select the Bhasa repository
   - Choose an agent from the dropdown menu
   - Submit your prompt

2. **When Creating PRs**:
   - Assign Copilot coding agent to an issue
   - Select your custom agent from the dropdown
   - The agent will use its specialized knowledge

### In GitHub Copilot CLI

```bash
# Use a specific agent
gh copilot /agent compiler-expert "Add OpAsync opcode for async/await"

# Reference agent in prompt
gh copilot "Using @bengali-language-expert, suggest a keyword for 'async'"
```

### In VS Code (as Chat Modes)

Custom agents appear as chat modes in the VS Code Chat view:
1. Open Chat view in VS Code
2. Click the mode dropdown
3. Select your custom agent (e.g., `compiler-expert`)
4. Chat with specialized context

## Example Use Cases

### Adding a New Language Feature

```
Prompt to @language-designer:
"Design syntax for pattern matching in Bhasa with Bengali keywords"

Then follow up with @bengali-language-expert:
"What's a natural Bengali keyword for 'match' in pattern matching?"

Then implement with @compiler-expert:
"Add OpMatch bytecode instruction for pattern matching"
```

### Debugging Compilation

```
Prompt to @compiler-expert:
"Why does struct instantiation crash at vm.go:419 with index out of range?"
```

### Creating VS Code Extension

```
Prompt to @vscode-extension-expert:
"Create TextMate grammar for Bengali keywords including ধরি, ফাংশন, যদি"

Then ask @bengali-language-expert:
"What Unicode ranges should I include for Bengali variable names?"
```

### Refactoring Go Code

```
Prompt to @go-expert:
"Refactor lexer.go to use rune slicing instead of byte slicing for Bengali text"
```

## Agent Tool Access

Each agent has access to specific tools defined in their YAML frontmatter:

| Agent | read | edit | search | grep | create | run |
|-------|------|------|--------|------|--------|-----|
| compiler-expert | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ |
| go-expert | ✅ | ✅ | ✅ | ✅ | ❌ | ✅ |
| vscode-extension-expert | ✅ | ✅ | ✅ | ❌ | ✅ | ❌ |
| bengali-language-expert | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ |
| language-designer | ✅ | ✅ | ✅ | ❌ | ✅ | ❌ |

## Best Practices

1. **Choose the Right Agent**: Match the agent's expertise to your task
2. **Combine Agents**: Use multiple agents for complex tasks (design → implementation)
3. **Be Specific**: Reference specific files, line numbers, or error messages
4. **Provide Context**: Mention related features or constraints
5. **Iterate**: Start with design agents, then move to implementation agents

## Contributing New Agents

To add a new specialized agent:

1. Create a new file in `.github/agents/` (e.g., `docs-expert.md`)
2. Add YAML frontmatter with name, description, and tools:
   ```yaml
   ---
   name: my-agent
   description: Brief description of specialization
   tools: ["read", "edit", "search"]
   ---
   ```
3. Write the agent's expertise and guidance in Markdown
4. Commit to the default branch
5. Refresh https://github.com/copilot/agents

### Agent Template

```markdown
---
name: my-agent
description: Brief description of specialization
tools: ["read", "edit", "search"]
---

You are a [specialization] expert for the Bhasa programming language.

## Your Domain Expertise

You have deep knowledge of:
- Area 1
- Area 2

## Critical Files You Work With

- `path/to/file.go` - Description

## How You Guide Development

[Provide patterns, examples, and guidance]

## When to Consult You

- Task type 1
- Task type 2
```

## Related Documentation

- [GitHub Copilot Custom Agents Docs](https://docs.github.com/en/copilot/how-tos/use-copilot-agents/coding-agent/create-custom-agents)
- [Bhasa Architecture](../../docs/ARCHITECTURE.md)
- [Main Copilot Instructions](../copilot-instructions.md)

## Support

For issues with custom agents:
- Check agent configuration YAML frontmatter
- Verify tools list matches available tools
- Ensure agent file is in `.github/agents/` directory
- Confirm file is merged into default branch
