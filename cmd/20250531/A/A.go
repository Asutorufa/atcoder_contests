package main

import "fmt"

func main() {
	var N, S int
	fmt.Scan(&N, &S)

	last := 0
	for i := 1; i <= N; i++ {
		var T int
		fmt.Scan(&T)

		if T-last > S {
			fmt.Println("No")
			return
		}

		last = T
	}

	fmt.Println("Yes")
}
