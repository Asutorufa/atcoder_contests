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

	args := make([]int, N)
	start, end := make(map[int]bool, N), make(map[int]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(br, &args[i])
		end[args[i]] = end[args[i]] + 1
	}

	max := 0
	for _, v := range args {
		start[v] = true

		x := end[v] - 1
		if x <= 0 {
			delete(end, v)
		} else {
			end[v] = x
		}

		sum := len(start) + len(end)

		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
}
