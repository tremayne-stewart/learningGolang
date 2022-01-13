package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	oddLength := (m+n)%2 == 1
	midMax := (m + n) / 2

	n1ptr := 0
	n2ptr := 0
	midPtr := 0
	midValue := 0
	prevMidValue := 0

	for midPtr <= midMax {
		var leftCompare int
		var rightCompare int

		if n1ptr < m {
			leftCompare = nums1[n1ptr]
		} else {
			leftCompare = math.MaxInt32
		}

		if n2ptr < n {
			rightCompare = nums2[n2ptr]
		} else {
			rightCompare = math.MaxInt32
		}

		prevMidValue = midValue
		if leftCompare < rightCompare {
			midValue = leftCompare
			n1ptr++
		} else {
			midValue = rightCompare
			n2ptr++
		}

		midPtr++
	}

	if oddLength {
		return float64(midValue)
	} else {
		return (float64(midValue) + float64(prevMidValue)) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := float64(2)
	fmt.Println(findMedianSortedArrays(nums1, nums2) == expected)

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	expected = 2.5
	fmt.Println(findMedianSortedArrays(nums1, nums2) == expected)

	nums1 = []int{1, 2}
	nums2 = []int{4, 5, 6, 7, 8}
	expected = 5
	fmt.Println(findMedianSortedArrays(nums1, nums2) == expected)
}
