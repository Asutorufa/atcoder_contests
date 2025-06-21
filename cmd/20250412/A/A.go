package main

import "fmt"

func main() {
	var S int
	fmt.Scan(&S)

	if S >= 200 && S <= 299 {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}
