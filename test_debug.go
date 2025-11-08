package main

import (
	"bhasa/compiler"
	"bhasa/lexer"
	"bhasa/parser"
	"bhasa/vm"
	"fmt"
)

func main() {
	input := `
লেখ("Test 1");
শ্রেণী Test {
    সার্বজনীন নাম: পাঠ্য;
    
    সার্বজনীন নির্মাতা() {
        লেখ("Constructor called");
    }
}
লেখ("Test 2");
ধরি t = নতুন Test();
লেখ("Test 3 - success!");
`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			fmt.Println(err)
		}
		return
	}
	
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Println("Compilation error:", err)
		return
	}
	
	bytecode := comp.Bytecode()
	
	machine := vm.New(bytecode)
	err = machine.Run()
	if err != nil {
		fmt.Println("Executing bytecode failed:\n", err)
	}
}
