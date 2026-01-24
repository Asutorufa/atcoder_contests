package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := readString()

	ret := 0
	for _, v := range s {
		if v == 'i' || v == 'j' {
			ret++
		}
	}

	fmt.Println(ret)
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
