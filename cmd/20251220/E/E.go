package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Entry struct {
	Arr []int
	Idx int
}

func main2() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	As := make([]*Entry, N+1)
	for i := range N + 1 {
		As[i] = &Entry{}
	}

	for i := range N {
		var x, y int
		fmt.Fscan(br, &x, &y)

		arrays := make([]int, len(As[x].Arr)+1)
		copy(arrays, As[x].Arr)
		As[i+1].Arr = append(arrays, y)
		As[i+1].Idx = i + 1
	}

	As = As[1:]

	sort.Slice(As, func(i, j int) bool {
		a, b := As[i], As[j]
		n := min(len(b.Arr), len(a.Arr))

		for k := range n {
			if a.Arr[k] != b.Arr[k] {
				return a.Arr[k] < b.Arr[k]
			}
		}

		if len(a.Arr) != len(b.Arr) {
			return len(a.Arr) < len(b.Arr)
		}

		return a.Idx < b.Idx
	})

	for _, v := range As {
		fmt.Print(v.Idx, " ")
	}
	fmt.Println()
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	// G: adjacency map, key = y, value = next node id
	G := make([]map[int]int, n+1)
	for i := range G {
		G[i] = make(map[int]int)
	}

	// vs: which sequence indices belong to this node
	vs := make([][]int, n+1)

	// pos[i]: which node sequence Ai corresponds to
	pos := make([]int, n+1)
	pos[0] = 0

	tmp := 1

	for i := 1; i <= n; i++ {
		var p, y int
		fmt.Fscan(in, &p, &y)

		parent := pos[p]
		if _, ok := G[parent][y]; !ok {
			G[parent][y] = tmp
			tmp++
		}
		pos[i] = G[parent][y]
		vs[pos[i]] = append(vs[pos[i]], i)
	}

	ans := make([]int, 0, n)

	var dfs func(int)
	dfs = func(i int) {
		// Same as: ans += vs[i]
		ans = append(ans, vs[i]...)

		// Sort children by key y
		keys := make([]int, 0, len(G[i]))
		for k := range G[i] {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		for _, k := range keys {
			dfs(G[i][k])
		}
	}

	dfs(0)

	out := bufio.NewWriter(os.Stdout)
	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
	out.Flush()
}
