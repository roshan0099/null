package parser

import (
	"null/token"
)

var Precedence = map[string]int{

	token.PLUS:    PLUSMINUS,
	token.MINUS:   PLUSMINUS,
	token.GREATER: LESSGREAT,
	token.LESSER:  LESSGREAT,
	token.EQUAL:   EQUAL,
	token.NEQUAL:  EQUAL,
	token.MULTI:   CROSSDIV,
	token.DIVIDE:  CROSSDIV,
}
