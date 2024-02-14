package main

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func compare(a *BinaryNode, b *BinaryNode) bool {
	// structural check
	if a == nil && b == nil {
		return true
	}

	// structural check
	if a == nil || b == nil {
		return false
	}

	// value check
	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}

func main() {
}
