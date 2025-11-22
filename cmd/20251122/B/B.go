package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var arrays = make([]int, 0, N)
	for range N {
		var S int
		fmt.Scan(&S)
		arrays = append(arrays, S)
	}

	for i1, v1 := range arrays {
		A := -1
		for i2, v2 := range arrays {
			if i1 == i2 {
				break
			}

			if v2 <= v1 {
				continue
			}

			A = i2 + 1
		}

		fmt.Println(A)
	}
}
