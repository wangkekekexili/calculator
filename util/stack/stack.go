package stack

type stack struct {
	values []interface{}
}

func New() *stack {
	return &stack{values: make([]interface{}, 0)}
}

func (stack *stack) Peek() interface{} {
	if len(stack.values) == 0 {
		return nil
	}
	return stack.values[len(stack.values)-1]
}

func (stack *stack) Pop() interface{} {
	if len(stack.values) == 0 {
		return nil
	}
	v := stack.values[len(stack.values)-1]
	stack.values = stack.values[:len(stack.values)-1]
	return v
}

func (stack *stack) Push(v interface{}) {
	stack.values = append(stack.values, v)
}

func (stack *stack) Size() int {
	return len(stack.values)
}
