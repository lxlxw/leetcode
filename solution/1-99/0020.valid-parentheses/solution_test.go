package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {

	var s string
	var ret bool
	s = "()"
	ret = true
	fmt.Printf("x = %v ret = %v\n", s, isValid(s))
	if isValid(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "(]"
	ret = false
	fmt.Printf("x = %v ret = %v\n", s, isValid(s))
	if isValid(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = "{[]}"
	ret = true
	fmt.Printf("x = %v ret = %v\n", s, isValid(s))
	if isValid(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}

	s = ""
	ret = true
	fmt.Printf("x = %v ret = %v\n", s, isValid(s))
	if isValid(s) != ret {
		t.Fatalf("case fails: %v\n", ret)
	}
}
