package main

/*
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5].
There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
*/
//Most efficient method:
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmp1 := headA
	tmp2 := headB

	if tmp1 == nil || tmp2 == nil {
		return nil
	}
	for tmp1 != tmp2 {
		if tmp1 != nil {
			tmp1 = tmp1.Next
		} else {
			tmp1 = headB
		}
		if tmp2 != nil {
			tmp2 = tmp2.Next
		} else {
			tmp2 = headA
		}
	}
	return tmp1
}
