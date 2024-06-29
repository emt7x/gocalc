package gocalc

import "fmt"

// Contain the data for a node object.
type Node interface {
	Type() string
	String() string
}

// Contains the information for a number node.
type NumberNode struct {
	Value *Token
}

func (n *NumberNode) Type() string {
	return "number"
}

func (n *NumberNode) String() string {
	return n.Value.Value
}

// Contains the information for a unary operator node.
type UnaryOperator struct {
	Prefix *Token
	Value  Node
}

func (n *UnaryOperator) Type() string {
	return "unaryOperator"
}

func (n *UnaryOperator) String() string {
	return fmt.Sprintf("(%s%s)", n.Prefix.Value, n.Value)
}

// Contain the data for a binary expression object.
type BinaryEpxression struct {
	Left     Node
	Right    Node
	Operator *Token
}

func (b *BinaryEpxression) Type() string {
	return "binaryExpression"
}

func (b *BinaryEpxression) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Operator.Value, b.Right.String())
}

// Create a new term node.
func (p *Parser) term() (Node, error) {
	// Fetch the base factor for the AST.
	factor, err := p.factor()
	if err != nil {
		return nil, err
	}

	var node Node = factor

	// Range until we break out of the BODMAS sequence.
	for p.Current != nil && (p.Current.Type == Multiply || p.Current.Type == Divide || p.Current.Type == Modulus) {
		// Create a binary expression for multiplication, division and modulus.
		switch p.Current.Type {
		case Multiply:
			op := p.Current
			p.advance()
			factor, err := p.factor()
			if err != nil {
				return nil, err
			}

			node = &BinaryEpxression{Left: node, Operator: op, Right: factor}
		case Divide:
			op := p.Current
			p.advance()
			factor, err := p.factor()
			if err != nil {
				return nil, err
			}

			node = &BinaryEpxression{Left: node, Operator: op, Right: factor}
		case Modulus:
			op := p.Current
			p.advance()
			factor, err := p.factor()
			if err != nil {
				return nil, err
			}

			node = &BinaryEpxression{Left: node, Operator: op, Right: factor}
		}
	}

	return node, nil
}

// Create a new factorial node.
func (p *Parser) factor() (Node, error) {
	var token *Token = p.Current

	if token == nil {
		return nil, fmt.Errorf("expected next token, instead got nil")
	}

	// Return the node type for the specified token.
	switch token.Type {
	case LeftParen:
		p.advance()

		// Parse the expression inside of the brackets.
		expression, err := p.parseExpression()
		if err != nil {
			return nil, err
		}

		if p.Current == nil || p.Current.Type != RightParen {
			return nil, fmt.Errorf("expected closing parenthesis, instead got %v", p.Current)
		}

		p.advance() // Update current token after parsing the expression inside parentheses.

		return expression, nil
	case Add:
		token := p.Current
		p.advance()
		factor, err := p.factor()
		if err != nil {
			return nil, err
		}

		return &UnaryOperator{Value: factor, Prefix: token}, nil
	case Minus:
		token := p.Current
		p.advance()
		factor, err := p.factor()
		if err != nil {
			return nil, err
		}

		return &UnaryOperator{Value: factor, Prefix: token}, nil
	case Integer:
		p.advance()
		return &NumberNode{Value: token}, nil
	default:
		return nil, fmt.Errorf("invalid syntax")
	}
}

// Parse an expression into a binary node.
func (p *Parser) parseExpression() (Node, error) {
	// Create the base node for the expression as a term to follow BODMAS.
	term, err := p.term()
	if err != nil {
		return nil, err
	}

	var node Node = term

	// Range until the rules of BODMAS are broken.
	for p.Current != nil && (p.Current.Type == Add || p.Current.Type == Minus) {
		switch p.Current.Type {
		case Add:
			op := p.Current
			p.advance()
			term, err := p.term()
			if err != nil {
				return nil, err
			}

			node = &BinaryEpxression{Left: node, Operator: op, Right: term}
		case Minus:
			op := p.Current
			p.advance()
			term, err := p.term()
			if err != nil {
				return nil, err
			}

			node = &BinaryEpxression{Left: node, Operator: op, Right: term}
		}
	}

	return node, nil
}
