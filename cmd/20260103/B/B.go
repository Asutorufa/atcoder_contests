package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	cache := map[int]bool{}

	for !cache[N] {
		cache[N] = true

		sum := 0

		for N > 0 {
			c := N % 10
			// fmt.Println(c)
			sum += c * c
			N = N / 10
		}
		// fmt.Println("sum", sum)

		if sum == 1 {
			fmt.Println("Yes")
			return
		}

		N = sum
	}

	fmt.Println("No")
}
