package parser

import (
	"fmt"
	"null/ast"
	"null/lexer"
	"null/token"
	"strconv"
)

const (
	_ = iota
	GENERAL
	EQUAL
	LESSGREAT
	PLUSMINUS
	CROSSDIV
	PREFIX
	CALL
)

//expression implementation
type (
	prefixFuncs func() ast.Expression

	infixFuncs func(ast.Expression) ast.Expression
)

type Parser struct {
	lex *lexer.Lexer
	err []string

	curToken  token.Token
	peekToken token.Token

	prefixParse map[string]prefixFuncs
	infixParse  map[string]infixFuncs
}

func New(lex *lexer.Lexer) *Parser {
	parse := &Parser{lex: lex}

	//setting up the map
	parse.prefixParse = make(map[string]prefixFuncs)
	parse.assignPrefix(token.IDENT, parse.identifierParse)
	parse.assignPrefix(token.INT, parse.intgerParse)
	parse.assignPrefix(token.MINUS, parse.parsePrefix)

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

	case token.RETURN:

		return p.ParseReturn()

	default:

		return p.ParseExpressionStmt()
		// fmt.Println("hello : ", p.curToken)
		// return nil
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

		// p.rollToken()
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

func (p *Parser) ParseReturn() *ast.ReturnStmt {

	returnStmt := &ast.ReturnStmt{
		Token: p.curToken,
	}

	//interim
	for !p.peekTokenCheck(token.SEMICOLON) {
		p.rollToken()
		fmt.Println("---")
		returnStmt.Exp = &ast.Identifier{
			Token: p.curToken,
		}
	}

	return returnStmt
}

//func to append error messages
func (p *Parser) errorMsg(errMessage string) {

	message := fmt.Sprintf("oops couldnt parse the token %s", errMessage)
	p.err = append(p.err, message)
}

//func to parse all the no statments in NULL which are generally expressions
func (p *Parser) ParseExpressionStmt() *ast.ParseExp {
	prgrmStmt := &ast.ParseExp{
		Token: p.curToken,
	}

	prgrmStmt.Exp = p.ParsingExpression(GENERAL)

	if p.peekTokenCheck(token.SEMICOLON) {
		p.rollToken()
	}

	return prgrmStmt

}

func (p *Parser) ParsingExpression(order int) ast.Expression {

	prefix := p.prefixParse[p.curToken.Type]

	if prefix == nil {
		p.errorMsg(p.curToken.Value)
		return nil
	}

	return prefix()
}

//for non integer expression
func (p *Parser) identifierParse() ast.Expression {

	return &ast.Identifier{Token: p.curToken}
}

//for integer expression
func (p *Parser) intgerParse() ast.Expression {
	integer := &ast.IntegralParse{
		Token: p.curToken,
	}

	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)

	if err != nil {
		p.errorMsg(p.curToken.Value)
	}

	integer.Val = value

	return integer
}

func (p *Parser) parsePrefix() ast.Expression {
	prefixStmt := &ast.PrefixExp{
		Token:    p.curToken,
		Operator: p.curToken.Value,
	}

	p.rollToken()

	rightStmt := p.ParsingExpression(GENERAL)

	prefixStmt.RightExp = rightStmt

	return prefixStmt
}

func (p *Parser) assignPrefix(token string, function prefixFuncs) {
	fmt.Println("assign infix", function())
	p.prefixParse[token] = function
}

func (p *Parser) assignInfix(token string, function infixFuncs) {
	p.infixParse[token] = function
}
