package lexer

import (
	_ "fmt"
	"null/token"
)

type Lexer struct {
	code            string
	currentPosition int
	nextPosition    int
	currentPoint    byte //current byte under examiine
}

func Create(input string) *Lexer {

	lex := &Lexer{code: input}
	lex.read()
	return lex
}

//read the code byte by byte
func (lex *Lexer) read() {
	if lex.nextPosition >= len(lex.code) {

		lex.currentPoint = 0
	} else {

		lex.currentPoint = lex.code[lex.nextPosition]
	}

	lex.currentPosition = lex.nextPosition
	lex.nextPosition++
}

//read the whole word
func (lex *Lexer) readWord() string {

	pinPostion := lex.currentPosition

	for isLetter(lex.currentPoint) {
		lex.read()
	}
	return lex.code[pinPostion:lex.currentPosition]
}

//read number
func (lex *Lexer) readNum() string {

	pinPostion := lex.currentPosition

	for isNumber(lex.currentPoint) {
		lex.read()
	}
	return lex.code[pinPostion:lex.currentPosition]
}

//to check if the current byte is a letter or not
func isLetter(ch byte) bool {

	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

//to check if the current byte a number or not
func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

//func to identify the word/symbol and derive meaning out of it
func (lex *Lexer) Identify() token.Token {

	// point := lex.currentPosition
	var info token.Token

	lex.skipVoid()
	// lex.read()

	switch lex.currentPoint {

	case '=':
		info = identifyingTokens(token.ASSIGN, lex.currentPoint)

	case '(':
		// info.Type = token.LBRACKET
		// info.Value = string(lex.currentPoint)

		info = identifyingTokens(token.LBRACKET, lex.currentPoint)
	case ')':
		// info.Type = token.RBRACKET
		// info.Value = string(lex.currentPoint)
		info = identifyingTokens(token.RBRACKET, lex.currentPoint)

	case ';':
		info = identifyingTokens(token.SEMICOLON, lex.currentPoint)

	case ',':
		info = identifyingTokens(token.COMMA, lex.currentPoint)

	case 0:
		info.Type = "END"
		info.Value = token.EOF

	default:
		if isLetter(lex.currentPoint) {
			// fmt.Println("yo got it default")
			word := lex.readWord()

			identifyWord := token.FindKey(word)

			info.Type = identifyWord
			info.Value = word
			lex.currentPoint = lex.code[lex.currentPosition-1]
			lex.nextPosition = lex.currentPosition
			lex.currentPosition = lex.currentPosition - 1

		} else if isNumber(lex.currentPoint) {

			word := lex.readNum()

			info.Type = token.INT
			info.Value = word
			lex.currentPoint = lex.code[lex.currentPosition-1]
			lex.nextPosition = lex.currentPosition
			lex.currentPosition = lex.currentPosition - 1

		}

		// fmt.Print("famm")

	}

	lex.read()

	return info
}

//func to skip unwanted space and stuff
func (lex *Lexer) skipVoid() {
	for lex.currentPoint == ' ' || lex.currentPoint == '\t' || lex.currentPoint == '\n' || lex.currentPoint == '\r' {
		lex.read()
	}
}

//to recognize the token and structure it
func identifyingTokens(identifier string, symbol byte) token.Token {
	return token.Token{Type: identifier, Value: string(symbol)}
}
