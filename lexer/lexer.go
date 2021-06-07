package lexer

import (
	"fmt"

	"github.com/lczm/as/token"
)

type Lexer struct {
	Keywords map[string]token.TokenType
}

func (l *Lexer) Scan(source string) []token.Token {
	var tokens []token.Token

	// Default to line 1
	currentLine := 1

	currentIndex := 0
	for currentIndex < len(source) {
		// Get the current character
		ch := source[currentIndex]
		// Increment index, as used by previous character
		currentIndex++

		switch ch {
		case ' ':
		case '\t': // Tabs
		case '\n': // New line
			currentLine++
		case '\r': // Carriage Return (CR)
			break
		// Operators
		case '+':
			// Handle the case of '++'
			if currentIndex < len(source) && source[currentIndex] == '+' {
				tokens = append(tokens, token.Token{
					Type:    token.INCREMENT,
					Literal: "++",
					Line:    currentLine,
				})
				currentIndex++
			} else if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_PLUS,
					Literal: "+=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.PLUS,
					Literal: "+",
					Line:    currentLine,
				})
			}
		case '-':
			if currentIndex < len(source) && source[currentIndex] == '-' {
				tokens = append(tokens, token.Token{
					Type:    token.DECREMENT,
					Literal: "--",
					Line:    currentLine,
				})
				currentIndex++
			} else if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_MINUS,
					Literal: "-=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.MINUS,
					Literal: "-",
					Line:    currentLine,
				})
			}
		case '!':
			// Handle the case of '!='
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.NOT_EQ,
					Literal: "!=",
					Line:    currentLine,
				})
				currentIndex++
			} else { // Handle the case of '!'
				tokens = append(tokens, token.Token{
					Type:    token.BANG,
					Literal: "!",
					Line:    currentLine,
				})
			}
		case '*':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_ASTERISK,
					Literal: "*=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.ASTERISK,
					Literal: "*",
					Line:    currentLine,
				})
			}
		case '/':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_SLASH,
					Literal: "/=",
					Line:    currentLine,
				})
				currentIndex++
			} else if currentIndex < len(source) && source[currentIndex] == '/' {
				tokens = append(tokens, token.Token{
					Type:    token.COMMENT,
					Literal: "//",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.SLASH,
					Literal: "/",
					Line:    currentLine,
				})
			}
		case '%':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_MODULUS,
					Literal: "%=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.MODULUS,
					Literal: "%",
					Line:    currentLine,
				})
			}
		// Comparison Operators
		case '<':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.LT_EQ,
					Literal: "<=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.LT,
					Literal: "<",
					Line:    currentLine,
				})
			}
		case '>':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.GT_EQ,
					Literal: ">=",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.GT,
					Literal: ">",
					Line:    currentLine,
				})
			}
		case '=':
			// Handle the case of '=='
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.EQ,
					Literal: "==",
					Line:    currentLine,
				})
				currentIndex++
			} else { // Handle the case of '='
				tokens = append(tokens, token.Token{
					Type:    token.ASSIGN,
					Literal: "=",
					Line:    currentLine,
				})
			}
		// Logical Comparisons
		case '&':
			if source[currentIndex] == '&' {
				tokens = append(tokens, token.Token{
					Type:    token.AND,
					Literal: "&&",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				panic("Single '&' character cannot be lexed")
			}
		case '|':
			if source[currentIndex] == '|' {
				tokens = append(tokens, token.Token{
					Type:    token.OR,
					Literal: "||",
					Line:    currentLine,
				})
				currentIndex++
			} else {
				panic("Single '|' character cannot be lexed")
			}
		// Delimiters
		case '.':
			tokens = append(tokens, token.Token{
				Type:    token.DOT,
				Literal: ".",
				Line:    currentLine,
			})
		case ',':
			tokens = append(tokens, token.Token{
				Type:    token.COMMA,
				Literal: ",",
				Line:    currentLine,
			})
		case ':':
			tokens = append(tokens, token.Token{
				Type:    token.COLON,
				Literal: ":",
				Line:    currentLine,
			})
		case ';':
			tokens = append(tokens, token.Token{
				Type:    token.SEMICOLON,
				Literal: ";",
				Line:    currentLine,
			})
		case '(':
			tokens = append(tokens, token.Token{
				Type:    token.LPAREN,
				Literal: "(",
				Line:    currentLine,
			})
		case ')':
			tokens = append(tokens, token.Token{
				Type:    token.RPAREN,
				Literal: ")",
				Line:    currentLine,
			})
		case '{':
			tokens = append(tokens, token.Token{
				Type:    token.LBRACE,
				Literal: "{",
				Line:    currentLine,
			})
		case '}':
			tokens = append(tokens, token.Token{
				Type:    token.RBRACE,
				Literal: "}",
				Line:    currentLine,
			})
		case '[':
			tokens = append(tokens, token.Token{
				Type:    token.LBRACKET,
				Literal: "[",
				Line:    currentLine,
			})
		case ']':
			tokens = append(tokens, token.Token{
				Type:    token.RBRACKET,
				Literal: "]",
				Line:    currentLine,
			})
		case '"':
			extendedIndex := currentIndex
			for extendedIndex < len(source) && source[extendedIndex] != '"' {
				extendedIndex++
			}

			stringValue := source[currentIndex:extendedIndex]
			currentIndex = extendedIndex + 1

			tokens = append(tokens, token.Token{
				Type:    token.STRING,
				Literal: stringValue,
				Line:    currentLine,
			})
		default:
			if l.isDigit(ch) { // Handle numeric case
				extendedIndex := currentIndex
				for extendedIndex < len(source) && l.isDigit(source[extendedIndex]) {
					extendedIndex++
				}

				tokens = append(tokens, token.Token{
					Type:    token.NUMBER,
					Literal: source[currentIndex-1 : extendedIndex],
					Line:    currentLine,
				})
				currentIndex = extendedIndex
			} else if l.isAlphaNumeric(ch) { // Handle alpha-numeric case
				// If it hits this branch, it means that it starts off with a
				// alphaNumeric, i.e. 'abc', 'bcd'
				// This can then get classified as an identifier
				extendedIndex := currentIndex
				for extendedIndex < len(source) &&
					l.isAlphaNumeric(source[extendedIndex]) {
					extendedIndex++
				}

				identifier := source[currentIndex-1 : extendedIndex]

				// Check if it is a keyword
				if l.isKeyword(identifier) {
					tokens = append(tokens, token.Token{
						Type:    l.Keywords[identifier],
						Literal: identifier,
						Line:    currentLine,
					})
				} else {
					tokens = append(tokens, token.Token{
						Type:    token.IDENTIFIER,
						Literal: identifier,
						Line:    currentLine,
					})
				}
				currentIndex = extendedIndex
			} else {
				// TODO : Do some form of error handling here
				fmt.Println("The lexer cannot handle this character : ", string(ch))
			}
		}
	}

	return tokens
}

func (l *Lexer) isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func (l *Lexer) isAlphaNumeric(b byte) bool {
	// Check that it is also a digit, as this will be useful
	// for during the extended index cases
	if (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		b == '_' || // Handle underscores as well
		l.isDigit(b) {
		return true
	}
	return false
}

func (l *Lexer) isKeyword(s string) bool {
	_, ok := l.Keywords[s]
	return ok
}

func New() *Lexer {
	keywords := make(map[string]token.TokenType)
	keywords["var"] = token.VAR
	keywords["if"] = token.IF
	keywords["else"] = token.ELSE
	keywords["while"] = token.WHILE
	keywords["for"] = token.FOR
	keywords["function"] = token.FUNCTION
	keywords["return"] = token.RETURN
	keywords["true"] = token.TRUE
	keywords["false"] = token.FALSE
	keywords["struct"] = token.STRUCT

	l := &Lexer{
		Keywords: keywords,
	}
	return l
}
