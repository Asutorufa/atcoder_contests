package main

import (
	"bufio"
	"fmt"
	"os"
)

type hen struct {
	X, Y int
}

var labels = map[hen]int{}
var points = map[int]*Point{}
var visited = map[int]bool{}

/*
dfs でかんたんに解ける

**単純連結無向グラフが与えられます**。

だけど　逆方向もだいじょうぶです、理解を間違えた🥲
*/

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	for i := 0; i < M; i++ {
		var x, y, l int
		fmt.Fscan(br, &x, &y, &l)

		labels[hen{x, y}] = l
		labels[hen{y, x}] = l

		p, ok := points[x]
		if !ok {
			p = &Point{Current: x, Next: map[int]*Point{}}
			points[x] = p
		}

		py, ok := points[y]
		if !ok {
			py = &Point{Current: y, Next: map[int]*Point{}}
			points[y] = py
		}

		p.Next[y] = py
		py.Next[x] = p
	}

	p := points[1]

	fmt.Println(p.Search(-1, N))
}

type Point struct {
	Current int
	Next    map[int]*Point
}

func (p *Point) Search(xor int, target int) int {
	visited[p.Current] = true
	defer func() {
		visited[p.Current] = false
	}()

	if p.Current == target {
		return xor
	}

	c := -1
	for i, v := range p.Next {
		if visited[i] {
			continue
		}

		label := labels[hen{p.Current, i}]

		var z int = label
		if xor != -1 {
			z = xor ^ label
		}

		// fmt.Println(p.Current, "->", i, xor, label, z)

		x := v.Search(z, target)
		if x != -1 {
			if c == -1 || c > x {
				c = x
			}
		}
	}

	return c
}
