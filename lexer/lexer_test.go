package lexer

import (
	"mylang/token"
	"testing"
)

func TestTokenSimple(t *testing.T) {
	input := `=+(){},;[]`

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
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
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

		// let two = 2;
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

func TestTokenAdvance2(t *testing.T) {
	input := `
  let ten = 10;
  let two = 2;
  let flag = false;

  let test = fn(x,y) {
  if (x>y) {
  return x*y;
  }else if (x<y) {
  return x/y;
  }else {
  flag = true;
  return x%y;
  }

  };

  let result = test(ten,two);
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

		// let two = 2;
		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		// let flag = false;
		{token.LET, "let"},
		{token.IDENT, "flag"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		// let test = fn (x,y) {
		{token.LET, "let"},
		{token.IDENT, "test"},
		{token.ASSIGN, "="},
		{token.FUNC, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		// if (x > y) {
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		// return x*y;
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PROD, "*"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		// } else if (x<y) {
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.LT, "<"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},

		// return x/y;
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.DIV, "/"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		// } else {
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},

		// flag = true;
		{token.IDENT, "flag"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		// return x%y;
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.MOD, "%"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},

		// }
		{token.RBRACE, "}"},
		// };
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = test(ten,two);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "test"},
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

func TestTokenAdvance3(t *testing.T) {
	input := `
  let ten = 10;
  let two = 2;

  ten <= two;
  ten >= two;
  ten == two;
  ten != two;

  ten & two;
  ten | two;

  true && false;
  true || false;
  "foobar"
  "foo bar"
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

		// let two = 2;
		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.LT_OR_EQ, "<="},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.GT_OR_EQ, ">="},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.EQ, "=="},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.NOT_EQ, "!="},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.BIT_AND, "&"},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.IDENT, "ten"},
		{token.BIT_OR, "|"},
		{token.IDENT, "two"},
		{token.SEMICOLON, ";"},

		{token.TRUE, "true"},
		{token.AND, "&&"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		{token.TRUE, "true"},
		{token.OR, "||"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

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
