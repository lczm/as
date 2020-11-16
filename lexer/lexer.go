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
		case '\r': // Carriage Return (CR)
			break
		// Operators
		case '+':
			// Handle the case of '++'
			if currentIndex < len(source) && source[currentIndex] == '+' {
				tokens = append(tokens, token.Token{
					Type:    token.INCREMENT,
					Literal: "++",
				})
				currentIndex++
			} else if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_PLUS,
					Literal: "+=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.PLUS,
					Literal: "+",
				})
			}
		case '-':
			if currentIndex < len(source) && source[currentIndex] == '-' {
				tokens = append(tokens, token.Token{
					Type:    token.DECREMENT,
					Literal: "--",
				})
				currentIndex++
			} else if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_MINUS,
					Literal: "-=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.MINUS,
					Literal: "-",
				})
			}
		case '!':
			// Handle the case of '!='
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.NOT_EQ,
					Literal: "!=",
				})
				currentIndex++
			} else { // Handle the case of '!'
				tokens = append(tokens, token.Token{
					Type:    token.BANG,
					Literal: "!",
				})
			}
		case '*':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_ASTERISK,
					Literal: "*=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.ASTERISK,
					Literal: "*",
				})
			}
		case '/':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_SLASH,
					Literal: "/=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.SLASH,
					Literal: "/",
				})
			}
		case '%':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.AUG_MODULUS,
					Literal: "%=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.MODULUS,
					Literal: "%",
				})
			}
		// Comparison Operators
		case '<':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.LT_EQ,
					Literal: "<=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.LT,
					Literal: "<",
				})
			}
		case '>':
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.GT_EQ,
					Literal: ">=",
				})
				currentIndex++
			} else {
				tokens = append(tokens, token.Token{
					Type:    token.GT,
					Literal: ">",
				})
			}
		case '=':
			// Handle the case of '=='
			if currentIndex < len(source) && source[currentIndex] == '=' {
				tokens = append(tokens, token.Token{
					Type:    token.EQ,
					Literal: "==",
				})
				currentIndex++
			} else { // Handle the case of '='
				tokens = append(tokens, token.Token{
					Type:    token.ASSIGN,
					Literal: "=",
				})
			}
		// Logical Comparisons
		case '&':
			if source[currentIndex] == '&' {
				tokens = append(tokens, token.Token{
					Type:    token.AND,
					Literal: "&&",
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
				})
				currentIndex++
			} else {
				panic("Single '|' character cannot be lexed")
			}
		// Delimiters
		case ',':
			tokens = append(tokens, token.Token{
				Type:    token.COMMA,
				Literal: ",",
			})
		case ':':
			tokens = append(tokens, token.Token{
				Type:    token.COLON,
				Literal: ":",
			})
		case ';':
			tokens = append(tokens, token.Token{
				Type:    token.SEMICOLON,
				Literal: ";",
			})
		case '(':
			tokens = append(tokens, token.Token{
				Type:    token.LPAREN,
				Literal: "(",
			})
		case ')':
			tokens = append(tokens, token.Token{
				Type:    token.RPAREN,
				Literal: ")",
			})
		case '{':
			tokens = append(tokens, token.Token{
				Type:    token.LBRACE,
				Literal: "{",
			})
		case '}':
			tokens = append(tokens, token.Token{
				Type:    token.RBRACE,
				Literal: "}",
			})
		case '[':
			tokens = append(tokens, token.Token{
				Type:    token.LBRACKET,
				Literal: "[",
			})
		case ']':
			tokens = append(tokens, token.Token{
				Type:    token.RBRACKET,
				Literal: "]",
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
					})
				} else {
					tokens = append(tokens, token.Token{
						Type:    token.IDENTIFIER,
						Literal: identifier,
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

	l := &Lexer{
		Keywords: keywords,
	}
	return l
}
