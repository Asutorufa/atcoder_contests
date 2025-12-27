package main

import (
	"fmt"
	"math"
)

func main() {
	var N, M int
	var S, T string
	fmt.Scan(&N, &M, &S, &T)

	ret := math.MaxInt
	for i := range N - M + 1 {
		t := S[i : i+M]

		steps := 0
		for i, v := range T {
			steps += step(maps[byte(v)], maps[t[i]])
		}

		ret = min(ret, steps)
	}

	fmt.Println(ret)
}

var maps = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

// x->y
func step(x, y int) int {
	if y >= x {
		return y - x
	}

	return y + 9 - x + 1
}
