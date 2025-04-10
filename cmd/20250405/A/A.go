package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if 400%N == 0 {
		fmt.Println(400 / N)
	} else {
		fmt.Println(-1)
	}
}
