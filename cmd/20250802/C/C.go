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

	arrays := make([]int, 0, N)

	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		arrays = append(arrays, A)
	}

	fmt.Println(countValidPairs(arrays))
}

/*
j - i = array[j] + array[i]

==> j - array[j] = i + array[i] && j > i

动态规划即可
*/
func countValidPairs(arr []int) int {
	mem := make(map[int]int)
	count := 0

	for i := len(arr) - 1; i >= 0; i-- {
		key := i + arr[i]
		count += mem[key]
		mem[i-arr[i]]++
	}

	return count
}
