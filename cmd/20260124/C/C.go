package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	M := readInt()

	Ns := make([]map[int]bool, N)
	for i := range N {
		Ns[i] = make(map[int]bool)
	}

	for range M {
		a := readInt() - 1
		b := readInt() - 1
		Ns[a][b] = true
		Ns[b][a] = true
	}

	// fmt.Println(Ns)

	for i := range N {
		c := N - len(Ns[i]) - 1
		if c < 3 {
			fmt.Println(0)
		} else {
			fmt.Println((c * (c - 1) * (c - 2)) / 6)
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

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
