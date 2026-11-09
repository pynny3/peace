# Peace

Peace is an experimental programming-language project. The repository explores
a small `.peace` language, a Common Lisp-hosted interpreter, and the idea of
generating Go code for cryptographic operations.

The project is at an early prototype stage. A small Go interpreter supports
the initial syntax in `hello.peace`; the Lisp experiments, Go code generator,
and cryptography sketch are not integrated into a single toolchain.

## Example

The intended starter syntax is shown in [`hello.peace`](./hello.peace):

```peace
set x = 10 : type integer
print x
print (+ x 5)
print (- x 3)
print (* x 2)
print (/ x 4)
```

Run the example with Go:

```sh
go run peace.go hello.peace
```

The interpreter also starts an interactive REPL when run without a file:

```sh
go run peace.go
```

## Project Status

| Component | Status |
| --- | --- |
| Language design | Draft; see [`SPEC.md`](./SPEC.md) |
| Go interpreter | Supports integer values, variables, re-assignment, `+ - * /`, printing, and `: type integer` |
| Lisp interpreter | Experimental sketch in `p.lisp`; not a complete implementation |
| Go lexer and parser | Tokenize and parse the core arithmetic syntax into an AST; the interpreter uses the same expression model |
| Go code generation | Planned; no working integrated compiler |
| Cryptography | Experimental sketches only; no supported Peace cryptographic API |
| Editor integration | No Peace-specific syntax highlighting or language extension |

The Go interpreter can run the initial `.peace` syntax shown above. The Lisp
files require a Common Lisp implementation and remain experimental.

## Repository Guide

- [`SPEC.md`](./SPEC.md) - Current core language specification for the Go
  implementation.
- [`hello.peace`](./hello.peace) - Example of the supported surface syntax.
- `archive/lisp/` - Archived Lisp experiments and design sketches; not part of
  the active language implementation.
- `lexer/` and `parser/` - Go packages for tokenizing and parsing the core
  syntax used by the interpreter.
- `peace.go` - The active Go interpreter for the initial Peace language core.
- `evaluator.go` - AST evaluator used by the language core.
- `tests/` - Current Go regression tests for the implemented core behavior.

## Development Direction

A useful next milestone is to connect the existing Go lexer and parser to a
testable evaluator for the syntax in `hello.peace`:

1. Evaluate the AST produced by the parser and test the behavior.
2. Keep the parser AST and interpreter semantics aligned.
3. Only then add a Go backend that consumes the same AST.

Cryptographic features should be designed and tested separately from the
language parser. Do not rely on the current prototypes to protect real data.

## License

This project is licensed under the [MIT License](./LICENSE).
