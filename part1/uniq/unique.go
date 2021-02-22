package uniq

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

func FindUnique(strings []string, options Options) (result []string, error error) {
	stringsOccurrences := make(map[string]int)

	for _, str := range strings {
		stringsOccurrences[str]++
	}

	for key, _ := range stringsOccurrences {
		result = append(result, key)
	}

	return result, nil
}
