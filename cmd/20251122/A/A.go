package main

import "fmt"

func main() {
	var X, Y, Z int
	fmt.Scan(&X, &Y, &Z)

	if X < Y {
		fmt.Println("No")
		return
	}

	for x := range 1000 {
		xx := X + x
		yy := Y + x

		if xx%yy == 0 && xx/yy == Z {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
