package interpreter

import (
	"fmt"
	"github.com/lczm/as/ast"
)

type Interpreter struct {
	Expressions []ast.Expression
}

func (i *Interpreter) Start() {
	if len(i.Expressions) < 0 {
		panic("Interpreter needs at least one expression to start")
	}

	expr := i.Expressions[0]
	i.Eval(expr)
}

func (i *Interpreter) Eval(expr ast.Expression) {
	fmt.Println("Eval")

	switch ast := expr.(type) {
	case *ast.BinaryExpression:
		fmt.Println(ast.Operator.Literal)
		fmt.Println("Eval : BinaryExpression")
	case *ast.UnaryExpression:
		fmt.Println(ast.Operator.Literal)
		fmt.Println("Eval : UnaryExpression")
	case *ast.NumberExpression:
		fmt.Println("Eval : NumberExpression")
	case *ast.GroupExpression:
		fmt.Println("Eval : GroupExpression")
	}
}

func New(expressions []ast.Expression) *Interpreter {
	i := &Interpreter{
		Expressions: expressions,
	}
	return i
}
