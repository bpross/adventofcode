package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

type key struct {
	values []string
	ridges []int
}

func (k *key) setRidges() {
	k.ridges = make([]int, len(k.values[0]))
	for i := 0; i < len(k.values[0]); i++ {
		cnt := 0
		for j := len(k.values) - 2; j >= 0; j-- {
			if k.values[j][i] == '#' {
				cnt++
			}
		}
		k.ridges[i] = cnt
	}
}

func (k *key) Print() {
	for _, v := range k.values {
		println(v)
	}
}

type lock struct {
	values []string
	ridges []int
}

func (l *lock) Print() {
	for _, v := range l.values {
		println(v)
	}
}

func (l *lock) setRidges() {
	l.ridges = make([]int, len(l.values[0]))
	for i := 0; i < len(l.values[0]); i++ {
		cnt := 0
		for j := 1; j < len(l.values); j++ {
			if l.values[j][i] == '#' {
				cnt++
			}
		}
		l.ridges[i] = cnt
	}
}

func checkOVerlap(k key, l lock) bool {
	for i := 0; i < len(k.ridges); i++ {
		if k.ridges[i]+l.ridges[i] > 5 {
			return false
		}
	}
	return true
}

func part1() {
	locks := make([]lock, 0)
	keys := make([]key, 0)
	lineFunc := func(lines []string, _ []int) error {
		// check if the top line is all #
		// if so its a lock
		l := true
		for _, c := range lines[0] {
			if c != '#' {
				l = false
				break
			}
		}
		if l {
			locks = append(locks, lock{values: lines})
			locks[len(locks)-1].setRidges()
		} else {
			keys = append(keys, key{values: lines})
			keys[len(keys)-1].setRidges()
		}

		return nil
	}

	err := utils.ReadFileInChunks("input.txt", 7, lineFunc)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, k := range keys {
		for _, l := range locks {
			if checkOVerlap(k, l) {
				total++
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
}
