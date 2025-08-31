package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
因为只有A和B所以最后的位置是固定的

要么 ABABABAB
要么 BABABABA

只需要把A或者B移到正确的位置，最后算两种情况移动次数的最小值
*/
func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	var S string
	fmt.Fscan(br, &S)

	A, B := 0, 0

	count := 0
	for i, v := range S {
		if v == 'B' {
			count++

			A += abs(((count * 2) - 1) - i)
			B += abs(((count - 1) * 2) - i)
		}

	}

	if A > B {
		fmt.Println(B)
	} else {
		fmt.Println(A)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
