package ast

import (
	"github.com/lczm/as/token"
)

type AstNode interface{}

type Statement interface {
	AstNode
	statement()
}

type Expression interface {
	AstNode
	expression()
	String() string
}

// Statements
type StatementExpression struct {
	Expr Expression
}

func (se *StatementExpression) statement() {}

type PrintStatement struct {
	Expr Expression
}

func (pe *PrintStatement) statement() {}

type IfStatement struct {
	Condition Expression
	Then      Statement
	Else      Statement
}

func (is *IfStatement) statement() {}

type VariableStatement struct {
	Name        token.Token
	Initializer Expression
}

func (vs *VariableStatement) statement() {}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

func (ws *WhileStatement) statement() {}

type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statement() {}

// Expressions
type AssignmentExpression struct {
	Name  token.Token
	Value Expression
}

func (ae *AssignmentExpression) expression() {}
func (ae *AssignmentExpression) String() string {
	return "Assignment Expression"
}

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (be *BinaryExpression) expression() {}
func (be *BinaryExpression) String() string {
	return "Binary Expression"
}

type UnaryExpression struct {
	Right    Expression
	Operator token.Token
}

func (ue *UnaryExpression) expression() {}
func (ue *UnaryExpression) String() string {
	return "Unary Expression"
}

type LogicalExpression struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (le *LogicalExpression) expression() {}
func (le *LogicalExpression) String() string {
	return "Logical Expression"
}

type NumberExpression struct {
	Value int
}

func (ne *NumberExpression) expression() {}
func (ne *NumberExpression) String() string {
	return "Number Expression"
}

type GroupExpression struct {
	Expr Expression
}

func (ge *GroupExpression) expression() {}
func (ge *GroupExpression) String() string {
	return "Group Expression"
}

type VariableExpression struct {
	Name token.Token
}

func (ve *VariableExpression) expression() {}
func (ve *VariableExpression) String() string {
	return "VariableExpression"
}
