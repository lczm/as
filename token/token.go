package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

// Available Tokens
const (
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	// Comparison Operators
	LT    = "<"
	LT_EQ = "<="
	GT    = ">"
	GT_EQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// Numeric values and identifiers
	NUMBER     = "NUMBER"
	IDENTIFIER = "IDENTIFIER"

	// Misc
	EOF = "EOF"
)
