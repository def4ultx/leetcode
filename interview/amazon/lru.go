package main

import (
	"container/list"
)

func main() {
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], 		[1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Output
	// [null, null, null, 1, null, -1, null, -1, 3, 4]

	// Input
	// ["LRUCache","put","put","get","put","put","get"]
	// [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
	// Output
	// [null,null,null,1,null,null,-1]
	// Expected
	// [null,null,null,2,null,null,-1]

	// lru := Constructor(2)
	// lru.Put(2, 1)
	// lru.Put(2, 2)
	// lru.Get(2)

	// lru.Put(1, 1)
	// lru.Put(2, 2)

	// lru.Put(1, 1)
	// lru.Put(2, 2)
	// lru.Get(1)

	// fmt.Println("lru", lru)
	// lru.Put(3, 3)
	// fmt.Println("lru", lru)
	// PrintList(lru.data)
	// lru.Get(2)

	// ll := &LinkedList{}

	// n1 := &Node{1, nil, nil}
	// n2 := &Node{2, nil, nil}
	// n3 := &Node{3, nil, nil}
	// n4 := &Node{4, nil, nil}
	// n5 := &Node{5, nil, nil}

	// ll.AddLast(n1)
	// ll.AddLast(n2)
	// ll.AddLast(n3)
	// ll.AddLast(n4)
	// ll.AddLast(n5)

	// PrintList(ll)
	// ll.RemoveHead()
	// PrintList(ll)
	// ll.RemoveHead()
	// ll.RemoveHead()
	// ll.RemoveHead()
	// PrintList(ll)
	// ll.RemoveHead()
	// PrintList(ll)
	// ll.RemoveHead()
	// ll.RemoveHead()
	// PrintList(ll)

	// ll.Remove(n5)
	// fmt.Println("last", ll.last)
	// ll.Remove(n4)
	// fmt.Println("last", ll.last)
	// ll.Remove(n3)
	// fmt.Println("last", ll.last)
	// ll.Remove(n2)
	// fmt.Println("last", ll.last)
	// ll.Remove(n1)
	// fmt.Println("last", ll.last)

	// ll.Remove(n3)

	// curr = ll.last
	// for curr != nil {
	// 	fmt.Print(curr.val, " ")
	// 	curr = curr.prev
	// }

	// fmt.Println()
	// ll.AddLast(n3)
	// curr = ll.last
	// for curr != nil {
	// 	fmt.Print(curr.val, " ")
	// 	curr = curr.prev
	// }

	// fmt.Println("--")
	// ll.Remove(n1)
	// ll.AddLast(n1)
	// curr = ll.last
	// for curr != nil {
	// 	fmt.Print(curr.val, " ")
	// 	curr = curr.prev
	// }
	// fmt.Println("--")
	// ll.Remove(n5)
	// ll.AddLast(n5)
	// curr = ll.last
	// for curr != nil {
	// 	fmt.Print(curr.val, " ")
	// 	curr = curr.prev
	// }
	// fmt.Println("--")
	// ll.Remove(n2)
	// ll.AddLast(n2)
	// curr = ll.last
	// for curr != nil {
	// 	fmt.Print(curr.val, " ")
	// 	curr = curr.prev
	// }

}

// func PrintList(ll *LinkedList) {
// 	curr := ll.head
// 	for curr != nil {
// 		fmt.Print(curr.val, " ")
// 		curr = curr.next
// 	}
// 	fmt.Println()
// }

// type LinkedList struct {
// 	head, last *Node
// }

// func (l *LinkedList) Remove(node *Node) {
// 	// fmt.Println("remove", node, node.prev, node.next)
// 	if l.last == node {
// 		l.last = node.prev
// 	}
// 	if l.head == node {
// 		l.head = node.next
// 	}
// 	if node.prev != nil {
// 		node.prev.next = node.next
// 	}
// 	if node.next != nil {
// 		node.next.prev = node.prev
// 	}
// 	node.next = nil
// 	node.prev = nil
// }

// func (l *LinkedList) AddLast(node *Node) {
// 	if l.last != nil {
// 		l.last.next = node
// 		node.prev = l.last
// 	}
// 	if l.head == nil {
// 		l.head = node
// 	}
// 	l.last = node
// }

// func (l *LinkedList) RemoveHead() *Node {
// 	curr := l.head
// 	if l.head != nil {
// 		l.head = l.head.next
// 	}
// 	if l.head != nil {
// 		l.head.prev = nil
// 	}
// 	return curr
// }

// type Node struct {
// 	key        int
// 	val        int
// 	next, prev *Node
// }

// type LRUCache struct {
// 	capacity int
// 	size     int

// 	mapping map[int]*Node
// 	data    *LinkedList
// }

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{
// 		capacity: capacity,
// 		size:     0,
// 		mapping:  make(map[int]*Node),
// 		data:     &LinkedList{},
// 	}
// }

// func (this *LRUCache) Get(key int) int {
// 	node, ok := this.mapping[key]
// 	if !ok {
// 		fmt.Println("return -1")
// 		return -1
// 	}

// 	this.data.Remove(node)
// 	this.data.AddLast(node)

// 	fmt.Println("return node", node)
// 	return node.val
// }

// func (this *LRUCache) Put(key int, value int) {
// 	node, ok := this.mapping[key]
// 	if !ok {
// 		n := &Node{key, value, nil, nil}

// 		this.size++
// 		this.mapping[key] = n
// 		this.data.AddLast(n)
// 	} else {
// 		node.val = value
// 		this.data.Remove(node)
// 		this.data.AddLast(node)
// 	}

// 	if this.size > this.capacity {
// 		head := this.data.RemoveHead()
// 		this.size--
// 		delete(this.mapping, head.key)
// 	}
// }

type LRUCache struct {
	capacity int
	size     int

	mapping map[int]*list.Element
	data    *list.List
}

type Item struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		mapping:  make(map[int]*list.Element),
		data:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.mapping[key]
	if !ok {
		return -1
	}
	this.data.MoveToBack(elem)
	return elem.Value.(*Item).value
}

func (this *LRUCache) Put(key int, value int) {
	elem, ok := this.mapping[key]
	if !ok {
		item := &Item{key, value}
		elem := this.data.PushBack(item)

		this.size++
		this.mapping[key] = elem
	} else {
		elem.Value.(*Item).value = value
		this.data.MoveToBack(elem)
	}

	if this.size > this.capacity {
		head := this.data.Front()
		this.size--
		this.data.Remove(head)
		delete(this.mapping, head.Value.(*Item).key)
	}
}
