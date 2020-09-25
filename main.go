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
        } else {
            return fib(n - 2) + fib(n - 1);
        }
    }

    function sum_two(a, b) {
        return a + b;
    }

    function sum_four(a, b, c, d) {
        if (a + b <= 5) {
            return sum_two(10, 10);
        }
        return sum_two(a + b, c + d);
    }

    function test_return(a) {
        if (a <= 5) {
            return 10;
        } 
        return 20;
    }

    var c = fib(10);
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
