package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// res:=head
	next := head.Next

	head.Next = next.Next
	next.Next = head
	swap(head, head.Next)
	return next
}
func swap(head, next *ListNode) {
	if next == nil || next.Next == nil {
		return
	}
	// exchange tmp with next
	tmp := next.Next
	head.Next = tmp
	next.Next = tmp.Next
	tmp.Next = next

	swap(next, next.Next)

}
