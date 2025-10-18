package main

import "fmt"

func main() {
	var S, A, B, X int
	fmt.Scan(&S, &A, &B, &X)

	c := X / (A + B)

	remain := X % (A + B)
	if remain > A {
		remain = A
	}

	fmt.Println((c*A + remain) * S)
}
