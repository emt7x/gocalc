package gocalc

import "fmt"

// Contain the data required for a parser object.
type Parser struct {
	Tokens   []*Token
	Current  *Token
	Position int
}

// Create a new parser method.
func NewParser(l *Lexer) *Parser {
	return &Parser{
		Tokens:   l.Tokens,
		Current:  nil,
		Position: -1,
	}
}

// Advances onto the next token.
func (p *Parser) advance() {
	if p.Position+1 >= len(p.Tokens) {
		p.Current = nil
	} else {
		p.Position++
		p.Current = p.Tokens[p.Position]
	}
}

// Parse all of the provided parser tokens.
func (p *Parser) Parse() (Node, error) {
	p.advance()

	// Check if there are any tokens to parse.
	if len(p.Tokens) < 1 {
		return nil, nil
	}

	node, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	// Check if there has been a syntax error.
	if p.Current != nil {
		return nil, fmt.Errorf("invalid expression syntax")
	}

	return node, nil
}
