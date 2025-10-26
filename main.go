package main

import (
	"bhasa/compiler"
	"bhasa/lexer"
	"bhasa/parser"
	"bhasa/repl"
	"bhasa/vm"
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

	comp := compiler.New()
	err = comp.Compile(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compilation failed:\n %s\n", err)
		os.Exit(1)
	}

	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Executing bytecode failed:\n %s\n", err)
		os.Exit(1)
	}
}

