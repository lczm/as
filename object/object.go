package object

import (
	"fmt"
)

// Types
const (
	BOOL     = "BOOL"
	INTEGER  = "INTEGER"
	FUNCTION = "FUNCTION"
)

// All types implement this interface
type Object interface {
	Type() string
	String() string
}

// All the call-able objects will implement this interface
// i.e. functions
type Callable interface {
	Call() Object
}

// Boolean type
type Bool struct {
	Value bool
}

func (b *Bool) Type() string {
	return BOOL
}

func (b *Bool) String() string {
	if b.Value == true {
		return "true"
	}
	return "false"
}

// Integer type
type Integer struct {
	Value int64
}

func (i *Integer) Type() string {
	return INTEGER
}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}
