package lexer

import "monkey/token"

type Lexer struct {
	input string
	position int //current position
	readPosition int //next one
	ch byte
}

func New (input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	//if readPosition >= string range then equal ch to 0
	//else then ch = string[position]
	//position = readPosition
	//readPosition++

	if l.readPosition >= len(l.input) {
		l.ch = 0 //signals NUL, “we haven’t read anything yet” or “end of file”
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

//make nextToken and newToken
func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	switch (l.ch) {//{}()+=;
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case '+':
			tok = newToken(token.ADD, l.ch)
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			//if it is a character then run readIdent
			if isLetter(l.ch) {
				tok.Literal = l.readIdent()
				tok.Type = token.IDENT //TODO: look for it
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tok

}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool{
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func (l *Lexer) readIdent() string {
	var tok token.Token
	switch (l.input) {
		case "let":
			tok = newToken(token.LET, "let")
		
			
	}
}