package main

import (
	"fmt"
	"github.com/lczm/as/ast"
	"github.com/lczm/as/interpreter"
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
	expressions := parser.Parse()

	fmt.Println("Length : ", len(expressions))
	expr, _ := expressions[0].(*ast.BinaryExpression)

	fmt.Println("Expr : ", expr.String())
	fmt.Println("Left : ", expr.Left.String())
	fmt.Println("Right : ", expr.Right.String())
	fmt.Println("Operator : ", expr.Operator.Literal)

	interpreter := interpreter.New(expressions)
	interpreter.Start()
}
