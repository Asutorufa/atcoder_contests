package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

type P struct {
	C, V int
}

func main() {
	N, K, M := readInt(), readInt(), readInt()

	p := make([]P, N)

	for i := range N {
		p[i] = P{readInt(), readInt()}
	}

	slices.SortFunc(p, func(a, b P) int {
		return cmp.Compare(b.V, a.V)
	})

	result := p[:K]

	maps := make(map[int]int, K)

	for _, v := range result {
		maps[v.C]++
	}

	fn := p[K:]
	nn := 0
	findNext := func() int {
		for ; ; nn++ {
			if maps[fn[nn].C] == 0 {
				return nn
			}
		}
	}

	pn := K - 1
	for i := pn; i >= 0; i-- {
		if len(maps) >= M {
			break
		}

		z := p[i]

		if maps[z.C] > 1 {
			fi := findNext()

			maps[z.C]--
			maps[fn[fi].C]++
			result[i] = fn[fi]
		}
	}

	var ans int

	for _, v := range result {
		ans += v.V
	}

	fmt.Println(ans)
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
