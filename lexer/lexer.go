package lexer

import "duskterpreter/token"

type Lexer struct { //lexing is converting source code to tokens
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

//make nextToken and newToken
func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

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
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '*':
			tok = newToken(token.MULTIPLY, l.ch)
		case '/':
			tok = newToken(token.DIVIDE, l.ch)
		case '!':
			/*
				TODO:
				If we were to start supporting more
				two-character tokens in Monkey, we should probably abstract the behaviour away in a method
				called makeTwoCharToken that peeks and advances if it found the right token
			*/
			if l.peekChar() == '=' {
				tok = token.Token{Type: token.NOTEQ, Literal: string(l.ch) + string(l.peekChar())}
				l.readChar()
			} else {
				tok = newToken(token.EXCLAIM, l.ch)
			}
		case '=':
			if l.peekChar() == '='{
				tok = token.Token{Type: token.EQ, Literal: string(l.ch) + string(l.peekChar())}
				l.readChar()
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
		case '<':
			tok = newToken(token.LSTHAN, l.ch)
		case '>':
			tok = newToken(token.GRTHAN, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			//if it is a character then run readIdent
			if isLetter(l.ch) {
				tok.Literal = l.readIdent()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isInt(l.ch) {
				tok.Literal = l.readInt()
				tok.Type = token.INT
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

/*
	Identify the type of character
*/
func isLetter(ch byte) bool{
	if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'  || ch == '_'{
		return true
	}
	return false
}

func isInt(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}

/*
	Reading methods
*/
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

func (l *Lexer) readIdent() string {
	//position to new position
	position := l.position

	for isLetter(l.ch){
		l.readChar()
	}

	return l.input[position:l.position]
}
//TODO: Add in float, hexa, decimal functionality
func (l *Lexer) readInt() string {
	//position to new position
	position := l.position

	for isInt(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l*Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r'{
		l.readChar()
	}
}

func (l*Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}