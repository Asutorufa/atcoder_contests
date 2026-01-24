package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 理论直接用浮点数也行，但是精度不够过不了
// 所以用外积表示角度
// 大数的东西，早忘了

type Monster struct {
	Index int
	X, Y  int
}

func half(m Monster) int {
	if m.Y > 0 || (m.Y == 0 && m.X > 0) {
		return 0
	}
	return 1
}

func cross(a, b Monster) int64 {
	return int64(a.X)*int64(b.Y) - int64(a.Y)*int64(b.X)
}

func dot(a, b Monster) int64 {
	return int64(a.X)*int64(b.X) + int64(a.Y)*int64(b.Y)
}

func sameDir(a, b Monster) bool {
	return cross(a, b) == 0 && dot(a, b) > 0
}

func main() {
	N := readInt()
	Q := readInt()

	monsters := make([]Monster, N)
	for i := range monsters {
		x := readInt()
		y := readInt()
		monsters[i] = Monster{
			Index: i,
			X:     x,
			Y:     y,
		}
	}

	sort.Slice(monsters, func(i, j int) bool {
		hi := half(monsters[i])
		hj := half(monsters[j])
		if hi != hj {
			return hi < hj
		}
		c := cross(monsters[i], monsters[j])
		if c != 0 {
			return c > 0
		}
		return false
	})

	dirIndex := make([]int, N)
	dirCount := make([]int, 0)

	for i := 0; i < N; {
		j := i
		for j+1 < N && sameDir(monsters[i], monsters[j+1]) {
			j++
		}
		dirCount = append(dirCount, j-i+1)
		dirID := len(dirCount) - 1
		for k := i; k <= j; k++ {
			dirIndex[monsters[k].Index] = dirID
		}
		i = j + 1
	}

	M := len(dirCount)
	prefix := make([]int, M+1)
	for i := range M {
		prefix[i+1] = prefix[i] + dirCount[i]
	}

	ret := make([]int, 0, Q)
	for range Q {
		a := readInt() - 1
		b := readInt() - 1
		da := dirIndex[a]
		db := dirIndex[b]

		var ans int
		if db <= da {
			ans = prefix[da+1] - prefix[db]
		} else {
			ans = (prefix[M] - prefix[db]) + prefix[da+1]
		}
		ret = append(ret, ans)
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
