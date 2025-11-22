package main

import "fmt"

func main() {
	var N string
	fmt.Scan(&N)

	var first int
	var second = -1
	var firstCount, secondCount int

	first = int(N[0] - '0')
	firstCount = 1

	for _, v := range N[1:] {
		if int(v-'0') == first {
			firstCount++
		} else {
			break
		}
	}

	// fmt.Println(first, firstCount, N[firstCount:])

	var ret int
	for _, v := range N[firstCount:] {
		if second != -1 && int(v-'0') != second {
			first = second
			firstCount = secondCount
			secondCount = 0
		}

		second = int(v - '0')
		secondCount++

		// fmt.Println(first, second, firstCount, secondCount)

		if isOdd(secondCount * 2) {
			continue
		}

		if secondCount > firstCount {
			continue
		}

		if first+1 != second {
			continue
		}

		ret++
	}

	fmt.Println(ret)
}

func isOdd(n int) bool {
	return n%2 != 0
}
