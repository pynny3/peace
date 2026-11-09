// Package lexer converts Peace source text into tokens.
package lexer

// TokenType identifies the kind of a token.
type TokenType int

const (
	TokenIllegal TokenType = iota
	TokenEOF
	TokenIdentifier
	TokenNumber
	TokenString
	TokenOperator
	TokenKeyword
	TokenComment
)

// Token is a lexical unit and its starting position in the source.
type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

// Lexer scans source text and returns one token at a time.
type Lexer struct {
	input    string
	position int
	line     int
	column   int
}

var keywords = map[string]bool{
	"else":  true,
	"if":    true,
	"let":   true,
	"print": true,
	"set":   true,
	"type":  true,
}

// NewLexer creates a lexer for input.
func NewLexer(input string) *Lexer {
	return &Lexer{input: input, line: 1, column: 1}
}

// NextToken returns the next token, including TokenEOF at the end of input.
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	line, column := l.line, l.column

	if l.position >= len(l.input) {
		return Token{Type: TokenEOF, Line: line, Column: column}
	}

	ch := l.peek()
	if isIdentifierStart(ch) {
		value := l.readIdentifier()
		tokenType := TokenIdentifier
		if keywords[value] {
			tokenType = TokenKeyword
		}
		return Token{Type: tokenType, Value: value, Line: line, Column: column}
	}

	if isDigit(ch) {
		return Token{
			Type:   TokenNumber,
			Value:  l.readNumber(),
			Line:   line,
			Column: column,
		}
	}

	if ch == ';' {
		start := l.position
		for l.position < len(l.input) && l.peek() != '\n' {
			l.advance()
		}
		return Token{
			Type:   TokenComment,
			Value:  l.input[start:l.position],
			Line:   line,
			Column: column,
		}
	}

	l.advance()
	switch ch {
	case '+', '-', '*', '/', '=', '(', ')', '{', '}', '[', ']', ',', '.', ':':
		return Token{Type: TokenOperator, Value: string(ch), Line: line, Column: column}
	default:
		return Token{Type: TokenIllegal, Value: string(ch), Line: line, Column: column}
	}
}

func (l *Lexer) skipWhitespace() {
	for l.position < len(l.input) {
		switch l.peek() {
		case ' ', '\t', '\r', '\n':
			l.advance()
		default:
			return
		}
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for l.position < len(l.input) && isIdentifierPart(l.peek()) {
		l.advance()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readNumber() string {
	start := l.position
	for l.position < len(l.input) && isDigit(l.peek()) {
		l.advance()
	}
	return l.input[start:l.position]
}

func (l *Lexer) peek() byte {
	return l.input[l.position]
}

func (l *Lexer) advance() byte {
	ch := l.input[l.position]
	l.position++
	if ch == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
	return ch
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
