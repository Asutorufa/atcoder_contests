package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	var S, T = make([]string, 0, N), make([]string, 0, N)

	for i := 0; i < N; i++ {
		var s string
		fmt.Fscan(br, &s)
		S = append(S, s)
	}

	for i := 0; i < N; i++ {
		var t string
		fmt.Fscan(br, &t)
		T = append(T, t)
	}

	res := math.MaxInt

	for i, f := range []func([]string) []string{rotate360, rotate90, rotate180, rotate270} {
		// fmt.Println(diff(f(S), T) + i)

		res = min(res, diff(f(S), T)+i)
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(s, t []string) int {
	var res int
	for i := range s {
		for j := range s[i] {
			if s[i][j] != t[i][j] {
				res++
			}
		}
	}

	return res
}

func rotate360(s []string) []string {
	return s
}

func rotate90(s []string) []string {
	var res []string = make([]string, len(s[0]))

	for i, v := range s {
		for j := range v {
			res[j] = string(s[i][j]) + res[j]
		}
	}

	return res
}

func rotate180(s []string) []string {
	var res []string = make([]string, len(s[0]))

	for i, v := range s {
		res[len(v)-1-i] = reverse(v)
	}

	return res
}

func rotate270(s []string) []string {
	return rotate180(rotate90(s))
}

func reverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
