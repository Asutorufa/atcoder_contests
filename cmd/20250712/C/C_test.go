package main

import "testing"

func TestC(t *testing.T) {
	n := int64(1000000)
	k := int64(2)
	t.Log(sumPalindrome(n, k))
}
