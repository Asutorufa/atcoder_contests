package main

import "fmt"

func main() {
	var A, B int
	fmt.Scanln(&A, &B)
	fmt.Println((A + B) * (A + B))
}
