package object

import (
	"fmt"
	"strings"
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

type BlockStmt struct {
	Block []Object
}

func (b *BlockStmt) Type() string { return "Block" }
func (b *BlockStmt) Inspect() string {

	var merge []string
	// return strings.Join(b.Block[:], "\n")
	for _, point := range b.Block {
		if point != nil {
			merge = append(merge, point.Inspect())
		}
	}
	return strings.Join(merge, "\n")
}

type StringType struct {
	word string
}

// func (s *StringType)
