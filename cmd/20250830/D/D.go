package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var Rt, Ct, Ra, Ca int
	fmt.Fscan(br, &Rt, &Ct, &Ra, &Ca)

	var N, M, L int
	fmt.Fscan(br, &N, &M, &L)

	var taka, aoi []Step

	for i := 0; i < M; i++ {
		var S string
		var X int
		fmt.Fscan(br, &S, &X)
		taka = append(taka, Step{S: S, X: X})
	}

	for i := 0; i < L; i++ {
		var S string
		var X int
		fmt.Fscan(br, &S, &X)
		aoi = append(aoi, Step{S: S, X: X})
	}

	taka, aoi = Split(aoi, taka)

	fmt.Println(Calulate(Point{X: Rt, Y: Ct}, Point{X: Ra, Y: Ca}, taka, aoi))
}

type Step struct {
	S string
	X int
}

func Split(a, t []Step) ([]Step, []Step) {
	nexta := func() (Step, bool) {
		if len(a) == 0 {
			return Step{}, false
		}
		x := a[0]
		a = a[1:]
		return x, true
	}

	nextt := func() (Step, bool) {
		if len(t) == 0 {
			return Step{}, false
		}
		x := t[0]
		t = t[1:]
		return x, true
	}

	var newa, newt []Step

	var ca, cb Step
	for {
		var ok bool
		if ca.X == 0 {
			ca, ok = nexta()
			if !ok {
				if cb.X > 0 {
					newt = append(newt, cb)
				}
				break
			}
		}

		if cb.X == 0 {
			cb, ok = nextt()
			if !ok {
				if ca.X > 0 {
					newa = append(newa, ca)
				}
				break
			}
		}

		if ca.X > cb.X {
			newt = append(newt, cb)
			newa = append(newa, Step{
				S: ca.S,
				X: cb.X,
			})
			ca.X -= cb.X
			cb.X = 0
		} else {
			newa = append(newa, ca)
			newt = append(newt, Step{
				S: cb.S,
				X: ca.X,
			})
			cb.X -= ca.X
			ca.X = 0
		}
	}

	for {
		ca, ok := nexta()
		if !ok {
			break
		}
		newa = append(newa, ca)
	}

	for {
		cb, ok := nextt()
		if !ok {
			break
		}
		newt = append(newt, cb)
	}

	return newt, newa
}

type Point struct {
	X, Y int
}

func (p Point) Aruku(s Step) Point {
	// fmt.Println(p, s)
	switch s.S {
	case "U":
		return Point{X: p.X - s.X, Y: p.Y}
	case "D":
		return Point{X: p.X + s.X, Y: p.Y}
	case "L":
		return Point{X: p.X, Y: p.Y - s.X}
	case "R":
		return Point{X: p.X, Y: p.Y + s.X}
	}

	return p
}

type All [2]string

var (
	UU = All{"U", "U"}
	DD = All{"D", "D"}
	LL = All{"L", "L"}
	RR = All{"R", "R"}

	UD = All{"U", "D"}
	DU = All{"D", "U"}
	LR = All{"L", "R"}
	RL = All{"R", "L"}

	// check Intersect

	UL = All{"U", "L"}
	UR = All{"U", "R"}

	DL = All{"D", "L"}
	DR = All{"D", "R"}

	LU = All{"L", "U"}
	LD = All{"L", "D"}

	RU = All{"R", "U"}
	RD = All{"R", "D"}
)

func Calulate(t, a Point, ts, as []Step) int {
	l := min(len(as), len(ts))

	startt, starta := t, a

	ret := 0
	for i := 0; i < l; i++ {
		a := as[i]
		t := ts[i]

		// fmt.Println(t, a, startt, endt, starta, enda)

		if startt.X == starta.X && startt.Y == starta.Y {
			if a.S == t.S {
				ret += t.X
			}
		} else {
			// 这里 dist 是两点的 曼哈顿距离（|x1-x2| + |y1-y2|）。因为每次只能上下左右走，走的路径长度就是曼哈顿距离。
			dist := abs(starta.X-startt.X) + abs(starta.Y-startt.Y)
			/*
				为什么「最多走到中点」？

					如果两人相向而行：
						距离是 dist = |x2 - x1|。
						每秒距离缩短 |d1 - d2|（相向就是 2）。
						他们会在 t = dist/2 的时候相遇，正好在中点。
						超过中点以后，他们就会“错开”并交错而过，不会再重合。

					如果两人同向而行（速度相同）：
						d1 = d2，公式分母为 0。
						他们要么永远平行、永远不相遇；要么一开始就在一起。
						所以检查中点也不会丢失情况。
						如果速度不同（假设 |d1| = |d2| = 1）：
						一个人快，一个人慢。
						那么相遇时刻要么发生在快的追上慢的过程里，要么永远追不上。
						这个过程也只可能在“追赶的过程中”发生，一旦超过，距离会一直保持正数，不会再缩回零。


					动画 (t = 时间)
							t=0:   A(0)                   B(10)
							t=1:     A(1)               B(9)
							t=2:       A(2)           B(8)
							t=3:         A(3)       B(7)
							t=4:           A(4)   B(6)
							t=5:             A(5)B(5)   ← 相遇点
							t=6:               A(6) B(4)
							t=7:                 A(7)B(3)
							t=8:                   A(8)B(2)
			*/
			p := a.X
			if dist/2 < p {
				p = dist / 2
			}
			nx1 := starta.Aruku(Step{S: a.S, X: p})
			nx2 := startt.Aruku(Step{S: t.S, X: p})

			if nx1 == nx2 {
				ret += 1
			}
		}

		starta = starta.Aruku(a)
		startt = startt.Aruku(t)
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
