package tests

import (
	"strconv"
	"testing"

	"github.com/lczm/as/interpreter"
	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func TestAugmentedAssignsFunc(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			`
            var output = 10;
            output += 10;
            `,
			20,
		},
		{
			`
            var output = 10;
            output -= 10;
            `,
			0,
		},
		{
			`
            var output = 10;
            output *= 10;
            `,
			100,
		},
		{
			`
            var output = 10;
            output /= 2;
            `,
			5,
		},
		{
			`
            var output = 100;
            output %= 10;
            `,
			0,
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
			// The object returns with strings, this should be fixed in the future, but
			// for the time being this should do well enough for tests

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
