// Author: Tremayne Stewart
// 13 Jan 2022
package main

import (
	"testing"
)

func makeListNode(list []int) *ListNode {
	var root = ListNode{}
	head := &root
	for i := 0; i < len(list); i++ {
		head.Next = &ListNode{}
		head.Next.Val = list[i]
		head = head.Next
	}
	return root.Next
}

func TestMergeKSortedLists(t *testing.T) {
	// list1 = [1,4,5]
	var list1 = *makeListNode([]int{1, 4, 5})

	// list2 = [1,3,4]
	var list2 = *makeListNode([]int{1, 3, 4})

	// list3 = [2,6]
	var list3 = *makeListNode([]int{2, 6})

	// expectedList = [1,1,2,3,4,4,5,6]
	var expectedList = *makeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})
	expectedListHead := &expectedList

	resultRoot := mergeKLists([]*ListNode{&list1, &list2, &list3})
	mergedPrintHead := resultRoot
	for mergedPrintHead != nil {
		if expectedListHead.Val != mergedPrintHead.Val {
			t.Fatalf("Mismatch values: %d != %d\n", expectedListHead.Val, mergedPrintHead.Val)
			break
		}
		mergedPrintHead = mergedPrintHead.Next
		expectedListHead = expectedListHead.Next
	}
}
