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
	require.Equal(t, polish, "1 2 + 3 * 3 + 6 /")
}

func TestCalcAddition(t *testing.T) {
	polishInput := "100 42 +"
	result, err := calc.Calculate(polishInput)
	if err != nil {
		t.Fatal("addition failed")
	}
	require.Equal(t, result, 142.0)
}

func TestCalcSubtraction(t *testing.T) {
	polishInput := "100 42 -"
	result, err := calc.Calculate(polishInput)
	if err != nil {
		t.Fatal("subtraction failed")
	}
	require.Equal(t, result, 58.0)
}

func TestCalcMultiplication(t *testing.T) {
	polishInput := "100 42 *"
	result, err := calc.Calculate(polishInput)
	if err != nil {
		t.Fatal("multiplication failed")
	}
	require.Equal(t, result, 4200.0)
}

func TestCalcDivision(t *testing.T) {
	polishInput := "1000 50 /"
	result, err := calc.Calculate(polishInput)
	if err != nil {
		t.Fatal("division failed")
	}
	require.Equal(t, result, 20.0)
}

func TestCalcZeroDivision(t *testing.T) {
	polishInput := "1000 0 /"
	_, err := calc.Calculate(polishInput)
	if err == nil {
		t.Fatal("zero division didn't return error")
	}
}
