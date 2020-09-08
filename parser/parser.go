package parser

import "fmt"

type Parser struct{}

func (p *Parser) Parse() {
	fmt.Println("Parse...")
}

func New() *Parser {
	p := &Parser{}
	return p
}
