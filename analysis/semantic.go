package analysis

import (
	"fmt"

	"github.com/lczm/as/ast"
)

type SemanticAnalyzer struct {
	variables  map[string]bool
	statements []ast.Statement
}

// Run the analyzer, and update globals
func (s *SemanticAnalyzer) Analyze() {
	for _, stmt := range s.statements {
		s.Eval(stmt)
	}
}

func (s *SemanticAnalyzer) Eval(astNode ast.AstNode) {
	switch node := astNode.(type) {
	case *ast.VariableStatement:
		_, found := s.variables[node.Name.Literal]
		if found {
			// Shadow warning
			fmt.Printf("Shadow warning at line %d, Declaring an already declared variable: \"%s\"", node.Name.Line, node.Name.Literal)
		} else {
			s.variables[node.Name.Literal] = true
		}
	}
}

func New(statements []ast.Statement) *SemanticAnalyzer {
	s := &SemanticAnalyzer{
		variables:  make(map[string]bool),
		statements: statements,
	}
	return s
}
