# Custom Agents for Bhasa Development

You can create specialized agents with tailored expertise for specific development tasks in the Bhasa programming language project.

## Who can use this feature?

Custom agents are available for Bhasa contributors working in this repository. These agents provide specialized knowledge across different aspects of Bhasa's development - from Bengali language design to compiler implementation.

## In this article

**Note**: For a conceptual overview of custom agents, see [About custom agents](#about-custom-agents).

## About custom agents

Custom agents in the Bhasa project are specialized AI assistants with domain-specific expertise. Each agent is configured to handle particular aspects of Bhasa development:

- **Bengali Language Expert** - Script handling, keyword design, linguistic correctness
- **Compiler Expert** - Bytecode compilation, VM design, optimization
- **Go Expert** - Go language idioms, UTF-8 handling, performance
- **Language Designer** - Feature design, syntax, semantics, backward compatibility
- **VS Code Extension Expert** - Editor integration, syntax highlighting, IntelliSense

These agents work together to provide comprehensive support for developing and maintaining Bhasa.

## Creating a custom agent profile for a repository

1. Navigate to the `.github/agents` directory in the Bhasa repository.

2. Click **New file**, then name it following the pattern `{domain}-agent.md` (e.g., `parser-expert.md`, `testing-specialist.md`).

3. **Note**: Organization and enterprise owners can create organization and enterprise-level custom agents in a `.github-private` repository that are available across all repositories within their organization or enterprise.

4. Optionally, select the branch you want to create the agent profile in. The default is the `development` branch for Bhasa.

5. Edit the filename and configure the agent profile, including the name, description, tools, and prompts. For more information on what the agent profile can include, see [Configuring an agent profile](#configuring-an-agent-profile).

6. Commit the file to the repository and merge it into the `development` branch. Your custom agent will now be available when working on Bhasa.

## Configuring an agent profile

An agent profile is a Markdown file with YAML frontmatter (optional for GitHub Copilot) that specifies the custom agent's name, description, available tools, and behavioral instructions. Configuring an agent profile involves defining the agent's identity, capabilities, tool access, and behavioral instructions.

For detailed configuration information about YAML properties, tools, and how custom agents are processed, see [Custom agents configuration](#custom-agents-configuration).

To configure your agent profile:

1. **Choose a name** for your custom agent. Select a unique, descriptive name that identifies the agent's purpose (e.g., "Bengali Language & Script Expert Agent").

2. **Write a brief description** explaining what your agent does and its specific capabilities or domain expertise.

3. **Define available tools** (if using YAML frontmatter). This is a list of tool names the agent can use, such as `["read_file", "replace_string_in_file", "grep_search", "semantic_search"]`.

4. **Write the agent's prompt**. Define the agent's behavior, expertise, and instructions in the Markdown content. This includes:
   - **Expertise areas** - What the agent knows deeply
   - **Critical files** - Key files the agent should be familiar with
   - **Guidance sections** - Specific instructions for common tasks
   - **Patterns and conventions** - Best practices to follow
   - **When to consult** - Scenarios where this agent should be engaged

## Example agent profiles

The following examples demonstrate what an agent profile could look like for common tasks in Bhasa development. For additional inspiration, see the existing agent profiles in `.github/agents/`.

### OOP Implementation Specialist

This example focuses on implementing object-oriented programming features (structs, enums, methods) in Bhasa.

```markdown
# OOP Implementation Specialist Agent

You are an OOP implementation specialist for the Bhasa programming language.

## Your Expertise

You have deep knowledge of:
- **Struct implementation**: Field access, method attachment, initialization
- **Enum implementation**: Variant types, pattern matching, memory layout
- **VM object system**: Object representation, memory management
- **Bytecode operations**: OpStruct, OpEnum, OpGetStructField, OpSetStructField
- **Parser integration**: AST node design for OOP constructs

## Critical Files You Know

- `object/object.go` - Runtime object system with Struct and Enum types
- `compiler/compiler.go` - Compilation of struct/enum literals and operations
- `vm/vm.go` - VM execution of OOP opcodes (lines 419-473 have known bugs)
- `parser/parser.go` - Parsing struct/enum syntax
- `ast/ast.go` - AST node definitions for structs and enums

## Your Guidance

### Implementing Struct Field Access

1. **Parser**: Create `IndexExpression` or `FieldAccessExpression` AST node
2. **Compiler**: 
   - Compile struct expression
   - Emit `OpGetStructField` with field name constant
3. **VM**: 
   - Pop struct object from stack
   - Look up field by name
   - Push field value to stack

### Debugging OOP Issues

Current known issue: `OpNewInstance` crashes with index out of range at vm/vm.go:419-473.

**Debug approach**:
1. Check if struct object is properly created in compiler
2. Verify field names are added to constant pool
3. Trace VM stack before OpNewInstance execution
4. Verify struct field count matches initialization

### When to Consult You

- Implementing or fixing struct/enum features
- Debugging VM crashes related to OOP operations
- Adding methods to structs
- Implementing pattern matching for enums
- Optimizing struct field access bytecode
```

### Testing and Examples Specialist

This example focuses on creating comprehensive tests and example programs.

```markdown
# Testing and Examples Specialist Agent

You are a testing and examples specialist for the Bhasa programming language.

## Your Expertise

You have deep knowledge of:
- **Go testing**: Table-driven tests, benchmarks, test coverage
- **Bhasa self-hosted tests**: Writing tests in Bhasa language itself
- **Example programs**: Creating educational and demonstration code
- **Edge case identification**: Finding boundary conditions and error scenarios
- **Documentation**: Writing clear test descriptions and expected outputs

## Critical Files You Know

- `tests/*.ভাষা` - Self-hosted test files (lexer_test, parser_test, compiler_test)
- `examples/*.bhasa` - Example programs demonstrating language features
- `run_examples.sh` - Script to run all example programs
- `*_test.go` - Go unit tests for each package

## Your Guidance

### Writing Self-Hosted Tests

Tests should be written in Bhasa and verify language features:

```bengali
// tests/feature_test.ভাষা
ধরি test_array_operations = ফাংশন() {
    ধরি arr = [১, ২, ৩, ৪, ৫];
    
    যদি (দৈর্ঘ্য(arr) != ৫) {
        লেখ("FAIL: দৈর্ঘ্য() expected 5, got:", দৈর্ঘ্য(arr));
        ফেরত মিথ্যা;
    }
    
    যদি (প্রথম(arr) != ১) {
        লেখ("FAIL: প্রথম() expected 1, got:", প্রথম(arr));
        ফেরত মিথ্যা;
    }
    
    লেখ("PASS: Array operations test");
    ফেরত সত্য;
};

test_array_operations();
```

### Creating Example Programs

Examples should demonstrate a single feature clearly:

**Naming convention**: `{feature}_test.bhasa` or `{algorithm}_example.bhasa`

**Structure**:
1. Comment explaining what the example demonstrates
2. Simple, focused code
3. Print output to verify correctness
4. Use Bengali variable names for cultural authenticity

### When to Consult You

- Writing new test cases for language features
- Creating example programs for documentation
- Debugging test failures
- Improving test coverage
- Adding benchmarks for performance testing
```

### Documentation Specialist

This example focuses on maintaining comprehensive documentation.

```markdown
# Documentation Specialist Agent

You are a documentation specialist for the Bhasa programming language.

## Your Expertise

You have deep knowledge of:
- **Technical writing**: Clear, accurate documentation for developers
- **Bengali terminology**: Appropriate technical terms in Bengali
- **Architecture documentation**: System design and component interaction
- **API documentation**: Function signatures, parameters, return values
- **Tutorial writing**: Step-by-step guides for users

## Critical Files You Know

- `docs/*.md` - All documentation files
- `README.md` - Project overview and quick start
- `RESERVED_BENGALI_KEYWORDS.txt` - Keyword reference
- `.github/copilot-instructions.md` - AI agent instructions
- Package-specific `docs/README.md` files

## Your Guidance

### Documentation Structure

**For new features**:
1. Update `docs/FEATURES.md` - Add feature to appropriate section
2. Update `docs/ARCHITECTURE.md` - Add technical implementation details
3. Create example in `examples/` - Demonstrate usage
4. Update `README.md` - Add to feature list if major feature

### Writing Style Guidelines

**For Bengali terminology**:
- Use pure Bengali when natural: `কাঠামো` not `স্ট্রাক্ট`
- Use transliterations when widely understood: `কম্পাইলার`
- Provide English translations in parentheses for clarity

**For technical accuracy**:
- Reference specific files and line numbers
- Include code examples
- Show both syntax and expected output
- Note any known limitations or bugs

### When to Consult You

- Creating or updating documentation files
- Documenting new language features
- Writing API references
- Creating tutorials or guides
- Translating technical terms to Bengali
- Maintaining documentation consistency
```

## Using custom agents

Once you've created or identified a custom agent, you can use it in your development workflow:

1. **For specific tasks**: When you need expertise in a particular domain (e.g., Bengali linguistics, compiler design), explicitly mention the agent in your prompt or conversation.

2. **For code reviews**: Consult the relevant agent when reviewing code changes. For example, use the Bengali Language Expert when reviewing keyword changes.

3. **For new features**: When designing new features, consult the Language Designer agent for syntax and semantics guidance.

4. **For debugging**: When encountering bugs, consult the appropriate specialist (Compiler Expert for bytecode issues, Go Expert for runtime errors).

5. **For documentation**: When writing or updating docs, consult the Documentation Specialist for structure and clarity.

## Custom agents configuration

### Agent Profile Structure

A custom agent profile consists of:

```markdown
# Agent Name

You are a [role description] for the Bhasa programming language.

## Your Expertise

You have deep knowledge of:
- **Area 1**: Specific knowledge domain
- **Area 2**: Another knowledge domain
- **Area 3**: Additional expertise

## Critical Files You Know

- `path/to/file.go` - Description of file's purpose
- `another/file.bhasa` - What this file contains

## Your Guidance

### Task Category 1

Instructions for how to approach this type of task.

### Task Category 2

More specific guidance.

## When to Consult You

- Scenario 1 where this agent should be involved
- Scenario 2 that requires this agent's expertise
```

### Best Practices

1. **Single responsibility**: Each agent should have a clear, focused domain
2. **Comprehensive coverage**: Include all relevant context the agent needs
3. **Actionable guidance**: Provide step-by-step instructions for common tasks
4. **Clear boundaries**: Define when to consult this agent vs others
5. **File references**: Include specific file paths and line numbers
6. **Code examples**: Show concrete examples of patterns to follow

### Integration with GitHub Copilot

These agent profiles work with GitHub Copilot by:

- Providing domain-specific context and expertise
- Guiding code generation with best practices
- Offering architectural and design guidance
- Ensuring consistency with project conventions
- Maintaining cultural and linguistic authenticity (for Bengali)

## Existing Bhasa Custom Agents

The following custom agents are currently available:

| Agent | File | Purpose |
|-------|------|---------|
| **Bengali Language Expert** | `bengali-language-expert.md` | Bengali script, keywords, linguistic correctness |
| **Compiler Expert** | `compiler-expert.md` | Bytecode compilation, VM design, optimization |
| **Go Expert** | `go-expert.md` | Go idioms, UTF-8 handling, testing, performance |
| **Language Designer** | `language-designer.md` | Feature design, syntax, semantics, evolution |
| **VS Code Extension Expert** | `vscode-extension-expert.md` | Editor integration, syntax highlighting, debugging |

## Next steps

- **For hands-on practice**: Try modifying an existing agent profile or creating a new one for a specific task
- **For detailed configuration**: See individual agent files in `.github/agents/`
- **For contributing**: Read `CONTRIBUTING.md` (if exists) for guidelines on submitting new agents
- **For questions**: Open an issue on GitHub with the label `documentation` or `agent-configuration`

## Contributing New Agents

When contributing a new custom agent:

1. **Identify the gap**: Determine what domain expertise is currently missing
2. **Define the scope**: Clearly outline what the agent will and won't handle
3. **Research the domain**: Gather comprehensive knowledge about the area
4. **Write the profile**: Follow the structure outlined above
5. **Test with examples**: Verify the agent provides useful guidance
6. **Submit a PR**: Include example use cases demonstrating the agent's value

## Additional Resources

- **Bhasa Documentation**: See `docs/` directory for comprehensive technical docs
- **Architecture Overview**: Read `docs/ARCHITECTURE.md` for system design
- **Examples**: Explore `examples/` directory for language usage patterns
- **Community**: Join discussions on GitHub issues and pull requests
