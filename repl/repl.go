package repl

import (
	"bhasa/evaluator"
	"bhasa/lexer"
	"bhasa/object"
	"bhasa/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

const BANNER = `
╔═══════════════════════════════════════════════════╗
║   ভাষা (Bhasa) - Bengali Programming Language   ║
║              Built with Go                     ║
╚═══════════════════════════════════════════════════╝

Welcome! Type your Bengali code below.
Commands:
  - Type 'প্রস্থান' or 'exit' to quit
  - Use Bengali keywords: ধরি, ফাংশন, যদি, নাহলে, ফেরত
  - Built-in functions: লেখ(), দৈর্ঘ্য(), প্রথম(), শেষ()

Example:
  ধরি x = ৫;
  লেখ(x);

`

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Fprint(out, BANNER)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Exit commands
		if line == "প্রস্থান" || line == "exit" || line == "quit" {
			fmt.Fprintln(out, "আবার দেখা হবে! (Goodbye!)")
			return
		}

		if line == "" {
			continue
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "ত্রুটি (Errors):\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

