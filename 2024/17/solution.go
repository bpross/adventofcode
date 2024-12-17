package main

import (
	"fmt"

	"github.com/bpross/adventofcode/utils"
)

const (
	// opcodes
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type register struct {
	value int
}

type machine struct {
	a, b, c        register
	instructions   []int
	instructionPtr int
}

func (m machine) String() string {
	return fmt.Sprintf("\na: %d, b: %d, c: %d, instructions: %v, instructionPtr: %d", m.a, m.b, m.c, m.instructions, m.instructionPtr)
}

func (m machine) operandToValue(operand int) int {
	switch operand {
	case 0, 1, 2, 3, 7:
		return operand
	case 4:
		return m.a.value
	case 5:
		return m.b.value
	case 6:
		return m.c.value
	}
	return -1
}

func (m *machine) run() {
	for m.instructionPtr < len(m.instructions) {
		opcode := m.instructions[m.instructionPtr]
		operand := m.instructions[m.instructionPtr+1]

		val := m.operandToValue(operand)

		switch opcode {
		case adv:
			denom := utils.Pow(2, val)
			result := m.a.value / denom
			m.a.value = result
			m.instructionPtr += 2
		case bxl:
			result := m.b.value ^ operand
			m.b.value = result
			m.instructionPtr += 2
		case bst:
			val = val % 8
			m.b.value = val
			m.instructionPtr += 2
		case jnz:
			if m.a.value == 0 {
				m.instructionPtr += 2
			} else {
				m.instructionPtr = operand
			}
		case bxc:
			result := m.b.value ^ m.c.value
			m.b.value = result
			m.instructionPtr += 2
		case out:
			val = val % 8
			fmt.Printf("%d,", val)
			m.instructionPtr += 2
		case bdv:
			denom := utils.Pow(2, val)
			result := m.a.value / denom
			m.b.value = result
			m.instructionPtr += 2
		case cdv:
			denom := utils.Pow(2, val)
			result := m.a.value / denom
			m.c.value = result
			m.instructionPtr += 2
		}
	}
}

func newMachine(a, b, c register, instructions []int) machine {
	return machine{a, b, c, instructions, 0}
}

func part1() {
	// small input so just easy to just construct
	instructions := []int{2, 4, 1, 3, 7, 5, 0, 3, 4, 1, 1, 5, 5, 5, 3, 0}
	machine := newMachine(register{45483412}, register{0}, register{0}, instructions)

	machine.run()
	fmt.Println(machine)
}

func main() {
	part1()
}
