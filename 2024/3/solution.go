package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

func part1() {
	total := 0
	lineFunc := func(line string, _ int) error {
		mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := mul.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			j, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			total += i * j
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}

func main() {
	part1()
}
