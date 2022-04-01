package main

import "fmt"

// LFU in O(1)
func main() {
	capacity := 2
	cache := Constructor(capacity)
	cache.Put(1, 1)
	fmt.Println(cache)
}

type Node struct {
	next, prev *Node
	head, tail *Entry
	freq       int
}

type Entry struct {
	next, prev *Entry
	node       *Node
	value      int
}

type LFUCache struct {
	capacity   int
	size       int
	mapping    map[int]*Entry
	head, tail *Node
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		size:     0,
		mapping:  make(map[int]*Entry),
	}
}

func (this *LFUCache) Get(key int) int {
	entry, ok := this.mapping[key]
	if !ok {
		return -1
	}

	// if entry.node.head == entry {
	// 	entry.node.head = entry.next
	// 	entry.next.prev = nil
	// } else if entry.node.tail == entry {
	// 	entry.node.tail = entry.prev
	// 	entry.prev.next = nil
	// } else {
	// 	entry.prev.next = entry.next
	// 	entry.next.prev = entry.prev
	// }

	// if entry.node.next != nil {
	// 	currentFreq := entry.node.freq
	// 	nextFreq := entry.node.next.freq

	// 	if currentFreq+1 != nextFreq {
	// 		// new node
	// 	} else {
	// 		// append to next node tail
	// 	}
	// } else {
	// 	// new node
	// }

	// remove entry from node list
	// check freq of next node value
	// if increase from current and equally, push to the last item from that node
	// else create new node, append to between current and next node
	// if next node is nil, update the tail

	return entry.value
}

func (this *LFUCache) Put(key int, value int) {
	entry, ok := this.mapping[key]
	if ok {
		entry.value = value
		return
	}

	entry = &Entry{
		value: value,
	}
	this.mapping[key] = entry

	if this.head == nil {
		node := &Node{
			head: entry,
			tail: entry,
			freq: 1,
		}
		this.head = node
		this.tail = node
		this.size++
		return
	}

	if this.head.freq == 1 {
		node := this.head
		entry.prev = node.tail
		node.tail.next = entry
		node.tail = entry
	} else {

	}

	if this.size+1 < this.capacity {
		return
	}
	this.size++
	this.head.head

	// add item to first node where freq = 1
	// if freq != 1, add new node, update the head
	// if size reach capacity, remove item from first node
	// if first node is empty, remove the node and update the head

}
