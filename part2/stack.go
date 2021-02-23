package part2

import "errors"

type Stack struct {
	buffer []interface{}
}

func (stack *Stack) GetSize() int {
	return len(stack.buffer)
}

func (stack *Stack) Push(newElement interface{}) {
	stack.buffer = append(stack.buffer, newElement)
}

func (stack *Stack) Top() (interface{}, error) {
	size := stack.GetSize()
	if size == 0 {
		return nil, errors.New("Stack is empty")
	}
	return stack.buffer[size-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	size := stack.GetSize()
	if size == 0 {
		return nil, errors.New("Stack is empty")
	}
	popped := stack.buffer[size-1]
	stack.buffer = stack.buffer[:size-1]
	return popped, nil
}
