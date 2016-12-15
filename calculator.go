package calculator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wangkekekexili/calculator/util/stack"
)

type tokenType int

const (
	_ tokenType = iota
	tokenTypeNumber
	tokenTypePlus
	tokenTypeMinus
	tokenTypeMultiple
	tokenTypeDivide
	tokenTypeEOF
	tokenTypeError
)

var (
	tokenTypeValueToString = map[tokenType]string{
		tokenTypeNumber:   "number",
		tokenTypePlus:     "plus",
		tokenTypeMinus:    "minus",
		tokenTypeMultiple: "multiple",
		tokenTypeDivide:   "devide",
		tokenTypeEOF:      "EOF",
		tokenTypeError:    "error",
	}
)

type unexpectedTokenError struct {
	index              int
	expectedTokenTypes []tokenType
	token              *token
}

func newUnexpectedTokenError(index int, token *token, expectedTokenTypes ...tokenType) *unexpectedTokenError {
	return &unexpectedTokenError{
		index:              index,
		expectedTokenTypes: expectedTokenTypes,
		token:              token,
	}
}

func (err *unexpectedTokenError) Error() string {
	targetTypeStr := tokenTypeValueToString[err.token.tokenType]
	if len(err.expectedTokenTypes) == 0 {
		return fmt.Sprintf("Index: %d. Unexpected token %v.", err.index, targetTypeStr)
	}
	var expectedTypeStrs []string
	for _, t := range err.expectedTokenTypes {
		expectedTypeStrs = append(expectedTypeStrs, tokenTypeValueToString[t])
	}
	return fmt.Sprintf("Index: %d. Expected to get tokens of type %v. Got token %v.", err.index, strings.Join(expectedTypeStrs, " or "), targetTypeStr)
}

type token struct {
	value     float64
	tokenType tokenType
}

func (t *token) String() string {
	return fmt.Sprintf("token type: %v; value: %v", tokenTypeValueToString[t.tokenType], t.value)
}

type calculator struct {
	input           string
	currentPosition int
}

func newCalculator(input string) *calculator {
	return &calculator{input: input}
}

func (c *calculator) advance() {
	if c.currentPosition < len(c.input) {
		c.currentPosition++
	}
}

func (c *calculator) skipWhitespace() {
	for c.currentPosition < len(c.input) && c.input[c.currentPosition] == ' ' {
		c.currentPosition++
	}
}

// getNextNumber gets the next number token. The caller should guarantee that the next token is a number.
func (c *calculator) getNextNumber() float64 {
	valueEndIndex := c.currentPosition + 1
	for valueEndIndex < len(c.input) {
		ch := c.input[valueEndIndex]
		if ch >= '0' && ch <= '9' {
			valueEndIndex++
			continue
		}
		break
	}
	value, err := strconv.ParseInt(c.input[c.currentPosition:valueEndIndex], 10, 32)
	if err != nil {
		panic(err)
	}
	c.currentPosition = valueEndIndex
	return float64(value)
}

func (c *calculator) getNextToken() *token {
	c.skipWhitespace()
	if c.currentPosition >= len(c.input) {
		return &token{tokenType: tokenTypeEOF}
	}
	ch := c.input[c.currentPosition]
	switch {
	case ch >= '0' && ch <= '9':
		return &token{value: c.getNextNumber(), tokenType: tokenTypeNumber}
	case ch == '+':
		c.advance()
		return &token{tokenType: tokenTypePlus}
	case ch == '-':
		c.advance()
		return &token{tokenType: tokenTypeMinus}
	case ch == '*':
		c.advance()
		return &token{tokenType: tokenTypeMultiple}
	case ch == '/':
		c.advance()
		return &token{tokenType: tokenTypeDivide}
	default:
		return &token{tokenType: tokenTypeError}
	}
}

func (c *calculator) getNextTokenWithExpectedType(expectedTokenType tokenType) (*token, error) {
	nextToken := c.getNextToken()
	if nextToken.tokenType != expectedTokenType {
		return nil, newUnexpectedTokenError(c.currentPosition, nextToken, expectedTokenType)
	}
	return nextToken, nil
}

func (c *calculator) calculate() (float64, error) {
	s := stack.New()
	firstNumber, err := c.getNextTokenWithExpectedType(tokenTypeNumber)
	if err != nil {
		return 0, err
	}
	s.Push(firstNumber)
	for {
		nextToken := c.getNextToken()
		switch nextToken.tokenType {
		case tokenTypeEOF:
			var result float64
			reverse := stack.New()
			for s.Size() != 0 {
				reverse.Push(s.Pop())
			}
			for reverse.Size() != 0 {
				first := reverse.Pop().(*token)
				if first.tokenType == tokenTypeNumber {
					result = first.value
				} else if first.tokenType == tokenTypePlus {
					second := reverse.Pop().(*token)
					result += second.value
				} else if first.tokenType == tokenTypeMinus {
					second := reverse.Pop().(*token)
					result -= second.value
				}
			}
			return result, nil
		case tokenTypeNumber:
			return 0, newUnexpectedTokenError(c.currentPosition, nextToken)
		case tokenTypePlus, tokenTypeMinus:
			s.Push(nextToken)
			nextNumberToken, err := c.getNextTokenWithExpectedType(tokenTypeNumber)
			if err != nil {
				return 0, err
			}
			s.Push(nextNumberToken)
		case tokenTypeMultiple, tokenTypeDivide:
			nextNumberToken, err := c.getNextTokenWithExpectedType(tokenTypeNumber)
			if err != nil {
				return 0, err
			}
			lastNumberToken := s.Pop().(*token)
			if lastNumberToken.tokenType != tokenTypeNumber {
				return 0, newUnexpectedTokenError(c.currentPosition, nextNumberToken)
			}
			if nextToken.tokenType == tokenTypeMultiple {
				s.Push(&token{value: lastNumberToken.value * nextNumberToken.value, tokenType: tokenTypeNumber})
			} else {
				s.Push(&token{value: lastNumberToken.value / nextNumberToken.value, tokenType: tokenTypeNumber})
			}
		}
	}
}

// Do performs arithmetic calculation based on the input string.
func Do(input string) (float64, error) {
	return newCalculator(input).calculate()
}
