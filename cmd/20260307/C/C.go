package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Ball struct {
	Index int
	Value int
}

func main() {
	N := readInt()
	Q := readInt()

	balls := make([]Ball, 0, N)
	for i := range N {
		balls = append(balls, Ball{Index: i, Value: readInt()})
	}

	sort.Slice(balls, func(i, j int) bool {
		return balls[i].Value < balls[j].Value
	})

	query := map[int]bool{}

	ret := []int{}
	for range Q {
		K := readInt()
		clear(query)

		for range K {
			B := readInt() - 1
			query[B] = true
		}

		for _, v := range balls {
			if !query[v.Index] {
				ret = append(ret, v.Value)
				break
			}
		}
	}

	for _, v := range ret {
		fmt.Println(v)
	}
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
