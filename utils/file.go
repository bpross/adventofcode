package utils

import (
	"bufio"
	"os"
)

type LineReader func(line string, lineNumber int) error

func ReadFile(filename string, lineFunc LineReader) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := lineFunc(scanner.Text(), i)
		if err != nil {
			return err
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
