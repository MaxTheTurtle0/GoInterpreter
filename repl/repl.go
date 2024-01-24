package repl

import (
	"bufio"
	"fmt"
    "go/interpreter/evaluator"	
    "go/interpreter/lexer"
	"go/interpreter/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Printf(PROMPT) 
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        line := scanner.Text()

        if line == ".quit" {
            return
        } 

        l := lexer.New(line)
        p := parser.New(l)
        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors()) 
        }
       
        evaluated := evaluator.Eval(program)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}

func printParserErrors(out io.Writer, errors []string) {
    io.WriteString(out, "Woops! We ran into some error here!\n")
    io.WriteString(out, " parser errors:\n")
    
    for _, msg := range errors {
        io.WriteString(out, "\t"+msg+"\n")
    }
}
