package builtin

import (
	"fmt"

	"github.com/lczm/as/environment"
	"github.com/lczm/as/errors"
	"github.com/lczm/as/object"
)

func TypeFunc() object.Object {
	function := &object.BuiltinFunction{
		Name: "type",
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				errors.DefaultError("type() can only take in one parameter at a time.")
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
			if len(args) != 1 {
				errors.DefaultError("len() needs at least one parameter.")
			}

			obj := args[0]
			switch obj := obj.(type) {
			case *object.Integer:
				errors.DefaultError("len() cannot be used on an integer.")
			case *object.String:
				return &object.Integer{Value: int64(len(obj.String()))}
			case *object.List:
				return &object.Integer{Value: int64(len(obj.Value))}
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

// This is to piggy-back off the golang 'append' function
// Mostly as a way to add more elements to a list.
// In the future if classes are implemented
// maybe lists can be a natively implemented class such that
// .append() or .remove() would work seamlessly.
func AppendFunc() object.Object {
	function := &object.BuiltinFunction{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				fmt.Println("append() takes in two parameters, the appendee and the element.")
				return nil
			}

			// TODO : Support more than just lists, possibly hashmaps
			// Extract the list value out of the list object
			list := args[0].(*object.List).Value
			element := args[1]

			// Add the list item
			list = append(list, element)

			// Wrap it back into an object
			return &object.List{Value: list}
		},
	}
	return function
}

// Removes an element at the specified index
func RemoveAtFunc() object.Object {
	function := &object.BuiltinFunction{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				fmt.Println("removeAt() takes in two parameters, the list object and the index to remove at.")
				return nil
			}

			list := args[0].(*object.List).Value
			index := args[1].(*object.Integer).Value

			list = append(list[:index], list[index+1:]...)

			return &object.List{Value: list}
		},
	}
	return function
}

func PopulateEnvironment(env *environment.Environment) {
	env.Define("type", TypeFunc())
	env.Define("len", LenFunc())
	env.Define("print", PrintFunc())
	env.Define("append", AppendFunc())
	env.Define("removeAt", RemoveAtFunc())
}
