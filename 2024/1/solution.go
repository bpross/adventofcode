package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

func part1() {
	right := []int{}
	left := []int{}

	lineFunc := func(line string) error {
		nums := strings.Split(line, "   ")
		i, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, i)
		i, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right = append(right, i)

		return nil
	}

	err := utils.ReadFile("input1.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	total := 0

	for i := 0; i < len(left); i++ {
		dist := utils.Abs(left[i], right[i])
		total += dist
	}

	fmt.Println(total)
}

func part2() {
	left := []int{}
	right := map[int]int{}

	lineFunc := func(line string) error {
		nums := strings.Split(line, "   ")
		i, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, i)
		i, err = strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		right[i] += 1

		return nil
	}

	err := utils.ReadFile("input1.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0

	for i := 0; i < len(left); i++ {
		cnt := right[left[i]]
		total += cnt * left[i]
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
