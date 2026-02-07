package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	N := readInt()
	D := readInt()
	A := make([]int, N)
	for i := range N {
		A[i] = readInt()
	}

	tree := redblacktree.NewWithIntComparator()

	add := func(val int) {
		count, found := tree.Get(val)
		if found {
			tree.Put(val, count.(int)+1)
		} else {
			tree.Put(val, 1)
		}
	}

	remove := func(val int) {
		count, found := tree.Get(val)
		if !found {
			return
		}
		c := count.(int)
		if c > 1 {
			tree.Put(val, c-1)
		} else {
			tree.Remove(val)
		}
	}

	left := 0
	ans := 0

	for right := range N {
		val := A[right]

		minQuery := val - D + 1
		maxLimit := val + D

		for {
			node, found := tree.Ceiling(minQuery)

			if !found || node.Key.(int) >= maxLimit {
				break
			}

			remove(A[left])
			left++
		}

		add(val)

		ans += (right - left + 1)
	}

	fmt.Println(ans)
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
