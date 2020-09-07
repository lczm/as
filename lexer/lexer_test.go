package lexer

import (
	"testing"

	"github.com/lczm/as/token"
)

func TestIndividualTokenScan(t *testing.T) {
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
				t.Fatalf("Test : [%d] - Wrong TokenType, expected=%q, got=%q",
					i, tests[i].expectedType, token.Type)
			}
			if token.Literal != tests[i].expectedLiteral {
				t.Fatalf("Test : [%d] - Wrong Literal, expected=%q, got=%q",
					i, tests[i].expectedLiteral, token.Literal)
			}
		}
	}
}

func TestMultipleTokenScan(t *testing.T) {
	tests := []struct {
		input            string
		expectedTypes    []token.TokenType
		expectedLiterals []string
	}{
		{
			`= + - ! * /`,
			[]token.TokenType{token.ASSIGN, token.PLUS, token.MINUS, token.BANG, token.ASTERISK, token.SLASH},
			[]string{"=", "+", "-", "!", "*", "/"},
		},
	}

	lexer := New()
	for i, test := range tests {
		tokens := lexer.Scan(test.input)

		if len(test.expectedTypes) != len(test.expectedLiterals) {
			t.Fatalf("Test : [%d] - Mismatch amount of types and literals, expectedTypes=%d, expectedLiterals=%d",
				i, len(test.expectedTypes), len(test.expectedLiterals))
		}

		if len(tokens) != len(test.expectedTypes) {
			t.Fatalf("Test : [%d] - Mismatch amount of scanned tokens and expectedTokens,"+
				"scanned=%d, expectedTokens=%d",
				i, len(tokens), len(test.expectedTypes))
		}

		for j, token := range tokens {
			if token.Type != tests[i].expectedTypes[j] {
				t.Fatalf("Test : [%d - %d] - Wrong TokenType, expected=%q, got=%q",
					i, j, tests[i].expectedTypes[j], token.Type)
			}
			if token.Literal != tests[i].expectedLiterals[j] {
				t.Fatalf("Test : [%d - %d] - Wrong Literal, expected=%q, got=%q",
					i, j, tests[i].expectedLiterals[j], token.Literal)
			}
		}
	}
}
