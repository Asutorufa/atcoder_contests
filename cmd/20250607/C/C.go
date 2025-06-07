package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var N, L int
	fmt.Fscan(r, &N, &L)

	// 如果 L 不是 3 的倍数，直接输出 0
	if L%3 != 0 {
		fmt.Println(0)
		return
	}

	// 读取数组 d
	// 此处拿到的是两个点之间的距离， 点i+1 与点 i 之间的距离
	// 点 i+1 は点 i から時計回りに円周上を di​ 進んだ位置にあります。
	d := make([]int, N)
	for i := 0; i < N-1; i++ {
		var D int
		fmt.Fscan(r, &D)
		d[i] = D
	}

	// 前缀和模 L
	cnt := make([]int, L)
	x := 0
	for i := 0; i < N; i++ {
		if i != 0 {
			x += d[i-1]
		}
		x %= L
		cnt[x]++
	}

	// fmt.Println(cnt)

	// 三个等间距的位置组合统计
	// 等边三角形表示三个点的模等长增长
	ans := 0
	k := L / 3
	for i := 0; i < k; i++ {
		ans += cnt[i] * cnt[i+k] * cnt[i+2*k]
	}

	fmt.Println(ans)
}
