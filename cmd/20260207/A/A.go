package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	S := readString()

	first := S[0]

	for i := range S {
		if S[i] != first {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
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
