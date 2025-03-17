package main

import (
	"fmt"
	"strconv"
)

func main() {
	var AA string
	fmt.Scan(&AA)

	AA = AA[:4]

	A, _ := strconv.ParseFloat(AA, 64)

	if A >= 38 {
		fmt.Println(1)
	} else if A >= 37.5 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}
