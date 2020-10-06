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
			fmt.Println("Hello from typefunc")
			return nil
		},
	}
	return function
}

func PopulateEnvironment(env *environment.Environment) {
	env.Define("type", TypeFunc())
}
