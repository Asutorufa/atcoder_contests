package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := readString()

	if len(s)%5 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
