package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	Count            bool
	Duplicate        bool
	Unique           bool
	CaseInsensitive  bool
	IgnoreFields     bool
	IgnoredFieldsNum int
	IgnoreChars      bool
	IgnoredCharsNum  int
}

func FindUnique(lines []string, options Options) (result []string, error error) {
	stringsOccurrences := make(map[string]int)

	// TODO: fix -i, add -f and -c

	for _, str := range lines {
		if options.CaseInsensitive {
			str = strings.ToLower(str)
		}

		stringsOccurrences[str]++
	}

	if options.Duplicate {
		for key, value := range stringsOccurrences {
			if value > 1 {
				result = append(result, key)
			}
		}
	} else if options.Unique {
		for key, value := range stringsOccurrences {
			if value == 1 {
				result = append(result, key)
			}
		}
	} else if options.Count {
		for key, value := range stringsOccurrences {
			result = append(result, strconv.Itoa(value)+" "+key)
		}
	} else {
		for key, _ := range stringsOccurrences {
			result = append(result, key)
		}
	}

	return result, nil
}
