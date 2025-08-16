package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	switch S {
	case "red":
		fmt.Println("SSS")
	case "blue":
		fmt.Println("FFF")
	case "green":
		fmt.Println("MMM")
	default:
		fmt.Println("Unknown")
	}
}
