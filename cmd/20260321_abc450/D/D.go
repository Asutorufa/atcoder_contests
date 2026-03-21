package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	N, K := readInt(), readInt()

	array := make([]int, 0, N)

	for range N {
		array = append(array, readInt()%K)
	}

	sort.Ints(array)

	ret := array[N-1] - array[0]

	for i := 0; i < N-1; i++ {
		n := array[i] + K - array[i+1]
		if n < ret {
			ret = n
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
