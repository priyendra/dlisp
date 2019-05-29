package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/priyendra/dlisp/expression"
	"github.com/priyendra/dlisp/value"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Read eval print loop
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			break
		}
		// TODO(deshwal): Support multi-line inputs
		input := scanner.Text()
		if input == "exit" {
			break
		}
		expr, err := expression.Parse(input)
		if err != nil {
			fmt.Println(err)
		} else {
			val, err := expression.Eval(expr)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value.ToString(val))
			}
		}
	}
}
