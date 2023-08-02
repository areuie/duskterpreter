package ast

import (
	"bytes"
	"duskterpreter/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string { //creates a buffer and returns the string() output
	var out bytes.Buffer //what is this format

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() { }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {return i.Value}

//return type

type ReturnStatement struct{
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() { }
func (rs *ReturnStatement) TokenLiteral() string {return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

//expression statement

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() { }
func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

//integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string { return il.Token.Literal }

//prefix expression

type PrefixExpression struct {
	Token token.Token //prefix token !, -
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() { }
func (pe *PrefixExpression) TokenLiteral() string {return pe.Token.Literal}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
