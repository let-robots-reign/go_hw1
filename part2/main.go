package main

import (
	"bufio"
	"fmt"
	"github.com/let-robots-reign/go_hw1/part2/calc"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please enter expression:")

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

	fmt.Println("The result of your expression is:", result)
}
