package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/priyendra/dlisp/builtins"
	"github.com/priyendra/dlisp/expression"
)

type StdEnvironment struct {
	vars map[string]expression.Expression
}

func (env StdEnvironment) LookupVar(name string) (
	expression.Expression, error) {
	val, ok := env.vars[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Undefined value: '%s'", name))
	}
	return val, nil
}

func (env *StdEnvironment) SetVar(name string, value expression.Expression) {
	env.vars[name] = value
}

func stdEnv() expression.Environment {
	env := StdEnvironment{map[string]expression.Expression{}}

	// List functions
	env.vars["append"] = builtins.Append
	env.vars["car"] = builtins.Car
	env.vars["cdr"] = builtins.Cdr
	env.vars["cons"] = builtins.Cons
	env.vars["len"] = builtins.Len
	env.vars["map"] = builtins.Map

	// Math functions
	env.vars["abs"] = builtins.Abs
	env.vars["acos"] = builtins.Acos
	env.vars["acosh"] = builtins.Acosh
	env.vars["asin"] = builtins.Asin
	env.vars["asinh"] = builtins.Asinh
	env.vars["atan"] = builtins.Atan
	env.vars["atanh"] = builtins.Atanh
	env.vars["cbrt"] = builtins.Cbrt
	env.vars["ceil"] = builtins.Ceil
	env.vars["cos"] = builtins.Cos
	env.vars["cosh"] = builtins.Cosh
	env.vars["erf"] = builtins.Erf
	env.vars["erfc"] = builtins.Erfc
	env.vars["erfcinv"] = builtins.Erfcinv
	env.vars["erfinv"] = builtins.Erfinv
	env.vars["exp"] = builtins.Exp
	env.vars["exp2"] = builtins.Exp2
	env.vars["floor"] = builtins.Floor
	env.vars["gamma"] = builtins.Gamma
	env.vars["log"] = builtins.Log
	env.vars["log10"] = builtins.Log10
	env.vars["log2"] = builtins.Log2
	env.vars["round"] = builtins.Round
	env.vars["sqrt"] = builtins.Sqrt
	env.vars["trunc"] = builtins.Trunc
	env.vars["max"] = builtins.Max
	env.vars["min"] = builtins.Min
	env.vars["mod"] = builtins.Mod
	env.vars["sin"] = builtins.Sin
	env.vars["sinh"] = builtins.Sinh
	env.vars["tan"] = builtins.Tan
	env.vars["tanh"] = builtins.Tanh
	env.vars["+"] = builtins.Plus
	env.vars["-"] = builtins.Minus
	env.vars["*"] = builtins.Multiply
	env.vars["/"] = builtins.Divide
	env.vars["%"] = builtins.Mod

	// Relational & logical operators
	env.vars[">"] = builtins.GreaterThan
	env.vars[">="] = builtins.GreaterEqual
	env.vars["<"] = builtins.LessThan
	env.vars["<="] = builtins.LessEqual
	env.vars["="] = builtins.Equal
	env.vars["!="] = builtins.NotEqual
	env.vars["&&"] = builtins.LogicalAnd
	env.vars["||"] = builtins.LogicalOr
	env.vars["!"] = builtins.LogicalNot
	env.vars["if"] = builtins.If
	env.vars["true"] = expression.Bool(true)
	env.vars["false"] = expression.Bool(false)

	// General functions
	env.vars["pi"] = expression.Float(math.Pi)
	env.vars["identity"] = builtins.Identity
	return &env
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
			val, err := expression.Eval(expr, env)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(expression.ToString(val))
			}
		}
	}
}
