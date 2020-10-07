package object

import (
	"fmt"

	"github.com/lczm/as/ast"
)

// Types
const (
	BOOL     = "BOOL"
	INTEGER  = "INTEGER"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	STRING   = "STRING"
	BUILTIN  = "BULITIN" // builtin functions from the host language
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
	return fmt.Sprintf("<type: %s>", INTEGER)
}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// String type
type String struct {
	Value string
}

func (s *String) Type() string {
	return fmt.Sprintf("<type: %s>", STRING)
}

func (s *String) String() string {
	return s.Value
}

// Function type, it is an Object as well as a Callable
type Function struct {
	FunctionStatement ast.FunctionStatement
}

func (f *Function) Type() string {
	return fmt.Sprintf("<type: %s>", STRING)
}

func (f *Function) String() string {
	return fmt.Sprintf("Function : <%s>", f.FunctionStatement.Name.Literal)
}

type BuiltinFunction struct {
	Name string
	Fn   func(args ...Object) Object
}

func (bf *BuiltinFunction) Type() string {
	return fmt.Sprintf("<type: %s>", BUILTIN)
}

func (bf *BuiltinFunction) String() string {
	return fmt.Sprintf("BulitinFunction: <%s>", bf.Name)
}

// The call functions should return an object
// in the case of something like
// var x = function(...)
func (f *Function) Call() {}

// Return type, this is only for the interpreter and is not for use normally.
type Return struct {
	Value Object
}

func (r *Return) Type() string {
	return RETURN
}

func (r *Return) String() string {
	return r.Value.String()
}
