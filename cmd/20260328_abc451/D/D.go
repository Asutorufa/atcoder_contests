package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

var MAX = int(math.Pow(10, 9))

var parts = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912}

func main() {
	N := readInt()

	seen := make(map[int]bool)

	a := &IntHeap{}
	heap.Init(a)

	for _, p := range parts {
		seen[p] = true
		heap.Push(a, p)
	}

	count := 0
	for a.Len() > 0 {
		x := heap.Pop(a).(int)
		count++
		if count == N {
			fmt.Println(x)
			return
		}

		for _, p := range parts {
			nn := ConcatNumbers(x, p)
			if nn <= MAX {
				if !seen[nn] {
					seen[nn] = true
					heap.Push(a, nn)
				}
			} else {
				break
			}
		}
	}
}

func ConcatNumbers(a, b int) int {
	if b == 0 {
		return a * 10
	}

	multiplier := 1
	tempB := b
	for tempB > 0 {
		multiplier *= 10
		tempB /= 10
	}

	return a*multiplier + b
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
