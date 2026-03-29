package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	Q := readInt()

	a := &IntHeap{}
	heap.Init(a)

	for range Q {
		m := readInt()
		h := readInt()

		switch m {
		case 1:
			heap.Push(a, h)
		case 2:
			for a.Len() > 0 && (*a)[0] <= h {
				heap.Pop(a)
			}
		}

		fmt.Println(a.Len())
	}
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
