package calculator

type node struct {
	left, right *node
	token       *token
}

func newNode(left, right *node, token *token) *node {
	return &node{
		left:  left,
		right: right,
		token: token,
	}
}

type parser struct {
	lexer        *lexer
	currentToken *token
}

func newParser(input string) *parser {
	i := &parser{
		lexer: newLexer(input),
	}
	i.currentToken = i.lexer.GetToken()
	return i
}

// eat checks if the current token has the expected token type and fetches the next token.
func (c *parser) eat(t tokenType) error {
	if c.currentToken.tokenType != t {
		return newUnexpectedTokenError(c.lexer.pos, c.currentToken, t)
	}
	c.currentToken = c.lexer.GetToken()
	return nil
}

func (c *parser) number() (*node, error) {
	token := c.currentToken
	var result *node
	var err error

	switch token.tokenType {
	case tokenTypeNumber:
		c.eat(tokenTypeNumber)
		result = newNode(nil, nil, token)
	case tokenTypeLParen:
		c.eat(tokenTypeLParen)
		result, err = c.expr()
		if err != nil {
			return nil, err
		}
		if err = c.eat(tokenTypeRParen); err != nil {
			return nil, err
		}
	default:
		return nil, newUnexpectedTokenError(c.lexer.pos, token, tokenTypeNumber, tokenTypeLParen)
	}

	return result, nil
}

func (c *parser) factor() (*node, error) {
	result, err := c.number()
	if err != nil {
		return nil, err
	}

	op := c.currentToken
	if op.tokenType == tokenTypePower {
		c.eat(tokenTypePower)
		second, err := c.number()
		if err != nil {
			return nil, err
		}
		result = newNode(result, second, op)
	}

	return result, nil
}

func (c *parser) term() (*node, error) {
	result, err := c.factor()
	if err != nil {
		return nil, err
	}
	for c.currentToken.tokenType == tokenTypeMultiple || c.currentToken.tokenType == tokenTypeDivide {
		op := c.currentToken
		switch op.tokenType {
		case tokenTypeMultiple:
			c.eat(tokenTypeMultiple)
		case tokenTypeDivide:
			c.eat(tokenTypeDivide)
		}
		n, err := c.factor()
		if err != nil {
			return nil, err
		}
		result = newNode(result, n, op)
	}
	return result, nil
}

func (c *parser) expr() (*node, error) {
	result, err := c.term()
	if err != nil {
		return nil, err
	}
	for c.currentToken.tokenType == tokenTypePlus || c.currentToken.tokenType == tokenTypeMinus {
		op := c.currentToken
		switch op.tokenType {
		case tokenTypePlus:
			c.eat(tokenTypePlus)
		case tokenTypeMinus:
			c.eat(tokenTypeMinus)
		}
		n, err := c.term()
		if err != nil {
			return nil, err
		}
		result = newNode(result, n, op)
	}
	return result, nil
}

func (c *parser) parse() (*node, error) {
	n, err := c.expr()
	if err != nil {
		return nil, err
	}
	if err = c.eat(tokenTypeEOF); err != nil {
		return nil, err
	}
	return n, nil
}
