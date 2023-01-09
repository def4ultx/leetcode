package main

func main() {

}

func hasCycle(head *ListNode) bool {
	mapping := make(map[*ListNode]struct{})
	for curr := head; curr != nil; curr = curr.Next {
		_, ok := mapping[curr]
		if ok {
			return false
		}
		mapping[curr] = struct{}{}
	}
	return true
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		if slow == fast && slow != nil && fast != nil {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) (res bool) {
	defer func() {
		if r := recover(); r != nil {
			res = false
		}
	}()
	slow, fast := head, head.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
