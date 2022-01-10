package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root = new(ListNode)
	head := root
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			head.Next = list2
			list2 = list2.Next
		} else {
			head.Next = list1
			list1 = list1.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	} else {
		head.Next = list2
	}
	return root.Next
}

func main() {
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
			fmt.Printf("Mismatch values: %d != %d\n", expectedListHead.Val, mergedPrintHead.Val)
			break
		}
		fmt.Println(mergedPrintHead.Val)
		mergedLengthCounter++
		mergedPrintHead = mergedPrintHead.Next
		if expectedListHead.Next != nil {
			expectedListHead = expectedListHead.Next
		}
	}
	if expectedListLength != mergedLengthCounter {
		fmt.Printf("Expected list and merged list are different lenghts: %d != %d", expectedListLength, mergedLengthCounter)
	}
}
