package leetcode

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	var ret int
	var num int

	num = 38
	ret = 2
	if ret != addDigits(num) {
		t.Fatalf("case fails %v\n", ret)
	}
}
