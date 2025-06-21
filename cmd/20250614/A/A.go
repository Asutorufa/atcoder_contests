package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	arrays := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arrays[i])
	}

	var K int
	fmt.Scan(&K)

	count := 0
	for i := 0; i < N; i++ {
		if arrays[i] >= K {
			count++
		}
	}

	fmt.Println(count)
}
