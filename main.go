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

    function fib(n) {
        if (n <= 1) {
            return n;
        }
        return fib(n - 2) + fib(n - 1);
    }

    function test(a) {
        if (a <= 1) {
            return a;
        } else {
            return a;
        }
    }

    function sum(a, b) {
        var c = a + b;
        return c;
    }

    var a = sum(1, 2);
    print a;

    var b = test(1);
    print b;

    var c = fib(2);
    print c;
`

	fmt.Println("Input : ", input)

	lexer := lexer.New()
	tokens := lexer.Scan(input)

	parser := parser.New(tokens)
	expressions := parser.Parse()

	interpreter := interpreter.New(expressions)
	interpreter.Start()
}
