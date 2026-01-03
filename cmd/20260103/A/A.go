package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(int(math.Pow(2, float64(N)) - 2*float64(N)))
}
