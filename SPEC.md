# Peace Language Specification

**Status:** Core specification for the current Go implementation  
**Version:** 0.2  
**Last updated:** 2026-11-09

This document defines the language core that is currently represented by the Go
implementation in [peace.go](peace.go), [lexer/lexer.go](lexer/lexer.go),
[parser/parser.go](parser/parser.go), and [evaluator.go](evaluator.go).

The older Lisp sketches and experiments have been archived under
[archive/lisp](archive/lisp) and are not the source of truth for the language
definition.

## 1. Goals

Peace is a small experimental language whose initial target is a minimal
integer language for:

- variable assignment
- integer arithmetic
- printing values
- basic runtime type validation

The implemented core intentionally excludes functions, control flow, strings,
scopes, concurrency, and cryptographic features.

## 2. Source Model

A Peace program is a sequence of newline-delimited statements.

- Blank lines are ignored.
- A line whose first non-whitespace character is `;` is ignored as a comment.
- Each statement is evaluated in source order.
- A program shares one global environment across statements.

The canonical example is:

```peace
set x = 10 : type integer
print x
print (+ x 5)
```

## 3. Lexical Conventions

The current Go lexer recognizes the following tokens:

- identifiers
- integer literals
- keywords: `set`, `print`, `type`, `if`, `else`, `let`
- operators: `+`, `-`, `*`, `/`, `=`, `(`, `)`, `{`, `}`, `[`, `]`, `,`, `.`, `:`
- comments beginning with `;` and continuing to the end of the line
- whitespace

Rules:

- Identifiers begin with an ASCII letter or `_`.
- Subsequent identifier characters may include ASCII letters, digits, or `_`.
- Integer literals are decimal digits only.
- Keywords are recognized when the identifier text matches a reserved keyword.
- Unknown characters produce an illegal token error.

## 4. Grammar

The implemented core grammar is:

```ebnf
program          = { statement } ;
statement        = set-statement | print-statement ;
set-statement    = "set", identifier, "=", expression,
                   [ ":", "type", "integer" ] ;
print-statement  = "print", expression ;
expression       = integer | identifier | addition-expression ;
addition-expression = "(", "+", expression, expression, ")" ;
identifier       = letter-or-underscore, { letter-or-underscore | digit } ;
integer          = [ "-" ], digit, { digit } ;
```

This is the grammar implemented by the current Go parser.

## 5. Runtime Semantics

### 5.1 Variables

- A variable is introduced by `set name = expression`.
- Variable names are case-sensitive.
- A variable must be previously bound before it is read.
- Reading an unbound variable is an error.

### 5.2 Assignment

The statement:

```peace
set x = 10
```

evaluates the right-hand expression and stores the result in the environment under
`x`.

### 5.3 Type Annotation

The annotation:

```peace
: type integer
```

is optional. If present, the assigned value must be an integer. If the runtime
value is not an integer, assignment fails with an error and the environment is
left unchanged.

The current implementation recognizes only `integer` as a valid type name.

### 5.4 Printing

The statement:

```peace
print expr
```

evaluates `expr` and writes its value to standard output followed by a newline.

### 5.5 Expressions

The currently implemented expression forms are:

- integer literal
- variable reference
- binary arithmetic: `(+ left right)`, `(- left right)`, `(* left right)`, `(/ left right)`

Examples:

```peace
10
x
(+ x 5)
(- x 3)
(* x 2)
(/ x 4)
```

### 5.6 Reassignment

A variable may be assigned more than once in the same environment. Reassignment
updates the existing binding rather than creating a second scope.

Example:

```peace
set x = 10
set x = 20
print x
```

This evaluates to `20`.

### 5.7 Evaluation Order

Evaluation occurs from left to right within each binary expression, and
statements execute in source order.

## 6. Errors

The implementation must fail fast on the following conditions:

- malformed syntax
- unknown statement
- missing variable
- unsupported expression form
- invalid type annotation
- illegal token

The error message format is not yet standardized beyond the current Go runtime,
but the error must be reported to the caller and must not silently continue with
an invalid program.

## 7. Non-Goals for the Current Core

The following features are intentionally out of scope for the current language
core and are not part of the implemented semantics:

- local scopes
- function definitions
- strings
- booleans
- comparison operators
- loops
- imports
- modules
- cryptographic primitives
- editor integration
- automatic code generation to Go

These features may be specified in later versions after the core language is
stable.

## 8. Implementation Status

| Area | Status |
| --- | --- |
| Lexical analysis | Implemented in Go |
| Parsing | Implemented for the current core grammar |
| Evaluation | Implemented for integers, variables, and addition |
| Type annotation | Implemented for `: type integer` |
| Program execution | Implemented through the Go interpreter |
| Lisp experiments | Experimental only |
| Cryptography | Not implemented |
| Editor integration | Not implemented |

This specification reflects the language as currently implemented by the Go
source files and should be treated as the authoritative definition for the
initial language core.
