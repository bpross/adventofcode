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

func checkAs(graph [][]string, i, j int) bool {
	if i-1 < 0 || j-1 < 0 || i+1 >= len(graph) || j+1 >= len(graph[i]) {
		return false
	}
	topLeft := graph[i-1][j-1]
	topRight := graph[i-1][j+1]
	bottomLeft := graph[i+1][j-1]
	bottomRight := graph[i+1][j+1]

	firstValid, secondValid := false, false
	if (topLeft == "S" && bottomRight == "M") || (topLeft == "M" && bottomRight == "S") {
		firstValid = true
	}

	if (topRight == "S" && bottomLeft == "M") || (topRight == "M" && bottomLeft == "S") {
		secondValid = true
	}

	return firstValid && secondValid
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

func part2() {
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
			if c == "A" {
				if graph[i][j] == "A" && checkAs(graph, i, j) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
