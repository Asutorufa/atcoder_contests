package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(br, &H, &W)

	mas := make([]string, H)

	wap := map[byte][]Point{}

	for i := range mas {
		var s string
		fmt.Fscan(br, &s)
		mas[i] = s

		for j := range s {
			if s[j] != '.' && s[j] != '#' {
				wap[s[j]] = append(wap[s[j]], Point{X: i, Y: j})
			}
		}
	}

	dirs := [][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	dist := make([][]int, H)
	for i := range dist {
		dist[i] = make([]int, W)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[0][0] = 0

	// fmt.Println(mas)

	visited := make([][]bool, H)
	for i := range visited {
		visited[i] = make([]bool, W)
	}
	visited[0][0] = true

	queue := []Point{
		{X: 0, Y: 0},
	}

	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]

		// fmt.Println("-->", p)

		if p.X == H-1 && p.Y == W-1 {
			fmt.Println(dist[p.X][p.Y])
			return
		}

		for _, d := range dirs {
			nx, ny := p.X+d[0], p.Y+d[1]
			if nx < 0 || nx >= H || ny < 0 || ny >= W {
				continue
			}

			if mas[nx][ny] == '#' {
				continue
			}

			if visited[nx][ny] {
				continue
			}
			visited[nx][ny] = true
			dist[nx][ny] = dist[p.X][p.Y] + 1
			queue = append(queue, Point{X: nx, Y: ny})
		}

		z := mas[p.X][p.Y]
		t := wap[z]
		if t == nil {
			continue
		}

		for _, v := range t {
			if visited[v.X][v.Y] {
				continue
			}
			visited[v.X][v.Y] = true
			dist[v.X][v.Y] = dist[p.X][p.Y] + 1
			queue = append(queue, v)
		}
		wap[z] = nil
	}

	fmt.Println(-1)
}
