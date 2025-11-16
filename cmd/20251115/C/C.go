package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, X, Y int64
	fmt.Fscan(br, &N, &X, &Y)

	var ns []int64
	for range N {
		var n int64
		fmt.Fscan(br, &n)
		ns = append(ns, n)
	}

	fmt.Println(solve(X, Y, ns))
}

func solve(x, y int64, a []int64) int64 {
	n := len(a)
	if n == 0 {
		return 0
	}

	mina := a[0]
	for i := range n {
		mina = min(mina, a[i])
	}

	sum := int64(0)

	for i := range n {
		vs := (a[i] - mina) * y
		vt := y - x
		if vs%vt != 0 {
			return -1
		}
		minor := vs / vt
		if minor > a[i] {
			return -1
		}
		sum += a[i] - minor
	}

	return sum
}

/*
假设：

  - X = 6, Y = 8 → g = 2

  - A = [11, 10, 13]

    1.	检查模类：

    11*6 = 66 === 0 mod 2
    10*6 = 60 === 0 mod 2
    13*6 = 78 === 0 mod 2

    都是 0 mod 2 → 有可能有解

    2.	交集区间：

    11*6 ~ 11*8 = 66 ~ 88
    10*6 ~ 10*8 = 60 ~ 80
    13*6 ~ 13*8 = 78 ~ 104

    交集区间 [78, 80]

    3.	找 W ≤ 80 且 ≡ 0 mod 2 → W = 80
    4.	计算 b_i：

    b_1 = (80 - 66)/2 = 7
    b_2 = (80 - 60)/2 = 10
    b_3 = (80 - 78)/2 = 1
    都在 0~A[i] 内 → 合法
    sum b_i = 18
*/
func solve2(X, Y int64, A []int64) int64 {
	N := len(A)

	g := Y - X

	/*
		because

		 Sum(A[i]) = A[i] * X + Z*g

		{A[i]*X, A[i]*X + g, A[i]*X + 2g, ...}

		 so
		  g must mod of ervery Sum
	*/
	var modClass = ((A[0] % g) * (X % g)) % g
	for i := range N {
		if i == 0 {
			continue
		}

		if mi := ((A[i] % g) * (X % g)) % g; modClass != mi {
			return -1
		}
	}

	/*
		每个 i 的 W 可取值区间是 [A_i*X, A_i*Y]。
		所有 i 的 W 必须落在交集区间 [L, R] 里。
		如果交集为空（L > R），无解。
	*/
	var L = A[0] * X
	var R = A[0] * Y
	for i := range N {
		if i == 0 {
			continue
		}
		if lx := A[i] * X; lx > L {
			L = lx
		}
		if rx := A[i] * Y; rx < R {
			R = rx
		}
	}
	if L > R {
		return -1
	}

	delta := (R - modClass) % g
	if delta < 0 {
		delta += g
	}
	W := R - delta
	if W < L {
		return -1
	}

	var total int64 = 0
	for i := range N {
		bi := (W - A[i]*X) / g
		if bi < 0 || bi > A[i] {
			return -1
		}
		total += bi
	}

	return total
}
