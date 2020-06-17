package leetcode

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {

	var x, ret int
	x = 123
	ret = 321
	fmt.Printf("x = %v ret = %v\n", x, reverse(x))
	if reverse(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	x = -123
	ret = -321
	fmt.Printf("x = %v ret = %v\n", x, reverse(x))
	if reverse(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	x = 120
	ret = 21
	fmt.Printf("x = %v ret = %v\n", x, reverse(x))
	if reverse(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	x = 1 << 31
	ret = 0
	fmt.Printf("x = %v ret = %v\n", x, reverse(x))
	if reverse(x) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}
