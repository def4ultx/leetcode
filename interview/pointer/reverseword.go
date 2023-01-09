package main

import "fmt"

func main() {
	// reverseWords("a good   example")
	// reverseWords("   hello world   ")
	reverseWords("   hello    world   ")
}

func reverseWords(s string) string {
	str := " "
	for i := 0; i < len(s); i++ {
		if str[len(str)-1] == ' ' && s[i] == ' ' {
			continue
		}
		str += string(s[i])
	}
	str = str[1:]
	fmt.Println(str)

	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != ' ' && i != 0 {
			continue
		}

		fmt.Println(str[i], string(str[i]))

		idx := i + 1
		if i == 0 {
			idx = 0
		}
		for idx < len(str) && str[idx] != ' ' {
			result += string(str[idx])
			idx++
		}
		result += " "
	}

	fmt.Println(str[0] == ' ', str[0])
	if result[0] == ' ' {
		result = result[1:]
	}

	fmt.Println(result[:len(result)-1])
	return result[:len(result)-1]
}

// func reverseWords(s string) string {
// 	str := " "
// 	for i := 0; i < len(s); i++ {
// 		if str[len(str)-1] == ' ' && s[i] == ' ' {
// 			continue
// 		}
// 		str += string(s[i])
// 	}
//     str = strings.TrimLeft(str, " ")
//     str = strings.TrimRight(str, " ")
//     words := strings.Split(str, " ")
//     for i := 0;i<len(words)/2;i++ {
//         words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
//     }
//     return strings.Join(words, " ")
// }
