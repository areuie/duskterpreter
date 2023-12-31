package parser

import (
	"duskterpreter/ast"
	"duskterpreter/lexer"
	"duskterpreter/token"
	"strconv"
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

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

type ( //just to have better naming
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

const (
	_ int = iota //gives following constants incrementing nums as values, blank _ is value 0 (following is 1-7)
	LOWEST
	EQUALS // == 
	LESSGREATER // < or >
	SUM // +
	PRODUCT // *
	PREFIX // -X or !X
	CALL //myFunction(X)
)

var precedences = map[token.TokenType]int {
	token.EQ: EQUALS,
	token.NOTEQ: EQUALS,
	token.LSTHAN: LESSGREATER,
	token.GRTHAN: LESSGREATER,
	token.ADD: SUM,
	token.MINUS: SUM,
	token.DIVIDE: PRODUCT,
	token.MULTIPLY: PRODUCT,
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:l,
		errors: []string{},
	} //putting l into l variable

	p.nextToken()
	p.nextToken()

	//don't understand
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	//need to add this bcs parseExpression can't find prefixParseFn for token of type tokten.INT
	//to make it pass, we need to "register" the parseIntegerLiteral method
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.EXCLAIM, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOTEQ, p.parseInfixExpression)
	p.registerInfix(token.LSTHAN, p.parseInfixExpression)
	p.registerInfix(token.GRTHAN, p.parseInfixExpression)
	p.registerInfix(token.ADD, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.DIVIDE, p.parseInfixExpression)
	p.registerInfix(token.MULTIPLY, p.parseInfixExpression)

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
		case token.RETURN:
			return p.parseReturnStatement()
		default:
			return p.parseExpressionStatement()
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
	return false // --> return 12345;
}

//parse return statement
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}
	p.nextToken()

	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

//errors
func (p *Parser) Errors() []string{
	return p.errors
}

func (p *Parser) peekError(t token.TokenType){
	msg := fmt.Sprintf("expected next token to be %s, got %s instead.",t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

//infix/prefix
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

//identifier statement
// 1 + 1
// 2 + 2
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement{
	stmt := &ast.ExpressionStatement{Token: p.currToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

//checks whethere there is a parsing function associated with the token type
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.currToken.Type)
		return nil
	}

	leftExp := prefix() //where does this come from
	//GET BACK TO THIS TO UNDERSTAND
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

//don't understand
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

//parse integer

func (p *Parser) parseIntegerLiteral() ast.Expression {
	//create a new ast.expression, initalize the token
	//then create a ok assign, to catch error, return nil to prematurely exit and then assign value if no error
	lit := &ast.IntegerLiteral{Token: p.currToken}

	//basically an catch block for errors when converting
	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

//parse prefix expression
func (p *Parser) parsePrefixExpression() ast.Expression {
	//create a new instance of a prefix expression and set it as the curernt token
	//token, operator, right
	exp := &ast.PrefixExpression{
		Token: p.currToken,
		Operator: p.currToken.Literal,
	}
	p.nextToken()
	//what does this mean
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

//create a function with infix, where there a param for left side already evaluated and right side need to put in
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	//create a new instance of it
	exp := &ast.InfixExpression{Token: p.currToken, Operator: p.currToken.Literal, Left: left,}
	precedence := p.currPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

//peek and cur precedence
func (p *Parser) currPrecedence() int{
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int{
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) parseBoolean () ast.Expression {
	return &ast.Boolean{Token: p.currToken, Value: p.currTokenIs(token.TRUE)}
}