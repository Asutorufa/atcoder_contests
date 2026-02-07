package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	N := readInt()

	maxA := 0

	cnt := make([]int, 1000005)
	for range N {
		A := readInt()
		cnt[A]++
		maxA = max(maxA, A)
	}

	digits := make([]int, 1000005)
	currentSum := 0

	for i := maxA; i >= 1; i-- {
		currentSum += cnt[i]
		digits[i-1] = currentSum
	}

	carry := 0
	var str bytes.Buffer
	for i := 0; i < maxA || carry > 0; i++ {
		val := carry + digits[i]
		str.WriteByte(byte(val%10) + '0')
		carry = val / 10
	}

	slices.Reverse(str.Bytes())
	fmt.Println(str.String())
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
