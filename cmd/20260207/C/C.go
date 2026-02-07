package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	N := readInt()

	minium := 0
	sum := 0
	array := make([]int, 0, N)
	for range N {
		A := readInt()

		array = append(array, A)

		sum += A
		if A > minium {
			minium = A
		}
	}

	sort.Ints(array)

	all := map[int]bool{}

	M := array[N-1]

	all[M] = true

	for i := 0; i < N-1; i++ {
		all[M+array[i]] = true
	}

	for _, L := range array {
		all[L] = true
	}

	ans := []int{}

	for L := range all {
		if sum%L != 0 {
			continue
		}

		if check(array, L) {
			ans = append(ans, L)
		}
	}

	sort.Ints(ans)

	for _, v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func check(A []int, L int) bool {
	l, r := 0, len(A)-1
	for l <= r {
		if A[l]+A[r] == L {
			l++
		} else if A[r] != L {
			return false
		}
		r--
	}
	return true
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
