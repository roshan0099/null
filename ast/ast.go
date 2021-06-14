package ast

import (
	"bytes"
	_ "fmt"
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

	concatInfo.WriteString(v.TokenLiteral() + " " + v.Name.String() + " ")

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
	// fmt.Println("infix : ", concatInfo.String())
	return concatInfo.String()
}

type BooleanValue struct {
	Token token.Token
	Value bool
}

func (b *BooleanValue) expressionNode()      {}
func (b *BooleanValue) TokenLiteral() string { return b.Token.Value }
func (b *BooleanValue) String() string {

	return b.Token.Value
}

type IfStatement struct {
	Token     token.Token
	Condition Expression
	Body      *BodyStatement
	ElseBody  *BodyStatement
}

func (I *IfStatement) TokenLiteral() string {
	return I.Token.Value
}

func (I *IfStatement) expressionNode() {}

func (I *IfStatement) String() string {

	var concatInfo bytes.Buffer

	concatInfo.WriteString("if")

	concatInfo.WriteString(I.Condition.String() + "{")

	concatInfo.WriteString(I.Body.String() + "}")

	if I.ElseBody == nil {
		return concatInfo.String()
	}

	concatInfo.WriteString("else" + "{" + I.ElseBody.String() + "}")

	return concatInfo.String()

}

type BodyStatement struct {
	Token     token.Token
	Statement []Statement
}

func (b *BodyStatement) TokenLiteral() string {
	return b.Token.Value
}

func (b *BodyStatement) statementNode() {}

func (b *BodyStatement) String() string {

	var concatInfo bytes.Buffer

	for _, val := range b.Statement {

		concatInfo.WriteString(val.String())

	}

	return concatInfo.String()
}

type LoopStmt struct {
	Token     token.Token
	Condition Expression
	Body      *BodyStatement
}

func (l *LoopStmt) statementNode() {}
func (l *LoopStmt) TokenLiteral() string {
	return "while loop"
}

func (l *LoopStmt) String() string {
	var concatInfo bytes.Buffer

	concatInfo.WriteString(l.Token.Value)

	concatInfo.WriteString("(" + l.Condition.String() + ")")

	concatInfo.WriteString("{" + l.Body.String() + "}")

	return concatInfo.String()
}

type Sample struct {
	SampleText string
}

func (s *Sample) statementNode() {}
func (s *Sample) TokenLiteral() string {
	return "this is Sample "
}
func (s *Sample) String() string {
	return "this is Sample's string"
}

type StringLine struct {
	Token token.Token
	Line  string
}

func (s *StringLine) expressionNode() {}
func (s *StringLine) TokenLiteral() string {
	return token.STRING
}
func (s *StringLine) String() string {
	return s.Line
}

type FunctionDeclare struct {
	Token     token.Token
	Arguments []*Identifier
	Body      *BodyStatement
}

func (f *FunctionDeclare) TokenLiteral() string { return f.Token.Type }
func (f *FunctionDeclare) expressionNode()      {}
func (f *FunctionDeclare) String() string {

	var concatInfo bytes.Buffer

	concatInfo.WriteString(f.Token.Value + "(")

	for _, val := range f.Arguments {
		concatInfo.WriteString(val.String() + " ")
	}

	concatInfo.WriteString(")" + "{ \n" + f.Body.String() + " \n}")

	return concatInfo.String()
}

type FunctionCall struct {
	Token         token.Token
	FunctionName  Expression
	ArgumentsCall []Expression
}

func (f *FunctionCall) TokenLiteral() string {
	return f.Token.Type
}

func (f *FunctionCall) expressionNode() {}
func (f *FunctionCall) String() string {

	var concatInfo bytes.Buffer

	for _, val := range f.ArgumentsCall {
		concatInfo.WriteString(val.String() + " ")
	}
	return concatInfo.String()
}

type ArrayType struct {
	Token     token.Token
	ArrayBody []Expression
}

func (a *ArrayType) TokenLiteral() string { return a.Token.Value }
func (a *ArrayType) String() string {

	var concatInfo bytes.Buffer

	for _, val := range a.ArrayBody {

		concatInfo.WriteString(val.String() + " ")
	}

	return concatInfo.String()
}
func (a *ArrayType) expressionNode() {}
