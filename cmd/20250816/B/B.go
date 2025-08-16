package main

import (
	"fmt"
	"sort"
)

func main() {
	var Q int
	fmt.Scan(&Q)

	var array []int

	for i := 0; i < Q; i++ {
		var T int
		fmt.Scan(&T)

		switch T {
		case 1:
			var X int
			fmt.Scan(&X)

			array = append(array, X)
		case 2:
			if len(array) == 0 {
				continue
			}

			sort.Ints(array)
			fmt.Println(array[0])
			array = array[1:]
		}
	}
}
