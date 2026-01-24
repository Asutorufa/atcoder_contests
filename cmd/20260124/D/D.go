package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	Q := readInt()

	array := make([]int, N)
	prefixArray := make([]int, N+1)

	for i := range N {
		array[i] = readInt()
		prefixArray[i+1] = array[i] + prefixArray[i]
	}

	for range Q {
		cmd := readInt()
		switch cmd {
		case 1:
			x := readInt() - 1
			prefixArray[x+1] = prefixArray[x+1] - (array[x] - array[x+1])
			array[x], array[x+1] = array[x+1], array[x]
		case 2:
			l := readInt() - 1
			r := readInt() - 1
			fmt.Println(prefixArray[r+1] - prefixArray[l])
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
