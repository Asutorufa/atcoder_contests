package main

import (
	"testing"

	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/stretchr/testify/assert"
)

func TestRBTree(t *testing.T) {
	rbt := redblacktree.NewWithIntComparator()

	rbt.Put(1, 1)
	rbt.Put(3, 3)
	rbt.Put(119999, 11111)
	rbt.Put(11111, 11111)
	rbt.Put(40, 40)
	rbt.Put(100, 100)
	rbt.Put(20, 20)

	fn, ok := rbt.Floor(4)
	assert.True(t, ok)

	iter := rbt.IteratorAt(fn)

	t.Log(fn.Key)
	for iter.Next() {
		t.Log(iter.Key())
	}

	iter = rbt.IteratorAt(fn)
	for iter.Prev() {
		t.Log(iter.Key())
	}
}
