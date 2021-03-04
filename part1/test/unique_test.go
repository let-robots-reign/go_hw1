package test

import (
	"errors"
	"github.com/stretchr/testify/require"
	"go_hw1/part1/uniq"
	"testing"
)

func TestWithoutOptions(t *testing.T) {
	testInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        false,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		" ",
		"I love music of Kartik.",
		"I love music.",
		"Thanks.",
	}
	require.Equal(t, result, answer, "launch with no positional arguments")
}

func TestDuplicates(t *testing.T) {
	testInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        true,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		"I love music of Kartik.",
		"I love music.",
	}
	require.Equal(t, result, answer, "launch with -d flag")
}

func TestUnique(t *testing.T) {
	testInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        false,
		Unique:           true,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		" ",
		"Thanks.",
	}
	require.Equal(t, result, answer, "launch with -u flag")
}

func TestCount(t *testing.T) {
	testInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            true,
		Duplicate:        false,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		"1  ",
		"1 Thanks.",
		"2 I love music of Kartik.",
		"3 I love music.",
	}
	require.Equal(t, result, answer, "launch with -c flag")
}

func TestCaseInsensitive(t *testing.T) {
	testInput := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		" ",
		"I love MuSIc of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        false,
		Unique:           false,
		CaseInsensitive:  true,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		" ",
		"I LOVE MUSIC.",
		"I love MuSIc of Kartik.",
		"Thanks.",
	}
	require.Equal(t, result, answer, "launch with -i flag")
}

func TestIgnoreFields(t *testing.T) {
	testInput := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        false,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 1,
		IgnoredCharsNum:  0,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		"I love music of Kartik.",
		"Thanks.",
		"We love music.",
	}
	require.Equal(t, result, answer, "launch with -f 1 flag")
}

func TestIgnoreChars(t *testing.T) {
	testInput := []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        false,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 0,
		IgnoredCharsNum:  1,
	}

	result, _ := uniq.FindUnique(testInput, options)

	answer := []string{
		"",
		"I love music of Kartik.",
		"I love music.",
		"Thanks.",
		"We love music of Kartik.",
	}
	require.Equal(t, result, answer, "launch with -s 1 flag")
}

func TestErrors(t *testing.T) {
	testInput := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	options := uniq.Options{
		Count:            false,
		Duplicate:        true,
		Unique:           false,
		CaseInsensitive:  false,
		IgnoredFieldsNum: 5,
		IgnoredCharsNum:  0,
	}

	_, err := uniq.FindUnique(testInput, options)
	if err != nil {
		incorrectArgumentError := errors.New("incorrect -f argument")
		require.Equal(t, err, incorrectArgumentError)
	}
}
