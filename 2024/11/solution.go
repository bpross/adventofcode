package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bpross/adventofcode/utils"
)

var cache = map[int][]int{}

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
	if res, ok := cache[s]; ok {
		return res
	}

	if s == 0 {
		cache[s] = []int{1}
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
		cache[s] = res
		return res
	}

	// multiple value by 2024
	n := 2024 * s
	res := []int{n}
	cache[s] = res
	return res
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

	for i := 0; i < blinks; i++ {
		start := time.Now()
		newints := []int{}
		for _, s := range ints {
			newints = append(newints, performRules(s)...)
		}

		end := time.Now()
		fmt.Println("Blink", i, "took", end.Sub(start), "len", len(newints))
		ints = newints
	}

	fmt.Println(len(ints))
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

	blinks := 75

	for i := 0; i < blinks; i++ {
		start := time.Now()
		newints := []int{}
		for _, s := range ints {
			newints = append(newints, performRules(s)...)
		}
		end := time.Now()
		fmt.Println("Blink", i, "took", end.Sub(start), "len", len(newints))
		ints = newints
	}

	fmt.Println(len(ints))
}

func main() {
	part1()
	part2()
}
