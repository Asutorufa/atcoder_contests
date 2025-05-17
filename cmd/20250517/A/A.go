package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scanln(&A, &B, &C, &D)

	if C < A {
		fmt.Println("Yes")
		return
	}

	if C <= A && D <= B {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
