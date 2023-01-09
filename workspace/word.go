package main

import (
	"fmt"
	"math"
)

func main() {
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	beginWord := "cet"
	endWord := "ism"
	wordList := []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"}

	count := ladderLength(beginWord, endWord, wordList)
	fmt.Println("total count:", count)
}

var memory = make(map[string]int)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	found := false
	for _, v := range wordList {
		if v == endWord {
			found = true
		}
	}
	if !found {
		return 0
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

	visited := make(map[string]struct{})
	// memory := make(map[string]int)
	count := findLadderLenght(beginWord, endWord, mapping, visited)
	if count >= math.MaxInt32 {
		return 0
	}
	return count
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

func findLadderLenght(start, end string, mapping map[string]map[string]struct{}, visited map[string]struct{}) int {
	if start == end {
		return 1
	}
	mem, ok := memory[start]
	if ok {
		return mem
	}

	counts := make([]int, 0)

	newVisited := make(map[string]struct{})
	for k, v := range visited {
		newVisited[k] = v
	}
	newVisited[start] = struct{}{}

	words := mapping[start]
	for w := range words {
		if _, ok := visited[w]; !ok {
			count := findLadderLenght(w, end, mapping, newVisited) + 1
			if count < math.MaxInt32 {
				counts = append(counts, count)
			}
		}
	}

	min := findMin(counts)
	memory[start] = min
	return min
}

func findMin(numbers []int) int {
	min := math.MaxInt32
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}

// package main

// import (
// 	"fmt"
// 	"math"
// 	"sync"
// )

// func main() {
// 	// beginWord := "hit"
// 	// endWord := "cog"
// 	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
// 	beginWord := "cet"
// 	endWord := "ism"
// 	wordList := []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"}

// 	count := ladderLength(beginWord, endWord, wordList)
// 	fmt.Println("total count:", count)
// }

// // // var memory = make(map[string]int)
// // var memory sync.Map
// // var mutex sync.Mutex

// // func ladderLength(beginWord string, endWord string, wordList []string) int {
// // 	// found := false
// // 	// for _, v := range wordList {
// // 	// 	if v == endWord {
// // 	// 		found = true
// // 	// 	}
// // 	// }
// // 	// if !found {
// // 	// 	return 0
// // 	// }

// // 	wordList = append(wordList, beginWord)

// // 	mapping := make(map[string]map[string]struct{})
// // 	for _, word := range wordList {
// // 		mapping[word] = make(map[string]struct{})
// // 	}

// // 	for _, src := range wordList {
// // 		for _, dst := range wordList {
// // 			if src == dst {
// // 				continue
// // 			}

// // 			if diffByOne(src, dst) {
// // 				mapping[src][dst] = struct{}{}
// // 				mapping[dst][src] = struct{}{}
// // 			}
// // 		}
// // 	}

// // 	mm := make(map[string][]string)
// // 	for k, v := range mapping {
// // 		list := make([]string, 0)
// // 		for w := range v {
// // 			list = append(list, w)
// // 		}
// // 		mm[k] = list
// // 	}

// // 	// fmt.Println(mapping)

// // 	visited := make([]string, 0)
// // 	// memory := make(map[string]int)
// // 	count := findLadderLenght(beginWord, endWord, mm, visited)
// // 	if count >= 99999 {
// // 		return 0
// // 	}
// // 	return count
// // }

// // func diffByOne(start, end string) bool {
// // 	var count int
// // 	for i := 0; i < len(start); i++ {
// // 		if start[i] != end[i] {
// // 			count++
// // 		}
// // 	}
// // 	return count == 1
// // }

// // func findLadderLenght(start, end string, mapping map[string][]string, visited []string) int {
// // 	// fmt.Println("start", start)

// // 	if start == end {
// // 		return 1
// // 	}
// // 	// mem, ok := memory[start]
// // 	// if ok {
// // 	// 	// fmt.Println("found memory", start, mem)
// // 	// 	return mem
// // 	// }
// // 	mem, ok := memory.Load(start)
// // 	if ok {
// // 		return mem.(int)
// // 	}

// // 	counts := make([]int, 0)

// // 	newVisited := make([]string, 0)
// // 	for _, v := range visited {
// // 		newVisited = append(newVisited, v)
// // 	}
// // 	newVisited = append(newVisited, start)

// // 	for _, w := range mapping[start] {
// // 		if !exist(w, visited) {
// // 			count := findLadderLenght(w, end, mapping, newVisited) + 1
// // 			if count < 99999 {
// // 				counts = append(counts, count)
// // 			}
// // 		}
// // 	}

// // 	min := findMin(counts)
// // 	// memory[start] = min

// // 	mutex.Lock()
// // 	defer mutex.Unlock()
// // 	if min < 99999 {
// // 	}
// // 	mem, ok = memory.Load(start)
// // 	if ok {
// // 		value := mem.(int)
// // 		if min < value {
// // 			memory.Store(start, min)
// // 		}
// // 	} else {
// // 		memory.Store(start, min)
// // 	}

// // 	return min
// // }

// // func exist(word string, words []string) bool {
// // 	for _, v := range words {
// // 		if v == word {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// // func findMin(numbers []int) int {
// // 	min := 99999
// // 	for _, v := range numbers {
// // 		if v < min {
// // 			min = v
// // 		}
// // 	}
// // 	return min
// // }

// // func ladderLength22(beginWord string, endWord string, wordList []string) int {
// // 	m := make(map[string]bool)
// // 	for _, v := range wordList {
// // 		m[v] = true
// // 	}
// // 	list := []string{}
// // 	list = append(list, beginWord)
// // 	rst := 1
// // 	for len(list) != 0 {
// // 		l := len(list)
// // 		for i := 0; i < l; i++ {
// // 			word := list[0]
// // 			list = list[1:]
// // 			if word == endWord {
// // 				return rst
// // 			}
// // 			m[word] = false
// // 			for i, _ := range word {
// // 				for j := 0; j < 26; j++ {
// // 					tmp := []rune(word)
// // 					tmp[i] = rune('a' + j)
// // 					ts := string(tmp)
// // 					if v, ok := m[ts]; ok && v && ((len(list) != 0 && ts != list[len(list)-1]) || len(list) == 0) {
// // 						list = append(list, ts)
// // 					}
// // 				}
// // 			}
// // 		}
// // 		rst++
// // 	}
// // 	return 0
// // }

// // // func main() {
// // // 	// beginWord := "hit"
// // // 	// endWord := "cog"
// // // 	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

// // // 	beginWord := "cet"
// // // 	endWord := "ism"
// // // 	wordList := []string{"kid", "tag", "pup", "ail", "tun", "woo", "erg", "luz", "brr", "gay", "sip", "kay", "per", "val", "mes", "ohs", "now", "boa", "cet", "pal", "bar", "die", "war", "hay", "eco", "pub", "lob", "rue", "fry", "lit", "rex", "jan", "cot", "bid", "ali", "pay", "col", "gum", "ger", "row", "won", "dan", "rum", "fad", "tut", "sag", "yip", "sui", "ark", "has", "zip", "fez", "own", "ump", "dis", "ads", "max", "jaw", "out", "btu", "ana", "gap", "cry", "led", "abe", "box", "ore", "pig", "fie", "toy", "fat", "cal", "lie", "noh", "sew", "ono", "tam", "flu", "mgm", "ply", "awe", "pry", "tit", "tie", "yet", "too", "tax", "jim", "san", "pan", "map", "ski", "ova", "wed", "non", "wac", "nut", "why", "bye", "lye", "oct", "old", "fin", "feb", "chi", "sap", "owl", "log", "tod", "dot", "bow", "fob", "for", "joe", "ivy", "fan", "age", "fax", "hip", "jib", "mel", "hus", "sob", "ifs", "tab", "ara", "dab", "jag", "jar", "arm", "lot", "tom", "sax", "tex", "yum", "pei", "wen", "wry", "ire", "irk", "far", "mew", "wit", "doe", "gas", "rte", "ian", "pot", "ask", "wag", "hag", "amy", "nag", "ron", "soy", "gin", "don", "tug", "fay", "vic", "boo", "nam", "ave", "buy", "sop", "but", "orb", "fen", "paw", "his", "sub", "bob", "yea", "oft", "inn", "rod", "yam", "pew", "web", "hod", "hun", "gyp", "wei", "wis", "rob", "gad", "pie", "mon", "dog", "bib", "rub", "ere", "dig", "era", "cat", "fox", "bee", "mod", "day", "apr", "vie", "nev", "jam", "pam", "new", "aye", "ani", "and", "ibm", "yap", "can", "pyx", "tar", "kin", "fog", "hum", "pip", "cup", "dye", "lyx", "jog", "nun", "par", "wan", "fey", "bus", "oak", "bad", "ats", "set", "qom", "vat", "eat", "pus", "rev", "axe", "ion", "six", "ila", "lao", "mom", "mas", "pro", "few", "opt", "poe", "art", "ash", "oar", "cap", "lop", "may", "shy", "rid", "bat", "sum", "rim", "fee", "bmw", "sky", "maj", "hue", "thy", "ava", "rap", "den", "fla", "auk", "cox", "ibo", "hey", "saw", "vim", "sec", "ltd", "you", "its", "tat", "dew", "eva", "tog", "ram", "let", "see", "zit", "maw", "nix", "ate", "gig", "rep", "owe", "ind", "hog", "eve", "sam", "zoo", "any", "dow", "cod", "bed", "vet", "ham", "sis", "hex", "via", "fir", "nod", "mao", "aug", "mum", "hoe", "bah", "hal", "keg", "hew", "zed", "tow", "gog", "ass", "dem", "who", "bet", "gos", "son", "ear", "spy", "kit", "boy", "due", "sen", "oaf", "mix", "hep", "fur", "ada", "bin", "nil", "mia", "ewe", "hit", "fix", "sad", "rib", "eye", "hop", "haw", "wax", "mid", "tad", "ken", "wad", "rye", "pap", "bog", "gut", "ito", "woe", "our", "ado", "sin", "mad", "ray", "hon", "roy", "dip", "hen", "iva", "lug", "asp", "hui", "yak", "bay", "poi", "yep", "bun", "try", "lad", "elm", "nat", "wyo", "gym", "dug", "toe", "dee", "wig", "sly", "rip", "geo", "cog", "pas", "zen", "odd", "nan", "lay", "pod", "fit", "hem", "joy", "bum", "rio", "yon", "dec", "leg", "put", "sue", "dim", "pet", "yaw", "nub", "bit", "bur", "sid", "sun", "oil", "red", "doc", "moe", "caw", "eel", "dix", "cub", "end", "gem", "off", "yew", "hug", "pop", "tub", "sgt", "lid", "pun", "ton", "sol", "din", "yup", "jab", "pea", "bug", "gag", "mil", "jig", "hub", "low", "did", "tin", "get", "gte", "sox", "lei", "mig", "fig", "lon", "use", "ban", "flo", "nov", "jut", "bag", "mir", "sty", "lap", "two", "ins", "con", "ant", "net", "tux", "ode", "stu", "mug", "cad", "nap", "gun", "fop", "tot", "sow", "sal", "sic", "ted", "wot", "del", "imp", "cob", "way", "ann", "tan", "mci", "job", "wet", "ism", "err", "him", "all", "pad", "hah", "hie", "aim", "ike", "jed", "ego", "mac", "baa", "min", "com", "ill", "was", "cab", "ago", "ina", "big", "ilk", "gal", "tap", "duh", "ola", "ran", "lab", "top", "gob", "hot", "ora", "tia", "kip", "han", "met", "hut", "she", "sac", "fed", "goo", "tee", "ell", "not", "act", "gil", "rut", "ala", "ape", "rig", "cid", "god", "duo", "lin", "aid", "gel", "awl", "lag", "elf", "liz", "ref", "aha", "fib", "oho", "tho", "her", "nor", "ace", "adz", "fun", "ned", "coo", "win", "tao", "coy", "van", "man", "pit", "guy", "foe", "hid", "mai", "sup", "jay", "hob", "mow", "jot", "are", "pol", "arc", "lax", "aft", "alb", "len", "air", "pug", "pox", "vow", "got", "meg", "zoe", "amp", "ale", "bud", "gee", "pin", "dun", "pat", "ten", "mob"}

// // // 	count := ladderLength(beginWord, endWord, wordList)
// // // 	fmt.Println("total count:", count)
// // // }

// var memory = make(map[string]int)
// var mutex sync.Mutex

// func ladderLength(beginWord string, endWord string, wordList []string) int {
// 	found := false
// 	for _, v := range wordList {
// 		if v == endWord {
// 			found = true
// 		}
// 	}
// 	if !found {
// 		return 0
// 	}

// 	wordList = append(wordList, beginWord)

// 	mapping := make(map[string]map[string]struct{})
// 	for _, word := range wordList {
// 		mapping[word] = make(map[string]struct{})
// 		memory[word] = math.MaxInt32
// 	}

// 	for _, src := range wordList {
// 		for _, dst := range wordList {
// 			if src == dst {
// 				continue
// 			}

// 			if diffByOne(src, dst) {
// 				mapping[src][dst] = struct{}{}
// 				mapping[dst][src] = struct{}{}
// 			}
// 		}
// 	}

// 	// fmt.Println(mapping)

// 	visited := make(map[string]struct{})
// 	// memory := make(map[string]int)
// 	count := findLadderLenght(beginWord, endWord, mapping, visited)
// 	if count >= math.MaxInt32 {
// 		return 0
// 	}
// 	return count
// }

// func diffByOne(start, end string) bool {
// 	var count int
// 	for i := 0; i < len(start); i++ {
// 		if start[i] != end[i] {
// 			count++
// 		}
// 	}
// 	return count == 1
// }

// func findLadderLenght(start, end string, mapping map[string]map[string]struct{}, visited map[string]struct{}) int {
// 	// fmt.Println("start", start)

// 	if start == end {
// 		return 1
// 	}
// 	key := start + end
// 	fmt.Println(key)
// 	mem, ok := memory[key]
// 	// if ok && mem != math.MaxInt32 {
// 	if ok {
// 		fmt.Println("found memory", key, mem)
// 		return mem
// 	}
// 	if _, ok := visited[key]; ok {
// 		return math.MaxInt32
// 	}

// 	counts := make([]int, 0)

// 	newVisited := make(map[string]struct{})
// 	for k, v := range visited {
// 		newVisited[k] = v
// 	}
// 	newVisited[start] = struct{}{}

// 	// var visit int
// 	words := mapping[start]
// 	for w := range words {
// 		count := findLadderLenght(w, end, mapping, newVisited) + 1
// 		counts = append(counts, count)
// 		// if _, ok := visited[w]; !ok {
// 		// 	count := findLadderLenght(w, end, mapping, newVisited) + 1
// 		// 	if count < math.MaxInt32 {
// 		// 		counts = append(counts, count)
// 		// 	}
// 		// 	// if w == "hot" {
// 		// 	// 	fmt.Println("found hot", start, count)
// 		// 	// }
// 		// 	// fmt.Println(w, count)

// 		// 	visit++
// 		// }
// 	}

// 	min := findMin(counts)
// 	// if min < memory[key] {
// 	// 	memory[key] = min
// 	// }
// 	memory[key] = min
// 	// if min < mem && mem != 0 {
// 	// 	memory[start] = min
// 	// }
// 	// if !ok && len(counts) != 0 {
// 	// 	memory[start] = min
// 	// }

// 	// fmt.Println(start, end, counts, words, visited, min)

// 	return min
// }

// func findMin(numbers []int) int {
// 	min := math.MaxInt32
// 	for _, v := range numbers {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }
