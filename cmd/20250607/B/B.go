package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	max := 0
	for x := 0; x <= N; x++ {
		count := 0
		for _, v := range A {
			if v >= x {
				count++
			}
		}
		if count >= x {
			max = x
		}
	}

	fmt.Println(max)
}
