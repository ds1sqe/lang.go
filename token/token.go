package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"  // add, x, y, ...
	INT    = "INT"    // Integer
	STRING = "STRING" // String

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	PROD   = "*" // Product
	DIV    = "/" // Divide
	MOD    = "%" // Modular
	BANG   = "!"

	// Compare
	LT       = "<"
	LT_OR_EQ = "<="
	GT       = ">"
	GT_OR_EQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// BIT And, Or
	BIT_AND = "&"
	BIT_OR  = "|"

	// And , Or
	AND = "&&"
	OR  = "||"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	LBRACKET = "["
	RBRACKET = "]"

	// keywords
	FUNC   = "FUNC" // Function
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func GetTokenType(word string) TokenType {
	if tok, ok := keywords[word]; ok {
		return tok
	}
	return IDENT
}
