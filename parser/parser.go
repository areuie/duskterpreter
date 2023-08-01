package parser

import (
	"duskterpreter/lexer"
	"duskterpreter/ast"
	"duskterpreter/token"
)

//create type parser struct with current and peek token and lexer variable
//new function
//next token
//parse program

type Parser struct {
	l *lexer.Lexer
	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l} //putting l into l variable

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser)ParseProgram() *ast.Program {
	return nil

	//create newProgramASTNode()
	//continues to parse tokens (advanceTokens()) until EOF
	//parse let, if, return statements
	//push the statement into the statements array
	//if it is a parseExpression, it can be broken down into parseOperatorExpression() or parseIntegerLiteral()

}


