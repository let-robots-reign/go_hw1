package main

import (
	"flag"
	"fmt"
	"go_hw1/part1/filesIO"
	"go_hw1/part1/uniq"
)

var flagCount = flag.Bool("c", false, "count each string's occurrences")
var flagDuplicate = flag.Bool("d", false, "return only repeating strings")
var flagUnique = flag.Bool("u", false, "return only not repeating strings")
var flagIgnoreFields = flag.Int("f", 0, "ignore first N fields in a string")
var flagIgnoreChars = flag.Int("s", 0, "ignore first N characters in a string")
var flagCaseInsensitive = flag.Bool("i", false, "enable case insensitivity")

func main() {
	flag.Parse()

	ignoreFields := false
	if *flagIgnoreFields != 0 {
		ignoreFields = true
	}

	ignoreChars := false
	if *flagIgnoreChars != 0 {
		ignoreChars = true
	}

	positionalArgs := uniq.Options{
		Count:            *flagCount,
		Duplicate:        *flagDuplicate,
		Unique:           *flagUnique,
		CaseInsensitive:  *flagCaseInsensitive,
		IgnoreFields:     ignoreFields,
		IgnoredFieldsNum: *flagIgnoreFields,
		IgnoreChars:      ignoreChars,
		IgnoredCharsNum:  *flagIgnoreChars,
	}

	infile := ""
	outfile := ""
	args := flag.Args() // оставшиеся аргументы
	if len(args) == 2 {
		infile = args[0]
		outfile = args[1]
	} else if len(args) == 1 {
		infile = args[0]
	}

	strings, readErr := filesIO.Read(infile)
	if readErr != nil {
		fmt.Println("Error while reading lines")
		return
	}
	result, findErr := uniq.FindUnique(strings, positionalArgs)
	if findErr != nil {
		fmt.Println("Error while finding unique")
		return
	}
	writeErr := filesIO.Write(outfile, result)
	if writeErr != nil {
		fmt.Println("Error while writing to file")
	}
}
