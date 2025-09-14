package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, R int
	fmt.Fscan(in, &N, &R)

	L := make([]int, N+1) // doors indexed 1..N
	for i := 1; i <= N; i++ {
		fmt.Fscan(in, &L[i])
	}

	// 收集 0 门的信息
	left0, right0 := -1, -1
	zeros := 0
	for i := 1; i <= N; i++ {
		if L[i] == 0 {
			if left0 == -1 {
				left0 = i
			}
			right0 = i
			zeros++
		}
	}

	if zeros == 0 {
		fmt.Println(0)
		return
	}

	// 计算房间区间 [x,y]
	x := R
	if left0 < x {
		x = left0
	}
	y := R
	if right0-1 > y {
		y = right0 - 1
	}

	// 统计区间内部原本为 1 的门的数量（门索引 j = x+1 .. y）
	ones := 0
	for j := x + 1; j <= y; j++ {
		if j >= 1 && j <= N && L[j] == 1 {
			ones++
		}
	}

	ans := zeros + 2*ones
	fmt.Println(ans)
}

/*
8 2
0 1 1 0 1 0 1 1

9

8 2
1 1 1 0 1 0 1 1

6


8 0
1 1 1 0 1 0 1 1

10


8 7
1 1 1 0 1 0 0 0

6


8 8
1 1 1 0 1 0 0 0

6


8 8
1 1 1 0 1 0 0 1

7


*/
