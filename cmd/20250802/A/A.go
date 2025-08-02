package main

import "fmt"

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)

	var S string
	fmt.Scan(&S)

	fmt.Println(S[A : len(S)-B])
}
