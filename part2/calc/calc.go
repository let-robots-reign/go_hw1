package calc

import (
	"errors"
	"go_hw1/part2/utils"
	"strings"
)

func getOperationPriority(op rune) int {
	switch op {
	case '(':
		return 0
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return -1
	}
}

func applyOperation(lhs float64, rhs float64, op string) (float64, error) {
	switch op {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if rhs == 0 {
			return 0, errors.New("zero division")
		}
		return lhs / rhs, nil
	}
	return 0, errors.New("invalid operation")
}

func GetPolishNotation(expr string) (string, error) {
	operations := &utils.Stack{Buffer: make([]interface{}, 0)}
	resultPolish := make([]string, 0)

	for _, char := range expr {
		if '0' <= char && char <= '9' {
			resultPolish = append(resultPolish, string(char))
		} else {
			switch char {
			case '(':
				operations.Push(char)
			case ')':
				for operations.GetSize() != 0 {
					popped, _ := operations.Pop()
					if popped != '(' {
						resultPolish = append(resultPolish, string(popped.(rune))) // type assertion to convert interface{} to rune
					}
				}
			case '+', '-', '*', '/':
				if operations.GetSize() == 0 {
					operations.Push(char)
				} else {
					top, _ := operations.Top()
					if getOperationPriority(char) > getOperationPriority(top.(rune)) {
						operations.Push(char)
					} else {
						popped, _ := operations.Pop()
						resultPolish = append(resultPolish, string(popped.(rune)))
						operations.Push(char)
					}
				}
			default:
				return "", errors.New("invalid rune")
			}
		}
	}

	for operations.GetSize() != 0 {
		popped, _ := operations.Pop()
		resultPolish = append(resultPolish, string(popped.(rune)))
	}

	return strings.Join(resultPolish[:], " "), nil
}

func Calculate(expr string) (float64, error) {
	if expr == "" {
		return 0, nil
	}

	return 0, nil
}
