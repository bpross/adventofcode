package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

// key pad layout
// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//
//	| 0 | A |
//	+---+---+
var numericToMoves = map[string]map[string][][]string{
	"7": {
		"8": [][]string{{">", "A"}},
		"9": [][]string{{">", ">", "A"}},
		"4": [][]string{{"v", "A"}},
		"5": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
		"6": [][]string{{"v", ">", ">", "A"}, {">", "v", ">", "A"}, {">", ">", "v", "A"}},
		"1": [][]string{{"v", "v", "A"}},
		"2": [][]string{{"v", "v", ">", "A"}, {">", "v", "v", "A"}, {"v", ">", "v", "A"}},
		"3": [][]string{{"v", "v", ">", ">", "A"}, {">", "v", "v", ">", "A"}, {">", "v", ">", "v", "A"}, {">", ">", "v", "v", "A"}},
		"0": [][]string{{"v", "v", ">", "v", "A"}, {"v", ">", "v", "v", "A"}, {">", "v", "v", "v", "A"}},
		"A": [][]string{{"v", "v", ">", "v", ">", "A"}, {"v", ">", "v", "v", ">", "A"}, {"v", "v", ">", ">", "v", "A"}, {">", "v", "v", "v", ">", "A"}, {">", "v", ">", "v", "v", "A"}, {">", ">", "v", "v", "v", "A"}},
		"7": [][]string{{"A"}},
	},
	"8": {
		"8": [][]string{{"A"}},
		"9": [][]string{{">", "A"}},
		"7": [][]string{{"<", "A"}},
		"4": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		"5": [][]string{{"v", "A"}},
		"6": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
		"1": [][]string{{"v", "v", "<", "A"}, {"<", "v", "v", "A"}, {"v", "<", "v", "A"}},
		"2": [][]string{{"v", "v", "A"}},
		"3": [][]string{{"v", "v", ">", "A"}, {">", "v", "v", "A"}, {"v", ">", "v", "A"}},
		"0": [][]string{{"v", "v", "v", "A"}},
		"A": [][]string{{"v", "v", "v", ">", "A"}, {">", "v", "v", "v", "A"}, {"v", "v", ">", "v", "A"}, {"v", ">", "v", "v", "A"}},
	},
	"9": {
		"9": [][]string{{"A"}},
		"8": [][]string{{"<", "A"}},
		"7": [][]string{{"<", "<", "A"}},
		"4": [][]string{{"v", "<", "<", "A"}, {"<", "v", "<", "A"}, {"<", "<", "v", "A"}},
		"5": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		"6": [][]string{{"v", "A"}},
		"1": [][]string{{"v", "v", "<", "<", "A"}, {"<", "v", "v", "<", "A"}, {"<", "<", "<", "v", "A"}},
		"2": [][]string{{"v", "v", "<", "A"}, {"<", "v", "v", "A"}},
		"3": [][]string{{"v", "v", "A"}},
		"0": [][]string{{"v", "v", "v", "<", "A"}, {"<", "v", "v", "v", "A"}, {"v", "<", "v", "v", "A"}, {"v", "v", "<", "v", "A"}},
		"A": [][]string{{"v", "v", "v", "A"}},
	},
	"4": {
		"4": [][]string{{"A"}},
		"5": [][]string{{">", "A"}},
		"6": [][]string{{">", ">", "A"}},
		"1": [][]string{{"v", "A"}},
		"2": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
		"3": [][]string{{"v", ">", ">", "A"}, {">", "v", ">", "A"}, {">", ">", "v", "A"}},
		"7": [][]string{{"^", "A"}},
		"8": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		"9": [][]string{{"^", ">", ">", "A"}, {">", "^", ">", "A"}, {">", ">", "^", "A"}},
		"0": [][]string{{"v", ">", "v", "A"}, {">", "v", "v", "A"}},
		"A": [][]string{{"v", ">", "v", ">", "A"}, {">", "v", "v", ">", "A"}, {">", "v", ">", "v", "A"}, {">", ">", "v", "v", "A"}},
	},
	"5": {
		"5": [][]string{{"A"}},
		"4": [][]string{{"<", "A"}},
		"6": [][]string{{">", "A"}},
		"1": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		"2": [][]string{{"v", "A"}},
		"3": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
		"7": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
		"8": [][]string{{"^", "A"}},
		"9": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		"0": [][]string{{"v", "v", "A"}},
		"A": [][]string{{"v", "v", ">", "A"}, {">", "v", "v", "A"}, {"v", ">", "v", "A"}},
	},
	"6": {
		"6": [][]string{{"A"}},
		"5": [][]string{{"<", "A"}},
		"4": [][]string{{"<", "<", "A"}},
		"1": [][]string{{"v", "<", "<", "A"}, {"<", "v", "<", "A"}, {"<", "<", "v", "A"}},
		"2": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		"3": [][]string{{"v", "A"}},
		"7": [][]string{{"^", "<", "<", "A"}, {"<", "^", "<", "A"}, {"<", "<", "^", "A"}},
		"8": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
		"9": [][]string{{"^", "A"}},
		"0": [][]string{{"v", "<", "v", "A"}, {"<", "v", "v", "A"}},
		"A": [][]string{{"v", "v", "A"}},
	},
	"1": {
		"1": [][]string{{"A"}},
		"2": [][]string{{">", "A"}},
		"3": [][]string{{">", ">", "A"}},
		"4": [][]string{{"^", "A"}},
		"5": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		"6": [][]string{{"^", ">", ">", "A"}, {">", "^", ">", "A"}, {">", ">", "^", "A"}},
		"7": [][]string{{"^", "^", "A"}},
		"8": [][]string{{"^", "^", ">", "A"}, {">", "^", "^", "A"}, {"^", ">", "^", "A"}},
		"9": [][]string{{"^", "^", ">", ">", "A"}, {">", "^", "^", ">", "A"}, {"^", ">", "^", ">", "A"}, {">", ">", "^", "^", "A"}},
		"0": [][]string{{">", "v", "A"}},
		"A": [][]string{{">", "v", ">", "A"}, {">", ">", "v", "A"}},
	},
	"2": {
		"2": [][]string{{"A"}},
		"1": [][]string{{"<", "A"}},
		"3": [][]string{{">", "A"}},
		"4": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
		"5": [][]string{{"^", "A"}},
		"6": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		"7": [][]string{{"^", "^", "<", "A"}, {"<", "^", "^", "A"}, {"^", "<", "^", "A"}},
		"8": [][]string{{"^", "^", "A"}},
		"9": [][]string{{"^", "^", ">", "A"}, {">", "^", "^", "A"}, {"^", ">", "^", "A"}},
		"0": [][]string{{"v", "A"}},
		"A": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
	},
	"3": {
		"3": [][]string{{"A"}},
		"2": [][]string{{"<", "A"}},
		"1": [][]string{{"<", "<", "A"}},
		"4": [][]string{{"^", "<", "<", "A"}, {"<", "^", "<", "A"}, {"<", "<", "^", "A"}},
		"5": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
		"6": [][]string{{"^", "A"}},
		"7": [][]string{{"<", "<", "^", "^", "A"}, {"<", "^", "<", "^", "A"}, {"^", "<", "<", "^", "A"}, {"^", "<", "^", "<", "A"}, {"^", "^", "<", "<", "A"}},
		"8": [][]string{{"^", "^", "<", "A"}, {"<", "^", "^", "A"}, {"^", "<", "^", "A"}},
		"9": [][]string{{"^", "^", "A"}},
		"0": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		"A": [][]string{{"v", "A"}},
	},
	"0": {
		"0": [][]string{{"A"}},
		"A": [][]string{{">", "A"}},
		"1": [][]string{{"^", "<", "A"}},
		"2": [][]string{{"^", "A"}},
		"3": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		"4": [][]string{{"^", "^", "<", "A"}, {"^", "<", "^", "A"}},
		"5": [][]string{{"^", "^", "A"}},
		"6": [][]string{{"^", "^", ">", "A"}, {">", "^", "^", "A"}},
		"7": [][]string{{"^", "^", "^", "<", "A"}, {"^", "<", "^", "^", "A"}, {"^", "^", "<", "^", "A"}},
		"8": [][]string{{"^", "^", "^", "A"}},
		"9": [][]string{{"^", "^", "^", ">", "A"}, {">", "^", "^", "^", "A"}, {"^", ">", "^", "^", "A"}, {"^", "^", ">", "^", "A"}},
	},
	"A": {
		"A": [][]string{{"A"}},
		"0": [][]string{{"<", "A"}},
		"1": [][]string{{"^", "<", "<", "A"}, {"<", "^", "<", "A"}},
		"2": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
		"3": [][]string{{"^", "A"}},
		"4": [][]string{{"^", "^", "<", "<", "A"}, {"<", "^", "^", "<", "A"}, {"^", "<", "^", "<", "A"}, {"<", "^", "<", "^", "A"}},
		"5": [][]string{{"^", "^", "<", "A"}, {"<", "^", "^", "A"}, {"^", "<", "^", "A"}},
		"6": [][]string{{"^", "^", "A"}},
		"7": [][]string{{"^", "^", "^", "<", "<", "A"}, {"<", "^", "^", "^", "<", "A"}, {"^", "<", "^", "^", "<", "A"}, {"^", "<", "<", "^", "^", "A"}, {"^", "^", "<", "<", "^", "A"}},
		"8": [][]string{{"^", "^", "^", "<", "A"}, {"<", "^", "^", "^", "A"}, {"^", "<", "^", "^", "A"}, {"^", "^", "<", "^", "A"}},
		"9": [][]string{{"^", "^", "^", "A"}},
	},
}

// key pad layout
//
//	+---+---+
//	| ^ | A |
//
// +---+---+---+
// | < | v | > |
// +---+---+---+
var directionToMove = map[string]map[string][][]string{
	"^": {
		"^": [][]string{{"A"}},
		"A": [][]string{{">", "A"}},
		"<": [][]string{{"v", "<", "A"}},
		"v": [][]string{{"v", "A"}},
		">": [][]string{{"v", ">", "A"}, {">", "v", "A"}},
	},
	"A": {
		"A": [][]string{{"A"}},
		"^": [][]string{{"<", "A"}},
		"<": [][]string{{"v", "<", "<", "A"}, {"<", "v", "<", "A"}},
		"v": [][]string{{"v", "<", "A"}, {"<", "v", "A"}},
		">": [][]string{{"v", "A"}},
	},
	"<": {
		"<": [][]string{{"A"}},
		"A": [][]string{{">", ">", "^", "A"}, {">", "^", ">", "A"}},
		"^": [][]string{{">", "^", "A"}},
		"v": [][]string{{">", "A"}},
		">": [][]string{{">", ">", "A"}},
	},
	"v": {
		"v": [][]string{{"A"}},
		"<": [][]string{{"<", "A"}},
		"^": [][]string{{"^", "A"}},
		"A": [][]string{{"^", ">", "A"}, {">", "^", "A"}},
		">": [][]string{{">", "A"}},
	},
	">": {
		">": [][]string{{"A"}},
		"A": [][]string{{"^", "A"}},
		"<": [][]string{{"<", "<", "A"}},
		"v": [][]string{{"<", "A"}},
		"^": [][]string{{"^", "<", "A"}, {"<", "^", "A"}},
	},
}

var iStrToInt = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func codeToInt(code string) int {
	i := 0
	for _, c := range code {
		if _, ok := iStrToInt[string(c)]; !ok {
			continue
		}
		i = i*10 + iStrToInt[string(c)]
	}
	return i
}

func dfsNumeric(code string) []string {
	res := []string{}
	var dfs func(current, code string, path []string)
	dfs = func(current, code string, path []string) {
		if len(code) == 0 {
			p := strings.Join(path, "")
			res = append(res, p)
			return
		}
		for _, move := range numericToMoves[current][string(code[0])] {
			dfs(string(code[0]), code[1:], append(path, move...))
		}
	}
	dfs(string(code[0]), code[1:], []string{})
	return res
}

func dfsDirection(code string) []string {
	res := []string{}
	var dfs func(current, code string, path []string)
	dfs = func(current, code string, path []string) {
		if len(code) == 0 {
			p := strings.Join(path, "")
			res = append(res, p)
			return
		}
		for _, move := range directionToMove[current][string(code[0])] {
			dfs(string(code[0]), code[1:], append(path, move...))
		}
	}
	dfs(string(code[0]), code[1:], []string{})
	return res
}

func part1() {
	codes := []string{}

	lineFunc := func(line string, _ int) error {
		line = strings.TrimSpace(line)
		codes = append(codes, line)
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	for _, code := range codes {
		fmt.Println(code)
	}

	total := 0
	for _, code := range codes {
		paths := dfsNumeric("A" + code)
		ps := make(map[int][]string)
		min := math.MaxInt64
		for _, path := range paths {
			ps[len(path)] = append(ps[len(path)], path)
			if len(path) < min {
				min = len(path)
			}
		}

		paths = ps[min]

		newPaths := []string{}
		for _, path := range paths {
			newPaths = append(newPaths, dfsDirection("A"+path)...)
		}

		ps = make(map[int][]string)
		min = math.MaxInt64
		for _, path := range newPaths {
			ps[len(path)] = append(ps[len(path)], path)
			if len(path) < min {
				min = len(path)
			}
		}

		paths = ps[min]

		newPaths = []string{}
		for _, path := range paths {
			newPaths = append(newPaths, dfsDirection("A"+path)...)
		}

		ps = make(map[int][]string)
		min = math.MaxInt64
		for _, path := range newPaths {
			ps[len(path)] = append(ps[len(path)], path)
			if len(path) < min {
				min = len(path)
			}
		}

		num := codeToInt(code)
		fmt.Println(min, num)
		total += min * num
	}
	fmt.Println(total)
}

func main() {
	part1()
}
