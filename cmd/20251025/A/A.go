package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	for i := 1; i <= N; i++ {
		if i > M {
			fmt.Println("Too Many Requests")
		} else {
			fmt.Println("OK")
		}
	}
}
