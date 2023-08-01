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
	Value string
}

//implement return tokenliteral (etc let) and 
func (ls *LetStatement) expressionNode() { }
func (ls *LetStatement) tokenLiteral() string { return ls.Token.Literal }
// func (ls *LetStatement) push(s *Statement) {
	
// }

type Identifier struct {
	Token token.Token
	Name *Identifier
}

func (i *Identifier) statementNode() { }
func (i *Identifier) tokenLiteral() string { return i.Token.Literal }