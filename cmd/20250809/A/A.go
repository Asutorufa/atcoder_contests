package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	if strings.HasSuffix(S, "tea") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
