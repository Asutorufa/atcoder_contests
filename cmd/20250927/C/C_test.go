package main

import "testing"

func TestCircularPrefixSum(t *testing.T) {
	c := NewCircularPrefixSum([]int{3, 1, 4, 5})

	t.Log(c.rangeSum(0, 2))
	c.rotateLeft(1)
	t.Log(c.rangeSum(1, 2))
}
