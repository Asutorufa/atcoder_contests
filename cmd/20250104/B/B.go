package main

import "fmt"

func main() {

	var X int
	fmt.Scan(&X)

	sum := 0
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			z := i * j
			if z != X {
				sum += z
			}
		}
	}

	fmt.Println(sum)
}
