package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(br, &Q)

	current := []int{}

	for i := 0; i < Q; i++ {
		var T int
		fmt.Fscan(br, &T)

		switch T {
		case 1:
			var C string
			fmt.Fscan(br, &C)

			switch C {
			case "(":
				if len(current) == 0 {
					current = append(current, 1)
				} else {
					if current[len(current)-1] >= 0 {
						current = append(current, current[len(current)-1]+1)
					} else {
						current = append(current, current[len(current)-1])
					}
				}
			case ")":
				if len(current) == 0 {
					current = append(current, -1)
				} else {
					current = append(current, current[len(current)-1]-1)
				}
			}

		case 2:
			current = current[:len(current)-1]
		}

		// fmt.Println(current)
		if len(current) == 0 || current[len(current)-1] == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
