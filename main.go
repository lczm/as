package main

import (
	"fmt"
	"time"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	fmt.Println("as")

	input := `

    function fib(n) {
        if (n <= 1) {
            return n;
        }
        return fib(n - 2) + fib(n - 1);
    }

    var c = fib(20);
    print c;
`

	fmt.Println("Input : ", input)

	lexer := lexer.New()
	tokens := lexer.Scan(input)

	parser := parser.New(tokens)
	expressions := parser.Parse()

	start := time.Now()
	interpreter := interpreter.New(expressions)
	interpreter.Start()

	fmt.Println("Time taken : ", time.Since(start))
}
