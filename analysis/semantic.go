package analysis

import (
	"fmt"

	"github.com/lczm/as/ast"
)

type SemanticAnalyzer struct {
	// Array of variable mapping (Stack)
	variables  []map[string]bool
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
		_ = s.resolveVariable(node)
	}
}

// Resolves a variable into the scope
func (s *SemanticAnalyzer) resolveVariable(node *ast.VariableStatement) bool {
	var scope *map[string]bool
	for i := 0; i < len(s.variables)-1; i++ {
		scope = s.getScope(len(s.variables) - i - 1)
		_, found := (*scope)[node.Name.Literal]
		if found {
			// Shadow warning, variable name has been declared already but now it is declared again
			fmt.Printf("Shadow warning at line %d, Declaring an already declared variable: \"%s\"", node.Name.Line, node.Name.Literal)
			return false
		}
	}
	(*scope)[node.Name.Literal] = true
	return true
}

func (s *SemanticAnalyzer) pushScope() {
	s.variables = append(s.variables, make(map[string]bool))
}

func (s *SemanticAnalyzer) popScope() {
	// There should always be one scope, the global scope
	// So if there are exists more than one scope; then it is
	// safe to pop the scope out
	if len(s.variables) > 1 {
		s.variables = s.variables[:len(s.variables)-1]
	}
}

func (s *SemanticAnalyzer) getScope(index int) *map[string]bool {
	// Check if out of bounds
	if index < 0 || index > len(s.variables)-1 {
		panic("Tried to get a scope that is out of bounds")
	}
	return &s.variables[index]
}

func (s *SemanticAnalyzer) getLastScope() *map[string]bool {
	// This can include the global scope, unlike popScope()
	if len(s.variables) >= 1 {
		s.variables = s.variables[:len(s.variables)-1]
		return &s.variables[len(s.variables)-1]
	}
	// If this happens, something bugged out in the analyzer
	// just panic out of this.
	panic("Analyzer tried to get last scope, but there are no scopes remaining.")
}

func New(statements []ast.Statement) *SemanticAnalyzer {
	s := &SemanticAnalyzer{
		variables:  make([]map[string]bool, 1),
		statements: statements,
	}
	// Create global scope
	s.variables = append(s.variables, make(map[string]bool))
	return s
}
