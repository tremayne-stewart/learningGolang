// Author: Tremayne Stewart
// 13 Jan 2022
package main

// Leetcode doesn't allow access to math.MaxInt, but they
// do give the constraint that the value will be -10^4 < x < 10^
const MaxInt = 10001

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var answerRoot = ListNode{}
	answerPtr := &answerRoot
	minIndex := MaxInt
	minValue := MaxInt

	// Once minIndex is -1, there are no more nodes to touch.
	for minIndex != -1 {
		minValue = MaxInt
		minIndex = -1
		for i := 0; i < len(lists); i++ {

			// Check for nil ptr
			if lists[i] != nil {
				if lists[i].Val < minValue {
					minValue = lists[i].Val
					minIndex = i
				}
			}
		}
		if minIndex != -1 {
			answerPtr.Next = &ListNode{}
			answerPtr.Next.Val = minValue

			// Increment answerPtr and the consumed Listnode.
			answerPtr = answerPtr.Next
			lists[minIndex] = lists[minIndex].Next
		}
	}
	return answerRoot.Next
}
