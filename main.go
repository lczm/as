package main

import (
	"io/ioutil"
	"os"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
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
