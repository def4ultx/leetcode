package main

import "fmt"

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
	}
	findWords(board, nil)
}

func findWords(board [][]byte, words []string) []string {

	boardWords := make(map[string]struct{})
	for i := 0; i < len(board); i++ {
		w := ""
		for j := 0; j < len(board[i]); j++ {
			w += string(board[i][j])
			boardWords[w] = struct{}{}
		}
	}
	for i := 0; i < len(board[0]); i++ {
		w := ""
		for j := 0; j < len(board); j++ {
			w += string(board[j][i])
			boardWords[w] = struct{}{}
		}
	}
	fmt.Println(boardWords)

	trie := Constructor()
	for w := range boardWords {
		trie.Insert(w)
	}

	result := make([]string, 0)
	for _, v := range words {
		if trie.Search(v) {
			result = append(result, v)
		}
	}

	fmt.Println(result)
	return result
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}
