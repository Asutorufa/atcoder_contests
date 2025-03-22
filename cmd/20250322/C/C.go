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

	dul := make(map[int]bool, N)
	no := make(map[int]int, N)

	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		if dul[A] {
			continue
		}

		if _, ok := no[A]; ok {
			delete(no, A)
			dul[A] = true
			continue
		}

		no[A] = i + 1
	}

	max := -1
	index := -1
	for k, v := range no {
		if k > max {
			max = k
			index = v
		}
	}

	fmt.Println(index)
}
