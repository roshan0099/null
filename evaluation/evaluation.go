package evaluation

import (
	"fmt"
	"null/ast"
	"null/object"
)

func Eval(typeStruct ast.Node) object.Object {

	switch ch := typeStruct.(type) {

	case *ast.Program:

		return evaluate(ch.Statements)

	case *ast.ParseExp:

		fmt.Println("this is inside parse exp : ", ch.Token.Type, " -- ", ch.Exp)
		return Eval(ch.Exp)

	case *ast.IntegralParse:

		return &object.Integer{
			Val: ch.Val,
		}
	}
	return nil
}

func evaluate(stmt []ast.Statement) object.Object {

	var result object.Object
	fmt.Println("this is stat : ", stmt)
	for _, val := range stmt {
		fmt.Println("tis is inside evaluate  : ", val)
		result = Eval(val)

	}
	return result
}
