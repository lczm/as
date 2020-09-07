package main

import (
	"fmt"

	"github.com/lczm/as/lexer"
)

func main() {
	fmt.Println("as")

	lexer := lexer.New()
	tokens := lexer.Scan(">=")

	for _, token := range tokens {
		fmt.Println(token.Type, token.Literal)
	}
}
