package gocalc

import (
	"fmt"
	"strings"
)

// Contain the token type in an integer format for better storage.
type TokenType int

const (
	// Contain the tangeable digits for a number token.
	Digits string = "1234567890"

	// Contain a list of tokens for our calculator.
	Integer TokenType = iota
	Add
	Multiply
	Minus
	Divide
	Modulus
	LeftParen
	RightParen
)

// Contain the data for a token object.
type Token struct {
	Type  TokenType
	Value string
}

// Create a new parser token object.
func NewToken(t TokenType, l string) *Token {
	return &Token{
		Type:  t,
		Value: l,
	}
}

// Contain the data required for a lexer object.
type Lexer struct {
	Code     []string
	Tokens   []*Token
	Current  *string
	Position int
}

// Create a new lexer object.
func NewLexer(c []string) *Lexer {
	return &Lexer{
		Code:     c,
		Tokens:   make([]*Token, 0),
		Current:  nil,
		Position: -1,
	}
}

// Advances onto the next character in the code.
func (l *Lexer) advance() {
	if l.Position+1 >= len(l.Code) {
		l.Current = nil
	} else {
		l.Position++
		l.Current = &l.Code[l.Position]
	}
}

// Create an integer/number token.
func (l *Lexer) createInteger() (*Token, error) {
	// Create an empty number token.
	var token *Token = NewToken(Integer, "")

	// Append the number values to the token.
	for l.Current != nil && strings.Contains(Digits+".", *l.Current) {
		token.Value += *l.Current
		l.advance()
	}

	// Check if the token was improperly defined.
	if strings.Count(token.Value, ".") > 1 {
		return nil, fmt.Errorf("invalid number defined under %s", token.Value)
	}

	return token, nil
}

// Convert the provided code into parsable tokens.
func (l *Lexer) Tokenise() error {
	l.advance()

	for l.Current != nil {
		switch *l.Current {
		case " ", "\t", "\r", "\n":
			l.advance()
		case "+":
			l.Tokens = append(l.Tokens, NewToken(Add, "+"))
			l.advance()
		case "-":
			l.Tokens = append(l.Tokens, NewToken(Minus, "-"))
			l.advance()
		case "*":
			l.Tokens = append(l.Tokens, NewToken(Multiply, "*"))
			l.advance()
		case "/":
			l.Tokens = append(l.Tokens, NewToken(Divide, "/"))
			l.advance()
		case "%":
			l.Tokens = append(l.Tokens, NewToken(Modulus, "%"))
			l.advance()
		case "(":
			l.Tokens = append(l.Tokens, NewToken(LeftParen, "("))
			l.advance()
		case ")":
			l.Tokens = append(l.Tokens, NewToken(RightParen, ")"))
			l.advance()
		default:
			if strings.Contains(Digits, *l.Current) {
				number, err := l.createInteger()
				if err != nil {
					return err
				}

				l.Tokens = append(l.Tokens, number)
				continue
			}

			return fmt.Errorf("unexpected token defined under %s", *l.Current)
		}
	}

	return nil
}
