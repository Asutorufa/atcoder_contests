package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	As := make([]int, N)
	Bs := make([]int, N)

	for i := range As {
		fmt.Fscan(br, &As[i])
	}

	for i := range Bs {
		fmt.Fscan(br, &Bs[i])
	}

	current := make([]int, N)

	Sum := 0
	for i, A := range As {
		B := Bs[i]

		c := min(A, B)
		Sum += c
		current[i] = c
	}

	for i := 0; i < Q; i++ {
		var C string
		var X, V int
		fmt.Fscan(br, &C, &X, &V)
		X--

		switch C {
		case "A":
			As[X] = V
		case "B":
			Bs[X] = V
		}

		Sum -= current[X]
		current[X] = min(As[X], Bs[X])
		Sum += current[X]

		fmt.Println(Sum)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
