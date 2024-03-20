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

func (l *Lexer) peekChar() LexerChar {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return LexerChar(l.input[l.nextPos])
	}
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
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.PROD, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	case '%':
		tok = newToken(token.MOD, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_OR_EQ, Literal: literal}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_OR_EQ, Literal: literal}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.AND, Literal: literal}
		} else {
			tok = newToken(token.BIT_AND, l.ch)
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.OR, Literal: literal}
		} else {
			tok = newToken(token.BIT_OR, l.ch)
		}
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
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

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

func (l *Lexer) readString() string {
	position := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.pos]
}
