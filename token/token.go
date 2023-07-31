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

	GRTHAN = ">"
	LSTHAN = "<"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	EXCLAIM = "!"

	//DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"


	//DOUBLE CHARS
	EQ = "=="
	NOTEQ = "!="
	
	//KEYWORDS
	LET = "LET"
	FUNCTION = "FUNCTION"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"

)

var keywords = map[string] TokenType{
	"let": LET,
	"fn": FUNCTION,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
}

func LookupIdent(ident string) TokenType { //this is to get the token (TokenType)
	if tok, ok := keywords[ident]; ok { //token, 
		return tok
	}
	return IDENT
}
