package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var S, T string
	fmt.Scan(&S, &T)

	ii := 0

	for i, v := range S {
		if byte(v) != T[i] {
			ii++
		}
	}

	fmt.Println(ii)
}
