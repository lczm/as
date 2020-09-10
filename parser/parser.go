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

func (p *Parser) Parse() []ast.Expression {
	// var statements []ast.Statement
	var expressions []ast.Expression
	expressions = append(expressions, p.expression())

	return expressions
}

func (p *Parser) expression() ast.Expression {
	return p.equality()
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

	if p.match(token.LPAREN) {
		expr := p.expression()

		if p.peek().Type == token.RPAREN {
			p.current++
		} else {
			panic("Expect ')' after '('")
		}

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

func New(tokens []token.Token) *Parser {
	p := &Parser{
		current: 0,
		tokens:  tokens,
	}

	return p
}
