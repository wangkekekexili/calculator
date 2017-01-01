package calculator

import "math"

type interpreter struct {
	lexer        *lexer
	currentToken *token
}

func newInterpreter(input string) *interpreter {
	i := &interpreter{
		lexer: newLexer(input),
	}
	i.currentToken = i.lexer.GetToken()
	return i
}

// eat checks if the current token has the expected token type and fetches the next token.
func (c *interpreter) eat(t tokenType) error {
	if c.currentToken.tokenType != t {
		return newUnexpectedTokenError(c.lexer.pos, c.currentToken, t)
	}
	c.currentToken = c.lexer.GetToken()
	return nil
}

func (c *interpreter) number() (float64, error) {
	token := c.currentToken
	var result float64
	var err error

	switch token.tokenType {
	case tokenTypeNumber:
		c.eat(tokenTypeNumber)
		result = token.value
	case tokenTypeLParen:
		c.eat(tokenTypeLParen)
		result, err = c.expr()
		if err != nil {
			return 0, err
		}
		if err = c.eat(tokenTypeRParen); err != nil {
			return 0, err
		}
	default:
		return 0, newUnexpectedTokenError(c.lexer.pos, token, tokenTypeNumber, tokenTypeLParen)
	}

	return result, nil
}

func (c *interpreter) factor() (float64, error) {
	result, err := c.number()
	if err != nil {
		return 0, err
	}

	if c.currentToken.tokenType == tokenTypePower {
		c.eat(tokenTypePower)
		second, err := c.number()
		if err != nil {
			return 0, err
		}
		result = math.Pow(result, second)
	}

	return result, nil
}

func (c *interpreter) term() (float64, error) {
	result, err := c.factor()
	if err != nil {
		return 0, err
	}
	for c.currentToken.tokenType == tokenTypeMultiple || c.currentToken.tokenType == tokenTypeDivide {
		switch c.currentToken.tokenType {
		case tokenTypeMultiple:
			c.eat(tokenTypeMultiple)
			n, err := c.factor()
			if err != nil {
				return 0, err
			}
			result = result * n
		case tokenTypeDivide:
			c.eat(tokenTypeDivide)
			n, err := c.factor()
			if err != nil {
				return 0, err
			}
			result = result / n
		}
	}
	return result, nil
}

func (c *interpreter) expr() (float64, error) {
	result, err := c.term()
	if err != nil {
		return 0, err
	}
	for c.currentToken.tokenType == tokenTypePlus || c.currentToken.tokenType == tokenTypeMinus {
		switch c.currentToken.tokenType {
		case tokenTypePlus:
			c.eat(tokenTypePlus)
			n, err := c.term()
			if err != nil {
				return 0, err
			}
			result = result + n
		case tokenTypeMinus:
			c.eat(tokenTypeMinus)
			n, err := c.term()
			if err != nil {
				return 0, err
			}
			result = result - n
		}
	}
	return result, nil
}

func (c *interpreter) calculate() (float64, error) {
	value, err := c.expr()
	if err != nil {
		return 0, err
	}
	if err = c.eat(tokenTypeEOF); err != nil {
		return 0, err
	}
	return value, nil
}
