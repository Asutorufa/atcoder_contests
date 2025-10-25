package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M, C int
	fmt.Fscan(br, &N, &M, &C)

	arrays := map[int]int{}
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)
		arrays[A] = arrays[A] + 1
		arrays[A+M] = arrays[A+M] + 1
		arrays[A+M+M] = arrays[A+M+M] + 1
	}

	keys := make([]int, 0, len(arrays))
	for k := range arrays {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	array := make([]int, 0, len(keys))
	for i := range keys {
		array = append(array, arrays[keys[i]])
	}

	for i := range array {
		if i > 0 {
			array[i] += array[i-1]
		}
	}

	sum := 0
	i := 0
	for i < M {
		index := MaxLE(keys, i+1)
		same := keys[index+1] - keys[index]
		prefix := array[index]
		sum += same * (MinGE(array, C+prefix) - prefix)
		i = keys[index+1] - 1
	}

	fmt.Println(sum)
}

func MinGE(sortedArr []int, target int) int {
	idx := sort.Search(len(sortedArr), func(i int) bool {
		return sortedArr[i] >= target
	})
	if idx == len(sortedArr) {
		return sortedArr[idx-1]
	}
	return sortedArr[idx]
}

/*
1 1000 10
1 1000
*/

func MaxLE(sortedArr []int, target int) int {
	idx := sort.Search(len(sortedArr), func(i int) bool {
		return sortedArr[i] > target
	})
	if idx == 0 {
		return 0
	}
	return idx - 1
}
