package evaluation

import (
	_ "fmt"
	"null/ast"
	"null/object"
)

//creating types just to prevent from copies being made everytime bool is used
var (
	TRUE = &object.Boolean{
		Value: true,
	}
	FALSE = &object.Boolean{
		Value: false,
	}
)

func Eval(typeStruct ast.Node) object.Object {

	switch ch := typeStruct.(type) {

	case *ast.Program:

		return evaluate(ch.Statements)

	case *ast.ParseExp:

		// fmt.Println("this is inside parse exp : ", ch.Token.Type, " -- ", ch.Exp)
		return Eval(ch.Exp)

	case *ast.BooleanValue:
		return settingBoolean(ch.Value)

	case *ast.PrefixExp:
		rightExp := Eval(ch.RightExp)
		return prefixEval(ch.Operator, rightExp)

	case *ast.IntegralParse:

		return &object.Integer{
			Val: ch.Val,
		}
	}
	return nil
}

func evaluate(stmt []ast.Statement) object.Object {

	var result object.Object
	// fmt.Println("this is stat : ", stmt)
	for _, val := range stmt {
		// fmt.Println("tis is inside evaluate  : ", val)
		result = Eval(val)

	}
	return result
}

func settingBoolean(val bool) object.Object {
	if val {
		return TRUE
	} else {
		return FALSE
	}
}

func prefixEval(operator string, rightExp object.Object) object.Object {

	if operator == "!" {
		switch rightExp.Inspect() {
		case "false":
			return TRUE
		case "true":
			return FALSE
		}
	}

	return rightExp
}
