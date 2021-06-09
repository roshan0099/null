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

		for _, val := range eval.(*object.BlockStmts).Block {
			if valueCheck := val.Inspect(); valueCheck != "" {
				fmt.Println(valueCheck)
			}
		}

	}
}
