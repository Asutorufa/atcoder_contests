package main

import (
	"container/list"
	"testing"
)

func TestNode(t *testing.T) {
	node := NewNode()

	A1 := map[int]struct{}{1: {}}
	A2 := map[int]struct{}{2: {}}
	node.Add(A1)
	node.Add(A2)

	node.Range(func(z *list.Element) bool {
		t.Logf("%v", z.Value)
		return true
	})

	node.Remove(node.list.Front())

	node.Range(func(z *list.Element) bool {
		t.Logf("%v", z.Value)
		return true
	})
}
