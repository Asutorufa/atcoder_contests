package main

import "fmt"

func main() {
	var X, C int
	fmt.Scan(&X, &C)

	current := X / (1000 + C)
	fmt.Println(current * 1000)
}
