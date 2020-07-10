package leetcode

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var ss string
	var ts string
	var ret bool

	ss = "dmomkifm"
	ts = "skmokj"
	ret = false

	if ret != isAnagram(ss, ts) {
		t.Fatalf("case fails %v\n", ret)
	}

	ss = "ddss"
	ts = "ssdd"
	ret = true

	if ret != isAnagram(ss, ts) {
		t.Fatalf("case fails %v\n", ret)
	}

	ss = "ddsse"
	ts = "ssddd"
	ret = false

	if ret != isAnagram(ss, ts) {
		t.Fatalf("case fails %v\n", ret)
	}
}
