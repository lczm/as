package main

import (
	"fmt"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	fmt.Println("as")

	input := `
    for (var i = 0; i < 3; i = i + 1) {
    }
    print i;
`

	fmt.Println("Input : ", input)

	lexer := lexer.New()
	tokens := lexer.Scan(input)

	parser := parser.New(tokens)
	expressions := parser.Parse()

	interpreter := interpreter.New(expressions)
	interpreter.Start()
}
