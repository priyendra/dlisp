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
	env.Names["append"] = builtins.Append
	env.Names["car"] = builtins.Car
	env.Names["cdr"] = builtins.Cdr
	env.Names["cons"] = builtins.Cons
	env.Names["len"] = builtins.Len

	// Match functions
	env.Names["abs"] = builtins.Abs
	env.Names["acos"] = builtins.Acos
	env.Names["acosh"] = builtins.Acosh
	env.Names["asin"] = builtins.Asin
	env.Names["asinh"] = builtins.Asinh
	env.Names["atan"] = builtins.Atan
	env.Names["atanh"] = builtins.Atanh
	env.Names["cbrt"] = builtins.Cbrt
	env.Names["ceil"] = builtins.Ceil
	env.Names["cos"] = builtins.Cos
	env.Names["cosh"] = builtins.Cosh
	env.Names["erf"] = builtins.Erf
	env.Names["erfc"] = builtins.Erfc
	env.Names["erfcinv"] = builtins.Erfcinv
	env.Names["erfinv"] = builtins.Erfinv
	env.Names["exp"] = builtins.Exp
	env.Names["exp2"] = builtins.Exp2
	env.Names["floor"] = builtins.Floor
	env.Names["gamma"] = builtins.Gamma
	env.Names["log"] = builtins.Log
	env.Names["log10"] = builtins.Log10
	env.Names["log2"] = builtins.Log2
	env.Names["round"] = builtins.Round
	env.Names["sqrt"] = builtins.Sqrt
	env.Names["trunc"] = builtins.Trunc
	env.Names["max"] = builtins.Max
	env.Names["min"] = builtins.Min
	env.Names["mod"] = builtins.Mod
	env.Names["sin"] = builtins.Sin
	env.Names["sinh"] = builtins.Sinh
	env.Names["tan"] = builtins.Tan
	env.Names["tanh"] = builtins.Tanh
	env.Names["+"] = builtins.Plus
	env.Names["-"] = builtins.Minus
	env.Names["*"] = builtins.Multiply
	env.Names["/"] = builtins.Divide
	env.Names["%"] = builtins.Mod

	// Relational & logical operators
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
			val, err := expression.Eval(expr, &env)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression.ToString(val))
			}
		}
	}
}
