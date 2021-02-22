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

func findUnique(strings []string, options Options) (result []string, error error) {

}
