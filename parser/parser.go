package parser

import (
	"strconv"

	"github.com/lczm/as/ast"
	"github.com/lczm/as/token"
)

type Parser struct {
	current int
	tokens  []token.Token
}

func (p *Parser) Parse() []ast.Statement {
	// var expressions []ast.Expression
	// expressions = append(expressions, p.expression())

	var statements []ast.Statement
	for p.current != len(p.tokens) {
		statements = append(statements, p.declaration())
	}

	return statements
}

func (p *Parser) declaration() ast.Statement {
	if p.match(token.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) varDeclaration() ast.Statement {
	p.eat(token.IDENTIFIER, "Expect variable name")
	name := p.previous()

	// If there is an equals, this is an initializer
	// e.g. : var a = 2;
	variableStatement := &ast.VariableStatement{}
	if p.match(token.ASSIGN) {
		initializer := p.expression()
		p.eat(token.SEMICOLON, "Expect ';' after variable declaration'")

		variableStatement.Name = name
		variableStatement.Initializer = initializer
	} else { // If there is no equals, still have to check for ';'
		p.eat(token.SEMICOLON, "Expect ';' after variable declaration'")

		variableStatement.Name = name
		variableStatement.Initializer = nil
	}
	return variableStatement
}

func (p *Parser) statement() ast.Statement {
	if p.match(token.IF) {
		return p.ifStatement()
	}
	if p.match(token.PRINT) {
		return p.printStatement()
	}
	if p.match(token.LBRACE) {
		return p.blockStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) ifStatement() ast.Statement {
	return nil
}

func (p *Parser) printStatement() ast.Statement {
	expr := p.expression()

	p.eat(token.SEMICOLON, "Expect ';'")

	printStatement := &ast.PrintStatement{
		Expr: expr,
	}
	return printStatement
}

func (p *Parser) blockStatement() ast.Statement {
	var statements []ast.Statement

	// Keep going until it hits the right brace - '}'.
	for p.peek().Type != token.RBRACE {
		statements = append(statements, p.declaration())
	}

	// Once the right brace is hit, move the parser past the
	// right brace
	p.eat(token.RBRACE, "Expect '}' after block.")

	blockStatement := &ast.BlockStatement{
		Statements: statements,
	}
	return blockStatement
}

func (p *Parser) expressionStatement() ast.Statement {
	expr := p.expression()

	p.eat(token.SEMICOLON, "Expect ';'")

	statementExpression := &ast.StatementExpression{
		Expr: expr,
	}
	return statementExpression
}

func (p *Parser) expression() ast.Expression {
	return p.assignment()
}

func (p *Parser) assignment() ast.Expression {
	expr := p.equality()

	// Match for assignment
	for p.match(token.ASSIGN) {
		assignment := p.previous()
		value := p.assignment()

		if varExpr, ok := expr.(*ast.VariableExpression); ok {
			return &ast.AssignmentExpression{
				Name:  varExpr.Name,
				Value: value,
			}
		}

		// Error out here.
		panic(assignment)
	}
	return expr
}

func (p *Parser) equality() ast.Expression {
	expr := p.comparison()

	// Match for equality
	for p.match(token.NOT_EQ, token.EQ) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
	}
	return expr
}

func (p *Parser) comparison() ast.Expression {
	expr := p.addition()

	// Match for comparison
	for p.match(token.GT, token.GT_EQ, token.LT, token.LT_EQ) {
		operator := p.previous()
		right := p.addition()
		expr = &ast.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
	}
	return expr
}

func (p *Parser) addition() ast.Expression {
	expr := p.multiplication()

	// Match for addition
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right := p.multiplication()
		expr = &ast.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
	}
	return expr
}

func (p *Parser) multiplication() ast.Expression {
	expr := p.unary()

	// Match for multiplication
	for p.match(token.ASTERISK, token.SLASH) {
		operator := p.previous()
		right := p.unary()
		expr = &ast.BinaryExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
	}
	return expr
}

func (p *Parser) unary() ast.Expression {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &ast.UnaryExpression{
			Right:    right,
			Operator: operator,
		}
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expression {
	if p.match(token.NUMBER) {
		i, err := strconv.Atoi(p.previous().Literal)
		if err != nil {
			panic(err)
		}
		return &ast.NumberExpression{
			Value: i,
		}
	}

	if p.match(token.IDENTIFIER) {
		return &ast.VariableExpression{
			Name: p.previous(),
		}
	}

	if p.match(token.LPAREN) {
		expr := p.expression()

		p.eat(token.RPAREN, "Expect ')' after '('")
		return &ast.GroupExpression{
			Expr: expr,
		}
	}
	return nil
}

func (p *Parser) match(tokens ...token.TokenType) bool {
	if p.current < 0 || p.current >= len(p.tokens) {
		return false
	}
	for _, token := range tokens {
		if p.tokens[p.current].Type == token {
			p.current++
			return true
		}
	}
	return false
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	// Minus one as this is referring to the 'previous'
	current := p.current - 1

	// Check == len(p.tokens) as it is length, so it starts from 0
	if current < 0 || current == len(p.tokens) {
		panic("Parsing an out of range index")
	}

	return p.tokens[current]
}

func (p *Parser) eat(tokenType token.TokenType, message string) {
	if p.peek().Type == tokenType {
		p.current++
		return
	}

	// TODO : Throw an error with the message that is passed in
	panic(message)
}

func New(tokens []token.Token) *Parser {
	p := &Parser{
		current: 0,
		tokens:  tokens,
	}

	return p
}
