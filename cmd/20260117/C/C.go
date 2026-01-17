package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	N := readInt()
	K := readInt()
	X := readInt()

	as := make([]int, N)
	for i := range N {
		as[i] = readInt()
	}

	sort.Slice(as, func(i, j int) bool {
		return as[i] > as[j]
	})

	// fmt.Println(as)

	mizu := N - K

	ret := 0
	sum := 0
	for _, v := range as[mizu:] {
		ret++
		sum += v

		if sum >= X {
			break
		}
	}

	if sum >= X {
		fmt.Println(ret + mizu)
	} else {
		fmt.Println(-1)
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
