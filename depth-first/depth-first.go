package main

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func preWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	// pre
	path = append(path, curr.value)
	// recurse
	preWalk(curr.left, path)
	preWalk(curr.right, path)

	// post
	return path
}

func inWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	// pre
	preWalk(curr.left, path)

	// recurse
	path = append(path, curr.value)
	preWalk(curr.right, path)

	// post
	return path
}

func postWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}
	// pre
	preWalk(curr.left, path)

	// recurse
	preWalk(curr.right, path)

	// post
	path = append(path, curr.value)
	return path
}

func preOrderSearch(head *BinaryNode) []int {
	return preWalk(head, []int{})
}

func inOrderSearch(head *BinaryNode) []int {
	return inWalk(head, []int{})
}

func postOrderSearch(head *BinaryNode) []int {
	return postWalk(head, []int{})
}
