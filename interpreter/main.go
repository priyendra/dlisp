package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println("Execute:", input)
	}
}

func main() {
	repl()
}
