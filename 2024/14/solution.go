package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

var robotRe = regexp.MustCompile(`p=(\d{1,3}),(\d{1,3}) v=(-?)(\d{1,3}),(-?)(\d{1,3})`)

type robot struct {
	cv, rv int
	c, r   int
}

func (r *robot) move(width, height int) {
	newC := r.c + r.cv
	newR := r.r + r.rv
	if newC < 0 || newC >= width {
		// wrap around
		newC = (newC + width) % width
	}

	if newR < 0 || newR >= height {
		// wrap around
		newR = (newR + height) % height
	}

	r.c = newC
	r.r = newR
}

func (r *robot) getQuad(width, height int) int {
	badRow := height / 2
	badCol := width / 2
	if r.r == badRow || r.c == badCol {
		return 0
	}

	if r.c < width/2 {
		if r.r < height/2 {
			return 1
		}
		return 3
	}
	if r.r < height/2 {
		return 2
	}
	return 4
}

func part1() {
	robots := make([]*robot, 0)
	lineFunc := func(line string, _ int) error {
		matches := robotRe.FindStringSubmatch(line)
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		xNeg := matches[3] == "-"
		xv, err := strconv.Atoi(matches[4])
		if err != nil {
			panic(err)
		}
		if xNeg {
			xv = -xv
		}
		yNeg := matches[5] == "-"
		yv, err := strconv.Atoi(matches[6])
		if err != nil {
			panic(err)
		}
		if yNeg {
			yv = -yv
		}
		robots = append(robots, &robot{xv, yv, x, y})
		return nil
	}
	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	width, height := 101, 103
	moves := 100
	for i := 0; i < moves; i++ {
		for _, r := range robots {
			r.move(width, height)
		}
	}

	total := 1
	quads := make(map[int]int)

	for _, r := range robots {
		quads[r.getQuad(width, height)]++
	}

	delete(quads, 0)
	for _, v := range quads {
		total *= v
	}
	fmt.Println(total)
}

func main() {
	part1()
}
