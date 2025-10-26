package main

import (
	"bhasa/evaluator"
	"bhasa/lexer"
	"bhasa/object"
	"bhasa/parser"
	"bhasa/repl"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// Start REPL if no file is provided
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Run file
	filename := os.Args[1]
	runFile(filename)
}

func runFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Fprintln(os.Stderr, "Parser errors:")
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, "\t%s\n", msg)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		if evaluated.Type() == object.ERROR_OBJ {
			fmt.Fprintf(os.Stderr, "%s\n", evaluated.Inspect())
			os.Exit(1)
		}
	}
}

