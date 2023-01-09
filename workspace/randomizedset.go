package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randomizedSet := Constructor()
	randomizedSet.Insert(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	fmt.Println(randomizedSet)
	randomizedSet.Remove(2) // Returns false as 2 does not exist in the set.
	fmt.Println(randomizedSet)
	randomizedSet.Insert(2) // Inserts 2 to the set, returns true. Set now contains [1,2].
	fmt.Println(randomizedSet)
	randomizedSet.GetRandom() // getRandom() should return either 1 or 2 randomly.
	fmt.Println(randomizedSet)
	randomizedSet.Remove(1) // Removes 1 from the set, returns true. Set now contains [2].
	fmt.Println(randomizedSet)
	randomizedSet.Insert(2) // 2 was already in the set, so return false.
	fmt.Println(randomizedSet)
	randomizedSet.GetRandom() // Since 2 is the only number in the set, getRandom() will always return 2.
	fmt.Println(randomizedSet)
}

type RandomizedSet struct {
	mapping  map[int]int
	elements []int
	count    int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapping:  make(map[int]int),
		elements: make([]int, 0),
		count:    0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.mapping[val]
	if ok {
		return false
	}

	this.elements = append(this.elements, val)
	this.mapping[val] = this.count
	this.count++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.mapping[val]
	if !ok {
		return false
	}

	last := this.elements[this.count-1]
	this.mapping[last] = index
	this.elements[index] = last

	this.count--
	this.elements = this.elements[:this.count]
	delete(this.mapping, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	index := r.Intn(this.count)
	fmt.Println("return", this.elements[index], "from", this.count)
	return this.elements[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
