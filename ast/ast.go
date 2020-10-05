package ast

import (
	"fmt"

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

type ForStatement struct {
	Variable  Statement
	Condition Expression
	Effect    Expression
	Body      Statement
}

func (fs *ForStatement) statement() {}

type BlockStatement struct {
	Statements []Statement
}

func (bs *BlockStatement) statement() {}

type FunctionStatement struct {
	Name   token.Token
	Params []token.Token
	Body   BlockStatement
}

func (fs *FunctionStatement) statement() {}

type ReturnStatement struct {
	Keyword token.Token
	Value   Expression
}

func (rs *ReturnStatement) statement() {}

// Expressions
type AssignmentExpression struct {
	Name  token.Token
	Value Expression
}

func (ae *AssignmentExpression) expression() {}
func (ae *AssignmentExpression) String() string {
	return fmt.Sprintf("(AssignmentExpression) Name : %s, Value : %s\n",
		ae.Name, ae.Value.String())
}

type BinaryExpression struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (be *BinaryExpression) expression() {}
func (be *BinaryExpression) String() string {
	return fmt.Sprintf("(BinaryExpression) Left : %s, Right : %s, Operator : %s\n",
		be.Left.String(), be.Right.String(), be.Operator.Literal)
}

type UnaryExpression struct {
	Right    Expression
	Operator token.Token
}

func (ue *UnaryExpression) expression() {}
func (ue *UnaryExpression) String() string {
	return fmt.Sprintf("(UnaryExpression) Right : %s, Operator : %s\n",
		ue.Right.String(), ue.Operator.Literal)
}

type LogicalExpression struct {
	Left     Expression
	Right    Expression
	Operator token.Token
}

func (le *LogicalExpression) expression() {}
func (le *LogicalExpression) String() string {
	return fmt.Sprintf("(LogicalExpression) Left : %s, Right : %s, Operator : %s\n",
		le.Left.String(), le.Right.String(), le.Operator.Literal)
}

type NumberExpression struct {
	Value int
}

func (ne *NumberExpression) expression() {}
func (ne *NumberExpression) String() string {
	return fmt.Sprintf("(NumberExpression) Value : %d\n",
		ne.Value)
}

type StringExpression struct {
	Value string
}

func (se *StringExpression) expression() {}
func (se *StringExpression) String() string {
	return fmt.Sprintf("(StringExpression) Value : %s\n",
		se.Value)
}

type GroupExpression struct {
	Expr Expression
}

func (ge *GroupExpression) expression() {}
func (ge *GroupExpression) String() string {
	return fmt.Sprintf("(GroupExpression) Expr : %s\n",
		ge.Expr.String())
}

type VariableExpression struct {
	Name token.Token
}

func (ve *VariableExpression) expression() {}
func (ve *VariableExpression) String() string {
	return fmt.Sprintf("(VariableExpression) Name : %s\n",
		ve.Name.Literal)
}

type CallExpression struct {
	Callee    Expression
	Arguments []Expression
}

func (ce *CallExpression) expression() {}
func (ce *CallExpression) String() string {
	var arguments []string
	for i := 0; i < len(ce.Arguments); i++ {
		arguments = append(arguments, ce.Arguments[i].String())
	}
	return fmt.Sprintf("(CallExpression) Callee : %s, Arguments : %s\n",
		ce.Callee.String(), arguments)
}
