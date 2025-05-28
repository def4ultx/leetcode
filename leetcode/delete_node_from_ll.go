package main

import "sort"

func main() {

}

func modifiedList(nums []int, head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if len(nums) == 0 {
		return head
	}

	sets := make(map[int]struct{})
	for _, v := range nums {
		sets[v] = struct{}{}
	}

	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy

	for head != nil {
		if _, ok := sets[head.Val]; !ok {
			prev = head
			head = head.Next
			continue
		}

		prev.Next = head.Next
		head = head.Next

	}
	return dummy.Next
}

func modifiedListSort(nums []int, head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if len(nums) == 0 {
		return head
	}

	sort.Ints(nums)
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	prev := dummy

	n := len(nums)
	for head != nil {
		idx := sort.SearchInts(nums, head.Val)
		if idx == n || (idx < n && nums[idx] != head.Val) {
			prev = head
			head = head.Next
			continue
		}
		prev.Next = head.Next
		head = head.Next

	}
	return dummy.Next
}
