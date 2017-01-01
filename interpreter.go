package calculator

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

// factor checks the token to see if it is a number.
func (c *interpreter) factor() (float64, error) {
	token := c.currentToken
	if err := c.eat(tokenTypeNumber); err != nil {
		return 0, err
	}
	return token.value, nil
}

func (c *interpreter) expr() (float64, error) {
	result, err := c.factor()
	if err != nil {
		return 0, err
	}
	for c.currentToken.tokenType == tokenTypePlus || c.currentToken.tokenType == tokenTypeMinus {
		switch c.currentToken.tokenType {
		case tokenTypePlus:
			c.eat(tokenTypePlus)
			n, err := c.factor()
			if err != nil {
				return 0, err
			}
			result = result + n
		case tokenTypeMinus:
			c.eat(tokenTypeMinus)
			n, err := c.factor()
			if err != nil {
				return 0, err
			}
			result = result - n
		}
	}
	if c.currentToken.tokenType != tokenTypeEOF {
		return 0, newUnexpectedTokenError(c.lexer.pos, c.currentToken, tokenTypeEOF)
	}
	return result, nil
}
