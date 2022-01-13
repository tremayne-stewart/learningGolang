// Author: Tremayne Stewart
// 10 Jan 2022
package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	dest := 0
	for dest < x {
		dest = dest*10 + x%10
		x /= 10
	}
	return x == dest || x == dest/10
}
