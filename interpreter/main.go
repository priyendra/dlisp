package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/priyendra/dlisp/parser"
)

// Read-Eval-Print-Loop
func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "exit" {
			break
		}
		parsed, err := parser.Parse(input)
		if err != nil {
			fmt.Println(err)
		} else {
			parsed.PrintDebugString()
		}
	}
}

func main() {
	repl()
}
