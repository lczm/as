package interpreter

import (
	"strconv"
	"testing"

	"github.com/lczm/as/lexer"
	"github.com/lczm/as/object"
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
		// Directly hook into the eval function instead, as
		// the Start() method is self contained
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

func TestTruthy(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput bool
	}{
		{
			"0",
			false,
		},
		{
			"1",
			true,
		},
		{
			"1 > 2",
			false,
		},
		{
			"2 >=2",
			true,
		},
		{
			"3 > 2",
			true,
		},
	}

	lexer := lexer.New()
	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		obj := interpreter.Eval(statements[0])

		if obj.(*object.Bool).Value != test.expectedOutput {
			t.Fatalf("Test: [%d] - Incorrect value, expected=%t, got=%t",
				i, test.expectedOutput, obj.(*object.Bool).Value)
		}
	}
}

// func TestIfStatements(t *testing.T) {
// 	tests := []struct {
// 		input          string
// 		expectedOutput object.Object
// 	}{}
// }
