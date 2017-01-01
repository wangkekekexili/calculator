package calculator

import "math"

type interpreter struct {
	parser *parser
}

func newInterpreter(input string) *interpreter {
	return &interpreter{parser: newParser(input)}
}

func (i *interpreter) interpret() (float64, error) {
	root, err := i.parser.parse()
	if err != nil {
		return 0, err
	}
	return i.visit(root), nil
}

func (i *interpreter) visit(node *node) float64 {
	if node.left == nil && node.right == nil {
		return node.token.value
	}
	left := i.visit(node.left)
	right := i.visit(node.right)
	switch node.token.tokenType {
	case tokenTypePlus:
		return left + right
	case tokenTypeMinus:
		return left - right
	case tokenTypeMultiple:
		return left * right
	case tokenTypeDivide:
		return left / right
	case tokenTypePower:
		return math.Pow(left, right)
	}
	return 0
}
