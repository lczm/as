package lexer

import (
	"github.com/lczm/as/token"
)

type Lexer struct {
}

func (l Lexer) Scan(source string) []token.Token {
	var tokens []token.Token

	token := token.Token{
		Literal: "hello",
	}

	tokens = append(tokens, token)

	return tokens
}

func New() *Lexer {
	l := &Lexer{}
	return l
}
