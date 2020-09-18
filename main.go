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
    var jasper = 0;
    while (jasper < 10) {
        print 123;
        jasper = jasper + 1;
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

// 	var a = 3;
// 	var b = 5;
// 	print a;
// 	print b;
// 	print a + b;
//     a = 6;
//     print a;
//
//     {
//         print 1;
//         print 1 + 2;
//     }
//
//     if (1) {
//         print 500;
//     }
//
//     if (2 > 1) {
//         print 600;
//     }
//
//     if (1 > 2) {
//         print 700;
//     } else {
//         print 1000;
//     }
//
//     if (5 >= 5) {
//         print 1000000;
//     }
