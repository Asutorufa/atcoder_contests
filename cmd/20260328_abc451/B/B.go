package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	M := readInt()

	mn := make(map[int]int)
	ml := make(map[int]int)

	for range N {
		A, B := readX[int](), readX[int]()
		A--
		B--
		mn[A] += 1
		ml[B] += 1
	}

	for i := range M {
		fmt.Println(ml[i] - mn[i])
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
