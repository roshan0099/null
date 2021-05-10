package repl

import (
	"bufio"
	"fmt"
)

func Begin(inPoint *bufio.Scanner) {

	fmt.Println("R E P L ")

	for {

		fmt.Printf(">> ")
		_ = inPoint.Scan()

		scanLine := inPoint.Text()

		fmt.Println(scanLine)

	}
}
