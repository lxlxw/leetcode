package leetcode

import (
	"testing"
)

func TestCountPrimes(t *testing.T) {
	var ret int
	var n int
	ret = 4
	n = 10
	if ret != countPrimes(n) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 0
	n = 2
	if ret != countPrimes(n) {
		t.Fatalf("case fails %v\n", ret)
	}

}
