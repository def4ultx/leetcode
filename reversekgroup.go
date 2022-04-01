package main

import "fmt"

func main() {
	number1 := NewListNode(1)
	number2 := NewListNode(2)
	number3 := NewListNode(3)
	number4 := NewListNode(4)
	number5 := NewListNode(5)

	number1.Next = number2
	number2.Next = number3
	number3.Next = number4
	number4.Next = number5

	x := reverseKGroup(number1, 2)
	// x := Reverse(number1)
	fmt.Println("----")
	for {
		if x != nil {
			fmt.Println(x.Val)
			x = x.Next
		} else {
			break
		}
	}
}

func NewListNode(n int) *ListNode {
	return &ListNode{Val: n}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func CountList(head *ListNode) int {
	count := 0
	for {
		if head != nil {
			head = head.Next
			count++
		} else {
			break
		}
	}
	return count
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := CountList(head)
	totalReverseCount := count / k
	fmt.Println("total loop", totalReverseCount, count)

	temp := head
	tt := head
	for i := 0; i < totalReverseCount; i++ {
		current := temp
		var next *ListNode
		for j := 0; j < k; j++ {
			if current == nil {
				break
			}

			t := current.Next
			current.Next = next
			next = current
			current = t
		}
		fmt.Println(current, next, temp)
		if i == 0 {
			head = next
		} else {
			tt.Next = next
			tt = temp
		}

		// t := current.Next
		// temp.Next = current
		temp = current
	}
	tt.Next = temp

	return head
}

func Reverse(head *ListNode) *ListNode {
	current := head
	var next *ListNode
	for {
		if current == nil {
			break
		}

		t := current.Next
		current.Next = next
		next = current
		current = t
		fmt.Println(current, next)
	}
	fmt.Println("---", next.Val)
	return next
}

// func main() {
// 	// myAtoi("42")
// 	// myAtoi("-42")
// 	// myAtoi("000042")
// 	// myAtoi("42AAA42")
// 	// myAtoi("-91283472332")
// 	// myAtoi("+-12") // 0
// 	// myAtoi("00000-42a1234") // 0

// 	myAtoi("20000000000000000000")
// }

// func myAtoi(s string) int {

// 	str := strings.Trim(s, " ")
// 	// str = strings.TrimLeft(str, "0")
// 	sign := 0
// 	foundNumber := false

// 	var result string
// 	for _, v := range str {
// 		if v == '+' && sign == 0 && !foundNumber {
// 			sign = 1
// 			continue
// 		} else if v == '-' && sign == 0 && !foundNumber {
// 			sign = 2
// 			continue
// 		} else if v == '-' && sign != 0 {
// 			break
// 		} else if v == '+' && sign != 0 {
// 			break
// 		} else if v == '-' && sign == 0 && foundNumber {
// 			break
// 		} else if v == '+' && sign == 0 && foundNumber {
// 			break
// 		}

// 		if isNumber(v) {
// 			foundNumber = true
// 			result += string(v)
// 		} else {
// 			break
// 		}
// 	}
// 	result = strings.TrimLeft(result, "0")
// 	if len(result) > 10 {
// 		if sign == 2 {
// 			return -2147483648
// 		} else {
// 			return 2147483647
// 		}
// 	}
// 	if sign == 2 {
// 		result = "-" + result
// 	}

// 	// fmt.Println(str)
// 	// fmt.Println(sign)
// 	// fmt.Println(result)
// 	r, err := strconv.Atoi(result)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0
// 	}

// 	if r > 2147483647 {
// 		r = 2147483647
// 	}
// 	if r < -2147483648 {
// 		r = -2147483648
// 	}
// 	fmt.Println(r)
// 	return r
// }

// func isNumber(v rune) bool {
// 	switch v {
// 	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
// 		return true
// 	}
// 	return false
// }
