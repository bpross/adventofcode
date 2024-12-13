package main

import (
	"fmt"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	r, c  int
	plant string
}

func (p position) getNeighbors(g [][]string) []position {
	res := []position{}
	plant := g[p.r][p.c]

	pos := []position{
		{p.r - 1, p.c, plant},
		{p.r + 1, p.c, plant},
		{p.r, p.c - 1, plant},
		{p.r, p.c + 1, plant},
	}

	for _, n := range pos {
		if n.r >= 0 && n.r < len(g) && n.c >= 0 && n.c < len(g[0]) && g[n.r][n.c] == plant {
			res = append(res, n)
		}
	}
	return res
}

func (p position) perimeter(g [][]string) int {
	return 4 - len(p.getNeighbors(g))
}

type plot struct {
	id        int
	plant     string
	positions []position
}

func (p plot) getArea() int {
	return len(p.positions)
}

func (p plot) getPerimeter(g [][]string) int {
	perimeter := 0
	for _, pos := range p.positions {
		perimeter += pos.perimeter(g)
	}
	return perimeter
}

func part1() {
	grid := [][]string{}
	lineFunc := func(line string, _ int) error {
		row := strings.Split(line, "")
		grid = append(grid, row)
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	positionToPlot := map[position]plot{}
	plotID := 0
	plots := []*plot{}

	for r, row := range grid {
		for c, cell := range row {
			p := position{r, c, cell}
			q := []position{p}

			var pt plot
			if _, ok := positionToPlot[p]; !ok {
				pt = plot{plotID, cell, []position{}}
				plots = append(plots, &pt)
				plotID++
			} else {
				continue
			}

			// BFS our way through the grid
			// to create the plots
			var item position
			visited := map[position]struct{}{}
			for len(q) > 0 {
				item, q = q[0], q[1:]
				if _, ok := visited[item]; ok {
					continue
				}
				if _, ok := positionToPlot[item]; ok {
					// already assigned to a plot
					continue
				}
				visited[item] = struct{}{}
				positionToPlot[item] = pt
				pt.positions = append(pt.positions, item)
				q = append(q, item.getNeighbors(grid)...)
			}
		}
	}

	total := 0
	for _, p := range plots {
		total += p.getArea() * p.getPerimeter(grid)
	}
	fmt.Println(total)
}

func main() {
	part1()
}
