package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var Z string
	fmt.Scan(&Z)

	arrays := []int{}
	for _, v := range Z {
		arrays = append(arrays, int(v-'0'))
	}

	sort.Ints(arrays)

	for i, v := range arrays {
		if v != 0 {
			arrays[0], arrays[i] = arrays[i], arrays[0]
			break
		}
	}

	var str strings.Builder

	for _, v := range arrays {
		fmt.Fprint(&str, v)
	}

	fmt.Println(str.String())
}
