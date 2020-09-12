package interpreter

import (
	"strconv"
	"testing"

	"github.com/lczm/as/lexer"
	"github.com/lczm/as/parser"
)

func TestIntegerExpressions(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			"1 + 2;",
			3,
		},
		{
			"5 - 1 + 5;",
			9,
		},
		{
			"5 * 2;",
			10,
		},
		{
			"(5 - 2) * 2;",
			6,
		},
		{
			"5 - 1 + 2 - (2 * 2);",
			2,
		},
	}

	lexer := lexer.New()
	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		object := interpreter.Eval(statements[0])
		str := object.String()

		value, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}

		if test.expectedOutput != value {
			t.Fatalf("Test : [%d] - Mismatch in values, expected=%d, got=%q",
				i, test.expectedOutput, str)
		}
	}
}
