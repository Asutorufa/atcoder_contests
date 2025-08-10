package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Node struct {
	A     int
	Len   int
	Count int
	Sum   int
}

// 最后两分钟绝杀了， D都没看题目T_T
func main() {
	br := bufio.NewReader(os.Stdin)
	var N, Q int
	fmt.Fscan(br, &N, &Q)

	var array []*Node = []*Node{
		{},
	}
	var maps = make(map[int]*Node)

	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		if node, ok := maps[A]; ok {
			node.Len++
			continue
		}

		node := &Node{
			A:   A,
			Len: 1,
		}
		maps[A] = node
		array = append(array, node)
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i].A < array[j].A
	})

	sum := 0
	l := 0
	for _, v := range array {
		sum += v.Len * v.A
		l += v.Len

		v.Sum = sum
		v.Count = l
	}

	answer := make([]int, 0, Q)
	for i := 0; i < Q; i++ {
		var Q int
		fmt.Fscan(br, &Q)

		i := lastLessEqual(array, Q-1)

		one := 1
		if i == -1 {
			i = 0
			one = 0
		}

		v := array[i]

		if v.Count == N {
			answer = append(answer, -1)
		} else {
			answer = append(answer, (N-v.Count)*(Q-1)+one+v.Sum)
		}
	}

	for _, v := range answer {
		fmt.Println(v)
	}
}

/*
4 2
4 1 8 4
1
8
*/

func lastLessEqual(a []*Node, x int) int {
	// 找到第一个 > x 的位置
	idx := sort.Search(len(a), func(i int) bool {
		return a[i].A > x
	})
	if idx == 0 {
		return 0
	}
	return idx - 1 // idx-1 就是最后一个 ≤ x 的元素索引
}
