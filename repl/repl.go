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
	// fmt.Println("this", input)
	lex := lexer.Create(input)
	// fmt.Println("that")
	parse := parser.New(lex)
	// fmt.Println("kool")
	prgm := parse.ParseProgram()
	// fmt.Println("hmm then", prgm)
	eval := evaluation.Wrapper(prgm, env)
	// fmt.Println("lMO : ", eval)
	for _, val := range eval.(*object.BlockStmts).Block {
		// fmt.Println("kooi")
		if valueCheck := val.Inspect(); valueCheck != "" {
			fmt.Println(valueCheck)
		}
	}
}

func Begin(inPoint *bufio.Scanner, fileVal string) {
	// fmt.Println("reached")
	env := object.NewEnv()

	if fileVal != "" {
		// fmt.Println("reach")
		ReadingProcess(fileVal, env)
	} else {
		var scanLine string
		for {

			if !strings.Contains(scanLine, "ni(") && !strings.Contains(scanLine, "ns(") {
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
