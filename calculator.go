package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type tokenType int

const (
	_ tokenType = iota
	tokenTypeNumber
	tokenTypePlus
	tokenTypeMinus
	tokenTypeEOF
	tokenTypeError
)

var (
	tokenTypeValueToString = map[tokenType]string{
		tokenTypeNumber: "number",
		tokenTypePlus:   "plus",
		tokenTypeMinus:  "minus",
		tokenTypeEOF:    "EOF",
		tokenTypeError:  "error",
	}

	numberTypeSet = map[tokenType]bool{
		tokenTypeNumber: true,
	}
	operatorTypeSet = map[tokenType]bool{
		tokenTypePlus:  true,
		tokenTypeMinus: true,
	}
)

type unexpectedTokenError struct {
	index              int
	expectedTokenTypes map[tokenType]bool
	token              *token
}

func newUnexpectedTokenError(index int, token *token, expectedTokenTypes map[tokenType]bool) *unexpectedTokenError {
	return &unexpectedTokenError{
		index:              index,
		expectedTokenTypes: expectedTokenTypes,
		token:              token,
	}
}

func (err *unexpectedTokenError) Error() string {
	var expectedTypeStrs []string
	for t := range err.expectedTokenTypes {
		expectedTypeStrs = append(expectedTypeStrs, tokenTypeValueToString[t])
	}
	targetTypeStr := tokenTypeValueToString[err.token.tokenType]
	return fmt.Sprintf("Index: %d. Expected to get tokens of type %v. Got token %v.", err.index, strings.Join(expectedTypeStrs, " or "), targetTypeStr)
}

type token struct {
	value     int
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
		return &token{tokenType: tokenTypePlus}
	case ch == '-':
		c.currentPosition++
		return &token{tokenType: tokenTypeMinus}
	default:
		return &token{tokenType: tokenTypeError}
	}
}

func (c *calculator) getNextTokenWithExpectedType(expectedTokenTypeSet map[tokenType]bool) (*token, error) {
	token := c.getNextToken()
	if _, ok := expectedTokenTypeSet[token.tokenType]; !ok {
		return nil, newUnexpectedTokenError(c.currentPosition, token, expectedTokenTypeSet)
	}
	return token, nil
}

func (c *calculator) calculate() (int, error) {
	first, err := c.getNextTokenWithExpectedType(numberTypeSet)
	if err != nil {
		return 0, err
	}
	operator, err := c.getNextTokenWithExpectedType(operatorTypeSet)
	if err != nil {
		return 0, err
	}
	second, err := c.getNextTokenWithExpectedType(numberTypeSet)
	if err != nil {
		return 0, err
	}
	switch operator.tokenType {
	case tokenTypePlus:
		return first.value + second.value, nil
	case tokenTypeMinus:
		return first.value - second.value, nil
	}
	return 0, nil
}

// Do performs arithmetic calculation based on the input string.
func Do(input string) (int, error) {
	return newCalculator(input).calculate()
}
