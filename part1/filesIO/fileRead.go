package filesIO

import (
	"bufio"
	"os"
)

func scanLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func read(fileName string) ([]string, error) {
	var inputLines []string

	if fileName != "" {
		// если был передан файл, открываем его
		file, err := os.Open(fileName)
		if err != nil {
			return nil, err
		}
		defer file.Close() // закрываем файл перед выходом из функции

		scanner := bufio.NewScanner(file)
		inputLines = scanLines(scanner)
	} else {
		// если файла нет, то читаем из stdin
		scanner := bufio.NewScanner(os.Stdin)
		inputLines = scanLines(scanner)
	}

	return inputLines, nil
}
