package main

import (
	"errors"
	"fmt"
)

type number interface {
	int | int64 | float64
}

type Node[T number] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type LRU[K number | string, V number] struct {
	length        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]Node[V]
	reverseLookup map[Node[V]]K
}

func main() {
	lru := constructor[string, int](3)
	lru.update("foo", 10)
	v, _ := lru.get("foo")
	fmt.Println(v, lru.length)
	lru.update("bar", 14)
	v, _ = lru.get("bar")
	fmt.Println(v, lru.length)
	lru.update("baz", 88)
	v, _ = lru.get("baz")
	fmt.Println(v, lru.length)
	lru.update("ball", 69)
	v, _ = lru.get("ball")
	fmt.Println(v, lru.length)
	v, _ = lru.get("foo")
	fmt.Println(v, lru.length)
}

func createNode[T number](value T) *Node[T] {
	return &Node[T]{value: value}
}

func constructor[K number | string, V number](capacity ...int) *LRU[K, V] {
	this := new(LRU[K, V])
	if len(capacity) == 0 {
		capacity[0] = 10
	}
	this.capacity = capacity[0]
	this.length = 0
	this.head = nil
	this.tail = nil
	this.lookup = make(map[K]Node[V], capacity[0])
	this.reverseLookup = make(map[Node[V]]K, capacity[0])
	return this
}

func (this *LRU[K, V]) update(key K, value V) {
	node, ok := this.lookup[key]
	if !ok {
		newNode := createNode(value)
		this.length++
		this.prepend(newNode)
		this.trimCache()

		this.lookup[key] = *newNode
		this.reverseLookup[*newNode] = key
	} else {
		this.detach(&node)
		this.prepend(&node)
		p := &node
		p.value = value
	}
}

func (this *LRU[K, V]) get(key K) (V, error) {
	node, ok := this.lookup[key]
	pNode := &node
	if !ok {
		var result V
		return result, errors.New("The value for the given key does not exist")
	}

	this.detach(pNode)
	this.prepend(pNode)

	return node.value, nil
}

func (this *LRU[K, V]) detach(node *Node[V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if this.length == 1 {
		this.head = nil
		this.tail = nil
	}
	if this.head == node {
		this.head = this.head.next
	}
	if this.tail == node {
		this.tail = this.tail.prev
	}

	node.next = nil
	node.prev = nil
}

func (this *LRU[K, V]) prepend(node *Node[V]) {
	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	node.next = this.head
	this.head.prev = node
	this.head = node
}

func (this *LRU[K, V]) trimCache() {
	if this.length <= this.capacity {
		return
	}

	tail := this.tail
	this.detach(tail)

	key := this.reverseLookup[*tail]
	delete(this.lookup, key)
	delete(this.reverseLookup, *tail)
	this.length--
}
