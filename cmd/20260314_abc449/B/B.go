package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	H, W, Q := readX[int](), readX[int](), readX[int]()

	for range Q {
		c, r := readInt(), readInt()

		count := 0
		switch c {
		case 1:
			count = W * r
			H -= r

		case 2:
			count = H * r
			W -= r
		}

		fmt.Println(count)
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
