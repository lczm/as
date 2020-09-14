package environment

import "github.com/lczm/as/object"

type Environment struct {
	Values map[string]object.Object
	Parent *Environment
}

// This method can potentially take in other context parameters
// So that there can be a check for something like -Wshadow
func (e *Environment) Set(name string, value object.Object) {
	e.Values[name] = value
}

func (e *Environment) Get(name string) object.Object {
	// TODO : Check that the name exists before returning.
	return e.Values[name]
}

func New() *Environment {
	values := make(map[string]object.Object)

	e := &Environment{
		Values: values,
		Parent: nil,
	}
	return e
}

func NewChildEnvironment(parent *Environment) *Environment {
	values := make(map[string]object.Object)

	e := &Environment{
		Values: values,
		Parent: parent,
	}
	return e
}
