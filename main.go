package main

import (
	"bufio"
	"fmt"
	"go/interpreter/evaluator"
	"go/interpreter/lexer"
	"go/interpreter/object"
	"go/interpreter/parser"
	"go/interpreter/repl"
	"os"
	"os/user"
)

func main() {
    // Check if the user has provided a file path as an argument
    if len(os.Args) > 2 {
        // If the user has provided more than one argument, print a message and exit
        fmt.Printf("Either run this program without arguments or with the a valid input file\n")
        os.Exit(0)
    } else if len(os.Args) == 2 { 
        // If the user has provided a file path, run the interpreter

        filePath := os.Args[1]
        
        // check for right file extension
        if filePath[len(filePath)-7:] != ".turtls" {
            fmt.Printf("Please provide a file with the .turtls extension\n")
            os.Exit(0)
        }

        // Open the file
        file, err := os.Open(filePath)
        if err != nil {
            fmt.Println("Error opening file:", err)
            return
        }
        defer file.Close()

        // Create a scanner to read the file line by line
        scanner := bufio.NewScanner(file)

        // Create a variable to store the file content
        var fileContent string

        // Iterate over each line and concatenate it to the fileContent string
        for scanner.Scan() {
            fileContent += scanner.Text() + "\n"
        } 

        // Check for scanner errors
        if err := scanner.Err(); err != nil {
            fmt.Println("Error reading file:", err)
            return
        }

        env := object.NewEnvironment()
        l := lexer.New(fileContent)
        p := parser.New(l)

        program := p.ParseProgram()

        if len(p.Errors()) != 0 {
            p.PrintParserErrors(os.Stdout) 
        }

        evaluator.Eval(program, env)
        os.Exit(0)
    }

    // If the user has not provided any arguments, start the REPL
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! Welcome to TurtlScript!\n", user.Username)
    fmt.Printf("Type .quit to quit.\n")
    repl.Start(os.Stdin, os.Stdout)
}
