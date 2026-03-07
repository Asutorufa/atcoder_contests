package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	M := readInt()

	ks := make([]int, 0, M)
	for range M {
		K := readInt()
		ks = append(ks, K)
	}

	sum := 0

	for range N {
		A := readInt() - 1
		B := readInt()

		re := ks[A]

		if re > 0 {
			if re >= B {
				sum += B
				ks[A] = re - B
			} else {
				sum += re
				ks[A] = 0
			}
		}

	}

	fmt.Println(sum)
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
