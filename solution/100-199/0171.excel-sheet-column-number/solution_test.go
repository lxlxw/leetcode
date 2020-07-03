package leetcode

import (
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	var ret int
	var s string
	ret = 1
	s = "A"
	if ret != titleToNumber(s) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 28
	s = "AB"
	if ret != titleToNumber(s) {
		t.Fatalf("case fails %v\n", ret)
	}

	ret = 701
	s = "ZY"
	if ret != titleToNumber(s) {
		t.Fatalf("case fails %v\n", ret)
	}

}
