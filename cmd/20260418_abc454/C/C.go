package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	N, M := readInt(), readInt()

	maps := make(map[int]map[int]bool, N)

	for range M {
		A, B := readInt(), readInt()

		if _, ok := maps[A]; !ok {
			maps[A] = make(map[int]bool)
		}

		maps[A][B] = true
	}

	cache := make(map[int]bool, N)
	cache[1] = true
	ret := 1
	queue := []int{1}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for next := range maps[p] {
			if !cache[next] {
				cache[next] = true
				ret++
				queue = append(queue, next)
			}
		}
	}

	fmt.Println(ret)
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func insert(a []int, value int) []int {
	i := sort.SearchInts(a, value)
	a = append(a, 0)
	copy(a[i+1:], a[i:])
	a[i] = value
	return a
}

func deleteLessOrEqual(a []int, value int) []int {
	i := sort.Search(len(a), func(i int) bool {
		return a[i] > value
	})

	if i <= 0 {
		return a
	}

	return a[i:]
}
