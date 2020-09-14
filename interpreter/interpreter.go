package interpreter

import (
	"fmt"

	"github.com/lczm/as/ast"
	"github.com/lczm/as/environment"
	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

type Interpreter struct {
	Environment environment.Environment
	Statements  []ast.Statement
}

func (i *Interpreter) Start() {
	if len(i.Statements) < 0 {
		panic("Interpreter needs at least one statement to start")
	}

	for _, stmt := range i.Statements {
		i.Eval(stmt)
	}
}

// Eval has to take in an astNode and not an ast.Statement because
// this function will have to run recursively and deal with
// ast.Expression at times.
func (i *Interpreter) Eval(astNode ast.AstNode) object.Object {
	switch node := astNode.(type) {
	case *ast.StatementExpression:
		return i.Eval(node.Expr)
	case *ast.PrintStatement:
		return i.evalPrintStatement(node)
	case *ast.VariableStatement:
		i.evalVariableStatement(node)
	case *ast.VariableExpression:
		return i.evalVariableExpression(node)
	case *ast.AssignmentExpression:
		return i.evalAssignmentExpression(node)
	case *ast.BinaryExpression:
		return i.evalBinaryExpression(node)
	case *ast.UnaryExpression:
		return i.evalUnaryExpression(node)
	case *ast.NumberExpression:
		numberValue := int64(node.Value)
		return &object.Integer{Value: numberValue}
	case *ast.GroupExpression:
		return i.Eval(node.Expr)
	}

	return nil
}

func (i *Interpreter) evalPrintStatement(stmt *ast.PrintStatement) object.Object {
	// This is also a default value for returning values from a print statement.
	// Which allows for code such as `var a = print(3);` to work

	objectValue := i.Eval(stmt.Expr)
	fmt.Println(objectValue.String())
	return nil
}

func (i *Interpreter) evalVariableStatement(stmt *ast.VariableStatement) {
	// Create a default object, this also defines what a variable value/type will
	// be when it is not initialized
	// `var a;`, 'a' will be defined to what the default object is.
	// In this case, I think setting it to a 0 integer should be fine for now.
	// Perhaps in the future, this can be changed to a null value of some sort.
	if stmt.Initializer != nil {
		initializerValue := i.Eval(stmt.Initializer)
		i.Environment.Define(stmt.Name.Literal, initializerValue)
	} else {
		i.Environment.Define(stmt.Name.Literal, &object.Integer{Value: 0})
	}
}

func (i *Interpreter) evalVariableExpression(expr *ast.VariableExpression) object.Object {
	return i.Environment.Get(expr.Name.Literal)
}

func (i *Interpreter) evalAssignmentExpression(expr *ast.AssignmentExpression) object.Object {
	value := i.Eval(expr.Value)
	i.Environment.Set(expr.Name.Literal, value)

	return value
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
	environment := environment.New()

	i := &Interpreter{
		Statements:  statements,
		Environment: *environment,
	}
	return i
}
