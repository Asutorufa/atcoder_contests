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

	maps := make(map[int]int)
	for i := 0; i < N; i++ {
		var a int
		fmt.Fscan(br, &a)
		maps[a]++
	}

	sum := 0
	for _, v := range maps {
		if v < 2 {
			continue
		}

		s := ((v * (v - 1)) / 2) * (N - v)

		sum += s
	}

	fmt.Println(sum)
}
