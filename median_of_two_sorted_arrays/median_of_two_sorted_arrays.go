// Author: Tremayne Stewart
// 12 Jan 2022
package main

import (
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
