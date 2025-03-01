package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	minl := -1
	maps := map[int]int{}
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		if x, ok := maps[A]; ok {
			if minl == -1 || minl > i-x+1 {
				// fmt.Println(i, x, i-x+1)
				minl = i - x + 1
			}
		}

		maps[A] = i
	}

	fmt.Println(minl)
}

/*
10
1 2 3 5 8 1 13 21 34 55
*/
