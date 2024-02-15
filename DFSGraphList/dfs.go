package main

import "fmt"

type WeightedAdjacencyList [][][]int

func main() {
	list := WeightedAdjacencyList{
		{
			{2, 10},
		},
		{
			{0, 10},
			{2, 4},
			{3, 10},
		},
		{
			{0, 4},
			{1, 10},
		},
		{
			{},
		},
	}
	answer := dfs(list, 0, 3)
	fmt.Println("The answer is:", answer)
}

func walk(graph WeightedAdjacencyList, curr int, needle int, seen []bool, path *[]int) bool {
	if seen[curr] {
		return false
	}

	seen[curr] = true

	// recursion
	// pre
	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	// recurse
	list := graph[curr]
	for i := 0; i < len(list); i++ {
		edge := list[i]
		if walk(graph, edge[0], needle, seen, path) {
			return true
		}
	}

	// post
	_, *path = (*path)[len(*path)-1], (*path)[:len(*path)-1]

	return false
}

func dfs(graph WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	walk(graph, source, needle, seen, &path)

	return path
}
