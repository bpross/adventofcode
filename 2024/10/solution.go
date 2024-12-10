package main

import (
	"fmt"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	r, c int
}

func dfsFinalPosition(grid [][]int, current, r, c int) []position {
	fmt.Printf("r: %d, c: %d\n", r, c)
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
		fmt.Println("out of bounds")
		return []position{}
	}

	diff := grid[r][c] - current
	if diff != 1 {
		return []position{}
	}

	if grid[r][c] == 9 {
		fmt.Println("found 9")
		return []position{{r, c}}
	}

	total := []position{}
	total = append(total, dfsFinalPosition(grid, grid[r][c], r-1, c)...)
	total = append(total, dfsFinalPosition(grid, grid[r][c], r+1, c)...)
	total = append(total, dfsFinalPosition(grid, grid[r][c], r, c+1)...)
	total = append(total, dfsFinalPosition(grid, grid[r][c], r, c-1)...)

	return total
}

func dfs(grid [][]int, current, r, c int) int {
	fmt.Printf("r: %d, c: %d\n", r, c)
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
		fmt.Println("out of bounds")
		return 0
	}

	diff := grid[r][c] - current
	if diff != 1 {
		return 0
	}

	if grid[r][c] == 9 {
		fmt.Println("found 9")
		return 1
	}

	total := 0
	total += dfs(grid, grid[r][c], r-1, c)
	total += dfs(grid, grid[r][c], r+1, c)
	total += dfs(grid, grid[r][c], r, c+1)
	total += dfs(grid, grid[r][c], r, c-1)

	return total
}

func part1() {
	trailheads := []position{}
	grid := [][]int{}

	lineFunc := func(line string, r int) error {
		row := []int{}

		for c, char := range line {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row = append(row, i)
			if i == 0 {
				trailheads = append(trailheads, position{r, c})
			}
		}

		grid = append(grid, row)

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0

	for _, th := range trailheads {
		reachableNines := []position{}
		reachableNines = append(reachableNines, dfsFinalPosition(grid, 0, th.r-1, th.c)...)
		reachableNines = append(reachableNines, dfsFinalPosition(grid, 0, th.r+1, th.c)...)
		reachableNines = append(reachableNines, dfsFinalPosition(grid, 0, th.r, th.c-1)...)
		reachableNines = append(reachableNines, dfsFinalPosition(grid, 0, th.r, th.c+1)...)

		deduped := map[position]struct{}{}
		for _, p := range reachableNines {
			deduped[p] = struct{}{}
		}
		total += len(deduped)
	}

	fmt.Println(total)
}

func part2() {
	trailheads := []position{}
	grid := [][]int{}

	lineFunc := func(line string, r int) error {
		row := []int{}

		for c, char := range line {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row = append(row, i)
			if i == 0 {
				trailheads = append(trailheads, position{r, c})
			}
		}

		grid = append(grid, row)

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0

	for _, th := range trailheads {
		total += dfs(grid, 0, th.r-1, th.c)
		total += dfs(grid, 0, th.r+1, th.c)
		total += dfs(grid, 0, th.r, th.c-1)
		total += dfs(grid, 0, th.r, th.c+1)
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
