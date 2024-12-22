package main

import (
	"fmt"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

func mix(n, p int64) int64 {
	return n ^ p
}

func prune(n int64) int64 {
	return n % 16777216
}

func calculateNextSecretNumber(n int64) int64 {
	val := n

	// first multiply by 64
	val *= 64
	val = mix(val, n)
	val = prune(val)

	// divide by 32
	prev := val
	val /= 32
	val = mix(val, prev)
	val = prune(val)

	prev = val
	val *= 2048
	val = mix(val, prev)
	val = prune(val)

	return val
}

func part1() {
	initialNumbers := []int64{}
	lineFunc := func(line string, _ int) error {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		initialNumbers = append(initialNumbers, int64(i))

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	secretNumbers := make([]int64, len(initialNumbers))
	copy(secretNumbers, initialNumbers)

	for i := 0; i < 2000; i++ {
		for j := 0; j < len(initialNumbers); j++ {
			secretNumbers[j] = calculateNextSecretNumber(secretNumbers[j])
		}
	}
	total := int64(0)
	for _, n := range secretNumbers {
		total += n
	}
	fmt.Println(total)
}

func main() {
	part1()
}
