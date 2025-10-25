package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var arrays []int
	var sum int
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		arrays = append(arrays, a)
		sum += a
	}

	for _, v := range arrays {
		if sum-v == M {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
