// Author: Tremayne Stewart
// 12 Jan 2022
package main

import "testing"

func areArraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	expected := []int{0, 1}
	if !areArraysEqual(expected, twoSum(nums, 9)) {
		t.Fatal(expected, twoSum(nums, 9))
	}
}
