package main

import "testing"

func TestC(t *testing.T) {
	t.Log(Resolve(3, "0010000"))
	t.Log(Resolve(3, "0010110"))
	t.Log(Resolve(1, "1"))
	t.Log(Resolve(2, "100"))
	t.Log(Resolve(4, "001110010101110"))
}

func TestBC(t *testing.T) {
	b := NewBitCalculator()

	b.AddBit(0)
	t.Log(b.Value())
	b.AddBit(1)
	t.Log(b.Value())
	b.AddBit(2)
	t.Log(b.Value())
	b.Undo(2)
	t.Log(b.Value())
	b.Undo(1)
	t.Log(b.Value())
}
