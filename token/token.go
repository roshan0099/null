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
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	GREATER = ">"
	LESSER  = "<"
	MULTI   = "*"
	DIVIDE  = "/"

	LBRACKET    = "("
	RBRACKET    = ")"
	EXCLAMATORY = "!"
	EQUAL       = "=="
	NEQUAL      = "!="

	//boolean
	TRUE  = "TRUE"
	FALSE = "FALSE"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	COUT = "COUT"

	//declaration
	VAR = "VAR"

	RETURN = "RETURN"

	IF = "IF"
)

var keywords = map[string]string{
	"cout":   COUT,
	"var":    VAR,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
}

func FindKey(word string) string {
	if val, ok := keywords[word]; ok {
		return val
	}
	return IDENT
}
