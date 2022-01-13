// Author: Tremayne Stewart
// 10 Jan 2022
package main

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
