package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	H, W := readInt(), readInt()

	mas := make([][]byte, 0, H)
	for range H {
		mas = append(mas, []byte(readString()))
	}

	removeEdge(mas, H, W)

	fmt.Println(count(mas, H, W))
}

func count(mas [][]byte, H, W int) int {
	ret := 0

	for i, v := range mas {
		for j, c := range v {
			if c == '.' {
				ret++
				mark(mas, i, j, H, W)
			}
		}
	}

	return ret
}

func mark(mas [][]byte, i, j int, H, W int) {
	mas[i][j] = 'f'
	queue := [][2]int{{i, j}}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		for _, d := range dir {
			nx := q[0] + d[0]
			ny := q[1] + d[1]

			if nx >= 0 && nx < H && ny >= 0 && ny < W && mas[nx][ny] == '.' {
				mas[nx][ny] = 'f'
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}
}

func removeEdge(mas [][]byte, H, W int) {
	queue := [][2]int{}

	for i, v := range mas {
		if v[0] == '.' {
			mas[i][0] = 'x'
			queue = append(queue, [2]int{i, 0})
		}
		if v[W-1] == '.' {
			mas[i][W-1] = 'x'
			queue = append(queue, [2]int{i, W - 1})
		}
	}
	for j := range W {
		if mas[0][j] == '.' {
			mas[0][j] = 'x'
			queue = append(queue, [2]int{0, j})
		}
		if mas[H-1][j] == '.' {
			mas[H-1][j] = 'x'
			queue = append(queue, [2]int{H - 1, j})
		}
	}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]

		dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		for _, d := range dir {
			nx := q[0] + d[0]
			ny := q[1] + d[1]

			if nx >= 0 && nx < H && ny >= 0 && ny < W && mas[nx][ny] == '.' {
				mas[nx][ny] = 'x'
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}
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

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
