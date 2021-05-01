package parser

import (
	"null/ast"
	"null/lexer"
	"testing"
)

func TestParser(t *testing.T) {

	input := `
	var supes = 30 ;
	var 67;
	`

	test := []struct {
		// expectedInput string
		expectedToken string
		expectedVar   string
		expectedVal   string
	}{
		{"var", "supes", "30"},
		{"var", "smash", "30"},
	}

	lex := lexer.Create(input)

	parse := New(lex)

	parsedProgram := parse.ParseProgram()
	testParseError(t, parse)

	for i, val := range test {
		parsedVarStat := parsedProgram.Statements[i]
		if val.expectedToken != parsedVarStat.(*ast.VarStmt).Token.Value {
			// t.Fatalf("oops gone wrong %s", parse.Err)
			t.Fatalf("oops gone wrong 1")
		}

		if val.expectedVar != parsedVarStat.(*ast.VarStmt).Name.TokenLiteral() {
			t.Fatalf("oops gone wrong 2 %s  --- %s", val.expectedVar, parsedVarStat.(*ast.VarStmt).Name.TokenLiteral())
		}
		if val.expectedVal != parsedVarStat.(*ast.VarStmt).Value.TokenLiteral() {
			t.Fatalf("oops gone wrong 3 %s --- %s", val.expectedVal, parsedVarStat.(*ast.VarStmt).Value.TokenLiteral())
		}

	}

}

func testParseError(t *testing.T, parse *Parser) {
	if len(parse.Err()) == 0 {
		return
	} else {
		for _, v := range parse.Err() {
			t.Fatalf(v)
		}
		t.FailNow()
	}
}
