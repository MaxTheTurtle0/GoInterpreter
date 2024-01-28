package repl

import (
	"bufio"
	"fmt"
	"go/interpreter/evaluator"
	"go/interpreter/lexer"
	"go/interpreter/object"
	"go/interpreter/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    env := object.NewEnvironment()
    
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
            p.PrintParserErrors(out) 
        }
       
        evaluated := evaluator.Eval(program, env)
        if evaluated != nil {
            io.WriteString(out, evaluated.Inspect())
            io.WriteString(out, "\n")
        }
    }
}
