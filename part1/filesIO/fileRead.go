package filesIO

import (
	"bufio"
	"io"
)

func scanLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Read(reader io.Reader) []string {
	return scanLines(bufio.NewScanner(reader))
}
