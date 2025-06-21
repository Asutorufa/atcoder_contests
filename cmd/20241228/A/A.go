package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var xs []int
	readIntFunc(4, br, func(x, _ int) { xs = append(xs, x) })

	count := map[int]int{}

	for _, v := range xs {
		count[v] = count[v] + 1
	}

	if len(count) != 2 {
		fmt.Println("No")
		return
	}

	var keys []int
	for k := range count {
		keys = append(keys, k)
	}

	if count[keys[0]] == 2 && count[keys[1]] == 2 {
		fmt.Println("Yes")
		return
	} else if (count[keys[0]] == 3 && count[keys[1]] == 1) || (count[keys[0]] == 1 && count[keys[1]] == 3) {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}
