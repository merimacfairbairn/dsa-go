package main

import (
	"fmt"
	"math"
	"slices"
)

type GraphEdge map[string]int
type WeightedAdjacencyList [][]GraphEdge

func main() {
	graph := WeightedAdjacencyList{
		{
			{
				"to":     3,
				"weight": 38,
			},
			{
				"to":     2,
				"weight": 2,
			},
		},
		{
			{
				"to":     3,
				"weight": 10,
			},
		},
		{
			{
				"to":     1,
				"weight": 8,
			},
		},
		{
			{
				"to":     1,
				"weight": 1,
			},
		},
	}
	answer := dijkstra(0, 1, graph)
	fmt.Println("The answer is:", answer)
}

func dijkstra(source int, sink int, graph WeightedAdjacencyList) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}
	dists := make([]float64, len(graph))
	for i := range dists {
		dists[i] = math.Inf(0)
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge["to"]] {
				continue
			}

			dist := dists[curr] + float64(edge["weight"])
			if dist < dists[edge["to"]] {
				dists[edge["to"]] = float64(dist)
				prev[edge["to"]] = curr
			}
		}
	}

	var out []int
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}

func hasUnvisited(seen []bool, dists []float64) bool {
	has := false
	pHas := &has

	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < math.Inf(0) {
			*pHas = true
			break
		}
	}

	return has
}

func getLowestUnvisited(seen []bool, dists []float64) int {
	idx := -1
	pIdx := &idx
	lowestDistance := math.Inf(0)

	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}

		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			*pIdx = i
		}
	}

	return idx
}
