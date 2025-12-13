package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	cache := map[Point]bool{}

	ret := 0
_next:
	for range M {
		var x, y int
		fmt.Scan(&x, &y)
		ps := []Point{{x, y}, {x, y + 1}, {x + 1, y}, {x + 1, y + 1}}
		for _, v := range ps {
			if cache[v] {
				continue _next
			}
		}

		for _, v := range ps {
			cache[v] = true
		}

		ret++
	}

	fmt.Println(ret)
}
