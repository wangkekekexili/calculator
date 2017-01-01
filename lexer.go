package calculator

import "strconv"

type lexer struct {
	input string
	pos   int
}

func newLexer(input string) *lexer {
	return &lexer{input: input}
}

func (l *lexer) advance() {
	if l.pos < len(l.input) {
		l.pos++
	}
}

func (l *lexer) skipWhitespace() {
	for l.pos < len(l.input) && l.input[l.pos] == ' ' {
		l.pos++
	}
}

// getNumber gets the next number token. The caller should guarantee that the next token is a number.
func (l *lexer) getNumber() float64 {
	valueEndIndex := l.pos + 1
	for valueEndIndex < len(l.input) {
		ch := l.input[valueEndIndex]
		if ch >= '0' && ch <= '9' {
			valueEndIndex++
			continue
		}
		break
	}
	value, err := strconv.ParseFloat(l.input[l.pos:valueEndIndex], 64)
	if err != nil {
		panic(err)
	}
	l.pos = valueEndIndex
	return value
}

// GetToken gets the next token.
func (l *lexer) GetToken() *token {
	l.skipWhitespace()
	if l.pos >= len(l.input) {
		return &token{tokenType: tokenTypeEOF}
	}
	var resultToken *token
	ch := l.input[l.pos]
	switch {
	case ch >= '0' && ch <= '9':
		resultToken = &token{value: l.getNumber(), tokenType: tokenTypeNumber}
	case ch == '+':
		l.advance()
		resultToken = &token{tokenType: tokenTypePlus}
	case ch == '-':
		l.advance()
		resultToken = &token{tokenType: tokenTypeMinus}
	case ch == '*':
		l.advance()
		resultToken = &token{tokenType: tokenTypeMultiple}
	case ch == '/':
		l.advance()
		resultToken = &token{tokenType: tokenTypeDivide}
	case ch == '(':
		l.advance()
		resultToken = &token{tokenType: tokenTypeLParen}
	case ch == ')':
		l.advance()
		resultToken = &token{tokenType: tokenTypeRParen}
	default:
		resultToken = &token{tokenType: tokenTypeError}
	}
	return resultToken
}
