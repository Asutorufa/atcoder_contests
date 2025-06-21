package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	for i := 0; i < M; i++ {
		var x, y, l int
		fmt.Fscan(br, &x, &y, &l)

		points[x] = append(points[x], node{p: y, w: l})
	}

	/*
		顶点加倍/顶点记录历史值

		每次push进去的顶点都记录更新过的值，这样既可以重复走，又不怕死循环，确实是比较好的解决方法
	*/
	queue := points[1]
	visited[node{p: 1, w: 0}] = true

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if visited[n] {
			continue
		}

		visited[n] = true

		if n.p == N && (ans < 0 || ans > n.w) {
			ans = n.w
		}

		p, ok := points[n.p]
		if !ok {
			continue
		}

		for _, v := range p {
			queue = append(queue, node{p: v.p, w: n.w ^ v.w})
		}
	}

	fmt.Println(ans)
}

type node struct {
	p, w int
}

var points = map[int][]node{}
var visited = map[node]bool{}

var ans int = -1

// 不能直接记录边，会漏掉并且会死循环可能
// func (p *Point) Search(xor int, target int) {
// 	if p.Current == target {
// 		if ans == -1 || (xor != -1 && xor < ans) {
// 			ans = xor
// 		}
// 	}

// 	for i, v := range p.Next {
// 		h := hen{p.Current, i}

// 		label := labels[h]

// 		n := node{v.Current, label}

// 		if visited[n] {
// 			continue
// 		}

// 		visited[n] = true

// 		var z int = label
// 		if xor != -1 {
// 			z = xor ^ label
// 		}

// 		v.Search(z, target)
// 	}

// }
