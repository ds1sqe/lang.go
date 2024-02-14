package lexer

import (
	"mylang/token"
	"testing"
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

	for i, tst := range tests {
		tok := l.NextToken()
		if tok.Type != tst.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected:{%q}, got:{%q}",
				i, tst.expectedType, tok.Type)
		}
		if tok.Literal != tst.expectedLiteral {
			t.Fatalf("tests[%d] - token literal wrong. expected:{%q}, got:{%q}",
				i, tst.expectedType, tok.Literal)
		}
	}
}
