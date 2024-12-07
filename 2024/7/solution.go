package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type equation struct {
	total  int
	values []string
}

// generate all permutations of values and operations
func permutations(values []string, wantLength int) [][]string {
	found := make(map[string]bool)
	res := make([][]string, 0)

	var dfs func(path string, ops int)
	dfs = func(path string, ops int) {
		if ops == wantLength {
			if _, ok := found[path]; ok {
				return
			}
			found[path] = true
			var allOps []string
			idx := 0
			for idx < len(path) {
				if path[idx] == '|' {
					allOps = append(allOps, "||")
					idx += 2
				} else {
					allOps = append(allOps, string(path[idx]))
					idx++
				}
			}
			res = append(res, allOps)
			return
		}
		for i := 0; i < len(values); i++ {
			ops++
			dfs(path+values[i], ops)
			ops--
		}
	}
	dfs("", 0)

	return res
}

func checkPermutations(e equation, perms [][]string) bool {
	for _, p := range perms {
		total := 0
		idx := 0
		for idx < len(e.values) {
			k, err := strconv.Atoi(e.values[idx])
			if err != nil {
				panic(err)
			}

			if idx == 0 {
				total = k
				idx++
				continue
			}

			switch p[idx-1] {
			case "+":
				idx++
				total += k
			case "*":
				idx++
				total *= k
			case "||":
				t := strconv.Itoa(total)
				t += e.values[idx]
				total, err = strconv.Atoi(t)
				if err != nil {
					panic(err)
				}
				idx++
			}
		}

		if total == e.total {
			return true
		}
	}
	return false
}

func part1() {
	equations := make([]equation, 0)

	lineFunc := func(line string, _ int) error {
		nums := strings.Split(line, ":")
		total, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		p := strings.Trim(nums[1], " ")
		products := strings.Split(p, " ")
		vals := make([]string, 0)
		for _, v := range products {
			vals = append(vals, v)
		}

		e := equation{total: total, values: vals}
		equations = append(equations, e)
		return nil
	}
	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	trueEquations := make([]equation, 0)
	ops := []string{"+", "*"}
	for _, e := range equations {
		perms := permutations(ops, len(e.values)-1)
		if checkPermutations(e, perms) {
			trueEquations = append(trueEquations, e)
		}
	}

	fmt.Println(trueEquations)
	total := int64(0)
	for _, e := range trueEquations {
		total += int64(e.total)
	}

	fmt.Println(total)
}

func part2() {
	equations := make([]equation, 0)

	lineFunc := func(line string, _ int) error {
		nums := strings.Split(line, ":")
		total, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		p := strings.Trim(nums[1], " ")
		products := strings.Split(p, " ")
		vals := make([]string, 0)
		for _, v := range products {
			vals = append(vals, v)
		}

		e := equation{total: total, values: vals}
		equations = append(equations, e)
		return nil
	}
	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	trueEquations := make([]equation, 0)
	ops := []string{"+", "*", "||"}
	for _, e := range equations {
		perms := permutations(ops, len(e.values)-1)
		if checkPermutations(e, perms) {
			trueEquations = append(trueEquations, e)
		}
	}

	total := int64(0)
	for _, e := range trueEquations {
		total += int64(e.total)
	}

	fmt.Println(total)
}

func main() {
	// part1()
	part2()
}
