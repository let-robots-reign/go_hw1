package main

import (
	"bufio"
	"fmt"
	"go_hw1/part2/calc"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading")
		return
	}

	// it's easier to calculate after converting to postfix form
	polishNotation, err := calc.GetPolishNotation(strings.Trim(expression, "\n\t"))
	if err != nil {
		fmt.Println("Error while constructing polish notation")
		return
	}
	fmt.Println(polishNotation)

	result, calcErr := calc.Calculate(polishNotation)
	if calcErr != nil {
		fmt.Println("Error while calculating")
		return
	}

	fmt.Println(result)
}
