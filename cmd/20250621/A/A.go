package main

import "fmt"

func main() {
	var P string
	fmt.Scan(&P)

	var L int
	fmt.Scan(&L)

	if len(P) < L {
		fmt.Println("No")
		return
	}

	for _, v := range P {
		if !(v >= 'a' && v <= 'z') {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

/*
atc0der
7

*/
