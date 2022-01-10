package main

import (
	"fmt"
	"log"
)

func twoSum(nums []int, target int) []int {
	var memo = make(map[int]int)
	for i, val := range nums {
		log.Println(i, val)
		entry, exists := memo[target-val]
		if exists {
			return []int{entry, i}
		} else {
			memo[val] = i
		}
	}
	return nil
}

func main() {
	log.SetFlags(0)
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
