# Evaluator Documentation

Welcome to the Bhasa evaluator documentation!

## 📚 Documentation Files

### [Evaluator Documentation](./evaluator-documentation.md)
Comprehensive documentation for `evaluator.go`:
- Tree-walking interpreter architecture
- Complete evaluation process for all AST node types
- Statement evaluation (programs, blocks, returns, etc.)
- Expression evaluation (literals, operations, functions, etc.)
- Control flow evaluation (if-else, while loops)
- Function application and closures
- Environment management
- Error handling

**Recommended for**: Understanding the interpretation process, implementing new language features, debugging evaluation

### [Builtins Documentation](./builtins-documentation.md)
Complete reference for `builtins.go`:
- All built-in functions with Bengali names
- Function signatures and behavior
- Implementation details
- Usage examples
- Error handling

**Recommended for**: Using built-in functions, adding new builtins, understanding standard library

### [Quick Reference](./quick-reference.md)
Concise reference guide:
- Evaluation patterns
- Built-in function lookup
- Common operations
- Error handling
- Troubleshooting

**Recommended for**: Quick lookups, day-to-day development

### [Evaluation Examples](./evaluation-examples.md)
Visual learning guide:
- Step-by-step evaluation traces
- Environment state changes
- Stack traces
- Real-world patterns

**Recommended for**: Learning the evaluation process, visual learners

---

## 🎯 What is the Evaluator?

The Bhasa evaluator is a **tree-walking interpreter** that directly executes Abstract Syntax Tree (AST) nodes. It traverses the AST and evaluates each node, producing runtime values.

### Interpretation Pipeline

```
Source Code
    ↓
  Lexer
    ↓
  Tokens
    ↓
  Parser
    ↓
   AST
    ↓
Evaluator ← You are here
    ↓
 Result
```

---

## 🏗️ Architecture Overview

### Core Components

```
┌─────────────────────────────────────────────┐
│              Evaluator                      │
├─────────────────────────────────────────────┤
│  • Eval() function   - Main entry point    │
│  • Environment       - Variable storage     │
│  • Builtins          - Standard library     │
│  • Error handling    - Runtime errors       │
└─────────────────────────────────────────────┘
         │
         ├── Environment (variable scope)
         │   • Nested scopes
         │   • Variable binding
         │   • Closure support
         │
         └── Object System (runtime values)
             • Integer, String, Boolean
             • Array, Hash
             • Function, Closure
             • Return, Error, Null
```

---

## 📊 Key Concepts

### 1. Tree-Walking Interpretation

The evaluator recursively walks the AST:

```go
func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    case *ast.InfixExpression:
        left := Eval(node.Left, env)
        right := Eval(node.Right, env)
        return evalInfixExpression(node.Operator, left, right)
    // ... more cases
    }
}
```

### 2. Environment Chain

Variables are stored in nested environments:

```
Global Environment
    ↓
Function Environment
    ↓
Nested Function Environment
```

### 3. Object System

Everything evaluates to an object:

```
Expression → Evaluation → Object
5 + 3      → Eval()     → &Integer{Value: 8}
"hi"       → Eval()     → &String{Value: "hi"}
সত্য        → Eval()     → &Boolean{Value: true}
```

### 4. Error Propagation

Errors bubble up through the call stack:

```go
if isError(result) {
    return result  // Stop and propagate error
}
```

---

## 🔑 Key Types

### Eval Function

```go
func Eval(node ast.Node, env *object.Environment) object.Object
```

**Purpose**: Main evaluation entry point

**Parameters**:
- `node`: AST node to evaluate
- `env`: Current environment for variable lookup

**Returns**: Object representing the evaluated result

---

### Environment

```go
type Environment struct {
    store map[string]object.Object
    outer *object.Environment
}
```

**Purpose**: Stores variable bindings

**Methods**:
- `Get(name string) (object.Object, bool)` - Lookup variable
- `Set(name string, val object.Object)` - Bind variable
- `NewEnclosedEnvironment(outer)` - Create nested scope

---

### Object Types

All runtime values implement `object.Object`:

| Type | Description | Example |
|------|-------------|---------|
| `Integer` | Numbers | `42`, `-15` |
| `String` | Text | `"হ্যালো"` |
| `Boolean` | True/False | `সত্য`, `মিথ্যা` |
| `Array` | Lists | `[1, 2, 3]` |
| `Hash` | Dictionaries | `{"key": "value"}` |
| `Function` | User functions | `ফাংশন(x) { x + 1 }` |
| `Builtin` | Built-in functions | `len`, `দৈর্ঘ্য` |
| `ReturnValue` | Return wrapper | Wraps return values |
| `Error` | Errors | Runtime errors |
| `Null` | Null/void | `null` |

---

## 📖 Evaluation Process

### Example: Arithmetic

**Bhasa Code**:
```bhasa
5 + 3
```

**Evaluation Steps**:

1. **Eval InfixExpression**
2. **Eval left operand** → `&Integer{Value: 5}`
3. **Eval right operand** → `&Integer{Value: 3}`
4. **Apply operator** → `&Integer{Value: 8}`

### Example: Variable

**Bhasa Code**:
```bhasa
ধরি x = 10;
x + 5;
```

**Evaluation Steps**:

1. **Eval LetStatement**:
   - Eval value expression → `&Integer{Value: 10}`
   - Store in environment: `env.Set("x", &Integer{10})`

2. **Eval InfixExpression**:
   - Eval identifier "x" → lookup in env → `&Integer{10}`
   - Eval literal 5 → `&Integer{5}`
   - Apply + → `&Integer{15}`

### Example: Function Call

**Bhasa Code**:
```bhasa
ধরি add = ফাংশন(a, b) { a + b };
add(5, 3);
```

**Evaluation Steps**:

1. **Eval LetStatement**:
   - Create function object
   - Store in environment

2. **Eval CallExpression**:
   - Eval function → lookup "add"
   - Eval arguments → `[&Integer{5}, &Integer{3}]`
   - Apply function:
     - Create new environment (extends function's env)
     - Bind parameters: `a=5, b=3`
     - Eval body: `a + b` → `&Integer{8}`
     - Return result

---

## 🎓 Getting Started

### For New Developers

1. **Understand object system**: Review object types
2. **Study Eval() function**: Main entry point
3. **Learn environment chain**: Variable scoping
4. **Follow examples**: Work through evaluation traces
5. **Read detailed docs**: Study specific node evaluation

### For Feature Development

1. **Identify AST node** to evaluate
2. **Add case to Eval()** function
3. **Implement evaluation** logic
4. **Handle errors** properly
5. **Update environment** if needed
6. **Add tests**

### For Debugging

1. **Print intermediate values**:
```go
fmt.Printf("Evaluating: %T\n", node)
fmt.Printf("Result: %s\n", result.Inspect())
```

2. **Check environment**:
```go
fmt.Printf("Env: %+v\n", env.store)
```

3. **Trace evaluation**:
```go
func Eval(node ast.Node, env *object.Environment) object.Object {
    fmt.Printf("-> Eval: %T\n", node)
    result := evalNode(node, env)
    fmt.Printf("<- Result: %s\n", result.Inspect())
    return result
}
```

---

## 💡 Common Patterns

### Pattern 1: Evaluate and Check Error

```go
result := Eval(node, env)
if isError(result) {
    return result
}
// Use result...
```

### Pattern 2: Evaluate Multiple Expressions

```go
values := []object.Object{}
for _, expr := range expressions {
    val := Eval(expr, env)
    if isError(val) {
        return val
    }
    values = append(values, val)
}
```

### Pattern 3: Apply Binary Operation

```go
left := Eval(node.Left, env)
if isError(left) {
    return left
}
right := Eval(node.Right, env)
if isError(right) {
    return right
}
return applyOperator(operator, left, right)
```

### Pattern 4: Function Application

```go
fn := Eval(fnExpr, env)
if isError(fn) {
    return fn
}
args := evalExpressions(argExprs, env)
if len(args) == 1 && isError(args[0]) {
    return args[0]
}
return applyFunction(fn, args)
```

---

## 🔍 Built-in Functions

### Bengali Functions

| Bengali | English | Purpose | Example |
|---------|---------|---------|---------|
| `লেখ` | write/print | Print to console | `লেখ("হ্যালো")` |
| `দৈর্ঘ্য` | length | Get length | `দৈর্ঘ্য([1,2,3])` → 3 |
| `প্রথম` | first | First element | `প্রথম([1,2,3])` → 1 |
| `শেষ` | last | Last element | `শেষ([1,2,3])` → 3 |
| `বাকি` | rest | All but first | `বাকি([1,2,3])` → [2,3] |
| `যোগ` | push | Add to array | `যোগ([1,2], 3)` → [1,2,3] |
| `টাইপ` | type | Get type | `টাইপ(5)` → "INTEGER" |

---

## 📚 Environment Management

### Global Environment

```go
env := object.NewEnvironment()
env.Set("x", &object.Integer{Value: 5})
val, ok := env.Get("x")  // &Integer{5}, true
```

### Nested Environment

```go
outer := object.NewEnvironment()
outer.Set("x", &object.Integer{Value: 5})

inner := object.NewEnclosedEnvironment(outer)
inner.Set("y", &object.Integer{Value: 10})

// Inner can access outer
inner.Get("x")  // &Integer{5}, true
inner.Get("y")  // &Integer{10}, true

// Outer cannot access inner
outer.Get("y")  // nil, false
```

### Function Environment

```go
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
    env := object.NewEnclosedEnvironment(fn.Env)
    
    for i, param := range fn.Parameters {
        env.Set(param.Value, args[i])
    }
    
    return env
}
```

---

## 🚨 Error Handling

### Creating Errors

```go
func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{Message: fmt.Sprintf(format, a...)}
}
```

### Checking Errors

```go
func isError(obj object.Object) bool {
    if obj != nil {
        return obj.Type() == object.ERROR_OBJ
    }
    return false
}
```

### Propagating Errors

```go
result := Eval(node, env)
if isError(result) {
    return result  // Stop and return error
}
```

### Common Errors

- `"identifier not found: x"` - Undefined variable
- `"unknown operator: INTEGER + STRING"` - Type mismatch
- `"type mismatch: INTEGER + STRING"` - Incompatible types
- `"not a function: INTEGER"` - Calling non-function
- `"wrong number of arguments"` - Incorrect arg count
- `"division by zero"` - Division by zero

---

## 🔧 Extending the Evaluator

### Adding a New Built-in Function

1. **Add to builtins map**:

```go
var builtins = map[string]*object.Builtin{
    "নতুন_ফাংশন": {
        Fn: func(args ...object.Object) object.Object {
            // Validate arguments
            if len(args) != 1 {
                return newError("wrong number of arguments")
            }
            
            // Implement logic
            // ...
            
            return result
        },
    },
}
```

2. **Add tests**
3. **Update documentation**

---

### Adding New Expression Evaluation

1. **Add case to Eval()**:

```go
case *ast.MyNewExpression:
    return evalMyNewExpression(node, env)
```

2. **Implement evaluation function**:

```go
func evalMyNewExpression(node *ast.MyNewExpression, env *object.Environment) object.Object {
    // Evaluate sub-expressions
    val := Eval(node.SubExpr, env)
    if isError(val) {
        return val
    }
    
    // Process and return result
    return result
}
```

3. **Add tests**

---

## 📚 Related Documentation

- **[AST Documentation](../../ast/docs/)** - Input to evaluator
- **[Object System](../../object/docs/)** - Runtime values
- **[Parser Documentation](../../parser/docs/)** - AST generation

---

## 🧪 Testing

### Unit Tests

```go
func TestEvalIntegerExpression(t *testing.T) {
    tests := []struct {
        input    string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
        {"-5", -5},
        {"5 + 5", 10},
        {"5 - 3", 2},
    }
    
    for _, tt := range tests {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}
```

---

## 💭 Design Decisions

### Why Tree-Walking?

**Pros**:
- ✅ Simple implementation
- ✅ Easy to understand
- ✅ Direct AST interpretation
- ✅ Good for prototyping

**Cons**:
- ❌ Slower than bytecode VM
- ❌ No optimizations
- ❌ Higher memory usage

**Conclusion**: Good for initial implementation, can migrate to VM later

### Why Singleton Objects?

```go
var (
    NULL  = &object.Null{}
    TRUE  = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
}
```

**Benefits**:
- ✅ Reduce allocations
- ✅ Fast equality checks (pointer comparison)
- ✅ Memory efficient

---

## 🤝 Contributing

When modifying the evaluator:

1. ✅ **Handle all AST node types**
2. ✅ **Check for errors** before using values
3. ✅ **Propagate errors** up the call stack
4. ✅ **Update environment** correctly
5. ✅ **Add comprehensive tests**
6. ✅ **Update documentation**

---

## 📞 Need Help?

- Check [Evaluator Documentation](./evaluator-documentation.md) for detailed explanations
- Check [Builtins Documentation](./builtins-documentation.md) for built-in functions
- Check [Quick Reference](./quick-reference.md) for quick lookups
- Check [Evaluation Examples](./evaluation-examples.md) for visual examples
- Review test files for usage patterns

---

## 🗺️ Evaluator System Overview

```
                    ┌──────────────┐
                    │ Source Code  │
                    └──────┬───────┘
                           │
                           ↓
                    ┌──────────────┐
                    │   Parser     │
                    └──────┬───────┘
                           │ AST
                           ↓
        ┌──────────────────────────────────────┐
        │          Evaluator                   │
        ├──────────────────────────────────────┤
        │                                      │
        │  ┌────────────────┐  ┌────────────┐ │
        │  │  Environment   │  │  Builtins  │ │
        │  ├────────────────┤  ├────────────┤ │
        │  │ • Variables    │  │ • লেখ      │ │
        │  │ • Functions    │  │ • দৈর্ঘ্য   │ │
        │  │ • Scopes       │  │ • প্রথম     │ │
        │  └────────────────┘  │ • শেষ      │ │
        │                      │ • বাকি     │ │
        │  ┌────────────────┐  │ • যোগ      │ │
        │  │  Eval()        │  │ • টাইপ     │ │
        │  ├────────────────┤  └────────────┘ │
        │  │ • Type switch  │                 │
        │  │ • Recursion    │                 │
        │  │ • Error check  │                 │
        │  └────────────────┘                 │
        │                                      │
        └──────────────┬───────────────────────┘
                       │
                       ↓
            ┌──────────────────┐
            │  Runtime Values  │
            │  (Objects)       │
            └──────────────────┘
```

---

**Happy Evaluating! 🚀**

