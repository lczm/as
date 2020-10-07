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
			if len(args) != 1 {
				// TODO : Error out
				fmt.Println("type() can only take in one parameter at a time.")
			}

			obj := args[0]
			fmt.Println(obj.Type())
			return nil
		},
	}
	return function
}

func PopulateEnvironment(env *environment.Environment) {
	env.Define("type", TypeFunc())
}
