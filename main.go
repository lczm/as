package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

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

	lexer := lexer.New()
	tokens := lexer.Scan(input)

	parser := parser.New(tokens)
	expressions := parser.Parse()

	interpreter := interpreter.New(expressions)
	interpreter.Start()
}
