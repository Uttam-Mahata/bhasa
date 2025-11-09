package main

import (
	"bhasa/compiler"
	"bhasa/lexer"
	"bhasa/parser"
	"bhasa/repl"
	"bhasa/vm"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define CLI flags
	compileMode := flag.Bool("c", false, "Compile source to bytecode")
	outputFile := flag.String("o", "", "Output file for compiled bytecode")
	showHelp := flag.Bool("h", false, "Show help message")
	showVersion := flag.Bool("v", false, "Show version information")

	flag.Parse()

	// Show help
	if *showHelp {
		printHelp()
		return
	}

	// Show version
	if *showVersion {
		fmt.Println("Bhasa (ভাষা) Programming Language v1.0.0")
		fmt.Println("Bytecode Compiler & VM")
		return
	}

	// Get remaining arguments (non-flag arguments)
	args := flag.Args()

	if len(args) < 1 {
		// Start REPL if no file is provided
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	filename := args[0]

	// Check if file is bytecode or source
	if isBytecodeFile(filename) {
		// Execute pre-compiled bytecode
		runBytecode(filename)
	} else if *compileMode {
		// Compile source to bytecode
		compileFile(filename, *outputFile)
	} else {
		// Run source file directly
		runFile(filename)
	}
}

func printHelp() {
	fmt.Println("Bhasa (ভাষা) Programming Language - Usage:")
	fmt.Println()
	fmt.Println("  bhasa                         Start REPL (interactive mode)")
	fmt.Println("  bhasa <file>                  Run source file (.bhasa or .ভাষা)")
	fmt.Println("  bhasa <bytecode>              Execute bytecode file (.compiled or .সংকলিত)")
	fmt.Println("  bhasa -c <file>               Compile source to bytecode")
	fmt.Println("  bhasa -c -o <output> <file>   Compile with custom output name")
	fmt.Println("  bhasa -h                      Show this help message")
	fmt.Println("  bhasa -v                      Show version information")
	fmt.Println()
	fmt.Println("File Extensions:")
	fmt.Println("  Source:    .bhasa or .ভাষা")
	fmt.Println("  Bytecode:  .compiled or .সংকলিত")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  bhasa program.bhasa                   # Run source file")
	fmt.Println("  bhasa -c program.bhasa                # Compile to program.compiled")
	fmt.Println("  bhasa -c -o output.compiled program.bhasa")
	fmt.Println("  bhasa -c -o output.সংকলিত program.ভাষা   # Bengali extensions")
	fmt.Println("  bhasa program.compiled                # Execute compiled bytecode")
	fmt.Println()
	fmt.Println("Note: Flags must appear before the filename argument")
}

// isBytecodeFile checks if the file is a bytecode file based on extension
func isBytecodeFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".compiled" || ext == ".সংকলিত"
}

// runFile compiles and runs a source file
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

// compileFile compiles a source file to bytecode
func compileFile(filename string, outputFile string) {
	// Read source file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse source
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

	// Compile to bytecode
	comp := compiler.New()
	err = comp.Compile(program)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compilation failed:\n %s\n", err)
		os.Exit(1)
	}

	// Determine output filename
	if outputFile == "" {
		// Default: replace extension with .compiled
		ext := filepath.Ext(filename)
		outputFile = strings.TrimSuffix(filename, ext) + ".compiled"
	}

	// Create output file
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Serialize bytecode to file
	bytecode := comp.Bytecode()
	err = bytecode.Serialize(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error serializing bytecode: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled %s to %s\n", filename, outputFile)
}

// runBytecode executes a pre-compiled bytecode file
func runBytecode(filename string) {
	// Open bytecode file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bytecode file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Deserialize bytecode
	bytecode, err := compiler.Deserialize(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deserializing bytecode: %v\n", err)
		os.Exit(1)
	}

	// Execute bytecode in VM
	machine := vm.New(bytecode)
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Executing bytecode failed:\n %s\n", err)
		os.Exit(1)
	}
}
