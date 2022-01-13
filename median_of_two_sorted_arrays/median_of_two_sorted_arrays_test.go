// Author: Tremayne Stewart
// 12 Jan 2022
package main

import (
	"testing"
)

func TestOverlappingArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := float64(2)
	if expected != findMedianSortedArrays(nums1, nums2) {
		t.Fatal("doesn't work")
	}
}

func TestEvenDisjointArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expected := float64(2.5)
	if expected != findMedianSortedArrays(nums1, nums2) {
		t.Fatal("doesn't work")
	}
}

func TestDisjointArrays(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{4, 5, 6, 7, 8}
	expected := float64(5)
	if expected != findMedianSortedArrays(nums1, nums2) {
		t.Fatal("doesn't work")
	}
}
