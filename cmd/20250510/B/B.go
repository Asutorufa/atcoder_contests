package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var array = make([]int, 0, N)
	for i := 0; i < N; i++ {
		var x int

		fmt.Fscan(br, &x)
		array = append(array, x)
	}

	count := 0

	for {
		if len(array) == count {
			break
		}

		if !containAll(array[:len(array)-count], M) {
			break
		}

		count++
	}

	fmt.Println(count)
}

func containAll(a []int, M int) bool {
	dd := make(map[int]struct{}, M)
	for i := 1; i <= M; i++ {
		dd[i] = struct{}{}
	}

	for _, v := range a {
		delete(dd, v)
	}

	return len(dd) == 0
}

/*
1 1
1
*/
