package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	max := math.Pow10(9)

	sum := 0.0
	for i := 0; i <= M; i++ {
		sum += math.Pow(float64(N), float64(i))

		if sum > max {
			fmt.Println("inf")
			return
		}
	}

	fmt.Println(int(sum))
}
