package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(br, &T)

	answers := make([]int, 0, T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(br, &N)

		arrays := make([]int, 0, N)
		for i := 0; i < N; i++ {
			var A int
			fmt.Fscan(br, &A)
			arrays = append(arrays, A)
		}

		answers = append(answers, solve(arrays))
	}

	for _, v := range answers {
		fmt.Println(v)
	}
}

func solve(arrays []int) int {
	first := arrays[0]
	last := arrays[len(arrays)-1]

	if last <= first*2 {
		return 2
	}

	sort.Slice(arrays, func(i, j int) bool {
		return arrays[i] <= arrays[j]
	})

	start := -1
	end := -1

	if start == end+1 {
		return 2
	}

	for i, v := range arrays {
		if v == first {
			start = i
		}

		if v == last && (end == -1 || i < end) {
			end = i
		}

		if start != -1 && end != -1 {
			break
		}
	}

	arrays = arrays[start : end+1]

	if len(arrays) == 2 {
		if arrays[1] > arrays[0]*2 {
			return -1
		}

		return 2
	}

	// fmt.Println(arrays)

	last = arrays[0]

	answers := [][]int{
		{last},
	}

	current := []int{}

	for _, v := range arrays[1:] {
		if v <= last*2 {
			current = append(current, v)
		} else {
			if len(current) == 0 {
				// log.Println("last", last, "v", v)
				return -1
			}

			answers = append(answers, current)

			last = current[len(current)-1]
			current = []int{}

			if v <= last*2 {
				current = append(current, v)
			} else {
				// log.Println("last", last, "v", v)
				return -1
			}
		}
	}

	answers = append(answers, current)

	// fmt.Println(answers)
	return len(answers)
}
