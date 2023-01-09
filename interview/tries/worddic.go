package main

import "fmt"

func main() {
	d := Constructor()
	dic := &d

	dic.AddWord("a")
	dic.AddWord("a")
	// dic.Search(".")
	// dic.Search("a")
	// dic.Search("aa")
	// dic.Search("a")
	// dic.Search(".a")
	dic.Search("a.")

	// dic.AddWord("bad")
	// dic.AddWord("dad")
	// dic.AddWord("mad")
	// dic.Search("pad") // return False
	// dic.Search("bad") // return True
	// dic.Search(".ad") // return True
	// dic.Search("b..") // return True
}

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &WordDictionary{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	isEnd := false

	result := make([]bool, 0)
	curr := this
	for i, ch := range word {
		cc := 0

		for _, child := range curr.children {
			if child == nil {
				continue
			}
			cc++
		}

		fmt.Println(i, string(ch), word[i+1:], cc)
		if ch == '.' {
			for _, child := range curr.children {
				if child == nil {
					continue
				}
				r := child.Search(word[i+1:])
				result = append(result, r)
			}
			fmt.Println("result", result)

			isEnd = true
			break
		} else {
			idx := ch - 'a'
			fmt.Println(curr.children[idx])
			if curr.children[idx] == nil {
				fmt.Println("return false")
				isEnd = true
				break
			}
			curr = curr.children[idx]
		}
	}

	if !isEnd {
		fmt.Println("add isEnd", isEnd)
		result = append(result, curr.isEnd)
	}

	var found bool
	for _, v := range result {
		found = found || v
	}
	fmt.Println(result, curr.isEnd, found)
	return found
}
