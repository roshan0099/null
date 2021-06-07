package object

import (
	"null/ast"
)

type BuiltnCondition func(args []ast.Expression) Object

var Builtin = map[string]Object{

	"nout": &Wrapper{

		WrapperFunc: func(args []ast.Expression) Object {

			noutStore := &Nout{}

			for _, val := range args {

				noutStore.Statements = append(noutStore.Statements, val)
			}

			return noutStore

		},
	},
}
