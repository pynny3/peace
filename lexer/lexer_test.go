package lexer

import "testing"

func TestNextTokenScansPeaceStatements(t *testing.T) {
	lexer := NewLexer("set item2 = 10 : type integer\nprint (+ item2 5)")
	want := []Token{
		{Type: TokenKeyword, Value: "set", Line: 1, Column: 1},
		{Type: TokenIdentifier, Value: "item2", Line: 1, Column: 5},
		{Type: TokenOperator, Value: "=", Line: 1, Column: 11},
		{Type: TokenNumber, Value: "10", Line: 1, Column: 13},
		{Type: TokenOperator, Value: ":", Line: 1, Column: 16},
		{Type: TokenKeyword, Value: "type", Line: 1, Column: 18},
		{Type: TokenIdentifier, Value: "integer", Line: 1, Column: 23},
		{Type: TokenKeyword, Value: "print", Line: 2, Column: 1},
		{Type: TokenOperator, Value: "(", Line: 2, Column: 7},
		{Type: TokenOperator, Value: "+", Line: 2, Column: 8},
		{Type: TokenIdentifier, Value: "item2", Line: 2, Column: 10},
		{Type: TokenNumber, Value: "5", Line: 2, Column: 16},
		{Type: TokenOperator, Value: ")", Line: 2, Column: 17},
		{Type: TokenEOF, Line: 2, Column: 18},
	}

	for index, expected := range want {
		if got := lexer.NextToken(); got != expected {
			t.Errorf("token %d = %#v, want %#v", index, got, expected)
		}
	}
}

func TestNextTokenScansCommentsAndIllegalCharacters(t *testing.T) {
	lexer := NewLexer("; note\n@")
	want := []Token{
		{Type: TokenComment, Value: "; note", Line: 1, Column: 1},
		{Type: TokenIllegal, Value: "@", Line: 2, Column: 1},
		{Type: TokenEOF, Line: 2, Column: 2},
	}

	for index, expected := range want {
		if got := lexer.NextToken(); got != expected {
			t.Errorf("token %d = %#v, want %#v", index, got, expected)
		}
	}
}
