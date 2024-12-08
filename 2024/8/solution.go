package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	r, c int
}

func drawMap(frequencyPositions map[rune][]position, antiNodePositions map[position]struct{}, rows, cols int) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			p := position{r, c}
			if _, ok := antiNodePositions[p]; ok {
				fmt.Print("#")
			} else {
				found := false
				for k, v := range frequencyPositions {
					for _, pos := range v {
						if pos == p {
							fmt.Print(string(k))
							found = true
							break
						}
					}
					if found {
						break
					}
				}
				if !found {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func part2() {
	frequencyPositions := map[rune][]position{}
	rows, cols := 0, 0

	lineFunc := func(line string, r int) error {
		rows++
		cols = len(line)

		idx := 0
		for idx < len(line) {
			char := rune(line[idx])
			if char != '.' {
				frequencyPositions[char] = append(frequencyPositions[char], position{r, idx})
			}
			idx++
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	antiNodePositions := map[position]struct{}{}
	for _, positions := range frequencyPositions {
		if len(positions) == 1 {
			continue
		}

		for i, p := range positions {
			for _, p2 := range positions[i+1:] {
				antiNodePositions[p] = struct{}{}
				antiNodePositions[p2] = struct{}{}
				rDiff := p2.r - p.r // we know that p2 will be greater or equal to p
				cDiff := p.c - p2.c

				if cDiff > 0 {
					// check up to right and down to left
					r := p.r - rDiff
					c := p.c + cDiff
					for r >= 0 && c < cols {
						antiNodePositions[position{r, c}] = struct{}{}
						r -= rDiff
						c += cDiff
					}

					r = p2.r + rDiff
					c = p2.c - cDiff
					for r < rows && c >= 0 {
						antiNodePositions[position{r, c}] = struct{}{}
						r += rDiff
						c -= cDiff
					}
				} else {
					// check up to left and down to right
					cDiff = -cDiff

					r := p.r - rDiff
					c := p.c - cDiff
					for r >= 0 && c >= 0 {
						antiNodePositions[position{r, c}] = struct{}{}
						r -= rDiff
						c -= cDiff
					}

					r = p2.r + rDiff
					c = p2.c + cDiff

					for r < rows && c < cols {
						antiNodePositions[position{r, c}] = struct{}{}
						r += rDiff
						c += cDiff
					}
				}
			}
		}
	}

	fmt.Println(len(antiNodePositions))

	drawMap(frequencyPositions, antiNodePositions, rows, cols)
}

func part1() {
	frequencyPositions := map[rune][]position{}
	rows, cols := 0, 0

	lineFunc := func(line string, r int) error {
		rows++
		cols = len(line)

		idx := 0
		for idx < len(line) {
			char := rune(line[idx])
			if char != '.' {
				frequencyPositions[char] = append(frequencyPositions[char], position{r, idx})
			}
			idx++
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	antiNodePositions := map[position]struct{}{}
	for _, positions := range frequencyPositions {
		if len(positions) == 1 {
			continue
		}

		for i, p := range positions {
			for _, p2 := range positions[i+1:] {
				rDiff := p2.r - p.r // we know that p2 will be greater or equal to p
				cDiff := p.c - p2.c

				if cDiff > 0 {
					// check up to right and down to left
					if p.r-rDiff >= 0 && p.c+cDiff < cols {
						antiNodePositions[position{p.r - rDiff, p.c + cDiff}] = struct{}{}
					}

					if p2.r+rDiff < rows && p2.c-cDiff >= 0 {
						antiNodePositions[position{p2.r + rDiff, p2.c - cDiff}] = struct{}{}
					}
				} else {
					// check up to left and down to right
					cDiff = -cDiff
					if p.r-rDiff >= 0 && p.c-cDiff >= 0 {
						antiNodePositions[position{p.r - rDiff, p.c - cDiff}] = struct{}{}
					}
					if p2.r+rDiff < rows && p2.c+cDiff < cols {
						antiNodePositions[position{p2.r + rDiff, p2.c + cDiff}] = struct{}{}
					}
				}
			}
		}
	}

	fmt.Println(len(antiNodePositions))
}

func main() {
	// part1()
	part2()
}
