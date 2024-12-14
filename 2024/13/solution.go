package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

var (
	buttonRe = regexp.MustCompile(`Button [A|B]: X\+(\d{1,4}), Y\+(\d{1,4})`)
	prizeRe  = regexp.MustCompile(`Prize: X=(\d{1,7}), Y=(\d{1,7})`)
)

const defaultCost = 1000000

type button struct {
	x, y int64
}

func buttonFromString(s string) button {
	matches := buttonRe.FindStringSubmatch(s)
	x, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	return button{int64(x), int64(y)}
}

type prize struct {
	x, y int64
}

func prizeFromString(s string) prize {
	matches := prizeRe.FindStringSubmatch(s)
	x, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}
	return prize{int64(x), int64(y)}
}

type game struct {
	a button
	b button
	p prize
}

func (g game) Cramers() (float64, float64) {
	det := float64(g.a.x*g.b.y - g.a.y*g.b.x)
	if det == 0 {
		return 0, 0
	}
	detX := float64(g.p.x*g.b.y - g.p.y*g.b.x)
	detY := float64(g.a.x*g.p.y - g.a.y*g.p.x)

	ansX, ansY := detX/det, detY/det
	a, f := math.Modf(ansX)
	if f != 0.0 {
		return 0, 0
	}
	b, f := math.Modf(ansY)
	if f != 0.0 {
		return 0, 0
	}

	return a * 3, b
}

type move struct {
	button button
}

type state struct {
	moves              []move
	x, y               int64
	aPresses, bPresses int
}

type cacheKey struct {
	aPresses, bPresses int
}

func (s state) cost() int {
	return s.aPresses*3 + s.bPresses
}

func permutations(buttons []button, finalX, finalY int64) int {
	lowestCost := defaultCost
	cache := map[cacheKey]struct{}{}

	var dfs func(s state)
	dfs = func(s state) {
		if s.x == finalX && s.y == finalY {
			if s.cost() < lowestCost {
				lowestCost = s.cost()
			}
			return
		}

		if s.x > finalX || s.y > finalY {
			return
		}

		if s.aPresses > 100 || s.bPresses > 100 {
			return
		}

		if s.cost() > lowestCost {
			return
		}

		if _, ok := cache[cacheKey{s.aPresses, s.bPresses}]; ok {
			return
		}

		cache[cacheKey{s.aPresses, s.bPresses}] = struct{}{}

		for i, b := range buttons {
			if i == 0 {
				s.aPresses++
			} else {
				s.bPresses++
			}
			s.moves = append(s.moves, move{b})
			s.x += b.x
			s.y += b.y
			dfs(s)
			s.y -= b.y
			s.x -= b.x
			s.moves = s.moves[:len(s.moves)-1]
			if i == 0 {
				s.aPresses--
			} else {
				s.bPresses--
			}
		}
	}

	dfs(state{[]move{}, 0, 0, 0, 0})
	return lowestCost
}

func part1() {
	games := []game{}
	lineFunc := func(lines []string, _ []int) error {
		a := buttonFromString(lines[0])
		b := buttonFromString(lines[1])
		p := prizeFromString(lines[2])

		games = append(games, game{a, b, p})
		return nil
	}
	err := utils.ReadFileInChunks("input.txt", 3, lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, g := range games {
		outcome := permutations([]button{g.a, g.b}, g.p.x, g.p.y)
		if outcome != defaultCost {
			total += outcome
		}
	}

	fmt.Println(total)
}

func part2() {
	games := []game{}
	lineFunc := func(lines []string, _ []int) error {
		a := buttonFromString(lines[0])
		b := buttonFromString(lines[1])
		p := prizeFromString(lines[2])
		p.x += 10000000000000
		p.y += 10000000000000

		games = append(games, game{a, b, p})
		return nil
	}
	err := utils.ReadFileInChunks("input.txt", 3, lineFunc)
	if err != nil {
		panic(err)
	}

	total := float64(0)
	for _, g := range games {
		a, b := g.Cramers()
		total += a + b
	}

	fmt.Println(int64(total))
}

func main() {
	part1()
	part2()
}
