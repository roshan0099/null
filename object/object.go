package object

import (
	"fmt"
)

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
