package environment

import "github.com/lczm/as/object"

type Environment struct {
	Values map[string]object.Object
}

func (e *Environment) Add(name string, value object.Object) {
	e.Values[name] = value
}

func (e *Environment) Get(name string) object.Object {
	return e.Values[name]
}

func New() *Environment {
	e := &Environment{}
	return e
}
