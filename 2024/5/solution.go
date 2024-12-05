package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

func checkUpdate(rules map[int]map[int]struct{}, update []int) bool {
	for i, u := range update {
		for j, n := range update[i+1:] {
			if _, ok := rules[u][n]; !ok {
				fmt.Println(n, i, j+i+1)
				return false
			}
		}
	}
	return true
}

func reorderUpdate(rules map[int]map[int]struct{}, update []int) []int {
	for !checkUpdate(rules, update) {
	checkForSwap:
		for i, u := range update {
			for j, n := range update[i+1:] {
				if _, ok := rules[u][n]; !ok {
					// swap
					update[i], update[j+i+1] = update[j+i+1], update[i]
					break checkForSwap
				}
			}
		}
	}

	return update
}

func part1() {
	rules := map[int]map[int]struct{}{}
	updates := [][]int{}

	lineFunc := func(line string, _ int) error {
		if strings.Contains(line, "|") {
			// rule
			nums := strings.Split(line, "|")
			i, j := nums[0], nums[1]

			before, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			after, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			if _, ok := rules[before]; !ok {
				rules[before] = map[int]struct{}{}
			}
			rules[before][after] = struct{}{}

		} else if strings.Contains(line, ",") {
			// update
			nums := strings.Split(line, ",")
			update := []int{}
			for _, n := range nums {
				i, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				update = append(update, i)
			}
			updates = append(updates, update)
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, update := range updates {
		if checkUpdate(rules, update) {
			total += utils.GetMiddleVal(update)
		}
	}
	fmt.Println(total)
}

func part2() {
	rules := map[int]map[int]struct{}{}
	updates := [][]int{}

	lineFunc := func(line string, _ int) error {
		if strings.Contains(line, "|") {
			// rule
			nums := strings.Split(line, "|")
			i, j := nums[0], nums[1]

			before, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			after, err := strconv.Atoi(j)
			if err != nil {
				panic(err)
			}
			if _, ok := rules[before]; !ok {
				rules[before] = map[int]struct{}{}
			}
			rules[before][after] = struct{}{}

		} else if strings.Contains(line, ",") {
			// update
			nums := strings.Split(line, ",")
			update := []int{}
			for _, n := range nums {
				i, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				update = append(update, i)
			}
			updates = append(updates, update)
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, update := range updates {
		if checkUpdate(rules, update) {
			//	total += utils.GetMiddleVal(update)
		} else {
			fmt.Println(update)
			reordered := reorderUpdate(rules, update)
			fmt.Println(reordered)
			total += utils.GetMiddleVal(reordered)
			// Need to reorder
		}
	}
	fmt.Println(total)
}

func main() {
	// part1()
	part2()
}
