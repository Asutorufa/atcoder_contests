package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	X := readInt()

	for range N {
		A := readInt()

		if A < X {
			X = A
			fmt.Println("1")
		} else {
			fmt.Println("0")
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
