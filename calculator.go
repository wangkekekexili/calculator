package calculator

import "errors"

type tokenType int

var (
	errUnexpectedTokenType = errors.New("unexpected token type")
)

const (
	tokenTypeDoNotUse tokenType = iota
	tokenTypeNumber
	tokenTypeOperator
	tokenTypeEOF
	tokenTypeError
)

type token struct {
	value     float64
	operator  byte
	tokenType tokenType
}

type calculator struct {
	input           string
	currentPosition int
}

func newCalculator(input string) *calculator {
	return &calculator{input: input}
}

func (c *calculator) skipWhitespace() {
	for c.currentPosition < len(c.input) && c.input[c.currentPosition] == ' ' {
		c.currentPosition++
	}
}

func (c *calculator) getNextToken() *token {
	c.skipWhitespace()
	if c.currentPosition >= len(c.input) {
		return &token{tokenType: tokenTypeEOF}
	}
	ch := c.input[c.currentPosition]
	c.currentPosition++
	switch {
	case ch >= '0' && ch <= '9':
		return &token{value: float64(ch - '0'), tokenType: tokenTypeNumber}
	case ch == '+':
		return &token{operator: '+', tokenType: tokenTypeOperator}
	case ch == '-':
		return &token{operator: '-', tokenType: tokenTypeOperator}
	default:
		return &token{tokenType: tokenTypeError}
	}
}

func (c *calculator) getNextTokenWithExpectedType(expectedTokenType tokenType) (*token, error) {
	token := c.getNextToken()
	if token.tokenType != expectedTokenType {
		return nil, errUnexpectedTokenType
	}
	return token, nil
}

func (c *calculator) calculate() (float64, error) {
	first, err := c.getNextTokenWithExpectedType(tokenTypeNumber)
	if err != nil {
		return 0, err
	}
	operator, err := c.getNextTokenWithExpectedType(tokenTypeOperator)
	if err != nil {
		return 0, err
	}
	second, err := c.getNextTokenWithExpectedType(tokenTypeNumber)
	if err != nil {
		return 0, err
	}
	switch operator.operator {
	case '+':
		return first.value + second.value, nil
	case '-':
		return first.value - second.value, nil
	}
	return 0, nil
}

func Do(input string) float64 {
	result, err := newCalculator(input).calculate()
	if err != nil {
		panic(err)
	}
	return result
}
