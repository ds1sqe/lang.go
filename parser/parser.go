package parser

import (
	"fmt"
	"mylang/ast"
	"mylang/lexer"
	"mylang/token"
)

type Error string

type Parser struct {
	lex       *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []Error
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{lex: l, errors: []Error{}}

	p.nextToken() // setup curToken
	p.nextToken() //       and peekToken

	return p
}

func (p *Parser) Errors() []Error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stm := p.parseStatement()
		if stm != nil {
			program.Statements = append(program.Statements, stm)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stm := &ast.LetStatement{Token: p.curToken}

	// if next token is not 'IDENT'
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stm.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// if next token is not 'ASSIGN' aka '='
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stm
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, bot got %s instead.", t,
		p.peekToken.Type)
	p.errors = append(p.errors, Error(msg))
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
