package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
煙移動じゃなくで、高橋くんと焚き火を反対に移動する
*/
func main() {
	br := bufio.NewReader(os.Stdin)
	var N, R, C int
	fmt.Fscan(br, &N, &R, &C)

	MY := &Point{X: R, Y: C}
	BI := &Point{X: 0, Y: 0}

	Points := make(map[Point]struct{}, N)

	Points[*BI] = struct{}{}

	var S string
	fmt.Fscan(br, &S)

	for _, v := range S {
		switch v {
		case 'N':
			BI.N()
			MY.N()
		case 'S':
			BI.S()
			MY.S()
		case 'W':
			BI.W()
			MY.W()
		case 'E':
			BI.E()
			MY.E()
		}

		Points[*BI] = struct{}{}

		// fmt.Println(Points, MY)

		if _, ok := Points[*MY]; ok {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
}

func NewPoint(s string) *Point {
	p := &Point{}
	switch s {
	case "S":
		p.S()
	case "W":
		p.W()
	case "N":
		p.N()
	case "E":
		p.E()
	}
	return p
}

type Point struct {
	X, Y int
}

func (p *Point) S() {
	p.X -= 1
}

func (p *Point) W() {
	p.Y += 1
}

func (p *Point) N() {
	p.X += 1
}

func (p *Point) E() {
	p.Y -= 1
}
