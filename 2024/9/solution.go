package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

type block struct {
	size int
	data string
}

type disk struct {
	blocks []block
}

func (d disk) String() string {
	out := ""

	for _, b := range d.blocks {
		for i := 0; i < b.size; i++ {
			out += b.data
		}
	}

	return out
}

func advanceToNextFreeBlock(blocks []block, i int) int {
	if blocks[i].data == "." {
		i++
	}
	for i < len(blocks) && blocks[i].data != "." {
		i++
	}
	return i
}

func advanceToNextFileBlock(blocks []block, i int) int {
	if blocks[i].data != "." {
		i--
	}
	for i > 0 && blocks[i].data == "." {
		i--
	}
	return i
}

func (d disk) Compact() {
	// two pointers
	// find first empty block
	// move first file block to empty block
	// repeat until pointers equal each other
	i, j := 0, len(d.blocks)-1

	// advance i to first empty block
	i = advanceToNextFreeBlock(d.blocks, i)

	// advance j to first file block
	j = advanceToNextFileBlock(d.blocks, j)

	fmt.Println(i, j)

	for i < j {
		// move file block to empty block
		d.blocks[i], d.blocks[j] = d.blocks[j], d.blocks[i]
		// advance i to next empty block
		i = advanceToNextFreeBlock(d.blocks, i)
		// advance j to first file block
		j = advanceToNextFileBlock(d.blocks, j)
	}
}

// func (d *disk) ScanForFreeBlocks() {
// 	d.freeBlocks = map[int][]freeBlocks{}
// 	for i := 0; i < len(d.blocks); i++ {
// 		if d.blocks[i].data == "." {
// 			start := i
// 			for i < len(d.blocks) && d.blocks[i].data == "." {
// 				i++
// 			}
// 			d.freeBlocks[i-start] = append(d.freeBlocks[i-start], freeBlocks{start: start})
// 			if i-start > d.largestFreeSpace {
// 				fmt.Println("largest", i-start)
// 				d.largestFreeSpace = i - start
// 			}
// 		}
// 	}
// }
//
// func (d *disk) CompactByWholeFile() {
// 	d.ScanForFreeBlocks()
// 	fmt.Println(d.freeBlocks)
// 	i := advanceToNextFileBlock(d.blocks, len(d.blocks)-1)
// 	for i >= 0 {
// 		fmt.Println(d)
// 		fmt.Println(d.freeBlocks)
// 		char := d.blocks[i].data
// 		file := ""
// 		for i >= 0 && d.blocks[i].data != "." && d.blocks[i].data == char {
// 			file = d.blocks[i].data + file
// 			i--
// 		}
//
// 		fmt.Println("file", file)
//
// 		// find left most free blocks to swap file
// 		fileSize := len(file)
// 		startIdx := -1
// 		for fileSize <= d.largestFreeSpace {
// 			if val, ok := d.freeBlocks[fileSize]; ok {
// 				if val[0].start <= i {
// 					startIdx = val[0].start
// 					break
// 				}
// 			}
// 			fileSize++
// 		}
//
// 		fmt.Println("startIdx", startIdx)
//
// 		if startIdx != -1 && startIdx <= i {
//
// 			// clear out file blocks
// 			k := i + 1
// 			for k < i+1+len(file) && k < len(d.blocks) {
// 				d.blocks[k].data = "."
// 				k++
// 			}
//
// 			// move file to free space
// 			for j := 0; j < len(file); j++ {
// 				d.blocks[startIdx+j].data = string(file[j])
// 			}
// 		}
//
// 		i = advanceToNextFileBlock(d.blocks, i)
//
// 		d.ScanForFreeBlocks()
// 	}
// }

func (d *disk) CompactByWholeFile() {
	j := len(d.blocks) - 1
	for j >= 0 && d.blocks[j].data == "." {
		j--
	}
	i := 0
	for i < j && d.blocks[i].data != "." {
		i++
	}
	for j >= 0 {
		for i < j {
			if d.blocks[i].data != "." {
				i++
				continue
			}

			if d.blocks[i].size < d.blocks[j].size {
				i++
				continue
			}
			break
		}

		if i >= j {
			j--
			for j >= 0 && d.blocks[j].data == "." {
				j--
			}
			i = 0
			for i < j && d.blocks[i].data != "." {
				i++
			}
			continue
		}

		if d.blocks[i].size == d.blocks[j].size {
			d.blocks[i], d.blocks[j] = d.blocks[j], d.blocks[i]
		} else {
			// free space is larger than file
			if i == 0 {
				d.blocks = append([]block{d.blocks[j]}, d.blocks...)
			} else {
				d.blocks = slices.Insert(d.blocks, i, d.blocks[j])
				d.blocks[i+1].size = d.blocks[i+1].size - d.blocks[j+1].size
				d.blocks[j+1].data = "."
			}
		}
		j--
		for j >= 0 && d.blocks[j].data == "." {
			j--
		}
		i = 0
		for i < j && d.blocks[i].data != "." {
			i++
		}
	}
}

func (d disk) Checksum() int64 {
	fmt.Println(d)
	total := int64(0)
	elements := 0
	for i := 0; i < len(d.blocks); i++ {
		if d.blocks[i].data == "." {
			elements += d.blocks[i].size
			continue
		}
		id, err := strconv.Atoi(string(d.blocks[i].data))
		if err != nil {
			panic(err)
		}
		for j := 0; j < d.blocks[i].size; j++ {
			total += int64(id) * int64(elements)
			elements++
		}
	}
	fmt.Println("total", total)
	return total
}

// func part1() {
// 	d := disk{blocks: []block{}}
//
// 	// file is one line
// 	lineFunc := func(line string, _ int) error {
// 		id := 0
// 		for i, c := range line {
// 			blocks := []block{}
// 			size, err := strconv.Atoi(string(c))
// 			if err != nil {
// 				panic(err)
// 			}
// 			if i == 0 || i%2 == 0 {
// 				// file
// 				for j := 0; j < size; j++ {
// 					blocks = append(blocks, block{data: strconv.Itoa(id)})
// 				}
// 				id++
// 			} else {
// 				// free space
// 				for j := 0; j < size; j++ {
// 					blocks = append(blocks, block{data: "."})
// 				}
// 			}
// 			for _, b := range blocks {
// 				d.blocks = append(d.blocks, b)
// 			}
// 		}
// 		return nil
// 	}
//
// 	err := utils.ReadFile("input.txt", lineFunc)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	d.Compact()
// 	checksum := d.Checksum()
// 	fmt.Println(checksum)
// }

func part2() {
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
				blocks = append(blocks, block{size: size, data: strconv.Itoa(id)})
				id++
			} else {
				// free space
				blocks = append(blocks, block{size: size, data: "."})
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

	fmt.Println(d)
	d.CompactByWholeFile()
	fmt.Println(d)
	checksum := d.Checksum()
	fmt.Println(checksum)
}

func main() {
	// part1()
	part2()
}
