package interpreter

import (
	"fmt"

	"github.com/lczm/as/ast"
	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

type Interpreter struct {
	Expressions []ast.Expression
}

func (i *Interpreter) Start() {
	if len(i.Expressions) < 0 {
		panic("Interpreter needs at least one expression to start")
	}

	expr := i.Expressions[0]
	object := i.Eval(expr)
	fmt.Println(object.String())
}

func (i *Interpreter) Eval(expr ast.Expression) object.Object {
	// fmt.Println("Eval")

	switch ast := expr.(type) {
	case *ast.BinaryExpression:
		// fmt.Println(ast.Operator.Literal)
		return i.evalBinaryExpression(ast)
	case *ast.UnaryExpression:
		// fmt.Println(ast.Operator.Literal)
	case *ast.NumberExpression:
		numberValue := int64(ast.Value)
		return &object.Integer{Value: numberValue}
	case *ast.GroupExpression:
		fmt.Println("Eval : GroupExpression")
	}

	return &object.Integer{Value: 2}
}

func (i *Interpreter) evalBinaryExpression(expr *ast.BinaryExpression) object.Object {
	left := i.Eval(expr.Left)
	right := i.Eval(expr.Right)

	switch expr.Operator.Type {
	case token.PLUS: // Add
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: leftValue + rightValue}
		}
	case token.MINUS: // Subtract
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: leftValue - rightValue}
		}
	case token.ASTERISK: // Multiply
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			return &object.Integer{Value: leftValue * rightValue}
		}
	case token.SLASH: // Divide
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			return &object.Integer{Value: leftValue / rightValue}
		}
	}

	return nil
}

func New(expressions []ast.Expression) *Interpreter {
	i := &Interpreter{
		Expressions: expressions,
	}
	return i
}
