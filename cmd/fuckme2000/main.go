package main

import (
	"fmt"
	"os"
	"fuckme2000/internal/lexer"
	"fuckme2000/internal/parser"
	"fuckme2000/internal/vm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fuckme2000 <file.f2000>")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Step 1: Lexing
	l := lexer.New(string(source))
	var tokens []lexer.Token
	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}

	// Step 2: Parsing
	p := parser.New(tokens)
	instructions := p.ParseProgram()

	// Step 3: Executing
	machine := vm.New()
	machine.Run(instructions)
}
