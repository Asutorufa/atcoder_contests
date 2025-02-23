package main

import "testing"

func TestXxx(t *testing.T) {
	t.Log(string(replaceAll([]byte("([])<>()"), 1)))
	t.Log(string(replaceAll([]byte("<>()"), 0)))
}
