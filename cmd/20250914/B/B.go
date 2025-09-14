package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	L := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&L[i])
	}

	heya := make([]bool, N+1)

	heya[0] = true
	heya[N] = true

	for i := 0; i < N; i++ {
		if L[i] == 1 {
			break
		}

		heya[i+1] = true
	}

	for i := N; i > 0; i-- {
		if L[i-1] == 1 {
			break
		}

		heya[i-1] = true
	}

	count := 0
	for i := range heya {
		if !heya[i] {
			// fmt.Println(i)
			count++
		}
	}

	fmt.Println(count)
}
