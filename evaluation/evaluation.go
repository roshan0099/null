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

func Wrapper(typeStruct ast.Node, env *object.Env) object.Object {

	val := Eval(typeStruct, env)

	return val

}

func Eval(typeStruct ast.Node, env *object.Env) object.Object {

	switch ch := typeStruct.(type) {

	case *ast.Program:

		return evaluate(ch.Statements, env)
		// return nil

	case *ast.ParseExp:
		// fmt.Println(ch.Exp)
		// fmt.Println("this is inside parse exp : ", ch.Token.Type, " -- ", ch.Exp)
		return Eval(ch.Exp, env)

	case *ast.BooleanValue:
		return settingBoolean(ch.Value)

	case *ast.PrefixExp:
		rightExp := Eval(ch.RightExp, env)
		return prefixEval(ch.Operator, rightExp)

	case *ast.InfixExp:

		return infixEvaluationWrapper(ch, env)

	case *ast.IfStatement:

		conditionBool := evaluateIf(ch.Condition, ch.Body, ch.ElseBody, env)
		return conditionBool

	case *ast.VarStmt:
		// fmt.Println("it's here", ch.Token.Type, ch.Name, ch.Value)
		rightExp := Eval(ch.Value, env)

		StoreVal(ch.Name, rightExp, env)

	case *ast.BodyStatement:
		return evaluateBody(ch.Statement, env)

	case *ast.Identifier:

		return checkIdentifier(ch, env)

	case *ast.IntegralParse:

		return &object.Integer{
			Val: ch.Val,
		}

	case *ast.LoopStmt:

		EvalLoop(ch, env)

	case *ast.StringLine:
		fmt.Println("hemme")

	}

	return nil
}

func EvalLoop(choice *ast.LoopStmt, env *object.Env) {

	conditionLoop := Eval(choice.Condition, env)

	if conditionLoop.Inspect() != "false" {
		BodyStmtLoop := Eval(choice.Body, env)

		fmt.Println(BodyStmtLoop.Inspect())

		EvalLoop(choice, env)

	}

}

//function to check if the infix is just an expression or to change the value of a variable
func infixEvaluationWrapper(ch *ast.InfixExp, env *object.Env) object.Object {

	if ch.Operator == token.ASSIGN {

		switch choice := ch.Left.(type) {

		case *ast.Identifier:
			rightExp := Eval(ch.Right, env)

			ok := env.ChangeVal(choice.String(), rightExp)

			if !ok {
				return ErrorMsgUpdate("seems like the variable is not declared ")

			}
			return nil

		default:
			return ErrorMsgUpdate("Variable declaration not done right")
		}
	}

	leftExp := Eval(ch.Left, env)
	// fmt.Println("aavo meir")

	rightExp := Eval(ch.Right, env)
	return evaluateInfix(leftExp, rightExp, ch.Operator, env)

}

func checkIdentifier(choice *ast.Identifier, env *object.Env) object.Object {

	val, _ := env.GetEnv(choice.String())

	if val == nil {
		return ErrorMsgUpdate("bru its undefind alright ?")
	}

	return val
}

func ErrorMsgUpdate(message string) object.Object {

	return &object.Error{
		ErrorMsg: message,
	}
}

func StoreVal(name *ast.Identifier, exp object.Object, env *object.Env) {
	// fmt.Println("this is whats inside exp  --->: ", name.Token.Value, exp.Inspect())
	env.SetEnv(name.Token.Value, exp)
}

func evaluateIf(condition ast.Expression, body *ast.BodyStatement,
	elseBody *ast.BodyStatement, env *object.Env) object.Object {

	boolValue := Eval(condition, env)
	var returnVal object.Object

	switch boolValue {
	case TRUE:
		returnVal = Eval(body, env)

	case FALSE:

		if elseBody != nil {
			returnVal = Eval(elseBody, env)
			return returnVal
		}
		returnVal = &object.Null{}

	default:
		ErrorMsgUpdate("condition has some problems :( ")
	}

	return returnVal
}

func evaluate(stmt []ast.Statement, env *object.Env) object.Object {

	result := &object.BlockStmt{}

	for _, val := range stmt {

		if val := Eval(val, env); val != nil {
			result.Block = append(result.Block, val)
		}

	}

	return result
}

func evaluateBody(stmt []ast.Statement, env *object.Env) object.Object {

	result := &object.BlockStmt{}

	for _, val := range stmt {

		//filtering out the nil ones so as to prevent possible error

		if check := Eval(val, env); check != nil {
			result.Block = append(result.Block, check)
		}

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

func evaluateInfix(leftExp, rightExp object.Object, operator string, env *object.Env) object.Object {

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
