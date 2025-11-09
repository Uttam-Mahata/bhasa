# Bhasa Evaluator Documentation

## Table of Contents

1. [Overview](#overview)
2. [Core Functions](#core-functions)
3. [Statement Evaluation](#statement-evaluation)
4. [Expression Evaluation](#expression-evaluation)
5. [Operator Evaluation](#operator-evaluation)
6. [Function Application](#function-application)
7. [Environment Management](#environment-management)
8. [Error Handling](#error-handling)
9. [Helper Functions](#helper-functions)
10. [Evaluation Traces](#evaluation-traces)

---

## Overview

The Bhasa evaluator is a **tree-walking interpreter** that directly executes AST nodes by recursively traversing the tree and evaluating each node. It maintains an environment chain for variable scoping and supports closures through lexical scoping.

### Key Characteristics

- **Direct interpretation**: No intermediate bytecode
- **Recursive evaluation**: Each node evaluates its children
- **Environment chain**: Nested scopes for variables
- **Error propagation**: Errors bubble up immediately
- **Object-based**: All values are objects

---

## Core Functions

### Eval

```go
func Eval(node ast.Node, env *object.Environment) object.Object
```

**Purpose**: Main evaluation entry point - recursively evaluates any AST node

**Parameters**:
- `node`: The AST node to evaluate
- `env`: Current environment for variable lookup

**Returns**: An `object.Object` representing the evaluated result

**Structure**:
```go
func Eval(node ast.Node, env *object.Environment) object.Object {
    switch node := node.(type) {
    case *ast.Program:
        return evalProgram(node, env)
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    // ... more cases
    }
    return nil
}
```

**Design Pattern**: Big type switch on node type
- Each case handles one AST node type
- Delegates to helper functions for complex nodes
- Returns nil for unknown node types

---

## Statement Evaluation

### Program

```go
case *ast.Program:
    return evalProgram(node, env)
```

**Evaluates all statements** in sequence:

```go
func evalProgram(program *ast.Program, env *object.Environment) object.Object {
    var result object.Object
    
    for _, statement := range program.Statements {
        result = Eval(statement, env)
        
        switch result := result.(type) {
        case *object.ReturnValue:
            return result.Value  // Unwrap and stop
        case *object.Error:
            return result        // Propagate error
        }
    }
    
    return result  // Return last statement's value
}
```

**Key Points**:
- Evaluates statements sequentially
- **Returns immediately** on ReturnValue or Error
- Returns the last statement's value
- Unwraps ReturnValue before returning

**Example**:
```bhasa
ধরি x = 5;
ধরি y = 10;
x + y;  // Returns 15
```

---

### ExpressionStatement

```go
case *ast.ExpressionStatement:
    return Eval(node.Expression, env)
```

**Simply evaluates** the wrapped expression

**Example**:
```bhasa
5 + 3;  // ExpressionStatement wrapping InfixExpression
```

---

### BlockStatement

```go
case *ast.BlockStatement:
    return evalBlockStatement(node, env)
```

**Evaluates statements** in a block:

```go
func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
    var result object.Object
    
    for _, statement := range block.Statements {
        result = Eval(statement, env)
        
        if result != nil {
            rt := result.Type()
            if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
                return result  // Stop and return (don't unwrap)
            }
        }
    }
    
    return result
}
```

**Key Difference from Program**:
- Does **NOT unwrap** ReturnValue
- Allows return to propagate through nested blocks
- Returns immediately on Error

**Example**:
```bhasa
{
    ধরি x = 5;
    ধরি y = 10;
    x + y  // Block value
}
```

---

### LetStatement

```go
case *ast.LetStatement:
    val := Eval(node.Value, env)
    if isError(val) {
        return val
    }
    env.Set(node.Name.Value, val)
```

**Steps**:
1. Evaluate the value expression
2. Check for errors
3. Store in environment

**Example**:
```bhasa
ধরি x = 5 + 3;
// 1. Eval(5 + 3) → &Integer{8}
// 2. env.Set("x", &Integer{8})
```

**No Return Value**: LetStatement doesn't return a value

---

### AssignmentStatement

```go
case *ast.AssignmentStatement:
    val := Eval(node.Value, env)
    if isError(val) {
        return val
    }
    env.Set(node.Name.Value, val)
```

**Same as LetStatement**: Updates existing or creates new binding

**Example**:
```bhasa
x = 20;
// 1. Eval(20) → &Integer{20}
// 2. env.Set("x", &Integer{20})
```

---

### ReturnStatement

```go
case *ast.ReturnStatement:
    val := Eval(node.ReturnValue, env)
    if isError(val) {
        return val
    }
    return &object.ReturnValue{Value: val}
```

**Wraps value** in ReturnValue object:
- Signals to stop evaluation
- Propagates through blocks
- Unwrapped by evalProgram()

**Example**:
```bhasa
ফাংশন() {
    ফেরত 5 + 3;
    // Never reached
    10;
}
```

---

### WhileStatement

```go
case *ast.WhileStatement:
    return evalWhileStatement(node, env)
```

**Loops while condition is truthy**:

```go
func evalWhileStatement(ws *ast.WhileStatement, env *object.Environment) object.Object {
    var result object.Object = NULL
    
    for {
        // Evaluate condition
        condition := Eval(ws.Condition, env)
        if isError(condition) {
            return condition
        }
        
        // Check truthiness
        if !isTruthy(condition) {
            break
        }
        
        // Evaluate body
        result = Eval(ws.Body, env)
        if isError(result) {
            return result
        }
        
        // Handle return in loop
        if result != nil && result.Type() == object.RETURN_VALUE_OBJ {
            return result
        }
    }
    
    return result
}
```

**Key Points**:
- Re-evaluates condition each iteration
- Breaks on falsy condition
- Returns immediately on error or return
- Returns last iteration's value

**Example**:
```bhasa
ধরি i = 0;
যতক্ষণ (i < 5) {
    লেখ(i);
    i = i + 1;
}
// Prints: 0, 1, 2, 3, 4
```

---

## Expression Evaluation

### IntegerLiteral

```go
case *ast.IntegerLiteral:
    return &object.Integer{Value: node.Value}
```

**Creates Integer object** from AST node

**Example**:
```bhasa
42  // → &Integer{Value: 42}
```

---

### StringLiteral

```go
case *ast.StringLiteral:
    return &object.String{Value: node.Value}
```

**Creates String object** from AST node

**Example**:
```bhasa
"হ্যালো"  // → &String{Value: "হ্যালো"}
```

---

### Boolean

```go
case *ast.Boolean:
    return nativeBoolToBooleanObject(node.Value)
```

**Returns singleton** boolean objects:

```go
func nativeBoolToBooleanObject(input bool) *object.Boolean {
    if input {
        return TRUE   // Singleton
    }
    return FALSE  // Singleton
}
```

**Singletons**:
```go
var (
    TRUE  = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
)
```

**Benefits**:
- Only one TRUE and one FALSE object
- Fast equality checks (pointer comparison)
- Memory efficient

---

### Identifier

```go
case *ast.Identifier:
    return evalIdentifier(node, env)
```

**Looks up** variable in environment:

```go
func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
    // Check environment first
    if val, ok := env.Get(node.Value); ok {
        return val
    }
    
    // Check builtins
    if builtin, ok := builtins[node.Value]; ok {
        return builtin
    }
    
    // Not found
    return newError("identifier not found: " + node.Value)
}
```

**Lookup Order**:
1. Current environment (and parent scopes)
2. Builtins
3. Error if not found

**Example**:
```bhasa
ধরি x = 5;
x  // → Lookup "x" → &Integer{5}
```

---

### PrefixExpression

```go
case *ast.PrefixExpression:
    right := Eval(node.Right, env)
    if isError(right) {
        return right
    }
    return evalPrefixExpression(node.Operator, right)
```

**Evaluates** prefix operators:

```go
func evalPrefixExpression(operator string, right object.Object) object.Object {
    switch operator {
    case "!":
        return evalBangOperatorExpression(right)
    case "-":
        return evalMinusPrefixOperatorExpression(right)
    default:
        return newError("unknown operator: %s%s", operator, right.Type())
    }
}
```

**Supported Operators**:
- `!` - Logical NOT
- `-` - Negation

---

### InfixExpression

```go
case *ast.InfixExpression:
    left := Eval(node.Left, env)
    if isError(left) {
        return left
    }
    right := Eval(node.Right, env)
    if isError(right) {
        return right
    }
    return evalInfixExpression(node.Operator, left, right)
```

**Evaluates** binary operators:

```go
func evalInfixExpression(operator string, left, right object.Object) object.Object {
    switch {
    case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
        return evalIntegerInfixExpression(operator, left, right)
    case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
        return evalStringInfixExpression(operator, left, right)
    case operator == "==":
        return nativeBoolToBooleanObject(left == right)
    case operator == "!=":
        return nativeBoolToBooleanObject(left != right)
    case left.Type() != right.Type():
        return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
    default:
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
    }
}
```

**Type Dispatch**:
1. Both integers → integer arithmetic
2. Both strings → string concatenation
3. Equality operators → pointer comparison
4. Type mismatch → error
5. Unknown operator → error

---

### IfExpression

```go
case *ast.IfExpression:
    return evalIfExpression(node, env)
```

**Conditional evaluation**:

```go
func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
    // Evaluate condition
    condition := Eval(ie.Condition, env)
    if isError(condition) {
        return condition
    }
    
    // Check truthiness
    if isTruthy(condition) {
        return Eval(ie.Consequence, env)
    } else if ie.Alternative != nil {
        return Eval(ie.Alternative, env)
    } else {
        return NULL
    }
}
```

**Truthiness**:
```go
func isTruthy(obj object.Object) bool {
    switch obj {
    case NULL:
        return false
    case TRUE:
        return true
    case FALSE:
        return false
    default:
        return true  // Everything else is truthy
    }
}
```

**Example**:
```bhasa
যদি (x > 5) {
    "বড়"
} নাহলে {
    "ছোট"
}
```

---

### FunctionLiteral

```go
case *ast.FunctionLiteral:
    params := node.Parameters
    body := node.Body
    return &object.Function{Parameters: params, Env: env, Body: body}
```

**Creates Function object**:
- Stores parameters
- Captures current environment (closure)
- Stores body AST

**Example**:
```bhasa
ফাংশন(x, y) {
    x + y
}
// → &Function{
//     Parameters: [x, y],
//     Env: currentEnv,
//     Body: BlockStatement{...}
// }
```

**Key Point**: Environment is **captured** at function creation time (lexical scoping)

---

### CallExpression

```go
case *ast.CallExpression:
    function := Eval(node.Function, env)
    if isError(function) {
        return function
    }
    args := evalExpressions(node.Arguments, env)
    if len(args) == 1 && isError(args[0]) {
        return args[0]
    }
    return applyFunction(function, args)
```

**Steps**:
1. Evaluate function expression
2. Evaluate all arguments
3. Apply function to arguments

**Example**:
```bhasa
add(5, 3)
// 1. Eval(add) → &Function{...}
// 2. Eval([5, 3]) → [&Integer{5}, &Integer{3}]
// 3. applyFunction(...) → &Integer{8}
```

---

### ArrayLiteral

```go
case *ast.ArrayLiteral:
    elements := evalExpressions(node.Elements, env)
    if len(elements) == 1 && isError(elements[0]) {
        return elements[0]
    }
    return &object.Array{Elements: elements}
```

**Evaluates** all elements and creates Array object

**Example**:
```bhasa
[1, 2, 3]
// → &Array{Elements: [&Integer{1}, &Integer{2}, &Integer{3}]}
```

---

### IndexExpression

```go
case *ast.IndexExpression:
    left := Eval(node.Left, env)
    if isError(left) {
        return left
    }
    index := Eval(node.Index, env)
    if isError(index) {
        return index
    }
    return evalIndexExpression(left, index)
```

**Indexes** into arrays or hashes:

```go
func evalIndexExpression(left, index object.Object) object.Object {
    switch {
    case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
        return evalArrayIndexExpression(left, index)
    case left.Type() == object.HASH_OBJ:
        return evalHashIndexExpression(left, index)
    default:
        return newError("index operator not supported: %s", left.Type())
    }
}
```

**Example**:
```bhasa
arr[2]
// 1. Eval(arr) → &Array{...}
// 2. Eval(2) → &Integer{2}
// 3. Get arr.Elements[2]
```

---

### HashLiteral

```go
case *ast.HashLiteral:
    return evalHashLiteral(node, env)
```

**Creates hash** map:

```go
func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
    pairs := make(map[object.HashKey]object.HashPair)
    
    for keyNode, valueNode := range node.Pairs {
        // Evaluate key
        key := Eval(keyNode, env)
        if isError(key) {
            return key
        }
        
        // Check if key is hashable
        hashKey, ok := key.(object.Hashable)
        if !ok {
            return newError("unusable as hash key: %s", key.Type())
        }
        
        // Evaluate value
        value := Eval(valueNode, env)
        if isError(value) {
            return value
        }
        
        // Store pair
        hashed := hashKey.HashKey()
        pairs[hashed] = object.HashPair{Key: key, Value: value}
    }
    
    return &object.Hash{Pairs: pairs}
}
```

**Example**:
```bhasa
{"নাম": "রহিম", "বয়স": 30}
```

---

## Operator Evaluation

### Bang Operator (!)

```go
func evalBangOperatorExpression(right object.Object) object.Object {
    switch right {
    case TRUE:
        return FALSE
    case FALSE:
        return TRUE
    case NULL:
        return TRUE
    default:
        return FALSE
    }
}
```

**Truth Table**:
```
!সত্য    → মিথ্যা
!মিথ্যা   → সত্য
!null   → সত্য
!other  → মিথ্যা
```

---

### Minus Operator (-)

```go
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
    if right.Type() != object.INTEGER_OBJ {
        return newError("unknown operator: -%s", right.Type())
    }
    
    value := right.(*object.Integer).Value
    return &object.Integer{Value: -value}
}
```

**Example**:
```bhasa
-5  // → &Integer{Value: -5}
```

---

### Integer Infix Operators

```go
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
    leftVal := left.(*object.Integer).Value
    rightVal := right.(*object.Integer).Value
    
    switch operator {
    case "+":
        return &object.Integer{Value: leftVal + rightVal}
    case "-":
        return &object.Integer{Value: leftVal - rightVal}
    case "*":
        return &object.Integer{Value: leftVal * rightVal}
    case "/":
        if rightVal == 0 {
            return newError("division by zero")
        }
        return &object.Integer{Value: leftVal / rightVal}
    case "%":
        return &object.Integer{Value: leftVal % rightVal}
    case "<":
        return nativeBoolToBooleanObject(leftVal < rightVal)
    case ">":
        return nativeBoolToBooleanObject(leftVal > rightVal)
    case "<=":
        return nativeBoolToBooleanObject(leftVal <= rightVal)
    case ">=":
        return nativeBoolToBooleanObject(leftVal >= rightVal)
    case "==":
        return nativeBoolToBooleanObject(leftVal == rightVal)
    case "!=":
        return nativeBoolToBooleanObject(leftVal != rightVal)
    default:
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
    }
}
```

**Supported Operators**:
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `<`, `>`, `<=`, `>=`, `==`, `!=`

**Example**:
```bhasa
5 + 3   // → &Integer{8}
10 > 5  // → &Boolean{true}
8 / 0   // → &Error{"division by zero"}
```

---

### String Infix Operators

```go
func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
    if operator != "+" {
        return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
    }
    
    leftVal := left.(*object.String).Value
    rightVal := right.(*object.String).Value
    return &object.String{Value: leftVal + rightVal}
}
```

**Only supports** concatenation (`+`)

**Example**:
```bhasa
"হ্যালো " + "বিশ্ব"  // → &String{"হ্যালো বিশ্ব"}
```

---

## Function Application

### applyFunction

```go
func applyFunction(fn object.Object, args []object.Object) object.Object {
    switch fn := fn.(type) {
    case *object.Function:
        extendedEnv := extendFunctionEnv(fn, args)
        evaluated := Eval(fn.Body, extendedEnv)
        return unwrapReturnValue(evaluated)
    
    case *object.Builtin:
        return fn.Fn(args...)
    
    default:
        return newError("not a function: %s", fn.Type())
    }
}
```

**Two Function Types**:
1. **User functions**: Create new environment and evaluate body
2. **Builtin functions**: Call Go function directly

---

### extendFunctionEnv

```go
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
    env := object.NewEnclosedEnvironment(fn.Env)
    
    for paramIdx, param := range fn.Parameters {
        env.Set(param.Value, args[paramIdx])
    }
    
    return env
}
```

**Creates** new environment:
1. Encloses function's captured environment
2. Binds parameters to arguments

**Example**:
```
Function's Env: {x: 10}
Parameters: [a, b]
Arguments: [5, 3]

New Env:
  Outer: {x: 10}
  Store: {a: 5, b: 3}
```

---

### unwrapReturnValue

```go
func unwrapReturnValue(obj object.Object) object.Object {
    if returnValue, ok := obj.(*object.ReturnValue); ok {
        return returnValue.Value
    }
    return obj
}
```

**Unwraps** ReturnValue after function body evaluation

**Why?**:
- ReturnValue signals early return from function
- Must be unwrapped before returning to caller
- Otherwise return would propagate beyond function boundary

---

### evalExpressions

```go
func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
    var result []object.Object
    
    for _, e := range exps {
        evaluated := Eval(e, env)
        if isError(evaluated) {
            return []object.Object{evaluated}
        }
        result = append(result, evaluated)
    }
    
    return result
}
```

**Evaluates** list of expressions:
- Returns early on first error
- Used for function arguments and array elements

---

## Environment Management

### Environment Structure

```go
type Environment struct {
    store map[string]object.Object
    outer *object.Environment
}
```

**Methods**:
```go
// Get variable (checks outer scopes)
func (e *Environment) Get(name string) (object.Object, bool)

// Set variable (in current scope)
func (e *Environment) Set(name string, val object.Object) object.Object

// Create new environment
func NewEnvironment() *Environment

// Create enclosed environment
func NewEnclosedEnvironment(outer *Environment) *Environment
```

---

### Scope Chain

```
Global Environment
  ↓ outer
Function Environment
  ↓ outer
Nested Function Environment
```

**Lookup**:
1. Check current environment
2. If not found, check outer
3. Repeat until found or no outer
4. Return error if not found

---

### Closure Example

```bhasa
ধরি makeAdder = ফাংশন(x) {
    ফাংশন(y) {
        x + y
    }
};

ধরি add5 = makeAdder(5);
add5(3);  // 8
```

**Environment Chain**:

```
makeAdder called with x=5:
  Outer: Global
  Store: {x: 5}
  Returns: Function{Env: this environment}

add5 called with y=3:
  Outer: {x: 5} ← Captured environment
  Store: {y: 3}
  
  When evaluating x + y:
    - y found in current: 3
    - x found in outer: 5
    - Result: 8
```

---

## Error Handling

### Creating Errors

```go
func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{Message: fmt.Sprintf(format, a...)}
}
```

**Usage**:
```go
return newError("identifier not found: %s", name)
return newError("type mismatch: %s %s %s", left.Type(), op, right.Type())
```

---

### Checking Errors

```go
func isError(obj object.Object) bool {
    if obj != nil {
        return obj.Type() == object.ERROR_OBJ
    }
    return false
}
```

**Usage**:
```go
result := Eval(node, env)
if isError(result) {
    return result  // Propagate error
}
```

---

### Error Propagation

Errors **bubble up** through the call stack:

```go
// Evaluate sub-expression
val := Eval(subExpr, env)
if isError(val) {
    return val  // Stop and return error
}

// Use val...
```

**Example**:
```bhasa
ফাংশন() {
    {
        {
            unknownVar  // Error here
        }
        // Never reached
    }
    // Never reached
}
// Error propagates all the way up
```

---

## Helper Functions

### Array Index Expression

```go
func evalArrayIndexExpression(array, index object.Object) object.Object {
    arrayObject := array.(*object.Array)
    idx := index.(*object.Integer).Value
    max := int64(len(arrayObject.Elements) - 1)
    
    if idx < 0 || idx > max {
        return NULL  // Out of bounds returns NULL
    }
    
    return arrayObject.Elements[idx]
}
```

**Out of bounds** → NULL (not error)

---

### Hash Index Expression

```go
func evalHashIndexExpression(hash, index object.Object) object.Object {
    hashObject := hash.(*object.Hash)
    
    // Check if index is hashable
    key, ok := index.(object.Hashable)
    if !ok {
        return newError("unusable as hash key: %s", index.Type())
    }
    
    // Lookup in hash
    pair, ok := hashObject.Pairs[key.HashKey()]
    if !ok {
        return NULL  // Not found returns NULL
    }
    
    return pair.Value
}
```

**Not found** → NULL (not error)

---

## Evaluation Traces

### Example 1: Simple Arithmetic

**Code**:
```bhasa
5 + 3 * 2
```

**Trace**:
```
Eval(InfixExpression +)
├─ Eval(IntegerLiteral 5)
│  └─ Return: &Integer{5}
├─ Eval(InfixExpression *)
│  ├─ Eval(IntegerLiteral 3)
│  │  └─ Return: &Integer{3}
│  ├─ Eval(IntegerLiteral 2)
│  │  └─ Return: &Integer{2}
│  └─ Return: &Integer{6}
└─ Return: &Integer{11}
```

---

### Example 2: Variable Access

**Code**:
```bhasa
ধরি x = 10;
x + 5;
```

**Trace**:
```
Eval(LetStatement)
├─ Eval(IntegerLiteral 10)
│  └─ Return: &Integer{10}
└─ env.Set("x", &Integer{10})

Eval(InfixExpression +)
├─ Eval(Identifier "x")
│  ├─ env.Get("x")
│  └─ Return: &Integer{10}
├─ Eval(IntegerLiteral 5)
│  └─ Return: &Integer{5}
└─ Return: &Integer{15}
```

---

### Example 3: Function Call

**Code**:
```bhasa
ধরি add = ফাংশন(a, b) { a + b };
add(5, 3);
```

**Trace**:
```
Eval(LetStatement)
├─ Eval(FunctionLiteral)
│  └─ Return: &Function{params: [a,b], env: Global}
└─ env.Set("add", &Function{...})

Eval(CallExpression)
├─ Eval(Identifier "add")
│  └─ Return: &Function{...}
├─ Eval(IntegerLiteral 5)
│  └─ Return: &Integer{5}
├─ Eval(IntegerLiteral 3)
│  └─ Return: &Integer{3}
└─ applyFunction(&Function{...}, [&Integer{5}, &Integer{3}])
   ├─ extendFunctionEnv
   │  └─ New env: {a: 5, b: 3, outer: Global}
   ├─ Eval(BlockStatement) in new env
   │  └─ Eval(InfixExpression +)
   │     ├─ Eval(Identifier "a") → &Integer{5}
   │     ├─ Eval(Identifier "b") → &Integer{3}
   │     └─ Return: &Integer{8}
   └─ Return: &Integer{8}
```

---

### Example 4: Closure

**Code**:
```bhasa
ধরি makeCounter = ফাংশন() {
    ধরি count = 0;
    ফাংশন() {
        count = count + 1;
        count
    }
};

ধরি counter = makeCounter();
counter();
counter();
```

**Trace**:
```
Call makeCounter():
├─ New env: {outer: Global}
├─ Eval(LetStatement) - count = 0
│  └─ env.Set("count", &Integer{0})
├─ Eval(FunctionLiteral) - inner function
│  └─ Return: &Function{env: {count: 0}}  ← Captures environment
└─ Return: &Function{env: {count: 0}}

Store in counter variable

Call counter() [1st time]:
├─ New env: {outer: {count: 0}}
├─ Eval(Assignment) - count = count + 1
│  ├─ Get count from outer: &Integer{0}
│  ├─ Add 1: &Integer{1}
│  └─ Set count in outer: &Integer{1}
├─ Eval(Identifier "count")
│  └─ Get from outer: &Integer{1}
└─ Return: &Integer{1}

Call counter() [2nd time]:
├─ New env: {outer: {count: 1}}  ← Same outer environment!
├─ Eval(Assignment) - count = count + 1
│  ├─ Get count from outer: &Integer{1}
│  ├─ Add 1: &Integer{2}
│  └─ Set count in outer: &Integer{2}
├─ Eval(Identifier "count")
│  └─ Get from outer: &Integer{2}
└─ Return: &Integer{2}
```

**Key Point**: Both calls share the **same** captured environment

---

## Performance Considerations

### Singleton Objects

```go
var (
    NULL  = &object.Null{}
    TRUE  = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
}
```

**Benefits**:
- Reduced allocations
- Fast equality (pointer comparison)
- Memory efficient

---

### Allocation Hotspots

**High allocation**:
- Integer objects (every number creates new object)
- String concatenation
- Array/Hash creation

**Optimization Ideas**:
- Object pooling for integers
- Rope data structure for strings
- Copy-on-write for collections

---

## See Also

- [Builtins Documentation](./builtins-documentation.md)
- [Evaluation Examples](./evaluation-examples.md)
- [Quick Reference](./quick-reference.md)
- [Object System](../../object/docs/)
- [AST Documentation](../../ast/docs/)

