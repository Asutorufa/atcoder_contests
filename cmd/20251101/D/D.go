package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	var n int
	fmt.Fscan(br, &n)

	st := redblacktree.NewWithIntComparator()

	var x int
	fmt.Fscan(br, &x)
	st.Put(0, x)
	st.Put(x, x)
	res := x * 2

	fmt.Fprintln(bw, res)

	for i := 1; i < n; i++ {
		fmt.Fscan(br, &x)

		var dist = math.MaxInt

		if idx, ok := st.Floor(x); ok {
			key := idx.Key.(int)

			fd := x - key
			dist = min(dist, fd)

			if value := idx.Value.(int); fd < value {
				res -= value - fd
				st.Put(key, fd)
			}
		}

		if idx, ok := st.Ceiling(x); ok {
			key := idx.Key.(int)

			td := key - x
			dist = min(dist, td)

			if value := idx.Value.(int); td < value {
				res -= value - td
				st.Put(key, td)
			}
		}

		res += dist

		st.Put(x, dist)

		fmt.Fprintln(bw, res)
	}
}
