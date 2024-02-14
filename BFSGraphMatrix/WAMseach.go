package main

import (
	"fmt"
	"slices"
)

type WeightedAdjacencyMatrix [][]int

func main() {
	matrix := WeightedAdjacencyMatrix{
		//       0  1  2  3
		/* 0 */ {0, 2, 0, 10},
		/* 1 */ {10, 0, 0, 10},
		/* 2 */ {12, 12, 0, 10},
		/* 3 */ {12, 12, 10, 0},
	}
	answer := bfs(matrix, 0, 2)
	fmt.Printf("The answer for matrix %v is: %v\n", matrix, answer)
}

func bfs(graph WeightedAdjacencyMatrix, source int, needle int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	q := []int{source}

	i := 0
	for {
		curr := 0
		curr, q = q[0], q[1:]
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 || seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
		seen[curr] = true

		i++
		if i == len(q) {
			break
		}
	}

	// build it backwards

	if prev[needle] == -1 {
		return []int{}
	}

	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	slices.Reverse(out)
	return out
}
