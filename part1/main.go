package main

import (
	"flag"
	"fmt"
	"github.com/let-robots-reign/go_hw1/filesIO"
	"github.com/let-robots-reign/go_hw1/uniq"
	"os"
)

var flagCount = flag.Bool("c", false, "count each string's occurrences")
var flagDuplicate = flag.Bool("d", false, "return only repeating strings")
var flagUnique = flag.Bool("u", false, "return only not repeating strings")
var flagIgnoreFields = flag.Int("f", 0, "ignore first N fields in a string")
var flagIgnoreChars = flag.Int("s", 0, "ignore first N characters in a string")
var flagCaseInsensitive = flag.Bool("i", false, "enable case insensitivity")

func displayCorrectUsage() {
	fmt.Println("Correct usage of script: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}

func countTrue(b ...bool) int {
	count := 0
	for _, v := range b {
		if v {
			count++
		}
	}
	return count
}

func main() {
	flag.Parse()

	positionalArgs := uniq.Options{
		Count:            *flagCount,
		Duplicate:        *flagDuplicate,
		Unique:           *flagUnique,
		CaseInsensitive:  *flagCaseInsensitive,
		IgnoredFieldsNum: *flagIgnoreFields,
		IgnoredCharsNum:  *flagIgnoreChars,
	}

	if *flagIgnoreFields < 0 {
		fmt.Println("Negative -f flag not allowed")
		displayCorrectUsage()
		return
	}

	if *flagIgnoreChars < 0 {
		fmt.Println("Negative -s flag not allowed")
		displayCorrectUsage()
		return
	}

	// параллельно параметры -c, -d, -u не имеют смысла
	if countTrue(*flagCount, *flagDuplicate, *flagUnique) > 1 {
		fmt.Println("simultaneous usage of flags -c, -d, -u")
		displayCorrectUsage()
		return
	}

	infile := os.Stdin
	outfile := os.Stdout
	var errInfile, errOutfile error
	args := flag.Args() // оставшиеся аргументы
	if len(args) == 2 {
		infile, errInfile = os.Open(args[0])
		outfile, errOutfile = os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	} else if len(args) == 1 {
		infile, errInfile = os.Open(args[0])
	} else {
		fmt.Println("Wrong number of arguments")
		displayCorrectUsage()
		return
	}

	if errInfile != nil {
		fmt.Println("incorrect input file")
		return
	}
	if errOutfile != nil {
		fmt.Println("incorrect output file")
		return
	}

	defer infile.Close()
	defer outfile.Close()

	strings := filesIO.Read(infile)

	result, findErr := uniq.FindUnique(strings, positionalArgs)
	if findErr != nil {
		fmt.Println("Error while finding unique")
		return
	}
	writeErr := filesIO.Write(outfile, result)
	if writeErr != nil {
		fmt.Println("Error while writing to file", writeErr)
	}
}
