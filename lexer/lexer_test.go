package lexer

import (
	"null/token"
	"testing"
)

func TestLexer(t *testing.T) {

	input := "cout(2+3)"

	test := []struct {
		inputVal string
		output   string
	}{
		{token.COUT, "cout"},
		{token.LBRACKET, "("},
	}

	lex := Create(input)

	for _, val := range test {
		tok := lex.Identify()
		if tok.Value != val.output {
			t.Fatalf("good bye ")
		}

	}

}
