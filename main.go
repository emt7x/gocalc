package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Read the expression from the user.
		fmt.Print("Expression -> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		// Create a new lexer from the provided expression.
		lexer := NewLexer(strings.Split(scanner.Text(), ""))
		if err := lexer.Tokenise(); err != nil {
			fmt.Printf("lexer: %v\r\n", err)
			continue
		}

		// Create a new parser with the tokens gathered from the lexer.
		parser := NewParser(lexer)
		node, err := parser.Parse()
		if err != nil {
			fmt.Printf("parser: %v\r\n", err)
			continue
		}

		// Evaluate the expression node that has been formatted by the AST.
		evaluator := NewEvaluator(node)
		value, err := evaluator.Evaluate()
		if err != nil {
			fmt.Printf("evaluator: %v\r\n", err)
			continue
		}

		// Display the final outputs in the terminal.
		fmt.Println("    AST Result -> " + node.String())
		fmt.Println("    Solution   -> " + value.String())
	}
}
