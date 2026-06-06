package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	H, W, K := readInt(), readInt(), readInt()
	matrix := make([][]int, H)

	for i := range H {
		matrix[i] = make([]int, 0, W)

		S := readString()

		for _, v := range S {
			switch v {
			case '1':
				matrix[i] = append(matrix[i], 1)
			default:
				matrix[i] = append(matrix[i], 0)
			}
		}
	}

	fmt.Println(numSubmatrixSumTarget(matrix, K))
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n := len(matrix)
	m := len(matrix[0])
	ans := 0
	// 如果行比列多，转置一下，降低复杂度
	if n > m {
		trans := make([][]int, m)
		for i := range trans {
			trans[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				trans[j][i] = matrix[i][j]
			}
		}
		matrix = trans
		n, m = m, n
	}
	// 枚举上边界
	for top := 0; top < n; top++ {
		colSum := make([]int, m)
		// 枚举下边界
		for bottom := top; bottom < n; bottom++ {
			// 把这一行加进去
			for col := 0; col < m; col++ {
				colSum[col] += matrix[bottom][col]
			}
			// 统计 colSum 中和为 target 的连续子数组个数
			cnt := map[int]int{0: 1}
			prefix := 0
			for _, v := range colSum {
				prefix += v
				if c, ok := cnt[prefix-target]; ok {
					ans += c
				}
				cnt[prefix]++
			}
		}
	}
	return ans
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
