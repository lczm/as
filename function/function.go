package function

import (
	"github.com/lczm/as/ast"
	"github.com/lczm/as/environment"
	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/object"
)

// Function type, it is an Object as well as a Callable
type Function struct {
	FunctionStatement ast.FunctionStatement
}

func (f *Function) Type() string {
	return object.FUNCTION
}

func (f *Function) String() string {
	return "Function"
}

// The call functions should return an object
// in the case of something like
// var x = function(...)
func (f *Function) Call(i *interpreter.Interpreter,
	arguments []object.Object) object.Object {
	environment := environment.NewChildEnvironment(i.Environment)

	// For each of the arguments that is passed into the call,
	// the environment will have to define each of them into the function
	for i, argument := range arguments {
		environment.Define(f.FunctionStatement.Params[i].Literal, argument)
	}

	// Execute the statements that are defined inside the function
	i.ExecuteBlockStatements(f.FunctionStatement.Body, environment)

	// TODO : When return values are supported in functions,
	// they can be parsed and dealt with here rather than returning nil
	// by default.
	return nil
}
