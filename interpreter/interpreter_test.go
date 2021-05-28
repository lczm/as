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

func TestListIndexExpressions(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`
			var output = [0, 1, 2, 3];
			output[0] = 100;
			`,
			"[100, 1, 2, 3]",
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		// Directly hook into the eval function instead, as
		// the Start() method is self contained
		interpreter.Start()

		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)
			if test.expectedOutput != obj.String() {
				t.Fatalf("Test : [%d] - Mismatch in values, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}

func TestHashMapIndexExpressions(t *testing.T) {
	tests := []struct {
		input          string
		keys           []string
		expectedOutput []string
	}{
		{
			`
			var output = {0:10, 1:20, 2:30, 3:40};
			`,
			[]string{"0", "1", "2", "3"},
			[]string{"10", "20", "30", "40"},
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		// Directly hook into the eval function instead, as
		// the Start() method is self contained
		interpreter.Start()

		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)

			for j := 0; j < len(test.expectedOutput); j++ {
				key := test.keys[j]
				expectedOutput := test.expectedOutput[j]

				obj := obj.(*object.HashMap)
				intValue, _ := strconv.Atoi(key)
				keyStr := &object.Integer{Value: int64(intValue)}

				hashKey := object.HashKey{
					Type:  keyStr.RawType(),
					Value: keyStr.Hash().Value,
				}
				hashValue := obj.Value[hashKey]
				if expectedOutput != hashValue.Value.String() {
					t.Fatalf("Test : [%d] - Mismatch in values, expected=%s, got=%s",
						i, expectedOutput, hashValue.Value.String())
				}
			}
		}
	}
}

func TestIncrementDecrement(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput int
	}{
		{
			`
            var output = 10;
            output++;
            `,
			11,
		},
		{
			`
            var output = 10;
            output--;
            `,
			9,
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
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

func TestTruthy(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput bool
	}{
		{
			"0;",
			false,
		},
		{
			"1;",
			true,
		},
		{
			"1 > 2;",
			false,
		},
		{
			"2 >=2;",
			true,
		},
		{
			"3 > 2;",
			true,
		},
	}

	lexer := lexer.New()
	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		b := interpreter.IsTruthy(interpreter.Eval(statements[0]))

		if b != test.expectedOutput {
			t.Fatalf("Test: [%d] - Incorrect value, expected=%t, got=%t",
				i, test.expectedOutput, b)
		}
	}
}

func TestIfStatements(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`var output;
            if (5 > 1) {
                output = 5;
            } else {
                output = 1;
            }`,
			"5",
		},
		{
			`var output;
			if (5 >= 5) {
				output = 5;
			} else {
				output = 1;
			}`,
			"5",
		},
		{
			`var output;
			if (5 > 6) {
				output = 5;
			} else {
				output = 1;
			}`,
			"1",
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
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

func TestWhileStatements(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`var output = 0;
            while (output < 10) {
                output = output + 1;
            }`,
			"10",
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		interpreter.Start()

		// Directly hook into the environment to check for output variable
		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)
			if obj.String() != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}

func TestForStatements(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`
            for (var output = 0; output < 3; output = output + 1) {
            }`,
			"3",
		},
		{
			`
            for (var output = 10; output != 2; output = output - 1) {
            }
            `,
			"2",
		},
		{ // Test the more traditional c-style for loops with i++ iterators
			`
            for (var output = 0; output < 10; output++) {
            }
            `,
			"10",
		},
		{ // Test the more traditional c-style for loops with i-- iterators
			`
            for (var output = 10; output != 2; output--) {
            }
            `,
			"2",
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		interpreter.Start()

		// Directly hook into the environment to check for outpit variable
		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)
			if obj.String() != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}

func TestFunctionStatements(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`
			function fib(n) {
				if (n <= 1) {
					return n;
				}
				return fib(n - 2) + fib(n - 1);
			}
			var output = fib(6);
			`,
			"8",
		},
	}

	outputVariable := "output"
	lexer := lexer.New()

	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		parser := parser.New(tokens)
		statements := parser.Parse()

		interpreter := New(statements)
		interpreter.Start()

		// Directly hook into the environment to check for outpit variable
		if interpreter.Environment.Exists(outputVariable) {
			obj := interpreter.Environment.Get(outputVariable)
			if obj.String() != test.expectedOutput {
				t.Fatalf("Test: [%d] - Incorrect value, expected=%s, got=%s",
					i, test.expectedOutput, obj.String())
			}
		}
	}
}
