package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	r, c int
}

func positionValid(grid []string, pos position) bool {
	if pos.r < 0 || pos.r >= len(grid) {
		return false
	}
	if pos.c < 0 || pos.c >= len(grid[pos.r]) {
		return false
	}
	return true
}

func positionObstacle(grid []string, pos position) bool {
	if !positionValid(grid, pos) {
		return false
	}

	return grid[pos.r][pos.c] == '#'
}

func getMove(grid []string, pos position, direction rune) (position, rune) {
	switch direction {
	case '^':
		// up
		// check for obstacle
		pos = position{pos.r - 1, pos.c}
		if positionObstacle(grid, pos) {
			// turn right
			pos = position{pos.r + 1, pos.c + 1}
			direction = '>'
		}
	case 'v':
		// down
		pos = position{pos.r + 1, pos.c}
		if positionObstacle(grid, pos) {
			// turn right
			pos = position{pos.r - 1, pos.c - 1}
			direction = '<'
		}
	case '<':
		// left
		pos = position{pos.r, pos.c - 1}
		if positionObstacle(grid, pos) {
			// turn right
			pos = position{pos.r - 1, pos.c + 1}
			direction = '^'
		}
	case '>':
		pos = position{pos.r, pos.c + 1}
		if positionObstacle(grid, pos) {
			// turn right
			pos = position{pos.r + 1, pos.c - 1}
			direction = 'v'
		}

	}

	return pos, direction
}

func part1() {
	positions := map[position]struct{}{}
	grid := []string{}
	loc := position{-1, -1}
	direction := '^'

	lineFunc := func(line string, n int) error {
		grid = append(grid, line)

		for i, c := range line {
			if c == direction {
				loc = position{n, i}
			}
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	for positionValid(grid, loc) {
		fmt.Println(loc, direction)
		positions[loc] = struct{}{}
		loc, direction = getMove(grid, loc, direction)
	}

	fmt.Println(len(positions))
}

func main() {
	part1()
}
