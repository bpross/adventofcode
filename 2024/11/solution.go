package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bpross/adventofcode/utils"
)

type cacheKey struct {
	val, depth int
}

var cache = map[cacheKey]int{}

func splitValue(value int) []int {
	valueSplit := []int{}
	for value > 9 {
		modulo := value % 10
		valueSplit = append(valueSplit, modulo)
		value = (value - modulo) / 10
	}
	valueSplit = append(valueSplit, value)
	return valueSplit
}

func performRules(s int) []int {
	if s == 0 {
		return []int{1}
	}

	vs := splitValue(s)
	if len(vs)%2 == 0 {
		half := len(vs) / 2
		first, second := 0, 0
		idx := 0
		mult := 1
		for idx < half {
			first += vs[idx] * mult
			idx++
			mult *= 10
		}
		mult = 1
		for idx < len(vs) {
			second += vs[idx] * mult
			idx++
			mult *= 10
		}
		res := []int{second, first}
		return res
	}

	// multiple value by 2024
	n := 2024 * s
	res := []int{n}
	return res
}

func dfs(s, depth int) int {
	if depth == 75 {
		return 1
	}

	if val, ok := cache[cacheKey{s, depth}]; ok {
		return val
	}

	ns := performRules(s)
	total := 0
	for _, n := range ns {
		total += dfs(n, depth+1)
	}

	cache[cacheKey{s, depth}] = total
	return total
}

func part1() {
	ints := []int{}

	lineFunc := func(line string, _ int) error {
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		for _, c := range values {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			ints = append(ints, v)
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(ints)

	blinks := 25

	start := time.Now()
	for i := 0; i < blinks; i++ {
		newints := []int{}
		for _, s := range ints {
			newints = append(newints, performRules(s)...)
		}
		ints = newints
	}
	end := time.Now()
	fmt.Println("Total", len(ints), "took", end.Sub(start))
}

func part2() {
	ints := []int{}

	lineFunc := func(line string, _ int) error {
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		for _, c := range values {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			ints = append(ints, v)
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(ints)

	start := time.Now()
	total := 0
	for _, s := range ints {
		total += dfs(s, 0)
	}

	end := time.Now()
	fmt.Println("Total", total, "took", end.Sub(start))
}

func main() {
	part1()
	part2()
}
