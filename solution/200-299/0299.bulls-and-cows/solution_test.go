package leetcode

import (
	"testing"
)

func TestGetHint(t *testing.T) {
	var secret string
	var guess string
	secret = "1123"
	guess = "0111"
	getHint(secret, guess)

	secret = "011123112"
	guess = "012321544"
	getHint(secret, guess)
}
