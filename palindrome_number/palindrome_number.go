package main

import "fmt"

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

func testIsPalindrome(value int, expected bool) {
	isPalindromeResult := isPalindrome(value)
	if isPalindromeResult != expected {
		fmt.Printf("Expected value for %d (%v) is incorrect\n", value, expected)
		return
	}
	resultStr := "is"
	if isPalindromeResult == false {
		resultStr = "is not"
	}
	fmt.Printf("%d %v a palindrome\n", value, resultStr)
}

func main() {
	testIsPalindrome(10, true) // test value is incorrect.
	testIsPalindrome(10, false)
	testIsPalindrome(121, true)
	testIsPalindrome(1221, true)
	testIsPalindrome(12421, true)
	testIsPalindrome(100, false)
	testIsPalindrome(102, true)
	testIsPalindrome(101, true)
}
