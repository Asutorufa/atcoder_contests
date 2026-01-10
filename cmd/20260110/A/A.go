package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	for range Y {
		X *= 2
	}

	fmt.Println(X)
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
