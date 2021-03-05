package uniq

import (
	"sort"
	"strconv"
	"strings"
)

type Options struct {
	Count            bool
	Duplicate        bool
	Unique           bool
	CaseInsensitive  bool
	IgnoredFieldsNum int
	IgnoredCharsNum  int
}

type StringInfo struct {
	original      string // оригинальная строка до ToLower() и пропуска слов/символов
	originalIndex int    // позиция оригинальной строки среди входных строк
	repeats       int    // количество повторений
}

func NewStringInfo(original string, index int) StringInfo {
	return StringInfo{
		original:      original,
		originalIndex: index,
		repeats:       1,
	}
}

func FindUnique(lines []string, options Options) (result []string, error error) {
	stringsOccurrences := make(map[string]*StringInfo)

	for index, str := range lines {
		original := str

		if options.CaseInsensitive {
			str = strings.ToLower(str)
		}

		if options.IgnoredFieldsNum > 0 {
			fields := strings.Fields(str)
			if options.IgnoredFieldsNum < len(fields) {
				str = strings.Join(fields[options.IgnoredFieldsNum:], " ")
			}
		}

		if options.IgnoredCharsNum > 0 {
			if options.IgnoredCharsNum < len(str) {
				str = str[options.IgnoredCharsNum:]
			}
		}

		if _, exists := stringsOccurrences[str]; exists {
			stringsOccurrences[str].repeats++
		} else {
			newKey := NewStringInfo(original, index)
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

	sort.Strings(result)

	return result, nil
}
