package parser

import (
	"reflect"
	"testing"

	"peace/lexer"
)

func parseSource(source string) ([]Statement, error) {
	scanner := lexer.NewLexer(source)
	var tokens []lexer.Token
	for {
		token := scanner.NextToken()
		tokens = append(tokens, token)
		if token.Type == lexer.TokenEOF {
			break
		}
	}
	return NewParser(tokens).Parse()
}

func TestParseSetPrintAndAddition(t *testing.T) {
	statements, err := parseSource("set x = 10 : type integer\nprint x\nprint (+ x 5)")
	if err != nil {
		t.Fatal(err)
	}

	want := []Statement{
		{
			Kind:     StatementSet,
			Name:     "x",
			TypeName: "integer",
			Expr:     &Expr{Kind: ExprNumber, Value: "10"},
		},
		{
			Kind: StatementPrint,
			Expr: &Expr{Kind: ExprVariable, Value: "x"},
		},
		{
			Kind: StatementPrint,
			Expr: &Expr{
				Kind:  ExprAdd,
				Left:  &Expr{Kind: ExprVariable, Value: "x"},
				Right: &Expr{Kind: ExprNumber, Value: "5"},
			},
		},
	}
	if !reflect.DeepEqual(statements, want) {
		t.Fatalf("Parse() = %#v, want %#v", statements, want)
	}
}

func TestParseNegativeNumberAndComment(t *testing.T) {
	statements, err := parseSource("; comment\nset x = -2\nprint x")
	if err != nil {
		t.Fatal(err)
	}
	if got, want := statements[0].Expr.Value, "-2"; got != want {
		t.Fatalf("negative integer = %q, want %q", got, want)
	}
	if len(statements) != 2 {
		t.Fatalf("parsed %d statements, want 2", len(statements))
	}
}

func TestParseRejectsUnsupportedSyntax(t *testing.T) {
	for _, source := range []string{
		"let x = 1",
		"print (+ 1)",
		"set x 1",
		"set x = 1 : type string",
	} {
		if _, err := parseSource(source); err == nil {
			t.Errorf("Parse(%q) succeeded, want syntax error", source)
		}
	}
}
