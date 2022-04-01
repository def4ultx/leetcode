package main

import "fmt"

func main() {
	freqStack := Constructor()
	freqStack.Push(5) // The stack is [5]
	fmt.Println(freqStack)
	freqStack.Push(7) // The stack is [5,7]
	fmt.Println(freqStack)
	freqStack.Push(5) // The stack is [5,7,5]
	fmt.Println(freqStack)
	freqStack.Push(7) // The stack is [5,7,5,7]
	fmt.Println(freqStack)
	freqStack.Push(4) // The stack is [5,7,5,7,4]
	fmt.Println(freqStack)
	freqStack.Push(5) // The stack is [5,7,5,7,4,5]
	fmt.Println(freqStack)
	freqStack.Push(4) // The stack is [5,7,5,7,4,5,4]
	fmt.Println(freqStack)
	freqStack.Pop() // return 5, as 5 is the most frequent. The stack becomes [5,7,5,7,4].
	freqStack.Pop() // return 7, as 5 and 7 is the most frequent, but 7 is closest to the top. The stack becomes [5,7,5,4].
	freqStack.Pop() // return 5, as 5 is the most frequent. The stack becomes [5,7,4].
}

type FreqStack struct {
	maxCount int
	count    map[int]int
	elements map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		maxCount: 0,
		count:    make(map[int]int),
		elements: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	count, _ := this.count[val]

	this.count[val]++
	if count >= this.maxCount {
		this.maxCount++
	}

	maxCount := count + 1
	if len(this.elements[maxCount]) == 0 {
		this.elements[maxCount] = []int{val}
	} else {
		this.elements[maxCount] = append(this.elements[maxCount], val)
	}
}

func (this *FreqStack) Pop() int {
	var result int

	list := this.elements[this.maxCount]
	if len(list) > 1 {
		result = list[len(list)-1]
		this.elements[this.maxCount] = list[:len(list)-1]
	} else if len(list) == 1 {
		this.elements[this.maxCount] = []int{}
		this.maxCount--
		result = list[0]
	}

	this.count[result]--

	fmt.Println("result", result)
	return result
}
