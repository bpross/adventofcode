package utils

import (
	"bufio"
	"os"
)

type LineReader func(line string) error

func ReadFile(filename string, lineFunc LineReader) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err := lineFunc(scanner.Text())
		if err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
