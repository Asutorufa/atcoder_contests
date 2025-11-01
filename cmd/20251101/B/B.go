package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	var A []string
	for range N {
		var s string
		fmt.Scan(&s)
		A = append(A, s)
	}

	maps := map[string]int{}

	for x := 0; x <= N-M; x++ {
		for y := 0; y <= N-M; y++ {
			arrays := []string{}
			for j := y; j < y+M; j++ {
				arrays = append(arrays, A[j][x:x+M])
			}

			maps[strings.Join(arrays, "-")]++
		}
	}

	// fmt.Println(maps)

	fmt.Println(len(maps))
}
