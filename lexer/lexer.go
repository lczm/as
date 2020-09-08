package lexer

import (
	"fmt"
	"github.com/lczm/as/token"
)

type Lexer struct{}

func (l Lexer) Scan(source string) []token.Token {
	var tokens []token.Token

	currentIndex := 0
	for currentIndex < len(source) {
		// Get the current character
		ch := source[currentIndex]
		// Increment index, as used by previous character
		currentIndex++

		switch ch {
		// Operators
		case '+':
			tokens = append(tokens, token.Token{
				Type:    token.PLUS,
				Literal: "+",
			})
		case '-':
			tokens = append(tokens, token.Token{
				Type:    token.MINUS,
				Literal: "-",
			})
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
			tokens = append(tokens, token.Token{
				Type:    token.ASTERISK,
				Literal: "*",
			})
		case '/':
			tokens = append(tokens, token.Token{
				Type:    token.SLASH,
				Literal: "/",
			})
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
		default:
			if isDigit(ch) { // Handle numeric case
				extendedIndex := currentIndex
				for extendedIndex < len(source) && isDigit(source[extendedIndex]) {
					extendedIndex++
				}

				tokens = append(tokens, token.Token{
					Type:    token.NUMBER,
					Literal: source[currentIndex-1 : extendedIndex],
				})
				fmt.Println(source[currentIndex-1 : extendedIndex])
				currentIndex = extendedIndex
			} else { // Handle alpha-numeric case
			}
			fmt.Println("Default case")
		}
	}

	return tokens
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

func New() *Lexer {
	l := &Lexer{}
	return l
}
