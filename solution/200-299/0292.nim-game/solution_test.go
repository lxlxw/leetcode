package leetcode

import (
	"testing"
)

func TestCanWinNim(t *testing.T) {
	var ret bool
	var n int

	n = 30
	ret = true
	if ret != canWinNim(n) {
		t.Fatalf("case fails %v\n", ret)
	}

	n = 40
	ret = false
	if ret != canWinNim(n) {
		t.Fatalf("case fails %v\n", ret)
	}
}
