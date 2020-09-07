package main

import (
	"fmt"

	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	fmt.Println("as")

	lexer := lexer.New()
	tokens := lexer.Scan("Hello")

	for i, v := range tokens {
		fmt.Println(i, v)
	}

	parser.Parse()
}
