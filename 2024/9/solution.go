package main

import (
	"fmt"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

type block struct {
	data string
}

type disk struct {
	blocks []block
}

func (d disk) String() string {
	out := ""

	for _, b := range d.blocks {
		out += b.data
	}

	return out
}

func (d disk) Compact() {
	// two pointers
	// find first empty block
	// move first file block to empty block
	// repeat until pointers equal each other
	i, j := 0, len(d.blocks)-1

	// advance i to first empty block
	for i < j && d.blocks[i].data != "." {
		i++
	}

	// advance j to first file block
	for j > i && d.blocks[j].data == "." {
		j--
	}

	for i < j {
		// move file block to empty block
		d.blocks[i], d.blocks[j] = d.blocks[j], d.blocks[i]
		// advance i to next empty block
		for i < j && d.blocks[i].data != "." {
			i++
		}
		// advance j to first file block
		for j > i && d.blocks[j].data == "." {
			j--
		}
	}
}

func (d disk) Checksum() int64 {
	total := int64(0)
	cnt := 0
	for i := 0; i < len(d.blocks); i++ {
		if d.blocks[i].data == "." {
			continue
		}
		id, err := strconv.Atoi(string(d.blocks[i].data))
		if err != nil {
			panic(err)
		}
		total += int64(id) * int64(cnt)
		cnt++
	}
	return total
}

func part1() {
	d := disk{blocks: []block{}}

	// file is one line
	lineFunc := func(line string, _ int) error {
		id := 0
		for i, c := range line {
			blocks := []block{}
			size, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			if i == 0 || i%2 == 0 {
				// file
				for j := 0; j < size; j++ {
					blocks = append(blocks, block{data: strconv.Itoa(id)})
				}
				id++
			} else {
				// free space
				for j := 0; j < size; j++ {
					blocks = append(blocks, block{data: "."})
				}
			}
			for _, b := range blocks {
				d.blocks = append(d.blocks, b)
			}
		}
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	d.Compact()
	checksum := d.Checksum()
	fmt.Println(checksum)
}

func main() {
	part1()
}
