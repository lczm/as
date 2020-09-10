package object

import "fmt"

// Types
const (
	INTEGER = "INTEGER"
)

// All types implement this interface
type Object interface {
	Type() string
	String() string
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
