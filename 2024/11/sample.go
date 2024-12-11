package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type stone struct {
	val string
}

func trimZeroes(s string) string {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(v)
}

func performRules(stones stone) []stone {
	res := []stone{}

	if stones.val == "0" {
		res = append(res, stone{"1"})
		return res
	} else if len(stones.val)%2 == 0 {
		// break into two
		s1 := stone{trimZeroes(stones.val[:len(stones.val)/2])}
		s2 := stone{trimZeroes(stones.val[len(stones.val)/2:])}
		res = append(res, s1, s2)
		return res
	}
	// multiple value by 2024
	v, err := strconv.Atoi(stones.val)
	if err != nil {
		panic(err)
	}
	v *= 2024
	res = append(res, stone{strconv.Itoa(v)})
	return res
}

func part1() {
	stones := []stone{}

	lineFunc := func(line string, _ int) error {
		line = strings.TrimSpace(line)
		values := strings.Split(line, " ")
		for _, c := range values {
			stones = append(stones, stone{string(c)})
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(stones)

	blinks := 25

	for i := 0; i < blinks; i++ {
		newStones := []stone{}
		for _, s := range stones {
			newStones = append(newStones, performRules(s)...)
		}
		fmt.Println(newStones)
		stones = newStones
	}

	fmt.Println(len(stones))
}

func main() {
	part1()
	// part2()
}
