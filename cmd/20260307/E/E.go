package main

import (
	"bufio"
	"fmt"
	"os"
)

// TLE。。。 分からないです。。
func main() {
	K := readInt()
	M := readInt()

	const MOD = 10007
	remainder := 0
	quotientMod := 0

	for range K {
		digit := readInt()
		l := readInt()

		for range l {
			current := remainder*10 + digit
			stepQuotient := current / M
			remainder = current % M

			quotientMod = (quotientMod*10 + stepQuotient) % MOD
		}
	}

	fmt.Println(quotientMod)
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
