package parser

import (
	"duskterpreter/ast"
	"duskterpreter/lexer"
	"duskterpreter/token"
	"fmt"
)

//create type parser struct with current and peek token and lexer variable
//new function
//next token
//parse program

type Parser struct {
	l *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors []string
}


func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:l,
		errors: []string{},
	} //putting l into l variable

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser)ParseProgram() *ast.Program {

	//create new AST Program Node
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements= append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program

	//create newProgramASTNode()
	//continues to parse tokens (advanceTokens()) until EOF
	//parse let, if, return statements
	//push the statement into the statements array
	//if it is a parseExpression, it can be broken down into parseOperatorExpression() or parseIntegerLiteral()

}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
		case token.LET:
			return p.parseLetStatement()
		default:
			return nil
	}
} 

func (p *Parser) parseLetStatement() *ast.LetStatement { //why does this have a *
	letstmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	letstmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	
	return letstmt

}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

//errors
func (p *Parser) Errors() []string{
	return p.errors
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("expected next token to be %s, got %s instead.",t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

