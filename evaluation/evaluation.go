package evaluation

import (
	"fmt"
	"null/ast"
	"null/object"
	"null/token"
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

	case *ast.InfixExp:

		leftExp := Eval(ch.Left)
		rightExp := Eval(ch.Right)
		return evaluateInfix(leftExp, rightExp, ch.Operator)

	case *ast.IfStatement:
		conditionBool := evaluateIf(ch.Condition, ch.Body, ch.ElseBody)
		return conditionBool

	case *ast.BodyStatement:
		return evaluate(ch.Statement)

	case *ast.IntegralParse:

		return &object.Integer{
			Val: ch.Val,
		}
	}

	return nil
}

func evaluateIf(condition ast.Expression, body *ast.BodyStatement,
	elseBody *ast.BodyStatement) object.Object {

	boolValue := Eval(condition)
	var returnVal object.Object

	switch boolValue {
	case TRUE:
		returnVal = Eval(body)

	case FALSE:

		if elseBody != nil {
			returnVal = Eval(elseBody)
			return returnVal
		}
		returnVal = &object.Null{}

	default:
		fmt.Println("blaaaaahhhhh")
	}

	return returnVal
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

func evaluateInfix(leftExp, rightExp object.Object, operator string) object.Object {

	leftNumb, rightNumb := leftExp.(*object.Integer).Val, rightExp.(*object.Integer).Val
	if leftExp.Type() == "INTEGER" && rightExp.Type() == "INTEGER" {
		switch operator {
		case token.PLUS:
			return &object.Integer{Val: leftNumb + rightNumb}

		case token.MINUS:
			return &object.Integer{Val: leftNumb - rightNumb}

		case token.DIVIDE:
			return &object.Integer{Val: leftNumb / rightNumb}

		case token.MULTI:
			return &object.Integer{Val: leftNumb * rightNumb}

		case token.EQUAL:
			return settingBoolean(leftNumb == rightNumb)

		case token.GREATER:
			return settingBoolean(leftNumb > rightNumb)

		case token.LESSER:
			return settingBoolean(leftNumb < rightNumb)
		}
	}
	return nil
}

func prefixEval(operator string, rightExp object.Object) object.Object {

	if operator == "!" {
		switch rightExp.Inspect() {
		case "false":
			return TRUE
		case "true":
			return FALSE
		}
	} else if operator == "-" {
		switch rightExp.Inspect() {
		case "false":
			return TRUE
		case "true":
			return FALSE
		}
	}

	return rightExp
}
