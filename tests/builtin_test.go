package tests

import (
	"strconv"
	"strings"
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

func TestTypeFunc(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`var output = type("Hello");`,
			"<type: STRING>",
		},
		{
			`var output = type(1);`,
			"<type: INTEGER>",
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
			if obj.String() != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}

func TestAppendFunc(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`var output = [];`,
			"[]",
		},
		{
			`
            var output = [];
            output = append(output, 1);
            output = append(output, 2);
            `,
			"[1 2]",
		},
		{
			`
            var output = [];
            for (var a = 0; a < 5; a++) {
                output = append(output, a);
            }
            `,
			"[0 1 2 3 4]",
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
			if strings.TrimSuffix(obj.String(), "\n") != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}
