package evaluation

import (
	"fmt"
	"null/ast"
	"null/object"
	"null/token"
	"os"
	"strconv"
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

//collects all the interfaces
var collection object.MainObject

func Wrapper(typeStruct ast.Node, env *object.Env) object.MainObject {

	Eval(typeStruct, env)

	return collection

}

func Eval(typeStruct ast.Node, env *object.Env) object.Object {

	switch ch := typeStruct.(type) {

	case *ast.Program:
		collection = evaluate(ch.Statements, env)

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

	case *ast.FunctionCall:

		name := Eval(ch.FunctionName, env)
		return evalFtnCall(ch, name, env)

	case *ast.IfStatement:

		conditionBool := evaluateIf(ch.Condition, ch.Body, ch.ElseBody, env)
		return conditionBool

	case *ast.VarStmt:
		// fmt.Println("it's here", ch.Token.Type, ch.Name, ch.Value)
		fmt.Println("var")
		rightExp := Eval(ch.Value, env)
		fmt.Println("///////////", rightExp)
		StoreVal(ch, rightExp, env)

	case *ast.BodyStatement:
		return evaluateBody(ch.Statement, env)

	case *ast.Identifier:
		return checkIdentifier(ch, env)

	case *ast.ArrayType:
		return &object.ArrayContents{

			Body: evaluateCall(ch.ArrayBody, env),
		}

	case *ast.IntegralParse:
		return &object.Integer{
			Val: ch.Val,
		}

	case *ast.LoopStmt:

		fam := &object.LoopWrapper{
			Wrapper: func() {

				EvalLoop(ch, env)
			},
		}

		return fam

	case *ast.StringLine:

		return stringStore(ch, env)

	}

	return nil
}

func stringStore(ch *ast.StringLine, env *object.Env) object.Object {
	// fmt.Println("======", ch.Token)
	return &object.StringType{
		Word: ch.Line,
	}

}

func EvalLoop(choice *ast.LoopStmt, env *object.Env) {

	conditionLoop := Eval(choice.Condition, env)

	if conditionLoop.Inspect() != "false" {
		BodyStmtLoop := Eval(choice.Body, env)
		if BodyStmtLoop.Inspect() != "" {
			//printing the contents in the loop

			fmt.Println(BodyStmtLoop.Inspect())
		}
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
	} else if ch.Operator == token.LSQBRACKET {

		num := Eval(ch.Right, env)
		return arrayIndexVal(num, ch.Left, env)
		// return ErrorMsgUpdate("oooopppssss you on right track but we are woking on that ")
	}

	leftExp := Eval(ch.Left, env)

	rightExp := Eval(ch.Right, env)
	return evaluateInfix(leftExp, rightExp, ch.Operator, env)

}

////////////
//////////////////////
////////////////////////////////
func arrayIndexVal(index object.Object, name ast.Expression, env *object.Env) object.Object {
	val, _ := strconv.Atoi(index.Inspect())

	kal, _ := env.GetEnv(name.String())

	return kal.(*object.ArrayContents).Body[val]
}

///////////////////////////////////
////////////////////////////
//////////////

func checkIdentifier(choice *ast.Identifier, env *object.Env) object.Object {

	if val, _ := env.GetEnv(choice.String()); val != nil {

		return val
	}

	if val, _ := object.Builtin[choice.String()]; val != nil {
		nout := val.(*object.Wrapper)
		return nout
	}

	return ErrorMsgUpdate(" undeclared variable/function ")
}

func ErrorMsgUpdate(message string) object.Object {

	fmt.Println(message)
	os.Exit(0)
	return &object.Error{
		ErrorMsg: message,
	}
}

func StoreVal(name *ast.VarStmt, exp object.Object, env *object.Env) {
	// fmt.Println("this is whats inside exp  --->: ", name.Token.Value, exp.Inspect())

	//checking to know if the user input is required or not
	switch choice := name.Value.(type) {

	case *ast.FunctionCall:

		///////////////////
		//////////////////////////
		//////////////////////////////////
		fmt.Println("yo bithv : ", choice.FunctionName)
		jan := &object.StringType{}
		// var sa string
		fmt.Scanf("%s", &jan.Word)
		// fmt.Println(jan)
		env.SetEnv(name.Name.String(), jan)
		///////////////////////////////////////
		////////////////////////////////
		///////////////////////
	default:
		fmt.Println("ithaanu mone saanam", exp)
		env.SetEnv(name.Name.String(), exp)

	}
	// fmt.Println(name.String(), name.TokenLiteral())

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
		// returnVal = &object.Null{}

	default:
		ErrorMsgUpdate("condition has some problems :( ")

	}

	return returnVal
}

func evaluate(stmt []ast.Statement, env *object.Env) object.MainObject {

	result := &object.BlockStmts{}

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

		case token.MODULO:
			return &object.Integer{Val: leftNumb % rightNumb}

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

func evalFtnCall(choice ast.Expression, builtin object.Object, env *object.Env) object.Object {

	// fmt.Println("hey biatch : ", choice.(*ast.FunctionCall).ArgumentsCall)
	args := choice.(*ast.FunctionCall).ArgumentsCall

	analysedArgs := []object.Object{}

	for _, val := range args {
		analysedArgs = append(analysedArgs, Eval(val, env))
	}

	noutResponse := builtin.(*object.Wrapper)

	response := noutResponse.WrapperFunc(analysedArgs)

	// fmt.Println("there you g maite finally : ", kal)
	return response
}

func evaluateCall(ch []ast.Expression, env *object.Env) []object.Object {

	var body []object.Object
	for _, val := range ch {

		body = append(body, Eval(val, env))

	}

	return body
}
