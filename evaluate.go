package main

import (
	"fmt"
	"strconv"
)

// Contain the data for an evaulator object.
type Evaulater struct {
	Node Node
	New  Node
}

// Create a new evaluator object.
func NewEvaluator(n Node) *Evaulater {
	return &Evaulater{
		Node: n,
	}
}

// Contain the main number node.
type Number struct {
	Value float64
}

func (n *Number) Type() string   { return "number" }
func (n *Number) String() string { return fmt.Sprintf("%v", n.Value) }

// Evaulate the specific node.
func (e *Evaulater) Evaluate() (*Number, error) {
	return e.evalExpression(e.Node)
}

// Evaluate an expression into a final number node.
func (e *Evaulater) evalExpression(expr Node) (*Number, error) {
	switch expr.Type() {
	case "binaryExpression":
		left, err := e.evalExpression(expr.(*BinaryEpxression).Left)
		if err != nil {
			return nil, err
		}

		right, err := e.evalExpression(expr.(*BinaryEpxression).Right)
		if err != nil {
			return nil, err
		}

		// Create the mathematical value of the left and right values of the expression.
		switch expr.(*BinaryEpxression).Operator.Type {
		case Add:
			return &Number{Value: left.Value + right.Value}, nil
		case Minus:
			return &Number{Value: left.Value - right.Value}, nil
		case Multiply:
			return &Number{Value: left.Value * right.Value}, nil
		case Divide:
			if right.Value == 0 || left.Value == 0 {
				return nil, fmt.Errorf("integer cannot divide by zero")
			}

			return &Number{Value: left.Value / right.Value}, nil
		case Modulus:
			if right.Value == 0 || left.Value == 0 {
				return nil, fmt.Errorf("integer cannot divide by zero")
			}

			return &Number{Value: left.Value / right.Value}, nil
		default:
			println(expr.(*BinaryEpxression).Operator.Value)
			return nil, fmt.Errorf("invalid operator token")
		}
	case "unaryOperator":
		value, err := strconv.ParseFloat(expr.(*UnaryOperator).Value.(*NumberNode).Value.Value, 64)
		if err != nil {
			return nil, err
		}

		switch expr.(*UnaryOperator).Prefix.Type {
		case Add:
			return &Number{Value: +value}, nil
		case Minus:
			return &Number{Value: -value}, nil
		default:
			return nil, fmt.Errorf("invalid unary operation expression, should be '-' or '+'")
		}
	case "number":
		value, err := strconv.ParseFloat(expr.(*NumberNode).Value.Value, 64)
		if err != nil {
			return nil, err
		}

		return &Number{Value: value}, nil
	default:
		return nil, fmt.Errorf("unevaluable node provided")
	}
}
