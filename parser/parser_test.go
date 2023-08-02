package parser

import (
	"duskterpreter/ast"
	"duskterpreter/lexer"
	"testing"
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