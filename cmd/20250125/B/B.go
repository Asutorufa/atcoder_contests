package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	var z []int64
	readIntFunc(N, br, func(x int, index int) {
		z = append(z, int64(x))
	})

	last := z[1]
	ration := big.NewRat(z[1], z[0])

	for _, x := range z[2:] {
		if big.NewRat(x, last).Cmp(ration) != 0 {
			fmt.Println("No")
			return
		}
		last = x
	}

	fmt.Println("Yes")
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}
