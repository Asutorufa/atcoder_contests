package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(br, &n)

	array := []int{}

	for range n {
		var a int
		fmt.Fscan(br, &a)
		array = append(array, a)
	}

	left := map[int]int{}
	right := map[int]int{}

	/*
		刚开始时所有数都在右侧，然后从左到右遍历

		遍历到一个数时，把它从右侧移除
		计算左右两侧的数分别满足(7 * v) / 5，(3 * v) / 5的总数
		把当前数加入左侧
	*/
	for _, v := range array {
		right[v]++
	}

	ret := 0

	for _, v := range array {
		right[v]--

		if (7*v)%5 == 0 && (3*v)%5 == 0 {
			a := (7 * v) / 5
			c := (3 * v) / 5

			ret += right[a] * right[c]
			ret += left[a] * left[c]
		}

		left[v]++
	}

	fmt.Println(ret)
}
