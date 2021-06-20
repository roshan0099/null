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
	MODULO  = "%"

	LSQBRACKET  = "["
	RSQBRACKET  = "]"
	LBRACKET    = "("
	RBRACKET    = ")"
	LCURLYBRAC  = "{"
	RCURLYBRAC  = "}"
	EXCLAMATORY = "!"
	EQUAL       = "=="
	NEQUAL      = "!="

	STRING = "STRING"

	FUNCTION = "FUNCTION"

	//builtin
	LENGTH = "LENGTH"

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

	IF   = "IF"
	ELSE = "ELSE"

	//LOOP
	WHILE = "WHILE"
)

var keywords = map[string]string{
	"cout":   COUT,
	"var":    VAR,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"ftn":    FUNCTION,
	"len":    LENGTH,
}

func FindKey(word string) string {
	if val, ok := keywords[word]; ok {
		return val
	}
	return IDENT
}
