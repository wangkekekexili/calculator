package calculator

import (
	"fmt"
	"math"
	"reflect"
	"sync"
)

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
	return visit(root), nil
}

var nodeTypeNameToVisitFunc map[string]func(n node) float64
var nodeTypeNameToVisitFuncInit sync.Once

func getNodeTypeNameToVisitFunc() map[string]func(n node) float64 {
	nodeTypeNameToVisitFuncInit.Do(func() {
		nodeTypeNameToVisitFunc = map[string]func(n node) float64{
			reflect.TypeOf(binaryOperatorNode{}).Name(): visitBinaryOperatorNode,
			reflect.TypeOf(valueNode{}).Name():          visitValueNode,
		}
	})
	return nodeTypeNameToVisitFunc
}

func visit(n node) float64 {
	return getNodeTypeNameToVisitFunc()[n.getTypeName()](n)
}

func visitBinaryOperatorNode(n node) float64 {
	opNode := n.(*binaryOperatorNode)
	leftValue := visit(opNode.left)
	rightValue := visit(opNode.right)
	var value float64
	switch opNode.token.tokenType {
	case tokenTypePlus:
		value = leftValue + rightValue
	case tokenTypeMinus:
		value = leftValue - rightValue
	case tokenTypeMultiple:
		value = leftValue * rightValue
	case tokenTypeDivide:
		value = leftValue / rightValue
	case tokenTypePower:
		value = math.Pow(leftValue, rightValue)
	default:
		panic(fmt.Sprintf("programming error unexpected token: %v", opNode.token))
	}
	return value
}

func visitValueNode(n node) float64 {
	valueNode := n.(*valueNode)
	return valueNode.token.value
}
