package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

type Direction int

const (
	up Direction = iota
	down
	left
	right
	upLeft
	upRight
	downLeft
	downRight
)

var Directions = []Direction{up, down, left, right, upLeft, upRight, downLeft, downRight}

type steps struct {
	i, j int
}

func dfs(graph [][]string, i, j int, path string, s []steps, direction Direction) int {
	if len(s) > 4 {
		return 0
	}

	if path == "XMAS" {
		fmt.Println(i, j, path, s, direction)
		return 1
	}

	total := 0

	switch direction {
	case up:
		if i-1 < 0 {
			return 0
		}
		total += dfs(graph, i-1, j, path+graph[i-1][j], append(s, steps{i - 1, j}), up)
	case down:
		if i+1 >= len(graph) {
			return 0
		}
		total += dfs(graph, i+1, j, path+graph[i+1][j], append(s, steps{i + 1, j}), down)
	case left:
		if j-1 < 0 {
			return 0
		}
		total += dfs(graph, i, j-1, path+graph[i][j-1], append(s, steps{i, j - 1}), left)
	case right:
		if j+1 >= len(graph[i]) {
			return 0
		}
		total += dfs(graph, i, j+1, path+graph[i][j+1], append(s, steps{i, j + 1}), right)
	case upLeft:
		if i-1 < 0 || j-1 < 0 {
			return 0
		}
		total += dfs(graph, i-1, j-1, path+graph[i-1][j-1], append(s, steps{i - 1, j - 1}), upLeft)
	case upRight:
		if i-1 < 0 || j+1 >= len(graph[i]) {
			return 0
		}
		total += dfs(graph, i-1, j+1, path+graph[i-1][j+1], append(s, steps{i - 1, j + 1}), upRight)
	case downLeft:
		if i+1 >= len(graph) || j-1 < 0 {
			return 0
		}
		total += dfs(graph, i+1, j-1, path+graph[i+1][j-1], append(s, steps{i + 1, j - 1}), downLeft)
	case downRight:
		if i+1 >= len(graph) || j+1 >= len(graph[i]) {
			return 0
		}
		total += dfs(graph, i+1, j+1, path+graph[i+1][j+1], append(s, steps{i + 1, j + 1}), downRight)
	default:
		return 0
	}

	return total
}

func part1() {
	graph := [][]string{}

	lineFunc := func(line string, _ int) error {
		row := make([]string, len(line))
		for i, c := range line {
			row[i] = string(c)
		}
		graph = append(graph, row)
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0

	for i, row := range graph {
		for j, c := range row {
			if c == "X" {
				for _, d := range Directions {
					total += dfs(graph, i, j, graph[i][j], []steps{{i, j}}, d)
				}
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
}
