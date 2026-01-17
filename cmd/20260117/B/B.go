package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readInt()
	readInt()

	s := readString()
	t := readString()

	sm := make(map[rune]bool, len(s))
	tm := make(map[rune]bool, len(t))

	for _, c := range s {
		sm[c] = true
	}
	for _, c := range t {
		if sm[c] {
			delete(sm, c)
		} else {
			tm[c] = true
		}
	}

	Q := readInt()
	for i := 0; i < Q; i++ {
		c := readString()

		true1, true2 := false, false

		for _, v := range c {
			if sm[v] {
				true1 = true
			} else if tm[v] {
				true2 = true
			}
		}

		if true1 {
			fmt.Println("Takahashi")
		} else if true2 {
			fmt.Println("Aoki")
		} else {
			fmt.Println("Unknown")
		}
	}
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}
