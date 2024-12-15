package main

import (
	"fmt"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'

	wall = "#"
	box  = "O"
)

type position struct {
	x, y int
}

func display(grid [][]string) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func move(p position, dir rune, grid [][]string) (bool, position) {
	char := grid[p.x][p.y]
	newPos := p
	switch dir {
	case up:
		newPos = position{p.x - 1, p.y}
	case down:
		newPos = position{p.x + 1, p.y}
	case left:
		newPos = position{p.x, p.y - 1}
	case right:
		newPos = position{p.x, p.y + 1}
	}

	if grid[newPos.x][newPos.y] == wall {
		// can't move
		return false, p
	}

	if grid[newPos.x][newPos.y] == box {
		// try to move the box
		moved, _ := move(newPos, dir, grid)
		if !moved {
			// can't move the box
			return false, p
		}
	}

	grid[newPos.x][newPos.y] = char
	grid[p.x][p.y] = "."
	return true, newPos
}

func part1() {
	var start position
	var moves string
	grid := [][]string{}
	b := false

	lineFunc := func(line string, r int) error {
		if line == "" {
			b = true
			return nil
		}

		if b {
			line = strings.TrimSpace(line)
			moves += line
			return nil
		}
		row := strings.Split(line, "")
		for c, cell := range row {
			if cell == "@" {
				start = position{r, c}
			}
		}
		grid = append(grid, row)

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	for _, m := range moves {
		_, start = move(start, m, grid)
	}

	total := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == box {
				total += (100 * r) + c
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
}
