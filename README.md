# Peace

Peace is an experimental programming-language project. The repository explores
a small `.peace` language, a Common Lisp-hosted interpreter, and the idea of
generating Go code for cryptographic operations.

The project is at an early prototype stage. The language example, Lisp
experiments, Go sample, and cryptography sketch are not yet one working
toolchain. In particular, Peace source code is not currently parsed and
compiled to Go end to end.

## Example

The intended starter syntax is shown in [`hello.peace`](./hello.peace):

```peace
set x = 10 : type integer
print x
print (+ x 5)
```

This example documents the proposed syntax; it is not currently a verified
executable program.

## Project Status

| Component | Status |
| --- | --- |
| Language design | Draft; see [`SPEC.md`](./SPEC.md) |
| Interpreter | Experimental sketch in `p.lisp`; not a complete implementation |
| Lexer and parser | Incomplete and not integrated with the interpreter |
| Go code generation | Planned; no working integrated compiler |
| Cryptography | Experimental sketches only; no supported Peace cryptographic API |
| Editor integration | No Peace-specific syntax highlighting or language extension |

There is currently no supported command for running a `.peace` program. The
Lisp files require a Common Lisp implementation, but the current experiments
should not be treated as a working interpreter or compiler.

## Repository Guide

- [`SPEC.md`](./SPEC.md) - Draft syntax and semantics, with implementation
  status called out explicitly.
- [`hello.peace`](./hello.peace) - Example of the proposed surface syntax.
- `p.lisp` - An early interpreter sketch.
- `lexer.lisp`, `parser.lisp`, `evaluator.lisp` - Incomplete language
  implementation experiments.
- `crypt.lisp` - An experimental Lisp-to-Go code-generation sketch.
- `peace.go` - A separate Go demonstration; it is not generated from
  `.peace` source.
- `tests/` - Early tests for the Common Lisp `setq-safe` macro, not tests of a
  complete Peace implementation.

## Development Direction

A useful first milestone is a small, testable pipeline for the syntax in
`hello.peace`:

1. Tokenize source text.
2. Parse tokens into a documented AST.
3. Evaluate the AST and test the behavior.
4. Only then add a Go backend that consumes the same AST.

Cryptographic features should be designed and tested separately from the
language parser. Do not rely on the current prototypes to protect real data.

## License

This project is licensed under the [MIT License](./LICENSE).
