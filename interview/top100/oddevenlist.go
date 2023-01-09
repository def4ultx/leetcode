package main

func main() {

}

func oddEvenList(head *ListNode) *ListNode {
	left, right := &ListNode{0, nil}, &ListNode{0, nil}
	odd, even := left, right

	for curr, i := head, 0; curr != nil; curr, i = curr.Next, i+1 {
		if i%2 == 0 {
			odd.Next = curr
			odd = odd.Next
		} else {
			even.Next = curr
			even = even.Next
		}
	}

	odd.Next, even.Next = nil, nil
	odd = left
	for odd.Next != nil {
		odd = odd.Next
	}
	odd.Next = even.Next
	return left.Next
}
