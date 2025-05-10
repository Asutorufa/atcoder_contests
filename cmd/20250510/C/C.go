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

	var array = make([]int, 0, N)
	for i := 0; i < N; i++ {
		var x int

		fmt.Fscan(br, &x)
		array = append(array, x)
	}

	fmt.Println(solve2(array))
}

func solve2(array []int) int {
	count := array[len(array)-1]
	sum := 0

	for i := len(array) - 2; i >= 0; i-- {
		sum += count * array[i]
		count += array[i]
	}

	return sum
}
