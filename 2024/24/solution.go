package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

var (
	registerRegex  = regexp.MustCompile(`(\w\d\d): (1|0)`)
	operationRegex = regexp.MustCompile(`(\S{1,3}) (AND|XOR|OR) (\S{1,3}) -> (\S{1,3})`)
)

type register struct {
	name  string
	value int
}

type operation struct {
	in1, in2, out, op string
}

func and(a, b, out *register) {
	out.value = a.value & b.value
}

func or(a, b, out *register) {
	out.value = a.value | b.value
}

func xor(a, b, out *register) {
	out.value = a.value ^ b.value
}

func part1() {
	registers := make(map[string]*register)
	b := false
	operations := []operation{}

	lineFunc := func(line string, _ int) error {
		if line == "" {
			b = true
			return nil
		}
		if !b {
			// Parse register
			matches := registerRegex.FindStringSubmatch(line)
			val, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			registers[matches[1]] = &register{matches[1], val}
			return nil
		}

		// Parse operation
		matches := operationRegex.FindStringSubmatch(line)
		operations = append(operations, operation{matches[1], matches[3], matches[4], matches[2]})

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	var queue []operation
	queue = append(queue, operations...)
	for len(queue) > 0 {
		newQueue := []operation{}
		for _, op := range queue {
			// first check if in registers are defined
			// if not add it back to the queue
			if _, ok := registers[op.in1]; !ok {
				newQueue = append(newQueue, op)
				continue
			}
			if _, ok := registers[op.in2]; !ok {
				newQueue = append(newQueue, op)
				continue
			}

			// get the out register or create it
			var out *register
			var ok bool

			if out, ok = registers[op.out]; !ok {
				out = &register{op.out, 0}
			}

			// perform the operation
			switch op.op {
			case "AND":
				and(registers[op.in1], registers[op.in2], out)
			case "OR":
				or(registers[op.in1], registers[op.in2], out)
			case "XOR":
				xor(registers[op.in1], registers[op.in2], out)
			}

			// save the result
			registers[op.out] = out
		}
		queue = newQueue
	}

	zRegisters := []register{}
	for _, r := range registers {
		if strings.HasPrefix(r.name, "z") {
			zRegisters = append(zRegisters, *r)
		}
	}

	sort.Slice(zRegisters, func(i, j int) bool {
		return zRegisters[i].name > zRegisters[j].name
	})

	s := ""
	for _, r := range zRegisters {
		s += strconv.Itoa(r.value)
	}
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}

func main() {
	part1()
}
