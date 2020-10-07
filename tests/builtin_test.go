package tests

import (
	"strconv"
	"testing"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func TestLenFunc(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			`var output = len("Hello");`,
			5,
		},
		{
			`var output = len("test");`,
			4,
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := interpreter.New(statements)
		interpreter.Start()

		// Directly hook into the environment to check for output variable.
		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)

			value, err := strconv.Atoi(obj.String())
			if err != nil {
				panic(err)
			}

			if value != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%d, got=%d",
					i, test.expectedOutput, value)
			}
		}
	}
}
