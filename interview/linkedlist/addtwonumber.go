package main

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	var carry int
	for l1 != nil || l2 != nil {
		var left, right int
		if l1 != nil {
			left = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			right = l2.Val
			l2 = l2.Next
		}

		total := (left + right + carry) % 10
		carry = (left + right + carry) / 10
		curr.Next = &ListNode{total, nil}
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return head.Next
}
