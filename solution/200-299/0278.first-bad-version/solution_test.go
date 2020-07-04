package leetcode

import (
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	var ret int
	var num int

	num = 5
	ret = 4
	if ret != firstBadVersion(num) {
		t.Fatalf("case fails %v\n", ret)
	}
}
