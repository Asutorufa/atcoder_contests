package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	if C >= A && D < B {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
