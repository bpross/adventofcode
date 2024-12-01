package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func part1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	right := []int{}
	left := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
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
	}

	if err := scanner.Err(); err != nil {
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
		dist := abs(left[i], right[i])
		total += dist
	}

	fmt.Println(total)
}

func part2() {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	left := []int{}
	right := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
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
	}

	if err := scanner.Err(); err != nil {
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
