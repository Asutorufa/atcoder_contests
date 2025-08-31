package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	h := make([]string, 0, N)
	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)

		h = append(h, S)
	}

	var X int
	var Y string
	fmt.Scan(&X, &Y)

	X--

	if h[X] == Y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
