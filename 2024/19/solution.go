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

func main() {
	part1()
}
