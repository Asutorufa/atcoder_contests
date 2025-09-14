package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	Time int
	Size int
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(br, &N, &K)

	current := K
	currentTime := 0

	array := &Next{}

	for i := 0; i < N; i++ {
		var A, B, C int
		fmt.Fscan(br, &A, &B, &C)

		for current < C {
			p := heap.Pop(array).(*pair)
			current += p.Size
			currentTime = p.Time
		}

		if currentTime > A {
			A = currentTime
		}

		fmt.Println(A)

		current -= C
		heap.Push(array, &pair{Time: A + B, Size: C})
	}
}

var _ sort.Interface = (*Next)(nil)

type Next struct {
	ps []*pair
}

func (n *Next) Len() int           { return len(n.ps) }
func (n *Next) Less(i, j int) bool { return n.ps[i].Time < n.ps[j].Time }
func (n *Next) Swap(i, j int)      { n.ps[i], n.ps[j] = n.ps[j], n.ps[i] }
func (n *Next) Push(x interface{}) { n.ps = append(n.ps, x.(*pair)) }
func (n *Next) Pop() interface{} {
	x := n.ps[len(n.ps)-1]
	n.ps = n.ps[:len(n.ps)-1]
	return x
}

/*
4 10
30 300 3
60 45 4
90 45 5
120 45 2


*/
