package calculator

import "reflect"

type node interface {
	getToken() *token
	getTypeName() string
}

type unaryOperatorNode struct {
	child node
	token *token
}

var _ node = &unaryOperatorNode{}

func newUnaryOperatorNode(child node, token *token) *unaryOperatorNode {
	return &unaryOperatorNode{
		child: child,
		token: token,
	}
}

func (n *unaryOperatorNode) getToken() *token {
	return n.token
}

func (n *unaryOperatorNode) getTypeName() string {
	return reflect.TypeOf(*n).Name()
}

type binaryOperatorNode struct {
	left, right node
	token       *token
}

var _ node = &binaryOperatorNode{}

func newBinaryOperatorNode(left, right node, token *token) *binaryOperatorNode {
	return &binaryOperatorNode{
		left:  left,
		right: right,
		token: token,
	}
}

func (n *binaryOperatorNode) getToken() *token {
	return n.token
}

func (n *binaryOperatorNode) getTypeName() string {
	return reflect.TypeOf(*n).Name()
}

type valueNode struct {
	token *token
}

var _ node = &valueNode{}

func newValueNode(token *token) *valueNode {
	return &valueNode{token: token}
}

func (n *valueNode) getToken() *token {
	return n.token
}

func (n *valueNode) getTypeName() string {
	return reflect.TypeOf(*n).Name()
}
