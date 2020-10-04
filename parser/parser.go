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
	// This is a function declaration, it can be re-used to be parsed for methods as well.
	if p.match(token.FUNCTION) {
		return p.functionStatement("function")
	}
	if p.match(token.IF) {
		return p.ifStatement()
	}
	if p.match(token.PRINT) {
		return p.printStatement()
	}
	if p.match(token.RETURN) {
		return p.returnStatement()
	}
	if p.match(token.FOR) {
		return p.forStatement()
	}
	if p.match(token.WHILE) {
		return p.whileStatement()
	}
	if p.match(token.LBRACE) {
		return p.blockStatement()
	}

	return p.expressionStatement()
}

func (p *Parser) functionStatement(functionType string) ast.Statement {
	name := p.peek()
	if name.Type != token.IDENTIFIER {
		panic("Function name is not an identifier")
	}
	p.advance()

	// This will store the tokens in a function.
	// This is a token and not an expr array because
	// this does not evaluate.
	// function (a, b, c)
	var parameters []token.Token

	p.eat(token.LPAREN, "Expect '(' to start off function declaration")
	for !p.match(token.RPAREN) {
		parameter := p.peek()
		if parameter.Type != token.IDENTIFIER {
			panic("Expect identifiers within a function argument")
		}
		parameters = append(parameters, parameter)
		p.advance()
		if !p.match(token.COMMA) {
			break
		}
	}
	p.eat(token.RPAREN, "Expect ')' to end off function declaration")

	// Get the body of the function block statement
	// function(a, b, c) { }
	p.eat(token.LBRACE, "Expect '{' to start off the body of a function declaration")

	// Cast ast.Statement into a ast.BlockStatement
	body := p.blockStatement().(*ast.BlockStatement)

	functionStatement := &ast.FunctionStatement{
		Name:   name,
		Params: parameters,
		Body:   *body,
	}
	return functionStatement
}

// this function in the future should also support else if statements.
// this can be done by nesting if else {if else {if else}}
func (p *Parser) ifStatement() ast.Statement {
	// Condition
	p.eat(token.LPAREN, "Expect '(' to start off if condition")
	condition := p.expression()
	p.eat(token.RPAREN, "Expect ')' to end off if condition")

	// Then
	thenStatement := p.statement()

	// Else
	var elseStatement ast.Statement = nil
	if p.match(token.ELSE) {
		p.eat(token.LBRACE, "Expect '{' to start off if block")
		elseStatement = p.blockStatement()
	}

	ifStatement := &ast.IfStatement{
		Condition: condition,
		Then:      thenStatement,
		Else:      elseStatement,
	}

	return ifStatement
}

func (p *Parser) printStatement() ast.Statement {
	expr := p.expression()

	p.eat(token.SEMICOLON, "Expect ';'")

	printStatement := &ast.PrintStatement{
		Expr: expr,
	}
	return printStatement
}

func (p *Parser) returnStatement() ast.Statement {
	keyword := p.previous()

	// Default value to nil
	var value ast.Expression = nil

	if p.peek().Type != token.SEMICOLON {
		value = p.expression()
	}
	p.eat(token.SEMICOLON, "Expect ';' after return statement.")

	returnStatement := &ast.ReturnStatement{
		Keyword: keyword,
		Value:   value,
	}
	return returnStatement
}

func (p *Parser) forStatement() ast.Statement {
	p.eat(token.LPAREN, "Expect '(' after for.")

	// Variable section of for loops
	var variable ast.Statement = nil
	// for (var {x};
	if p.match(token.VAR) {
		// This will be a variableStatement
		variable = p.varDeclaration()
		if variable.(*ast.VariableStatement).Initializer == nil {
			panic("Cannot have uninitialized variable in a 'for' statement")
		}
	} else { // Existing variable, for({x};)
		variable = p.expressionStatement()
	}

	// Condition section of for loops
	var condition ast.Expression = nil
	condition = p.expression()
	p.eat(token.SEMICOLON, "Expect ';' after condition in 'for' statement")

	// Effect (Usually its the increment, but since it can also be decreasing
	// I thought that 'Effect' sounded better)
	var effect ast.Expression = nil
	effect = p.expression()

	p.eat(token.RPAREN, "Expect ')' after 'for' statement.")

	// Block statement
	body := p.statement()

	forStatement := &ast.ForStatement{
		Variable:  variable,
		Condition: condition,
		Effect:    effect,
		Body:      body,
	}
	return forStatement
}

func (p *Parser) whileStatement() ast.Statement {
	p.eat(token.LPAREN, "Expect '(' after while.")

	condition := p.expression()

	p.eat(token.RPAREN, "Expect, ')' after condition of while.")
	// This should most likely be a block statement
	body := p.statement()

	whileStatement := &ast.WhileStatement{
		Condition: condition,
		Body:      body,
	}
	return whileStatement
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
	expr := p.and()

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

func (p *Parser) and() ast.Expression {
	expr := p.or()

	for p.match(token.AND) {
		operator := p.previous()
		right := p.or()
		expr = &ast.LogicalExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
	}
	return expr
}

func (p *Parser) or() ast.Expression {
	expr := p.equality()

	for p.match(token.OR) {
		operator := p.previous()
		right := p.equality()
		expr = &ast.LogicalExpression{
			Left:     expr,
			Right:    right,
			Operator: operator,
		}
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
	for p.match(token.ASTERISK, token.SLASH, token.MODULUS) {
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
	return p.call()
}

func (p *Parser) call() ast.Expression {
	expr := p.primary()

	for {
		// If it is a left paren : '(<argument>, <argument>)
		if p.match(token.LPAREN) {
			var arguments []ast.Expression
			// While it is not a right paren to end off the function
			for !p.match(token.RPAREN) {
				// Add an argument as it is not empty
				arguments = append(arguments, p.expression())
				// If it does not match a comma, break out of this.
				if !p.match(token.COMMA) {
					break
				}
			}

			// Here is where the parser can check for other calls that may
			// be useful in the future.
			// i.e. Maximum len of arguments

			// Check that it ends with a ')'
			p.eat(token.RPAREN, "Expect ')' after arguments.")
			expr = &ast.CallExpression{
				Callee:    expr,
				Arguments: arguments,
			}
		} else {
			break
		}
	}
	return expr
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

func (p *Parser) advance() {
	p.current++
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
