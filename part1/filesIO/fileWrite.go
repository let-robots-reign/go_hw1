package filesIO

import (
	"io"
	"os"
)

func writeLines(file io.Writer, lines []string) error {
	for _, line := range lines {
		_, err := io.WriteString(file, line+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func write(fileName string, outputLines []string) error {
	var file *os.File

	if fileName != "" {
		// если был передан файл, пишем в него
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
	} else {
		file = os.Stdout
	}

	return writeLines(file, outputLines)
}
