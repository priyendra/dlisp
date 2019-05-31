package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/priyendra/dlisp/builtins"
	"github.com/priyendra/dlisp/expression"
)

func stdEnv() expression.Environment {
	env := expression.NewEnvironment()
	env.Names["identity"] = builtins.Identity
	env.Names["+"] = builtins.Plus
	env.Names["-"] = builtins.Minus
	env.Names["*"] = builtins.Multiply
	env.Names["/"] = builtins.Divide
	env.Names["%"] = builtins.Mod
	env.Names["pi"] = expression.Float(math.Pi)
	return env
}

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
			env := stdEnv()
			val, err := expression.Eval(expr, &env)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression.ToString(val))
			}
		}
	}
}
