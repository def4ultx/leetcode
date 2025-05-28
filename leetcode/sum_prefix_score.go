package main

import "fmt"

func main() {
	sumPrefixScores([]string{"abc", "ab", "bc", "b"})
}

func NewTrie() *Trie {
	return &Trie{
		children: [26]*Trie{},
		count:    0,
	}
}

type Trie struct {
	children [26]*Trie
	count    int
}

func (t *Trie) AddWord(word string) {
	curr := t
	for _, v := range word {
		idx := v - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = NewTrie()
		}
		curr = curr.children[idx]
		curr.count++
	}
}

func (t *Trie) GetScore(word string) int {
	total := 0
	curr := t
	for _, v := range word {
		curr = curr.children[v-'a']
		total += curr.count
	}
	return total
}

func sumPrefixScores(words []string) []int {
	trie := NewTrie()
	for _, v := range words {
		trie.AddWord(v)
	}

	result := make([]int, len(words))
	for i, v := range words {
		result[i] = trie.GetScore(v)
	}
	fmt.Println(result)
	return result
}
