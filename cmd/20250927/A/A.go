package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	sum := 0
	for i := 1; i <= N; i++ {

		sum += int(math.Pow(-1, float64(i))) * (int(math.Pow(float64(i), 3)))
	}

	fmt.Println(sum)
}
