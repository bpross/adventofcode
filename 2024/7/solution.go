package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type equation struct {
	total  int
	values []int
}

// generate all permutations of values and operations
func permutations(values []string, wantLength int) []string {
	found := make(map[string]bool)
	res := make([]string, 0)

	var dfs func(path string)
	dfs = func(path string) {
		if len(path) == wantLength {
			if _, ok := found[path]; ok {
				return
			}
			found[path] = true
			fmt.Println("Path:", path)
			res = append(res, path)
			return
		}
		for i := 0; i < len(values); i++ {
			fmt.Println("Path:", path, "Values:", values[i])
			dfs(path + values[i])
		}
	}
	dfs("")

	return res
}

func checkPermutations(e equation, perms []string) bool {
	for _, p := range perms {
		total := e.total
		for i, v := range e.values {
			if i == 0 {
				total = v
				continue
			}
			if p[i-1] == '+' {
				total += v
			} else {
				total *= v
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
		vals := make([]int, 0)
		for _, v := range products {
			val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			vals = append(vals, val)
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
		} else {
			fmt.Println(e)
			fmt.Println(perms)
		}
	}

	fmt.Println(trueEquations)
	total := int64(0)
	for _, e := range trueEquations {
		total += int64(e.total)
	}

	fmt.Println(total)
}

func main() {
	part1()
}
