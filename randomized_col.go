package main

import (
	"math/rand"
	"time"
)

func main() {
	col := Constructor()

	// col.Insert(1)
	// fmt.Println(col)
	// col.Insert(1)
	// fmt.Println(col)
	// col.Insert(2)
	// fmt.Println(col)
	// col.Insert(2)
	// fmt.Println(col)
	// col.Insert(2)
	// fmt.Println(col)
	// col.Remove(1)
	// fmt.Println(col)
	// col.Remove(1)
	// fmt.Println(col)
	// col.Remove(2)
	// fmt.Println(col)
	// col.Insert(1)
	// fmt.Println(col)
	// col.Remove(2)
	// fmt.Println(col)
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.GetRandom()
	// col.Insert(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	// fmt.Println(col)
	// col.Remove(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	// fmt.Println(col)
	// col.Insert(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	// fmt.Println(col)

	// col.Insert(1) // Inserts 1 to the set. Returns true as 1 was inserted successfully.
	// fmt.Println(col)
	// col.Remove(2) // Returns false as 2 does not exist in the set.
	// fmt.Println(col)
	// col.Insert(2) // Inserts 2 to the set, returns true. Set now contains [1,2].
	// fmt.Println(col)
	// col.Insert(2) // Inserts 2 to the set, returns true. Set now contains [1,2].
	// fmt.Println(col)
	// col.GetRandom() // getRandom() should return either 1 or 2 randomly.
	// fmt.Println(col)
	// col.Remove(1) // Removes 1 from the set, returns true. Set now contains [2].
	// fmt.Println(col)
	// col.Insert(2) // 2 was already in the set, so return false.
	// fmt.Println(col)
	// col.GetRandom() // Since 2 is the only number in the set, getRandom() will always return 2.
	// fmt.Println(col)
	// col.Insert(1) // 2 was already in the set, so return false.
	// fmt.Println(col)
	// col.GetRandom() // Since 2 is the only number in the set, getRandom() will always return 2.
	// fmt.Println(col)
}

type RandomizedCollection struct {
	mapping  map[int]map[int]struct{}
	elements []int
	count    int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		mapping:  make(map[int]map[int]struct{}),
		elements: make([]int, 0),
		count:    0,
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	_, found := this.mapping[val]
	if !found {
		this.mapping[val] = make(map[int]struct{})
	}

	this.mapping[val][this.count] = struct{}{}
	this.elements = append(this.elements, val)
	this.count++

	return !found
}

func (this *RandomizedCollection) Remove(val int) bool {
	indexes, ok := this.mapping[val]
	if !ok {
		return false
	}

	var index int
	for k := range indexes {
		index = k
		break
	}

	// swap index with last elem
	lastElemIndex := this.count - 1
	last := this.elements[lastElemIndex]

	delete(this.mapping[val], index)
	this.elements[index] = last
	this.mapping[last][index] = struct{}{}
	delete(this.mapping[last], lastElemIndex)

	// index := indexes[len(indexes)-1]
	// last := this.elements[this.count-1]
	// lastElementIndexes := this.mapping[last]

	// this.mapping[last][len(lastElementIndexes)-1] = index
	// fmt.Println(this.elements, last, index)
	// this.mapping[val] = this.mapping[val][:len(indexes)-1]

	// Need to fix this
	// sort.Ints(this.mapping[last])

	this.count--
	this.elements = this.elements[:this.count]
	if len(this.mapping[val]) == 0 {
		delete(this.mapping, val)
	}

	return true
}

func (this *RandomizedCollection) GetRandom() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	index := r.Intn(this.count)
	return this.elements[index]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
