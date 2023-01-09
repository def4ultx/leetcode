package main

import (
	"fmt"
	"strconv"
)

func main() {
	countAndSay(7)
}

func countAndSay(n int) string {

	current := "1"
	for i := 2; i <= n; i++ {

		fmt.Println("temp", i, current)
		temp := ""
		c, count := current[0], 1
		for k := 1; k < len(current); k++ {
			// fmt.Println(c, current[k], k, current)
			if c == current[k] {
				count++
				continue
			} else {
				temp = temp + toStr(c, count)
				count = 1
				c = current[k]
			}
		}
		temp = temp + toStr(c, count)
		fmt.Println("new str", i, temp, "|", c, count)
		current = temp
	}
	return current
}

func toStr(char byte, count int) string {
	return strconv.Itoa(count) + string(char)
}
