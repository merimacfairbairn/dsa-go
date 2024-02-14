package main

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func search(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.value == needle {
		return true
	}

	if curr.value < needle {
		return search(curr.left, needle)
	}

	return search(curr.right, needle)
}

func dfs(head *BinaryNode, needle int) bool {
	return search(head, needle)
}

func main() {
}
