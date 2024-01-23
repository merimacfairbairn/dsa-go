package main

type Point struct {
	x int
	y int
}

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func mazeSolver(maze []string, wall rune, start Point, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[0]))
	}
	path := make([]Point, 0)
	pPath := &path

	walk(maze, wall, start, end, seen, pPath)

	return path
}

func walk(maze []string, wall rune, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// 1. Base Case
	// Off the map
	if curr.x < 0 || curr.x >= len(maze[0]) || curr.y < 0 || curr.y >= len(maze) {
		return false
	}
	// on a wall
	if maze[curr.y][curr.x] == byte(wall) {
		return false
	}
	// got to the end
	if curr.x == end.x && curr.y == end.y {
		*path = append(*path, end)
		return true
	}
	// seen it
	if seen[curr.y][curr.x] {
		return false
	}

	// 3 recurse
	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// recurse
	for i := range dir {
		x, y := dir[i][0], dir[i][1]
		next := Point{x: curr.x + x, y: curr.y + y}
		if walk(maze, wall, next, end, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}
