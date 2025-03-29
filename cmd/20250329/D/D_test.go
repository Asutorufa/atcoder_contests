package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	dd := `
5
1 2 3 4 5 1 2 3 4 5
	`

	solveOne2(bufio.NewReader(strings.NewReader(dd)))
}
