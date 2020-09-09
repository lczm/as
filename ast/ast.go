package ast

import (
	"github.com/lczm/as/token"
)

type Statement interface {
	statement()
}

type Expression interface {
	expression()
}

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (be *BinaryExpression) expression() {}

type UnaryExpression struct {
	Right    Expression
	Operator token.Token
}

func (ue *UnaryExpression) expression() {}

type NumberExpression struct {
	Value int
}

func (ne *NumberExpression) expression() {}

type GroupExpression struct {
	Expr Expression
}

func (ge *GroupExpression) expression() {}
