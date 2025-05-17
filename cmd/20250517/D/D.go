package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W, N int
	fmt.Fscan(br, &H, &W, &N)

	cx := map[int]map[int]struct{}{}
	cy := map[int]map[int]struct{}{}

	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		if cx[x] == nil {
			cx[x] = map[int]struct{}{}
		}
		if cy[y] == nil {
			cy[y] = map[int]struct{}{}
		}

		cx[x][y] = struct{}{}
		cy[y][x] = struct{}{}
	}

	var Q int
	fmt.Fscan(br, &Q)

	for i := 0; i < Q; i++ {
		var o, ii int
		fmt.Fscan(br, &o, &ii)

		switch o {
		case 1:
			fmt.Println(len(cx[ii]))

			for j := range cx[ii] {
				delete(cy[j], ii)
			}

			cx[ii] = map[int]struct{}{}

		case 2:
			fmt.Println(len(cy[ii]))

			for j := range cy[ii] {
				delete(cx[j], ii)
			}

			cy[ii] = map[int]struct{}{}
		}
	}
}
