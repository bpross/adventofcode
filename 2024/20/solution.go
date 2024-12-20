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
)

type position struct {
	x, y int
}

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

func findAllWallRemovals(grid [][]string) map[position]struct{} {
	// find all walls where the previous and next are not walls
	// and the current is a wall
	positions := make(map[position]struct{}, 0)
	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[0])-1; j++ {
			// first look horizontally
			if j-1 < 0 || j+1 >= len(grid[0]) {
				continue
			}
			if grid[i][j] == wall && grid[i][j-1] != wall && grid[i][j+1] != wall {
				positions[position{i, j}] = struct{}{}
			}
			// then look vertically
			if i-1 < 0 || i+1 >= len(grid) {
				continue
			}
			if grid[i][j] == wall && grid[i-1][j] != wall && grid[i+1][j] != wall {
				positions[position{i, j}] = struct{}{}
			}
		}
	}
	return positions
}

func dijkstra(grid [][]string, start, end position, cheat *position) int {
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
				if cheat != nil && next != *cheat {
					continue
				} else if cheat == nil {
					continue
				}
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

	wallRemovals := findAllWallRemovals(grid)

	initialCost := dijkstra(grid, start, end, nil)
	fmt.Println(initialCost)

	secondsSaved := make(map[int][]position)
	idx := 0
	for p := range wallRemovals {
		cost := dijkstra(grid, start, end, &p)
		saved := initialCost - cost
		secondsSaved[saved] = append(secondsSaved[saved], p)
		idx++
	}

	total := 0
	for saved, positions := range secondsSaved {
		if saved >= 100 {
			total += len(positions)
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
}
