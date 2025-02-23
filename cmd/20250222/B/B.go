package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var N int
	fmt.Scanln(&N)

	type str struct {
		String string
		Length int
	}

	strs := make([]str, N)
	for i := 0; i < N; i++ {
		var s str
		fmt.Scanln(&s.String)
		s.Length = len(s.String)

		strs[i] = s
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i].Length < strs[j].Length
	})

	var ans strings.Builder
	for _, v := range strs {
		ans.WriteString(v.String)
	}

	fmt.Println(ans.String())
}
