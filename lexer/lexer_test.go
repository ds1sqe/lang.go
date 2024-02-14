package lexer

import (
	"mylang/token"
	"testing"
)

func TestTokenSimple(t *testing.T) {
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

func TestTokenAdvance(t *testing.T) {
	input := `
  let ten = 10;
  let two = 2;

  let add = fn(x,y) {
  x + y;
  };

  let result = add(ten,two);
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// let tow = 2;
		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		// let add = fn (x,y) {
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNC, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		// x + y;
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		// };
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = add(ten,two);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "ten"},
		{token.COMMA, ","},
		{token.IDENT, "two"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// EOF
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
