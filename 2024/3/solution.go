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

func part2() {
	total := 0
	multiply := true
	lineFunc := func(line string, _ int) error {
		mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
		matches := mul.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				fmt.Println("enabled", match)
				multiply = true
				continue
			} else if match[0] == "don't()" {
				fmt.Println("disabled", match)
				multiply = false
				continue
			}

			if multiply {
				fmt.Println("adding to total", match)
				i, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				j, err := strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}

				total += i * j
			} else {
				fmt.Println("skipped", match)
			}
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
	part2()
}
