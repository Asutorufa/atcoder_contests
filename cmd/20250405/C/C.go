package main

import (
	"fmt"
)

func main() {
	var N int64
	fmt.Scan(&N)

	// fmt.Println(solve(N))
	fmt.Println(solve2(N))

	return

	maps := make(map[int]struct{})

	for i := int64(1); i <= N; i++ {
		for j := int64(1); j <= N; j++ {
			z := pow(i, j)

			if z <= 1 {
				continue
			}

			if z > N {
				break
			}

			maps[int(z)] = struct{}{}
		}
	}

	fmt.Println(len(maps))
}

func pow(a, b int64) int64 {
	pre := pow2(a)
	suf := pow10(b)
	return pre * suf
}

var pow2Cache = make(map[int64]int64)

func pow2(a int64) int64 {
	if a == 0 {
		return 1
	}

	if a == 1 {
		return 2
	}

	if v, ok := pow2Cache[a]; ok {
		return v
	}

	v := pow2(a-1) * 2
	pow2Cache[a] = v
	return v
}

func pow10(a int64) int64 {
	return a * a
}

/*
定义与推导
1. “良い整数”的定义

	一个正整数 X 可以表示为 X=2a×b^2，其中 a 是非负整数，b 是正整数。
	根据 a 的奇偶性：
	    a 为偶数：
	        令 a=2k（k≥0）。
	        X=2^(2k)×b^2=(2k⋅b)^2×4=4×c^2，其中 c=2k⋅b 是正整数。
	        X 是某个平方数（c^2）的 4 倍。
	    a 为奇数：
	        令 a=2k+1（k≥0）。
	        X=2^(2k+1)×b^2 = 2×2^(2k)×b^2= 2×(2k⋅b)^2=2×d^2，其中 d=2k⋅b 是正整数。
	        X 是某个平方数（d^2）的 2 倍。
	因此，“良い整数”要么是“平方数的 2 倍”（如 2, 8, 18），要么是“平方数的 4 倍”（如 4, 16, 36）。
*/
func solve(n int64) int64 {
	// 计算 |T| = ⌊√(⌊N/2⌋)⌋
	t := isqrt(n / 2)
	// 计算 |U| = ⌊√(⌊N/4⌋)⌋
	u := isqrt(n / 4)
	return t + u
}

func isqrt(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}

/*
この問題の難しいところは，n=2^a*b^2 を満たす (a,b) が複数存在するような n があるため，重複して数えるのを避ける工夫が必要であるという点です．公式解説では

a=1,2 に限定することで重複を避けて数えています．

これとは別に，「b は奇数」という条件を加えることでも重複を避けることが可能です．

良い整数は正整数 a および正の奇数 b を用いてただ一通りに 2^a*b^2 とあらわすことが可能なことが証明できます．よって，2^a*b^2≤N を満たす正整数 a および正の奇数 b の組の個数を数えられれば良いです．a を固定すると b ≤ sqrt(N/(2^a))
​ となるため，これを満たす正の奇数 b の個数は (sqrt(N/(2^a))+1)/2 です．
この値を a=1,2,…,⌊log2​N⌋ に対して求め，それらを足し合わせることで答えを求めることができます．
*/
// 其中奇数在自然数中为等差数列，所以奇数的个数是 (n+1)/2

func solve2(n int64) int64 {
	ans := int64(0)

	for i := int64(1); i < 61; i++ {
		ans += (isqrt(n/pow2(i)) + 1) / 2
	}

	return ans
}
