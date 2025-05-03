package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var br = bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscan(br, &N, &M)

	points := make([]*point, N+1)
	for i := range points {
		points[i] = &point{next: map[int]*point{}}
	}

	for i := 0; i < M; i++ {
		var x, y int
		fmt.Fscan(br, &x, &y)

		x--
		y--
		points[x].next[y] = points[y]
		points[y].next[x] = points[x]
	}

	history := map[int]bool{}

	next := points[0]
	current := 0

	for {
		if current != 0 && len(next.next) != 1 {
			fmt.Println("No")
			return
		}

		if current == 0 && len(next.next) != 2 {
			fmt.Println("No")
			return
		}

		for x, v := range next.next {
			if x == 0 && len(history) != N-1 {
				fmt.Println("No")
				return
			}

			if current == 0 && x == 0 {
				fmt.Println("No")
				return
			}

			if x == 0 && len(history) == N-1 {
				fmt.Println("Yes")
				return
			}

			if history[x] {
				fmt.Println("No")
				return
			}

			// graph = append(graph, x)

			delete(v.next, current)
			current = x
			history[x] = true
			next = v

			break
		}
	}
}

type point struct {
	next map[int]*point
}

/*
4 3
3 1
2 1
2 4

*/
