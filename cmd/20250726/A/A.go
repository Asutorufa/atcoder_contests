package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, L, R int
	fmt.Scan(&N, &L, &R)

	var S string
	fmt.Scan(&S)

	L--
	S = S[L:R]

	// fmt.Println(S)

	if strings.Contains(S, "x") {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
