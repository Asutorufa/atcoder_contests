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

	Ss := make([]string, N)

	for i := range Ss {
		fmt.Fscan(br, &Ss[i])
	}

	type point struct {
		num   int
		count int
	}
	points := make([]*point, N)
	for i := range points {
		points[i] = &point{
			num:   i,
			count: 0,
		}
	}

	for i := 0; i < M; i++ {
		cache := [2][]int{}
		for j, s := range Ss {
			if s[i] == '0' {
				cache[0] = append(cache[0], j)
			} else {
				cache[1] = append(cache[1], j)
			}
		}

		if len(cache[0]) == 0 || len(cache[1]) == 0 {
			for _, v := range cache[0] {
				points[v].count++
			}

			for _, v := range cache[1] {
				points[v].count++
			}

			continue
		}

		if len(cache[0]) < len(cache[1]) {
			for _, v := range cache[0] {
				points[v].count++
			}
		} else if len(cache[0]) > len(cache[1]) {
			for _, v := range cache[1] {
				points[v].count++
			}
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].count < points[j].count
	})

	i := points[N-1]

	resp := []int{}
	for _, v := range points {
		if v.count == i.count {
			resp = append(resp, v.num+1)
		}
	}

	sort.Ints(resp)

	for _, v := range resp {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
