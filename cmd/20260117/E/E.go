package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountSubstrings(s string) int {
	n := len(s)
	p := make([]int, n+1)
	curr := 0
	for i := range n {
		val := 0
		switch s[i] {
		case 'A':
			val = 1
		case 'B':
			val = -1
		}
		curr += val
		p[i+1] = curr
	}

	return mergeSortAndCount(p, 0, len(p)-1)
}

func mergeSortAndCount(arr []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	count := mergeSortAndCount(arr, left, mid) + mergeSortAndCount(arr, mid+1, right)

	temp := make([]int, 0, right-left+1)
	i, j := left, mid+1

	ptr := left
	for j <= right {
		for ptr <= mid && arr[ptr] < arr[j] {
			ptr++
		}
		count += (ptr - left)
		j++
	}

	i, j = left, mid+1
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	temp = append(temp, arr[i:mid+1]...)
	temp = append(temp, arr[j:right+1]...)
	copy(arr[left:right+1], temp)

	return count
}

func CountSubstringsv1(s string) int {
	n := len(s)
	totalCount := 0

	for i := range n {
		balance := 0

		for j := i; j < n; j++ {
			switch s[j] {
			case 'A':
				balance++
			case 'B':
				balance--
			}
			if balance > 0 {
				totalCount++
			}
		}
	}

	return totalCount
}

func main() {
	readInt()
	S := readString()

	fmt.Println(CountSubstrings(S))
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
