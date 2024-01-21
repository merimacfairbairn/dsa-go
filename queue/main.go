package main

import "fmt"

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
    fmt.Println(q.peek())
    q.enqueue(11)
    fmt.Println(q.peek())
    dequed := q.deque()
    fmt.Println("Value dequed:", dequed)
    fmt.Println(q.peek())
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

func (q *Queue) deque() int {
    if q.head == nil {
        return q.head.value
    }

    q.length--
    head := q.head
    q.head = q.head.next

    head.next = nil
    return head.value
}

func (q *Queue) peek() int {
    return q.head.value
}
