package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	var x int
	var ret bool = true

	x = 121
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	x = 123
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) == ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	x = 1221
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}
