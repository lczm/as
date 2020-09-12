package interpreter

import (
	"github.com/lczm/as/ast"
	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

type Interpreter struct {
	// Expressions []ast.Expression
	Statements []ast.Statement
}

func (i *Interpreter) Start() string {
	if len(i.Statements) < 0 {
		panic("Interpreter needs at least one statement to start")
	}

	stmt := i.Statements[0]
	object := i.Eval(stmt)

	return object.String()
}

func (i *Interpreter) Eval(astNode ast.AstNode) object.Object {
	// fmt.Println("Eval")

	switch node := astNode.(type) {
	case *ast.BinaryExpression:
		// fmt.Println(ast.Operator.Literal)
		return i.evalBinaryExpression(node)
	case *ast.UnaryExpression:
		// fmt.Println(ast.Operator.Literal)
		return i.evalUnaryExpression(node)
	case *ast.NumberExpression:
		numberValue := int64(node.Value)
		return &object.Integer{Value: numberValue}
	case *ast.GroupExpression:
		return i.Eval(node.Expr)
	}

	return nil
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

func (i *Interpreter) evalUnaryExpression(expr *ast.UnaryExpression) object.Object {
	right := i.Eval(expr.Right)

	switch expr.Operator.Type {
	case token.MINUS:
		// Inverse the value
		if right.Type() == object.INTEGER {
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: -rightValue}
		}
	}
	return nil
}

func New(statements []ast.Statement) *Interpreter {
	i := &Interpreter{
		Statements: statements,
	}
	return i
}
