package main

import (
	"fmt"
	"math"
)

func main() {
	stack := Constructor()
	// stack.Push(-10)
	// stack.Push(14)
	// stack.GetMin()
	// stack.GetMin()
	// stack.Push(-20)
	// stack.GetMin()
	// stack.GetMin()
	// stack.Top()
	// stack.GetMin()
	// stack.Pop()
	// stack.Push(10)
	// stack.Push(-7)
	// stack.GetMin()

	// stack.Push(-7)
	// stack.Pop()
	// stack.Top()
	// stack.GetMin()
	// stack.Pop()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.GetMin()
	stack.Pop()
	stack.Top()
	stack.GetMin()

	fmt.Println(stack)
}

type Item struct {
	val int
	min int
}

type MinStack struct {
	stack []Item
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]Item, 0),
	}
}

func (this *MinStack) Push(val int) {

	min := math.MaxInt64
	if len(this.stack) >= 1 {
		prev := this.stack[len(this.stack)-1]
		fmt.Println("prev", prev, val)
		min = prev.min
	}
	fmt.Println("val", val, min)
	item := Item{
		val: val,
		min: Min(min, val),
	}
	this.stack = append(this.stack, item)
}

func (this *MinStack) Pop() {
	// if len(this.stack) == 1 {
	// 	this.min = math.MaxInt64
	// }
	// if len(this.stack) > 1 {
	// 	item := this.stack[len(this.stack)-1]
	// 	prev := this.stack[len(this.stack)-2]

	// 	if prev.minSoFar > item.minSoFar {
	// 		this.min = prev.minSoFar
	// 	}
	// }
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
