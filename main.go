package main

import (
	"fmt"

	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	fmt.Println("as")

	lexer := lexer.New()
	tokens := lexer.Scan("1 + 2")

	// for _, token := range tokens {
	// 	fmt.Println(token.Type, token.Literal)
	// }

	parser := parser.New(tokens)
	parser.Parse()
}
