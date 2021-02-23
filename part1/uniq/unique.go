package uniq

import (
	"errors"
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

type StringInfo struct {
	original string // оригинальная строка до ToLower() и пропуска слов/символов
	repeats  int    // количество повторений
}

func NewStringInfo(original string) StringInfo {
	return StringInfo{
		original: original,
		repeats:  1,
	}
}

func FindUnique(lines []string, options Options) (result []string, error error) {
	stringsOccurrences := make(map[string]*StringInfo)

	for _, str := range lines {
		original := str

		if options.CaseInsensitive {
			str = strings.ToLower(str)
		}

		if options.IgnoreFields {
			fields := strings.Fields(str)
			if options.IgnoredFieldsNum <= len(fields) {
				str = strings.Join(fields[options.IgnoredFieldsNum:], " ")
			} else if str != "" {
				return nil, errors.New("incorrect -f argument")
			}
		}

		if options.IgnoreChars {
			if options.IgnoredCharsNum <= len(str) {
				str = str[options.IgnoredCharsNum:]
			} else if str != "" {
				return nil, errors.New("incorrect -c argument")
			}
		}

		if _, exists := stringsOccurrences[str]; exists {
			stringsOccurrences[str].repeats++
		} else {
			newKey := NewStringInfo(original)
			stringsOccurrences[str] = &newKey
		}
	}

	for _, value := range stringsOccurrences {
		switch {
		case options.Duplicate:
			if value.repeats > 1 {
				result = append(result, value.original)
			}
		case options.Unique:
			if value.repeats == 1 {
				result = append(result, value.original)
			}
		case options.Count:
			result = append(result, strconv.Itoa(value.repeats)+" "+value.original)
		default:
			result = append(result, value.original)
		}
	}

	return result, nil
}
