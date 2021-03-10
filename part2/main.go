package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/let-robots-reign/go_hw1/part2/calc"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var expression string
	if len(args) == 1 {
		expression = args[0]
	} else if len(args) == 0 {
		fmt.Println("Please enter expression:")

		reader := bufio.NewReader(os.Stdin)
		var err error
		expression, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading")
			return
		}
	} else {
		fmt.Println("Wrong number of command line arguments (1 allowed)")
		return
	}

	result, calcErr := calc.Calculate(strings.Trim(expression, "\n\r"))
	if calcErr != nil {
		fmt.Println("Error while calculating:", calcErr)
		return
	}

	fmt.Println("The result of your expression is:", result)
}
