package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscanln(br, &H, &W)

	var start *Point
	mas := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscanln(br, &mas[i])

		if start == nil {
			z := strings.IndexByte(mas[i], 'S')
			if z != -1 {
				start = &Point{X: i, Y: z}
			}
		}
	}

	// 其次XY缓存得分开，不然过不了
	// 按理说以题目的说明应该是可以走回头路的，不知道是我想多了，还是题目不严谨
	cacheX := make(map[Point]bool, 1000)
	cacheY := make(map[Point]bool, 1000)

	queue := []Step{
		{
			Point: *start,
			IsY:   true,
		},
		{
			Point: *start,
			IsY:   false,
		},
	}

	// 广度优先，用队列
	// 深度优先是递归
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if mas[p.X][p.Y] == 'G' {
			fmt.Println(p.Step)
			return
		}

		if p.IsY {
			if p.Y > 0 && mas[p.X][p.Y-1] != '#' && !cacheY[Point{X: p.X, Y: p.Y - 1}] {
				cacheY[Point{X: p.X, Y: p.Y - 1}] = true
				queue = append(queue, Step{
					Point: Point{X: p.X, Y: p.Y - 1},
					IsY:   false,
					Step:  p.Step + 1,
				})
			}

			if p.Y < W-1 && mas[p.X][p.Y+1] != '#' && !cacheY[Point{X: p.X, Y: p.Y + 1}] {
				cacheY[Point{X: p.X, Y: p.Y + 1}] = true
				queue = append(queue, Step{
					Point: Point{X: p.X, Y: p.Y + 1},
					IsY:   false,
					Step:  p.Step + 1,
				})
			}
		} else {
			if p.X > 0 && mas[p.X-1][p.Y] != '#' && !cacheX[Point{X: p.X - 1, Y: p.Y}] {
				cacheX[Point{X: p.X - 1, Y: p.Y}] = true
				queue = append(queue, Step{
					Point: Point{X: p.X - 1, Y: p.Y},
					IsY:   true,
					Step:  p.Step + 1,
				})
			}

			if p.X < H-1 && mas[p.X+1][p.Y] != '#' && !cacheX[Point{X: p.X + 1, Y: p.Y}] {
				cacheX[Point{X: p.X + 1, Y: p.Y}] = true
				queue = append(queue, Step{
					Point: Point{X: p.X + 1, Y: p.Y},
					IsY:   true,
					Step:  p.Step + 1,
				})
			}
		}
	}

	fmt.Println(-1)

	// 错误解法，不能用DFS
	// step, ok := (&Process{
	// 	Start: *start,
	// 	Mas:   mas,
	// 	H:     H,
	// 	W:     W,
	// }).Do()

	// if ok {
	// 	fmt.Println(step)
	// } else {
	// 	fmt.Println(-1)
	// }
}

type Step struct {
	Point
	IsY  bool
	Step int
}

type Point struct {
	X, Y int
}

type Process struct {
	Start Point
	H, W  int
	Mas   []string
}

var (
	cacheX = map[Point]bool{}
	cacheY = map[Point]bool{}
)

func (p *Process) Do() (int, bool) {
	s := p.Start

	var step int
	var ok bool
	cacheX = map[Point]bool{}
	cacheY = map[Point]bool{}
	l, lok := p.Left(s, 0, cacheY)
	cacheX = map[Point]bool{}
	cacheY = map[Point]bool{}
	r, rok := p.Right(s, 0, cacheY)
	cacheX = map[Point]bool{}
	cacheY = map[Point]bool{}
	u, uok := p.Up(s, 0, cacheX)
	cacheX = map[Point]bool{}
	cacheY = map[Point]bool{}
	d, dok := p.Down(s, 0, cacheX)

	// fmt.Println(u, d, l, r)
	if uok {
		step = u
		ok = true
	}

	if dok {
		if step == 0 || d < step {
			step = d
		}
		ok = true
	}

	if lok {
		if step == 0 || l < step {
			step = l
		}
		ok = true
	}

	if rok {
		if step == 0 || r < step {
			step = r
		}
		ok = true
	}

	return step, ok
}

func (p *Process) lOrR(c Point, step int) (int, bool) {
	l, lok := p.Left(c, step+1, cacheY)
	r, rok := p.Right(c, step+1, cacheY)

	if lok && rok {
		if l < r {
			return l, true
		} else {
			return r, true
		}
	} else if lok {
		return l, true
	} else if rok {
		return r, true
	} else {
		return 0, false
	}
}

func (p *Process) uOrD(c Point, step int) (int, bool) {
	u, uok := p.Up(c, step+1, cacheX)
	d, dok := p.Down(c, step+1, cacheX)

	if uok && dok {
		if u < d {
			return u, true
		} else {
			return d, true
		}
	} else if uok {
		return u, true
	} else if dok {
		return d, true
	} else {
		return 0, false
	}
}

func (p *Process) Up(c Point, step int, cache map[Point]bool) (int, bool) {
	if p.Mas[c.X][c.Y] == 'G' {
		return step, true
	}

	if c.X == 0 {
		return 0, false
	}

	c.X--

	if p.Mas[c.X][c.Y] == '#' {
		return 0, false
	}

	if cache[c] {
		return 0, false
	}

	cache[c] = true

	step, ok := p.lOrR(c, step)
	if !ok {
		cache[c] = false
	}

	return step, ok
}

func (p *Process) Down(c Point, step int, cache map[Point]bool) (int, bool) {
	if p.Mas[c.X][c.Y] == 'G' {
		return step, true
	}

	if c.X == p.H-1 {
		return 0, false
	}

	c.X++

	if p.Mas[c.X][c.Y] == '#' {
		return 0, false
	}

	if cache[c] {
		return 0, false
	}

	cache[c] = true

	step, ok := p.lOrR(c, step)
	if !ok {
		cache[c] = false
	}

	return step, ok
}

func (p *Process) Left(c Point, step int, cache map[Point]bool) (int, bool) {
	if p.Mas[c.X][c.Y] == 'G' {
		return step, true
	}

	if c.Y == 0 {
		return 0, false
	}

	c.Y--

	if p.Mas[c.X][c.Y] == '#' {
		return 0, false
	}

	if cache[c] {
		return 0, false
	}

	cache[c] = true

	step, ok := p.uOrD(c, step)
	if !ok {
		cache[c] = false
	}

	return step, ok
}

func (p *Process) Right(c Point, step int, cache map[Point]bool) (int, bool) {
	if p.Mas[c.X][c.Y] == 'G' {
		return step, true
	}

	if c.Y == p.W-1 {
		return 0, false
	}

	c.Y++

	if p.Mas[c.X][c.Y] == '#' {
		return 0, false
	}

	if cache[c] {
		return 0, false
	}

	cache[c] = true

	step, ok := p.uOrD(c, step)
	if !ok {
		cache[c] = false
	}

	return step, ok
}
