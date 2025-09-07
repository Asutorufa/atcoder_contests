package main

import "testing"

func TestC(t *testing.T) {
	t.Log(calc(2, 5, 2), 2)
	t.Log(calc(2, 2, 5), 2)
	t.Log(calc(2, 2, 2), 2)
	t.Log(calc(2, 2, 1), 1)
	t.Log(calc(1, 1, 2), 1)

	t.Log(calc(0, 1, 2), 0)
	t.Log(calc(1, 0, 2), 1)
	t.Log(calc(1, 1, 0), 0)

	t.Log(calc(22, 1, 10), 10)

	t.Log(calc(22, 1, 22), 12)

	t.Log(calc(22, 1, 50), 22)

	t.Log(calc(40, 0, 21), 20)

	t.Log(calc(21, 0, 21), 11)

	t.Log(calc(21, 2, 21), 12)

	t.Log(calc(5, 1, 4), 3)
}
