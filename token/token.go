package token

type Token struct {
	Type  string
	Value string
}

const (
	EOF = "EOF"

	//random words and numbers
	IDENT = "IDENT"
	INT   = "INT"

	//arthmetic operators
	ASSIGN = "="
	PLUS   = "+"

	LBRACKET = "("
	RBRACKET = ")"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	COUT = "COUT"

	//declaration
	VAR = "VAR"
)

var keywords = map[string]string{
	"cout": COUT,
	"var":  VAR,
}

func FindKey(word string) string {
	if val, ok := keywords[word]; ok {
		return val
	}
	return IDENT
}
