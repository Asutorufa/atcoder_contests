package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var array []int
	for i := 0; i < N-1; i++ {
		var a int
		fmt.Scan(&a)
		array = append(array, a)
	}

	var resp [][]int

	for i := 0; i < N-1; i++ {
		var r []int
		c := 0
		for j := i; j < N-1; j++ {
			c += array[j]
			r = append(r, c)
		}
		resp = append(resp, r)
	}

	for i := range resp {
		str := fmt.Sprintf("%v", resp[i])
		fmt.Println(str[1 : len(str)-1])
	}
}
