package main

import (
	// "bufio"
	"fmt"
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

	input := `
		if (5+3){
			var rust = 43;
		}

		return abhaya;
		`

	// lex := lexer.Create(input)

	lex := lexer.Create(input)
	parse := parser.New(lex)
	// err := parse.Err
	sam := *parse.ParseProgram()
	fmt.Println("this is pever : ", sam.Statements[0])
	// for tok := lex.Identify(); tok.Value != token.EOF; tok = lex.Identify() {
	// 	fmt.Printf("%+v\n", tok)
	// }

	// fmt.Println("hey meite : ", sam.Statements[0].String(), err())
	// fmt.Println("hey meite : ", sam.Statements[1].String(), err())
	// fmt.Println("=> ", sam.Statements[0].(*ast.VarStmt).Name)
	fmt.Println("statement : ", sam)
	// for index, val := range sam.Statements {

	// 	// 	// fmt.Println("+> ", err())
	// 	fmt.Println(index, " -- -- -- ", val.String())
	// }

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
