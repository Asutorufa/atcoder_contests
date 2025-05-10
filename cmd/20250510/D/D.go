package main

import (
	"bufio"
	"fmt"
	"os"
)

var mas = make([][]byte, 0)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(br, &H, &W)

	mas = make([][]byte, 0, H)

	for i := 0; i < H; i++ {
		var s string
		fmt.Fscan(br, &s)
		mas = append(mas, []byte(s))
	}

	// bfs就可以，傻了，虽然写出来了，但是走了多次
	reverse(H, W)
	// for i := 0; i < H; i++ {
	// 	for j := 0; j < W; j++ {
	// 		if mas[i][j] == '.' {
	// 			findAndChange(Point{X: i, Y: j}, H, W)
	// 		}
	// 	}
	// }

	for i := range mas {
		fmt.Println(string(mas[i]))
	}
}

func reverse(H, W int) {
	queue := []Point{}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if mas[i][j] == 'E' {
				queue = append(queue, Point{X: i, Y: j})
			}
		}
	}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if np, ok := p.Down(H, W); ok && mas[np.X][np.Y] == '.' {
			queue = append(queue, np)

			mas[np.X][np.Y] = '^'
		}

		if np, ok := p.Right(H, W); ok && mas[np.X][np.Y] == '.' {
			queue = append(queue, np)

			mas[np.X][np.Y] = '<'
		}

		if np, ok := p.Up(H, W); ok && mas[np.X][np.Y] == '.' {
			queue = append(queue, np)

			mas[np.X][np.Y] = 'v'
		}

		if np, ok := p.Left(H, W); ok && mas[np.X][np.Y] == '.' {
			queue = append(queue, np)

			mas[np.X][np.Y] = '>'
		}

	}
}

var (
	queue []Point
)

func findAndChange(start Point, H, W int) {

	queue = queue[:0]
	queue = append(queue, start)

	prev := map[Point]Point{}
	visited := map[Point]bool{}

	var target Point
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if visited[p] {
			continue
		}

		visited[p] = true

		if mas[p.X][p.Y] == 'E' /*|| mas[p.X][p.Y] == '>' || mas[p.X][p.Y] == 'v' || mas[p.X][p.Y] == '^' || mas[p.X][p.Y] == '<'*/ {
			target = p
			break
		}

		if np, ok := p.Up(H, W); ok {
			if !visited[np] {
				prev[np] = p
				queue = append(queue, np)
			}
		}

		if np, ok := p.Down(H, W); ok {
			if !visited[np] {
				prev[np] = p
				queue = append(queue, np)
			}
		}

		if np, ok := p.Left(H, W); ok {
			if !visited[np] {
				prev[np] = p
				queue = append(queue, np)
			}
		}

		if np, ok := p.Right(H, W); ok {
			if !visited[np] {
				prev[np] = p
				queue = append(queue, np)
			}
		}
	}

	// fmt.Println(prev, target)
	for {
		if p, ok := prev[target]; !ok {
			break
		} else {
			change(target, p)
			target = p
			// fmt.Println(p)
		}
	}
}

func change(prev Point, current Point) {
	if mas[current.X][current.Y] == 'E' {
		return
	}

	if prev.X == current.X+1 {
		mas[current.X][current.Y] = 'v'
	} else if prev.X == current.X-1 {
		mas[current.X][current.Y] = '^'
	} else if prev.Y == current.Y+1 {
		mas[current.X][current.Y] = '>'
	} else if prev.Y == current.Y-1 {
		mas[current.X][current.Y] = '<'
	}
}

type Point struct {
	X, Y int
}

func (p Point) Up(H, W int) (Point, bool) {
	if p.X == 0 {
		return Point{}, false
	}
	if mas[p.X-1][p.Y] == '#' {
		return Point{}, false
	}
	return Point{X: p.X - 1, Y: p.Y}, true
}

func (p Point) Down(H, W int) (Point, bool) {
	if p.X == H-1 {
		return Point{}, false
	}
	if mas[p.X+1][p.Y] == '#' {
		return Point{}, false
	}
	return Point{X: p.X + 1, Y: p.Y}, true
}

func (p Point) Left(H, W int) (Point, bool) {
	if p.Y == 0 {
		return Point{}, false
	}
	if mas[p.X][p.Y-1] == '#' {
		return Point{}, false
	}
	return Point{X: p.X, Y: p.Y - 1}, true
}

func (p Point) Right(H, W int) (Point, bool) {
	if p.Y == W-1 {
		return Point{}, false
	}
	if mas[p.X][p.Y+1] == '#' {
		return Point{}, false
	}
	return Point{X: p.X, Y: p.Y + 1}, true
}
