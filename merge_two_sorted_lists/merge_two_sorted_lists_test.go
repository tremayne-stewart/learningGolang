// Author: Tremayne Stewart
// 12 Jan 2022
package main

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	// list1 = [1,2,4]
	var list1 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	// list2 = [1,2,4]
	var list2 = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	// expectedList = [1,1,2,3,4]
	var expectedList = ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	}
	expectedListHead := &expectedList
	const expectedListLength = 6

	mergedRoot := mergeTwoLists(&list1, &list2)
	mergedLengthCounter := 0
	mergedPrintHead := mergedRoot
	for mergedPrintHead != nil {
		if expectedListHead.Val != mergedPrintHead.Val {
			t.Fatalf("Mismatch values: %d != %d\n", expectedListHead.Val, mergedPrintHead.Val)
			break
		}
		mergedLengthCounter++
		mergedPrintHead = mergedPrintHead.Next
		if expectedListHead.Next != nil {
			expectedListHead = expectedListHead.Next
		}
	}
	if expectedListLength != mergedLengthCounter {
		t.Fatalf("Expected list and merged list are different lenghts: %d != %d", expectedListLength, mergedLengthCounter)
	}
}
