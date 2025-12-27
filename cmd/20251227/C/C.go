package main

import (
	"bufio"
	"fmt"
	"os"
)

func minLength(A []int) int {
	type Node struct {
		V     int
		Count int
	}

	As := make([]Node, 0)

	for _, x := range A {
		n := len(As)
		if n > 0 && As[n-1].V == x {
			As[n-1].Count++
			if As[n-1].Count == 4 {
				As = As[:n-1]
			}
		} else {
			As = append(As, Node{V: x, Count: 1})
		}
	}

	ans := 0
	for _, node := range As {
		ans += node.Count
	}
	return ans
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	As := make([]int, 0, N)
	for range N {
		var a int
		fmt.Fscan(br, &a)
		As = append(As, a)
	}

	fmt.Println(minLength(As))
}
