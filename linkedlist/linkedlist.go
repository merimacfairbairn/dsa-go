package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type List struct {
	length int
	head   *Node
	tail   *Node
}

func main() {
	l := constructor()
	l.prepend(12)
	l.append(420)
	fmt.Println(l.get(1))
	l.prepend(29)
	fmt.Println(l.get(1))
	l.remove(12)
	fmt.Println(l.get(1))
	l.insertAt(69, 1)
	fmt.Println(l.get(1))
	l.removeAt(1)
	fmt.Println(l.get(1))
}

func constructor() *List {
	return &List{0, nil, nil}
}

func (this *List) prepend(item int) {
	node := &Node{value: item}
	this.length++
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}
	node.next = this.head
	this.head.prev = node
	this.head = node
}

func (this *List) append(item int) {
	this.length++
	node := &Node{value: item}

	if this.tail == nil {
		this.head = node
		this.tail = node
		return
	}

	node.prev = this.tail
	this.tail.next = node
	this.tail = node
}

func (this *List) insertAt(item int, idx int) error {
	if idx > this.length {
		return errors.New("Oh no")
	} else if idx == this.length {
		this.append(item)
		return nil
	} else if idx == 0 {
		this.prepend(item)
		return nil
	}
	this.length++
	curr := this.getAt(idx)

	node := &Node{value: item}

	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if node.prev != nil {
		node.prev.next = node
	}

	return nil
}

func (this *List) remove(item int) (int, error) {
	curr := this.head
	for i := 0; curr != nil && i < this.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}
	if curr == nil {
		return -1, errors.New("Oh no")
	}

	return this.removeNode(curr)
}

func (this *List) get(idx int) (int, error) {
	node := this.getAt(idx)
	if node == nil {
		return -1, errors.New("Oh no")
	}

	return node.value, nil
}

func (this *List) removeAt(idx int) (int, error) {
	node := this.getAt(idx)

	if node == nil {
		return -1, errors.New("Oh no")
	}

	return this.removeNode(node)
}

func (this *List) removeNode(node *Node) (int, error) {
	this.length--

	if this.length == 0 {
		out := this.head.value
		this.head = nil
		return out, nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == this.head {
		this.head = node.next
	}

	if node == this.tail {
		this.tail = node.prev
	}

	node.prev = nil
	node.next = nil

	return node.value, nil
}

func (this *List) getAt(idx int) *Node {
	curr := this.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}
	return curr
}
