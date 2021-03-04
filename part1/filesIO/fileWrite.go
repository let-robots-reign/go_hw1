package filesIO

import "io"

func Write(writer io.Writer, outputLines []string) error {
	for _, line := range outputLines {
		_, writeErr := io.WriteString(writer, line+"\n")
		if writeErr != nil {
			return writeErr
		}
	}
	return nil
}
