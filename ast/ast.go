package ast

import (
	"bytes"
	"fmt"
	"null/token"
)

//the main node
type Node interface {
	TokenLiteral() string
	String() string
}

//sub node that contains the whole statement
type Statement interface {
	Node
	statementNode()
}

//node that contains expression
type Expression interface {
	Node
	expressionNode()
}

// struct that the whole program in the form of array
type Program struct {
	Statements []Statement
}

//for debugging
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//outputs the whole program
func (p *Program) String() string {
	var concatInfo bytes.Buffer

	for _, s := range p.Statements {
		concatInfo.WriteString(s.String() + " ")
	}
	return concatInfo.String()
}

//struct to identify "var" statements
type VarStmt struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

//satisfying the interface conditions
func (v *VarStmt) statementNode() {}
func (v *VarStmt) TokenLiteral() string {
	return v.Token.Value
}

func (v *VarStmt) String() string {
	var concatInfo bytes.Buffer

	concatInfo.WriteString(v.TokenLiteral() + " " + v.Name.String())

	if v.Value != nil {
		concatInfo.WriteString(v.Value.String())
	}
	// concatInfo.WriteString(";")
	return concatInfo.String()
}

//identifier struct that takes in the expression node in most of the other structs
type Identifier struct {
	Token token.Token
}

func (i *Identifier) TokenLiteral() string {
	fmt.Println("identi : ", i.Token.Value)
	return i.Token.Value
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {
	return i.Token.Value
}

//struct for return statements
type ReturnStmt struct {
	Token token.Token
	Exp   Expression
}

func (r *ReturnStmt) statementNode() {}

func (r *ReturnStmt) TokenLiteral() string { return r.Token.Value }

func (r *ReturnStmt) String() string {

	var concatInfo bytes.Buffer

	concatInfo.WriteString(r.TokenLiteral() + " ")

	if r.Exp != nil {
		concatInfo.WriteString(r.Exp.String())
	}

	return concatInfo.String()
}

//Parse statement

type ParseExp struct {
	Token token.Token
	Exp   Expression
}

func (p *ParseExp) statementNode()       {}
func (p *ParseExp) TokenLiteral() string { return p.Token.Value }
func (p *ParseExp) String() string {
	var concatInfo bytes.Buffer

	if p.Exp != nil {
		concatInfo.WriteString(p.Exp.String())
	}

	return concatInfo.String()
}

//struct for integral parsing
type IntegralParse struct {
	Token token.Token
	Val   int64
}

func (i *IntegralParse) expressionNode()      {}
func (i *IntegralParse) TokenLiteral() string { return i.Token.Value }
func (i *IntegralParse) String() string       { return i.Token.Value }

//struct for prefix
type PrefixExp struct {
	Token    token.Token
	Operator string
	RightExp Expression
}

func (p *PrefixExp) expressionNode() {}
func (p *PrefixExp) TokenLiteral() string {
	return p.Token.Value
}
func (p *PrefixExp) String() string {
	var concatInfo bytes.Buffer

	concatInfo.WriteString("(" + p.Operator + " " + p.RightExp.String() + ")")
	return concatInfo.String()

}

//struct for infix Exp
type InfixExp struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExp) expressionNode()      {}
func (i *InfixExp) TokenLiteral() string { return i.Token.Value }
func (i *InfixExp) String() string {

	var concatInfo bytes.Buffer

	concatInfo.WriteString("(" + i.Left.String() + " " + i.Operator + " " + i.Right.String() + ")")
	fmt.Println("infix : ", concatInfo.String())
	return concatInfo.String()
}
