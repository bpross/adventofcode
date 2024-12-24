package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bpross/adventofcode/utils"
)

type Graph struct {
	vertices int
	edges    map[string][]string
}

// Create a new graph
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make(map[string][]string),
	}
}

// Add an edge to the graph
func (g *Graph) AddEdge(v, w string) {
	g.edges[v] = append(g.edges[v], w)
	g.edges[w] = append(g.edges[w], v)
}

// Function to find all cycles in the graph
func (g *Graph) FindAllCycles() [][]string {
	var cycles [][]string
	visited := make(map[string]bool)
	var stack []string

	var dfs func(v, parent string)
	dfs = func(v, parent string) {
		if len(stack) > 4 {
			return
		}

		visited[v] = true
		stack = append(stack, v)

		for _, i := range g.edges[v] {
			if !visited[i] {
				dfs(i, v)
			} else if i != parent && len(stack) >= 3 {
				// Cycle detected, extract the cycle path
				cycle := extractCycle(stack, i)
				cycles = append(cycles, cycle)
			}
		}

		stack = stack[:len(stack)-1]
		visited[v] = false
	}

	for v := range g.edges {
		if !visited[v] {
			dfs(v, "")
		}
	}

	return cycles
}

// Helper function to extract the cycle path from the stack
func extractCycle(stack []string, start string) []string {
	var cycle []string
	for i := len(stack) - 1; i >= 0; i-- {
		cycle = append(cycle, stack[i])
		if stack[i] == start {
			break
		}
	}
	return cycle
}

func part1() {
	graph := NewGraph(0)

	lineFunc := func(line string, _ int) error {
		l := strings.Split(line, "-")
		left, right := l[0], l[1]
		graph.AddEdge(left, right)
		return nil
	}

	err := utils.ReadFile("input.txt", lineFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println("finding cycles")
	cycles := graph.FindAllCycles()
	fmt.Println("done finding cycles")
	dedupe := make(map[string][]string)
	for _, cycle := range cycles {
		if len(cycle) == 3 {
			// sort the cycle
			sort.Slice(cycle, func(i, j int) bool {
				return cycle[i] < cycle[j]
			})
			dedupe[strings.Join(cycle, "-")] = cycle
		}
	}
	total := 0
	for _, cycle := range dedupe {
		for _, c := range cycle {
			if strings.HasPrefix(c, "t") {
				total += 1
				break
			}
		}
	}

	fmt.Println(total)
}

func main() {
	part1()
}
