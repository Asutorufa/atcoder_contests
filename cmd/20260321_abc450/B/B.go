package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()

	cost := make([][]int, N)
	for i := range N {
		cost[i] = make([]int, N)
	}

	for i := range N {
		for j := i + 1; j < N; j++ {
			cost[i][j] = readInt()
		}
	}

	// fmt.Println(cost)

	for i := range N {
		for j := i + 1; j < N; j++ {
			dcost := cost[i][j]

			for z := i + 1; z < j; z++ {
				// fmt.Println(i, j, z, dcost, cost[i][z]+cost[z][j])
				if cost[i][z]+cost[z][j] < dcost {
					fmt.Println("Yes")
					return
				}
			}
		}
	}

	fmt.Println("No")
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
