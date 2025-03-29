package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	arrs := make([]int, 0, N)
	for i := 0; i < N; i++ {
		var P int
		fmt.Fscan(br, &P)

		arrs = append(arrs, P)
	}

	ca := make([]int, N)
	copy(ca, arrs)

	sort.Slice(arrs, func(i, j int) bool {
		return arrs[j] < arrs[i]
	})

	var mm = map[int]int{}

	i := 1

	for _, v := range arrs {
		_, ok := mm[v]
		if !ok {
			mm[v] = i
		}
		i++
	}

	for _, v := range ca {
		fmt.Println(mm[v])
	}
}
