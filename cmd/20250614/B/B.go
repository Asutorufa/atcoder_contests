package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)

	hakos := make([]int, N)

	result := make([]int, 0, N)
	for i := 0; i < Q; i++ {
		var X int
		fmt.Scan(&X)

		if X >= 1 {
			hakos[X-1] += 1
			result = append(result, X)
		} else {
			ii := findMinimum(hakos)
			hakos[ii] += 1
			result = append(result, ii+1)
		}
	}

	str := fmt.Sprint(result)
	str = str[1 : len(str)-1]
	fmt.Println(str)
}

func findMinimum(arr []int) int {
	min := arr[0]
	index := 0

	for i, v := range arr {
		if v < min {
			min = v
			index = i
		}
	}

	return index
}
