package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(br, &T)

	for i := 0; i < T; i++ {
		var c, d int
		fmt.Fscan(br, &c, &d)

		fmt.Println(solve(c, d))
	}
}

func solve(c, d int) int {
	xmin, xmax, cshift := 1, 9, 10

	ret := 0
	for xmin <= c+d {
		l := max(xmin, c+1)
		r := min(xmax, c+d)

		if l <= r {
			start := c*cshift + l
			end := c*cshift + r
			ret += countv2(start, end)
		}

		xmin *= 10
		xmax = (xmax+1)*10 - 1
		cshift *= 10
	}

	return ret
}

func isqrt(x int64) int64 {
	y := int64(math.Sqrt(float64(x)))
	for y*y > x {
		y--
	}
	for (y+1)*(y+1) <= x {
		y++
	}
	return y
}

func countv2(min, max int) int {
	start := isqrt(int64(min - 1))
	end := isqrt(int64(max))

	return int(end - start)
}

// !!!!!! 会越界，坑爹呀。。
func count(min, max int) int {
	start := math.Ceil(math.Sqrt(float64(min)))
	end := math.Floor(math.Sqrt(float64(max)))

	return int(end - start + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
