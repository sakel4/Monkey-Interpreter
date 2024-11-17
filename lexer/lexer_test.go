package lexer

import (
	"testing"

	"github.com/sakel4/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		currentToken := l.NextToken()

		if currentToken.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q", i, tt.expectedType, currentToken.Type)
		}

		if currentToken.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%d", i, tt.expectedLiteral, currentToken.Literal)
		}
	}
}
