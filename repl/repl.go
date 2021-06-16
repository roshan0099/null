package repl

import (
	"bufio"
	"fmt"
	"null/evaluation"
	"null/lexer"
	"null/object"
	"null/parser"
	"strings"
)

func ReadingProcess(input string, env *object.Env) {
	lex := lexer.Create(input)
	parse := parser.New(lex)
	prgm := parse.ParseProgram()

	eval := evaluation.Wrapper(prgm, env)

	for _, val := range eval.(*object.BlockStmts).Block {
		if valueCheck := val.Inspect(); valueCheck != "" {
			fmt.Println(valueCheck)
		}
	}
}

func Begin(inPoint *bufio.Scanner, fileVal string) {

	env := object.NewEnv()

	if fileVal != "" {
		ReadingProcess(fileVal, env)
	} else {
		var scanLine string
		for {

			if !strings.Contains(scanLine, "nin(") {
				fmt.Printf(">> ")
			}
			_ = inPoint.Scan()
			scanLine = inPoint.Text()
			if scanLine == "bye" {
				break
			}
			ReadingProcess(scanLine, env)

		}
	}

}
