package repl

import (
	"bufio"
	"fmt"
	"null/evaluation"
	"null/lexer"
	"null/parser"
)

func Begin(inPoint *bufio.Scanner) {

	fmt.Println("R E P L ")

	for {

		fmt.Printf(">> ")
		_ = inPoint.Scan()

		scanLine := inPoint.Text()

		lex := lexer.Create(scanLine)
		parse := parser.New(lex)

		prgm := parse.ParseProgram()

		eval := evaluation.Eval(prgm)

		fmt.Println(eval.Inspect())

	}
}
