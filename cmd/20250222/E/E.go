package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		0 -> 0
		0 -> 1
		1 -> 2
		2 -> 3
		3 -> 0

	*/
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	var strings []string

	root := map[int]*Point{}

	for i := 0; i < N; i++ {
		var S string
		fmt.Fscanln(br, &S)

		strings = append(strings, S)

		for j, v := range S {
			if v != '-' {
				p := root[i]
				if p == nil {
					p = &Point{
						Next: map[int]*Point{},
						Moji: map[int]string{},
					}
					root[i] = p
				}

				jp := root[j]
				if jp == nil {
					jp = &Point{
						Next: map[int]*Point{},
						Moji: map[int]string{},
					}
					root[j] = jp
				}

				p.Next[j] = jp
				p.Moji[j] = string(v)
			}
		}
	}

	// fmt.Print(findPath(root[0].Next, map[int]bool{0: true}, 3), " ")

	for i, v := range strings {
		for j := range v {
			if i == j && root[i].Next[j] != nil {
				fmt.Print(0, " ")
			} else {
				Z := findPath(root[i], map[int]bool{i: true}, j)
				if Z != "A" {
					fmt.Print(len(Z), " ")
				} else {
					fmt.Print(-1, " ")
				}
			}
		}
		fmt.Print("\n")
	}
}

type Point struct {
	Moji map[int]string
	Next map[int]*Point
}

func findPath(maps *Point, history map[int]bool, end int) string {
	next := maps.Next[end]
	if next != nil {
		return next.Moji[end]
	}

	current := "A"
	for z, v := range maps.Next {
		if history[z] {
			continue
		}

		history[z] = true

		x := findPath(v, history, end)
		if x != "A" && (current == "A" || len(x) < len(current)) {
			current = maps.Moji[z] + x
		}
	}

	// fmt.Println(current)
	return current
}
