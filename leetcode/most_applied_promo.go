package main

import "fmt"

func main() {
	pp := &MostPromo{counter: make(map[int]int)}
	pp.Apply(1)
	pp.Apply(2)
	pp.Apply(3)
	pp.Apply(1)
	fmt.Println(pp.Query())
	pp.Apply(2)
	pp.Apply(2)
	fmt.Println(pp.Query())
	pp.Apply(1)
	pp.Apply(3)
	pp.Apply(3)
	pp.Apply(3)
	fmt.Println(pp.Query())
}

type MostPromo struct {
	counter map[int]int
	max     int
	maxId   int
}

func (p *MostPromo) Apply(id int) {
	p.counter[id]++

	if p.counter[id] > p.max {
		p.max = p.counter[id]
		p.maxId = id
	}
}

func (p *MostPromo) Query() int {
	return p.maxId
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
