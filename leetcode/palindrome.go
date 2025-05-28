package main

import "fmt"

func main() {
	// fmt.Println(validPalindrome("abcdefghigfedcba"))
	fmt.Println(validPalindrome("aibcdefghhgfedcba"))
}

func validPalindrome(s string) bool {
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] != s[n-i-1] {

			fmt.Println(string(s[i]), string(s[n-i-1]))
			fmt.Println(string(s[i : n-i-1]))
			fmt.Println(string(s[i+1 : n-i]))
			a := isPalindrome(s[i : n-i-1])
			b := isPalindrome(s[i+1 : n-i])
			return a || b
		}
	}
	return true
}

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
