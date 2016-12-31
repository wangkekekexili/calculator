package calculator

import (
	"fmt"
	"strings"
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
	tokenType tokenType
	value     float64
}

func (t *token) String() string {
	return fmt.Sprintf("token type: %v; value: %v", tokenTypeValueToString[t.tokenType], t.value)
}
