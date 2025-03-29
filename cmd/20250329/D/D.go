package main

import (
	"bufio"
	"fmt"
	"os"
)

// 令人无语的题目，原来是找两两不相邻的情侣的数量
// 还以为是找换座位的次数。。。
//
// N 組のカップルが一列に座っています。2 組のカップルであって、もともと両方のカップルは隣り合わせで座っておらず、かつ 4 人の間で席を交換することで両方のカップルが隣り合わせで座れるようになる組の個数を数えてください。
//
// ** 組の個数を数えてください **
func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(br, &T)

	for i := 0; i < T; i++ {
		solveOne2(br)
	}
}

func solveOne2(br *bufio.Reader) {
	var N int
	fmt.Fscan(br, &N)

	a := make([]int, N*2)
	pos := make(map[int]int, N*2)
	for i := range a {
		fmt.Fscan(br, &a[i])

		if _, ok := pos[a[i]]; ok {
			pos[-a[i]] = i
			a[i] = -a[i]
		} else {
			pos[a[i]] = i
		}
	}

	ans := 0
	for v := 1; v <= N; v++ {
		i := pos[v]
		j := pos[-v]

		if i+1 == j {
			continue
		}

		ii := i + 1
		if ii < 0 || ii >= 2*N {
			continue
		}

		for _, dj := range []int{-1, 1} {
			jj := j + dj

			if jj < 0 || jj >= 2*N || ii-jj == 1 || jj-ii == 1 {
				continue
			}

			// fmt.Println("ssss", i, j, ii, jj, a[ii], a[jj])
			if a[ii] == -a[jj] {
				// fmt.Println("ccc", i, j, ii, jj)
				ans++
			}
		}
	}

	fmt.Println(ans)
}
