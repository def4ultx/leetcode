package main

import "math"

func main() {
	// findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"})
	words := []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	findLadders("qa", "sq", words)
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	found := false
	for _, v := range wordList {
		if v == endWord {
			found = true
		}
	}
	if !found {
		return nil
	}

	wordList = append(wordList, beginWord)

	mapping := make(map[string]map[string]struct{})
	for _, word := range wordList {
		mapping[word] = make(map[string]struct{})
	}

	for _, src := range wordList {
		for _, dst := range wordList {
			if src == dst {
				continue
			}

			if diffByOne(src, dst) {
				mapping[src][dst] = struct{}{}
				mapping[dst][src] = struct{}{}
			}
		}
	}

	// count := findLadderLenght(beginWord, endWord, mapping)

	result := make([][]string, 0)
	x := math.MaxInt32
	findLadder(beginWord, endWord, []string{}, map[string]struct{}{}, &x, mapping, &result)

	res := make([][]string, 0)
	for _, v := range result {
		if len(v) == x {
			res = append(res, v)
		}
	}

	return res
}

func diffByOne(start, end string) bool {
	var count int
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			count++
		}
	}
	return count == 1
}

func findLadder(current, end string, path []string, visited map[string]struct{}, maxdepth *int, mapping map[string]map[string]struct{}, result *[][]string) {
	if current == end {
		if len(path) <= *maxdepth {

			path = append(path, end)
			temp := make([]string, len(path))
			copy(temp, path)
			*result = append(*result, temp)
			*maxdepth = len(path)
		}
		return
	}
	if len(path) >= *maxdepth {
		return
	}

	if _, ok := visited[current]; ok {
		return
	}

	words, ok := mapping[current]
	if !ok {
		return
	}

	visited[current] = struct{}{}
	path = append(path, current)

	for w := range words {
		findLadder(w, end, path, visited, maxdepth, mapping, result)
	}

	delete(visited, current)
	path = path[:len(path)-1]

}

func exists(xs []string, w string) bool {
	for _, v := range xs {
		if v == w {
			return true
		}
	}
	return false
}

func findLadderLenght(start, end string, mapping map[string]map[string]struct{}) int {

	type node struct {
		level int
		word  string
	}

	visited := make(map[string]struct{})

	queue := []node{
		{1, start},
	}

	for {
		if len(queue) == 0 {
			break
		}

		current := queue[0]
		queue = queue[1:]

		words, ok := mapping[current.word]
		if !ok {
			continue
		}

		for w := range words {
			if w == end {
				return current.level + 1
			}

			if _, ok := visited[w]; !ok {
				visited[w] = struct{}{}
				queue = append(queue, node{current.level + 1, w})
			}
		}
	}

	return 0
}
