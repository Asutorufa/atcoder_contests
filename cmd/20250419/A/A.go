package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	for _, v := range S {
		if v >= 'A' && v <= 'Z' {
			fmt.Print(string(v))
		}
	}
	fmt.Println()
}
