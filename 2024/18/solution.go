package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type position struct {
	x, y int
}

// Pulled all of this from: the golang docs.
// Just swaped the Less method to return the lowest value
// instead of the highest

// An Item is something we manage in a priority queue.
type Item struct {
	p        position
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (i Item) String() string {
	return fmt.Sprintf("%v: %d", i.p, i.priority)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on priority
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func dijkstra(grid [][]string, start, end position) int {
	wall := "#"
	queue := make(PriorityQueue, 1)
	queue[0] = &Item{start, 0, 0}
	heap.Init(&queue)
	cameFrom := make(map[position]position)
	costSoFar := make(map[position]int)
	cameFrom[start] = position{-1, -1}
	costSoFar[start] = 0

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*Item)

		if current.p == end {
			break
		}

		for _, next := range []position{
			{current.p.x, current.p.y - 1},
			{current.p.x, current.p.y + 1},
			{current.p.x - 1, current.p.y},
			{current.p.x + 1, current.p.y},
		} {
			if next.x < 0 || next.x >= len(grid) || next.y < 0 || next.y >= len(grid[0]) {
				continue
			}
			if grid[next.x][next.y] == wall {
				continue
			}
			newCost := costSoFar[current.p] + 1
			if _, ok := costSoFar[next]; !ok || newCost < costSoFar[next] {
				costSoFar[next] = newCost
				priority := newCost + 0 // heuristic
				heap.Push(&queue, &Item{next, priority, 0})
				cameFrom[next] = current.p
			}
		}
	}

	return costSoFar[end]
}

func part1() {
	grid := make([][]string, 0)

	size := 71
	for i := 0; i < size; i++ {
		row := make([]string, size)
		for j := 0; j < size; j++ {
			row[j] = "."
		}
		grid = append(grid, row)
	}

	lineFunc := func(line string, _ int) error {
		cords := strings.Split(line, ",")
		c, err := strconv.Atoi(cords[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(cords[1])
		if err != nil {
			panic(err)
		}

		grid[r][c] = "#"
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}
	for _, row := range grid {
		fmt.Println(strings.Join(row, ""))
	}

	start := position{0, 0}
	end := position{size - 1, size - 1}
	fmt.Println(start, end)
	fmt.Println(dijkstra(grid, start, end))
}

func main() {
	part1()
}
