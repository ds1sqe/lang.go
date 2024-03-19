package repl

import (
	"bufio"
	"fmt"
	"io"
	"mylang/evaluator"
	"mylang/lexer"
	"mylang/parser"
)

const PROMPT = "-> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")

		}

	}
}

func printParserErrors(out io.Writer, errors []parser.Error) {

	io.WriteString(out, "Error Occur.\n")
	io.WriteString(out, "parser errors:\n")
	for _, err := range errors {
		io.WriteString(out, "\t"+string(err)+"\n")
	}
}
