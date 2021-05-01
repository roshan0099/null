package main

import (
	// "bufio"
	"fmt"
	"null/ast"
	"null/lexer"
	"null/parser"
	_ "null/token"
	// "os"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- NULL ---")
	// val, err := reader.ReadString('\n')

	// if err != nil {

	// 	fmt.Println("something went wrnog")
	// 	return
	// }

	// java := lexer.Crete(val)

	input := `var sal = 39;			`

	// lex := lexer.Create(input)

	lex := lexer.Create(input)
	parse := parser.New(lex)
	// err := parse.Err
	sam := *parse.ParseProgram()

	// for tok := lex.Identify(); tok.Value != token.EOF; tok = lex.Identify() {
	// 	fmt.Printf("%+v\n", tok)
	// }

	// fmt.Println("hey meite : ", sam, err())
	// fmt.Println("=> ", sam.Statements[0].(*ast.VarStmt).Name)
	for index, val := range sam.Statements {

		// fmt.Println("+> ", err())
		fmt.Println(index, " -- -- -- ", val.(*ast.VarStmt).Token.Value)
		fmt.Println(index, " -- -- -- ", val.(*ast.VarStmt).Name.TokenLiteral())
		fmt.Println(index, " -- -- -- ", val.(*ast.VarStmt).Value.TokenLiteral())

	}

	// for {
	// 	tok := lex.Identify()
	// 	if tok.Value == token.EOF {

	// 		fmt.Printf("the type : %T ", token.EOF)
	// 		// fmt.Println(tok.Value, "meir")
	// 		break
	// 	}
	// 	fmt.Printf("%+v\n", tok)

	// }

}
