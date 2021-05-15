package main

import (
	"bufio"
	"fmt"
	_ "null/lexer"
	_ "null/parser"
	"null/repl"
	"os"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- NULL ---")

	read := os.Stdin

	scan := bufio.NewScanner(read)

	repl.Begin(scan)

	// val, err := reader.ReadString('\n')

	// if err != nil {

	// 	fmt.Println("something went wrnog")
	// 	return
	// }

	// java := lexer.Crete(val)

	// input := `
	// if(1==0){
	// 	10+1;
	// }
	// var sam = 2;
	// `

	// // lex := lexer.Create(input)

	// lex := lexer.Create(input)
	// parse := parser.New(lex)
	// // // err := parse.Err
	// sam := *parse.ParseProgram()
	// fmt.Println("this is pever : ", sam)
	// for tok := lex.Identify(); tok.Value != token.EOF; tok = lex.Identify() {
	// 	fmt.Printf("%+v\n", tok)
	// }

	// fmt.Println("hey meite : ", sam.Statements[0].String(), err())
	// fmt.Println("hey meite : ", sam.Statements[1].String(), err())
	// fmt.Println("=> ", sam.Statements[0].(*ast.VarStmt).Name)
	// fmt.Println("statement : ", sam)

	// for index, val := range sam.Statements {
	// 	fmt.Println("this is inside loop : ", index, " ---- ", val)
	// }
	// 	// 	// 	// fmt.Println("+> ", err())
	// 	switch ch := val.(type) {
	// 	case *ast.ParseExp:
	// 		fmt.Println(index, " this is ", ch.String())

	// 	case *ast.VarStmt:
	// 		fmt.Println(index, " this is ", ch.Value)

	// 	default:
	// 		fmt.Println("ngaaa ")

	// 	}

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
