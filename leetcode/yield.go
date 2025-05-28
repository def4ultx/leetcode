package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for v := range IterItem(xs) {
		fmt.Println("print", v)
	}

	for v := range Iter2 {
		if v > 6 {
			break
		}
		fmt.Println(v)
	}

	next, stop := iter.Pull(IterItem(xs))
	for v, ok := next(); ok; v, ok = next() {
		fmt.Println("pull", v)
	}
	stop()

	for v := range slices.Backward(xs) {
		fmt.Println("print", v)
	}

	for v := range slices.All(xs) {
		v = v * 2
	}

	fmt.Println(xs)
}

func Iter2(yield func(int) bool) {
	for _, v := range []int{2, 4, 6, 8, 10} {
		if !yield(v) {
			return
		}
	}
}

func IterItem[V any](xs []V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range xs {
			fmt.Println("process", v)
			if !yield(v) {
				return
			}
			fmt.Println("after process", v)
		}
	}
}
