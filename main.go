package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	_ "null/lexer"
	_ "null/parser"
	"null/repl"
	"os"
	"strings"
)

func main() {

	var fileVal []byte
	args := os.Args

	argsFileName := args[1:]

	read := os.Stdin

	scan := bufio.NewScanner(read)

	if len(argsFileName) == 1 {

		fileName := string(argsFileName[0])
		if strings.Split(fileName, ".")[1] == "nl" {
			fileVal, _ = ioutil.ReadFile(fileName)
			repl.Begin(scan, string(fileVal))
		}
		os.Exit(0)
	}

	fmt.Println("--- NULL ---")
	repl.Begin(scan, string(fileVal))

}
