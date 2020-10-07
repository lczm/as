package builtin

import (
	"fmt"

	"github.com/lczm/as/environment"
	"github.com/lczm/as/object"
)

func TypeFunc() object.Object {
	function := &object.BuiltinFunction{
		Name: "type",
		Fn: func(args ...object.Object) object.Object {
			// TODO : Error out
			if len(args) != 1 {
				fmt.Println("type() can only take in one parameter at a time.")
			}

			obj := args[0]
			return &object.String{Value: obj.Type()}
		},
	}
	return function
}

func LenFunc() object.Object {
	function := &object.BuiltinFunction{
		Name: "len",
		Fn: func(args ...object.Object) object.Object {
			// TODO : Error out, len() can only take in one parameter
			if len(args) != 1 {
				fmt.Println("len() needs at least one parameter.")
			}

			obj := args[0]
			switch obj := obj.(type) {
			case *object.Integer:
				// TODO Format this properly
				fmt.Println("len() cannot be used on an integer.")
			case *object.String:
				return &object.Integer{Value: int64(len(obj.String()))}
			default:
				return nil
			}
			return nil
		},
	}
	return function
}

func PrintFunc() object.Object {
	function := &object.BuiltinFunction{
		Name: "print",
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.String())
			}
			return nil
		},
	}
	return function
}

func PopulateEnvironment(env *environment.Environment) {
	env.Define("type", TypeFunc())
	env.Define("len", LenFunc())
	env.Define("print", PrintFunc())
}
