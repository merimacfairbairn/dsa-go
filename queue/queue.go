package main

import (
	"errors"
	"fmt"
)

type Number interface{ int8 | int16 | int32 | int64 | float32 | float64 }

type Node[T Number] struct {
    value   T
    next    *Node[T]
}

type Queue[T Number] struct {
    length  int
    head    *Node[T]
    tail    *Node[T]
}

func main() {
    q := constructor[int8]()
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

func constructor[T Number]() *Queue[T] {
    q := new(Queue[T])
    q.length = 0
    q.head, q.tail = nil, nil
    return q
}

func (q *Queue[T]) enqueue(item T) {
    q.length++
    node := new(Node[T])
    node.value = item
    node.next = nil
    if q.tail == nil {
        q.head, q.tail = node, node
        return
    }
    q.tail.next = node
    q.tail = node
}

func (q *Queue[T]) deque() (T, error) {
    if q.head == nil {
        return 0, errors.New("Queue[T] is empty")
    }
    q.length--
    head := q.head
    q.head = q.head.next

    head.next = nil
    return head.value, nil
}

func (q *Queue[T]) peek() (T, error) {
    if q.head == nil {
        return 0, errors.New("Queue[T] is empty")
    }
    return q.head.value, nil
}
