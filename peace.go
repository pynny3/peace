package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"strings"
)

// SetqSafeValue represents a type-safe variable with optional encryption
type SetqSafeValue struct {
	Name  string
	Value interface{}
	Type  string
	Doc   string
	Debug bool
}

// SetqSafe creates a type-safe variable
func SetqSafe(name string, value interface{}, typ string, doc string, debug bool) SetqSafeValue {
	if debug {
		fmt.Printf("[debug] setting %s to %v\n", name, value)
	}
	return SetqSafeValue{
		Name:  name,
		Value: value,
		Type:  typ,
		Doc:   doc,
		Debug: debug,
	}
}

// Top returns the maximum of two values
func Top(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Bottom returns the minimum of two values
func Bottom(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Middle returns the median of three values
func Middle(a, b, c int) int {
	if (a >= b && a <= c) || (a >= c && a <= b) {
		return a
	}
	if (b >= a && b <= c) || (b >= c && b <= a) {
		return b
	}
	return c
}

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply multiplies two integers
func Multiply(a, b int) int {
	return a * b
}

// Demo demonstrates the original Go helper functions
func Demo() {
	fmt.Println("=== Peace Language Demo (Go Implementation) ===")

	x := SetqSafe("x", 10, "integer", "test variable", true)
	fmt.Printf("x = %v\n", x.Value)
	fmt.Printf("x (doc): %s\n", x.Doc)

	fmt.Println("\n=== Math Operations ===")
	fmt.Printf("Top(15, 8) = %d\n", Top(15, 8))
	fmt.Printf("Bottom(15, 8) = %d\n", Bottom(15, 8))
	fmt.Printf("Middle(3, 8, 5) = %d\n", Middle(3, 8, 5))
	fmt.Printf("Add(5, 3) = %d\n", Add(5, 3))
	fmt.Printf("Subtract(10, 3) = %d\n", Subtract(10, 3))
	fmt.Printf("Multiply(4, 5) = %d\n", Multiply(4, 5))

	fmt.Println("\n=== Type Safety ===")
	y := SetqSafe("y", 42, "integer", "another test variable", true)
	fmt.Printf("y = %v (type: %s, doc: %s)\n", y.Value, y.Type, y.Doc)

	fmt.Println("\n=== Demo Complete ===")
}

// Interpreter evaluates Peace statements in a persistent environment.
type Interpreter struct {
	environment map[string]*big.Int
}

// NewInterpreter creates an interpreter with an empty environment.
func NewInterpreter() *Interpreter {
	return &Interpreter{environment: make(map[string]*big.Int)}
}

// RunLine evaluates one Peace statement and writes print results to out.
func (i *Interpreter) RunLine(line string, out io.Writer) error {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, ";") {
		return nil
	}

	if rest, ok := statementRemainder(line, "print"); ok {
		if strings.TrimSpace(rest) == "" {
			return errors.New("Peace: expected expression after `print`")
		}
		value, err := i.evaluateExpression(rest)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, value.String())
		return err
	}

	if rest, ok := statementRemainder(line, "set"); ok {
		return i.setVariable(rest)
	}

	return fmt.Errorf("Peace: unknown statement: %s", line)
}

func statementRemainder(line, keyword string) (string, bool) {
	if !strings.HasPrefix(line, keyword) || len(line) == len(keyword) {
		return "", false
	}
	if line[len(keyword)] != ' ' && line[len(keyword)] != '\t' {
		return "", false
	}
	return strings.TrimSpace(line[len(keyword):]), true
}

func (i *Interpreter) setVariable(statement string) error {
	equal := strings.IndexByte(statement, '=')
	if equal < 0 {
		return errors.New("Peace: expected `set name = expression`")
	}

	name := strings.TrimSpace(statement[:equal])
	if !validIdentifier(name) {
		return fmt.Errorf("Peace: invalid variable name: %s", name)
	}

	expression := strings.TrimSpace(statement[equal+1:])
	if annotation := strings.LastIndex(expression, ": type "); annotation >= 0 {
		typeName := strings.TrimSpace(expression[annotation+len(": type "):])
		if typeName != "integer" {
			return fmt.Errorf("Peace: unsupported type: %s", typeName)
		}
		expression = strings.TrimSpace(expression[:annotation])
	}
	if expression == "" {
		return errors.New("Peace: expected expression in assignment")
	}

	value, err := i.evaluateExpression(expression)
	if err != nil {
		return err
	}

	i.environment[name] = new(big.Int).Set(value)
	return nil
}

func (i *Interpreter) evaluateExpression(text string) (*big.Int, error) {
	parser := expressionParser{input: text, environment: i.environment}
	value, err := parser.parse()
	if err != nil {
		return nil, fmt.Errorf("Peace: invalid expression: %s: %w", text, err)
	}
	parser.skipWhitespace()
	if parser.position != len(parser.input) {
		return nil, fmt.Errorf("Peace: invalid expression: unexpected input %q", parser.input[parser.position:])
	}
	return value, nil
}

type expressionParser struct {
	input       string
	position    int
	environment map[string]*big.Int
}

func (p *expressionParser) parse() (*big.Int, error) {
	p.skipWhitespace()
	if p.position >= len(p.input) {
		return nil, errors.New("expected expression")
	}

	switch ch := p.input[p.position]; {
	case ch == '(':
		p.position++
		p.skipWhitespace()
		if p.position >= len(p.input) {
			return nil, errors.New("expected binary operator")
		}

		op := p.input[p.position]
		if !isBinaryOperator(op) {
			return nil, fmt.Errorf("unsupported operator: %q", op)
		}
		p.position++

		left, err := p.parse()
		if err != nil {
			return nil, err
		}
		right, err := p.parse()
		if err != nil {
			return nil, err
		}
		p.skipWhitespace()
		if p.position >= len(p.input) || p.input[p.position] != ')' {
			return nil, errors.New("expected `)`")
		}
		p.position++

		switch op {
		case '+':
			return new(big.Int).Add(left, right), nil
		case '-':
			return new(big.Int).Sub(left, right), nil
		case '*':
			return new(big.Int).Mul(left, right), nil
		case '/':
			if right.Sign() == 0 {
				return nil, errors.New("division by zero")
			}
			return new(big.Int).Quo(left, right), nil
		default:
			return nil, fmt.Errorf("unsupported operator: %q", op)
		}
	case ch == '-' || isDigit(ch):
		start := p.position
		if ch == '-' {
			p.position++
			if p.position >= len(p.input) || !isDigit(p.input[p.position]) {
				return nil, errors.New("expected digits after `-`")
			}
		}
		for p.position < len(p.input) && isDigit(p.input[p.position]) {
			p.position++
		}
		value, ok := new(big.Int).SetString(p.input[start:p.position], 10)
		if !ok {
			return nil, errors.New("invalid integer")
		}
		return value, nil
	case isIdentifierStart(ch):
		start := p.position
		p.position++
		for p.position < len(p.input) && isIdentifierPart(p.input[p.position]) {
			p.position++
		}
		name := p.input[start:p.position]
		value, ok := p.environment[name]
		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", name)
		}
		return new(big.Int).Set(value), nil
	default:
		return nil, fmt.Errorf("unexpected character %q", ch)
	}
}

func (p *expressionParser) skipWhitespace() {
	for p.position < len(p.input) {
		switch p.input[p.position] {
		case ' ', '\t', '\r', '\n':
			p.position++
		default:
			return
		}
	}
}

func validIdentifier(name string) bool {
	if name == "" || !isIdentifierStart(name[0]) {
		return false
	}
	for position := 1; position < len(name); position++ {
		if !isIdentifierPart(name[position]) {
			return false
		}
	}
	return true
}

func isIdentifierStart(ch byte) bool {
	return ch == '_' || 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isIdentifierPart(ch byte) bool {
	return isIdentifierStart(ch) || isDigit(ch)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isBinaryOperator(ch byte) bool {
	switch ch {
	case '+', '-', '*', '/':
		return true
	default:
		return false
	}
}

// RunFile executes a Peace source file, preserving the environment between lines.
func (i *Interpreter) RunFile(path string, out io.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		if err := i.RunLine(scanner.Text(), out); err != nil {
			return fmt.Errorf("%s:%d: %w", path, lineNumber, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}

// RunREPL reads and evaluates statements until EOF or the :quit command.
func (i *Interpreter) RunREPL(in io.Reader, out, errOut io.Writer) error {
	fmt.Fprintln(out, "Peace REPL (:quit end)")
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, "peace> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == ":quit" {
			break
		}
		if err := i.RunLine(line, out); err != nil {
			fmt.Fprintf(errOut, "Error: %v\n", err)
		}
	}
	return scanner.Err()
}

func main() {
	interpreter := NewInterpreter()
	switch len(os.Args) {
	case 1:
		if err := interpreter.RunREPL(os.Stdin, os.Stdout, os.Stderr); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	case 2:
		if err := interpreter.RunFile(os.Args[1], os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Usage: %s [file.peace]\n", os.Args[0])
		os.Exit(2)
	}
}
