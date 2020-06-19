package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	var x string
	var ret bool

	ret = false
	x = "race a car"
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	ret = true
	x = ".,"
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	ret = true
	x = "A man, a plan, a canal: Panama"
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	ret = true
	x = "......a....."
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	ret = false
	x = "a.b,."
	fmt.Printf("x = %v ret = %v\n", x, isPalindrome(x))
	if isPalindrome(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}
