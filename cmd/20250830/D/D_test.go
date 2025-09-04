package main

import "testing"

func TestSplit(t *testing.T) {
	A := []Step{
		{"A", 100},
		Step{"B", 100},
	}

	B := []Step{
		{"A", 50},
		{"B", 50},
		{"C", 70},
		{"D", 50},
		{"E", 50},
	}

	t.Log(Split(A, B))
}
