package leetcode

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	var x int
	var ret int
	x = 8
	ret = 2
	if ret != mySqrt(x) {
		t.Fatalf("case fails %v\n", ret)
	}

	x = 1
	ret = 1
	if ret != mySqrt(x) {
		t.Fatalf("case fails %v\n", ret)
	}

	x = 101
	ret = 10
	if ret != mySqrt(x) {
		t.Fatalf("case fails %v\n", ret)
	}
}
