package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	var minx, miny int = math.MaxInt, math.MaxInt
	var maxx, maxy int = -1, -1

	points := make(map[[2]int]bool, N)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		points[[2]int{x, y}] = true

		if x < minx {
			minx = x
		}

		if y < miny {
			miny = y
		}

		if x > maxx {
			maxx = x
		}

		if y > maxy {
			maxy = y
		}
	}

	x := (maxx-minx)/2 + minx
	y := (maxy-miny)/2 + miny

	ans := 0

	// fmt.Println(minx, miny, maxx, maxy)
	// fmt.Println(x, y)

	for _, v := range [][2]int{
		{minx, miny},
		{minx, maxy},
		{maxx, miny},
		{maxx, maxy},
	} {
		if v[0]-x > ans {
			ans = v[0] - x
		}

		if v[1]-y > ans {
			ans = v[1] - y
		}
	}

	fmt.Println(ans)
}
