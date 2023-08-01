package parser

import (
	"duskterpreter/ast"
	"duskterpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x 5;
	let = 10;
	let 838383;
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