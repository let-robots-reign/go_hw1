package filesIO

import (
	"io"
	"os"
)

func Write(fileName string, outputLines []string) error {
	if fileName != "" {
		// если был передан файл, пишем в него
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		for _, line := range outputLines {
			_, err := file.WriteString(line + "\n")
			if err != nil {
				return err
			}
		}
	} else {
		for _, line := range outputLines {
			_, err := io.WriteString(os.Stdout, line+"\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
