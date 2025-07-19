package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	arrays := make([]int, 0, len(S))
	for i, v := range S {
		if v == '#' {
			arrays = append(arrays, i+1)
		}
	}

	for i := 0; i < len(arrays); i += 2 {
		fmt.Printf("%d,%d\n", arrays[i], arrays[i+1])
	}
}
