package parser

import (
	"fmt"
	"null/ast"
	"null/lexer"
	"null/token"
)

type Parser struct {
	lex *lexer.Lexer
	err []string

	curToken  token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	parse := &Parser{lex: lex}

	//to set both cur and peek
	parse.rollToken()
	parse.rollToken()

	return parse
}

//checking the next token and changing the place of cur and peek respectively
func (p *Parser) rollToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.Identify()
}

func (p *Parser) ParseProgram() *ast.Program {

	program := &ast.Program{}

	program.Statements = []ast.Statement{}

	for p.curToken.Value != token.EOF {

		parseStat := p.ParseStat()
		if parseStat != nil {

			// fmt.Println(parseStat)

			program.Statements = append(program.Statements, parseStat)
		}
		p.rollToken()
	}

	return program

}

//func to parse statment based on what it is
func (p *Parser) ParseStat() ast.Statement {

	switch p.curToken.Type {
	case token.VAR:

		return p.ParseVar()

	default:

		return nil
	}

}

//func to parse statement that starts with var keyword
func (p *Parser) ParseVar() *ast.VarStmt {

	VarParse := &ast.VarStmt{Token: p.curToken}

	if !p.expectingToken(token.IDENT) {

		return nil

	}

	VarParse.Name = &ast.Identifier{
		Token: p.curToken,
	}

	// p.rollToken()

	if !p.expectingToken(token.ASSIGN) {

		return nil
	}

	// if p.peekToken.Type != token.COMMA {
	for p.peekToken.Type != token.SEMICOLON {
		// fmt.Println(" this is 21 : ", p.peekToken)
		p.rollToken()
		if p.curToken.Type == token.VAR {

			// fmt.Println("return 1")

			p.ErrorValidity(token.SEMICOLON)
			return nil
		}

		VarParse.Value = &ast.Identifier{
			Token: p.curToken,
		}
	}

	// } else {
	// 	fmt.Println("hmm else aayi ")
	// 	return nil
	// }
	// fmt.Println("parse var thing 3 : ", VarParse.Value)
	return VarParse

}

//func to check if the coming up token is what we expected or not
func (p *Parser) expectingToken(tokenMatch string) bool {
	if p.peekTokenCheck(tokenMatch) {
		// fmt.Println("this is expecting section ! ", p.curToken)
		p.rollToken()
		// fmt.Println("this is expecting section ", p.curToken)
		return true
	} else {
		p.ErrorValidity(tokenMatch)
		return false
	}
}

func (p *Parser) peekTokenCheck(tokenMatch string) bool {
	return p.peekToken.Type == tokenMatch
}

//error validation should be changed
func (p *Parser) ErrorValidity(tokenMatch string) {
	// fmt.Println("error validity !!")
	message := fmt.Sprintf("oops was expecting %s but got %s :( ", tokenMatch, p.peekToken.Value)

	p.err = append(p.err, message)
	// fmt.Println(p.err)
}

//just to return error
func (p *Parser) Err() []string {
	// fmt.Println("in here : ", p.err)
	return p.err
}
