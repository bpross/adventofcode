package main

import (
	"fmt"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

func canBeMade(patterns []string, design string) bool {
	var dfs func(path string) bool
	dfs = func(path string) bool {
		if path == design {
			return true
		}
		if !strings.HasPrefix(design, path) {
			return false
		}

		for _, p := range patterns {
			if dfs(path + p) {
				return true
			}
		}
		return false
	}

	return dfs("")
}

func waysCanBeMade(patterns []string, design string) int {
	var dfs func(path string) int
	cache := map[string]int{}

	dfs = func(path string) int {
		if path == design {
			return 1
		}
		if !strings.HasPrefix(design, path) {
			return 0
		}
		if v, ok := cache[path]; ok {
			return v
		}

		total := 0
		for _, p := range patterns {
			newPath := path + p
			t := dfs(newPath)
			cache[newPath] = t
			total += t
		}
		return total
	}

	return dfs("")
}

func part1() {
	patterns := []string{}
	designs := []string{}
	b := false

	lineFunc := func(line string, r int) error {
		if line == "" {
			b = true
			return nil
		}

		if b {
			line = strings.TrimSpace(line)
			designs = append(designs, line)
			return nil
		}

		row := strings.Split(line, ",")
		for _, p := range row {
			ps := strings.TrimSpace(p)
			patterns = append(patterns, ps)
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, d := range designs {
		if canBeMade(patterns, d) {
			total++
		}
	}
	fmt.Println(total)
}

func part2() {
	patterns := []string{}
	designs := []string{}
	b := false

	lineFunc := func(line string, r int) error {
		if line == "" {
			b = true
			return nil
		}

		if b {
			line = strings.TrimSpace(line)
			designs = append(designs, line)
			return nil
		}

		row := strings.Split(line, ",")
		for _, p := range row {
			ps := strings.TrimSpace(p)
			patterns = append(patterns, ps)
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, d := range designs {
		total += waysCanBeMade(patterns, d)
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
