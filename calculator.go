package calculator

import "strconv"

type interpreter struct {
	input        string
	pos          int
	currentToken *token
}

func newInterpreter(input string) *interpreter {
	c := &interpreter{input: input}
	c.getToken()
	return c
}

func (c *interpreter) advance() {
	if c.pos < len(c.input) {
		c.pos++
	}
}

func (c *interpreter) skipWhitespace() {
	for c.pos < len(c.input) && c.input[c.pos] == ' ' {
		c.pos++
	}
}

// getNumber gets the next number token. The caller should guarantee that the next token is a number.
func (c *interpreter) getNumber() float64 {
	valueEndIndex := c.pos + 1
	for valueEndIndex < len(c.input) {
		ch := c.input[valueEndIndex]
		if ch >= '0' && ch <= '9' {
			valueEndIndex++
			continue
		}
		break
	}
	value, err := strconv.ParseInt(c.input[c.pos:valueEndIndex], 10, 32)
	if err != nil {
		panic(err)
	}
	c.pos = valueEndIndex
	return float64(value)
}

func (c *interpreter) getToken() {
	c.skipWhitespace()
	if c.pos >= len(c.input) {
		c.currentToken = &token{tokenType: tokenTypeEOF}
		return
	}
	ch := c.input[c.pos]
	switch {
	case ch >= '0' && ch <= '9':
		c.currentToken = &token{value: c.getNumber(), tokenType: tokenTypeNumber}
	case ch == '+':
		c.advance()
		c.currentToken = &token{tokenType: tokenTypePlus}
	case ch == '-':
		c.advance()
		c.currentToken = &token{tokenType: tokenTypeMinus}
	case ch == '*':
		c.advance()
		c.currentToken = &token{tokenType: tokenTypeMultiple}
	case ch == '/':
		c.advance()
		c.currentToken = &token{tokenType: tokenTypeDivide}
	default:
		c.currentToken = &token{tokenType: tokenTypeError}
	}
}

func (c *interpreter) getCurrentToken() *token {
	return c.currentToken
}

// eat checks if the current token has the expected token type and fetches the next token.
func (c *interpreter) eat(t tokenType) error {
	if c.currentToken.tokenType != t {
		return newUnexpectedTokenError(c.pos, c.currentToken, t)
	}
	c.getToken()
	return nil
}

func (c *interpreter) calculate() (float64, error) {
	firstNumber := c.getCurrentToken()
	if err := c.eat(tokenTypeNumber); err != nil {
		return 0, err
	}
	operator := c.getCurrentToken()
	switch operator.tokenType {
	case tokenTypePlus:
		c.eat(tokenTypePlus)
	case tokenTypeMinus:
		c.eat(tokenTypeMinus)
	case tokenTypeMultiple:
		c.eat(tokenTypeMultiple)
	case tokenTypeDivide:
		c.eat(tokenTypeDivide)
	default:
		return 0, newUnexpectedTokenError(c.pos, c.currentToken, tokenTypePlus, tokenTypeMinus, tokenTypeMultiple, tokenTypeDivide)
	}
	secondNumber := c.getCurrentToken()
	if err := c.eat(tokenTypeNumber); err != nil {
		return 0, err
	}

	switch operator.tokenType {
	case tokenTypePlus:
		return firstNumber.value + secondNumber.value, nil
	case tokenTypeMinus:
		return firstNumber.value - secondNumber.value, nil
	case tokenTypeMultiple:
		return firstNumber.value * secondNumber.value, nil
	case tokenTypeDivide:
		return firstNumber.value / secondNumber.value, nil
	}
	return 0, nil
}

// Do performs arithmetic calculation based on the input string.
func Do(input string) (float64, error) {
	return newInterpreter(input).calculate()
}
