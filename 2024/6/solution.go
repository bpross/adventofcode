package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	r, c int
}

type cyclePosition struct {
	position
	direction rune
}

var positionToRight = map[rune]rune{
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func positionValid(grid [][]string, pos position) bool {
	if pos.r < 0 || pos.r >= len(grid) {
		return false
	}
	if pos.c < 0 || pos.c >= len(grid[pos.r]) {
		return false
	}
	return true
}

func positionObstacle(grid [][]string, pos position) bool {
	if !positionValid(grid, pos) {
		return false
	}

	return grid[pos.r][pos.c] == "#"
}

func getMove(grid [][]string, pos position, direction rune) (position, rune) {
	up := position{pos.r - 1, pos.c}
	down := position{pos.r + 1, pos.c}
	left := position{pos.r, pos.c - 1}
	right := position{pos.r, pos.c + 1}

	switch direction {
	case '^':
		// up
		pos = up
		direction = '^'
		if positionObstacle(grid, up) {
			pos = right
			direction = '>'
		} else {
			break
		}
		if positionObstacle(grid, right) {
			pos = down
			direction = 'v'
		}
	case 'v':
		// down
		pos = down
		direction = 'v'
		if positionObstacle(grid, down) {
			pos = left
			direction = '<'
		} else {
			break
		}
		if positionObstacle(grid, left) {
			pos = up
			direction = '^'
		}
	case '<':
		// left
		pos = left
		direction = '<'
		if positionObstacle(grid, left) {
			pos = up
			direction = '^'
		} else {
			break
		}
		if positionObstacle(grid, up) {
			pos = right
			direction = '>'
		}
	case '>':
		pos = right
		direction = '>'
		if positionObstacle(grid, right) {
			pos = down
			direction = 'v'
		} else {
			break
		}
		if positionObstacle(grid, down) {
			pos = left
			direction = '<'
		}
	}

	return pos, direction
}

func checkForCycle(grid [][]string, loc, newObj position, direction rune) bool {
	path := map[cyclePosition]struct{}{}
	path[cyclePosition{loc, direction}] = struct{}{}

	if grid[newObj.r][newObj.c] == "#" {
		return false
	}
	grid[newObj.r][newObj.c] = "#"

	defer func() {
		grid[newObj.r][newObj.c] = "."
	}()

	for positionValid(grid, loc) {
		loc, direction = getMove(grid, loc, direction)
		if _, ok := path[cyclePosition{loc, direction}]; ok {
			return true
		}
		path[cyclePosition{loc, direction}] = struct{}{}
	}

	return false
}

func part1() {
	positions := map[position]struct{}{}
	grid := [][]string{}
	loc := position{-1, -1}
	direction := '^'

	lineFunc := func(line string, n int) error {
		row := []string{}
		for i, c := range line {
			row = append(row, string(c))
			if c == direction {
				loc = position{n, i}
			}
		}
		grid = append(grid, row)

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	for positionValid(grid, loc) {
		positions[loc] = struct{}{}
		loc, direction = getMove(grid, loc, direction)
	}

	fmt.Println(len(positions))
}

func part2() {
	grid := [][]string{}
	loc := position{-1, -1}
	direction := '^'

	lineFunc := func(line string, n int) error {
		row := []string{}
		for i, c := range line {
			row = append(row, string(c))
			if c == direction {
				loc = position{n, i}
			}
		}
		grid = append(grid, row)

		return nil
	}
	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	cycles := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" {
				continue
			}
			if loc.r == i && loc.c == j {
				continue
			}

			if checkForCycle(grid, loc, position{i, j}, '^') {
				cycles++
			}
		}
	}

	fmt.Println(cycles)
}

func main() {
	part1()
	part2()
}
