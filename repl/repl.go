package repl

import (
	"bufio"
	"fmt"
	"null/evaluation"
	"null/lexer"
	"null/object"
	"null/parser"
)

func Begin(inPoint *bufio.Scanner) {

	fmt.Println("R E P L ")

	env := object.NewEnv()

	for {

		fmt.Printf(">> ")
		_ = inPoint.Scan()

		scanLine := inPoint.Text()
		if scanLine == "bye" {
			break
		}
		lex := lexer.Create(scanLine)
		parse := parser.New(lex)

		prgm := parse.ParseProgram()

		eval := evaluation.Wrapper(prgm, env)
		// if eval == nil {
		// 	fmt.Println("kkk -> ", eval)
		// 	continue
		// } else {
		// 	// for sam := range eval {
		// 	// 	fmt.Println(sam)
		// 	// }
		// 	fmt.Println("==== ", eval)
		// }
		if eval.Inspect() != "" {
			fmt.Println(eval.Inspect())
		}
	}
}
