package main

import "fmt"

func main() {
	var X, Y string
	fmt.Scan(&X, &Y)

	maps := map[string]int{
		"Ocelot": 1,
		"Serval": 2,
		"Lynx":   3,
	}

	if maps[X] >= maps[Y] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
