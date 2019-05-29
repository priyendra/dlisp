package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/priyendra/dlisp/builtins"
	"github.com/priyendra/dlisp/parser"
	"github.com/priyendra/dlisp/value"
)

type Environment struct {
	names map[string]value.Value
}

func stdEnv() Environment {
	env := Environment{
		map[string]value.Value{},
	}
	env.names["identity"] = builtins.Identity
	return env
}

// TODO(deshwal): Support expressions that evalute to non-integral values.
func eval(env Environment, expr *parser.AstNode) (value.Value, error) {
	switch expr.Type() {
	case parser.AST_SYMBOL:
		val, ok := env.names[expr.Symbol()]
		if !ok {
			return nil, errors.New(fmt.Sprintf("Symbol '%s' undefined", expr.Symbol()))
		}
		return val, nil
	case parser.AST_INT_LITERAL:
		return value.Int(expr.Literal()), nil
	case parser.AST_LIST:
		if expr.NumChildren() == 0 {
			return nil, errors.New("Cannot eval empty list")
		}
		child0, err := eval(env, expr.Child(0))
		if err != nil {
			return nil, errors.New("Could not eval first child of list expression")
		}
		if value.ToType(child0) != value.FUNCTION {
			return nil, errors.New(
				"First child of list function must eval to type function")
		}
		args := make([]value.Value, expr.NumChildren()-1)
		for i := 1; i < expr.NumChildren(); i++ {
			args[i-1], err = eval(env, expr.Child(1))
			if err != nil {
				return nil, errors.New(
					fmt.Sprintf("Could not eval %d-th arg of list expression", i-1))
			}
		}
		return child0.(value.Function).Eval(args)
	}
	return value.Int(-1), errors.New("Cannot eval:\n" + expr.DebugString())
}

// Read-Eval-Print-Loop
func repl() {
	env := stdEnv()
	scanner := bufio.NewScanner(os.Stdin)
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
		parsed, err := parser.Parse(input)
		if err != nil {
			fmt.Println(err)
		} else {
			val, err := eval(env, &parsed)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value.ToString(val))
			}
		}
	}
}

func main() {
	repl()
}
