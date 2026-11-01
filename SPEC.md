# Peace Language Specification

**Status:** Draft  
**Version:** 0.1  
**Last updated:** 2026-10-31

This document describes the intended small core of the Peace language. It is
not a claim that the current repository implements the specification. The
interpreter, lexer, parser, and code generator are experimental and are not yet
connected into a working toolchain.

## 1. Goals

Peace is an experimental language project with two intended ways to use a
program:

1. Interpret it with a Peace interpreter.
2. Eventually compile supported programs to Go.

The initial language core is deliberately small: integer values, variables,
addition, optional integer type annotations, and printing. Cryptographic
operations and Go code generation are future work, not part of the implemented
language core.

## 2. Source Files

Peace source files use the `.peace` extension. A program consists of
statements, one per line. Blank lines and lines whose first non-whitespace
character is `;` are ignored.

The current example is:

```peace
set x = 10 : type integer
print x
print (+ x 5)
```

## 3. Lexical Conventions

- Identifiers begin with an ASCII letter or underscore.
- Subsequent identifier characters may also include decimal digits.
- Keywords and identifiers are written in lowercase in this draft.
- Whitespace separates tokens.
- A semicolon begins a full-line comment.

Inline comments, escape sequences, Unicode identifiers, and multiline
expressions are not specified yet.

## 4. Syntax

The following EBNF describes the proposed initial syntax:

```ebnf
program       = { statement, newline } ;
statement     = set-statement | print-statement ;
set-statement = "set", identifier, "=", expression,
                [":", "type", "integer"] ;
print-statement = "print", expression ;
expression    = integer | identifier | "(", "+", expression, expression, ")" ;
identifier    = letter-or-underscore, { letter-or-underscore | digit } ;
integer       = [ "-" ], digit, { digit } ;
```

The grammar is a draft. In particular, the repository does not yet contain a
working lexer and parser that enforce it.

## 5. Core Semantics

These are the intended semantics for the initial core:

- `set name = expression` evaluates the expression and binds its value to
  `name`.
- `: type integer` requests a runtime check that the assigned value is an
  integer. A failed check is an error and must not silently assign the value.
- `print expression` evaluates the expression and writes its value followed by
  a newline.
- `(+ left right)` evaluates both operands as integers and returns their sum.
- Reading an unbound variable is an error.
- Statements execute in source order in a single program-level environment.

Reassignment, local scopes, functions, control flow, other operators, and
additional types are not specified in version 0.1.

## 6. Errors

A conforming implementation should report an error for malformed syntax,
unknown statements, reads of unbound variables, and failed type checks. Error
wording and source locations are not standardized yet.

## 7. Go Backend and Cryptography

Generating Go code from Peace programs is a project goal, but no working
Peace-to-Go compiler is currently available. The experimental `crypt.lisp`
file is not integrated with a lexer, parser, or Go backend.

If cryptographic support is added, it should use established cryptographic
libraries in the target language rather than implementing cryptographic
primitives in Peace. A future design must specify key handling, nonce handling,
authentication, and error behavior before claiming to provide secure
encryption.

## 8. Implementation Status

| Area | Status |
| --- | --- |
| Example `.peace` syntax | Present as a small example |
| Interpreter sketch | Present in `p.lisp`; not a complete or safe language implementation |
| Lexer and parser | Experimental, incomplete, and not connected |
| AST evaluator | Incomplete and not connected |
| Go code generation | No working integrated backend |
| Cryptographic operations | No working Peace cryptographic feature |
| Editor support for `.peace` | No language extension or syntax definition |

This table describes the repository as of the date above. It should be updated
when implementation behavior changes.
