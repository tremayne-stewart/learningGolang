// Author: Tremayne Stewart
// 12 Jan 2022
package main

import (
	"testing"
)

func testIsPalindromFactory(t *testing.T) func(int, bool) {
	return func(value int, expected bool) {
		isPalindromeResult := isPalindrome(value)
		if isPalindromeResult != expected {
			t.Fatalf("Expected value for %d (%v) is incorrect\n", value, expected)
		}
	}
}

func TestIsPalindromeMultipleValues(t *testing.T) {
	testIsPalindrome := testIsPalindromFactory(t)
	testIsPalindrome(10, false)
	testIsPalindrome(121, true)
	testIsPalindrome(1221, true)
	testIsPalindrome(12421, true)
	testIsPalindrome(100, false)
	testIsPalindrome(102, false)
	testIsPalindrome(101, true)
}
