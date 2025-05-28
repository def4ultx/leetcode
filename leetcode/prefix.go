package main

func main() {

	// longestCommonPrefix([]int{56554, 5655359, 43456}, []int{565, 1223})

	// longestCommonPrefix([]int{1, 10, 100}, []int{1000, 2})
	longestCommonPrefix([]int{1, 2, 3}, []int{4, 4, 4})
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {

	// left, right := NewTrie(), NewTrie()
	// for _, v := range arr1 {
	// 	left.Append(v)
	// }
	// for _, v := range arr2 {
	// 	right.Append(v)
	// }

	// n := MaxDepth(left, right, 0)
	// fmt.Println(n)

	// fmt.Println(left)
	// fmt.Println(left.children[4])
	// fmt.Println(left.children[4].children[3])
	// fmt.Println(left.children[4].children[3].children[4])
	// fmt.Println(left.children[4].children[3].children[4].children[5])
	// fmt.Println(left.children[4].children[3].children[4].children[5].children[6])
	// fmt.Println(right)

	left := NewTrie()
	for _, v := range arr1 {
		left.Append(v)
	}

	max := 0
	// xs := make([]int, 10)
	for _, v := range arr2 {
		idx := 0
		for ; v > 0; idx++ {
			xs[idx] = v % 10
			v = v / 10
		}

		count := 0
		current := left
		for i := idx - 1; i >= 0; i-- {
			digit := xs[i]
			if current.children[digit] == nil {
				break
			}
			current = current.children[digit]

			count++
		}

		max = Max(max, count)

	}

	return max
}

// func MaxDepth(left, right *Trie, level int) int {
// 	if left == nil || right == nil {
// 		return 0
// 	}

// 	max := 0
// 	for i := 0; i < 10; i++ {
// 		if left.children[i] != nil && right.children[i] != nil {
// 			fmt.Println("level", level, "idx", i)
// 			n := MaxDepth(left.children[i], right.children[i], level+1)
// 			max = Max(max, n+1)
// 		}
// 	}
// 	return max
// }

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Trie struct {
	children [10]*Trie
	end      bool
}

func NewTrie() *Trie {
	return &Trie{children: [10]*Trie{}}
}

var xs = make([]int, 10)

func (t *Trie) Append(n int) {
	idx := 0
	for ; n > 0; idx++ {
		xs[idx] = n % 10
		n = n / 10
	}

	// fmt.Println(xs, idx)

	current := t
	for i := idx - 1; i >= 0; i-- {
		digit := xs[i]
		if current.children[digit] == nil {
			current.children[digit] = NewTrie()
		}
		current = current.children[digit]
	}
	current.end = true
}
