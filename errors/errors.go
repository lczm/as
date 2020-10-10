package errors

import (
	"fmt"
	"os"

	"github.com/lczm/as/object"
	"github.com/lczm/as/token"
)

// TODO : The lexer has to keep track of the tokens
// line numbers and column numbers, they can then be passed into
// the specific errors to give more detialed messages

// TODO : Proper exit codes? as these are errors they should not
// be 0 escapes

// Initially this was implemented as with interfaces and structs
// but there is really no need for that as it would complicate
// things, so just simple functions to deal with errors would be enough.

// Syntax errors would take in tokens as arguments as
// it will take place in the parsing phase
func SyntaxError(tokenType token.TokenType, message string) {
	fmt.Printf("Syntax Error at '%s' : %s\n", tokenType, message)

	// Quit the program without panicing
	os.Exit(0)
}

// Runtime errors will take in objects as they are
// taken place during the interpreting phase
func RuntimeError(obj object.Object, message string) {
	fmt.Printf("Runtime Error : %s at %s\n", obj.String(), message)

	// Quit the program without panicing
	os.Exit(0)
}

// This is for error messages that are not the easiest to pass
// objects into, i.e.
// in the case where there is a need to handle multiple parameters
// and they are not entirely relevant (builtin functions)
// A simple message to show is okay.
func DefaultError(message string) {
	fmt.Printf("Error : %s\n", message)

	// Quit the program without panicing
	os.Exit(0)
}
