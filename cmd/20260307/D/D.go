package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	N := readInt()
	numbers := make([]int, 0, N)
	for range N {
		numbers = append(numbers, readInt())
	}

	paths := make([][]int, N)
	for range N - 1 {
		U := readInt() - 1
		V := readInt() - 1
		paths[U] = append(paths[U], V)
		paths[V] = append(paths[V], U)
	}

	ans := make([]bool, N)
	currentCount := make(map[int]int)
	globalCount := 0

	var dfs func(u int, parent int)
	dfs = func(u int, parent int) {
		val := numbers[u]
		currentCount[val]++
		if currentCount[val] == 2 {
			globalCount++
		}

		if globalCount > 0 {
			ans[u] = true
		}

		for _, v := range paths[u] {
			if v != parent {
				dfs(v, u)
			}
		}

		if currentCount[val] == 2 {
			globalCount--
		}
		currentCount[val]--
	}

	dfs(0, -1)

	for i := range N {
		if ans[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
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

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
