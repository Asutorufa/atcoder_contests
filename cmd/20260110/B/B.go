package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	N := readInt()

	type T struct {
		S  int
		ID int
	}
	Ts := make([]T, 0, N)
	for i := range N {
		Ts = append(Ts, T{readInt(), i + 1})
	}

	sort.Slice(Ts, func(i, j int) bool {
		return Ts[i].S < Ts[j].S
	})

	for i := range 3 {
		fmt.Print(Ts[i].ID, " ")
	}
	fmt.Println()
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
