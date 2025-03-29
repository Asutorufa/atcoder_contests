package main

import (
	"bufio"
	"fmt"
	"os"
)

// 素集合データ構造(Union-Find) そしゅうごうデータこうぞう
// 并查集/dfs 即可 完全忘了

/*
#include <iostream>
using namespace std;
#include "atcoder/dsu.hpp"

	int main() {
	  int N, M;
	  cin >> N >> M;
	  atcoder::dsu uf(N);
	  int K = N;
	  for (int i = 0; i < M; i++) {
	    int u, v;
	    cin >> u >> v;
	    --u, --v;
	    if (!uf.same(u, v)) uf.merge(u, v), K--;
	  }
	  cout << M - (N - K) << "\n";
	}


	これが残すことができる辺の本数の最大なので、元のグラフを森にするために削除すべき辺の本数の最小値

	先计算一棵树最多能有几条边，利用并查集合并，最后用已有的减去总的，得到的就是能删减的边的数量
*/
func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	uf := NewUnionFind(N)

	K := 0
	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		if uf.same(x-1, y-1) {
			K++
		} else {
			uf.union(x-1, y-1)
		}
	}

	fmt.Println(K)
}

func Dfs() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var points = make([]*Point, N+1)

	for i := range points {
		points[i] = NewPoint(i)
	}

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		points[x].Add(y, points[y])
		points[y].Add(x, points[x])
	}

	cache := map[int]bool{}

	var dfs func(last []int, p *Point)

	dfs = func(last []int, p *Point) {
		cache[p.X] = true

		for x := range p.Next {
			if !cache[x] {
				dfs(append(last, p.X), p.Get(x))
			}
		}
	}

	conntectd := 0
	for i := 1; i <= N; i++ {
		if !cache[i] {
			dfs([]int{i}, points[i])
			conntectd++
		}
	}

	if N == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(M - N + conntectd)
	}
}

type Point struct {
	X    int
	Next map[int]*Point
}

func NewPoint(x int) *Point {
	return &Point{
		X:    x,
		Next: map[int]*Point{},
	}
}

func (p *Point) Add(x int, next *Point) {
	p.Next[x] = next
}

func (p *Point) Get(x int) *Point {
	return p.Next[x]
}

func (p *Point) Remove(x int) {
	delete(p.Next, x)
}

type UnionFind struct {
	node []int
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		node: make([]int, n),
	}

	for i := range u.node {
		u.node[i] = -1
	}

	return u
}

func (u *UnionFind) root(x int) int {
	if u.node[x] < 0 {
		return x
	}

	u.node[x] = u.root(u.node[x])
	return u.node[x]
}

func (u *UnionFind) union(x, y int) {
	xr := u.root(x)
	yr := u.root(y)

	if xr == yr {
		return
	}

	u.node[xr] += u.node[yr]
	u.node[yr] = xr
}

func (u *UnionFind) size(x int) int {
	return -u.node[u.root(x)]
}

func (u *UnionFind) same(x, y int) bool {
	return u.root(x) == u.root(y)
}
