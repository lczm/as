package environment

type Environment struct{}

func New() *Environment {
	e := &Environment{}
	return e
}
