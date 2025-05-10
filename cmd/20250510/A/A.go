package main

import "fmt"

func main() {
	var R, X int
	fmt.Scanln(&R, &X)

	switch X {
	case 1:
		if R >= 1600 && R <= 2999 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	case 2:
		if R >= 1200 && R <= 2399 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
