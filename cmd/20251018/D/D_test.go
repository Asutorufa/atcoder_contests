package main

import (
	"math"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Log(math.Ceil(math.Sqrt(10)))
	t.Log(math.Floor(math.Sqrt(101)))

	t.Log(count(45, 49))
	t.Log(count(9, 9))

	t.Log(solve(4, 80))
	t.Log(solve(183, 5000))
	t.Log(solve(18, 10))
	t.Log(solve(824, 5000000000))

	t.Log(solve(49, 70))
}

func TestCount(t *testing.T) {
	t.Log(count(45, 49))
	t.Log(count(9, 9))

}
