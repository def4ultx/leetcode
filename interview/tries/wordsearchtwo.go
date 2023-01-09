package main

import "fmt"

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	findWords(board, words)
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() *Trie {
	return &Trie{}
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

func findWords(board [][]byte, words []string) []string {
	root := Constructor()
	for _, w := range words {
		root.Insert(w)
	}
	fmt.Println("words", root)
	found := make(map[string]struct{})

	var dfs func(row, col int, current string, trie *Trie)
	dfs = func(row, col int, current string, trie *Trie) {
		fmt.Println("visit", row, col, current)
		if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) {
			return
		}
		if board[row][col] == '#' {
			return
		}
		if trie == nil {
			return
		}

		c := board[row][col]
		current += string(c)
		tt := trie.children[c-'a']
		if tt != nil && tt.isEnd {
			found[current] = struct{}{}
		}

		tmp := board[row][col]
		board[row][col] = '#'

		direction := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, v := range direction {
			dfs(row+v.x, col+v.y, current, tt)
		}
		board[row][col] = tmp
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, "", root)
		}
	}

	result := make([]string, 0)
	for k := range found {
		result = append(result, k)
	}

	fmt.Println("result")
	fmt.Println(result)

	return result
}
