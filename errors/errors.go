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
func RuntimeError(rgs ...object.Object) {
	fmt.Println("Hello from RuntimeError")
}
