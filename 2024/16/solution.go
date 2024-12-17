package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

const (
	startChar = "S"
	endChar   = "E"
	wall      = "#"

	up    = "^"
	down  = "v"
	left  = "<"
	right = ">"
)

var directionsToCost = map[string]map[string]int{
	up:    {left: 1001, up: 1, right: 1001, down: 2001},
	down:  {left: 1001, up: 2001, right: 1001, down: 1},
	left:  {left: 1, up: 1001, right: 2001, down: 1001},
	right: {left: 2001, up: 1001, right: 1, down: 1001},
}

type position struct {
	x, y int
}

type move struct {
	p         position
	direction string
}

// Pulled all of this from: the golang docs.
// Just swaped the Less method to return the lowest value
// instead of the highest

// An Item is something we manage in a priority queue.
type Item struct {
	m        move
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (i Item) String() string {
	return fmt.Sprintf("%v: %d", i.m, i.priority)
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

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, m move, priority int) {
	item.m = m
	item.priority = priority
	heap.Fix(pq, item.index)
}

func dijkstra(grid [][]string, start, end position) int {
	queue := make(PriorityQueue, 1)
	queue[0] = &Item{move{direction: right, p: start}, 0, 0}
	heap.Init(&queue)
	cameFrom := make(map[position]position)
	costSoFar := make(map[position]int)
	cameFrom[start] = position{-1, -1}
	costSoFar[start] = 0

	for queue.Len() > 0 {
		current := heap.Pop(&queue).(*Item)

		if current.m.p == end {
			break
		}

		for _, next := range []move{
			{position{current.m.p.x, current.m.p.y - 1}, left},
			{position{current.m.p.x, current.m.p.y + 1}, right},
			{position{current.m.p.x - 1, current.m.p.y}, up},
			{position{current.m.p.x + 1, current.m.p.y}, down},
		} {
			if next.p.x < 0 || next.p.x >= len(grid) || next.p.y < 0 || next.p.y >= len(grid[0]) {
				continue
			}
			if grid[next.p.x][next.p.y] == wall {
				continue
			}
			newCost := costSoFar[current.m.p] + directionsToCost[current.m.direction][next.direction]
			if _, ok := costSoFar[next.p]; !ok || newCost < costSoFar[next.p] {
				costSoFar[next.p] = newCost
				priority := newCost + 0 // heuristic
				heap.Push(&queue, &Item{next, priority, 0})
				cameFrom[next.p] = current.m.p
			}
		}
	}

	return costSoFar[end]
}

func part1() {
	grid := make([][]string, 0)
	var start position
	var end position

	lineFunc := func(line string, r int) error {
		row := strings.Split(line, "")
		for i, c := range row {
			if c == startChar {
				start = position{r, i}
			}
			if c == endChar {
				end = position{r, i}
			}
		}
		grid = append(grid, row)
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println(start, end)
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println()
	fmt.Println(dijkstra(grid, start, end))
}

func main() {
	part1()
}
