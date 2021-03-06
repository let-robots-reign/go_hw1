package test

import (
	"github.com/let-robots-reign/go_hw1/part1/uniq"
	"github.com/stretchr/testify/require"
	"testing"
)

var uniqTests = []struct {
	input           []string
	options         uniq.Options
	answer          []string
	testDescription string
}{
	{
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			"I love music.",
			" ",
			"I love music of Kartik.",
			"Thanks.",
		},

		"launch with no positional arguments",
	},

	{
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            true,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			"3 I love music.",
			"1  ",
			"2 I love music of Kartik.",
			"1 Thanks.",
		},

		"launch with -c flag",
	},

	{
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        true,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			"I love music.",
			"I love music of Kartik.",
		},

		"launch with -d flag",
	},

	{
		[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			" ",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        false,
			Unique:           true,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			" ",
			"Thanks.",
		},

		"launch with -u flag",
	},

	{
		[]string{
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			" ",
			"I love MuSIC of Kartik.",
			"I love music of kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  true,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			"I LOVE MUSIC.",
			" ",
			"I love MuSIC of Kartik.",
			"Thanks.",
		},

		"launch with -i flag",
	},

	{
		[]string{
			"We love music.",
			"I love music.",
			"They love music.",
			" ",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 1,
			IgnoredCharsNum:  0,
		},

		[]string{
			"We love music.",
			" ",
			"I love music of Kartik.",
			"Thanks.",
		},

		"launch with -f 1 flag",
	},

	{
		[]string{
			"I love music.",
			"A love music.",
			"C love music.",
			" ",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            false,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  1,
		},

		[]string{
			"I love music.",
			" ",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},

		"launch with -s 1 flag",
	},

	{
		[]string{
			"I LOVE music.",
			"I love music.",
			"I love MUSic.",
			"THANKS.",
			" ",
			"I love music of Kartik.",
			"I loVe music of KARTIK.",
			"Thanks.",
		},

		uniq.Options{
			Count:            true,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  true,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  0,
		},

		[]string{
			"3 I LOVE music.",
			"2 THANKS.",
			"1  ",
			"2 I love music of Kartik.",
		},

		"launch with -c and -i flags",
	},

	{
		[]string{
			"I LOVE music.",
			"I love music.",
			"I love MUSic.",
			"We LoVe mUsIc.",
			"THANKS.",
			" ",
			"I love music of Kartik.",
			"I loVe music of KARTIK.",
			"Thanks.",
		},

		uniq.Options{
			Count:            true,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  true,
			IgnoredFieldsNum: 1,
			IgnoredCharsNum:  0,
		},

		[]string{
			"4 I LOVE music.",
			"2 THANKS.",
			"1  ",
			"2 I love music of Kartik.",
		},

		"launch with -c, -f 1 and -i flags",
	},

	{
		[]string{
			"I love music.",
			"A love music.",
			"C love music.",
			" ",
			"I love music of Kartik.",
			"O love music of Kartik.",
			"Thanks.",
		},

		uniq.Options{
			Count:            true,
			Duplicate:        false,
			Unique:           false,
			CaseInsensitive:  false,
			IgnoredFieldsNum: 0,
			IgnoredCharsNum:  1,
		},

		[]string{
			"3 I love music.",
			"1  ",
			"2 I love music of Kartik.",
			"1 Thanks.",
		},

		"launch with -s 1 and -c flags",
	},
}

func TestUniq(t *testing.T) {
	for _, testCase := range uniqTests {
		result, _ := uniq.FindUnique(testCase.input, testCase.options)
		require.Equal(t, testCase.answer, result, testCase.testDescription)
	}
}
