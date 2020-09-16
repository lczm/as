package main

import (
	"fmt"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func main() {
	fmt.Println("as")

	// input := "print (5 * 2);"
	// input := "var a = 3; print a;"
	input := `
	var a = 3;
	var b = 5;
	print a;
	print b;
	print a + b;
    a = 6;
    print a;

    {
        print 1;
        print 1 + 2;
    }

    if (1) {
        print 500;
    }
`

	fmt.Println("Input : ", input)

	lexer := lexer.New()
	tokens := lexer.Scan(input)

	parser := parser.New(tokens)
	expressions := parser.Parse()

	interpreter := interpreter.New(expressions)
	interpreter.Start()
}
