package main

import "testing"

func TestXxx(t *testing.T) {
	t.Log(string(replaceBeforeAll([]byte("WWWWWWA"), 5)))
}
