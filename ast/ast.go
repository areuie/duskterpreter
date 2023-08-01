package ast

import "duskterpreter/token"

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	//this is different
	return ""
}

//statement types

type LetStatement struct {
	//token, identifier, value
	Token token.Token
	Name *Identifier
	Value Expression
}

//implement return tokenliteral (etc let) and 
func (ls *LetStatement) statementNode() { }
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
// func (ls *LetStatement) push(s *Statement) {
	
// }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() { }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }