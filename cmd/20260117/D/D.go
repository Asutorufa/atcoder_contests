package main

import (
	"bufio"
	"fmt"
	"os"
)

type hen struct {
	to   int
	cost int
}

var (
	N, M, L int
	S, T    int
	g       [][]hen
	ok      []bool
)

func dfs(v, step, cost int) {
	if cost > T {
		return
	}
	if step == L {
		if cost >= S && cost <= T {
			ok[v] = true
		}
		return
	}
	for _, e := range g[v] {
		dfs(e.to, step+1, cost+e.cost)
	}
}

func main() {
	fmt.Fscan(br, &N, &M, &L, &S, &T)

	g = make([][]hen, N)
	for i := 0; i < M; i++ {
		var u, v, c int
		fmt.Fscan(br, &u, &v, &c)
		u--
		v--
		g[u] = append(g[u], hen{v, c})
	}

	ok = make([]bool, N)

	dfs(0, 0, 0)

	for i := 0; i < N; i++ {
		if ok[i] {
			fmt.Print(i+1, " ")
		}
	}
	fmt.Println()
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
