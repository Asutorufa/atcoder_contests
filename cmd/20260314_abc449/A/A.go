package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	d := readX[float64]() / 2

	fmt.Println(d * d * math.Pi)
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
