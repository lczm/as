package interpreter

import (
	"fmt"
	"github.com/lczm/as/ast"
)

type Interpreter struct {
	Expressions []ast.Expression
}

func (i Interpreter) Hello() {
	fmt.Println(len(i.Expressions))
}

func New(expressions []ast.Expression) *Interpreter {
	i := &Interpreter{
		Expressions: expressions,
	}
	return i
}
