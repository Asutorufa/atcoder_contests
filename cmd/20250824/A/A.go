package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	z := (X + Y) % 12
	if z == 0 {
		z = 12
	}
	fmt.Println(z)
}
