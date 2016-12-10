package calculator

import (
	"errors"
	"strconv"
)

type tokenType int

var (
	errUnexpectedTokenType = errors.New("unexpected token type")
)

const (
	_ tokenType = iota
	tokenTypeNumber
	tokenTypeOperator
	tokenTypeEOF
	tokenTypeError
)

var (
	tokenTypeValueToString = map[tokenType]string{
		tokenTypeNumber:   "number",
		tokenTypeOperator: "operator",
		tokenTypeEOF:      "EOF",
		tokenTypeError:    "error",
	}
)

type token struct {
	value     int
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
	switch {
	case ch >= '0' && ch <= '9':
		valueEndIndex := c.currentPosition + 1
		for valueEndIndex < len(c.input) {
			ch = c.input[valueEndIndex]
			if ch >= '0' && ch <= '9' {
				valueEndIndex++
				continue
			}
			break
		}
		value, _ := strconv.ParseInt(c.input[c.currentPosition:valueEndIndex], 10, 32)
		c.currentPosition = valueEndIndex
		return &token{value: int(value), tokenType: tokenTypeNumber}
	case ch == '+':
		c.currentPosition++
		return &token{operator: '+', tokenType: tokenTypeOperator}
	case ch == '-':
		c.currentPosition++
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

func (c *calculator) calculate() (int, error) {
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

func Do(input string) int {
	result, err := newCalculator(input).calculate()
	if err != nil {
		panic(err)
	}
	return result
}
