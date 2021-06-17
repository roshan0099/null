package object

import (
	"bytes"
	"fmt"
	"strings"
)

type MainObject interface {
	Type() string
	Inspect() []Object
}

type Object interface {
	Type() string
	Inspect() string
}

type Integer struct {
	Val int64
}

func (I *Integer) Type() string    { return "INTEGER" }
func (I *Integer) Inspect() string { return fmt.Sprintf("%d", I.Val) }

type Error struct {
	ErrorMsg string
}

func (e *Error) Type() string    { return "ERROR" }
func (e *Error) Inspect() string { return "Error -> " + e.ErrorMsg }

type Boolean struct {
	Value bool
}

func (B *Boolean) Type() string    { return "BOOLEAN" }
func (B *Boolean) Inspect() string { return fmt.Sprintf("%t", B.Value) }

type Null struct{}

func (n *Null) Type() string    { return "NULL" }
func (n *Null) Inspect() string { return "null" }

type BlockStmt struct {
	Block []Object
}

func (b *BlockStmt) Type() string { return "Block" }
func (b *BlockStmt) Inspect() string {

	var merge []string

	for _, point := range b.Block {
		if point != nil {
			merge = append(merge, point.Inspect())
		}
	}
	return strings.Join(merge, "\n")
}

type BlockStmts struct {
	Block []Object
}

func (b *BlockStmts) Type() string { return "Block" }
func (b *BlockStmts) Inspect() []Object {

	var merge []Object

	for _, point := range b.Block {
		if point != nil {
			merge = append(merge, point)
		}
	}
	return merge
}

type StringType struct {
	Word string
}

func (s *StringType) Type() string { return "STRING" }
func (s *StringType) Inspect() string {

	return s.Word
}

type Nout struct {
	Statements []Object
}

func (n *Nout) Type() string { return "nout function" }
func (n *Nout) Inspect() string {

	var concatInfo bytes.Buffer

	for index, val := range n.Statements {

		if index != len(n.Statements)-1 {
			concatInfo.WriteString(val.Inspect() + " ")
		} else {
			concatInfo.WriteString(val.Inspect())
		}
	}
	return concatInfo.String()
}

//Wrapper for function so as to return object

type Wrapper struct {
	Name        string
	WrapperFunc BuiltnCondition
}

func (w *Wrapper) Type() string { return "Wrapper function" }

func (w *Wrapper) Inspect() string {

	return w.Name
}

type WrapCondition func()
type LoopWrapper struct {
	Wrapper WrapCondition
}

func (l *LoopWrapper) Type() string { return "Just a sample function" }
func (l *LoopWrapper) Inspect() string {

	l.Wrapper()
	return ""
}

type ArrayContents struct {
	Body []Object
}

func (a *ArrayContents) Type() string { return "Array Content" }
func (a *ArrayContents) Inspect() string {

	var (
		outinfo   bytes.Buffer
		storeInfo []string
	)
	outinfo.WriteString("[")
	for _, val := range a.Body {

		storeInfo = append(storeInfo, val.Inspect())
	}
	outinfo.WriteString(strings.Join(storeInfo, " "))
	outinfo.WriteString("]")
	return outinfo.String()
}
