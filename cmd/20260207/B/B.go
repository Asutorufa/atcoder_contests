package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	K := readInt()

	ret := 0
	for i := 1; i <= N; i++ {
		if sum(i) == K {
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

func sum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}
