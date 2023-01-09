package main

import "fmt"

func main() {
	// isPalindrome("A man, a plan, a canal: Panama")
	isPalindrome("race a car")
	isPalindrome("0P")
}

func isPalindrome(s string) bool {

	l, r := 0, len(s)-1
	fmt.Println(l, r)
	for l < r {
		fmt.Println(l, r)
		if !isCharacter(s[l]) {
			l++
			continue
		}
		if !isCharacter(s[r]) {
			r--
			continue
		}
		fmt.Println(l, r)
		if !isSameChar(s[l], s[r]) {
			fmt.Println(s[l], s[r], l, r)
			return false
		}
		l++
		r--
	}

	return true
}

func isCharacter(r byte) bool {

	// if r >= 97 {
	// 	r = r - 32
	// }
	// if r >= 65 && r <= 90 {
	// 	return true
	// }
	// if r >= 48 && r <= 57 {
	// 	return true
	// }

	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}

func isSameChar(a, b byte) bool {
	if a >= 97 {
		a = a - 32
	}
	if b >= 97 {
		b = b - 32
	}
	return a == b
}
