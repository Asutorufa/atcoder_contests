package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	L, R, D, U := readInt(), readInt(), readInt(), readInt()
	xRanges := splitRange(L, R)
	yRanges := splitRange(D, U)

	ret := 0
	for _, xr := range xRanges {
		for _, yr := range yRanges {
			ret += calcRegion(xr[0], xr[1], yr[0], yr[1])
		}
	}

	fmt.Println(ret)
}

func splitRange(start, end int) [][2]int {
	var res [][2]int
	if start <= 0 && end >= 0 {
		res = append(res, [2]int{0, end})
		if start < 0 {
			res = append(res, [2]int{1, -start})
		}
	} else if end < 0 {
		res = append(res, [2]int{-end, -start})
	} else {
		res = append(res, [2]int{start, end})
	}
	return res
}

func calcRegion(x1, x2, y1, y2 int) int {
	if x1 > x2 || y1 > y2 {
		return 0
	}
	return calc(x2, y2) - calc(x1-1, y2) - calc(x2, y1-1) + calc(x1-1, y1-1)
}

func calc(x, y int) int {
	if x < 0 || y < 0 {
		return 0
	}

	minv, maxv := x, y
	if y < x {
		minv, maxv = y, x
	}

	// 有几个偶数圈
	ss := minv / 2
	// 等差数列求和找出正方形中所有满足的点
	squarePoints := (ss + 1) * (ss*2 + 1)

	evenCount := maxv/2 - ss
	// 如果是长方形，还有加上没被正方形包含的点
	extraPoints := evenCount * (minv + 1)

	return squarePoints + extraPoints
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
