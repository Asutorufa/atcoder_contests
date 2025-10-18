package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N, K int
	var S string
	fmt.Scan(&N, &K, &S)

	max := 0
	maps := map[string]int{}
	for i := 0; i < N-K+1; i++ {
		maps[S[i:i+K]]++
		if maps[S[i:i+K]] > max {
			max = maps[S[i:i+K]]
		}
	}

	rets := []string{}

	for k, v := range maps {
		if v == max {
			rets = append(rets, k)
		}
	}

	sort.Strings(rets)
	fmt.Println(max)
	fmt.Println(strings.Join(rets, " "))
}
