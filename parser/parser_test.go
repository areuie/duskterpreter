package parser

import (
	"duskterpreter/ast"
	"duskterpreter/lexer"
	"testing"
	"fmt"
)

//test let statements
func TestLetStatements(t *testing.T) {
	input := `
		let x = 2;
		let y = 10;

		let foobar = 543824;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	//return error if it is nil or not length of 3

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}


	if len(program.Statements) != 3 {
		t.Fatalf("The program.Statements provided does not have 3 statements. got=%d", len(program.Statements))
	}

	//define the tests identifier structs (expectedIdentifier string)

	tests := []struct {
		expectedIdentifier string
	} {
		{"x"},
		{"y"},
		{"foobar"},
	}

	//COME BACK TO DON'T REALLY UNDERSTAND
	for i, tt := range tests { //range returns i (index) and tt (the value)
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

//tests whether or not it is a let statement
func testLetStatement(t *testing.T, s ast.Statement, name string) bool { //testing env, statement given, expected name
	//if it is not let
	//is it a let statement?
	//value
	//name of the identifier
	if s.TokenLiteral() != "let" {
		t.Errorf(" s.TokenLiteral() not a let statement got:%q", s.TokenLiteral())
		return false
	}
	letstmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s is not an *ast.LetStatement")
		return false
	}
	if letstmt.Name.Value != name { //help
		t.Errorf("letstmt.Name.Value is not '%q'. got:'%q'", name, letstmt.Name.Value)
		return false
	}
	if letstmt.Name.TokenLiteral() != name {//help
		t.Errorf("letstmt.Name is not '%q'. got: '%q'", name, letstmt.Name)
		return false
	}

	return true
}

//test return statements
func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 5;
		return 389432;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("number of statements is not 3. got: %d", len(program.Statements))
	}

	//TestReturnStatement()
	//test if it is the right type, and then value
	
	for _, stmt := range program.Statements{
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not a return statement. got:%T", stmt)
			continue
		}
		if returnStmt.Token.Literal != "return" {
			t.Errorf("returnStmt token literal is not the same. got:%q", returnStmt.TokenLiteral())
		}
	}
}

//test identifier expression

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	//check length, if it is the same type, value
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got: %d", len(program.Statements))
	}
	//check is it a statement and an expression
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression is not an *ast.Identifier. got:%d", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value is not %s. got: %s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not %s. got: %s", "foobar", ident.TokenLiteral())
	}
}


func checkParserErrors (t *testing.T, p *Parser) {
	error := p.errors

	if len(error) == 0 {
		return
	}

	t.Errorf("there are %d errors", len(error))

	for _, msg := range error {
		t.Errorf("parser error: %q",msg)
	}

	t.FailNow()
}

func TestIntegerLiteralExpression (t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	//test length, test if it is ok valid with all the types, test if the values are ok
	if len(program.Statements) != 1 {
		t.Fatalf("program statements not 1. got: %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got: %T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt.Expression is not an IntegerLiteral. got: %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value is not %d. got: %d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral() is not %s. got: %s", "5", literal.TokenLiteral())
	}

}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool{
	//check if is type is integer literal
	//check value and tokenLiteral() is same
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("il not *ast.IntegerLiteral. got: %T", il)
		return false
	}

	if integ.Value != value {
		t.Fatalf("integ.Value not %d. got:%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Fatalf("integ.TokenLiteral() not %d. got: %s", value, integ.TokenLiteral())
		return false
	}

	return true

}

func TestParsingPrefixExpressions (t *testing.T) {
	prefixTests := []struct {
		input string
		operator string
		integerValue int64
	} {
		{"!5;","!", 5,},
		{"-15;","-", 15,},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got = %d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatment. got: %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.PrefixExpression. got: %T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got: %s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
	
}

func TestInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input string
		leftValue int64
		operator string //why is it a string?
		rightValue int64
	} {
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}
	for _, tt := range infixTests {
		//create a lexer, parser, create a program and test for errors, and then check for logic errors
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		//len, make sure value matches up with the ans key
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements not %d. got: %d", 1, len(program.Statements))
		}
		//do ok statements, matches the right datatype ExpressionStatement and InfixStatement
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not amt.ExpressionStatement. got: %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.InfixExpression. got: %T", stmt.Expression)
		}
		//now test the individual values, integer literal of the left and right got/expected and also same opperator
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return //why is it just return????
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got='%s'", tt.operator, exp.Operator) 
		}
	}
}