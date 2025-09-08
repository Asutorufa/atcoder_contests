package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	var ps [][2]int
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Fscan(br, &a, &b)
		ps = append(ps, [2]int{a, b})
	}

	/*
		ax+by+c=0

		ax1+by1+c=ax2+by2+c
			==> ax1-ax2=by2-by1
			==> a(x1-x2)=b(y2-y1)
			=> a/b = (y2-y1)/(x1-x2)

		由上可得
			==> a = y2 - y1
			==》b = x1 - x2
		再将a,b 带入最开始的方程ax1+by1+c=0，可得c
		 (y2-y1)x1 + (x1-x2)y1 = -c
		 y2x1 - y1x1 - x2y1 + x1y1 = -c
		 y2x1 - x2y1 = -c
		 x2y1 - y2x1 = c

		或
			a = y1 - y2
			b = x2 - x1
		 (y1-y2)x1 + (x2-x1)y1 = c
		 y1x1 - y2x1 + x2y1 - x1y1 = -c
		 x2y1 - y2x1 = -c
		 y2x1 - y1x2 = c
	*/

	for i := 0; i < 100; i++ {
		var i, j int
		for i == j {
			i = rand.Intn(N)
			j = rand.Intn(N)
		}

		x1, y1 := ps[i][0], ps[i][1]
		x2, y2 := ps[j][0], ps[j][1]

		a := y2 - y1
		b := x1 - x2
		c := x2*y1 - y2*x1

		num := 0
		for _, v := range ps {
			if a*v[0]+b*v[1]+c == 0 {
				num++
			}
		}

		if num*2 > N {
			fmt.Println("Yes")
			fmt.Println(a, b, c)
			return
		}
	}

	fmt.Println("No")
}

/*
1️⃣ RANSAC 的基本原理

RANSAC 是一种 随机抽样一致性算法，用于在含噪数据中拟合模型，核心思想：
数据中大多数点符合某种模型（内点 inlier）
剩余点可能是噪声或异常点（外点 outlier）
随机选最少点子集拟合模型，然后统计有多少点落在这个模型上
重复多次，选内点最多的模型作为最终结果

特点：

可以容忍大量异常点
不需要对所有点都拟合
成功概率随随机次数增加而接近 100%

2️⃣ 在求多数点共线问题中的应用

场景：有 n 个点，保证存在一条直线穿过 超过一半点。
目标：找到这条直线。

RANSAC 思路：

每次随机选 两点（直线最少需要两点确定）

计算这条直线方程 ax+by+c=0

遍历所有点，统计落在直线上的点数（或容差内）
如果点数 > n/2，直接返回这条直线
如果没有，重复 k 次，直到找到

3️⃣ 为什么只需要随机两点
假设多数派直线有 M 个点（M > n/2）
随机选两点，其中都落在多数派直线的概率： p=(2n)(2M)
	​
当 M > n/2 时，概率p>1/4（比如 n=100, M=60, p ≈ 0.357）

多次尝试（k 次），失败概率=(1−p)k只要 k 足够大，几乎必然选到正确直线
*/
