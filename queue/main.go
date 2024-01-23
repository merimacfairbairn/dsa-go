package main

import (
	"errors"
	"fmt"
)

type Node struct {
    value   int
    next    *Node
}

type Queue struct {
    length  int
    head    *Node
    tail    *Node
}

func main() {
    q := constructor()
    q.enqueue(10)
    val, err := q.peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
    q.enqueue(11)
    val, err = q.peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
    dequed, err := q.deque()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Value dequed:", dequed)
    val, err = q.peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func constructor() *Queue {
    q := new(Queue)
    q.length = 0
    q.head, q.tail = nil, nil
    return q
}

func (q *Queue) enqueue(item int) {
    q.length++
    node := new(Node)
    node.value = item
    node.next = nil
    if q.tail == nil {
        q.head, q.tail = node, node
        return
    }
    q.tail.next = node
    q.tail = node
}

func (q *Queue) deque() (int, error) {
    if q.head == nil {
        return 0, errors.New("Queue is empty")
    }
    q.length--
    head := q.head
    q.head = q.head.next

    head.next = nil
    return head.value, nil
}

func (q *Queue) peek() (int, error) {
    if q.head == nil {
        return 0, errors.New("Queue is empty")
    }
    return q.head.value, nil
}
