package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scanln(&A, &B, &C)

	if A == B && B == C {
		fmt.Println("Yes")
		return
	}

	if (A+B == C) || (A+C == B) || (B+C == A) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
