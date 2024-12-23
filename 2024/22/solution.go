package main

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/bpross/adventofcode/utils"
)

// An Item is something we manage in a priority queue.
type Item struct {
	sequence string
	priority int64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (i Item) String() string {
	return fmt.Sprintf("%v: %d", i.sequence, i.priority)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on priority
	return pq[i].priority > pq[j].priority
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
func (pq *PriorityQueue) update(item *Item, priority int64) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func sequenceToString(seq []int64) string {
	s := ""
	for _, n := range seq {
		s += strconv.Itoa(int(n))
	}
	return s
}

func mix(n, p int64) int64 {
	return n ^ p
}

func prune(n int64) int64 {
	return n % 16777216
}

func calculateNextSecretNumber(n int64) int64 {
	val := n

	// first multiply by 64
	val *= 64
	val = mix(val, n)
	val = prune(val)

	// divide by 32
	prev := val
	val /= 32
	val = mix(val, prev)
	val = prune(val)

	prev = val
	val *= 2048
	val = mix(val, prev)
	val = prune(val)

	return val
}

func part1() {
	initialNumbers := []int64{}
	lineFunc := func(line string, _ int) error {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		initialNumbers = append(initialNumbers, int64(i))

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	secretNumbers := make([]int64, len(initialNumbers))
	copy(secretNumbers, initialNumbers)

	for i := 0; i < 2000; i++ {
		for j := 0; j < len(initialNumbers); j++ {
			secretNumbers[j] = calculateNextSecretNumber(secretNumbers[j])
		}
	}
	total := int64(0)
	for _, n := range secretNumbers {
		total += n
	}
	fmt.Println(total)
}

func part2() {
	initialNumbers := []int64{}
	lineFunc := func(line string, _ int) error {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		initialNumbers = append(initialNumbers, int64(i))

		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	secretNumbers := make([]int64, len(initialNumbers))
	copy(secretNumbers, initialNumbers)

	prices := make([][]int64, len(initialNumbers))
	priceChanges := make([][]int64, len(initialNumbers))

	for i := 0; i < 2000; i++ {
		for j := 0; j < len(initialNumbers); j++ {
			prev := secretNumbers[j]
			secretNumbers[j] = calculateNextSecretNumber(secretNumbers[j])
			price := secretNumbers[j] % 10
			prices[j] = append(prices[j], price)
			if i == 0 {
				priceChanges[j] = append(priceChanges[j], prices[j][i]-prev%10)
				continue
			}
			priceChanges[j] = append(priceChanges[j], prices[j][i]-prices[j][i-1])
		}
	}

	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	sequenceToItem := make(map[string]*Item)

	for i := 0; i < len(initialNumbers); i++ {
		q := []int64{}
		seen := make(map[string]bool)
		for j := 0; j < len(priceChanges[i]); j++ {
			if len(q) < 4 {
				q = append(q, priceChanges[i][j])
				continue
			}
			seq := sequenceToString(q)
			if _, ok := seen[seq]; ok {
				q = q[1:]
				q = append(q, priceChanges[i][j])
				continue
			}
			seen[seq] = true
			if _, ok := sequenceToItem[seq]; !ok {
				item := &Item{sequence: seq, priority: prices[i][j-1]}
				sequenceToItem[seq] = item
				heap.Push(&queue, item)
			} else {
				item := sequenceToItem[seq]
				queue.update(item, item.priority+prices[i][j-1])
			}
			q = q[1:]
			q = append(q, priceChanges[i][j])
		}
	}

	item := heap.Pop(&queue).(*Item)
	fmt.Println(item)
}

func main() {
	// part1()
	part2()
}
