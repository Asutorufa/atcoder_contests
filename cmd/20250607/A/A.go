package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var T, A string
	fmt.Scan(&T, &A)

	for i := 0; i < N; i++ {
		if T[i] != 'o' {
			continue
		}

		if T[i] == A[i] {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
