package parser

import (
	"testing"
	"duskterpreter/ast"
	"duskterpreter/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 2;
		let y = 10;

		let foobar = 543824;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

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

	funct testLetStatement(t *testing.T, s ast.Statement, name string) bool {
		if 
	}


}