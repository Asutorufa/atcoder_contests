package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	count := 0
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scan(&A, &B)
		if B > A {
			count++
		}
	}

	fmt.Println(count)
}
