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
	env.Names["pi"] = expression.Float(math.Pi)
	env.Names["identity"] = builtins.Identity
	env.Names["+"] = builtins.Plus
	env.Names["-"] = builtins.Minus
	env.Names["*"] = builtins.Multiply
	env.Names["/"] = builtins.Divide
	env.Names["%"] = builtins.Mod
	env.Names[">"] = builtins.GreaterThan
	env.Names[">="] = builtins.GreaterEqual
	env.Names["<"] = builtins.LessThan
	env.Names["<="] = builtins.LessEqual
	env.Names["="] = builtins.Equal
	env.Names["!="] = builtins.NotEqual
	env.Names["&&"] = builtins.LogicalAnd
	env.Names["||"] = builtins.LogicalOr
	env.Names["!"] = builtins.LogicalNot
	env.Names["if"] = builtins.If
	env.Names["true"] = expression.Bool(true)
	env.Names["false"] = expression.Bool(false)
	return env
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	env := stdEnv()
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
			val, err := expression.Eval(expr, &env)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression.ToString(val))
			}
		}
	}
}
