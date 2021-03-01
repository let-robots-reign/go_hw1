package part2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getOperationPriority(op string) int {
	switch op {
	case "(":
		return 0
	case "+", "-":
		return 1
	case "*", "/":
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

func getPolishNotation(expr string) string {

}

func calculate(expr string) (float64, error) {
	if expr == "" {
		return 0, nil
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading")
		return
	}

	polishNotation, err := getPolishNotation(strings.Trim(expression, "\n\t"))
	if err != nil {
		fmt.Println("Error while constructing polish notation")
		return
	}

	result, calcErr := calculate(polishNotation)
	if calcErr != nil {
		fmt.Println("Error while calculating")
		return
	}

	fmt.Println(result)
}
