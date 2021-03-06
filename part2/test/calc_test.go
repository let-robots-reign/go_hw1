package test

import (
	"github.com/let-robots-reign/go_hw1/part2/calc"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPolishNotation(t *testing.T) {
	expr := "((1+2)*3+3)/6"
	polish, err := calc.GetPolishNotation(expr)
	if err != nil {
		t.Fatal("conversion to polish notation failed")
	}
	require.Equal(t, "1 2 + 3 * 3 + 6 /", polish)
}

var calcTest = []struct {
	input           string
	answer          float64
	testDescription string
}{
	{
		"100+42",
		142.0,
		"test addition",
	},

	{
		"100-42",
		58.0,
		"test subtraction",
	},

	{
		"100*42",
		4200.0,
		"test multiplication",
	},

	{
		"1000/50",
		20.0,
		"test division",
	},

	{
		"((1+2)*3+3)/6",
		2.0,
		"test braces #1",
	},

	{
		"(12*(5+7)/12-12)*100",
		0.0,
		"test braces #2",
	},

	{
		"(0-100)+1*5-20/4*10",
		-145.0,
		"test braces #3",
	},
}

func TestCalc(t *testing.T) {
	for _, testCase := range calcTest {
		result, _ := calc.Calculate(testCase.input)
		require.Equal(t, testCase.answer, result, testCase.testDescription)
	}
}

func TestCalcZeroDivision(t *testing.T) {
	expr := "1000/0"
	_, err := calc.Calculate(expr)
	if err == nil {
		t.Fatal("zero division didn't return error")
	}
}
