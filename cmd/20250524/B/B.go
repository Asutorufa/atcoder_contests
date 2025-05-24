package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	x := map[p]struct{}{}

	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i+j >= X {
				x[p{A: i, B: j}] = struct{}{}
			}
			if i-j >= Y || j-i >= Y {
				x[p{A: i, B: j}] = struct{}{}
			}
		}
	}

	fmt.Println(float64(len(x)) / 36)
}

type p struct {
	A, B int
}
