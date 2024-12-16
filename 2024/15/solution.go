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

	wall    = "#"
	boxChar = "O"
)

var directionToOpposite = map[rune]rune{
	up:    down,
	down:  up,
	left:  right,
	right: left,
}

type position struct {
	x, y int
}

type robot struct {
	pos position
}

func (r *robot) move(dir rune, grid [][]string, boxes map[position]box) {
	newPos := r.pos
	switch dir {
	case up:
		newPos = position{r.pos.x - 1, r.pos.y}
	case down:
		newPos = position{r.pos.x + 1, r.pos.y}
	case left:
		newPos = position{r.pos.x, r.pos.y - 1}
	case right:
		newPos = position{r.pos.x, r.pos.y + 1}
	}

	if grid[newPos.x][newPos.y] == wall {
		// can't move
		return
	}

	if b, ok := boxes[newPos]; ok {
		// try to move the box
		canMove := b.checkMove(dir, grid, boxes)
		if !canMove {
			// can't move the box
			return
		}
		moved := b.move(dir, grid, boxes)
		if !moved {
			panic("should be able to move the box")
		}
	}
	r.pos = newPos
}

type box struct {
	left, right position
}

func (b box) move(dir rune, grid [][]string, boxes map[position]box) bool {
	newRight, newLeft := b.right, b.left

	delete(boxes, b.right)
	delete(boxes, b.left)

	switch dir {
	case up:
		newRight = position{b.right.x - 1, b.right.y}
		newLeft = position{b.left.x - 1, b.left.y}
	case down:
		newRight = position{b.right.x + 1, b.right.y}
		newLeft = position{b.left.x + 1, b.left.y}
	case left:
		newRight = position{b.right.x, b.right.y - 1}
		newLeft = position{b.left.x, b.left.y - 1}
	case right:
		newRight = position{b.right.x, b.right.y + 1}
		newLeft = position{b.left.x, b.left.y + 1}
	}

	if grid[newRight.x][newRight.y] == wall || grid[newLeft.x][newLeft.y] == wall {
		// can't move
		boxes[b.right] = b
		boxes[b.left] = b
		return false
	}

	if br, ok := boxes[newRight]; ok {
		// try to move the box
		brMoved := br.move(dir, grid, boxes)
		if !brMoved {
			boxes[b.right] = b
			boxes[b.left] = b
			return false
		}
	}

	if bl, ok := boxes[newLeft]; ok {
		// try to move the box
		blMoved := bl.move(dir, grid, boxes)
		if !blMoved {
			boxes[b.right] = b
			boxes[b.left] = b
			return false
		}
	}

	b.right = newRight
	b.left = newLeft
	boxes[b.right] = b
	boxes[b.left] = b

	return true
}

func (b box) checkMove(dir rune, grid [][]string, boxes map[position]box) bool {
	newRight, newLeft := b.right, b.left

	delete(boxes, b.right)
	delete(boxes, b.left)

	defer func() {
		boxes[b.right] = b
		boxes[b.left] = b
	}()

	switch dir {
	case up:
		newRight = position{b.right.x - 1, b.right.y}
		newLeft = position{b.left.x - 1, b.left.y}
	case down:
		newRight = position{b.right.x + 1, b.right.y}
		newLeft = position{b.left.x + 1, b.left.y}
	case left:
		newRight = position{b.right.x, b.right.y - 1}
		newLeft = position{b.left.x, b.left.y - 1}
	case right:
		newRight = position{b.right.x, b.right.y + 1}
		newLeft = position{b.left.x, b.left.y + 1}
	}

	if grid[newRight.x][newRight.y] == wall || grid[newLeft.x][newLeft.y] == wall {
		// cant move
		return false
	}

	// check if there is a box in the new position
	if br, ok := boxes[newRight]; ok {
		// see if the box can move
		canMove := br.checkMove(dir, grid, boxes)
		if !canMove {
			return false
		}
	}

	if bl, ok := boxes[newLeft]; ok {
		canMove := bl.checkMove(dir, grid, boxes)
		if !canMove {
			return false
		}
	}

	return true
}

func (b box) getGPS() int {
	// get the GRS coordinates of the box from the closet edge
	total := (100 * b.left.x)

	total += utils.Abs(0, b.left.y)

	return total
}

func display(grid [][]string) {
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}
}

func redraw(grid [][]string, robo robot, boxes map[position]box) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" {
				continue
			}
			grid[i][j] = "."
		}
	}

	grid[robo.pos.x][robo.pos.y] = "@"
	for _, b := range boxes {
		grid[b.left.x][b.left.y] = "["
		grid[b.right.x][b.right.y] = "]"
	}
	display(grid)
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

	if grid[newPos.x][newPos.y] == boxChar {
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
			if cell == boxChar {
				total += (100 * r) + c
			}
		}
	}

	fmt.Println(total)
}

func part2() {
	var moves string
	var robo robot
	boxes := map[position]box{}

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
		var newRow []string
		row := strings.Split(line, "")
		for _, cell := range row {
			switch cell {
			case "@":
				robo = robot{position{r, len(newRow)}}
				newRow = append(newRow, []string{"@", "."}...)
			case wall:
				newRow = append(newRow, []string{"#", "#"}...)
			case boxChar:
				b := box{position{r, len(newRow)}, position{r, len(newRow) + 1}}
				boxes[b.left] = b
				boxes[b.right] = b
				newRow = append(newRow, []string{"[", "]"}...)
			default:
				newRow = append(newRow, []string{".", "."}...)
			}
		}
		grid = append(grid, newRow)

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	for _, m := range moves {
		robo.move(m, grid, boxes)
	}

	total := 0
	onlyBoxes := map[box]struct{}{}
	for _, b := range boxes {
		onlyBoxes[b] = struct{}{}
	}

	for b := range onlyBoxes {
		total += b.getGPS()
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
