package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
	Line    int
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
	MODULUS  = "%"

	// Augmented assignment operators
	AUG_PLUS     = "+="
	AUG_MINUS    = "-="
	AUG_ASTERISK = "*="
	AUG_SLASH    = "/="
	AUG_MODULUS  = "%="

	// Comparison Operators
	LT    = "<"
	LT_EQ = "<="
	GT    = ">"
	GT_EQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Logical Operators
	AND = "&&"
	OR  = "||"

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

	// Quotes, usually for strings
	SINGLE_QUOTE = "'"
	DOUBLE_QUOTE = "\""

	// Single line comments
	COMMENT = "//"

	// Increment / Decrement
	INCREMENT = "++"
	DECREMENT = "--"

	// Booleans
	TRUE  = "TRUE"
	FALSE = "FALSE"

	// Numeric values and identifiers
	NUMBER     = "NUMBER"
	STRING     = "STRING"
	IDENTIFIER = "IDENTIFIER"

	// Statements
	PRINT    = "PRINT"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	WHILE    = "WHILE"
	FOR      = "FOR"
	FUNCTION = "FUNCTION"
	STRUCT   = "STRUCT"
	RETURN   = "RETURN"

	// Misc
	EOF = "EOF"
)
