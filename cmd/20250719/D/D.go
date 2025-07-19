package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	type AB struct {
		A, B int
	}

	arr := make([]AB, M)

	for i := range arr {
		fmt.Fscan(br, &arr[i].A, &arr[i].B)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].A-arr[i].B < arr[j].A-arr[j].B
	})

	var sum int
	for _, v := range arr {
		if v.A > N {
			continue
		}

		sub := v.A - v.B
		more := N - v.A
		remove := more / sub
		remove++
		sum += remove
		N -= remove * sub
	}

	fmt.Println(sum)
}
