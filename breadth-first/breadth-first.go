package main

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func bfs(head *BinaryNode, needle int) bool {
	q := []*BinaryNode{head}

	for len(q) != 0 {
		var curr *BinaryNode
		curr, q = q[0], q[1:]

		// serch
		if curr.value == needle {
			return true
		}

		if curr.left != nil {
			q = append(q, curr.left)
		}

		if curr.left != nil {
			q = append(q, curr.right)
		}
	}

	return false
}
