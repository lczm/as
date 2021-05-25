package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/lczm/as/analysis"
	"github.com/lczm/as/globals"
	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	// If there are no arguments passed into the binary, don't prompt a help
	// message and exit from the program, don't continue to do anything
	if len(os.Args) == 1 {
		userOs := runtime.GOOS
		fmt.Println("No files found.")
		switch userOs {
		case "windows": // Windows
			fmt.Println("Usage: as {file}")
		default: // Mac, Linux
			fmt.Println("Usage: ./as {file}")
		}
		os.Exit(0)
	}

	// Grab all the arguments
	arguments := os.Args[1:]
	if len(arguments) > 1 {
		os.Exit(1)
	}

	name := arguments[0]
	data, _ := ioutil.ReadFile(name)

	input := string(data)

	// Lex the program into tokens
	lexer := lexer.New()
	tokens := lexer.Scan(input)

	// Parse the tokens into an AST of statements
	parser := parser.New(tokens)
	statements := parser.Parse()

	// Analyze the values
	semanticAnalyzer := analysis.New(statements)
	semanticAnalyzer.Analyze()

	// TODO : if it is more than 0, and there is some form of strict flag
	// this should not continue running
	if len(globals.ErrorList) > 0 {
		// If there are any errors that are detected
		for _, error := range globals.ErrorList {
			error.Describe()
		}
		// TODO : Find the correct error code to exit from an error
		os.Exit(1)
	}

	// TODO : Some form of flag to determine whether this should be continued or not
	if len(globals.WarningList) > 0 {
		for _, warning := range globals.WarningList {
			warning.Describe()
		}
		// If there is a flag to determine that this should not be continued;
		// then this should exited
		// os.Exit(1)
	}

	interpreter := interpreter.New(statements)
	interpreter.Start()
}
