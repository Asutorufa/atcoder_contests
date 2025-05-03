package main

import "testing"

func TestRotate(t *testing.T) {
	var s = []string{
		"###.",
		"..#.",
		"..#.",
		"..#.",
	}

	for _, v := range rotate270(s) {
		t.Log(v)
	}

	for _, v := range rotate180(s) {
		t.Log(v)
	}

	for _, v := range rotate90(s) {
		t.Log(v)
	}
}
