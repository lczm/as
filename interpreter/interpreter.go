package interpreter

import (
	"fmt"

	"github.com/lczm/as/ast"
	"github.com/lczm/as/environment"
	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

type Interpreter struct {
	Environment *environment.Environment
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
	case *ast.IfStatement:
		return i.evalIfStatement(node)
	case *ast.ForStatement:
		i.evalForStatement(node)
	case *ast.WhileStatement:
		i.evalWhileStatement(node)
	case *ast.BlockStatement:
		return i.evalBlockStatement(node)
	case *ast.FunctionStatement:
		i.evalFunctionStatement(node)
	case *ast.PrintStatement:
		return i.evalPrintStatement(node)
	case *ast.ReturnStatement:
		return i.evalReturnStatement(node)
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
	case *ast.LogicalExpression:
		return i.evalLogicalExpression(node)
	case *ast.NumberExpression:
		numberValue := int64(node.Value)
		return &object.Integer{Value: numberValue}
	case *ast.GroupExpression:
		return i.Eval(node.Expr)
	case *ast.CallExpression:
		return i.evalCallExpression(node)
	}

	return nil
}

func (i *Interpreter) evalIfStatement(stmt *ast.IfStatement) object.Object {
	if i.IsTruthy(i.Eval(stmt.Condition)) {
		return i.Eval(stmt.Then)
	}

	if stmt.Else != nil {
		return i.Eval(stmt.Else)
	}
	return nil
}

func (i *Interpreter) evalForStatement(stmt *ast.ForStatement) {
	// Initialize the variable first.
	i.Eval(stmt.Variable)

	for i.IsTruthy(i.Eval(stmt.Condition)) {
		// Evaluate the body expression
		i.Eval(stmt.Body)

		// Afterwards run the effect
		// This is also where a pre vs post increment can be done.
		i.Eval(stmt.Effect)
	}
}

func (i *Interpreter) evalWhileStatement(stmt *ast.WhileStatement) {
	for i.IsTruthy(i.Eval(stmt.Condition)) {
		i.Eval(stmt.Body)
	}
}

func (i *Interpreter) evalBlockStatement(stmt *ast.BlockStatement) object.Object {
	childEnvironment := environment.NewChildEnvironment(i.Environment)

	return i.ExecuteBlockStatements(stmt.Statements, childEnvironment)
}

func (i *Interpreter) evalFunctionStatement(stmt *ast.FunctionStatement) {
	functionObject := &object.Function{
		FunctionStatement: *stmt,
	}
	i.Environment.Define(stmt.Name.Literal, functionObject)
}

func (i *Interpreter) evalPrintStatement(stmt *ast.PrintStatement) object.Object {
	// This is also a default value for returning values from a print statement.
	// Which allows for code such as `var a = print(3);` to work

	objectValue := i.Eval(stmt.Expr)
	str := objectValue.String()

	fmt.Println(str)
	return nil
}

func (i *Interpreter) evalReturnStatement(stmt *ast.ReturnStatement) object.Object {
	if stmt.Value == nil {
		return nil
	}
	return &object.Return{Value: i.Eval(stmt.Value)}
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
	case token.GT: // Greater than
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue > rightValue}
		}
	case token.GT_EQ: // Greater equal than
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue >= rightValue}
		}
	case token.LT: // Lesser than
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue < rightValue}
		}
	case token.LT_EQ: // Lesser equal than
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue <= rightValue}
		}
	case token.EQ: // Equals '=='
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue == rightValue}
		}
	case token.NOT_EQ: // Not equals '!='
		if left.Type() == object.INTEGER && right.Type() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue != rightValue}
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

func (i *Interpreter) evalLogicalExpression(expr *ast.LogicalExpression) object.Object {
	if expr.Operator.Type == token.AND {
		return &object.Bool{
			Value: i.IsTruthy(i.Eval(expr.Left)) && i.IsTruthy(i.Eval(expr.Right)),
		}
	} else if expr.Operator.Type == token.OR {
		return &object.Bool{
			Value: i.IsTruthy(i.Eval(expr.Left)) || i.IsTruthy(i.Eval(expr.Right)),
		}
	} else {
		panic("LogicalExpression has a operator that is not supported")
	}
}

func (i *Interpreter) evalCallExpression(expr *ast.CallExpression) object.Object {
	function, ok := i.Eval(expr.Callee).(*object.Function)
	if !ok {
		panic("Call expression callee is not a declared function")
	}

	var evaluatedArguments []object.Object
	for _, argument := range expr.Arguments {
		evaluatedArguments = append(evaluatedArguments, i.Eval(argument))
	}

	environment := environment.NewChildEnvironment(i.Environment)
	for i, argument := range evaluatedArguments {
		environment.Define(function.FunctionStatement.Params[i].Literal,
			argument)
	}

	obj := i.ExecuteBlockStatements(function.FunctionStatement.Body.Statements, environment)
	// If the object is a return value
	returnObj, ok := obj.(*object.Return)
	if ok {
		return returnObj.Value
	}
	return obj
}

// ---  Utility functions
// This function will take in an environment as a block is scoped
// to it's own environment.
func (i *Interpreter) ExecuteBlockStatements(
	statements []ast.Statement,
	environment *environment.Environment) object.Object {

	// Go does everything by value and not reference so this is fine.
	previousEnvironment := i.Environment
	i.Environment = environment

	for index, stmt := range statements {
		obj := i.Eval(stmt)
		fmt.Println("Hello", index, obj.String(), obj.Type())

		// If the object returned from evaluation is a return object
		// break out of the evaluation loop and return it.
		returnObj, ok := obj.(*object.Return)
		if ok {
			// Reset the environment back to the previous one.
			i.Environment = previousEnvironment
			return returnObj
		}
	}

	i.Environment = previousEnvironment
	return nil
}

// This is where it is important to define what is truthy and what is not.
func (i *Interpreter) IsTruthy(obj object.Object) bool {
	// Check for booleans
	if obj.Type() == object.BOOL {
		return obj.(*object.Bool).Value
	}

	if obj.Type() == object.INTEGER {
		// If the value of the object is 0, it is falsey
		if obj.(*object.Integer).Value == 0 {
			return false
		}
		// If the value of the object is not 0, it is truthy
		return true
	}

	return false
}

func New(statements []ast.Statement) *Interpreter {
	environment := environment.New()

	i := &Interpreter{
		Statements:  statements,
		Environment: environment,
	}
	return i
}
