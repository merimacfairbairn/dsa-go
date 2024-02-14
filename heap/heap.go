package main

import "fmt"

type MinHeap struct {
	length int
	data   []int
}

func main() {
	heap := &MinHeap{0, make([]int, 2)}
	heap.insert(12)
	fmt.Println(heap.data)
	heap.insert(83)
	fmt.Println(heap.data)
	deleted := heap.delete()
	fmt.Println("Deleted:", deleted)
}

func (this *MinHeap) insert(value int) {
	this.data[this.length] = value
	this.heapifyUp(this.length)
	this.length++
}

func (this *MinHeap) delete() int {
	if this.length == 0 {
		return -1
	}

	out := this.data[0]
	this.length--

	if this.length == 0 {
		this.data = []int{}
		return out
	}

	this.data[0] = this.data[this.length]
	this.heapifyDown(0)

	return out
}

func (this *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := this.parent(idx)
	parentV := this.data[p]
	v := this.data[idx]

	if parentV > v {
		this.data[idx] = parentV
		this.data[p] = v
		this.heapifyUp(p)
	}
}

func (this *MinHeap) heapifyDown(idx int) {
	lIdx := this.leftChild(idx)
	rIdx := this.rightChild(idx)

	if lIdx >= this.length || rIdx >= this.length {
		return
	}

	lV := this.data[lIdx]
	rV := this.data[rIdx]
	v := this.data[idx]

	if lV > rV && v > rV {
		this.data[idx] = rV
		this.data[rIdx] = v
		this.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		this.data[idx] = lV
		this.data[lIdx] = v
		this.heapifyDown(lIdx)
	}
}

func (this *MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (this *MinHeap) leftChild(idx int) int {
	return idx*2 + 1
}

func (this *MinHeap) rightChild(idx int) int {
	return idx*2 + 2
}
