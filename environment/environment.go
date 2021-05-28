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

	// If it does not exist in the current environment, go up
	// the parents
	if e.Parent != nil {
		e.Parent.Set(name, value)
		return
	}

	panic("Name not found in environment : Set")
}

func (e *Environment) SetIndex(name string, index object.Object, value object.Object) {
	_, ok := e.Values[name]
	if ok {
		list, listOk := e.Values[name].(*object.List)
		listIndex, listIndexOk := index.(*object.Integer)
		if listOk && listIndexOk {
			list.Value[listIndex.Value] = value
		}

		hash, hashOk := e.Values[name].(*object.HashMap)
		_, hashable := index.(object.Hashable)
		if hashOk && hashable {
			// Construct the hashkey here.
			hashKey := object.HashKey{
				Type:  index.RawType(),
				Value: index.(object.Hashable).Hash().Value,
			}
			hashValue := object.HashValue{
				Key:   index,
				Value: value,
			}
			hash.Value[hashKey] = hashValue
		}
		return
	}

	// If it does not exist in the current environment, go up
	// the parents
	if e.Parent != nil {
		e.Parent.SetIndex(name, index, value)
		return
	}

	panic("Name not found in environment : SetIndex")
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

// Check if a name exists within the environment.
func (e *Environment) Exists(name string) bool {
	_, ok := e.Values[name]
	if ok {
		return true
	}
	return false
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
