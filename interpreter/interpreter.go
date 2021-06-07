package interpreter

import (
	"github.com/lczm/as/ast"
	"github.com/lczm/as/builtin"
	"github.com/lczm/as/environment"
	"github.com/lczm/as/errors"
	"github.com/lczm/as/globals"
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
	case *ast.StructStatement:
		i.evalStructStatement(node)
	case *ast.ReturnStatement:
		return i.evalReturnStatement(node)
	case *ast.VariableStatement:
		i.evalVariableStatement(node)
	case *ast.VariableExpression:
		return i.evalVariableExpression(node)
	case *ast.AssignmentExpression:
		return i.evalAssignmentExpression(node)
	case *ast.AssignmentIndexExpression:
		return i.evalAssignmentIndexExpression(node)
	case *ast.AssignmentStruct:
		return i.evalAssignmentStruct(node)
	case *ast.BinaryExpression:
		return i.evalBinaryExpression(node)
	case *ast.UnaryExpression:
		return i.evalUnaryExpression(node)
	case *ast.LogicalExpression:
		return i.evalLogicalExpression(node)
	case *ast.NumberExpression:
		return &object.Integer{Value: int64(node.Value)}
	case *ast.ListExpression:
		return i.evalListExpression(node)
	case *ast.HashMapExpression:
		return i.evalHashMapExpression(node)
	case *ast.StringExpression:
		return &object.String{Value: node.Value}
	case *ast.BoolExpression:
		return &object.Bool{Value: node.Value}
	case *ast.GroupExpression:
		return i.Eval(node.Expr)
	case *ast.CallExpression:
		return i.evalCallExpression(node)
	case *ast.GetExpression:
		return i.evalGetExpression(node)
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

func (i *Interpreter) evalStructStatement(stmt *ast.StructStatement) {
	attributes := make(map[string]object.Object)
	methods := make(map[string]object.Object)

	for attributeStmt := range stmt.Attributes {
		attributes[attributeStmt.Literal] = i.Eval(attributeStmt)
	}
	for methodStmt := range stmt.Methods {
		attributes[methodStmt.Literal] = i.Eval(methodStmt)
	}

	structObject := &object.Struct{
		Name:       stmt.Name.Literal,
		Attributes: attributes,
		Methods:    methods,
	}
	i.Environment.Define(stmt.Name.Literal, structObject)
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

func (i *Interpreter) evalAssignmentIndexExpression(expr *ast.AssignmentIndexExpression) object.Object {
	value := i.Eval(expr.Value)
	index := i.Eval(expr.Index)

	i.Environment.SetIndex(expr.Name.Literal, index, value)
	return value
}

func (i *Interpreter) evalAssignmentStruct(expr *ast.AssignmentStruct) object.Object {
	value := i.Eval(expr.Value)
	// Need to convert from a generic 'Expression' into a ast.VariableExpression
	// to access Name.Literal
	variableExpression := expr.Attribute.(*ast.VariableExpression)
	i.Environment.SetStruct(expr.Name.Literal, variableExpression.Name.Literal, value)
	return value
}

func (i *Interpreter) evalBinaryExpression(expr *ast.BinaryExpression) object.Object {
	left := i.Eval(expr.Left)
	right := i.Eval(expr.Right)

	switch expr.Operator.Type {
	case token.PLUS: // Add
		// Integers
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: leftValue + rightValue}
		}
		// Strings
		if left.RawType() == object.STRING && right.RawType() == object.STRING {
			leftValue := left.(*object.String).Value
			rightValue := right.(*object.String).Value
			return &object.String{Value: leftValue + rightValue}
		}
	case token.MINUS: // Subtract
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: leftValue - rightValue}
		}
	case token.ASTERISK: // Multiply
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			return &object.Integer{Value: leftValue * rightValue}
		}
	case token.SLASH: // Divide
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			return &object.Integer{Value: leftValue / rightValue}
		}
	case token.MODULUS: // Modulus
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			return &object.Integer{Value: leftValue % rightValue}
		}
	case token.GT: // Greater than
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue > rightValue}
		}
	case token.GT_EQ: // Greater equal than
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue >= rightValue}
		}
	case token.LT: // Lesser than
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue < rightValue}
		}
	case token.LT_EQ: // Lesser equal than
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue <= rightValue}
		}
	case token.EQ: // Equals '=='
		// Integers
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue == rightValue}
		}
		// Strings
		if left.RawType() == object.STRING && right.RawType() == object.STRING {
			leftValue := left.(*object.String).Value
			rightValue := right.(*object.String).Value
			return &object.Bool{Value: leftValue == rightValue}
		}
		// Bools
		if left.RawType() == object.BOOL && right.RawType() == object.BOOL {
			leftValue := left.(*object.Bool).Value
			rightValue := right.(*object.Bool).Value
			return &object.Bool{Value: leftValue == rightValue}
		}
	case token.NOT_EQ: // Not equals '!='
		// Integers
		if left.RawType() == object.INTEGER && right.RawType() == object.INTEGER {
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value
			return &object.Bool{Value: leftValue != rightValue}
		}
		// Strings
		if left.RawType() == object.STRING && right.RawType() == object.STRING {
			leftValue := left.(*object.String).Value
			rightValue := right.(*object.String).Value
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
		if right.RawType() == object.INTEGER {
			rightValue := right.(*object.Integer).Value
			return &object.Integer{Value: -rightValue}
		}
	case token.BANG:
		// If evaluated condition is true, inverse the result
		if i.IsTruthy(right) {
			return &object.Bool{Value: false}
		}
		// If it is not true, inverse the result here
		return &object.Bool{Value: true}
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

func (i *Interpreter) evalListExpression(expr *ast.ListExpression) object.Object {
	var evaluatedExpressions []object.Object
	for _, expression := range expr.Values {
		evaluatedExpressions = append(evaluatedExpressions, i.Eval(expression))
	}
	return &object.List{
		Value: evaluatedExpressions,
	}
}

func (i *Interpreter) evalHashMapExpression(expr *ast.HashMapExpression) object.Object {
	var hashMap map[object.HashKey]object.HashValue
	hashMap = make(map[object.HashKey]object.HashValue)
	for k, v := range expr.Values {
		evaluatedKey := i.Eval(k)
		evaluatedKeyHashable, ok := evaluatedKey.(object.Hashable)
		if !ok {
			panic("Object is not hashable")
		}

		evaluatedValue := i.Eval(v)

		objHash := evaluatedKeyHashable.Hash()
		hashValue := object.HashValue{
			Key:   evaluatedKey,
			Value: evaluatedValue,
		}
		hashMap[objHash] = hashValue
	}
	return &object.HashMap{
		Value: hashMap,
	}
}

func (i *Interpreter) evalCallExpression(expr *ast.CallExpression) object.Object {
	switch callee := i.Eval(expr.Callee).(type) {
	// If it is a function that the user has defined somewhere,
	// evaluate the arguments in the environment and pass the
	// arguments over to the function
	case *object.Function:
		// function, ok := i.Eval(expr.Callee).(*object.Function)
		// if !ok {
		// 	panic("Call expression callee is not a declared function")
		// }

		var evaluatedArguments []object.Object
		for _, argument := range expr.Arguments {
			evaluatedArguments = append(evaluatedArguments, i.Eval(argument))
		}

		environment := environment.NewChildEnvironment(i.Environment)
		for i, argument := range evaluatedArguments {
			environment.Define(callee.FunctionStatement.Params[i].Literal,
				argument)
		}

		obj := i.ExecuteBlockStatements(callee.FunctionStatement.Body.Statements, environment)
		// If the object is a return value
		returnObj, ok := obj.(*object.Return)
		if ok {
			return returnObj.Value
		}
		return obj
		// If it is a builtin function that is being called, evaluate the arguments
		// and pass it to the built in function
	case *object.BuiltinFunction:
		var evaluatedArguments []object.Object
		for _, argument := range expr.Arguments {
			evaluatedArguments = append(evaluatedArguments, i.Eval(argument))
		}

		// Pass the array as a variadic argument
		obj := callee.Fn(evaluatedArguments...)
		// If the object is a return value
		returnObj, ok := obj.(*object.Return)
		if ok {
			return returnObj.Value
		}
		return obj
	// This is to initialize a struct from nothing-ness
	// TODO: Add in parameters when initializing the structs
	// the parameters have to match up with a init function in the
	// struct itself. Reference can be how python objects are initialized
	case *object.Struct:
		// Make a copy of it, TODO: need to check if there are other ways to
		// make this work
		newCallee := callee
		newCallee.Attributes = make(map[string]object.Object)
		newCallee.Methods = make(map[string]object.Object)
		return newCallee
	// If the callee is a list, then the following is what is parsed
	// (List)[1]
	// Where the '1' is now the argument to the 'callee'
	case *object.List:
		// Can do this as we know that it will only parse a single expression
		evaluatedIndex := i.Eval(expr.Arguments[0])

		intIndex, ok := evaluatedIndex.(*object.Integer)
		if !ok {
			globals.ErrorList = append(globals.ErrorList,
				errors.NewRuntimeError(evaluatedIndex, "Indexed operation on a list expression is not an integer"))
		}

		obj := callee.Value[intIndex.Value]
		return obj
	case *object.HashMap:
		// Get the object
		objectIndex := i.Eval(expr.Arguments[0])

		// TODO : Refactor this
		objectHashable, ok := objectIndex.(object.Hashable)
		if ok {
			obj, found := callee.Value[objectHashable.Hash()]
			if found {
				return obj.Value
			}
		}

		// object not found
		return nil
	case *object.String:
		// Same as above, on the list, by the time it reaches here,
		// it is known that there is only one expression
		evaluatedIndex := i.Eval(expr.Arguments[0])

		intIndex, ok := evaluatedIndex.(*object.Integer)
		if !ok {
			globals.ErrorList = append(globals.ErrorList,
				errors.NewRuntimeError(evaluatedIndex, "Indexed operation on a list expression is not an integer"))
		}

		return &object.String{Value: string(callee.Value[intIndex.Value])}
	default:
		return nil
	}
}

func (i *Interpreter) evalGetExpression(expr *ast.GetExpression) object.Object {
	switch callee := i.Eval(expr.Callee).(type) {
	case *object.Struct:
		// Check if attribute exists
		variableExpression := expr.Attribute.(*ast.VariableExpression)
		obj, ok := callee.Attributes[variableExpression.Name.Literal]
		if ok {
			return obj
		}
		return nil
	default:
		return nil
	}
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

	for _, stmt := range statements {
		obj := i.Eval(stmt)

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
	if obj.RawType() == object.BOOL {
		return obj.(*object.Bool).Value
	}

	if obj.RawType() == object.INTEGER {
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

	// Populate the environment with all the built in functions
	builtin.PopulateEnvironment(environment)

	i := &Interpreter{
		Statements:  statements,
		Environment: environment,
	}
	return i
}
