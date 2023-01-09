package main

import "fmt"

func main() {
	// fmt.Println(checkInclusion("ab", "eidbaooo"))
	// fmt.Println(checkInclusion("ab", "eidboaoo"))
	// fmt.Println(checkInclusion("ab", "eidboooooo"))
	// fmt.Println(checkInclusion("abc", "cccddccbabbbaaaa"))
	// fmt.Println(checkInclusion("bba", "bbbba"))
	// fmt.Println(checkInclusion("bba", "bbycyyyyyyyy"))
	fmt.Println(checkInclusion("ba", "cbyyyyyy"))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	if len(s1) == 0 {
		return false
	}

	s1mapping := make([]int, 26)
	s2mapping := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1mapping[s1[i]-'a']++
		s2mapping[s2[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if s1mapping[i] == s2mapping[i] {
			count++
		}
	}

	// fmt.Println(count)
	// fmt.Println(s1mapping)
	// fmt.Println(s2mapping)

	for i := len(s1); i < len(s2); i++ {
		if count == 26 {
			return true
		}

		fmt.Println(string(s2[i]), count)
		fmt.Println(s1mapping)
		fmt.Println(s2mapping)

		s2mapping[s2[i]-'a']++
		fmt.Println(s2mapping)
		if s2mapping[s2[i]-'a'] == s1mapping[s2[i]-'a'] {
			count++
			// } else {
			// 	fmt.Println("remove")
			// 	count--
			// }
		} else if s2mapping[s2[i]-'a']-1 == s1mapping[s2[i]-'a'] {
			fmt.Println("remove!?")
			count--
		}

		s2mapping[s2[i-len(s1)]-'a']--
		fmt.Println(s2mapping)
		if s2mapping[s2[i-len(s1)]-'a'] == s1mapping[s2[i-len(s1)]-'a'] {
			count++
		} else if s2mapping[s2[i-len(s1)]-'a']+1 == s1mapping[s2[i-len(s1)]-'a'] {
			fmt.Println("remove?")
			count--
		}
		fmt.Println("change count to", count)
	}

	fmt.Println("count", count)

	return count == 26
}
