package utils

import "errors"

type Stack struct {
	Buffer []interface{}
}

func (stack *Stack) GetSize() int {
	return len(stack.Buffer)
}

func (stack *Stack) Push(newElement interface{}) {
	stack.Buffer = append(stack.Buffer, newElement)
}

func (stack *Stack) Top() (interface{}, error) {
	size := stack.GetSize()
	if size == 0 {
		return nil, errors.New("Stack is empty")
	}
	return stack.Buffer[size-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	size := stack.GetSize()
	if size == 0 {
		return nil, errors.New("Stack is empty")
	}
	popped := stack.Buffer[size-1]
	stack.Buffer = stack.Buffer[:size-1]
	return popped, nil
}
