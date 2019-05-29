package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/priyendra/dlisp/parser"
)

// TODO(deshwal): Support expressions that evalute to non-integral values.
func eval(expr parser.AstNode) (int, error) {
	switch expr.Type() {
	case parser.AST_LIST:
	case parser.AST_SYMBOL:
	case parser.AST_LITERAL:
		return expr.Literal(), nil
	}
	return -1, errors.New("Cannot eval:\n" + expr.DebugString())
}

// Read-Eval-Print-Loop
func repl() {
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
			val, err := eval(parsed)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}
	}
}

func main() {
	repl()
}
