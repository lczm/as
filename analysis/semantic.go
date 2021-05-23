package analysis

import (
	"github.com/lczm/as/ast"
)

type SemanticAnalyzer struct {
	statements []ast.Statement
}

// Run the analyzer, and update globals
func (s *SemanticAnalyzer) Analyze() {
}

func New(statements []ast.Statement) *SemanticAnalyzer {
	s := &SemanticAnalyzer{
		statements: statements,
	}
	return s
}
