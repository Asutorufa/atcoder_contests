package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var TN int
	fmt.Fscan(br, &TN)

	ret := make([]int, 0, TN)
	for range TN {
		var N int
		fmt.Fscan(br, &N)
		ts := make([]T, N)
		w := 0
		for i := range ts {
			fmt.Fscan(br, &ts[i].W, &ts[i].P)
			w += ts[i].W
		}
		ret = append(ret, solve(w, ts))
	}

	for _, v := range ret {
		fmt.Println(v)
	}
}

type T struct {
	W, P int
}

func solve(W int, ts []T) int {
	sort.Slice(ts, func(i, j int) bool {
		ii := ts[i].P + ts[i].W
		jj := ts[j].P + ts[j].W

		return ii > jj
	})

	ps := 0
	count := len(ts)

	for _, v := range ts {
		ps += v.P
		W -= v.W
		count--
		if ps >= W {
			return count
		}
	}

	return 0
}
