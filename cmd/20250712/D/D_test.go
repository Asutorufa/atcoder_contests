package main

import (
	"fmt"
	"testing"
)

func TestSplit(t *testing.T) {
	points := []int{5, 8, 10, 14, 15, 15, 20}
	k := 3
	result := minTotalSpan(points, k)
	fmt.Println(result)

	points = []int{5, 10, 15, 20, 8, 14, 15}
	k = 2
	result = minTotalSpan(points, k)

	fmt.Println(result)
}
