package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

func part1() {
	total := 0

	lineFunc := func(line string, lineNumber int) error {
		fmt.Println("line", lineNumber)
		increase, decrease := false, false
		nums := strings.Split(line, " ")

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
				fmt.Println("Equal", nums[j], nums[i], line)
				// equal, needs to increase OR decrease
				return nil
			}
			if jInt < iInt && increase {
				fmt.Println("not increasing", nums[j], nums[i], line)
				return nil
			}
			if jInt > iInt && decrease {
				fmt.Println("not decreasing", nums[j], nums[i], line)
				return nil
			}

			diff := utils.Abs(iInt, jInt)
			if diff < 1 || diff >= 4 {
				fmt.Println("diff", iInt, jInt, line, diff)
				return nil
			}

			i++
			j++
		}

		// Safe
		total += 1
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
}
