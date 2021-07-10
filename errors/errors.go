package errors

import (
	"fmt"

	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

// TODO : The lexer has to keep track of the tokens
// line numbers and column numbers, they can then be passed into
// the specific errors to give more detialed messages

// TODO : Proper exit codes? as these are errors they should not
// be 0 escapes

type Error interface {
	Describe()
}

// Initially this was implemented as with interfaces and structs
// but there is really no need for that as it would complicate
// things, so just simple functions to deal with errors would be enough.

// Syntax errors would take in tokens as arguments as
// it will take place in the parsing phase
type SyntaxError struct {
	Error
	token   token.Token
	message string
}

func NewSyntaxError(token token.Token, message string) SyntaxError {
	se := SyntaxError{
		token:   token,
		message: message,
	}
	return se
}

func (se SyntaxError) Describe() {
	fmt.Printf("Syntax Error at line '%d' : %s\n", se.token.Line, se.message)
}

// Runtime errors will take in objects as they are
// taken place during the interpreting phase
type RuntimeError struct {
	Error
	object  object.Object
	message string
}

func NewRuntimeError(obj object.Object, message string) RuntimeError {
	re := RuntimeError{
		object:  obj,
		message: message,
	}
	return re
}

func (re RuntimeError) Describe() {
	fmt.Printf("Runtime Error : %s at %s\n", re.object.String(), re.message)
}

// This is for error messages that are not the easiest to pass
// objects into, i.e.
// in the case where there is a need to handle multiple parameters
// and they are not entirely relevant (builtin functions)
// A simple message to show is okay.
type DefaultError struct {
	Error
	message string
}

func NewDefaultError(message string) DefaultError {
	de := DefaultError{
		message: message,
	}
	return de
}

func (de DefaultError) Describe() {
	fmt.Printf("Error : %s\n", de.message)
}

type ShadowWarning struct {
	Error
	line         int
	variableName string
}

func NewShadowWarning(line int, variableName string) ShadowWarning {
	sw := ShadowWarning{
		line:         line,
		variableName: variableName,
	}
	return sw
}

func (sw ShadowWarning) Describe() {
	fmt.Printf("Shadow warning at line %d, Declaring an already declared variable: \"%s\"", sw.line, sw.variableName)
}
