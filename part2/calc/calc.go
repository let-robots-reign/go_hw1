package calc

import (
	"errors"
	"github.com/let-robots-reign/go_hw1/part2/utils"
	"strconv"
	"strings"
)

func contains(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func isOperator(r rune) bool {
	return contains([]rune{'+', '-', '*', '/'}, r)
}

func validateExpression(expr string) error {
	bracesCount := 0
	previousWasOperator := false

	for _, char := range expr {
		if !('0' <= char && char <= '9') && !contains([]rune{'+', '-', '*', '/', '(', ')'}, char) {
			return errors.New("invalid chars")
		}

		if isOperator(char) && previousWasOperator {
			return errors.New("several operators in succession")
		}

		if isOperator(char) {
			previousWasOperator = true
		} else {
			previousWasOperator = false
		}

		if char == '(' {
			bracesCount++
		} else if char == ')' {
			bracesCount--
		}

		if bracesCount < 0 {
			return errors.New("closing bracket doesn't match any opening bracket")
		}
	}

	if bracesCount != 0 {
		return errors.New("extra or not enough braces")
	}

	return nil
}

func getOperationPriority(op rune) int {
	switch op {
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
	var postfix string
	previousWasOperator := false

	for _, char := range expr {
		if '0' <= char && char <= '9' {
			if previousWasOperator == true {
				postfix += " "
			}
			postfix += string(char)
			previousWasOperator = false
			continue
		}

		if char == '(' {
			operations.Push(char)
		} else if char == ')' {
			top, _ := operations.Top()
			for operations.GetSize() != 0 && top != '(' {
				popped, _ := operations.Pop()
				postfix += " " + string(popped.(rune))

				top, _ = operations.Top()
			}

			top, _ = operations.Top()
			if top == '(' {
				_, _ = operations.Pop()
			}
		} else {
			previousWasOperator = true
			top, _ := operations.Top()
			for operations.GetSize() != 0 && getOperationPriority(char) <= getOperationPriority(top.(rune)) {
				popped, _ := operations.Pop()
				postfix += " " + string(popped.(rune))

				top, _ = operations.Top()
			}
			operations.Push(char)
		}
	}

	for operations.GetSize() != 0 {
		popped, _ := operations.Pop()
		postfix += " " + string(popped.(rune))
	}

	return postfix, nil
}

func Calculate(expr string) (float64, error) {
	if expr == "" {
		return 0, nil
	}

	err := validateExpression(expr)
	if err != nil {
		return 0, err
	}

	// it's easier to calculate after converting to postfix form
	polishNotation, err := GetPolishNotation(expr)
	if err != nil {
		return 0, errors.New("error while constructing polish notation")
	}

	calcStack := &utils.Stack{Buffer: make([]interface{}, 0)}
	tokens := strings.Split(polishNotation, " ")

	for _, token := range tokens {
		number, parseError := strconv.ParseFloat(token, 64)
		if parseError == nil {
			// token is a number
			calcStack.Push(number)
		} else {
			// token is an operator
			rhs, lhsPopErr := calcStack.Pop()
			lhs, rhsPopErr := calcStack.Pop()
			if lhsPopErr != nil || rhsPopErr != nil {
				return 0, errors.New("invalid expression")
			}

			result, applyErr := applyOperation(lhs.(float64), rhs.(float64), token)
			if applyErr != nil {
				return 0, applyErr
			}

			calcStack.Push(result)
		}
	}

	result, _ := calcStack.Pop()
	return result.(float64), nil
}
