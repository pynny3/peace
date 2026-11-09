// Package parser parses Peace tokens into an abstract syntax tree.
package parser

import (
	"fmt"

	"peace/lexer"
)

// ExprKind identifies an expression node.
type ExprKind uint8

const (
	ExprNumber ExprKind = iota
	ExprVariable
	ExprAdd
)

// Expr is an integer, variable reference, or addition expression.
type Expr struct {
	Kind  ExprKind
	Value string
	Left  *Expr
	Right *Expr
}

// StatementKind identifies a statement node.
type StatementKind uint8

const (
	StatementSet StatementKind = iota
	StatementPrint
)

// Statement is a variable assignment or print statement.
type Statement struct {
	Kind     StatementKind
	Name     string
	Expr     *Expr
	TypeName string
}

// Parser parses a sequence of lexer tokens.
type Parser struct {
	tokens []lexer.Token
	pos    int
}

// NewParser creates a parser for tokens.
func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

// Parse parses the token stream and returns its statements.
func (p *Parser) Parse() ([]Statement, error) {
	var statements []Statement
	for token := p.current(); token.Type != lexer.TokenEOF; token = p.current() {
		statement, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, statement)
	}
	return statements, nil
}

func (p *Parser) parseStatement() (Statement, error) {
	token := p.current()
	if token.Type != lexer.TokenKeyword {
		return Statement{}, p.syntaxError(token, "set or print")
	}

	switch token.Value {
	case "set":
		return p.parseSet()
	case "print":
		return p.parsePrint()
	default:
		return Statement{}, p.syntaxError(token, "supported statement (set or print)")
	}
}

func (p *Parser) parseSet() (Statement, error) {
	p.advance() // set

	name, err := p.expect(lexer.TokenIdentifier, "")
	if err != nil {
		return Statement{}, err
	}
	if _, err := p.expect(lexer.TokenOperator, "="); err != nil {
		return Statement{}, err
	}

	expression, err := p.parseExpression()
	if err != nil {
		return Statement{}, err
	}

	statement := Statement{
		Kind: StatementSet,
		Name: name.Value,
		Expr: expression,
	}
	if token := p.current(); token.Type == lexer.TokenOperator && token.Value == ":" {
		p.advance()
		if _, err := p.expect(lexer.TokenKeyword, "type"); err != nil {
			return Statement{}, err
		}
		typeName, err := p.expect(lexer.TokenIdentifier, "")
		if err != nil {
			return Statement{}, err
		}
		if typeName.Value != "integer" {
			return Statement{}, p.syntaxError(typeName, "integer")
		}
		statement.TypeName = typeName.Value
	}

	return statement, nil
}

func (p *Parser) parsePrint() (Statement, error) {
	p.advance() // print

	expression, err := p.parseExpression()
	if err != nil {
		return Statement{}, err
	}
	return Statement{Kind: StatementPrint, Expr: expression}, nil
}

func (p *Parser) parseExpression() (*Expr, error) {
	token := p.current()
	switch token.Type {
	case lexer.TokenNumber:
		p.advance()
		return &Expr{Kind: ExprNumber, Value: token.Value}, nil
	case lexer.TokenIdentifier:
		p.advance()
		return &Expr{Kind: ExprVariable, Value: token.Value}, nil
	case lexer.TokenOperator:
		if token.Value == "-" {
			p.advance()
			number, err := p.expect(lexer.TokenNumber, "")
			if err != nil {
				return nil, err
			}
			return &Expr{Kind: ExprNumber, Value: "-" + number.Value}, nil
		}
		if token.Value != "(" {
			return nil, p.syntaxError(token, "integer, identifier, or addition expression")
		}
		p.advance() // (

		if _, err := p.expect(lexer.TokenOperator, "+"); err != nil {
			return nil, err
		}
		left, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		right, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		if _, err := p.expect(lexer.TokenOperator, ")"); err != nil {
			return nil, err
		}
		return &Expr{Kind: ExprAdd, Left: left, Right: right}, nil
	default:
		return nil, p.syntaxError(token, "integer, identifier, or addition expression")
	}
}

func (p *Parser) expect(tokenType lexer.TokenType, value string) (lexer.Token, error) {
	token := p.current()
	if token.Type != tokenType || (value != "" && token.Value != value) {
		expected := value
		if expected == "" {
			expected = fmt.Sprintf("token type %d", tokenType)
		}
		return lexer.Token{}, p.syntaxError(token, expected)
	}
	p.advance()
	return token, nil
}

func (p *Parser) current() lexer.Token {
	for p.pos < len(p.tokens) && p.tokens[p.pos].Type == lexer.TokenComment {
		p.pos++
	}
	if p.pos >= len(p.tokens) {
		return lexer.Token{Type: lexer.TokenEOF, Line: 1, Column: 1}
	}
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	if p.pos < len(p.tokens) {
		p.pos++
	}
}

func (p *Parser) syntaxError(token lexer.Token, expected string) error {
	return fmt.Errorf(
		"%d:%d: expected %s, got %q",
		token.Line,
		token.Column,
		expected,
		token.Value,
	)
}
