package leetcode

import (
	"testing"
)

func TestGenerate(t *testing.T) {

	var target int
	var ret [][]int

	ret = [][]int{}
	target = 0
	for k, nums := range generate(target) {
		for i, v := range nums {
			if v != ret[k][i] {
				t.Fatalf("case fails: %v\n", ret)
			}
		}
	}

	ret = [][]int{{1}}
	target = 1
	for k, nums := range generate(target) {
		for i, v := range nums {
			if v != ret[k][i] {
				t.Fatalf("case fails: %v\n", ret)
			}
		}
	}

	ret = [][]int{{1}, {1, 1}}
	target = 2
	for k, nums := range generate(target) {
		for i, v := range nums {
			if v != ret[k][i] {
				t.Fatalf("case fails: %v\n", ret)
			}
		}
	}

	ret = [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}
	target = 4
	for k, nums := range generate(target) {
		for i, v := range nums {
			if v != ret[k][i] {
				t.Fatalf("case fails: %v\n", ret)
			}
		}
	}
}
