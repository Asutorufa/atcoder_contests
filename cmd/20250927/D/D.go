package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(br, &H, &W)

	mas := make([][]byte, H)
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(br, &s)
		mas[i] = []byte(s)
	}

	queue := make([]Point, 0, H*W)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if mas[i][j] == '#' {
				queue = append(queue, Point{i, j})
			}
		}
	}

	ret := len(queue)

	for len(queue) > 0 {
		newQueue := make([]Point, 0, len(queue))

		for _, v := range queue {
			ws := Whites(mas, H, W, v.x, v.y)
			ret += len(ws)
			newQueue = append(newQueue, ws...)
		}

		for _, v := range newQueue {
			mas[v.x][v.y] = '#'
		}

		queue = newQueue
	}

	fmt.Println(ret)
}

func Whites(mas [][]byte, H, W, x, y int) []Point {
	res := make([]Point, 0, 4)
	for _, move := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		rx, ry := x+move[0], y+move[1]
		if rx < 0 || rx >= H || ry < 0 || ry >= W {
			continue
		}

		if mas[rx][ry] != '.' {
			continue
		}

		if IsOnlyOne(mas, H, W, rx, ry) {
			res = append(res, Point{rx, ry})
		}
	}

	return res
}

func IsOnlyOne(mas [][]byte, H, W, x, y int) bool {
	count := 0
	for _, move := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		rx, ry := x+move[0], y+move[1]
		if rx < 0 || rx >= H || ry < 0 || ry >= W {
			continue
		}
		if mas[rx][ry] == '#' {
			count++
		}
	}

	return count == 1
}
