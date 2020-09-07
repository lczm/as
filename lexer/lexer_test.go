package lexer

import (
	"testing"

	"github.com/lczm/as/token"
)

func TestIndividualScan(t *testing.T) {
	tests := []struct {
		input           string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Operators
		{"=", token.ASSIGN, "="},
		{"+", token.PLUS, "+"},
		{"-", token.MINUS, "-"},
		{"!", token.BANG, "!"},
		{"*", token.ASTERISK, "*"},
		{"/", token.SLASH, "/"},

		// Comparison Operators
		{"<", token.LT, "<"},
		{"<=", token.LT_EQ, "<="},
		{">", token.GT, ">"},
		{">=", token.GT_EQ, ">="},
		{"==", token.EQ, "=="},
		{"!=", token.NOT_EQ, "!="},

		// Delimiters
		{",", token.COMMA, ","},
		{":", token.COLON, ":"},
		{";", token.SEMICOLON, ";"},
		{"(", token.LPAREN, "("},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"}", token.RBRACE, "}"},
		{"[", token.LBRACKET, "["},
		{"]", token.RBRACKET, "]"},
	}

	lexer := New()
	for i, test := range tests {
		tokens := lexer.Scan(test.input)
		for _, token := range tokens {
			if token.Type != tests[i].expectedType {
				t.Fatalf("[%d] - Wrong TokenType, expected=%q, got=%q",
					i, tests[i].expectedType, token.Type)
			}
			if token.Literal != tests[i].expectedLiteral {
				t.Fatalf("[%d] - Wrong Literal, expected=%q, got=%q",
					i, tests[i].expectedLiteral, token.Literal)
			}
		}
	}
}
