package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

func checkIfSafe(nums []string) bool {
	increase, decrease := false, false
	i, j := 0, 1
	iInt, err := strconv.Atoi(nums[i])
	if err != nil {
		panic(err)
	}

	jInt, err := strconv.Atoi(nums[j])
	if err != nil {
		panic(err)
	}
	if iInt > jInt {
		decrease = true
	} else {
		increase = true
	}

	for j < len(nums) {
		iInt, err = strconv.Atoi(nums[i])
		if err != nil {
			panic(err)
		}

		jInt, err = strconv.Atoi(nums[j])
		if err != nil {
			panic(err)
		}

		if iInt == jInt {
			return false
		}
		if jInt < iInt && increase {
			return false
		}
		if jInt > iInt && decrease {
			return false
		}

		diff := utils.Abs(iInt, jInt)
		if diff < 1 || diff >= 4 {
			return false
		}

		i++
		j++
	}

	return true
}

func part1() {
	total := 0

	lineFunc := func(line string, _ int) error {
		nums := strings.Split(line, " ")

		safe := checkIfSafe(nums)
		if safe {
			total += 1
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func part2() {
	total := 0

	lineFunc := func(line string, _ int) error {
		nums := strings.Split(line, " ")

		safe := checkIfSafe(nums)
		if safe {
			total += 1
			return nil
		}

		for i := 0; i < len(nums); i++ {
			modified := nums
			modified = utils.RemoveIndex(modified, i)
			safe = checkIfSafe(modified)
			if safe {
				total += 1
				break
			}
		}

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func main() {
	part1()
	part2()
}
