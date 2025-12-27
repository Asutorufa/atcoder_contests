package main

import "fmt"

func main() {
	var D, F int
	fmt.Scan(&D, &F)
	fmt.Println(7 - ((D - F) % 7))
}
