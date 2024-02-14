package lexer

import "mylang/token"

type LexerChar byte

type Lexer struct {
	input   string
	pos     int       // current position
	nextPos int       // next position
	ch      LexerChar // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = LexerChar(l.input[l.nextPos])
	}
	l.pos = l.nextPos
	l.nextPos += 1
}

func newToken(tokenType token.TokenType, ch LexerChar) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch LexerChar) bool {
	return (('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_')
}
func isDigit(ch LexerChar) bool {
	return '0' <= ch && ch <= '9'
}

// read identifier and return it
func (l *Lexer) readId() string {
	pos := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) readNum() string {
	pos := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readId()
			tok.Type = token.GetTokenType(tok.Literal)
			return tok

		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNum()
			return tok

		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
