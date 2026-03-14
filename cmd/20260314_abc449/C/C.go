package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	_, L, R := readInt(), readInt(), readInt()
	S := readString()

	maps := make(map[rune][]int)

	for i, s := range S {
		maps[s] = append(maps[s], i)
	}

	ans := 0
	for _, nums := range maps {
		ans += solve(nums, L, R)
	}

	fmt.Println(ans)
}

func solve(nums []int, L int, R int) int {
	count := 0
	p1, p2 := 0, 0

	for j := range nums {
		for p1 < j && nums[p1] < nums[j]-R {
			p1++
		}

		for p2 < j && nums[p2] <= nums[j]-L {
			p2++
		}

		count += p2 - p1
	}

	return count
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
