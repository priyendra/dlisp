package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/priyendra/dlisp/expression"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	env := NewStdEnvironment()
	// Read eval print loop
	for {
		fmt.Print("$ ")
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
			val, err := expression.Eval(expr, env)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression.ToString(val))
			}
		}
	}
}
