package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	cache := map[int]int{}

	for x := 1; x <= N; x++ {
		for y := x + 1; y <= N; y++ {
			n := x*x + y*y
			if n > N {
				break
			}

			cache[n] = cache[n] + 1
		}
	}

	array := []int{}
	for k, v := range cache {
		if v == 1 {
			array = append(array, k)
		}
	}

	sort.Ints(array)

	fmt.Println(len(array))
	fmt.Println(strings.Trim(fmt.Sprint(array), "[]"))
}
