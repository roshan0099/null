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

//func that returns the next byte
func (lex *Lexer) nextPoint() byte {
	nextpoint := lex.currentPosition

	if nextpoint >= len(lex.code) {
		return 0
	}
	return lex.code[nextpoint+1]
}

//func to return the curent byte thats been read as string
func (lex *Lexer) byteString() string {
	return string(lex.code[lex.currentPosition])
}

//func to identify a string
func (lex *Lexer) StringIdentifier() string {

	lex.read()
	var line string

	for lex.currentPoint != '"' {
		line += string(lex.currentPoint)
		lex.read()
	}
	return line
}

//func to identify the word/symbol and derive meaning out of it
func (lex *Lexer) Identify() token.Token {

	// point := lex.currentPosition
	var info token.Token

	lex.skipVoid()
	// lex.read()

	switch lex.currentPoint {

	case '=':
		if lex.nextPoint() == '=' {

			equalToken := lex.byteString()

			lex.read()

			info = token.Token{
				Type:  token.EQUAL,
				Value: equalToken + string(lex.currentPoint),
			}

		} else {
			info = identifyingTokens(token.ASSIGN, lex.currentPoint)
		}

	case '(':
		// info.Type = token.LBRACKET
		// info.Value = string(lex.currentPoint)

		info = identifyingTokens(token.LBRACKET, lex.currentPoint)

	case ')':
		// info.Type = token.RBRACKET
		// info.Value = string(lex.currentPoint)
		info = identifyingTokens(token.RBRACKET, lex.currentPoint)

	case '"':
		info.Type = token.STRING
		info.Value = lex.StringIdentifier()

	case ';':
		info = identifyingTokens(token.SEMICOLON, lex.currentPoint)

	case ',':
		info = identifyingTokens(token.COMMA, lex.currentPoint)

	case '!':
		if lex.nextPoint() == '=' {
			word := lex.byteString()

			lex.read()

			info = token.Token{
				Type:  token.NEQUAL,
				Value: word + lex.byteString(),
			}
		} else {
			info = identifyingTokens(token.EXCLAMATORY, lex.currentPoint)
		}

	case '+':
		info = identifyingTokens(token.PLUS, lex.currentPoint)

	case '*':
		info = identifyingTokens(token.MULTI, lex.currentPoint)

	case '/':
		info = identifyingTokens(token.DIVIDE, lex.currentPoint)

	case '{':
		info = identifyingTokens(token.LCURLYBRAC, lex.currentPoint)

	case '}':
		info = identifyingTokens(token.RCURLYBRAC, lex.currentPoint)

	case '>':
		info = identifyingTokens(token.GREATER, lex.currentPoint)

	case '<':
		info = identifyingTokens(token.LESSER, lex.currentPoint)

	case 0:
		info.Type = "EOF"
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
