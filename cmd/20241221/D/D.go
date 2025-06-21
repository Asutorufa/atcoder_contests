package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M, Sx, Sy int
	fmt.Fscanln(br, &N, &M, &Sx, &Sy)

	// Xmin := map[int]int{}
	// Xmax := map[int]int{}
	// Ymin := map[int]int{}
	// Ymax := map[int]int{}

	Yies := make(map[int][]Point, 1000)
	Xies := make(map[int][]Point, 1000)
	for i := 0; i < N; i++ {
		var x, y int
		fmt.Fscanln(br, &x, &y)

		if _, ok := Xies[x]; !ok {
			Xies[x] = make([]Point, 0, 1000)
		}

		Xies[x] = append(Xies[x], Point{X: x, Y: y})
		Yies[y] = append(Yies[y], Point{X: x, Y: y})
	}

	cache := make(map[Point]bool, 1000)
	count := 0
	for i := 0; i < M; i++ {
		var D string
		var C int
		fmt.Fscanln(br, &D, &C)
		switch D {
		case "L":
			start := Sx
			Sx -= C

			for _, v := range Yies[Sy] {
				if v.X <= start && v.X >= Sx {
					if !cache[v] {
						count++
						cache[v] = true
					}
				}
			}
		case "R":
			start := Sx
			Sx += C

			for _, v := range Yies[Sy] {
				if v.X >= start && v.X <= Sx {
					if !cache[v] {
						count++
						cache[v] = true
					}
				}
			}
		case "U":
			start := Sy
			Sy += C

			for _, v := range Xies[Sx] {
				if v.Y >= start && v.Y <= Sy {
					if !cache[v] {
						count++
						cache[v] = true
					}
				}
			}
		case "D":
			start := Sy
			Sy -= C

			for _, v := range Xies[Sx] {
				if v.Y <= start && v.Y >= Sy {
					if !cache[v] {
						count++
						cache[v] = true
					}
				}
			}
		}

		// fmt.Println(Sx, Sy, count)
	}

	fmt.Println(Sx, Sy, count)
}

type Point struct {
	X, Y int
}
