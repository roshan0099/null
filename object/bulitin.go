package object

import "null/ast"

type builtnCondition func(args ...ast.Expression) Object

var builtin = map[string]builtnCondition{
	"nout": func(args ...ast.Expression) Object {

		return nil

	},
}
