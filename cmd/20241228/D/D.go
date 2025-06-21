package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscanln(br, &N, &M)

	X := make(map[int][]Point, N)
	Y := make(map[int][]Point, N)

	for i := 0; i < M; i++ {
		var x, y int
		var C string
		fmt.Fscanln(br, &x, &y, &C)

		X[x] = append(X[x], Point{X: x, Y: y, C: C})
		Y[y] = append(Y[y], Point{X: x, Y: y, C: C})
	}

	for _, v := range X {
		sort.Sort(&Points{Ps: v, LessF: func(i, j Point) bool {
			return Compare(i.Y, j.Y)
		}})
	}

	for _, v := range Y {
		sort.Sort(&Points{Ps: v, LessF: func(i, j Point) bool {
			return Compare(i.X, j.X)
		}})
	}

	// fmt.Println(X, Y)

	for _, v := range X {
		hasW := false
		for _, vv := range v {
			switch vv.C {
			case "B":
				if hasW {
					fmt.Println("No")
					return
				}
			case "W":
				hasW = true

				for k, v := range Y {
					if k > vv.Y {
						for _, vvY := range v {
							if vvY.X > vv.X && vvY.C == "B" {
								fmt.Println("No")
								return
							}
						}
					}
				}
			}
		}
	}

	for _, v := range Y {
		hasW := false
		for _, vv := range v {
			switch vv.C {
			case "B":
				if hasW {
					fmt.Println("No")
					return
				}
			case "W":
				hasW = true
			}
		}
	}

	fmt.Println("Yes")
}

type Point struct {
	X, Y int
	C    string
}

type Points struct {
	Ps    []Point
	LessF func(i, j Point) bool
}

func (p *Points) Len() int {
	return len(p.Ps)
}

func (p *Points) Less(i, j int) bool {
	return p.LessF(p.Ps[i], p.Ps[j])
}

func (p Points) Swap(i, j int) {
	p.Ps[i], p.Ps[j] = p.Ps[j], p.Ps[i]
}

func Compare(x, y int) bool {
	return x < y
}
