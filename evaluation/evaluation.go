package evaluation

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

		conditionBool := ifConditionCheck(ch, env)
		return conditionBool

	case *ast.VarStmt:

		rightExp := Eval(ch.Value, env)

		StoreVal(ch, rightExp, env)

	case *ast.BodyStatement:
		return evaluateBody(ch.Statement, env)

	case *ast.ArrayCall:

		index := Eval(ch.Index, env)
		value := indexVal(index, ch.Name, env)
		return value

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

	case *ast.FileHandler:
		evalFile(ch, env)
		// return nil
	}

	return nil
}

func stringStore(ch *ast.StringLine, env *object.Env) object.Object {

	return &object.StringType{
		Word: ch.Line,
	}

}

func EvalLoop(choice *ast.LoopStmt, env *object.Env) {

	conditionLoop := whileConditionCheck(choice.Condition, env)

	if conditionLoop {
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

			switch choice2 := ch.Right.(type) {

			case *ast.FunctionCall:

				if choice2.FunctionName.String() == "ni" || choice2.FunctionName.String() == "ns" {
					niAndns(choice.String(), choice2, env)

				} else if choice2.FunctionName.String() == "chr" {

					env.ChangeVal(choice.String(), rightExp)

				}

			default:

				ok := env.ChangeVal(choice.String(), rightExp)
				if !ok {
					return ErrorMsgUpdate("seems like the variable is not declared ")

				}

			}
			// return nil
		case *ast.InfixExp:

			index, _ := strconv.Atoi(choice.Right.String())
			value := &object.StringType{
				Word: ch.Right.String(),
			}

			env.IndexValChange(choice.Left.String(), index, value)

			return nil

		case *ast.ArrayCall:
			// leftExp := Eval(ch.Left, env)
			rightExp := Eval(ch.Right, env)

			//pending : should check if the value has been changed or not
			arrayValChange(ch.Left.(*ast.ArrayCall), rightExp, env)
		default:
			return ErrorMsgUpdate("Variable declaration not done right")
		}

	} else if ch.Operator == token.LSQBRACKET {

		num := Eval(ch.Right, env)

		return indexVal(num, ch.Left.String(), env)
		// return ErrorMsgUpdate("oooopppssss you on right track but we are woking on that ")
	} else {

		leftExp := Eval(ch.Left, env)

		rightExp := Eval(ch.Right, env)

		return evaluateInfix(leftExp, rightExp, ch.Operator, env)
	}

	return nil
}

func arrayValChange(arrObj ast.Expression, val object.Object, env *object.Env) bool {

	// valInt,_ := strconv.Atoi(val.Inspect())
	indexTemp := Eval(arrObj.(*ast.ArrayCall).Index, env)
	index, _ := strconv.Atoi(indexTemp.Inspect())

	env.IndexValChange(arrObj.(*ast.ArrayCall).Name, index, val)

	return false
}

func indexVal(index object.Object, name string, env *object.Env) object.Object {

	val, _ := strconv.Atoi(index.Inspect())

	nameVar, _ := env.GetEnv(name)

	switch nameVar.(type) {

	case *object.StringType:

		indexedWord := nameVar.(*object.StringType).Word[val]
		return &object.StringType{
			Word: string(indexedWord),
		}

	default:
		return nameVar.(*object.ArrayContents).Body[val]
	}
}

func checkIdentifier(choice *ast.Identifier, env *object.Env) object.Object {
	if val, _ := env.GetEnv(choice.String()); val != nil {

		return val
	}

	if val, _ := object.Builtin[choice.String()]; val != nil {
		builtin := val.(*object.Wrapper)

		return builtin
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

	//checking to know if the user input is required or not
	switch choice := name.Value.(type) {

	//to check if it's funtion to accept user input
	case *ast.FunctionCall:
		if choice.FunctionName.String() == "ns" || choice.FunctionName.String() == "ni" {
			niAndns(name.Name.String(), choice, env)
		} else if choice.FunctionName.String() == "len" {

			length := builtinLen(choice.ArgumentsCall[0].String(), env)

			env.SetEnv(name.Name.String(), length)
		} else if choice.FunctionName.String() == "chr" {

			env.SetEnv(name.Name.String(), exp)
		}
	default:

		env.SetEnv(name.Name.String(), exp)

	}

}

func ifConditionCheck(ifChoice *ast.IfStatement, env *object.Env) object.Object {

	for _, condition := range ifChoice.ElfStmt {

		boolValue := Eval(condition.Condition, env)
		if boolValue.(*object.Boolean).Value {

			return Eval(condition.Body, env)
		}
	}

	if ifChoice.ElseBody != nil {
		return Eval(ifChoice.ElseBody, env)
	}
	return nil
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

	switch leftExp.(type) {
	case *object.StringType:

		leftStmt, rightStmt := leftExp.(*object.StringType).Word, rightExp.(*object.StringType).Word

		switch operator {

		case token.NEQUAL:

			return settingBoolean(leftStmt != rightStmt)

		case token.EQUAL:

			return settingBoolean(leftStmt == rightStmt)

		default:

			ErrorMsgUpdate("Condition not supported by null langauge")
			return nil
		}

	case *object.Integer:

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

			case token.NEQUAL:
				return settingBoolean(leftNumb != rightNumb)
			}
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

	args := choice.(*ast.FunctionCall).ArgumentsCall

	//to avoid unnecessary declaration
	if len(args) > 0 {

		noutResponse := builtin.(*object.Wrapper)

		if noutResponse.Name == "np" {
			analysedArgs := []object.Object{}

			for _, val := range args {
				analysedArgs = append(analysedArgs, Eval(val, env))
			}

			response := noutResponse.WrapperFunc(analysedArgs)
			fmt.Println(response.Inspect())

		} else if noutResponse.Name == "len" {

			return builtinLen(args[0].String(), env)

		} else if noutResponse.Name == "chr" {

			res := Eval(args[0], env)

			return asciiChar(res, env)
		} else if noutResponse.Name == "push" {

			if len(args) == 2 {

				intVal, _ := strconv.Atoi(args[1].String())

				appendVal := &object.Integer{
					Val: int64(intVal),
				}
				variable, _ := env.GetEnv(args[0].String())

				variable.(*object.ArrayContents).Body = append(variable.(*object.ArrayContents).Body, appendVal)

				env.SetEnv(args[0].String(), variable)

			} else {

				ErrorMsgUpdate("Syntx Error \nThis function only supports 2 parameters")
			}
		}
	}
	return nil
}

func evaluateCall(ch []ast.Expression, env *object.Env) []object.Object {

	var body []object.Object
	for _, val := range ch {

		body = append(body, Eval(val, env))

	}

	return body
}

func niAndns(name string, choice *ast.FunctionCall, env *object.Env) {

	if choice.FunctionName.String() == "ns" {

		input := &object.StringType{}

		//checking if there is any arguments
		if len(choice.ArgumentsCall) > 0 {

			fmt.Println(choice.String())

		}
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		input.Word = in.Text()

		env.SetEnv(name, input)

	} else if choice.FunctionName.String() == "ni" {

		input := &object.Integer{}

		if len(choice.ArgumentsCall) > 0 {

			fmt.Println(choice.String())

		}

		fmt.Scanln(&input.Val)

		env.SetEnv(name, input)

	}
}

func builtinLen(name string, env *object.Env) object.Object {

	variable, _ := env.GetEnv(name)

	switch choice := variable.(type) {

	case *object.ArrayContents:

		return &object.Integer{
			Val: int64(len(choice.Body)),
		}

	default:
		return &object.Integer{

			Val: int64(len(variable.Inspect())),
		}
	}

}

func asciiChar(word object.Object, env *object.Env) object.Object {

	switch ch := word.(type) {

	case *object.Integer:

		return &object.StringType{
			Word: string(ch.Val),
		}

	default:
		fmt.Println("oops this function is currently under work")
		os.Exit(0)
	}

	return nil
}

func evalFile(filePoint *ast.FileHandler, env *object.Env) {

	if filePoint.Arguments[1].String() == "r" {
		text, err := ioutil.ReadFile(filePoint.Arguments[0].String())

		if err != nil {
			ErrorMsgUpdate("Error : couldnt find the file \nUsual Reasons : Wrong directory/mispelled file name ")
		}

		//converting the text to object
		evalText := &object.StringType{
			Word: string(text),
		}

		env.SetEnv(filePoint.FileName, evalText)

	} else if filePoint.Arguments[1].String() == "w" {
		ErrorMsgUpdate("Evaluation Error : This function is currently under development")
	} else {
		ErrorMsgUpdate("Error : Invalid file mehtod")
	}
}

func whileConditionCheck(condition ast.Expression, env *object.Env) bool {

	parsedCondition := Eval(condition, env)

	return parsedCondition.(*object.Boolean).Value
}
