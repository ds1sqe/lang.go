package parser

import (
	"mylang/ast"
	"mylang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
  let x = 5;
  let y = 10;
  let foo = 31324;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statement. got =%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"foo"},
	}

	for i, test := range tests {
		stm := program.Statements[i]
		if !testLetStatements(t, stm, test.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral is not 'let', got = %q", s.TokenLiteral())
		return false
	}
	letStm, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got = %T", s)
		return false
	}

	if letStm.Name.Value != name {
		t.Errorf("letStm.Name.Value != '%s' , got = %s ",
			name, letStm.Name.Value)
		return false
	}

	if letStm.Name.TokenLiteral() != name {
		t.Errorf("letStm.Name.TokenLiteral() != '%s' , got = %s ",
			name, letStm.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))

	for _, error := range errors {
		t.Errorf("parser error: %q", error)
	}
	t.FailNow()
}
