package main

import (
	"bufio"
	"fmt"
	"os"
)

// 直接并查集
func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	dsu := NewDsu(N + 1)

	c := make([]int, N+1)
	s := make([]int, N+1)

	for i := 0; i < Q; i++ {
		var T int
		fmt.Fscan(br, &T)

		switch T {
		case 1:
			var u, v int
			fmt.Fscan(br, &u, &v)
			u = dsu.Leader(u)
			v = dsu.Leader(v)

			if u != v {
				dsu.Merge(u, v)
				w := dsu.Leader(u)
				s[w] = s[u] + s[v]

				// 清空旧的那个不用的leader, u和v中有一个和w相等
				// 所以 u^v^w 就是剩余那个不用的leader
				//
				// 在 dsu.Merge(u, v) 之后，新的代表元（Leader）变成了 w。
				// 但是我们一开始拿到的 u 和 v 都是 Leader（u = dsu.Leader(u)）。
				// 所以合并前，集合代表是 u 和 v，合并后，集合代表是 w。
				s[u^v^w] = 0
			}
		case 2:
			var u int
			fmt.Fscan(br, &u)

			/*
				等于下面的代码，更新根节点黑色的计数

				if c[u] == 0 {
					// 原来是白，现在变黑
					c[u] = 1
					s[dsu.Leader(u)] += 1   // 集合里的黑点数量 +1
				} else {
					// 原来是黑，现在变白
					c[u] = 0
					s[dsu.Leader(u)] -= 1   // 集合里的黑点数量 -1
				}
			*/
			s[dsu.Leader(u)] -= c[u]
			c[u] ^= 1
			s[dsu.Leader(u)] += c[u]
		case 3:
			var u int
			fmt.Fscan(br, &u)

			if s[dsu.Leader(u)] > 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

//
// Disjoint Set Union: Union Find Tree
//

// DSU :
type DSU struct {
	parentOrSize []int
	n            int
}

// NewDsu :
func NewDsu(n int) *DSU {
	var d DSU
	d.n = n
	d.parentOrSize = make([]int, n)
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return &d
}

// Merge :
func (d DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

// Same :
func (d DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

// Leader :
func (d DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

// Size :
func (d DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

// Groups :
func (d DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}
