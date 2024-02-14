package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, x, y, ...
	INT   = "INT"   // Integer

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	FUNC = "FUNC" // Function
	LET  = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNC,
	"let": LET,
}

func GetTokenType(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return IDENT
}
