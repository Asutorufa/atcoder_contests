package main

import "fmt"

func main() {
	var counter = map[int]int{}
	for i := 0; i < 7; i++ {
		var a int
		fmt.Scan(&a)

		counter[a] = counter[a] + 1
	}

	three, two := false, false

	for _, v := range counter {
		if !three && v >= 3 {
			three = true
		} else if !two && v >= 2 {
			two = true
		}

		if three && two {
			break
		}
	}

	if three && two {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
