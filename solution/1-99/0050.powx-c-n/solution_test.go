package leetcode

import (
	"testing"
)

func TestMyPow(t *testing.T) {
	var ret float64
	var x float64
	var n int
	ret = 0
	x = 0.00001
	n = 2147483647
	if ret != myPow(x, n) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 1.0000
	x = 2
	n = 0
	if ret != myPow(x, n) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 10.0000
	x = 10.0000
	n = 1
	if ret != myPow(x, n) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 0.25000
	x = 2.0000
	n = -2
	if ret != myPow(x, n) {
		t.Fatalf("case fails %v\n", ret)
	}
}
