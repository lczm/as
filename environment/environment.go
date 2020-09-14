package environment

import "github.com/lczm/as/object"

type Environment struct {
	Values map[string]object.Object
	Parent *Environment
}

func (e *Environment) Define(name string, value object.Object) {
	e.Values[name] = value
}

// This method can potentially take in other context parameters
// So that there can be a check for something like -Wshadow
func (e *Environment) Set(name string, value object.Object) {
	_, ok := e.Values[name]
	if ok {
		e.Values[name] = value
		return
	}

	if e.Parent != nil {
		e.Parent.Set(name, value)
		return
	}

	panic("Name not found in environment : Set")
}

func (e *Environment) Get(name string) object.Object {
	object, ok := e.Values[name]
	if ok {
		return object
	}

	// Go up the parent environments to get the string.
	// The parent environment will deal with the error handling
	// through the panic call.
	if e.Parent != nil {
		return e.Parent.Get(name)
	}

	panic("Name not found in environment : Get")
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
