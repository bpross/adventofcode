package utils

import (
	"bufio"
	"os"
)

type (
	LineReader      func(line string, lineNumber int) error
	MultiLineReader func(lines []string, lineNumbers []int) error
)

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

func ReadFileInChunks(filename string, chunkSize int, lineFunc MultiLineReader) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	chunk := make([]string, chunkSize)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			continue
		}
		chunk[i] = t
		if i == chunkSize-1 {
			err := lineFunc(chunk, make([]int, chunkSize))
			if err != nil {
				return err
			}
			chunk = make([]string, chunkSize)
			i = 0
		} else {
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
