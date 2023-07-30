package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	ADD      = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	//KEYWORDS
	LET = "LET"
	FUNCTION = "FUNCTION"

)

var keywords = map[string] TokenType{
	"let": LET,
	"fn": FUNCTION,
}

func LookupIdent(ident string) TokenType {
	
}
