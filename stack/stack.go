package main

import (
	"errors"
	"fmt"
	"math"
)

type Number interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

type Node[T Number] struct {
	value T
	prev  *Node[T]
}

type Stack[T Number] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func main() {
	s := constructor[int8]()
	s.push(13)
	fmt.Println(s.peek())
	s.push(1)
	fmt.Println(s.peek())
	s.push(10)
	fmt.Println(s.peek())
	val := s.pop()
	fmt.Println(val)
}

func constructor[T Number]() *Stack[T] {
	s := new(Stack[T])
	s.length = 0
	s.head = nil
	s.tail = nil
	return s
}

func (s *Stack[T]) push(item T) {
	node := &Node[T]{value: item, prev: nil}
	s.length++
	if s.head == nil {
		s.head = node
		return
	}

	node.prev = s.head
	s.head = node
}

func (s *Stack[T]) pop() T {
	s.length = int(math.Max(0, float64(s.length)-1))
	head := s.head
	if s.length == 0 {
		s.head = nil
		return head.value
	}

	s.head = s.head.prev
	return head.value
}

func (s *Stack[T]) peek() (T, error) {
	if s.head == nil {
		return 0, errors.New("Stack is empty")
	}
	return s.head.value, nil
}
