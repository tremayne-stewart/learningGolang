// Author: Tremayne Stewart
// 10 Jan 2022
package main

func twoSum(nums []int, target int) []int {
	var memo = make(map[int]int)
	for i, val := range nums {
		entry, exists := memo[target-val]
		if exists {
			return []int{entry, i}
		} else {
			memo[val] = i
		}
	}
	return nil
}
