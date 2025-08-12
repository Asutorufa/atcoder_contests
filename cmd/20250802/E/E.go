package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(br, &T)

	for i := 0; i < T; i++ {
		var N, M, X, Y int
		fmt.Fscan(br, &N, &M, &X, &Y)

		X--
		Y--

		Graph := make([][]int, N)
		for i := range Graph {
			Graph[i] = make([]int, N)
		}

		for i := 0; i < M; i++ {
			var U, V int
			fmt.Fscan(br, &U, &V)

			U--
			V--
			Graph[U][V] = 1
			Graph[V][U] = 1
		}

		vistied := make([]bool, N)
		vistied[X] = true
		result := Answer(Graph, X, Y, vistied)

		result.Array = append(result.Array, X)

		result.Print()
	}
}

type Queue struct {
	True  bool
	Array []int
}

func (q Queue) Print() {
	for i := len(q.Array) - 1; i >= 0; i-- {
		fmt.Print(q.Array[i] + 1)
		if i != 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}

func Answer(graph [][]int, X, Y int, visited []bool) Queue {
	if X == Y {
		return Queue{
			True:  true,
			Array: make([]int, 0, len(graph)),
		}
	}

	start := graph[X]

	for i, v := range start {
		if v != 1 {
			continue
		}

		if visited[i] {
			continue
		}

		// 因为已经访问过的百分百确认不能到达终点，所以不需要恢复
		visited[i] = true

		a := Answer(graph, i, Y, visited)
		if a.True {
			a.Array = append(a.Array, i)
			return a
			// 按从小到大的顺序，第一个能找到的必定是最小的
		}
	}

	return Queue{}
}

/*
1
6 10 3 5
1 2
1 3
1 5
1 6
2 4
2 5
2 6
3 4
3 5
5 6
*/
